// Package translator содержит реализацию транслятора в Z3
package translator

import (
	"fmt"
	"strconv"
	"symbolic-execution-course/internal/symbolic"

	"github.com/ebukreev/go-z3/z3"
)

// Z3Translator транслирует символьные выражения в Z3 формулы
type Z3Translator struct {
	ctx    *z3.Context
	slv    *z3.Solver
	config *z3.Config

	// Var cache
	vars map[string]z3.Value

	// Sorts
	stringSort z3.Sort
	bvSort     z3.Sort
}

// NewZ3Translator создаёт новый экземпляр Z3 транслятора
func NewZ3Translator() *Z3Translator {
	config := &z3.Config{}
	ctx := z3.NewContext(config)
	slv := z3.NewSolver(ctx)

	return &Z3Translator{
		ctx:        ctx,
		slv:        slv,
		config:     config,
		vars:       make(map[string]z3.Value),
		stringSort: ctx.UninterpretedSort("String"),
		bvSort:     ctx.BVSort(64),
	}
}

func (zt *Z3Translator) Reset() {
	zt.config = &z3.Config{}
	zt.ctx = z3.NewContext(zt.config)
	zt.slv = z3.NewSolver(zt.ctx)

	zt.vars = make(map[string]z3.Value)
	zt.stringSort = zt.ctx.UninterpretedSort("String")
}

func (zt *Z3Translator) Assert(expr interface{}) {
	switch v := expr.(type) {
	case z3.Bool:
		zt.slv.Assert(v)
	case z3.Value:
		if b, ok := v.(z3.Bool); ok {
			zt.slv.Assert(b)
			return
		}
		panic(fmt.Sprintf("AssertExpression: non-boolean z3.Value (%T) passed", v))
	default:
		panic(fmt.Sprintf("AssertExpression: unsupported type %T", expr))
	}
}

func (zt *Z3Translator) AssertExpr(e symbolic.SymbolicExpression) {
	v, err := zt.TranslateExpression(e)
	if err != nil {
		panic("AssertExpr: translate error: " + err.Error())
	}
	zt.Assert(v)
}

func (zt *Z3Translator) SolverToString() string {
	return zt.slv.String()
}

func (zt *Z3Translator) IsSatisfiable() (bool, error) {
	return zt.slv.Check()
}

func (zt *Z3Translator) Model() *z3.Model {
	return zt.slv.Model()
}

// GetIntValue получает значение целочисленной переменной из модели
func (zt *Z3Translator) GetIntValue(m *z3.Model, variable *symbolic.SymbolicVariable) (int64, error) {
	val := m.Eval(zt.vars[variable.Name], false)
	if val == nil {
		return 0, fmt.Errorf("variable %s not in model", variable.Name)
	}

	return strconv.ParseInt(val.String(), 10, 64)
}

// GetBoolValue получает значение булевой переменной из модели
func (zt *Z3Translator) GetBoolValue(model *z3.Model, variable z3.Bool) (bool, error) {
	value := model.Eval(variable, false)
	if value == nil {
		return false, fmt.Errorf("variable not found in model")
	}

	// Проверяем строковое представление
	str := value.String()
	switch str {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, fmt.Errorf("unexpected boolean value: %s", str)
	}
}

// GetContext возвращает Z3 контекст
func (zt *Z3Translator) GetContext() interface{} {
	return zt.ctx
}

// Close освобождает ресурсы
func (zt *Z3Translator) Close() {
	// Z3 контекст закрывается автоматически
}

// TranslateExpression транслирует символьное выражение в Z3
func (zt *Z3Translator) TranslateExpression(expr symbolic.SymbolicExpression) (interface{}, error) {
	return expr.Accept(zt), nil
}

// VisitVariable транслирует символьную переменную в Z3
func (zt *Z3Translator) VisitVariable(expr *symbolic.SymbolicVariable) interface{} {
	if z3CacheVar, ok := zt.vars[expr.Name]; ok {
		return z3CacheVar
	}

	// Нету в кэше
	z3Var := zt.createZ3Variable(expr.Name, expr.ExprType)
	zt.vars[expr.Name] = z3Var
	return z3Var
}

