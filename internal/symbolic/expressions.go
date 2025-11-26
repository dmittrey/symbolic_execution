// Package symbolic содержит конкретные реализации символьных выражений
package symbolic

import (
	"fmt"
	"strings"
)

// SymbolicExpression - базовый интерфейс для всех символьных выражений
type SymbolicExpression interface {
	// Type возвращает тип выражения
	Type() ExpressionType

	// String возвращает строковое представление выражения
	String() string

	// Accept принимает visitor для обхода дерева выражений
	Accept(visitor Visitor) interface{}
}

// SymbolicVariable представляет символьную переменную
type SymbolicVariable struct {
	Name     string
	ExprType ExpressionType
}

// NewSymbolicVariable создаёт новую символьную переменную
func NewSymbolicVariable(name string, exprType ExpressionType) *SymbolicVariable {
	return &SymbolicVariable{
		Name:     name,
		ExprType: exprType,
	}
}

// Type возвращает тип переменной
func (sv *SymbolicVariable) Type() ExpressionType {
	return sv.ExprType
}

// String возвращает строковое представление переменной
func (sv *SymbolicVariable) String() string {
	return sv.Name
}

// Accept реализует Visitor pattern
func (sv *SymbolicVariable) Accept(visitor Visitor) interface{} {
	return visitor.VisitVariable(sv)
}

// IntConstant представляет целочисленную константу
type IntConstant struct {
	Value int64
}

// NewIntConstant создаёт новую целочисленную константу
func NewIntConstant(value int64) *IntConstant {
	return &IntConstant{Value: value}
}

// Type возвращает тип константы
func (ic *IntConstant) Type() ExpressionType {
	return IntType
}

// String возвращает строковое представление константы
func (ic *IntConstant) String() string {
	return fmt.Sprintf("%d", ic.Value)
}

// Accept реализует Visitor pattern
func (ic *IntConstant) Accept(visitor Visitor) interface{} {
	return visitor.VisitIntConstant(ic)
}

// BoolConstant представляет булеву константу
type BoolConstant struct {
	Value bool
}

// NewBoolConstant создаёт новую булеву константу
func NewBoolConstant(value bool) *BoolConstant {
	return &BoolConstant{Value: value}
}

// Type возвращает тип константы
func (bc *BoolConstant) Type() ExpressionType {
	return BoolType
}

// String возвращает строковое представление константы
func (bc *BoolConstant) String() string {
	return fmt.Sprintf("%t", bc.Value)
}

// Accept реализует Visitor pattern
func (bc *BoolConstant) Accept(visitor Visitor) interface{} {
	return visitor.VisitBoolConstant(bc)
}

// BinaryOperation представляет бинарную операцию
type BinaryOperation struct {
	Left     SymbolicExpression
	Right    SymbolicExpression
	Operator BinaryOperator
}

// NewBinaryOperation создаёт новую бинарную операцию
func NewBinaryOperation(left, right SymbolicExpression, op BinaryOperator) *BinaryOperation {
	lt, rt := left.Type(), right.Type()

	switch {
	case lt == IntType && rt == IntType:
		return &BinaryOperation{Left: left, Right: right, Operator: op}

	case lt == BoolType && rt == BoolType && (op == EQ || op == NE):
		return &BinaryOperation{Left: left, Right: right, Operator: op}

	case lt == StringType && rt == StringType && (op == EQ || op == NE):
		return &BinaryOperation{Left: left, Right: right, Operator: op}

	case lt == ArrayType && rt == ArrayType && (op == EQ || op == NE):
		return &BinaryOperation{Left: left, Right: right, Operator: op}

	default:
		panic(fmt.Errorf("binary op %v not supported for %v and %v", op, lt, rt))
	}
}

func (bo *BinaryOperation) Type() ExpressionType {
	switch bo.Operator {
	case EQ, NE, LT, LE, GT, GE:
		return BoolType
	case ADD, SUB, MUL, DIV, MOD:
		return IntType
	default:
		panic("unknown operator")
	}
}

// String возвращает строковое представление операции
func (bo *BinaryOperation) String() string {
	return fmt.Sprintf("(%s %s %s)",
		bo.Left.String(),
		bo.Operator.String(),
		bo.Right.String(),
	)
}

// Accept реализует Visitor pattern
func (bo *BinaryOperation) Accept(visitor Visitor) interface{} {
	return visitor.VisitBinaryOperation(bo)
}

