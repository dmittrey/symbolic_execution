
==================== SimpleFormula ====================
CFG построен для функции с 5 блоками
interpretDynamically: fst < 100:int
interpretDynamically: if t0 goto 1 else 2
resolveExpression: fst < 100:int
resolveExpression: parameter fst : int
resolveExpression: 100:int
interpretDynamically: snd < 100:int
interpretDynamically: if t1 goto 3 else 4
resolveExpression: snd < 100:int
resolveExpression: parameter snd : int
resolveExpression: 100:int
interpretDynamically: return 0:int
resolveExpression: 0:int
interpretDynamically: return 0:int
resolveExpression: 0:int
interpretDynamically: fst + 5:int
interpretDynamically: half(snd)
resolveExpression: parameter snd : int
interpretDynamically: a / 2:int
interpretDynamically: return t0
resolveExpression: a / 2:int
resolveExpression: parameter a : int
resolveExpression: 2:int
interpretDynamically: mult(t2, t3)
resolveExpression: fst + 5:int
resolveExpression: parameter fst : int
resolveExpression: 5:int
resolveExpression: half(snd)
interpretDynamically: a * b
interpretDynamically: return t0
resolveExpression: a * b
resolveExpression: parameter a : int
resolveExpression: parameter b : int
interpretDynamically: return t4
resolveExpression: mult(t2, t3)
States: 3
---- State #0 ----
Interpreter{
  PC: ((true & (!(IntSegment_sel0 < 100))) & (IntSegment_sel1 < 100))
  ReturnResults:
    [0] 0
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=2, store=0, select=2)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_1, res=IntSegment_sel1
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_1, IntSegment_sel1)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
addr_1 -> #xffffffffffffffff
addr_0 -> #x0000000000000000
IntSegment -> (store ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x48f9ba1d6e4b6864)
                     #xffffffffffffffff
                     #x8000000000000060)
IntSegment_sel1 -> #x8000000000000060
IntSegment_sel0 -> #x48f9ba1d6e4b6864

---- State #1 ----
Interpreter{
  PC: (true & (IntSegment_sel0 < 100))
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
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000063)
addr_0 -> #x0000000000000000
IntSegment_sel0 -> #x0000000000000063

---- State #2 ----
Interpreter{
  PC: ((true & (!(IntSegment_sel0 < 100))) & (!(IntSegment_sel1 < 100)))
  ReturnResults:
    [0] ((IntSegment_sel3 + 5) * (IntSegment_sel2 / 2))
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_1, res=IntSegment_sel1
        op[2]: select idx=addr_1, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_1, IntSegment_sel1).select(addr_1, IntSegment_sel2).select(addr_0, IntSegment_sel3)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
addr_1 -> #x0000000000000000
addr_0 -> #x0000000000000000
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000070)
IntSegment_sel2 -> #x0000000000000070
IntSegment_sel1 -> #x0000000000000070
IntSegment_sel3 -> #x0000000000000070
IntSegment_sel0 -> #x0000000000000070


==================== CreateObjectFromValue ====================
CFG построен для функции с 3 блоками
interpretDynamically: value == 0:int
interpretDynamically: if t0 goto 1 else 2
resolveExpression: value == 0:int
resolveExpression: parameter value : int
resolveExpression: 0:int
interpretDynamically: phi [0: value, 1: 1:int] #value
resolveExpression: parameter value : int
interpretDynamically: jump 2
interpretDynamically: phi [0: value, 1: 1:int] #value
resolveExpression: 1:int
interpretDynamically: initialize(t1)
resolveExpression: phi [0: value, 1: 1:int] #value
interpretDynamically: initialize(t1)
resolveExpression: phi [0: value, 1: 1:int] #value
interpretDynamically: new InvokeClass (complit)
interpretDynamically: &t0.Value [#0]
interpretDynamically: *t1 = value
resolveExpression: parameter value : int
resolveExpression: new InvokeClass (complit)
interpretDynamically: return t0
resolveExpression: new InvokeClass (complit)
interpretDynamically: return t2
resolveExpression: initialize(t1)
interpretDynamically: new InvokeClass (complit)
interpretDynamically: &t0.Value [#0]
interpretDynamically: *t1 = value
resolveExpression: parameter value : int
resolveExpression: new InvokeClass (complit)
interpretDynamically: return t0
resolveExpression: new InvokeClass (complit)
interpretDynamically: return t2
resolveExpression: initialize(t1)
States: 2
---- State #0 ----
Interpreter{
  PC: (true & (!(IntSegment_sel0 == 0)))
  ReturnResults:
    [0] &(addr_1)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 1, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    AllocationsList:
      A[0]: &(addr_1) (refType=struct, type=InvokeClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=2, store=0, select=2)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      InvokeClass:
        field[0]: InvokeClass_0 (elem=int, ops=1, store=1, select=0)
        op[0]: store idx=addr_1, val=IntSegment_sel1
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_1)
    C[1]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1)
    C[2]: BoolSegment
    C[3]: StringSegment
    C[4]: FloatSegment
    C[5]: ArraySegment
    C[6]: InvokeClass_0.store(addr_1, IntSegment_sel1)
}