// VisitIntConstant транслирует целочисленную константу в Z3
func (zt *Z3Translator) VisitIntConstant(expr *symbolic.IntConstant) interface{} {
	return zt.ctx.FromInt(expr.Value, zt.bvSort)
}

// VisitBoolConstant транслирует булеву константу в Z3
func (zt *Z3Translator) VisitBoolConstant(expr *symbolic.BoolConstant) interface{} {
	return zt.ctx.FromBool(expr.Value)
}

// VisitBinaryOperation транслирует бинарную операцию в Z3
func (zt *Z3Translator) VisitBinaryOperation(expr *symbolic.BinaryOperation) interface{} {
	leftZ3, err := zt.TranslateExpression(expr.Left)
	if err != nil {
		panic("VisitBinaryOperation: Unable to translate left operand.")
	}

	rightZ3, err := zt.TranslateExpression(expr.Right)
	if err != nil {
		panic("VisitBinaryOperation: Unable to translate right operand.")
	}

	if leftZ3i, ok := leftZ3.(z3.BV); ok {
		rightZ3i, ok := rightZ3.(z3.BV)
		if !ok {
			panic("VisitBinaryOperation: both operands must be BV")
		}

		switch expr.Operator {
		// Арифметические операторы
		case symbolic.ADD:
			return leftZ3i.Add(rightZ3i)
		case symbolic.SUB:
			return leftZ3i.Sub(rightZ3i)
		case symbolic.MUL:
			return leftZ3i.Mul(rightZ3i)
		case symbolic.DIV:
			return leftZ3i.SDiv(rightZ3i)
		case symbolic.REM:
			return leftZ3i.SRem(rightZ3i)

		// Логические операторы
		case symbolic.AND:
			return leftZ3i.And(rightZ3i)
		case symbolic.OR:
			return leftZ3i.Or(rightZ3i)
		case symbolic.XOR:
			return leftZ3i.Xor(rightZ3i)
		case symbolic.SHL:
			return leftZ3i.Lsh(rightZ3i)
		case symbolic.SHR:
			return leftZ3i.SRsh(rightZ3i)
		case symbolic.AND_NOT:
			// x &^ y в Go = x & (^y)
			return leftZ3i.And(rightZ3i.Not())

		// Операторы сравнения
		case symbolic.EQ:
			return leftZ3i.Eq(rightZ3i)
		case symbolic.NE:
			return leftZ3i.NE(rightZ3i)
		case symbolic.LT:
			return leftZ3i.SLT(rightZ3i)
		case symbolic.LE:
			return leftZ3i.SLE(rightZ3i)
		case symbolic.GT:
			return leftZ3i.SGT(rightZ3i)
		case symbolic.GE:
			return leftZ3i.SGE(rightZ3i)

		default:
			panic("Not impl.")
		}
	}

	if leftZ3b, ok := leftZ3.(z3.Bool); ok {
		rightZ3b, ok := rightZ3.(z3.Bool)
		if !ok {
			panic("VisitBinaryOperation: both operands must be Bool")
		}

		switch expr.Operator {
		// Логические операторы
		case symbolic.AND:
			return leftZ3b.And(rightZ3b)
		case symbolic.OR:
			return leftZ3b.Or(rightZ3b)
		case symbolic.XOR:
			return leftZ3b.Xor(rightZ3b)
		case symbolic.AND_NOT:
			return leftZ3b.And(rightZ3b).Not()
		case symbolic.IMPLIES:
			return leftZ3b.Implies(rightZ3b)

		// Операторы сравнения
		case symbolic.EQ:
			return leftZ3b.Eq(rightZ3b)
		case symbolic.NE:
			return leftZ3b.NE(rightZ3b)

		default:
			panic("Not impl.")
		}
	}

	if expr.Left.Type() == symbolic.StringType && expr.Right.Type() == symbolic.StringType {
		lu, ok := leftZ3.(z3.Uninterpreted)
		if !ok {
			panic("VisitBinaryOperation: left string is not z3.Uninterpreted")
		}
		ru, ok := rightZ3.(z3.Uninterpreted)
		if !ok {
			panic("VisitBinaryOperation: right string is not z3.Uninterpreted")
		}

		switch expr.Operator {
		case symbolic.EQ:
			return lu.Eq(ru)
		case symbolic.NE:
			return lu.NE(ru)
		default:
			panic("VisitBinaryOperation: unsupported operator for string")
		}
	}

	if leftZ3a, ok := leftZ3.(z3.Array); ok {
		rightZ3a, ok := rightZ3.(z3.Array)
		if !ok {
			panic("VisitBinaryOperation: both operands must be Array")
		}

		switch expr.Operator {
		// Операторы сравнения
		case symbolic.EQ:
			return leftZ3a.Eq(rightZ3a)
		case symbolic.NE:
			return leftZ3a.NE(rightZ3a)

		default:
			panic("Not impl.")
		}
	}

	panic("Not impl type.")
}

