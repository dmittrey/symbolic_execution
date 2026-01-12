package main

import (
	"fmt"
	"log"
	"symbolic-execution-course/internal"
	"symbolic-execution-course/internal/translator"

	"github.com/ebukreev/go-z3/z3"
)

func main() {
	source := `
package final_tests

// ==== aliasing ====

type Foo struct {
	a int
}

func IntAliasing(foo1 int, foo2 int) int {
	if foo2 == foo1 {
		return 4
	}
	return 5
}

func BoolAliasing(foo1 bool, foo2 bool) int {
	if foo2 == foo1 {
		return 4
	}
	return 5
}

func StringAliasing(foo1 string, foo2 string) int {
	if foo2 == foo1 {
		return 4
	}
	return 5
}

func IntRefAliasing(foo1 *int, foo2 *int) int {
	*foo2 = 5
	*foo1 = 2
	if *foo2 == 2 {
		return 4
	}
	return 5
}

func BoolRefAliasing(foo1 *bool, foo2 *bool) int {
	*foo2 = true
	*foo1 = false
	if *foo2 == false {
		return 4
	}
	return 5
}

func StringRefAliasing(foo1 *string, foo2 *string) int {
	*foo2 = "hi"
	*foo1 = "kek"
	if *foo2 == "kek" {
		return 4
	}
	return 5
}

func StructAliasing(foo1 *Foo, foo2 *Foo) int {
	foo2.a = 5
	foo1.a = 2
	if foo2.a == 2 {
		return 4
	}
	return 5
}

func ArrayAliasing(arr1 []int, arr2 []int) int {
	arr2[1] = 5
	arr1[1] = 2
	if arr2[1] == 2 {
		return 4
	}
	return 5
}

// ==== arrays ====

func DefaultBooleanValues() []bool {
	array := make([]bool, 3)
	if !array[1] {
		return array
	}
	return array
}

func ByteArray(a []byte, x byte, len_a int) byte {
	if len_a != 2 {
		return 255
	}
	a[0] = 5
	a[1] = x
	if a[0] + a[1] > 20 {
		return 1
	}
	return 0
}

func CharSizeAndIndex(a []rune, x rune, len_a int) byte {
	if a == nil || len_a <= int(x) || x < 1 {
		return 255
	}
	b := make([]rune, x)
	b[0] = 5
	a[x] = x
	if b[0] + a[x] > 7 {
		return 1
	}
	return 0
}

func BooleanArray(arr []bool, len_arr int) int {
	if len_arr == 0 {
		return 1
	}
	if arr[0] {
		return 2
	}
	arr[0] = true
	return 3
}

func CreateArray(x, y, length, i int) []*ObjectWithPrimitivesClass {
	if length < 3 {
		return nil
	}

	array := make([]*ObjectWithPrimitivesClass, length)

	for i := 0; i < length; i++ {
		array[i] = &ObjectWithPrimitivesClass{}
		array[i].x = x + i
		array[i].y = y + i
	}

	return array
}

func IsIdentityMatrix(matrix [][]int, len_matrix1 []int, len_matrix int) bool {
	if len_matrix < 3 {
		return false
	}
	for i := 0; i < len_matrix; i++ {
		if len_matrix1[i] != len_matrix {
			return false
		}
		for j := 0; j < len_matrix1[i]; j++ {
			if i == j && matrix[i][j] != 1 {
				return false
			}
			if i != j && matrix[i][j] != 0 {
				return false
			}
		}
	}
	return true
}

func ReallyMultiDimensionalArray(array [][][]int) [][][]int {
	if array[1][2][3] != 12345 {
		array[1][2][3] = 12345
	} else {
		array[1][2][3] -= 12345 * 2
	}
	return array
}

func FillMultiArrayWithArray(value []int, len_value int) [][]int {
	if len_value < 2 {
		return make([][]int, 0)
	}
	for i := range value {
		value[i] += i
	}
	length := 3
	array := make([][]int, length)
	for i := 0; i < length; i++ {
		array[i] = value
	}
	return array
}

// ==== bit ====

func Complement(x int) bool {
	return ^x == 1
}

func Xor(x, y int) bool {
	return (x ^ y) == 0
}

func Or(val int) bool {
	return (val | 7) == 15
}

func And(value int) bool {
	return (value & (value - 1)) == 0
}

func BooleanNot(boolA, boolB bool) int {
	d := boolA && boolB
	e := (!boolA) || boolB
	if d && e {
		return 100
	}
	return 200
}

func BooleanXorCompare(aBool, bBool bool) int {
	if aBool != bBool {
		return 1
	}
	return 0
}

func ShlWithBigLongShift(shift int64) int {
	if shift < 40 {
		return 1
	}
	if (0x77777777 << shift) == 0x77777770 {
		return 2
	}
	return 3
}

// ==== calls ====

type InvokeClass struct {
	Value int
}

func (i *InvokeClass) DivBy(den int) int {
	return i.Value / den
}

func (i *InvokeClass) UpdateValue(newValue int) {
	i.Value = newValue
}

func mult(a, b int) int {
	return a * b
}

func half(a int) int {
	return a / 2
}

func SimpleFormula(fst, snd int) int {
	if fst < 100 {
		return 0
	} else if snd < 100 {
		return 0
	}

	x := fst + 5
	y := half(snd)

	return mult(x, y)
}

func initialize(value int) *InvokeClass {
	objectValue := &InvokeClass{
		Value: value,
	}
	return objectValue
}

func CreateObjectFromValue(value int) *InvokeClass {
	if value == 0 {
		value = 1
	}
	return initialize(value)
}

func changeValue(objectValue *InvokeClass, value int) {
	objectValue.Value = value
}

func ChangeObjectValueByMethod(objectValue *InvokeClass) *InvokeClass {
	objectValue.Value = 1
	changeValue(objectValue, 4)
	return objectValue
}

func getFive() int {
	return 5
}

func getTwo() int {
	return 2
}

func ParticularValue(invokeObject *InvokeClass) *InvokeClass {
	if invokeObject.Value < 0 {
		return nil
	}
	x := getFive() * getTwo()
	y := getFive() / getTwo()

	invokeObject.Value = x + y
	return invokeObject
}

func getNull() *InvokeClass {
	return nil
}

func GetNullOrValue(invokeObject *InvokeClass) *InvokeClass {
	if invokeObject.Value < 100 {
		return getNull()
	}
	invokeObject.Value = getFive()
	return invokeObject
}

// ==== doubles ====

func CompareWithDiv(a, b float64) float64 {
	z := a + 0.5
	if (a / z) > b {
		return 1.0
	} else {
		return 0.0
	}
}

func Mul(a, b float64) float64 {
	if a*b > 33.32 && a*b < 33.333 {
		return 1.1
	} else if a*b > 33.333 && a*b < 33.7592 {
		return 1.2
	} else {
		return 1.3
	}
}

// ==== loops ====

func LoopWithConcreteBound(n int) int {
	result := 0
	for i := 0; i < 10; i++ {
		result += i
	}
	return result
}

func LoopWithSymbolicBound(n int) int {
	if n > 10 {
		return -1
	}

	result := 0
	for i := 0; i < n; i++ {
		result += i
	}
	return result
}

func LoopWithSymbolicBoundAndSymbolicBranching(n int, condition bool) int {
	if n > 10 {
		return -1
	}

	result := 0
	for i := 0; i < n; i++ {
		if condition && i%2 == 0 {
			result += i
		}
	}
	return result
}

func LoopWithSymbolicBoundAndComplexControlFlow(n int, condition bool) int {
	if n > 10 {
		return -1
	}

	result := 0
	for i := 0; i < n; i++ {
		if condition && i == 3 {
			break
		}
		if i%2 != 0 {
			continue
		}
		result += i
	}
	return result
}

func WhileCycle(x int) int {
	i := 0
	sum := 0
	for i < x {
		sum += i
		i++
	}
	return sum
}

func LoopInsideLoop(x int) int {
	for i := x - 5; i < x; i++ {
		if i < 0 {
			return 2
		} else {
			for j := i; j < x+i; j++ {
				if j == 7 {
					return 1
				}
			}
		}
	}
	return -1
}

// ==== recursion ====

func Factorial(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	result := Factorial(n - 1)
	return n * result
}

func InfiniteRecursion(i int) {
	InfiniteRecursion(i + 1)
}

// ==== structs ====

type ObjectWithPrimitivesClass struct {
	ValueByDefault int
	x, y           int
	ShortValue     int16
	Weight         float64
}

func NewObjectWithPrimitivesClass() *ObjectWithPrimitivesClass {
	return &ObjectWithPrimitivesClass{ValueByDefault: 5}
}

func Max(fst, snd *ObjectWithPrimitivesClass) *ObjectWithPrimitivesClass {
	if fst.x > snd.x && fst.y > snd.y {
		return fst
	} else if fst.x < snd.x && fst.y < snd.y {
		return snd
	}
	return fst
}

func Example(value *ObjectWithPrimitivesClass) *ObjectWithPrimitivesClass {
	if value.x == 1 {
		return value
	}
	value.x = 1
	return value
}

func CreateObject(a, b int, objectExample *ObjectWithPrimitivesClass) *ObjectWithPrimitivesClass {
	object := &ObjectWithPrimitivesClass{ValueByDefault: 5}
	object.x = a + 5
	object.y = b + 6
	object.Weight = objectExample.Weight
	if object.Weight < 0 {
		return nil
	}
	return object
}

func Memory(objectExample *ObjectWithPrimitivesClass, value int) *ObjectWithPrimitivesClass {
	if value > 0 {
		objectExample.x = 1
		objectExample.y = 2
		objectExample.Weight = 1.2
	} else {
		objectExample.x = -1
		objectExample.y = -2
		objectExample.Weight = -1.2
	}
	return objectExample
}

func CompareTwoNullObjects(value int) int {
	fst := &ObjectWithPrimitivesClass{ValueByDefault: 5}
	snd := &ObjectWithPrimitivesClass{ValueByDefault: 5}

	fst.x = value + 1
	snd.x = value + 2

	if fst.x == value+1 {
		fst = nil
	}
	if snd.x == value+2 {
		snd = nil
	}

	if fst == snd {
		return 1
	}
	return 0
}

type SimpleDataClass struct {
	a int
	b int
}

type ObjectWithRefFieldClass struct {
	x, y       int
	Weight     float64
	RefField   *SimpleDataClass
	ArrayField []int
}

func WriteToRefTypeField(objectExample *ObjectWithRefFieldClass, value int) *ObjectWithRefFieldClass {
	if value != 42 {
		return nil
	}

	simpleDataClass := &SimpleDataClass{
		a: value,
		b: value * 2,
	}
	objectExample.RefField = simpleDataClass
	return objectExample
}

func WriteToArrayField(objectExample *ObjectWithRefFieldClass, length int) *ObjectWithRefFieldClass {
	if length < 3 {
		return nil
	}

	array := make([]int, length)
	for i := 0; i < length; i++ {
		array[i] = i + 1
	}

	objectExample.ArrayField = array
	objectExample.ArrayField[length-1] = 100

	return objectExample
}

func ReadFromArrayField(objectExample *ObjectWithRefFieldClass, value int) int {
	if len(objectExample.ArrayField) > 2 && objectExample.ArrayField[2] == value {
		return 1
	}
	return 2
}

func CompareTwoDifferentObjectsFromArguments(fst, snd *ObjectWithRefFieldClass) int {
	if fst.x > 0 && snd.x < 0 {
		if fst == snd {
			return 0
		} else {
			return 1
		}
	}

	fst.x = snd.x
	fst.y = snd.y
	fst.Weight = snd.Weight

	if fst == snd {
		return 2
	}

	return 3
}

func CompareTwoObjectsWithTheSameRefField(fst, snd *ObjectWithRefFieldClass) int {
	simpleDataClass := &SimpleDataClass{}

	fst.RefField = simpleDataClass
	snd.RefField = simpleDataClass
	fst.x = snd.x
	fst.y = snd.y
	fst.Weight = snd.Weight

	if fst == snd {
		return 1
	}
	return 2
}

type RecursiveTypeClass struct {
	Next  *RecursiveTypeClass
	Value int
}

func NextValue(node *RecursiveTypeClass, value int) *RecursiveTypeClass {
	if value == 0 {
		return nil
	}
	if node.Next != nil && node.Next.Value == value {
		return node.Next
	}
	return nil
}

func WriteObjectField(node *RecursiveTypeClass) *RecursiveTypeClass {
	if node.Next == nil {
		node.Next = &RecursiveTypeClass{}
	}
	node.Next.Value = node.Next.Value + 1
	return node
}

type Person struct {
	Name string
	Age  int
	ID   int
}

func TestPathConstraintMutability(p Person) int {
	if p.Age != 18 {
		p.Age = 18
		if p.Age != 18 {
			return 0
		} else {
			return 1
		}
	}
	return 2
}
`

	funcs := []string{
		// aliasing
		"IntAliasing",
		"BoolAliasing",
		"StringAliasing",
		"IntRefAliasing",
		"BoolRefAliasing",
		"StringRefAliasing",
		"StructAliasing",
		"ArrayAliasing",

		// arrays
		"DefaultBooleanValues",
		"ByteArray",
		"CharSizeAndIndex",
		"BooleanArray",
		"CreateArray",
		"IsIdentityMatrix",
		"ReallyMultiDimensionalArray",
		"FillMultiArrayWithArray",

		// bit
		"Complement",
		"Xor",
		"Or",
		"And",
		"BooleanNot",
		"BooleanXorCompare",
		"ShlWithBigLongShift",

		// calls (future work in HW5)
		"SimpleFormula",
		"CreateObjectFromValue",
		"ChangeObjectValueByMethod",
		"ParticularValue",
		"GetNullOrValue",

		// doubles
		"CompareWithDiv",
		"Mul",

		// loops
		"LoopWithConcreteBound",
		"LoopWithSymbolicBound",
		"LoopWithSymbolicBoundAndSymbolicBranching",
		"LoopWithSymbolicBoundAndComplexControlFlow",
		"WhileCycle",
		"LoopInsideLoop",

		// recursion (future work in HW5)
		"Factorial",
		"InfiniteRecursion",

		// structs
		"NewObjectWithPrimitivesClass",
		"Max",
		"Example",
		"CreateObject",
		"Memory",
		"CompareTwoNullObjects",
		"WriteToRefTypeField",
		"WriteToArrayField",
		"ReadFromArrayField",
		"CompareTwoDifferentObjectsFromArguments",
		"CompareTwoObjectsWithTheSameRefField",
		"NextValue",
		"WriteObjectField",
		"TestPathConstraintMutability",
	}

	analyser := internal.NewAnalyser()

	for _, fn := range funcs {
		fmt.Printf("\n==================== %s ====================\n", fn)
		result := analyser.Analyse(source, fn)

		fmt.Printf("States: %d\n", len(result))
		for i, st := range result {
			tr := translator.NewZ3Translator()
			defer tr.Close()

			fmt.Printf("---- State #%d ----\n", i)
			fmt.Println(st.Dump())

			tr.AssertExpr(st.PathCondition)
			for _, c := range st.Heap.Constraints() {
				tr.AssertExpr(c)
			}

			if m := mustSat(tr); m != nil {
				fmt.Printf("Model:\n%s\n", m.String())
			}
		}
	}
}

func mustSat(tr *translator.Z3Translator) *z3.Model {
	sat, err := tr.IsSatisfiable()
	if err != nil {
		log.Fatalf("Check error: %v", err)
	}
	if !sat {
		fmt.Println("UNSAT")
		return nil
	}
	fmt.Println("SAT")
	return tr.Model()
}