SAT
Model:
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #xffffffffffffffff)
addr_1 -> #xffffffffffffffff
addr_0 -> #x0000000000000000
IntSegment_sel1 -> #xffffffffffffffff
IntSegment_sel0 -> #xffffffffffffffff

---- State #1 ----
Interpreter{
  PC: (true & (IntSegment_sel0 == 0))
  ReturnResults:
    [0] &(addr_1)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 1, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    AllocationsList:
      A[0]: &(addr_1) (refType=struct, type=InvokeClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=1, store=0, select=1)
        op[0]: select idx=addr_0, res=IntSegment_sel0
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      InvokeClass:
        field[0]: InvokeClass_0 (elem=int, ops=1, store=1, select=0)
        op[0]: store idx=addr_1, val=1
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_1)
    C[1]: IntSegment.select(addr_0, IntSegment_sel0)
    C[2]: BoolSegment
    C[3]: StringSegment
    C[4]: FloatSegment
    C[5]: ArraySegment
    C[6]: InvokeClass_0.store(addr_1, 1)
}

SAT
Model:
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
addr_1 -> #xffffffffffffffff
addr_0 -> #x0000000000000000
IntSegment_sel0 -> #x0000000000000000


==================== ChangeObjectValueByMethod ====================
CFG построен для функции с 1 блоками
interpretDynamically: &objectValue.Value [#0]
interpretDynamically: *t0 = 1:int
resolveExpression: 1:int
resolveExpression: parameter objectValue : *InvokeClass
interpretDynamically: changeValue(objectValue, 4:int)
resolveExpression: parameter objectValue : *InvokeClass
resolveExpression: 4:int
interpretDynamically: &objectValue.Value [#0]
interpretDynamically: *t0 = value
resolveExpression: parameter value : int
resolveExpression: parameter objectValue : *InvokeClass
interpretDynamically: return
interpretDynamically: return objectValue
resolveExpression: parameter objectValue : *InvokeClass
States: 1
---- State #0 ----
Interpreter{
  PC: true
  ReturnResults:
    [0] &(addr_0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=InvokeClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      InvokeClass:
        field[0]: InvokeClass_0 (elem=int, ops=2, store=2, select=0)
        op[0]: store idx=addr_0, val=1
        op[1]: store idx=addr_0, val=4
  }
  HeapConstraints: 6
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: InvokeClass_0.store(addr_0, 1).store(addr_0, 4)
}

SAT
Model:


==================== ParticularValue ====================
CFG построен для функции с 3 блоками
interpretDynamically: &invokeObject.Value [#0]
interpretDynamically: *t0
interpretDynamically: t1 < 0:int
interpretDynamically: if t2 goto 1 else 2
resolveExpression: t1 < 0:int
resolveExpression: *t0
resolveExpression: parameter invokeObject : *InvokeClass
resolveExpression: 0:int
interpretDynamically: getFive()
interpretDynamically: return 5:int
resolveExpression: 5:int
interpretDynamically: getTwo()
interpretDynamically: return 2:int
resolveExpression: 2:int
interpretDynamically: t3 * t4
interpretDynamically: getFive()
interpretDynamically: return 5:int
resolveExpression: 5:int
interpretDynamically: getTwo()
interpretDynamically: return 2:int
resolveExpression: 2:int
interpretDynamically: t6 / t7
interpretDynamically: t5 + t8
interpretDynamically: &invokeObject.Value [#0]
interpretDynamically: *t10 = t9
resolveExpression: t5 + t8
resolveExpression: t3 * t4
resolveExpression: getFive()
resolveExpression: getTwo()
resolveExpression: t6 / t7
resolveExpression: getFive()
resolveExpression: getTwo()
resolveExpression: parameter invokeObject : *InvokeClass
interpretDynamically: return invokeObject
resolveExpression: parameter invokeObject : *InvokeClass
interpretDynamically: return nil:*InvokeClass
resolveExpression: nil:*main.InvokeClass
States: 2
---- State #0 ----
Interpreter{
  PC: (true & (!(InvokeClass_0_sel0 < 0)))
  ReturnResults:
    [0] &(addr_0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=InvokeClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      InvokeClass:
        field[0]: InvokeClass_0 (elem=int, ops=2, store=1, select=1)
        op[0]: select idx=addr_0, res=InvokeClass_0_sel0
        op[1]: store idx=addr_0, val=((5 * 2) + (5 / 2))
  }
  HeapConstraints: 6
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: InvokeClass_0.select(addr_0, InvokeClass_0_sel0).store(addr_0, ((5 * 2) + (5 / 2)))
}

SAT
Model:
InvokeClass_0 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
addr_0 -> #x0000000000000000
InvokeClass_0_sel0 -> #x0000000000000000

---- State #1 ----
Interpreter{
  PC: (true & (InvokeClass_0_sel0 < 0))
  ReturnResults:
    [0] nil
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=InvokeClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      InvokeClass:
        field[0]: InvokeClass_0 (elem=int, ops=1, store=0, select=1)
        op[0]: select idx=addr_0, res=InvokeClass_0_sel0
  }
  HeapConstraints: 6
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: InvokeClass_0.select(addr_0, InvokeClass_0_sel0)
}

SAT
Model:
InvokeClass_0 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #xffffffffffffffff)
addr_0 -> #x0000000000000000
InvokeClass_0_sel0 -> #xffffffffffffffff


==================== GetNullOrValue ====================
CFG построен для функции с 3 блоками
interpretDynamically: &invokeObject.Value [#0]
interpretDynamically: *t0
interpretDynamically: t1 < 100:int
interpretDynamically: if t2 goto 1 else 2
resolveExpression: t1 < 100:int
resolveExpression: *t0
resolveExpression: parameter invokeObject : *InvokeClass
resolveExpression: 100:int
interpretDynamically: getFive()
interpretDynamically: getNull()
interpretDynamically: return 5:int
resolveExpression: 5:int
interpretDynamically: &invokeObject.Value [#0]
interpretDynamically: *t5 = t4
resolveExpression: getFive()
resolveExpression: parameter invokeObject : *InvokeClass
interpretDynamically: return invokeObject
resolveExpression: parameter invokeObject : *InvokeClass
interpretDynamically: return nil:*InvokeClass
resolveExpression: nil:*main.InvokeClass
interpretDynamically: return t3
resolveExpression: getNull()
States: 2
---- State #0 ----
Interpreter{
  PC: (true & (!(InvokeClass_0_sel0 < 100)))
  ReturnResults:
    [0] &(addr_0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=InvokeClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      InvokeClass:
        field[0]: InvokeClass_0 (elem=int, ops=2, store=1, select=1)
        op[0]: select idx=addr_0, res=InvokeClass_0_sel0
        op[1]: store idx=addr_0, val=5
  }
  HeapConstraints: 6
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: InvokeClass_0.select(addr_0, InvokeClass_0_sel0).store(addr_0, 5)
}

SAT
Model:
InvokeClass_0 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000064)
addr_0 -> #x0000000000000000
InvokeClass_0_sel0 -> #x0000000000000064

---- State #1 ----
Interpreter{
  PC: (true & (InvokeClass_0_sel0 < 100))
  ReturnResults:
    [0] nil
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=InvokeClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      InvokeClass:
        field[0]: InvokeClass_0 (elem=int, ops=1, store=0, select=1)
        op[0]: select idx=addr_0, res=InvokeClass_0_sel0
  }
  HeapConstraints: 6
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: InvokeClass_0.select(addr_0, InvokeClass_0_sel0)
}

SAT
Model:
InvokeClass_0 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000063)
addr_0 -> #x0000000000000000
InvokeClass_0_sel0 -> #x0000000000000063

