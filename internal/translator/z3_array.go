package translator

import (
	"fmt"
	"symbolic-execution-course/internal/symbolic"

	"github.com/ebukreev/go-z3/z3"
)

func (zt *Z3Translator) VisitArray(expr *symbolic.Array) interface{} {
	// ВАЖНО: многомерность будет работать только если range-sort массива = ArraySort,
	// иначе Select никогда не вернёт z3.Array и каст не сработает.

	idxSort := zt.bvSort
	elemSort := zt.bvSort

	arrSort := zt.ctx.ArraySort(idxSort, elemSort)

	// базовый массив (константа)
	cur := zt.ctx.Const(expr.Name, arrSort).(z3.Array)

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
			sel := cur.Select(idx) // z3.Value

			// Если вернулся массив (multi-dim как массив массивов) — продолжаем цепочку
			if next, ok := sel.(z3.Array); ok {
				cur = next
				continue
			}

			// Иначе это скаляр — досрочно выходим
			return sel

		default:
			panic(fmt.Sprintf("VisitArray: unknown op kind %v at #%d", op.Kind, i))
		}
	}

	return cur
}
