
==================== Factorial ====================
CFG построен для функции с 5 блоками
interpretDynamically: n < 0:int
interpretDynamically: if t0 goto 1 else 2
resolveExpression: n < 0:int
resolveExpression: parameter n : int
resolveExpression: 0:int
interpretDynamically: return 0:int
resolveExpression: 0:int
interpretDynamically: n == 0:int
interpretDynamically: if t1 goto 3 else 4
resolveExpression: n == 0:int
resolveExpression: parameter n : int
resolveExpression: 0:int
interpretDynamically: n - 1:int
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: Factorial(t2)
resolveExpression: n - 1:int
resolveExpression: parameter n : int
resolveExpression: 1:int
interpretDynamically: n < 0:int
interpretDynamically: if t0 goto 1 else 2
resolveExpression: n < 0:int
resolveExpression: parameter n : int
resolveExpression: 0:int
interpretDynamically: n == 0:int
interpretDynamically: if t1 goto 3 else 4
resolveExpression: n == 0:int
resolveExpression: parameter n : int
resolveExpression: 0:int
interpretDynamically: return 0:int
resolveExpression: 0:int
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: n - 1:int
interpretDynamically: Factorial(t2)
resolveExpression: n - 1:int
resolveExpression: parameter n : int
resolveExpression: 1:int
interpretDynamically: n < 0:int
interpretDynamically: if t0 goto 1 else 2
resolveExpression: n < 0:int
resolveExpression: parameter n : int
resolveExpression: 0:int
interpretDynamically: n == 0:int
interpretDynamically: return 0:int
resolveExpression: 0:int
interpretDynamically: if t1 goto 3 else 4
resolveExpression: n == 0:int
resolveExpression: parameter n : int
resolveExpression: 0:int
interpretDynamically: n - 1:int
interpretDynamically: Factorial(t2)
resolveExpression: n - 1:int
resolveExpression: parameter n : int
resolveExpression: 1:int
interpretDynamically: n < 0:int
interpretDynamically: if t0 goto 1 else 2
resolveExpression: n < 0:int
resolveExpression: parameter n : int
resolveExpression: 0:int
interpretDynamically: n == 0:int
interpretDynamically: if t1 goto 3 else 4
resolveExpression: n == 0:int
resolveExpression: parameter n : int
resolveExpression: 0:int
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: n - 1:int
interpretDynamically: Factorial(t2)
resolveExpression: n - 1:int
resolveExpression: parameter n : int
resolveExpression: 1:int
interpretDynamically: return 0:int
resolveExpression: 0:int
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: n < 0:int
interpretDynamically: if t0 goto 1 else 2
resolveExpression: n < 0:int
resolveExpression: parameter n : int
resolveExpression: 0:int
interpretDynamically: n == 0:int
interpretDynamically: if t1 goto 3 else 4
resolveExpression: n == 0:int
resolveExpression: parameter n : int
resolveExpression: 0:int
interpretDynamically: n - 1:int
interpretDynamically: Factorial(t2)
resolveExpression: n - 1:int
resolveExpression: parameter n : int
resolveExpression: 1:int
interpretDynamically: return 0:int
resolveExpression: 0:int
interpretDynamically: return 1:int
resolveExpression: 1:int
States: 10
---- State #0 ----
Interpreter{
  PC: (true & (IntSegment_sel0 < 0))
  ReturnResults:
    [0] 0
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=1, store=0, select=1)
        op[0]: select idx=addr_0, res=IntSegment_sel0
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #xffffffffffffffff)
addr_0 -> #x0000000000000000
IntSegment_sel0 -> #xffffffffffffffff

---- State #1 ----
Interpreter{
  PC: ((true & (!(IntSegment_sel0 < 0))) & (IntSegment_sel1 == 0))
  ReturnResults:
    [0] 1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=2, store=0, select=2)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
IntSegment_sel1 -> #x0000000000000000
IntSegment_sel0 -> #x0000000000000000

---- State #2 ----
Interpreter{
  PC: (((true & (!(IntSegment_sel0 < 0))) & (!(IntSegment_sel1 == 0))) & ((IntSegment_sel2 - 1) < 0))
  ReturnResults:
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=3, store=0, select=3)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #3 ----
Interpreter{
  PC: (((((true & (!(IntSegment_sel0 < 0))) & (!(IntSegment_sel1 == 0))) & (!((IntSegment_sel2 - 1) < 0))) & (!((IntSegment_sel2 - 1) == 0))) & ((IntSegment_sel2 - 1) < 0))
  ReturnResults:
  CallStack: 1 frame(s)
    Frame[0]: main.Factorial
      Locals: 1
        n = (IntSegment_sel2 - 1)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=3, store=0, select=3)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #4 ----
Interpreter{
  PC: ((((true & (!(IntSegment_sel0 < 0))) & (!(IntSegment_sel1 == 0))) & (!((IntSegment_sel2 - 1) < 0))) & ((IntSegment_sel2 - 1) == 0))
  ReturnResults:
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=3, store=0, select=3)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000001)
IntSegment_sel2 -> #x0000000000000001
IntSegment_sel1 -> #x0000000000000001
IntSegment_sel0 -> #x0000000000000001

