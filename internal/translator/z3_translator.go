// Package translator содержит реализацию транслятора в Z3
package translator

import (
	"fmt"
	"math/big"
	"strconv"
	"symbolic-execution-course/internal/symbolic"

	"github.com/ebukreev/go-z3/z3"
)

// Z3Translator транслирует символьные выражения в Z3 формулы
type Z3Translator struct {
	ctx    *z3.Context
	slv    *z3.Solver
	config *z3.Config
	vars   map[string]z3.Value // Кэш переменных
}

// NewZ3Translator создаёт новый экземпляр Z3 транслятора
func NewZ3Translator() *Z3Translator {
	config := &z3.Config{}
	ctx := z3.NewContext(config)
	slv := z3.NewSolver(ctx)

	return &Z3Translator{
		ctx:    ctx,
		slv:    slv,
		config: config,
		vars:   make(map[string]z3.Value),
	}
}

// Reset сбрасывает состояние транслятора
func (zt *Z3Translator) Reset() {
	zt.config = &z3.Config{}
	zt.ctx = z3.NewContext(zt.config)
	zt.slv = z3.NewSolver(zt.ctx)

	zt.vars = make(map[string]z3.Value)
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
	return zt.ctx.FromBigInt(big.NewInt(expr.Value), zt.ctx.IntSort())
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

	if leftZ3i, ok := leftZ3.(z3.Int); ok {
		rightZ3i, ok := rightZ3.(z3.Int)
		if !ok {
			panic("VisitBinaryOperation: both operands must be Int")
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
			return leftZ3i.Div(rightZ3i)
		case symbolic.MOD:
			return leftZ3i.Mod(rightZ3i)

		// Операторы сравнения
		case symbolic.EQ:
			return leftZ3i.Eq(rightZ3i)
		case symbolic.NE:
			return leftZ3i.NE(rightZ3i)
		case symbolic.LT:
			return leftZ3i.LT(rightZ3i)
		case symbolic.LE:
			return leftZ3i.LE(rightZ3i)
		case symbolic.GT:
			return leftZ3i.GT(rightZ3i)
		case symbolic.GE:
			return leftZ3i.GE(rightZ3i)

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
		// Операторы сравнения
		case symbolic.EQ:
			return leftZ3b.Eq(rightZ3b)
		case symbolic.NE:
			return leftZ3b.NE(rightZ3b)

		default:
			panic("Not impl.")
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

// VisitLogicalOperation транслирует логическую операцию в Z3
func (zt *Z3Translator) VisitLogicalOperation(expr *symbolic.LogicalOperation) interface{} {
	ops := make([]z3.Bool, 0, len(expr.Operands))
	for _, opnd := range expr.Operands {
		v, err := zt.TranslateExpression(opnd)
		if err != nil {
			panic("VisitLogicalOperation: translation error")
		}
		b, ok := v.(z3.Bool)
		if !ok {
			panic("VisitLogicalOperation: expected boolean operand")
		}
		ops = append(ops, b)
	}

	switch expr.Operator {
	case symbolic.AND:
		return ops[0].And(ops[1:]...)
	case symbolic.OR:
		return ops[0].Or(ops[1:]...)
	case symbolic.NOT:
		if len(ops) != 1 {
			panic("NOT expects 1 operand")
		}
		return ops[0].Not()
	case symbolic.IMPLIES:
		if len(ops) != 2 {
			panic("IMPLIES expects 2 operands")
		}
		return ops[0].Implies(ops[1])

	default:
		panic("Not impl.")
	}
}

func (zt *Z3Translator) VisitFunctionCall(expr *symbolic.FunctionCall) interface{} {
	// 1) Переводим аргументы в Z3
	argVals := make([]z3.Value, len(expr.Args))
	for i, a := range expr.Args {
		v, err := zt.TranslateExpression(a)
		if err != nil {
			panic(fmt.Sprintf("VisitFunctionCall(%s): translate arg %d: %v", expr.Name, i, err))
		}
		switch vv := v.(type) {
		case z3.Int:
			argVals[i] = z3.Value(vv)
		case z3.Bool:
			argVals[i] = z3.Value(vv)
		case z3.Value:
			argVals[i] = vv
		default:
			panic(fmt.Sprintf("VisitFunctionCall(%s): unsupported arg %d type %T", expr.Name, i, v))
		}
	}

	// 2) Формируем доменные сорта из типов аргументов
	domain := make([]z3.Sort, len(expr.Args))
	for i, a := range expr.Args {
		switch a.Type() {
		case symbolic.IntType:
			domain[i] = zt.ctx.IntSort()
		case symbolic.BoolType:
			domain[i] = zt.ctx.BoolSort()
		default:
			panic(fmt.Sprintf("VisitFunctionCall(%s): unsupported arg sort #%d (%v)", expr.Name, i, a.Type()))
		}
	}

	// 3) Результирующий сорт по RetType
	var rng z3.Sort
	switch expr.RetType {
	case symbolic.IntType:
		rng = zt.ctx.IntSort()
	case symbolic.BoolType:
		rng = zt.ctx.BoolSort()
	default:
		panic(fmt.Sprintf("VisitFunctionCall(%s): unsupported return type %v", expr.Name, expr.RetType))
	}

	// 4) Объявляем uninterpreted function и делаем аппликацию
	fd := zt.ctx.FuncDecl(expr.Name, domain, rng)

	// Apply creates a Value representing the result of applying f to args
	return fd.Apply(argVals...)
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
}

// Вспомогательные методы

// createZ3Variable создаёт Z3 переменную соответствующего типа
func (zt *Z3Translator) createZ3Variable(name string, exprType symbolic.ExpressionType) z3.Value {
	switch exprType {
	case symbolic.IntType:
		return zt.ctx.IntConst(name)
	case symbolic.BoolType:
		return zt.ctx.BoolConst(name)
	default:
		panic("Not impl.")
	}
}
