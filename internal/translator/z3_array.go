package translator

import (
	"fmt"
	"symbolic-execution-course/internal/symbolic"

	"github.com/ebukreev/go-z3/z3"
)

func (zt *Z3Translator) VisitArray(expr *symbolic.Array) interface{} {
	idxSort := zt.bvSort
	elemSort := zt.arrayElemSort(expr.ElemType)

	arrSort := zt.ctx.ArraySort(idxSort, elemSort)

	// базовый массив (константа)
	cur := zt.ctx.Const(expr.Name, arrSort).(z3.Array)

	cs := make([]z3.Bool, 0)

	for i, op := range expr.Ops {
		idxV, err := zt.TranslateExpression(op.Idx)
		if err != nil {
			panic(fmt.Sprintf("VisitArray: translate op #%d index: %v", i, err))
		}
		idx, ok := idxV.(z3.Value)
		if !ok {
			panic(fmt.Sprintf("VisitArray: op #%d index is not z3.Value (got %T)", i, idxV))
		}

		switch op.Kind {
		case symbolic.OpStore:
			valV, err := zt.TranslateExpression(op.Val)
			if err != nil {
				panic(fmt.Sprintf("VisitArray: translate store #%d value: %v", i, err))
			}
			val, ok := valV.(z3.Value)
			if !ok {
				panic(fmt.Sprintf("VisitArray: store #%d value is not z3.Value (got %T)", i, valV))
			}

			cur = cur.Store(idx, val)

		case symbolic.OpSelect:
			// истинное чтение из текущего состояния массива
			sel := cur.Select(idx) // z3.Value

			// op.Val — результат Access: переменная, которую надо приравнять к sel
			resV, err := zt.TranslateExpression(op.Val)
			if err != nil {
				panic(fmt.Sprintf("VisitArray: translate select #%d result: %v", i, err))
			}

			// ВАЖНО: здесь тип может быть BV / Bool / Uninterpreted / Array
			switch resT := resV.(type) {

			case z3.BV:
				selT, ok := sel.(z3.BV)
				if !ok {
					panic(fmt.Sprintf("VisitArray: select #%d type mismatch: res is BV, sel is %T", i, sel))
				}
				cs = append(cs, resT.Eq(selT))

			case z3.Bool:
				selT, ok := sel.(z3.Bool)
				if !ok {
					panic(fmt.Sprintf("VisitArray: select #%d type mismatch: res is Bool, sel is %T", i, sel))
				}
				cs = append(cs, resT.Eq(selT))

			case z3.Uninterpreted:
				selT, ok := sel.(z3.Uninterpreted)
				if !ok {
					panic(fmt.Sprintf("VisitArray: select #%d type mismatch: res is Uninterpreted, sel is %T", i, sel))
				}
				cs = append(cs, resT.Eq(selT))

			case z3.Array:
				selT, ok := sel.(z3.Array)
				if !ok {
					panic(fmt.Sprintf("VisitArray: select #%d type mismatch: res is Array, sel is %T", i, sel))
				}
				cs = append(cs, resT.Eq(selT))

			case z3.Value:
				// fallback: если вдруг библиотека возвращает какой-то Value-тип,
				// у которого Eq есть только после конкретизации — лучше явно падать,
				// чтобы поймать новый тип и добавить ветку.
				panic(fmt.Sprintf("VisitArray: select #%d unsupported Value subtype: %T (sel=%T)", i, resV, sel))

			default:
				panic(fmt.Sprintf("VisitArray: select #%d result is not a z3.Value subtype (got %T)", i, resV))
			}

		default:
			panic(fmt.Sprintf("VisitArray: unknown op kind %v at #%d", op.Kind, i))
		}
	}

	// На этом шаге возвращаем булеву конъюнкцию ограничений
	if len(cs) == 0 {
		return zt.ctx.FromBool(true)
	}
	if len(cs) == 1 {
		return cs[0]
	}
	return cs[0].And(cs[1:]...)
}

func (zt *Z3Translator) arrayElemSort(et symbolic.ExpressionType) z3.Sort {
	switch et {
	case symbolic.IntType:
		return zt.bvSort
	case symbolic.BoolType:
		return zt.ctx.BoolSort()
	case symbolic.StringType:
		// если строки у тебя как uninterpreted:
		return zt.stringSort // либо zt.ctx.UninterpretedSort("String")
	default:
		// Для ArraySegment у тебя ElemType = IntType, так что сюда не должно попасть
		panic(fmt.Sprintf("unsupported array elem type: %v", et))
	}
}