// VisitUnaryOperation транслирует бинарную операцию в Z3
func (zt *Z3Translator) VisitUnaryOperation(expr *symbolic.UnaryOperation) interface{} {
	operandZ3, err := zt.TranslateExpression(expr.Operand)
	if err != nil {
		panic("VisitUnaryOperation: Unable to translate operand.")
	}

	if operandZ3i, ok := operandZ3.(z3.BV); ok {
		switch expr.Operator {
		case symbolic.NEG:
			return operandZ3i.Neg()
		case symbolic.UXOR:
			return operandZ3i.Not()
		default:
			panic("Not impl.")
		}
	}

	if operandZ3b, ok := operandZ3.(z3.Bool); ok {
		switch expr.Operator {
		case symbolic.NOT:
			return operandZ3b.Not()
		default:
			panic("Not impl.")
		}
	}

	panic("Not impl type.")
}

func (zt *Z3Translator) VisitConditionalExpression(expr *symbolic.ConditionalExpression) interface{} {
	ift, err := zt.TranslateExpression(expr.IfOperand)
	if err != nil {
		panic("VisitConditionalExpression: translation IfOperand error")
	}
	thent, err := zt.TranslateExpression(expr.ThenOperand)
	if err != nil {
		panic("VisitConditionalExpression: translation ThenOperand error")
	}
	elset, err := zt.TranslateExpression(expr.ElseOperand)
	if err != nil {
		panic("VisitConditionalExpression: translation ElseOperand error")
	}

	ifb, ok := ift.(z3.Bool)
	if !ok {
		panic("VisitConditionalExpression: expected boolean IfOperand operand")
	}

	if theni, ok := thent.(z3.Int); ok {
		elsei, ok := elset.(z3.Int)
		if !ok {
			panic("VisitConditionalExpression: then/else must both be Int or both be Bool")
		}
		return ifb.IfThenElse(theni, elsei)
	}
	if thenb, ok := thent.(z3.Bool); ok {
		elseb, ok := elset.(z3.Bool)
		if !ok {
			panic("VisitConditionalExpression: then/else must both be Int or both be Bool")
		}
		return ifb.IfThenElse(thenb, elseb)
	}
	panic("VisitConditionalExpression: unsupported branch type (only Int/Bool supported)")
}

func (zt *Z3Translator) VisitStringConstant(c *symbolic.StringConstant) interface{} {
	return zt.ctx.Const("str_"+c.Value, zt.stringSort)
}

// Вспомогательные методы

// createZ3Variable создаёт Z3 переменную соответствующего типа
func (zt *Z3Translator) createZ3Variable(name string, exprType symbolic.ExpressionType) z3.Value {
	switch exprType {
	case symbolic.IntType:
		return zt.ctx.Const(name, zt.bvSort)
	case symbolic.BoolType:
		return zt.ctx.BoolConst(name)
	case symbolic.StringType:
		return zt.ctx.Const(name, zt.stringSort)
	case symbolic.ArrayType:
		arrSort := zt.ctx.ArraySort(zt.bvSort, zt.bvSort)
		return zt.ctx.Const(name, arrSort)
	// TODO Возможно добавить BaseType в SymExpr и также VectorType
	default:
		panic("Not impl.")
	}
}