// LogicalOperation представляет логическую операцию
type LogicalOperation struct {
	Operands []SymbolicExpression
	Operator LogicalOperator
}

func NewLogicalOperation(operands []SymbolicExpression, op LogicalOperator) *LogicalOperation {
	for _, o := range operands {
		if o.Type() != BoolType {
			panic("logical operation expects bool operands")
		}
	}
	switch op {
	case NOT:
		if len(operands) != 1 {
			panic("NOT expects 1 operand")
		}
	case IMPLIES:
		if len(operands) != 2 {
			panic("IMPLIES expects 2 operands")
		}
	case AND, OR:
		if len(operands) < 2 {
			panic("AND/OR expect at least 2 operands")
		}
	default:
		panic("unknown logical operator")
	}
	return &LogicalOperation{Operands: operands, Operator: op}
}

// Type возвращает тип логической операции (всегда bool)
func (lo *LogicalOperation) Type() ExpressionType {
	return BoolType
}

// String возвращает строковое представление логической операции
func (lo *LogicalOperation) String() string {
	switch lo.Operator {
	case NOT:
		return fmt.Sprintf("!%s", lo.Operands[0].String())
	case IMPLIES:
		return fmt.Sprintf("(%s => %s)", lo.Operands[0].String(), lo.Operands[1].String())
	default:
		if len(lo.Operands) < 2 {
			panic("AND/OR expect at least 2 operand")
		}

		parts := make([]string, len(lo.Operands))
		for i, op := range lo.Operands {
			parts[i] = op.String()
		}
		sep := " " + lo.Operator.String() + " " // " && " или " || "
		return "(" + strings.Join(parts, sep) + ")"
	}
}

// Accept реализует Visitor pattern
func (lo *LogicalOperation) Accept(visitor Visitor) interface{} {
	return visitor.VisitLogicalOperation(lo)
}

// - FunctionCall (вызовы функций: f(x, y))
type FunctionCall struct {
	Name    string
	Args    []SymbolicExpression
	RetType ExpressionType
}

func NewFunctionCall(name string, args []SymbolicExpression, retType ExpressionType) *FunctionCall {
	return &FunctionCall{
		Name:    name,
		Args:    args,
		RetType: retType,
	}
}

func (fc *FunctionCall) Type() ExpressionType {
	return fc.RetType
}

func (fc *FunctionCall) String() string {
	parts := make([]string, len(fc.Args))
	for i, a := range fc.Args {
		parts[i] = a.String()
	}
	return "(" + "fn " + fc.Name + ": (" + strings.Join(parts, " -> ") + " -> " + fc.RetType.String() + "))"
}

func (fc *FunctionCall) Accept(visitor Visitor) interface{} {
	return visitor.VisitFunctionCall(fc)
}

// ConditionalExpression (тернарный оператор: condition ? true_expr : false_expr)
type ConditionalExpression struct {
	IfOperand   SymbolicExpression
	ThenOperand SymbolicExpression
	ElseOperand SymbolicExpression
}

func NewConditionalExpression(ifOperand SymbolicExpression, thenOperand SymbolicExpression, elseOperand SymbolicExpression) *ConditionalExpression {
	return &ConditionalExpression{
		IfOperand:   ifOperand,
		ThenOperand: thenOperand,
		ElseOperand: elseOperand,
	}
}

func (ce *ConditionalExpression) Type() ExpressionType {
	if ce.IfOperand.Type() != BoolType {
		panic("condition of ?: must be bool")
	}
	if ce.ThenOperand.Type() != ce.ElseOperand.Type() {
		panic("then/else types must match")
	}
	return ce.ThenOperand.Type()
}

func (ce *ConditionalExpression) String() string {
	return fmt.Sprintf("(%s ? %s : %s)",
		ce.IfOperand.String(), ce.ThenOperand.String(), ce.ElseOperand.String())
}

func (ce *ConditionalExpression) Accept(visitor Visitor) interface{} {
	return visitor.VisitConditionalExpression(ce)
}

type Array struct {
	Name    string
	ValType ExpressionType
	Stores  []struct {
		Idx SymbolicExpression
		Val SymbolicExpression
	}
}

func NewArray(name string, valType ExpressionType) *Array {
	if valType != IntType && valType != BoolType && valType != StringType {
		panic(fmt.Sprintf("Unexpected type %s", valType.String()))
	}

	return &Array{
		Name:    name,
		ValType: valType,
	}
}

