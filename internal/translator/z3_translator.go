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
	// TODO: Реализовать
	// Создать Z3 константу с помощью zt.ctx.FromBigInt или аналогичного метода

	panic("не реализовано")
}

// VisitBoolConstant транслирует булеву константу в Z3
func (zt *Z3Translator) VisitBoolConstant(expr *symbolic.BoolConstant) interface{} {
	// TODO: Реализовать
	// Использовать zt.ctx.FromBool для создания Z3 булевой константы

	panic("не реализовано")
}

// VisitBinaryOperation транслирует бинарную операцию в Z3
func (zt *Z3Translator) VisitBinaryOperation(expr *symbolic.BinaryOperation) interface{} {
	// TODO: Реализовать
	// 1. Транслировать левый и правый операнды
	// 2. В зависимости от оператора создать соответствующую Z3 операцию

	// Подсказки по операциям в Z3:
	// - Арифметические: left.Add(right), left.Sub(right), left.Mul(right), left.Div(right)
	// - Сравнения: left.Eq(right), left.LT(right), left.LE(right), etc.
	// - Приводите типы: left.(z3.Int), right.(z3.Int) для int операций

	panic("не реализовано")
}

// VisitLogicalOperation транслирует логическую операцию в Z3
func (zt *Z3Translator) VisitLogicalOperation(expr *symbolic.LogicalOperation) interface{} {
	// TODO: Реализовать
	// 1. Транслировать все операнды
	// 2. Применить соответствующую логическую операцию

	// Подсказки:
	// - AND: zt.ctx.And(operands...)
	// - OR: zt.ctx.Or(operands...)
	// - NOT: operand.Not() (для единственного операнда)
	// - IMPLIES: antecedent.Implies(consequent)

	panic("не реализовано")
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