---- State #5 ----
Interpreter{
  PC: (((((((true & (!(IntSegment_sel0 < 0))) & (!(IntSegment_sel1 == 0))) & (!((IntSegment_sel2 - 1) < 0))) & (!((IntSegment_sel2 - 1) == 0))) & (!((IntSegment_sel2 - 1) < 0))) & (!((IntSegment_sel2 - 1) == 0))) & ((IntSegment_sel2 - 1) < 0))
  ReturnResults:
  CallStack: 2 frame(s)
    Frame[0]: main.Factorial
      Locals: 1
        n = (IntSegment_sel2 - 1)
    Frame[1]: main.Factorial
      Locals: 1
        n = ((IntSegment_sel2 - 1) - 1)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=3, store=0, select=3)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #6 ----
Interpreter{
  PC: ((((((((true & (!(IntSegment_sel0 < 0))) & (!(IntSegment_sel1 == 0))) & (!((IntSegment_sel2 - 1) < 0))) & (!((IntSegment_sel2 - 1) == 0))) & (!((IntSegment_sel2 - 1) < 0))) & (!((IntSegment_sel2 - 1) == 0))) & (!((IntSegment_sel2 - 1) < 0))) & ((IntSegment_sel2 - 1) == 0))
  ReturnResults:
  CallStack: 2 frame(s)
    Frame[0]: main.Factorial
      Locals: 1
        n = (IntSegment_sel2 - 1)
    Frame[1]: main.Factorial
      Locals: 1
        n = ((IntSegment_sel2 - 1) - 1)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=3, store=0, select=3)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #7 ----
Interpreter{
  PC: ((((((true & (!(IntSegment_sel0 < 0))) & (!(IntSegment_sel1 == 0))) & (!((IntSegment_sel2 - 1) < 0))) & (!((IntSegment_sel2 - 1) == 0))) & (!((IntSegment_sel2 - 1) < 0))) & ((IntSegment_sel2 - 1) == 0))
  ReturnResults:
  CallStack: 1 frame(s)
    Frame[0]: main.Factorial
      Locals: 1
        n = (IntSegment_sel2 - 1)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=3, store=0, select=3)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #8 ----
Interpreter{
  PC: (((((((((true & (!(IntSegment_sel0 < 0))) & (!(IntSegment_sel1 == 0))) & (!((IntSegment_sel2 - 1) < 0))) & (!((IntSegment_sel2 - 1) == 0))) & (!((IntSegment_sel2 - 1) < 0))) & (!((IntSegment_sel2 - 1) == 0))) & (!((IntSegment_sel2 - 1) < 0))) & (!((IntSegment_sel2 - 1) == 0))) & ((IntSegment_sel2 - 1) < 0))
  ReturnResults:
  CallStack: 3 frame(s)
    Frame[0]: main.Factorial
      Locals: 1
        n = (IntSegment_sel2 - 1)
    Frame[1]: main.Factorial
      Locals: 1
        n = ((IntSegment_sel2 - 1) - 1)
    Frame[2]: main.Factorial
      Locals: 1
        n = ((IntSegment_sel2 - 1) - 1)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=3, store=0, select=3)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #9 ----
Interpreter{
  PC: ((((((((((true & (!(IntSegment_sel0 < 0))) & (!(IntSegment_sel1 == 0))) & (!((IntSegment_sel2 - 1) < 0))) & (!((IntSegment_sel2 - 1) == 0))) & (!((IntSegment_sel2 - 1) < 0))) & (!((IntSegment_sel2 - 1) == 0))) & (!((IntSegment_sel2 - 1) < 0))) & (!((IntSegment_sel2 - 1) == 0))) & (!((IntSegment_sel2 - 1) < 0))) & ((IntSegment_sel2 - 1) == 0))
  ReturnResults:
  CallStack: 3 frame(s)
    Frame[0]: main.Factorial
      Locals: 1
        n = (IntSegment_sel2 - 1)
    Frame[1]: main.Factorial
      Locals: 1
        n = ((IntSegment_sel2 - 1) - 1)
    Frame[2]: main.Factorial
      Locals: 1
        n = ((IntSegment_sel2 - 1) - 1)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=3, store=0, select=3)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT

==================== InfiniteRecursion ====================
CFG построен для функции с 1 блоками
interpretDynamically: i + 1:int
interpretDynamically: InfiniteRecursion(t0)
resolveExpression: i + 1:int
resolveExpression: parameter i : int
resolveExpression: 1:int
interpretDynamically: i + 1:int
interpretDynamically: InfiniteRecursion(t0)
resolveExpression: i + 1:int
resolveExpression: parameter i : int
resolveExpression: 1:int
interpretDynamically: i + 1:int
interpretDynamically: InfiniteRecursion(t0)
resolveExpression: i + 1:int
resolveExpression: parameter i : int
resolveExpression: 1:int
interpretDynamically: i + 1:int
interpretDynamically: InfiniteRecursion(t0)
resolveExpression: i + 1:int
resolveExpression: parameter i : int
resolveExpression: 1:int
interpretDynamically: i + 1:int
interpretDynamically: InfiniteRecursion(t0)
resolveExpression: i + 1:int
resolveExpression: parameter i : int
resolveExpression: 1:int
States: 0