func NewArrayStore(array *Array, idx SymbolicExpression, val SymbolicExpression) *Array {
	if val.Type() != array.ValType {
		panic(fmt.Sprintf("Array value must be %s", array.ValType.String()))
	}

	// Делаем ПЕРСИСТЕНТНУЮ версию: новый Array с копией старых stores + новый store
	newStores := make([]struct {
		Idx SymbolicExpression
		Val SymbolicExpression
	}, len(array.Stores)+1)

	copy(newStores, array.Stores)

	newStores[len(array.Stores)] = struct {
		Idx SymbolicExpression
		Val SymbolicExpression
	}{
		Idx: idx,
		Val: val,
	}

	return &Array{
		Name:    array.Name,
		ValType: array.ValType,
		Stores:  newStores,
	}
}

func (expr *Array) Type() ExpressionType {
	return ArrayType
}

func (expr *Array) String() string {
	return expr.Name + "[" + "]"
}

func (expr *Array) Accept(visitor Visitor) interface{} {
	return visitor.VisitArray(expr)
}

// ArrayAccess (доступ к элементам массива: arr[index])
type ArrayAccess struct {
	Arr *Array
	Idx SymbolicExpression
}

func NewArrayAccess(array *Array, idx SymbolicExpression) *ArrayAccess {
	if idx.Type() != IntType {
		panic("Array index must be Int")
	}
	return &ArrayAccess{Arr: array, Idx: idx}
}

func (expr *ArrayAccess) Type() ExpressionType {
	switch expr.Arr.ValType {
	case IntType, BoolType, StringType:
		return expr.Arr.ValType
	default:
		panic("unsupported array element type")
	}
}

func (expr *ArrayAccess) String() string {
	return expr.Arr.Name + "[" + expr.Idx.String() + "]"
}

func (expr *ArrayAccess) Accept(visitor Visitor) interface{} {
	return visitor.VisitArrayAccess(expr)
}

// Операторы для бинарных выражений
type BinaryOperator int

const (
	// Арифметические операторы
	ADD BinaryOperator = iota
	SUB
	MUL
	DIV
	MOD

	// Операторы сравнения
	EQ // равно
	NE // не равно
	LT // меньше
	LE // меньше или равно
	GT // больше
	GE // больше или равно
)

// String возвращает строковое представление оператора
func (op BinaryOperator) String() string {
	switch op {
	case ADD:
		return "+"
	case SUB:
		return "-"
	case MUL:
		return "*"
	case DIV:
		return "/"
	case MOD:
		return "%"
	case EQ:
		return "=="
	case NE:
		return "!="
	case LT:
		return "<"
	case LE:
		return "<="
	case GT:
		return ">"
	case GE:
		return ">="
	default:
		return "unknown"
	}
}

// Логические операторы
type LogicalOperator int

const (
	AND LogicalOperator = iota
	OR
	NOT
	IMPLIES
)

// String возвращает строковое представление логического оператора
func (op LogicalOperator) String() string {
	switch op {
	case AND:
		return "&&"
	case OR:
		return "||"
	case NOT:
		return "!"
	case IMPLIES:
		return "=>"
	default:
		return "unknown"
	}
}

type Ref struct {
	Idx     *SymbolicVariable
	RefType ExpressionType
}

func NewRef(idx int64, refType ExpressionType) *Ref {
	return &Ref{
		Idx:     NewSymbolicVariable(fmt.Sprint(idx), IntType),
		RefType: refType,
	}
}

func (r *Ref) Type() ExpressionType {
	return r.RefType
}

func (r *Ref) String() string {
	return fmt.Sprintf("&%s", r.Idx.String())
}

func (r *Ref) Accept(visitor Visitor) interface{} {
	panic("implement me")
}

// StringConstant представляет строковую константу
type StringConstant struct {
	Value string
}

// NewStringConstant создаёт новую строковую константу
func NewStringConstant(value string) *StringConstant {
	return &StringConstant{Value: value}
}

// Type возвращает тип константы
func (sc *StringConstant) Type() ExpressionType {
	return StringType
}

// String возвращает строковое представление константы
func (sc *StringConstant) String() string {
	// с кавычками, чтобы в дампах было видно, что это строка
	return fmt.Sprintf("%q", sc.Value)
}

// Accept реализует Visitor pattern
func (sc *StringConstant) Accept(visitor Visitor) interface{} {
	return visitor.VisitStringConstant(sc)
}
