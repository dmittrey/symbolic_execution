
==================== IntAliasing ====================
CFG построен для функции с 3 блоками
interpretDynamically: foo2 == foo1
interpretDynamically: if t0 goto 1 else 2
resolveExpression: foo2 == foo1
resolveExpression: parameter foo2 : int
resolveExpression: parameter foo1 : int
interpretDynamically: return 5:int
resolveExpression: 5:int
interpretDynamically: return 4:int
resolveExpression: 4:int
States: 2
---- State #0 ----
Interpreter{
  PC: (true & (!(IntSegment_sel0 == IntSegment_sel1)))
  ReturnResults:
    [0] 5
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
IntSegment -> (store ((as const (Array (_ BitVec 64) (_ BitVec 64))) #xffffffffffffffff)
                     #xffffffffffffffff
                     #x0000000000000000)
IntSegment_sel1 -> #x0000000000000000
IntSegment_sel0 -> #xffffffffffffffff

---- State #1 ----
Interpreter{
  PC: (true & (IntSegment_sel0 == IntSegment_sel1))
  ReturnResults:
    [0] 4
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
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
addr_1 -> #x0000000000000000
IntSegment_sel1 -> #x0000000000000000
IntSegment_sel0 -> #x0000000000000000


==================== BoolAliasing ====================
CFG построен для функции с 3 блоками
interpretDynamically: foo2 == foo1
interpretDynamically: if t0 goto 1 else 2
resolveExpression: foo2 == foo1
resolveExpression: parameter foo2 : bool
resolveExpression: parameter foo1 : bool
interpretDynamically: return 5:int
resolveExpression: 5:int
interpretDynamically: return 4:int
resolveExpression: 4:int
States: 2
---- State #0 ----
Interpreter{
  PC: (true & (BoolSegment_sel0 == BoolSegment_sel1))
  ReturnResults:
    [0] 4
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=bool)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=2, store=0, select=2)
        op[0]: select idx=addr_0, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment.select(addr_0, BoolSegment_sel0).select(addr_1, BoolSegment_sel1)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
BoolSegment -> ((as const (Array (_ BitVec 64) Bool)) false)
addr_1 -> #x0000000000000000
BoolSegment_sel0 -> false
BoolSegment_sel1 -> false

---- State #1 ----
Interpreter{
  PC: (true & (!(BoolSegment_sel0 == BoolSegment_sel1)))
  ReturnResults:
    [0] 5
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=bool)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=2, store=0, select=2)
        op[0]: select idx=addr_0, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment.select(addr_0, BoolSegment_sel0).select(addr_1, BoolSegment_sel1)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
addr_1 -> #xffffffffffffffff
addr_0 -> #x0000000000000000
BoolSegment -> (store ((as const (Array (_ BitVec 64) Bool)) false) #xffffffffffffffff true)
BoolSegment_sel1 -> true
BoolSegment_sel0 -> false


==================== StringAliasing ====================
CFG построен для функции с 3 блоками
interpretDynamically: foo2 == foo1
interpretDynamically: if t0 goto 1 else 2
resolveExpression: foo2 == foo1
resolveExpression: parameter foo2 : string
resolveExpression: parameter foo1 : string
interpretDynamically: return 5:int
resolveExpression: 5:int
interpretDynamically: return 4:int
resolveExpression: 4:int
States: 2
---- State #0 ----
Interpreter{
  PC: (true & (StringSegment_sel0 == StringSegment_sel1))
  ReturnResults:
    [0] 4
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=string)
      P[1]: &(addr_1) (refType=string)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=2, store=0, select=2)
        op[0]: select idx=addr_0, res=StringSegment_sel0
        op[1]: select idx=addr_1, res=StringSegment_sel1
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment.select(addr_0, StringSegment_sel0).select(addr_1, StringSegment_sel1)
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
addr_1 -> #x0000000000000000
StringSegment -> ((as const (Array (_ BitVec 64) String)) String!val!0)
addr_0 -> #x0000000000000000
StringSegment_sel1 -> String!val!0
StringSegment_sel0 -> String!val!0

---- State #1 ----
Interpreter{
  PC: (true & (!(StringSegment_sel0 == StringSegment_sel1)))
  ReturnResults:
    [0] 5
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=string)
      P[1]: &(addr_1) (refType=string)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=2, store=0, select=2)
        op[0]: select idx=addr_0, res=StringSegment_sel0
        op[1]: select idx=addr_1, res=StringSegment_sel1
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment.select(addr_0, StringSegment_sel0).select(addr_1, StringSegment_sel1)
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
addr_0 -> #x0000000000000001
StringSegment -> (store ((as const (Array (_ BitVec 64) String)) String!val!0)
                        #x0000000000000000
                        String!val!1)
addr_1 -> #x0000000000000000
StringSegment_sel1 -> String!val!1
StringSegment_sel0 -> String!val!0


==================== IntRefAliasing ====================
CFG построен для функции с 3 блоками
interpretDynamically: *foo2 = 5:int
resolveExpression: 5:int
resolveExpression: parameter foo2 : *int
interpretDynamically: *foo1 = 2:int
resolveExpression: 2:int
resolveExpression: parameter foo1 : *int
interpretDynamically: *foo2
interpretDynamically: t0 == 2:int
interpretDynamically: if t1 goto 1 else 2
resolveExpression: t0 == 2:int
resolveExpression: *foo2
resolveExpression: parameter foo2 : *int
resolveExpression: 2:int
interpretDynamically: return 5:int
resolveExpression: 5:int
interpretDynamically: return 4:int
resolveExpression: 4:int
States: 2
---- State #0 ----
Interpreter{
  PC: (true & (!(IntSegment_sel2 == 2)))
  ReturnResults:
    [0] 5
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=3, store=2, select=1)
        op[0]: store idx=addr_0, val=5
        op[1]: store idx=addr_1, val=2
        op[2]: select idx=addr_0, res=IntSegment_sel2
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.store(addr_0, 5).store(addr_1, 2).select(addr_0, IntSegment_sel2)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
addr_0 -> #x0000000000000001
addr_1 -> #x0000000000000000
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
IntSegment_sel2 -> #x0000000000000005

---- State #1 ----
Interpreter{
  PC: (true & (IntSegment_sel2 == 2))
  ReturnResults:
    [0] 4
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=3, store=2, select=1)
        op[0]: store idx=addr_0, val=5
        op[1]: store idx=addr_1, val=2
        op[2]: select idx=addr_0, res=IntSegment_sel2
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.store(addr_0, 5).store(addr_1, 2).select(addr_0, IntSegment_sel2)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
addr_0 -> #x0000000000000000
addr_1 -> #x0000000000000000
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
IntSegment_sel2 -> #x0000000000000002


==================== BoolRefAliasing ====================
CFG построен для функции с 3 блоками
interpretDynamically: *foo2 = true:bool
resolveExpression: true:bool
resolveExpression: parameter foo2 : *bool
interpretDynamically: *foo1 = false:bool
resolveExpression: false:bool
resolveExpression: parameter foo1 : *bool
interpretDynamically: *foo2
interpretDynamically: t0 == false:bool
interpretDynamically: if t1 goto 1 else 2
resolveExpression: t0 == false:bool
resolveExpression: *foo2
resolveExpression: parameter foo2 : *bool
resolveExpression: false:bool
interpretDynamically: return 5:int
resolveExpression: 5:int
interpretDynamically: return 4:int
resolveExpression: 4:int
States: 2
---- State #0 ----
Interpreter{
  PC: (true & (BoolSegment_sel2 == false))
  ReturnResults:
    [0] 4
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=bool)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=3, store=2, select=1)
        op[0]: store idx=addr_0, val=true
        op[1]: store idx=addr_1, val=false
        op[2]: select idx=addr_0, res=BoolSegment_sel2
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment.store(addr_0, true).store(addr_1, false).select(addr_0, BoolSegment_sel2)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
BoolSegment -> ((as const (Array (_ BitVec 64) Bool)) false)
addr_0 -> #x0000000000000000
addr_1 -> #x0000000000000000
BoolSegment_sel2 -> false

---- State #1 ----
Interpreter{
  PC: (true & (!(BoolSegment_sel2 == false)))
  ReturnResults:
    [0] 5
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=bool)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=3, store=2, select=1)
        op[0]: store idx=addr_0, val=true
        op[1]: store idx=addr_1, val=false
        op[2]: select idx=addr_0, res=BoolSegment_sel2
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment.store(addr_0, true).store(addr_1, false).select(addr_0, BoolSegment_sel2)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
BoolSegment -> ((as const (Array (_ BitVec 64) Bool)) false)
addr_0 -> #x0000000000000001
addr_1 -> #x0000000000000000
BoolSegment_sel2 -> true


==================== StringRefAliasing ====================
CFG построен для функции с 3 блоками
interpretDynamically: *foo2 = "hi":string
resolveExpression: "hi":string
resolveExpression: parameter foo2 : *string
interpretDynamically: *foo1 = "kek":string
resolveExpression: "kek":string
resolveExpression: parameter foo1 : *string
interpretDynamically: *foo2
interpretDynamically: t0 == "kek":string
interpretDynamically: if t1 goto 1 else 2
resolveExpression: t0 == "kek":string
resolveExpression: *foo2
resolveExpression: parameter foo2 : *string
resolveExpression: "kek":string
interpretDynamically: return 4:int
resolveExpression: 4:int
interpretDynamically: return 5:int
resolveExpression: 5:int
States: 2
---- State #0 ----
Interpreter{
  PC: (true & (!(StringSegment_sel2 == "kek")))
  ReturnResults:
    [0] 5
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=string)
      P[1]: &(addr_1) (refType=string)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=3, store=2, select=1)
        op[0]: store idx=addr_0, val="hi"
        op[1]: store idx=addr_1, val="kek"
        op[2]: select idx=addr_0, res=StringSegment_sel2
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment.store(addr_0, "hi").store(addr_1, "kek").select(addr_0, StringSegment_sel2)
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
str_hi -> String!val!1
str_kek -> String!val!0
addr_1 -> #x0000000000000001
addr_0 -> #x0000000000000000
StringSegment -> (store ((as const (Array (_ BitVec 64) String)) String!val!2)
                        #x0000000000000000
                        String!val!1)
StringSegment_sel2 -> String!val!1

---- State #1 ----
Interpreter{
  PC: (true & (StringSegment_sel2 == "kek"))
  ReturnResults:
    [0] 4
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=string)
      P[1]: &(addr_1) (refType=string)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=3, store=2, select=1)
        op[0]: store idx=addr_0, val="hi"
        op[1]: store idx=addr_1, val="kek"
        op[2]: select idx=addr_0, res=StringSegment_sel2
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment.store(addr_0, "hi").store(addr_1, "kek").select(addr_0, StringSegment_sel2)
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
str_hi -> String!val!1
str_kek -> String!val!0
addr_1 -> #x0000000000000000
addr_0 -> #x0000000000000000
StringSegment -> ((as const (Array (_ BitVec 64) String)) String!val!1)
StringSegment_sel2 -> String!val!0


==================== StructAliasing ====================
CFG построен для функции с 3 блоками
interpretDynamically: &foo2.a [#0]
interpretDynamically: *t0 = 5:int
resolveExpression: 5:int
resolveExpression: parameter foo2 : *Foo
interpretDynamically: &foo1.a [#0]
interpretDynamically: *t1 = 2:int
resolveExpression: 2:int
resolveExpression: parameter foo1 : *Foo
interpretDynamically: &foo2.a [#0]
interpretDynamically: *t2
interpretDynamically: t3 == 2:int
interpretDynamically: if t4 goto 1 else 2
resolveExpression: t3 == 2:int
resolveExpression: *t2
resolveExpression: parameter foo2 : *Foo
resolveExpression: 2:int
interpretDynamically: return 5:int
resolveExpression: 5:int
interpretDynamically: return 4:int
resolveExpression: 4:int
States: 2
---- State #0 ----
Interpreter{
  PC: (true & (Foo_0_sel2 == 2))
  ReturnResults:
    [0] 4
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=Foo, field=-1)
      P[1]: &(addr_1) (refType=struct, type=Foo, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      Foo:
        field[0]: Foo_0 (elem=int, ops=3, store=2, select=1)
        op[0]: store idx=addr_0, val=5
        op[1]: store idx=addr_1, val=2
        op[2]: select idx=addr_0, res=Foo_0_sel2
  }
  HeapConstraints: 6
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: Foo_0.store(addr_0, 5).store(addr_1, 2).select(addr_0, Foo_0_sel2)
}

SAT
Model:
Foo_0 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
addr_0 -> #x0000000000000000
addr_1 -> #x0000000000000000
Foo_0_sel2 -> #x0000000000000002

---- State #1 ----
Interpreter{
  PC: (true & (!(Foo_0_sel2 == 2)))
  ReturnResults:
    [0] 5
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=Foo, field=-1)
      P[1]: &(addr_1) (refType=struct, type=Foo, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      Foo:
        field[0]: Foo_0 (elem=int, ops=3, store=2, select=1)
        op[0]: store idx=addr_0, val=5
        op[1]: store idx=addr_1, val=2
        op[2]: select idx=addr_0, res=Foo_0_sel2
  }
  HeapConstraints: 6
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: Foo_0.store(addr_0, 5).store(addr_1, 2).select(addr_0, Foo_0_sel2)
}

SAT
Model:
Foo_0 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
addr_0 -> #x0000000000000001
addr_1 -> #x0000000000000000
Foo_0_sel2 -> #x0000000000000005


==================== ArrayAliasing ====================
CFG построен для функции с 3 блоками
interpretDynamically: &arr2[1:int]
interpretDynamically: *t0 = 5:int
resolveExpression: 5:int
resolveExpression: parameter arr2 : []int
resolveExpression: 1:int
interpretDynamically: &arr1[1:int]
interpretDynamically: *t1 = 2:int
resolveExpression: 2:int
resolveExpression: parameter arr1 : []int
resolveExpression: 1:int
interpretDynamically: &arr2[1:int]
interpretDynamically: *t2
interpretDynamically: t3 == 2:int
interpretDynamically: if t4 goto 1 else 2
resolveExpression: t3 == 2:int
resolveExpression: *t2
resolveExpression: parameter arr2 : []int
resolveExpression: 1:int
resolveExpression: 2:int
interpretDynamically: return 4:int
resolveExpression: 4:int
interpretDynamically: return 5:int
resolveExpression: 5:int
States: 2
---- State #0 ----
Interpreter{
  PC: (true & (IntSegment_sel2 == 2))
  ReturnResults:
    [0] 4
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=array, dim=0, base=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    IntSegment:    IntSegment (elem=int, ops=3, store=2, select=1)
        op[0]: store idx=(addr_0 + 1), val=5
        op[1]: store idx=(addr_1 + 1), val=2
        op[2]: select idx=(addr_0 + 1), res=IntSegment_sel2
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.store((addr_0 + 1), 5).store((addr_1 + 1), 2).select((addr_0 + 1), IntSegment_sel2)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
addr_0 -> #xffffffffffffffff
addr_1 -> #xffffffffffffffff
IntSegment_sel2 -> #x0000000000000002

---- State #1 ----
Interpreter{
  PC: (true & (!(IntSegment_sel2 == 2)))
  ReturnResults:
    [0] 5
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=array, dim=0, base=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    IntSegment:    IntSegment (elem=int, ops=3, store=2, select=1)
        op[0]: store idx=(addr_0 + 1), val=5
        op[1]: store idx=(addr_1 + 1), val=2
        op[2]: select idx=(addr_0 + 1), res=IntSegment_sel2
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.store((addr_0 + 1), 5).store((addr_1 + 1), 2).select((addr_0 + 1), IntSegment_sel2)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
addr_0 -> #xffffffffffffffff
addr_1 -> #x0000000000000000
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
IntSegment_sel2 -> #x0000000000000005


==================== DefaultBooleanValues ====================
CFG построен для функции с 3 блоками
interpretDynamically: new [3]bool (makeslice)
interpretDynamically: slice t0[:3:int]
interpretDynamically: &t1[1:int]
interpretDynamically: *t2
interpretDynamically: if t3 goto 2 else 1
resolveExpression: *t2
resolveExpression: slice t0[:3:int]
resolveExpression: new [3]bool (makeslice)
resolveExpression: 1:int
interpretDynamically: return t1
resolveExpression: slice t0[:3:int]
interpretDynamically: return t1
resolveExpression: slice t0[:3:int]
States: 2
---- State #0 ----
Interpreter{
  PC: (true & (!BoolSegment_sel0))
  ReturnResults:
    [0] &(addr_0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 0, Allocations: 1
    AllocationsList:
      A[0]: &(addr_0) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=1, store=0, select=1)
        op[0]: select idx=(addr_0 + 1), res=BoolSegment_sel0
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment.select((addr_0 + 1), BoolSegment_sel0)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
BoolSegment -> ((as const (Array (_ BitVec 64) Bool)) false)
addr_0 -> #xffffffffffffffff
BoolSegment_sel0 -> false

---- State #1 ----
Interpreter{
  PC: (true & BoolSegment_sel0)
  ReturnResults:
    [0] &(addr_0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 0, Allocations: 1
    AllocationsList:
      A[0]: &(addr_0) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=1, store=0, select=1)
        op[0]: select idx=(addr_0 + 1), res=BoolSegment_sel0
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment.select((addr_0 + 1), BoolSegment_sel0)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
BoolSegment -> ((as const (Array (_ BitVec 64) Bool)) true)
addr_0 -> #xffffffffffffffff
BoolSegment_sel0 -> true


==================== ByteArray ====================
CFG построен для функции с 5 блоками
interpretDynamically: len_a != 2:int
interpretDynamically: if t0 goto 1 else 2
resolveExpression: len_a != 2:int
resolveExpression: parameter len_a : int
resolveExpression: 2:int
interpretDynamically: return 255:byte
resolveExpression: 255:byte
interpretDynamically: &a[0:int]
interpretDynamically: *t1 = 5:byte
resolveExpression: 5:byte
resolveExpression: parameter a : []byte
resolveExpression: 0:int
interpretDynamically: &a[1:int]
interpretDynamically: *t2 = x
resolveExpression: parameter x : byte
resolveExpression: parameter a : []byte
resolveExpression: 1:int
interpretDynamically: &a[0:int]
interpretDynamically: *t3
interpretDynamically: &a[1:int]
interpretDynamically: *t5
interpretDynamically: t4 + t6
interpretDynamically: t7 > 20:byte
interpretDynamically: if t8 goto 3 else 4
resolveExpression: t7 > 20:byte
resolveExpression: t4 + t6
resolveExpression: *t3
resolveExpression: parameter a : []byte
resolveExpression: 0:int
resolveExpression: *t5
resolveExpression: parameter a : []byte
resolveExpression: 1:int
resolveExpression: 20:byte
interpretDynamically: return 0:byte
resolveExpression: 0:byte
interpretDynamically: return 1:byte
resolveExpression: 1:byte
States: 3
---- State #0 ----
Interpreter{
  PC: (true & (IntSegment_sel0 != 2))
  ReturnResults:
    [0] 255
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
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #xfffffffffffffffd)
addr_0 -> #x0000000000000000
IntSegment_sel0 -> #xfffffffffffffffd

---- State #1 ----
Interpreter{
  PC: ((true & (!(IntSegment_sel0 != 2))) & (!((IntSegment_sel4 + IntSegment_sel5) > 20)))
  ReturnResults:
    [0] 0
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 3, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
      P[2]: &(addr_2) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=6, store=2, select=4)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: store idx=(addr_1 + 0), val=5
        op[2]: select idx=addr_2, res=IntSegment_sel2
        op[3]: store idx=(addr_1 + 1), val=IntSegment_sel2
        op[4]: select idx=(addr_1 + 0), res=IntSegment_sel4
        op[5]: select idx=(addr_1 + 1), res=IntSegment_sel5
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).store((addr_1 + 0), 5).select(addr_2, IntSegment_sel2).store((addr_1 + 1), IntSegment_sel2).select((addr_1 + 0), IntSegment_sel4).select((addr_1 + 1), IntSegment_sel5)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
addr_2 -> #x2000000000000000
addr_1 -> #x0000000000000000
addr_0 -> #x0000000000000000
IntSegment -> (store (store ((as const (Array (_ BitVec 64) (_ BitVec 64)))
                              #x0000000000000005)
                            #x0000000000000000
                            #x0000000000000002)
                     #x2000000000000000
                     #xfffffffffffffff9)
IntSegment_sel2 -> #xfffffffffffffff9
IntSegment_sel5 -> #xfffffffffffffff9
IntSegment_sel4 -> #x0000000000000005
IntSegment_sel0 -> #x0000000000000002

---- State #2 ----
Interpreter{
  PC: ((true & (!(IntSegment_sel0 != 2))) & ((IntSegment_sel4 + IntSegment_sel5) > 20))
  ReturnResults:
    [0] 1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 3, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
      P[2]: &(addr_2) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=6, store=2, select=4)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: store idx=(addr_1 + 0), val=5
        op[2]: select idx=addr_2, res=IntSegment_sel2
        op[3]: store idx=(addr_1 + 1), val=IntSegment_sel2
        op[4]: select idx=(addr_1 + 0), res=IntSegment_sel4
        op[5]: select idx=(addr_1 + 1), res=IntSegment_sel5
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).store((addr_1 + 0), 5).select(addr_2, IntSegment_sel2).store((addr_1 + 1), IntSegment_sel2).select((addr_1 + 0), IntSegment_sel4).select((addr_1 + 1), IntSegment_sel5)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
addr_2 -> #x2000000000000000
addr_1 -> #x0000000000000000
addr_0 -> #x0000000000000000
IntSegment -> (store (store ((as const (Array (_ BitVec 64) (_ BitVec 64)))
                              #x0000000000000005)
                            #x0000000000000000
                            #x0000000000000002)
                     #x2000000000000000
                     #x7ffffffffffffff9)
IntSegment_sel2 -> #x7ffffffffffffff9
IntSegment_sel5 -> #x7ffffffffffffff9
IntSegment_sel4 -> #x0000000000000005
IntSegment_sel0 -> #x0000000000000002


==================== CharSizeAndIndex ====================
CFG построен для функции с 7 блоками
interpretDynamically: a == nil:[]rune
interpretDynamically: if t0 goto 1 else 4
resolveExpression: a == nil:[]rune
resolveExpression: parameter a : []rune
resolveExpression: nil:[]rune
interpretDynamically: return 255:byte
resolveExpression: 255:byte
interpretDynamically: convert int <- rune (x)
interpretDynamically: len_a <= t11
interpretDynamically: if t12 goto 1 else 3
resolveExpression: len_a <= t11
resolveExpression: parameter len_a : int
resolveExpression: convert int <- rune (x)
resolveExpression: parameter x : rune
interpretDynamically: x < 1:rune
interpretDynamically: return 255:byte
resolveExpression: 255:byte
interpretDynamically: if t10 goto 1 else 2
resolveExpression: x < 1:rune
resolveExpression: parameter x : rune
resolveExpression: 1:rune
interpretDynamically: make []rune x x
interpretDynamically: &t1[0:int]
interpretDynamically: *t2 = 5:rune
resolveExpression: 5:rune
resolveExpression: make []rune x x
resolveExpression: 0:int
interpretDynamically: &a[x]
interpretDynamically: *t3 = x
resolveExpression: parameter x : rune
resolveExpression: parameter a : []rune
resolveExpression: parameter x : rune
interpretDynamically: &t1[0:int]
interpretDynamically: return 255:byte
resolveExpression: 255:byte
interpretDynamically: *t4
interpretDynamically: &a[x]
interpretDynamically: *t6
interpretDynamically: t5 + t7
interpretDynamically: t8 > 7:rune
interpretDynamically: if t9 goto 5 else 6
resolveExpression: t8 > 7:rune
resolveExpression: t5 + t7
resolveExpression: *t4
resolveExpression: make []rune x x
resolveExpression: 0:int
resolveExpression: *t6
resolveExpression: parameter a : []rune
resolveExpression: parameter x : rune
resolveExpression: 7:rune
interpretDynamically: return 0:byte
resolveExpression: 0:byte
interpretDynamically: return 1:byte
resolveExpression: 1:byte
States: 5
---- State #0 ----
Interpreter{
  PC: (true & (addr_0 == 0))
  ReturnResults:
    [0] 255
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=array, dim=0, base=int)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
addr_0 -> #x0000000000000000

---- State #1 ----
Interpreter{
  PC: ((true & (!(addr_0 == 0))) & (IntSegment_sel0 <= IntSegment_sel1))
  ReturnResults:
    [0] 255
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 3, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=array, dim=0, base=int)
      P[1]: &(addr_1) (refType=int)
      P[2]: &(addr_2) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=2, store=0, select=2)
        op[0]: select idx=addr_1, res=IntSegment_sel0
        op[1]: select idx=addr_2, res=IntSegment_sel1
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_1, IntSegment_sel0).select(addr_2, IntSegment_sel1)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
addr_1 -> #xffffffffffffffff
addr_2 -> #x0000000000000000
addr_0 -> #xffffffffffffffff
IntSegment -> (store ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x8000000000000000)
                     #x0000000000000000
                     #x7fffffffffffffff)
IntSegment_sel1 -> #x7fffffffffffffff
IntSegment_sel0 -> #x8000000000000000

---- State #2 ----
Interpreter{
  PC: ((((true & (!(addr_0 == 0))) & (!(IntSegment_sel0 <= IntSegment_sel1))) & (!(IntSegment_sel2 < 1))) & (!((IntSegment_sel7 + IntSegment_sel9) > 7)))
  ReturnResults:
    [0] 0
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 4
    Params: 3, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=array, dim=0, base=int)
      P[1]: &(addr_1) (refType=int)
      P[2]: &(addr_2) (refType=int)
    AllocationsList:
      A[0]: &(addr_3) (refType=array, dim=0, base=int)
    IntSegment:    IntSegment (elem=int, ops=10, store=2, select=8)
        op[0]: select idx=addr_1, res=IntSegment_sel0
        op[1]: select idx=addr_2, res=IntSegment_sel1
        op[2]: select idx=addr_2, res=IntSegment_sel2
        op[3]: store idx=(addr_3 + 0), val=5
        op[4]: select idx=addr_2, res=IntSegment_sel4
        op[5]: select idx=addr_2, res=IntSegment_sel5
        op[6]: store idx=(addr_0 + IntSegment_sel5), val=IntSegment_sel4
        op[7]: select idx=(addr_3 + 0), res=IntSegment_sel7
        op[8]: select idx=addr_2, res=IntSegment_sel8
        op[9]: select idx=(addr_0 + IntSegment_sel8), res=IntSegment_sel9
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 8
    C[0]: (addr_0 != addr_3)
    C[1]: (addr_1 != addr_3)
    C[2]: (addr_2 != addr_3)
    C[3]: IntSegment.select(addr_1, IntSegment_sel0).select(addr_2, IntSegment_sel1).select(addr_2, IntSegment_sel2).store((addr_3 + 0), 5).select(addr_2, IntSegment_sel4).select(addr_2, IntSegment_sel5).store((addr_0 + IntSegment_sel5), IntSegment_sel4).select((addr_3 + 0), IntSegment_sel7).select(addr_2, IntSegment_sel8).select((addr_0 + IntSegment_sel8), IntSegment_sel9)
    C[4]: BoolSegment
    C[5]: StringSegment
    C[6]: FloatSegment
    C[7]: ArraySegment
}

SAT
Model:
addr_2 -> #x0020000000000000
addr_3 -> #x4000000000000001
addr_0 -> #x0000000000000001
addr_1 -> #x0000000000000000
IntSegment -> (store (store ((as const (Array (_ BitVec 64) (_ BitVec 64)))
                              #x0000000000000000)
                            #x0000000000000000
                            #x4000000000000001)
                     #x0020000000000000
                     #x4000000000000000)
IntSegment_sel1 -> #x4000000000000000
IntSegment_sel4 -> #x4000000000000000
IntSegment_sel7 -> #x4000000000000000
IntSegment_sel8 -> #x4000000000000000
IntSegment_sel2 -> #x4000000000000000
IntSegment_sel5 -> #x4000000000000000
IntSegment_sel9 -> #x4000000000000000
IntSegment_sel0 -> #x4000000000000001

---- State #3 ----
Interpreter{
  PC: (((true & (!(addr_0 == 0))) & (!(IntSegment_sel0 <= IntSegment_sel1))) & (IntSegment_sel2 < 1))
  ReturnResults:
    [0] 255
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 3, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=array, dim=0, base=int)
      P[1]: &(addr_1) (refType=int)
      P[2]: &(addr_2) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=3, store=0, select=3)
        op[0]: select idx=addr_1, res=IntSegment_sel0
        op[1]: select idx=addr_2, res=IntSegment_sel1
        op[2]: select idx=addr_2, res=IntSegment_sel2
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_1, IntSegment_sel0).select(addr_2, IntSegment_sel1).select(addr_2, IntSegment_sel2)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
addr_1 -> #xffffffffffffffff
addr_2 -> #x0000000000000000
addr_0 -> #xffffffffffffffff
IntSegment -> (store ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x555555aaaa800000)
                     #x0000000000000000
                     #x80000000007fffff)
IntSegment_sel2 -> #x80000000007fffff
IntSegment_sel1 -> #x80000000007fffff
IntSegment_sel0 -> #x555555aaaa800000

---- State #4 ----
Interpreter{
  PC: ((((true & (!(addr_0 == 0))) & (!(IntSegment_sel0 <= IntSegment_sel1))) & (!(IntSegment_sel2 < 1))) & ((IntSegment_sel7 + IntSegment_sel9) > 7))
  ReturnResults:
    [0] 1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 4
    Params: 3, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=array, dim=0, base=int)
      P[1]: &(addr_1) (refType=int)
      P[2]: &(addr_2) (refType=int)
    AllocationsList:
      A[0]: &(addr_3) (refType=array, dim=0, base=int)
    IntSegment:    IntSegment (elem=int, ops=10, store=2, select=8)
        op[0]: select idx=addr_1, res=IntSegment_sel0
        op[1]: select idx=addr_2, res=IntSegment_sel1
        op[2]: select idx=addr_2, res=IntSegment_sel2
        op[3]: store idx=(addr_3 + 0), val=5
        op[4]: select idx=addr_2, res=IntSegment_sel4
        op[5]: select idx=addr_2, res=IntSegment_sel5
        op[6]: store idx=(addr_0 + IntSegment_sel5), val=IntSegment_sel4
        op[7]: select idx=(addr_3 + 0), res=IntSegment_sel7
        op[8]: select idx=addr_2, res=IntSegment_sel8
        op[9]: select idx=(addr_0 + IntSegment_sel8), res=IntSegment_sel9
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 8
    C[0]: (addr_0 != addr_3)
    C[1]: (addr_1 != addr_3)
    C[2]: (addr_2 != addr_3)
    C[3]: IntSegment.select(addr_1, IntSegment_sel0).select(addr_2, IntSegment_sel1).select(addr_2, IntSegment_sel2).store((addr_3 + 0), 5).select(addr_2, IntSegment_sel4).select(addr_2, IntSegment_sel5).store((addr_0 + IntSegment_sel5), IntSegment_sel4).select((addr_3 + 0), IntSegment_sel7).select(addr_2, IntSegment_sel8).select((addr_0 + IntSegment_sel8), IntSegment_sel9)
    C[4]: BoolSegment
    C[5]: StringSegment
    C[6]: FloatSegment
    C[7]: ArraySegment
}

SAT
Model:
addr_2 -> #x0000000000000000
addr_3 -> #x0000000010000000
addr_0 -> #x0000000000800001
addr_1 -> #x0080000000000000
IntSegment -> (store (store ((as const (Array (_ BitVec 64) (_ BitVec 64)))
                              #x0000000000000000)
                            #x0080000000000000
                            #x4000000001800001)
                     #x0000000000000000
                     #x4000000001800000)
IntSegment_sel1 -> #x4000000001800000
IntSegment_sel4 -> #x4000000001800000
IntSegment_sel7 -> #x0000000000000005
IntSegment_sel8 -> #x4000000001800000
IntSegment_sel2 -> #x4000000001800000
IntSegment_sel5 -> #x4000000001800000
IntSegment_sel9 -> #x4000000001800000
IntSegment_sel0 -> #x4000000001800001


==================== BooleanArray ====================
CFG построен для функции с 5 блоками
interpretDynamically: len_arr == 0:int
interpretDynamically: if t0 goto 1 else 2
resolveExpression: len_arr == 0:int
resolveExpression: parameter len_arr : int
resolveExpression: 0:int
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: &arr[0:int]
interpretDynamically: *t1
interpretDynamically: if t2 goto 3 else 4
resolveExpression: *t1
resolveExpression: parameter arr : []bool
resolveExpression: 0:int
interpretDynamically: &arr[0:int]
interpretDynamically: return 2:int
resolveExpression: 2:int
interpretDynamically: *t3 = true:bool
resolveExpression: true:bool
resolveExpression: parameter arr : []bool
resolveExpression: 0:int
interpretDynamically: return 3:int
resolveExpression: 3:int
States: 3
---- State #0 ----
Interpreter{
  PC: (true & (IntSegment_sel0 == 0))
  ReturnResults:
    [0] 1
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
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
IntSegment_sel0 -> #x0000000000000000

---- State #1 ----
Interpreter{
  PC: ((true & (!(IntSegment_sel0 == 0))) & BoolSegment_sel0)
  ReturnResults:
    [0] 2
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=bool)
    IntSegment:    IntSegment (elem=int, ops=1, store=0, select=1)
        op[0]: select idx=addr_0, res=IntSegment_sel0
    BoolSegment:   BoolSegment (elem=bool, ops=1, store=0, select=1)
        op[0]: select idx=(addr_1 + 0), res=BoolSegment_sel0
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0)
    C[1]: BoolSegment.select((addr_1 + 0), BoolSegment_sel0)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #xffffffffffffffff)
BoolSegment -> ((as const (Array (_ BitVec 64) Bool)) true)
BoolSegment_sel0 -> true
addr_0 -> #x0000000000000000
IntSegment_sel0 -> #xffffffffffffffff

---- State #2 ----
Interpreter{
  PC: ((true & (!(IntSegment_sel0 == 0))) & (!BoolSegment_sel0))
  ReturnResults:
    [0] 3
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=bool)
    IntSegment:    IntSegment (elem=int, ops=1, store=0, select=1)
        op[0]: select idx=addr_0, res=IntSegment_sel0
    BoolSegment:   BoolSegment (elem=bool, ops=2, store=1, select=1)
        op[0]: select idx=(addr_1 + 0), res=BoolSegment_sel0
        op[1]: store idx=(addr_1 + 0), val=true
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0)
    C[1]: BoolSegment.select((addr_1 + 0), BoolSegment_sel0).store((addr_1 + 0), true)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #xffffffffffffffff)
BoolSegment -> ((as const (Array (_ BitVec 64) Bool)) false)
BoolSegment_sel0 -> false
addr_0 -> #x0000000000000000
IntSegment_sel0 -> #xffffffffffffffff


==================== CreateArray ====================
CFG построен для функции с 6 блоками
interpretDynamically: length < 3:int
interpretDynamically: if t0 goto 1 else 2
resolveExpression: length < 3:int
resolveExpression: parameter length : int
resolveExpression: 3:int
interpretDynamically: make []*ObjectWithPrimitivesClass length length
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 3: t12] #i
resolveExpression: 0:int
interpretDynamically: t13 < length
interpretDynamically: if t14 goto 3 else 4
resolveExpression: t13 < length
resolveExpression: phi [2: 0:int, 3: t12] #i
resolveExpression: parameter length : int
interpretDynamically: return t1
resolveExpression: make []*ObjectWithPrimitivesClass length length
interpretDynamically: new ObjectWithPrimitivesClass (complit)
interpretDynamically: &t1[t13]
interpretDynamically: *t3 = t2
resolveExpression: new ObjectWithPrimitivesClass (complit)
resolveExpression: make []*ObjectWithPrimitivesClass length length
resolveExpression: phi [2: 0:int, 3: t12] #i
interpretDynamically: return nil:[]*ObjectWithPrimitivesClass
resolveExpression: nil:[]*main.ObjectWithPrimitivesClass
interpretDynamically: &t1[t13]
interpretDynamically: *t4
interpretDynamically: x + t13
interpretDynamically: &t5.x [#1]
interpretDynamically: *t7 = t6
resolveExpression: x + t13
resolveExpression: parameter x : int
resolveExpression: phi [2: 0:int, 3: t12] #i
resolveExpression: *t4
resolveExpression: make []*ObjectWithPrimitivesClass length length
resolveExpression: phi [2: 0:int, 3: t12] #i
interpretDynamically: &t1[t13]
interpretDynamically: *t8
interpretDynamically: y + t13
interpretDynamically: &t9.y [#2]
interpretDynamically: *t11 = t10
resolveExpression: y + t13
resolveExpression: parameter y : int
resolveExpression: phi [2: 0:int, 3: t12] #i
resolveExpression: *t8
resolveExpression: make []*ObjectWithPrimitivesClass length length
resolveExpression: phi [2: 0:int, 3: t12] #i
interpretDynamically: t13 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 3: t12] #i
resolveExpression: t13 + 1:int
resolveExpression: phi [2: 0:int, 3: t12] #i
resolveExpression: 1:int
interpretDynamically: t13 < length
interpretDynamically: if t14 goto 3 else 4
resolveExpression: t13 < length
resolveExpression: phi [2: 0:int, 3: t12] #i
resolveExpression: parameter length : int
interpretDynamically: new ObjectWithPrimitivesClass (complit)
interpretDynamically: &t1[t13]
interpretDynamically: *t3 = t2
resolveExpression: new ObjectWithPrimitivesClass (complit)
resolveExpression: make []*ObjectWithPrimitivesClass length length
resolveExpression: phi [2: 0:int, 3: t12] #i
interpretDynamically: &t1[t13]
interpretDynamically: *t4
interpretDynamically: x + t13
interpretDynamically: &t5.x [#1]
interpretDynamically: *t7 = t6
resolveExpression: x + t13
resolveExpression: parameter x : int
resolveExpression: phi [2: 0:int, 3: t12] #i
resolveExpression: *t4
resolveExpression: make []*ObjectWithPrimitivesClass length length
resolveExpression: phi [2: 0:int, 3: t12] #i
interpretDynamically: &t1[t13]
interpretDynamically: *t8
interpretDynamically: y + t13
interpretDynamically: &t9.y [#2]
interpretDynamically: return t1
resolveExpression: make []*ObjectWithPrimitivesClass length length
interpretDynamically: *t11 = t10
resolveExpression: y + t13
resolveExpression: parameter y : int
resolveExpression: phi [2: 0:int, 3: t12] #i
resolveExpression: *t8
resolveExpression: make []*ObjectWithPrimitivesClass length length
resolveExpression: phi [2: 0:int, 3: t12] #i
interpretDynamically: t13 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 3: t12] #i
resolveExpression: t13 + 1:int
resolveExpression: phi [2: 0:int, 3: t12] #i
resolveExpression: 1:int
interpretDynamically: t13 < length
interpretDynamically: if t14 goto 3 else 4
resolveExpression: t13 < length
resolveExpression: phi [2: 0:int, 3: t12] #i
resolveExpression: parameter length : int
interpretDynamically: return t1
resolveExpression: make []*ObjectWithPrimitivesClass length length
interpretDynamically: new ObjectWithPrimitivesClass (complit)
interpretDynamically: &t1[t13]
interpretDynamically: *t3 = t2
resolveExpression: new ObjectWithPrimitivesClass (complit)
resolveExpression: make []*ObjectWithPrimitivesClass length length
resolveExpression: phi [2: 0:int, 3: t12] #i
interpretDynamically: &t1[t13]
interpretDynamically: *t4
interpretDynamically: x + t13
interpretDynamically: &t5.x [#1]
interpretDynamically: *t7 = t6
resolveExpression: x + t13
resolveExpression: parameter x : int
resolveExpression: phi [2: 0:int, 3: t12] #i
resolveExpression: *t4
resolveExpression: make []*ObjectWithPrimitivesClass length length
resolveExpression: phi [2: 0:int, 3: t12] #i
interpretDynamically: &t1[t13]
interpretDynamically: *t8
interpretDynamically: y + t13
interpretDynamically: &t9.y [#2]
interpretDynamically: *t11 = t10
resolveExpression: y + t13
resolveExpression: parameter y : int
resolveExpression: phi [2: 0:int, 3: t12] #i
resolveExpression: *t8
resolveExpression: make []*ObjectWithPrimitivesClass length length
resolveExpression: phi [2: 0:int, 3: t12] #i
interpretDynamically: t13 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 3: t12] #i
resolveExpression: t13 + 1:int
resolveExpression: phi [2: 0:int, 3: t12] #i
resolveExpression: 1:int
interpretDynamically: t13 < length
interpretDynamically: if t14 goto 3 else 4
resolveExpression: t13 < length
resolveExpression: phi [2: 0:int, 3: t12] #i
resolveExpression: parameter length : int
interpretDynamically: return t1
resolveExpression: make []*ObjectWithPrimitivesClass length length
interpretDynamically: new ObjectWithPrimitivesClass (complit)
interpretDynamically: &t1[t13]
interpretDynamically: *t3 = t2
resolveExpression: new ObjectWithPrimitivesClass (complit)
resolveExpression: make []*ObjectWithPrimitivesClass length length
resolveExpression: phi [2: 0:int, 3: t12] #i
interpretDynamically: &t1[t13]
interpretDynamically: *t4
interpretDynamically: x + t13
interpretDynamically: &t5.x [#1]
interpretDynamically: *t7 = t6
resolveExpression: x + t13
resolveExpression: parameter x : int
resolveExpression: phi [2: 0:int, 3: t12] #i
resolveExpression: *t4
resolveExpression: make []*ObjectWithPrimitivesClass length length
resolveExpression: phi [2: 0:int, 3: t12] #i
interpretDynamically: &t1[t13]
interpretDynamically: *t8
interpretDynamically: y + t13
interpretDynamically: &t9.y [#2]
interpretDynamically: *t11 = t10
resolveExpression: y + t13
resolveExpression: parameter y : int
resolveExpression: phi [2: 0:int, 3: t12] #i
resolveExpression: *t8
resolveExpression: make []*ObjectWithPrimitivesClass length length
resolveExpression: phi [2: 0:int, 3: t12] #i
interpretDynamically: t13 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 3: t12] #i
resolveExpression: t13 + 1:int
resolveExpression: phi [2: 0:int, 3: t12] #i
resolveExpression: 1:int
interpretDynamically: t13 < length
interpretDynamically: if t14 goto 3 else 4
resolveExpression: t13 < length
resolveExpression: phi [2: 0:int, 3: t12] #i
resolveExpression: parameter length : int
interpretDynamically: new ObjectWithPrimitivesClass (complit)
interpretDynamically: &t1[t13]
interpretDynamically: *t3 = t2
resolveExpression: new ObjectWithPrimitivesClass (complit)
resolveExpression: make []*ObjectWithPrimitivesClass length length
resolveExpression: phi [2: 0:int, 3: t12] #i
interpretDynamically: return t1
resolveExpression: make []*ObjectWithPrimitivesClass length length
interpretDynamically: &t1[t13]
interpretDynamically: *t4
interpretDynamically: x + t13
interpretDynamically: &t5.x [#1]
interpretDynamically: *t7 = t6
resolveExpression: x + t13
resolveExpression: parameter x : int
resolveExpression: phi [2: 0:int, 3: t12] #i
resolveExpression: *t4
resolveExpression: make []*ObjectWithPrimitivesClass length length
resolveExpression: phi [2: 0:int, 3: t12] #i
interpretDynamically: &t1[t13]
interpretDynamically: *t8
interpretDynamically: y + t13
interpretDynamically: &t9.y [#2]
interpretDynamically: *t11 = t10
resolveExpression: y + t13
resolveExpression: parameter y : int
resolveExpression: phi [2: 0:int, 3: t12] #i
resolveExpression: *t8
resolveExpression: make []*ObjectWithPrimitivesClass length length
resolveExpression: phi [2: 0:int, 3: t12] #i
interpretDynamically: t13 + 1:int
interpretDynamically: jump 5
States: 6
---- State #0 ----
Interpreter{
  PC: ((true & (!(IntSegment_sel0 < 3))) & (!(0 < IntSegment_sel1)))
  ReturnResults:
    [0] &(addr_1)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 1, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    AllocationsList:
      A[0]: &(addr_1) (refType=array, dim=0, base=struct(ObjectWithPrimitivesClass))
    IntSegment:    IntSegment (elem=int, ops=2, store=0, select=2)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=0, store=0, select=0)
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=0, store=0, select=0)
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=0, store=0, select=0)
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=0, store=0, select=0)
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=0, store=0, select=0)
  }
  HeapConstraints: 11
    C[0]: (addr_0 != addr_1)
    C[1]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1)
    C[2]: BoolSegment
    C[3]: StringSegment
    C[4]: FloatSegment
    C[5]: ArraySegment
    C[6]: ObjectWithPrimitivesClass_3
    C[7]: ObjectWithPrimitivesClass_4
    C[8]: ObjectWithPrimitivesClass_0
    C[9]: ObjectWithPrimitivesClass_1
    C[10]: ObjectWithPrimitivesClass_2
}

UNSAT
---- State #1 ----
Interpreter{
  PC: (true & (IntSegment_sel0 < 3))
  ReturnResults:
    [0] &(0)
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
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000002)
addr_0 -> #x0000000000000000
IntSegment_sel0 -> #x0000000000000002

---- State #2 ----
Interpreter{
  PC: (((true & (!(IntSegment_sel0 < 3))) & (0 < IntSegment_sel1)) & (!((0 + 1) < IntSegment_sel4)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 5
    Params: 3, Allocations: 2
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_3) (refType=int)
      P[2]: &(addr_4) (refType=int)
    AllocationsList:
      A[0]: &(addr_1) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
      A[1]: &(addr_2) (refType=array, dim=0, base=struct(ObjectWithPrimitivesClass))
    IntSegment:    IntSegment (elem=int, ops=5, store=0, select=5)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_3, res=IntSegment_sel2
        op[3]: select idx=addr_4, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=3, store=0, select=3)
        op[0]: select idx=(addr_2 + 0), res=ArraySegment_sel0
        op[1]: select idx=(addr_2 + 0), res=ArraySegment_sel1
        op[2]: select idx=(addr_2 + 0), res=ArraySegment_sel2
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=2, store=1, select=1)
        op[0]: select idx=addr_1, res=ObjectWithPrimitivesClass_0_sel0
        op[1]: store idx=ArraySegment_sel0, val=ObjectWithPrimitivesClass_0_sel0
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=3, store=2, select=1)
        op[0]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel0
        op[1]: store idx=ArraySegment_sel0, val=ObjectWithPrimitivesClass_1_sel0
        op[2]: store idx=ArraySegment_sel1, val=(IntSegment_sel2 + 0)
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=3, store=2, select=1)
        op[0]: select idx=addr_1, res=ObjectWithPrimitivesClass_2_sel0
        op[1]: store idx=ArraySegment_sel0, val=ObjectWithPrimitivesClass_2_sel0
        op[2]: store idx=ArraySegment_sel2, val=(IntSegment_sel3 + 0)
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=2, store=1, select=1)
        op[0]: select idx=addr_1, res=ObjectWithPrimitivesClass_3_sel0
        op[1]: store idx=ArraySegment_sel0, val=ObjectWithPrimitivesClass_3_sel0
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=2, store=1, select=1)
        op[0]: select idx=addr_1, res=ObjectWithPrimitivesClass_4_sel0
        op[1]: store idx=ArraySegment_sel0, val=ObjectWithPrimitivesClass_4_sel0
  }
  HeapConstraints: 17
    C[0]: (addr_0 != addr_1)
    C[1]: (addr_0 != addr_2)
    C[2]: (addr_3 != addr_1)
    C[3]: (addr_3 != addr_2)
    C[4]: (addr_4 != addr_1)
    C[5]: (addr_4 != addr_2)
    C[6]: (addr_1 != addr_2)
    C[7]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_3, IntSegment_sel2).select(addr_4, IntSegment_sel3).select(addr_0, IntSegment_sel4)
    C[8]: BoolSegment
    C[9]: StringSegment
    C[10]: FloatSegment
    C[11]: ArraySegment.select((addr_2 + 0), ArraySegment_sel0).select((addr_2 + 0), ArraySegment_sel1).select((addr_2 + 0), ArraySegment_sel2)
    C[12]: ObjectWithPrimitivesClass_4.select(addr_1, ObjectWithPrimitivesClass_4_sel0).store(ArraySegment_sel0, ObjectWithPrimitivesClass_4_sel0)
    C[13]: ObjectWithPrimitivesClass_0.select(addr_1, ObjectWithPrimitivesClass_0_sel0).store(ArraySegment_sel0, ObjectWithPrimitivesClass_0_sel0)
    C[14]: ObjectWithPrimitivesClass_1.select(addr_1, ObjectWithPrimitivesClass_1_sel0).store(ArraySegment_sel0, ObjectWithPrimitivesClass_1_sel0).store(ArraySegment_sel1, (IntSegment_sel2 + 0))
    C[15]: ObjectWithPrimitivesClass_2.select(addr_1, ObjectWithPrimitivesClass_2_sel0).store(ArraySegment_sel0, ObjectWithPrimitivesClass_2_sel0).store(ArraySegment_sel2, (IntSegment_sel3 + 0))
    C[16]: ObjectWithPrimitivesClass_3.select(addr_1, ObjectWithPrimitivesClass_3_sel0).store(ArraySegment_sel0, ObjectWithPrimitivesClass_3_sel0)
}

UNSAT
---- State #3 ----
Interpreter{
  PC: ((((true & (!(IntSegment_sel0 < 3))) & (0 < IntSegment_sel1)) & ((0 + 1) < IntSegment_sel4)) & (!(((0 + 1) + 1) < IntSegment_sel7)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 5
    Params: 3, Allocations: 2
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_3) (refType=int)
      P[2]: &(addr_4) (refType=int)
    AllocationsList:
      A[0]: &(addr_1) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
      A[1]: &(addr_2) (refType=array, dim=0, base=struct(ObjectWithPrimitivesClass))
    IntSegment:    IntSegment (elem=int, ops=8, store=0, select=8)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_3, res=IntSegment_sel2
        op[3]: select idx=addr_4, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_3, res=IntSegment_sel5
        op[6]: select idx=addr_4, res=IntSegment_sel6
        op[7]: select idx=addr_0, res=IntSegment_sel7
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=6, store=0, select=6)
        op[0]: select idx=(addr_2 + 0), res=ArraySegment_sel0
        op[1]: select idx=(addr_2 + 0), res=ArraySegment_sel1
        op[2]: select idx=(addr_2 + 0), res=ArraySegment_sel2
        op[3]: select idx=(addr_2 + (0 + 1)), res=ArraySegment_sel3
        op[4]: select idx=(addr_2 + (0 + 1)), res=ArraySegment_sel4
        op[5]: select idx=(addr_2 + (0 + 1)), res=ArraySegment_sel5
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=4, store=2, select=2)
        op[0]: select idx=addr_1, res=ObjectWithPrimitivesClass_0_sel0
        op[1]: store idx=ArraySegment_sel0, val=ObjectWithPrimitivesClass_0_sel0
        op[2]: select idx=addr_1, res=ObjectWithPrimitivesClass_0_sel2
        op[3]: store idx=ArraySegment_sel3, val=ObjectWithPrimitivesClass_0_sel2
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=6, store=4, select=2)
        op[0]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel0
        op[1]: store idx=ArraySegment_sel0, val=ObjectWithPrimitivesClass_1_sel0
        op[2]: store idx=ArraySegment_sel1, val=(IntSegment_sel2 + 0)
        op[3]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel3
        op[4]: store idx=ArraySegment_sel3, val=ObjectWithPrimitivesClass_1_sel3
        op[5]: store idx=ArraySegment_sel4, val=(IntSegment_sel5 + (0 + 1))
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=6, store=4, select=2)
        op[0]: select idx=addr_1, res=ObjectWithPrimitivesClass_2_sel0
        op[1]: store idx=ArraySegment_sel0, val=ObjectWithPrimitivesClass_2_sel0
        op[2]: store idx=ArraySegment_sel2, val=(IntSegment_sel3 + 0)
        op[3]: select idx=addr_1, res=ObjectWithPrimitivesClass_2_sel3
        op[4]: store idx=ArraySegment_sel3, val=ObjectWithPrimitivesClass_2_sel3
        op[5]: store idx=ArraySegment_sel5, val=(IntSegment_sel6 + (0 + 1))
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=4, store=2, select=2)
        op[0]: select idx=addr_1, res=ObjectWithPrimitivesClass_3_sel0
        op[1]: store idx=ArraySegment_sel0, val=ObjectWithPrimitivesClass_3_sel0
        op[2]: select idx=addr_1, res=ObjectWithPrimitivesClass_3_sel2
        op[3]: store idx=ArraySegment_sel3, val=ObjectWithPrimitivesClass_3_sel2
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=4, store=2, select=2)
        op[0]: select idx=addr_1, res=ObjectWithPrimitivesClass_4_sel0
        op[1]: store idx=ArraySegment_sel0, val=ObjectWithPrimitivesClass_4_sel0
        op[2]: select idx=addr_1, res=ObjectWithPrimitivesClass_4_sel2
        op[3]: store idx=ArraySegment_sel3, val=ObjectWithPrimitivesClass_4_sel2
  }
  HeapConstraints: 17
    C[0]: (addr_0 != addr_1)
    C[1]: (addr_0 != addr_2)
    C[2]: (addr_3 != addr_1)
    C[3]: (addr_3 != addr_2)
    C[4]: (addr_4 != addr_1)
    C[5]: (addr_4 != addr_2)
    C[6]: (addr_1 != addr_2)
    C[7]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_3, IntSegment_sel2).select(addr_4, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_3, IntSegment_sel5).select(addr_4, IntSegment_sel6).select(addr_0, IntSegment_sel7)
    C[8]: BoolSegment
    C[9]: StringSegment
    C[10]: FloatSegment
    C[11]: ArraySegment.select((addr_2 + 0), ArraySegment_sel0).select((addr_2 + 0), ArraySegment_sel1).select((addr_2 + 0), ArraySegment_sel2).select((addr_2 + (0 + 1)), ArraySegment_sel3).select((addr_2 + (0 + 1)), ArraySegment_sel4).select((addr_2 + (0 + 1)), ArraySegment_sel5)
    C[12]: ObjectWithPrimitivesClass_2.select(addr_1, ObjectWithPrimitivesClass_2_sel0).store(ArraySegment_sel0, ObjectWithPrimitivesClass_2_sel0).store(ArraySegment_sel2, (IntSegment_sel3 + 0)).select(addr_1, ObjectWithPrimitivesClass_2_sel3).store(ArraySegment_sel3, ObjectWithPrimitivesClass_2_sel3).store(ArraySegment_sel5, (IntSegment_sel6 + (0 + 1)))
    C[13]: ObjectWithPrimitivesClass_3.select(addr_1, ObjectWithPrimitivesClass_3_sel0).store(ArraySegment_sel0, ObjectWithPrimitivesClass_3_sel0).select(addr_1, ObjectWithPrimitivesClass_3_sel2).store(ArraySegment_sel3, ObjectWithPrimitivesClass_3_sel2)
    C[14]: ObjectWithPrimitivesClass_4.select(addr_1, ObjectWithPrimitivesClass_4_sel0).store(ArraySegment_sel0, ObjectWithPrimitivesClass_4_sel0).select(addr_1, ObjectWithPrimitivesClass_4_sel2).store(ArraySegment_sel3, ObjectWithPrimitivesClass_4_sel2)
    C[15]: ObjectWithPrimitivesClass_0.select(addr_1, ObjectWithPrimitivesClass_0_sel0).store(ArraySegment_sel0, ObjectWithPrimitivesClass_0_sel0).select(addr_1, ObjectWithPrimitivesClass_0_sel2).store(ArraySegment_sel3, ObjectWithPrimitivesClass_0_sel2)
    C[16]: ObjectWithPrimitivesClass_1.select(addr_1, ObjectWithPrimitivesClass_1_sel0).store(ArraySegment_sel0, ObjectWithPrimitivesClass_1_sel0).store(ArraySegment_sel1, (IntSegment_sel2 + 0)).select(addr_1, ObjectWithPrimitivesClass_1_sel3).store(ArraySegment_sel3, ObjectWithPrimitivesClass_1_sel3).store(ArraySegment_sel4, (IntSegment_sel5 + (0 + 1)))
}

UNSAT
---- State #4 ----
Interpreter{
  PC: (((((true & (!(IntSegment_sel0 < 3))) & (0 < IntSegment_sel1)) & ((0 + 1) < IntSegment_sel4)) & (((0 + 1) + 1) < IntSegment_sel7)) & (!((((0 + 1) + 1) + 1) < IntSegment_sel10)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 5
    Params: 3, Allocations: 2
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_3) (refType=int)
      P[2]: &(addr_4) (refType=int)
    AllocationsList:
      A[0]: &(addr_1) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
      A[1]: &(addr_2) (refType=array, dim=0, base=struct(ObjectWithPrimitivesClass))
    IntSegment:    IntSegment (elem=int, ops=11, store=0, select=11)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_3, res=IntSegment_sel2
        op[3]: select idx=addr_4, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_3, res=IntSegment_sel5
        op[6]: select idx=addr_4, res=IntSegment_sel6
        op[7]: select idx=addr_0, res=IntSegment_sel7
        op[8]: select idx=addr_3, res=IntSegment_sel8
        op[9]: select idx=addr_4, res=IntSegment_sel9
        op[10]: select idx=addr_0, res=IntSegment_sel10
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=9, store=0, select=9)
        op[0]: select idx=(addr_2 + 0), res=ArraySegment_sel0
        op[1]: select idx=(addr_2 + 0), res=ArraySegment_sel1
        op[2]: select idx=(addr_2 + 0), res=ArraySegment_sel2
        op[3]: select idx=(addr_2 + (0 + 1)), res=ArraySegment_sel3
        op[4]: select idx=(addr_2 + (0 + 1)), res=ArraySegment_sel4
        op[5]: select idx=(addr_2 + (0 + 1)), res=ArraySegment_sel5
        op[6]: select idx=(addr_2 + ((0 + 1) + 1)), res=ArraySegment_sel6
        op[7]: select idx=(addr_2 + ((0 + 1) + 1)), res=ArraySegment_sel7
        op[8]: select idx=(addr_2 + ((0 + 1) + 1)), res=ArraySegment_sel8
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=6, store=3, select=3)
        op[0]: select idx=addr_1, res=ObjectWithPrimitivesClass_0_sel0
        op[1]: store idx=ArraySegment_sel0, val=ObjectWithPrimitivesClass_0_sel0
        op[2]: select idx=addr_1, res=ObjectWithPrimitivesClass_0_sel2
        op[3]: store idx=ArraySegment_sel3, val=ObjectWithPrimitivesClass_0_sel2
        op[4]: select idx=addr_1, res=ObjectWithPrimitivesClass_0_sel4
        op[5]: store idx=ArraySegment_sel6, val=ObjectWithPrimitivesClass_0_sel4
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=9, store=6, select=3)
        op[0]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel0
        op[1]: store idx=ArraySegment_sel0, val=ObjectWithPrimitivesClass_1_sel0
        op[2]: store idx=ArraySegment_sel1, val=(IntSegment_sel2 + 0)
        op[3]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel3
        op[4]: store idx=ArraySegment_sel3, val=ObjectWithPrimitivesClass_1_sel3
        op[5]: store idx=ArraySegment_sel4, val=(IntSegment_sel5 + (0 + 1))
        op[6]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel6
        op[7]: store idx=ArraySegment_sel6, val=ObjectWithPrimitivesClass_1_sel6
        op[8]: store idx=ArraySegment_sel7, val=(IntSegment_sel8 + ((0 + 1) + 1))
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=9, store=6, select=3)
        op[0]: select idx=addr_1, res=ObjectWithPrimitivesClass_2_sel0
        op[1]: store idx=ArraySegment_sel0, val=ObjectWithPrimitivesClass_2_sel0
        op[2]: store idx=ArraySegment_sel2, val=(IntSegment_sel3 + 0)
        op[3]: select idx=addr_1, res=ObjectWithPrimitivesClass_2_sel3
        op[4]: store idx=ArraySegment_sel3, val=ObjectWithPrimitivesClass_2_sel3
        op[5]: store idx=ArraySegment_sel5, val=(IntSegment_sel6 + (0 + 1))
        op[6]: select idx=addr_1, res=ObjectWithPrimitivesClass_2_sel6
        op[7]: store idx=ArraySegment_sel6, val=ObjectWithPrimitivesClass_2_sel6
        op[8]: store idx=ArraySegment_sel8, val=(IntSegment_sel9 + ((0 + 1) + 1))
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=6, store=3, select=3)
        op[0]: select idx=addr_1, res=ObjectWithPrimitivesClass_3_sel0
        op[1]: store idx=ArraySegment_sel0, val=ObjectWithPrimitivesClass_3_sel0
        op[2]: select idx=addr_1, res=ObjectWithPrimitivesClass_3_sel2
        op[3]: store idx=ArraySegment_sel3, val=ObjectWithPrimitivesClass_3_sel2
        op[4]: select idx=addr_1, res=ObjectWithPrimitivesClass_3_sel4
        op[5]: store idx=ArraySegment_sel6, val=ObjectWithPrimitivesClass_3_sel4
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=6, store=3, select=3)
        op[0]: select idx=addr_1, res=ObjectWithPrimitivesClass_4_sel0
        op[1]: store idx=ArraySegment_sel0, val=ObjectWithPrimitivesClass_4_sel0
        op[2]: select idx=addr_1, res=ObjectWithPrimitivesClass_4_sel2
        op[3]: store idx=ArraySegment_sel3, val=ObjectWithPrimitivesClass_4_sel2
        op[4]: select idx=addr_1, res=ObjectWithPrimitivesClass_4_sel4
        op[5]: store idx=ArraySegment_sel6, val=ObjectWithPrimitivesClass_4_sel4
  }
  HeapConstraints: 17
    C[0]: (addr_0 != addr_1)
    C[1]: (addr_0 != addr_2)
    C[2]: (addr_3 != addr_1)
    C[3]: (addr_3 != addr_2)
    C[4]: (addr_4 != addr_1)
    C[5]: (addr_4 != addr_2)
    C[6]: (addr_1 != addr_2)
    C[7]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_3, IntSegment_sel2).select(addr_4, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_3, IntSegment_sel5).select(addr_4, IntSegment_sel6).select(addr_0, IntSegment_sel7).select(addr_3, IntSegment_sel8).select(addr_4, IntSegment_sel9).select(addr_0, IntSegment_sel10)
    C[8]: BoolSegment
    C[9]: StringSegment
    C[10]: FloatSegment
    C[11]: ArraySegment.select((addr_2 + 0), ArraySegment_sel0).select((addr_2 + 0), ArraySegment_sel1).select((addr_2 + 0), ArraySegment_sel2).select((addr_2 + (0 + 1)), ArraySegment_sel3).select((addr_2 + (0 + 1)), ArraySegment_sel4).select((addr_2 + (0 + 1)), ArraySegment_sel5).select((addr_2 + ((0 + 1) + 1)), ArraySegment_sel6).select((addr_2 + ((0 + 1) + 1)), ArraySegment_sel7).select((addr_2 + ((0 + 1) + 1)), ArraySegment_sel8)
    C[12]: ObjectWithPrimitivesClass_3.select(addr_1, ObjectWithPrimitivesClass_3_sel0).store(ArraySegment_sel0, ObjectWithPrimitivesClass_3_sel0).select(addr_1, ObjectWithPrimitivesClass_3_sel2).store(ArraySegment_sel3, ObjectWithPrimitivesClass_3_sel2).select(addr_1, ObjectWithPrimitivesClass_3_sel4).store(ArraySegment_sel6, ObjectWithPrimitivesClass_3_sel4)
    C[13]: ObjectWithPrimitivesClass_4.select(addr_1, ObjectWithPrimitivesClass_4_sel0).store(ArraySegment_sel0, ObjectWithPrimitivesClass_4_sel0).select(addr_1, ObjectWithPrimitivesClass_4_sel2).store(ArraySegment_sel3, ObjectWithPrimitivesClass_4_sel2).select(addr_1, ObjectWithPrimitivesClass_4_sel4).store(ArraySegment_sel6, ObjectWithPrimitivesClass_4_sel4)
    C[14]: ObjectWithPrimitivesClass_0.select(addr_1, ObjectWithPrimitivesClass_0_sel0).store(ArraySegment_sel0, ObjectWithPrimitivesClass_0_sel0).select(addr_1, ObjectWithPrimitivesClass_0_sel2).store(ArraySegment_sel3, ObjectWithPrimitivesClass_0_sel2).select(addr_1, ObjectWithPrimitivesClass_0_sel4).store(ArraySegment_sel6, ObjectWithPrimitivesClass_0_sel4)
    C[15]: ObjectWithPrimitivesClass_1.select(addr_1, ObjectWithPrimitivesClass_1_sel0).store(ArraySegment_sel0, ObjectWithPrimitivesClass_1_sel0).store(ArraySegment_sel1, (IntSegment_sel2 + 0)).select(addr_1, ObjectWithPrimitivesClass_1_sel3).store(ArraySegment_sel3, ObjectWithPrimitivesClass_1_sel3).store(ArraySegment_sel4, (IntSegment_sel5 + (0 + 1))).select(addr_1, ObjectWithPrimitivesClass_1_sel6).store(ArraySegment_sel6, ObjectWithPrimitivesClass_1_sel6).store(ArraySegment_sel7, (IntSegment_sel8 + ((0 + 1) + 1)))
    C[16]: ObjectWithPrimitivesClass_2.select(addr_1, ObjectWithPrimitivesClass_2_sel0).store(ArraySegment_sel0, ObjectWithPrimitivesClass_2_sel0).store(ArraySegment_sel2, (IntSegment_sel3 + 0)).select(addr_1, ObjectWithPrimitivesClass_2_sel3).store(ArraySegment_sel3, ObjectWithPrimitivesClass_2_sel3).store(ArraySegment_sel5, (IntSegment_sel6 + (0 + 1))).select(addr_1, ObjectWithPrimitivesClass_2_sel6).store(ArraySegment_sel6, ObjectWithPrimitivesClass_2_sel6).store(ArraySegment_sel8, (IntSegment_sel9 + ((0 + 1) + 1)))
}

SAT
Model:
addr_2 -> #x0000000000000001
addr_4 -> #x0000000000000000
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000003)
addr_3 -> #x0000000000000000
addr_0 -> #x0000000000000000
addr_1 -> #x0000000000000002
ObjectWithPrimitivesClass_0 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
ArraySegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
ObjectWithPrimitivesClass_0_sel4 -> #x0000000000000000
ObjectWithPrimitivesClass_2 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
ObjectWithPrimitivesClass_2_sel3 -> #x0000000000000000
ArraySegment_sel2 -> #x0000000000000000
IntSegment_sel5 -> #x0000000000000003
ObjectWithPrimitivesClass_0_sel0 -> #x0000000000000000
ObjectWithPrimitivesClass_4 -> ((as const (Array (_ BitVec 64) Real)) 0.0)
ObjectWithPrimitivesClass_4_sel4 -> 0.0
ObjectWithPrimitivesClass_1 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
ObjectWithPrimitivesClass_1_sel6 -> #x0000000000000000
ArraySegment_sel4 -> #x0000000000000000
ArraySegment_sel0 -> #x0000000000000000
IntSegment_sel10 -> #x0000000000000003
IntSegment_sel1 -> #x0000000000000003
IntSegment_sel7 -> #x0000000000000003
ObjectWithPrimitivesClass_1_sel3 -> #x0000000000000000
ObjectWithPrimitivesClass_0_sel2 -> #x0000000000000000
ArraySegment_sel5 -> #x0000000000000000
ObjectWithPrimitivesClass_2_sel6 -> #x0000000000000000
ArraySegment_sel1 -> #x0000000000000000
IntSegment_sel9 -> #x0000000000000003
IntSegment_sel0 -> #x0000000000000003
IntSegment_sel6 -> #x0000000000000003
ObjectWithPrimitivesClass_2_sel0 -> #x0000000000000000
ArraySegment_sel8 -> #x0000000000000000
IntSegment_sel4 -> #x0000000000000003
IntSegment_sel2 -> #x0000000000000003
ObjectWithPrimitivesClass_4_sel0 -> 0.0
ObjectWithPrimitivesClass_3 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
ObjectWithPrimitivesClass_3_sel2 -> #x0000000000000000
ArraySegment_sel6 -> #x0000000000000000
ArraySegment_sel3 -> #x0000000000000000
ObjectWithPrimitivesClass_4_sel2 -> 0.0
ObjectWithPrimitivesClass_3_sel4 -> #x0000000000000000
ObjectWithPrimitivesClass_1_sel0 -> #x0000000000000000
ObjectWithPrimitivesClass_3_sel0 -> #x0000000000000000
ArraySegment_sel7 -> #x0000000000000000
IntSegment_sel8 -> #x0000000000000003
IntSegment_sel3 -> #x0000000000000003

---- State #5 ----
Interpreter{
  PC: ((((((true & (!(IntSegment_sel0 < 3))) & (0 < IntSegment_sel1)) & ((0 + 1) < IntSegment_sel4)) & (((0 + 1) + 1) < IntSegment_sel7)) & ((((0 + 1) + 1) + 1) < IntSegment_sel10)) & (!(((((0 + 1) + 1) + 1) + 1) < IntSegment_sel13)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 5
    Params: 3, Allocations: 2
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_3) (refType=int)
      P[2]: &(addr_4) (refType=int)
    AllocationsList:
      A[0]: &(addr_1) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
      A[1]: &(addr_2) (refType=array, dim=0, base=struct(ObjectWithPrimitivesClass))
    IntSegment:    IntSegment (elem=int, ops=14, store=0, select=14)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_3, res=IntSegment_sel2
        op[3]: select idx=addr_4, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_3, res=IntSegment_sel5
        op[6]: select idx=addr_4, res=IntSegment_sel6
        op[7]: select idx=addr_0, res=IntSegment_sel7
        op[8]: select idx=addr_3, res=IntSegment_sel8
        op[9]: select idx=addr_4, res=IntSegment_sel9
        op[10]: select idx=addr_0, res=IntSegment_sel10
        op[11]: select idx=addr_3, res=IntSegment_sel11
        op[12]: select idx=addr_4, res=IntSegment_sel12
        op[13]: select idx=addr_0, res=IntSegment_sel13
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=12, store=0, select=12)
        op[0]: select idx=(addr_2 + 0), res=ArraySegment_sel0
        op[1]: select idx=(addr_2 + 0), res=ArraySegment_sel1
        op[2]: select idx=(addr_2 + 0), res=ArraySegment_sel2
        op[3]: select idx=(addr_2 + (0 + 1)), res=ArraySegment_sel3
        op[4]: select idx=(addr_2 + (0 + 1)), res=ArraySegment_sel4
        op[5]: select idx=(addr_2 + (0 + 1)), res=ArraySegment_sel5
        op[6]: select idx=(addr_2 + ((0 + 1) + 1)), res=ArraySegment_sel6
        op[7]: select idx=(addr_2 + ((0 + 1) + 1)), res=ArraySegment_sel7
        op[8]: select idx=(addr_2 + ((0 + 1) + 1)), res=ArraySegment_sel8
        op[9]: select idx=(addr_2 + (((0 + 1) + 1) + 1)), res=ArraySegment_sel9
        op[10]: select idx=(addr_2 + (((0 + 1) + 1) + 1)), res=ArraySegment_sel10
        op[11]: select idx=(addr_2 + (((0 + 1) + 1) + 1)), res=ArraySegment_sel11
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=8, store=4, select=4)
        op[0]: select idx=addr_1, res=ObjectWithPrimitivesClass_0_sel0
        op[1]: store idx=ArraySegment_sel0, val=ObjectWithPrimitivesClass_0_sel0
        op[2]: select idx=addr_1, res=ObjectWithPrimitivesClass_0_sel2
        op[3]: store idx=ArraySegment_sel3, val=ObjectWithPrimitivesClass_0_sel2
        op[4]: select idx=addr_1, res=ObjectWithPrimitivesClass_0_sel4
        op[5]: store idx=ArraySegment_sel6, val=ObjectWithPrimitivesClass_0_sel4
        op[6]: select idx=addr_1, res=ObjectWithPrimitivesClass_0_sel6
        op[7]: store idx=ArraySegment_sel9, val=ObjectWithPrimitivesClass_0_sel6
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=12, store=8, select=4)
        op[0]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel0
        op[1]: store idx=ArraySegment_sel0, val=ObjectWithPrimitivesClass_1_sel0
        op[2]: store idx=ArraySegment_sel1, val=(IntSegment_sel2 + 0)
        op[3]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel3
        op[4]: store idx=ArraySegment_sel3, val=ObjectWithPrimitivesClass_1_sel3
        op[5]: store idx=ArraySegment_sel4, val=(IntSegment_sel5 + (0 + 1))
        op[6]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel6
        op[7]: store idx=ArraySegment_sel6, val=ObjectWithPrimitivesClass_1_sel6
        op[8]: store idx=ArraySegment_sel7, val=(IntSegment_sel8 + ((0 + 1) + 1))
        op[9]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel9
        op[10]: store idx=ArraySegment_sel9, val=ObjectWithPrimitivesClass_1_sel9
        op[11]: store idx=ArraySegment_sel10, val=(IntSegment_sel11 + (((0 + 1) + 1) + 1))
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=12, store=8, select=4)
        op[0]: select idx=addr_1, res=ObjectWithPrimitivesClass_2_sel0
        op[1]: store idx=ArraySegment_sel0, val=ObjectWithPrimitivesClass_2_sel0
        op[2]: store idx=ArraySegment_sel2, val=(IntSegment_sel3 + 0)
        op[3]: select idx=addr_1, res=ObjectWithPrimitivesClass_2_sel3
        op[4]: store idx=ArraySegment_sel3, val=ObjectWithPrimitivesClass_2_sel3
        op[5]: store idx=ArraySegment_sel5, val=(IntSegment_sel6 + (0 + 1))
        op[6]: select idx=addr_1, res=ObjectWithPrimitivesClass_2_sel6
        op[7]: store idx=ArraySegment_sel6, val=ObjectWithPrimitivesClass_2_sel6
        op[8]: store idx=ArraySegment_sel8, val=(IntSegment_sel9 + ((0 + 1) + 1))
        op[9]: select idx=addr_1, res=ObjectWithPrimitivesClass_2_sel9
        op[10]: store idx=ArraySegment_sel9, val=ObjectWithPrimitivesClass_2_sel9
        op[11]: store idx=ArraySegment_sel11, val=(IntSegment_sel12 + (((0 + 1) + 1) + 1))
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=8, store=4, select=4)
        op[0]: select idx=addr_1, res=ObjectWithPrimitivesClass_3_sel0
        op[1]: store idx=ArraySegment_sel0, val=ObjectWithPrimitivesClass_3_sel0
        op[2]: select idx=addr_1, res=ObjectWithPrimitivesClass_3_sel2
        op[3]: store idx=ArraySegment_sel3, val=ObjectWithPrimitivesClass_3_sel2
        op[4]: select idx=addr_1, res=ObjectWithPrimitivesClass_3_sel4
        op[5]: store idx=ArraySegment_sel6, val=ObjectWithPrimitivesClass_3_sel4
        op[6]: select idx=addr_1, res=ObjectWithPrimitivesClass_3_sel6
        op[7]: store idx=ArraySegment_sel9, val=ObjectWithPrimitivesClass_3_sel6
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=8, store=4, select=4)
        op[0]: select idx=addr_1, res=ObjectWithPrimitivesClass_4_sel0
        op[1]: store idx=ArraySegment_sel0, val=ObjectWithPrimitivesClass_4_sel0
        op[2]: select idx=addr_1, res=ObjectWithPrimitivesClass_4_sel2
        op[3]: store idx=ArraySegment_sel3, val=ObjectWithPrimitivesClass_4_sel2
        op[4]: select idx=addr_1, res=ObjectWithPrimitivesClass_4_sel4
        op[5]: store idx=ArraySegment_sel6, val=ObjectWithPrimitivesClass_4_sel4
        op[6]: select idx=addr_1, res=ObjectWithPrimitivesClass_4_sel6
        op[7]: store idx=ArraySegment_sel9, val=ObjectWithPrimitivesClass_4_sel6
  }
  HeapConstraints: 17
    C[0]: (addr_0 != addr_1)
    C[1]: (addr_0 != addr_2)
    C[2]: (addr_3 != addr_1)
    C[3]: (addr_3 != addr_2)
    C[4]: (addr_4 != addr_1)
    C[5]: (addr_4 != addr_2)
    C[6]: (addr_1 != addr_2)
    C[7]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_3, IntSegment_sel2).select(addr_4, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_3, IntSegment_sel5).select(addr_4, IntSegment_sel6).select(addr_0, IntSegment_sel7).select(addr_3, IntSegment_sel8).select(addr_4, IntSegment_sel9).select(addr_0, IntSegment_sel10).select(addr_3, IntSegment_sel11).select(addr_4, IntSegment_sel12).select(addr_0, IntSegment_sel13)
    C[8]: BoolSegment
    C[9]: StringSegment
    C[10]: FloatSegment
    C[11]: ArraySegment.select((addr_2 + 0), ArraySegment_sel0).select((addr_2 + 0), ArraySegment_sel1).select((addr_2 + 0), ArraySegment_sel2).select((addr_2 + (0 + 1)), ArraySegment_sel3).select((addr_2 + (0 + 1)), ArraySegment_sel4).select((addr_2 + (0 + 1)), ArraySegment_sel5).select((addr_2 + ((0 + 1) + 1)), ArraySegment_sel6).select((addr_2 + ((0 + 1) + 1)), ArraySegment_sel7).select((addr_2 + ((0 + 1) + 1)), ArraySegment_sel8).select((addr_2 + (((0 + 1) + 1) + 1)), ArraySegment_sel9).select((addr_2 + (((0 + 1) + 1) + 1)), ArraySegment_sel10).select((addr_2 + (((0 + 1) + 1) + 1)), ArraySegment_sel11)
    C[12]: ObjectWithPrimitivesClass_0.select(addr_1, ObjectWithPrimitivesClass_0_sel0).store(ArraySegment_sel0, ObjectWithPrimitivesClass_0_sel0).select(addr_1, ObjectWithPrimitivesClass_0_sel2).store(ArraySegment_sel3, ObjectWithPrimitivesClass_0_sel2).select(addr_1, ObjectWithPrimitivesClass_0_sel4).store(ArraySegment_sel6, ObjectWithPrimitivesClass_0_sel4).select(addr_1, ObjectWithPrimitivesClass_0_sel6).store(ArraySegment_sel9, ObjectWithPrimitivesClass_0_sel6)
    C[13]: ObjectWithPrimitivesClass_1.select(addr_1, ObjectWithPrimitivesClass_1_sel0).store(ArraySegment_sel0, ObjectWithPrimitivesClass_1_sel0).store(ArraySegment_sel1, (IntSegment_sel2 + 0)).select(addr_1, ObjectWithPrimitivesClass_1_sel3).store(ArraySegment_sel3, ObjectWithPrimitivesClass_1_sel3).store(ArraySegment_sel4, (IntSegment_sel5 + (0 + 1))).select(addr_1, ObjectWithPrimitivesClass_1_sel6).store(ArraySegment_sel6, ObjectWithPrimitivesClass_1_sel6).store(ArraySegment_sel7, (IntSegment_sel8 + ((0 + 1) + 1))).select(addr_1, ObjectWithPrimitivesClass_1_sel9).store(ArraySegment_sel9, ObjectWithPrimitivesClass_1_sel9).store(ArraySegment_sel10, (IntSegment_sel11 + (((0 + 1) + 1) + 1)))
    C[14]: ObjectWithPrimitivesClass_2.select(addr_1, ObjectWithPrimitivesClass_2_sel0).store(ArraySegment_sel0, ObjectWithPrimitivesClass_2_sel0).store(ArraySegment_sel2, (IntSegment_sel3 + 0)).select(addr_1, ObjectWithPrimitivesClass_2_sel3).store(ArraySegment_sel3, ObjectWithPrimitivesClass_2_sel3).store(ArraySegment_sel5, (IntSegment_sel6 + (0 + 1))).select(addr_1, ObjectWithPrimitivesClass_2_sel6).store(ArraySegment_sel6, ObjectWithPrimitivesClass_2_sel6).store(ArraySegment_sel8, (IntSegment_sel9 + ((0 + 1) + 1))).select(addr_1, ObjectWithPrimitivesClass_2_sel9).store(ArraySegment_sel9, ObjectWithPrimitivesClass_2_sel9).store(ArraySegment_sel11, (IntSegment_sel12 + (((0 + 1) + 1) + 1)))
    C[15]: ObjectWithPrimitivesClass_3.select(addr_1, ObjectWithPrimitivesClass_3_sel0).store(ArraySegment_sel0, ObjectWithPrimitivesClass_3_sel0).select(addr_1, ObjectWithPrimitivesClass_3_sel2).store(ArraySegment_sel3, ObjectWithPrimitivesClass_3_sel2).select(addr_1, ObjectWithPrimitivesClass_3_sel4).store(ArraySegment_sel6, ObjectWithPrimitivesClass_3_sel4).select(addr_1, ObjectWithPrimitivesClass_3_sel6).store(ArraySegment_sel9, ObjectWithPrimitivesClass_3_sel6)
    C[16]: ObjectWithPrimitivesClass_4.select(addr_1, ObjectWithPrimitivesClass_4_sel0).store(ArraySegment_sel0, ObjectWithPrimitivesClass_4_sel0).select(addr_1, ObjectWithPrimitivesClass_4_sel2).store(ArraySegment_sel3, ObjectWithPrimitivesClass_4_sel2).select(addr_1, ObjectWithPrimitivesClass_4_sel4).store(ArraySegment_sel6, ObjectWithPrimitivesClass_4_sel4).select(addr_1, ObjectWithPrimitivesClass_4_sel6).store(ArraySegment_sel9, ObjectWithPrimitivesClass_4_sel6)
}

SAT
Model:
addr_2 -> #x0000000000000001
addr_4 -> #x0000000000000000
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000004)
addr_3 -> #x0000000000000000
addr_0 -> #x0000000000000000
addr_1 -> #x0000000000000002
ObjectWithPrimitivesClass_0 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
ArraySegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
ObjectWithPrimitivesClass_0_sel4 -> #x0000000000000000
ObjectWithPrimitivesClass_2 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
ObjectWithPrimitivesClass_2_sel3 -> #x0000000000000000
ArraySegment_sel9 -> #x0000000000000000
ArraySegment_sel2 -> #x0000000000000000
ObjectWithPrimitivesClass_4 -> ((as const (Array (_ BitVec 64) Real)) 0.0)
ObjectWithPrimitivesClass_4_sel6 -> 0.0
ArraySegment_sel10 -> #x0000000000000000
IntSegment_sel5 -> #x0000000000000004
ObjectWithPrimitivesClass_1 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
ObjectWithPrimitivesClass_1_sel6 -> #x0000000000000000
ObjectWithPrimitivesClass_0_sel0 -> #x0000000000000000
ObjectWithPrimitivesClass_4_sel4 -> 0.0
ArraySegment_sel4 -> #x0000000000000000
ArraySegment_sel0 -> #x0000000000000000
IntSegment_sel10 -> #x0000000000000004
IntSegment_sel1 -> #x0000000000000004
ObjectWithPrimitivesClass_1_sel9 -> #x0000000000000000
ObjectWithPrimitivesClass_1_sel3 -> #x0000000000000000
IntSegment_sel7 -> #x0000000000000004
ObjectWithPrimitivesClass_0_sel2 -> #x0000000000000000
ObjectWithPrimitivesClass_2_sel6 -> #x0000000000000000
ArraySegment_sel5 -> #x0000000000000000
IntSegment_sel12 -> #x0000000000000004
ArraySegment_sel1 -> #x0000000000000000
IntSegment_sel9 -> #x0000000000000004
IntSegment_sel0 -> #x0000000000000004
IntSegment_sel6 -> #x0000000000000004
ObjectWithPrimitivesClass_0_sel6 -> #x0000000000000000
ObjectWithPrimitivesClass_2_sel0 -> #x0000000000000000
ArraySegment_sel8 -> #x0000000000000000
IntSegment_sel4 -> #x0000000000000004
IntSegment_sel2 -> #x0000000000000004
ObjectWithPrimitivesClass_4_sel0 -> 0.0
ObjectWithPrimitivesClass_3 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
ObjectWithPrimitivesClass_3_sel2 -> #x0000000000000000
IntSegment_sel11 -> #x0000000000000004
ArraySegment_sel6 -> #x0000000000000000
ArraySegment_sel3 -> #x0000000000000000
IntSegment_sel13 -> #x0000000000000004
ObjectWithPrimitivesClass_1_sel0 -> #x0000000000000000
ObjectWithPrimitivesClass_4_sel2 -> 0.0
ObjectWithPrimitivesClass_3_sel6 -> #x0000000000000000
ObjectWithPrimitivesClass_3_sel4 -> #x0000000000000000
ObjectWithPrimitivesClass_3_sel0 -> #x0000000000000000
ArraySegment_sel7 -> #x0000000000000000
IntSegment_sel8 -> #x0000000000000004
ArraySegment_sel11 -> #x0000000000000000
ObjectWithPrimitivesClass_2_sel9 -> #x0000000000000000
IntSegment_sel3 -> #x0000000000000004


==================== IsIdentityMatrix ====================
CFG построен для функции с 17 блоками
interpretDynamically: len_matrix < 3:int
interpretDynamically: if t0 goto 1 else 2
resolveExpression: len_matrix < 3:int
resolveExpression: parameter len_matrix : int
resolveExpression: 3:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 9: t7] #i
resolveExpression: 0:int
interpretDynamically: t4 < len_matrix
interpretDynamically: if t5 goto 3 else 4
resolveExpression: t4 < len_matrix
resolveExpression: phi [2: 0:int, 9: t7] #i
resolveExpression: parameter len_matrix : int
interpretDynamically: return true:bool
resolveExpression: true:bool
interpretDynamically: &len_matrix1[t4]
interpretDynamically: *t1
interpretDynamically: t2 != len_matrix
interpretDynamically: return false:bool
resolveExpression: false:bool
interpretDynamically: if t3 goto 6 else 7
resolveExpression: t2 != len_matrix
resolveExpression: *t1
resolveExpression: parameter len_matrix1 : []int
resolveExpression: phi [2: 0:int, 9: t7] #i
resolveExpression: parameter len_matrix : int
interpretDynamically: return false:bool
resolveExpression: false:bool
interpretDynamically: jump 10
interpretDynamically: phi [7: 0:int, 15: t18] #j
resolveExpression: 0:int
interpretDynamically: &len_matrix1[t4]
interpretDynamically: *t9
interpretDynamically: t8 < t10
interpretDynamically: if t11 goto 8 else 9
resolveExpression: t8 < t10
resolveExpression: phi [7: 0:int, 15: t18] #j
resolveExpression: *t9
resolveExpression: parameter len_matrix1 : []int
resolveExpression: phi [2: 0:int, 9: t7] #i
interpretDynamically: t4 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 9: t7] #i
resolveExpression: t4 + 1:int
resolveExpression: phi [2: 0:int, 9: t7] #i
resolveExpression: 1:int
interpretDynamically: t4 < len_matrix
interpretDynamically: if t5 goto 3 else 4
resolveExpression: t4 < len_matrix
resolveExpression: phi [2: 0:int, 9: t7] #i
resolveExpression: parameter len_matrix : int
interpretDynamically: &len_matrix1[t4]
interpretDynamically: *t1
interpretDynamically: return true:bool
resolveExpression: true:bool
interpretDynamically: t2 != len_matrix
interpretDynamically: if t3 goto 6 else 7
resolveExpression: t2 != len_matrix
resolveExpression: *t1
resolveExpression: parameter len_matrix1 : []int
resolveExpression: phi [2: 0:int, 9: t7] #i
resolveExpression: parameter len_matrix : int
interpretDynamically: jump 10
interpretDynamically: return false:bool
resolveExpression: false:bool
interpretDynamically: phi [7: 0:int, 15: t18] #j
resolveExpression: 0:int
interpretDynamically: &len_matrix1[t4]
interpretDynamically: *t9
interpretDynamically: t8 < t10
interpretDynamically: if t11 goto 8 else 9
resolveExpression: t8 < t10
resolveExpression: phi [7: 0:int, 15: t18] #j
resolveExpression: *t9
resolveExpression: parameter len_matrix1 : []int
resolveExpression: phi [2: 0:int, 9: t7] #i
interpretDynamically: t4 == t8
interpretDynamically: if t6 goto 13 else 12
resolveExpression: t4 == t8
resolveExpression: phi [2: 0:int, 9: t7] #i
resolveExpression: phi [7: 0:int, 15: t18] #j
interpretDynamically: t4 + 1:int
interpretDynamically: jump 5
interpretDynamically: &matrix[t4]
interpretDynamically: phi [2: 0:int, 9: t7] #i
resolveExpression: t4 + 1:int
resolveExpression: phi [2: 0:int, 9: t7] #i
resolveExpression: 1:int
interpretDynamically: t4 != t8
interpretDynamically: if t12 goto 16 else 15
resolveExpression: t4 != t8
resolveExpression: phi [2: 0:int, 9: t7] #i
resolveExpression: phi [7: 0:int, 15: t18] #j
interpretDynamically: &matrix[t4]
interpretDynamically: t8 + 1:int
interpretDynamically: jump 10
interpretDynamically: t4 < len_matrix
interpretDynamically: phi [7: 0:int, 15: t18] #j
resolveExpression: t8 + 1:int
resolveExpression: phi [7: 0:int, 15: t18] #j
resolveExpression: 1:int
interpretDynamically: &len_matrix1[t4]
interpretDynamically: *t9
interpretDynamically: t8 < t10
interpretDynamically: if t5 goto 3 else 4
resolveExpression: t4 < len_matrix
resolveExpression: phi [2: 0:int, 9: t7] #i
resolveExpression: parameter len_matrix : int
interpretDynamically: &len_matrix1[t4]
interpretDynamically: *t1
interpretDynamically: t2 != len_matrix
interpretDynamically: if t3 goto 6 else 7
resolveExpression: t2 != len_matrix
resolveExpression: *t1
resolveExpression: parameter len_matrix1 : []int
resolveExpression: phi [2: 0:int, 9: t7] #i
resolveExpression: parameter len_matrix : int
interpretDynamically: jump 10
interpretDynamically: return false:bool
resolveExpression: false:bool
interpretDynamically: if t11 goto 8 else 9
resolveExpression: t8 < t10
resolveExpression: phi [7: 0:int, 15: t18] #j
resolveExpression: *t9
resolveExpression: parameter len_matrix1 : []int
resolveExpression: phi [2: 0:int, 9: t7] #i
interpretDynamically: t4 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 9: t7] #i
resolveExpression: t4 + 1:int
resolveExpression: phi [2: 0:int, 9: t7] #i
resolveExpression: 1:int
interpretDynamically: t4 < len_matrix
interpretDynamically: phi [7: 0:int, 15: t18] #j
resolveExpression: 0:int
interpretDynamically: &len_matrix1[t4]
interpretDynamically: *t9
interpretDynamically: t8 < t10
interpretDynamically: if t11 goto 8 else 9
resolveExpression: t8 < t10
resolveExpression: phi [7: 0:int, 15: t18] #j
resolveExpression: *t9
resolveExpression: parameter len_matrix1 : []int
resolveExpression: phi [2: 0:int, 9: t7] #i
interpretDynamically: t4 == t8
interpretDynamically: if t6 goto 13 else 12
resolveExpression: t4 == t8
resolveExpression: phi [2: 0:int, 9: t7] #i
resolveExpression: phi [7: 0:int, 15: t18] #j
interpretDynamically: t4 != t8
interpretDynamically: if t12 goto 16 else 15
resolveExpression: t4 != t8
resolveExpression: phi [2: 0:int, 9: t7] #i
resolveExpression: phi [7: 0:int, 15: t18] #j
interpretDynamically: &matrix[t4]
interpretDynamically: t8 + 1:int
interpretDynamically: jump 10
interpretDynamically: phi [7: 0:int, 15: t18] #j
resolveExpression: t8 + 1:int
resolveExpression: phi [7: 0:int, 15: t18] #j
resolveExpression: 1:int
interpretDynamically: &matrix[t4]
interpretDynamically: &len_matrix1[t4]
interpretDynamically: *t9
interpretDynamically: t8 < t10
interpretDynamically: return true:bool
resolveExpression: true:bool
interpretDynamically: t4 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 9: t7] #i
resolveExpression: t4 + 1:int
resolveExpression: phi [2: 0:int, 9: t7] #i
resolveExpression: 1:int
interpretDynamically: t4 < len_matrix
interpretDynamically: if t5 goto 3 else 4
resolveExpression: t4 < len_matrix
resolveExpression: phi [2: 0:int, 9: t7] #i
resolveExpression: parameter len_matrix : int
interpretDynamically: &len_matrix1[t4]
interpretDynamically: *t1
interpretDynamically: return true:bool
resolveExpression: true:bool
interpretDynamically: t4 == t8
interpretDynamically: if t6 goto 13 else 12
resolveExpression: t4 == t8
resolveExpression: phi [2: 0:int, 9: t7] #i
resolveExpression: phi [7: 0:int, 15: t18] #j
interpretDynamically: t4 != t8
interpretDynamically: if t12 goto 16 else 15
resolveExpression: t4 != t8
resolveExpression: phi [2: 0:int, 9: t7] #i
resolveExpression: phi [7: 0:int, 15: t18] #j
interpretDynamically: &matrix[t4]
interpretDynamically: t8 + 1:int
interpretDynamically: jump 10
interpretDynamically: phi [7: 0:int, 15: t18] #j
resolveExpression: t8 + 1:int
resolveExpression: phi [7: 0:int, 15: t18] #j
resolveExpression: 1:int
interpretDynamically: *t19
interpretDynamically: &t20[t8]
interpretDynamically: *t21
interpretDynamically: t22 != 0:int
interpretDynamically: if t23 goto 14 else 15
resolveExpression: t22 != 0:int
resolveExpression: *t21
resolveExpression: *t19
resolveExpression: parameter matrix : [][]int
resolveExpression: phi [2: 0:int, 9: t7] #i
resolveExpression: phi [7: 0:int, 15: t18] #j
resolveExpression: 0:int
interpretDynamically: return false:bool
resolveExpression: false:bool
interpretDynamically: t8 + 1:int
interpretDynamically: jump 10
interpretDynamically: phi [7: 0:int, 15: t18] #j
resolveExpression: t8 + 1:int
resolveExpression: phi [7: 0:int, 15: t18] #j
resolveExpression: 1:int
interpretDynamically: &len_matrix1[t4]
interpretDynamically: *t9
interpretDynamically: &len_matrix1[t4]
interpretDynamically: *t9
interpretDynamically: t8 < t10
interpretDynamically: if t11 goto 8 else 9
resolveExpression: t8 < t10
resolveExpression: phi [7: 0:int, 15: t18] #j
resolveExpression: *t9
resolveExpression: parameter len_matrix1 : []int
resolveExpression: phi [2: 0:int, 9: t7] #i
interpretDynamically: &matrix[t4]
interpretDynamically: *t13
interpretDynamically: &t14[t8]
interpretDynamically: t4 == t8
interpretDynamically: if t6 goto 13 else 12
resolveExpression: t4 == t8
resolveExpression: phi [2: 0:int, 9: t7] #i
resolveExpression: phi [7: 0:int, 15: t18] #j
interpretDynamically: t4 != t8
interpretDynamically: &matrix[t4]
interpretDynamically: *t13
interpretDynamically: &t14[t8]
interpretDynamically: *t15
interpretDynamically: t16 != 1:int
interpretDynamically: if t17 goto 11 else 12
resolveExpression: t16 != 1:int
resolveExpression: *t15
resolveExpression: *t13
resolveExpression: parameter matrix : [][]int
resolveExpression: phi [2: 0:int, 9: t7] #i
resolveExpression: phi [7: 0:int, 15: t18] #j
resolveExpression: 1:int
interpretDynamically: return false:bool
resolveExpression: false:bool
interpretDynamically: t4 != t8
interpretDynamically: if t12 goto 16 else 15
resolveExpression: t4 != t8
resolveExpression: phi [2: 0:int, 9: t7] #i
resolveExpression: phi [7: 0:int, 15: t18] #j
interpretDynamically: t8 + 1:int
interpretDynamically: &matrix[t4]
interpretDynamically: jump 10
interpretDynamically: phi [7: 0:int, 15: t18] #j
resolveExpression: t8 + 1:int
resolveExpression: phi [7: 0:int, 15: t18] #j
resolveExpression: 1:int
interpretDynamically: &len_matrix1[t4]
interpretDynamically: *t9
interpretDynamically: if t12 goto 16 else 15
resolveExpression: t4 != t8
resolveExpression: phi [2: 0:int, 9: t7] #i
resolveExpression: phi [7: 0:int, 15: t18] #j
interpretDynamically: t8 + 1:int
interpretDynamically: jump 10
interpretDynamically: phi [7: 0:int, 15: t18] #j
resolveExpression: t8 + 1:int
resolveExpression: phi [7: 0:int, 15: t18] #j
resolveExpression: 1:int
interpretDynamically: &len_matrix1[t4]
interpretDynamically: *t9
interpretDynamically: &matrix[t4]
interpretDynamically: *t19
interpretDynamically: &t20[t8]
interpretDynamically: *t21
interpretDynamically: t22 != 0:int
interpretDynamically: if t23 goto 14 else 15
resolveExpression: t22 != 0:int
resolveExpression: *t21
resolveExpression: *t19
resolveExpression: parameter matrix : [][]int
resolveExpression: phi [2: 0:int, 9: t7] #i
resolveExpression: phi [7: 0:int, 15: t18] #j
resolveExpression: 0:int
interpretDynamically: t8 + 1:int
interpretDynamically: jump 10
interpretDynamically: phi [7: 0:int, 15: t18] #j
resolveExpression: t8 + 1:int
resolveExpression: phi [7: 0:int, 15: t18] #j
resolveExpression: 1:int
interpretDynamically: &len_matrix1[t4]
interpretDynamically: *t9
interpretDynamically: t8 < t10
interpretDynamically: if t11 goto 8 else 9
resolveExpression: t8 < t10
resolveExpression: phi [7: 0:int, 15: t18] #j
resolveExpression: *t9
resolveExpression: parameter len_matrix1 : []int
resolveExpression: phi [2: 0:int, 9: t7] #i
interpretDynamically: t4 == t8
interpretDynamically: if t6 goto 13 else 12
resolveExpression: t4 == t8
resolveExpression: phi [2: 0:int, 9: t7] #i
resolveExpression: phi [7: 0:int, 15: t18] #j
Analyse stopped: context deadline exceeded
States: 10
---- State #0 ----
Interpreter{
  PC: ((true & (!(IntSegment_sel0 < 3))) & (!(0 < IntSegment_sel1)))
  ReturnResults:
    [0] true
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

UNSAT
---- State #1 ----
Interpreter{
  PC: (true & (IntSegment_sel0 < 3))
  ReturnResults:
    [0] false
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
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000002)
addr_0 -> #x0000000000000000
IntSegment_sel0 -> #x0000000000000002

---- State #2 ----
Interpreter{
  PC: (((true & (!(IntSegment_sel0 < 3))) & (0 < IntSegment_sel1)) & (IntSegment_sel2 != IntSegment_sel3))
  ReturnResults:
    [0] false
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    IntSegment:    IntSegment (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=(addr_1 + 0), res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select((addr_1 + 0), IntSegment_sel2).select(addr_0, IntSegment_sel3)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
addr_0 -> #xffffffffffffffff
addr_1 -> #x0000000000000000
IntSegment -> (store ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x5016d276f85d9f82)
                     #xffffffffffffffff
                     #x4000000000000000)
IntSegment_sel2 -> #x5016d276f85d9f82
IntSegment_sel1 -> #x4000000000000000
IntSegment_sel3 -> #x4000000000000000
IntSegment_sel0 -> #x4000000000000000

---- State #3 ----
Interpreter{
  PC: (((((true & (!(IntSegment_sel0 < 3))) & (0 < IntSegment_sel1)) & (!(IntSegment_sel2 != IntSegment_sel3))) & (!(0 < IntSegment_sel4))) & (!((0 + 1) < IntSegment_sel5)))
  ReturnResults:
    [0] true
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    IntSegment:    IntSegment (elem=int, ops=6, store=0, select=6)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=(addr_1 + 0), res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=(addr_1 + 0), res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select((addr_1 + 0), IntSegment_sel2).select(addr_0, IntSegment_sel3).select((addr_1 + 0), IntSegment_sel4).select(addr_0, IntSegment_sel5)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #4 ----
Interpreter{
  PC: ((((((true & (!(IntSegment_sel0 < 3))) & (0 < IntSegment_sel1)) & (!(IntSegment_sel2 != IntSegment_sel3))) & (!(0 < IntSegment_sel4))) & ((0 + 1) < IntSegment_sel5)) & (IntSegment_sel6 != IntSegment_sel7))
  ReturnResults:
    [0] false
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    IntSegment:    IntSegment (elem=int, ops=8, store=0, select=8)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=(addr_1 + 0), res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=(addr_1 + 0), res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=(addr_1 + (0 + 1)), res=IntSegment_sel6
        op[7]: select idx=addr_0, res=IntSegment_sel7
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select((addr_1 + 0), IntSegment_sel2).select(addr_0, IntSegment_sel3).select((addr_1 + 0), IntSegment_sel4).select(addr_0, IntSegment_sel5).select((addr_1 + (0 + 1)), IntSegment_sel6).select(addr_0, IntSegment_sel7)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #5 ----
Interpreter{
  PC: (((((((((true & (!(IntSegment_sel0 < 3))) & (0 < IntSegment_sel1)) & (!(IntSegment_sel2 != IntSegment_sel3))) & (!(0 < IntSegment_sel4))) & ((0 + 1) < IntSegment_sel5)) & (!(IntSegment_sel6 != IntSegment_sel7))) & (!(0 < IntSegment_sel8))) & (((0 + 1) + 1) < IntSegment_sel9)) & (IntSegment_sel10 != IntSegment_sel11))
  ReturnResults:
    [0] false
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    IntSegment:    IntSegment (elem=int, ops=12, store=0, select=12)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=(addr_1 + 0), res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=(addr_1 + 0), res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=(addr_1 + (0 + 1)), res=IntSegment_sel6
        op[7]: select idx=addr_0, res=IntSegment_sel7
        op[8]: select idx=(addr_1 + (0 + 1)), res=IntSegment_sel8
        op[9]: select idx=addr_0, res=IntSegment_sel9
        op[10]: select idx=(addr_1 + ((0 + 1) + 1)), res=IntSegment_sel10
        op[11]: select idx=addr_0, res=IntSegment_sel11
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select((addr_1 + 0), IntSegment_sel2).select(addr_0, IntSegment_sel3).select((addr_1 + 0), IntSegment_sel4).select(addr_0, IntSegment_sel5).select((addr_1 + (0 + 1)), IntSegment_sel6).select(addr_0, IntSegment_sel7).select((addr_1 + (0 + 1)), IntSegment_sel8).select(addr_0, IntSegment_sel9).select((addr_1 + ((0 + 1) + 1)), IntSegment_sel10).select(addr_0, IntSegment_sel11)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #6 ----
Interpreter{
  PC: ((((((((true & (!(IntSegment_sel0 < 3))) & (0 < IntSegment_sel1)) & (!(IntSegment_sel2 != IntSegment_sel3))) & (!(0 < IntSegment_sel4))) & ((0 + 1) < IntSegment_sel5)) & (!(IntSegment_sel6 != IntSegment_sel7))) & (!(0 < IntSegment_sel8))) & (!(((0 + 1) + 1) < IntSegment_sel9)))
  ReturnResults:
    [0] true
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    IntSegment:    IntSegment (elem=int, ops=10, store=0, select=10)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=(addr_1 + 0), res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=(addr_1 + 0), res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=(addr_1 + (0 + 1)), res=IntSegment_sel6
        op[7]: select idx=addr_0, res=IntSegment_sel7
        op[8]: select idx=(addr_1 + (0 + 1)), res=IntSegment_sel8
        op[9]: select idx=addr_0, res=IntSegment_sel9
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select((addr_1 + 0), IntSegment_sel2).select(addr_0, IntSegment_sel3).select((addr_1 + 0), IntSegment_sel4).select(addr_0, IntSegment_sel5).select((addr_1 + (0 + 1)), IntSegment_sel6).select(addr_0, IntSegment_sel7).select((addr_1 + (0 + 1)), IntSegment_sel8).select(addr_0, IntSegment_sel9)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #7 ----
Interpreter{
  PC: (((((((((((true & (!(IntSegment_sel0 < 3))) & (0 < IntSegment_sel1)) & (!(IntSegment_sel2 != IntSegment_sel3))) & (!(0 < IntSegment_sel4))) & ((0 + 1) < IntSegment_sel5)) & (!(IntSegment_sel6 != IntSegment_sel7))) & (!(0 < IntSegment_sel8))) & (((0 + 1) + 1) < IntSegment_sel9)) & (!(IntSegment_sel10 != IntSegment_sel11))) & (!(0 < IntSegment_sel12))) & (!((((0 + 1) + 1) + 1) < IntSegment_sel13)))
  ReturnResults:
    [0] true
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    IntSegment:    IntSegment (elem=int, ops=14, store=0, select=14)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=(addr_1 + 0), res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=(addr_1 + 0), res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=(addr_1 + (0 + 1)), res=IntSegment_sel6
        op[7]: select idx=addr_0, res=IntSegment_sel7
        op[8]: select idx=(addr_1 + (0 + 1)), res=IntSegment_sel8
        op[9]: select idx=addr_0, res=IntSegment_sel9
        op[10]: select idx=(addr_1 + ((0 + 1) + 1)), res=IntSegment_sel10
        op[11]: select idx=addr_0, res=IntSegment_sel11
        op[12]: select idx=(addr_1 + ((0 + 1) + 1)), res=IntSegment_sel12
        op[13]: select idx=addr_0, res=IntSegment_sel13
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select((addr_1 + 0), IntSegment_sel2).select(addr_0, IntSegment_sel3).select((addr_1 + 0), IntSegment_sel4).select(addr_0, IntSegment_sel5).select((addr_1 + (0 + 1)), IntSegment_sel6).select(addr_0, IntSegment_sel7).select((addr_1 + (0 + 1)), IntSegment_sel8).select(addr_0, IntSegment_sel9).select((addr_1 + ((0 + 1) + 1)), IntSegment_sel10).select(addr_0, IntSegment_sel11).select((addr_1 + ((0 + 1) + 1)), IntSegment_sel12).select(addr_0, IntSegment_sel13)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #8 ----
Interpreter{
  PC: (((((((((((((true & (!(IntSegment_sel0 < 3))) & (0 < IntSegment_sel1)) & (!(IntSegment_sel2 != IntSegment_sel3))) & (!(0 < IntSegment_sel4))) & ((0 + 1) < IntSegment_sel5)) & (!(IntSegment_sel6 != IntSegment_sel7))) & (0 < IntSegment_sel8)) & (!((0 + 1) == 0))) & (!((0 + 1) != 0))) & ((0 + 1) < IntSegment_sel9)) & (!((0 + 1) == (0 + 1)))) & ((0 + 1) != (0 + 1))) & (IntSegment_sel10 != 0))
  ReturnResults:
    [0] false
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 3, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
      P[2]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=11, store=0, select=11)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=(addr_1 + 0), res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=(addr_1 + 0), res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=(addr_1 + (0 + 1)), res=IntSegment_sel6
        op[7]: select idx=addr_0, res=IntSegment_sel7
        op[8]: select idx=(addr_1 + (0 + 1)), res=IntSegment_sel8
        op[9]: select idx=(addr_1 + (0 + 1)), res=IntSegment_sel9
        op[10]: select idx=(ArraySegment_sel0 + (0 + 1)), res=IntSegment_sel10
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=1, store=0, select=1)
        op[0]: select idx=(addr_2 + (0 + 1)), res=ArraySegment_sel0
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select((addr_1 + 0), IntSegment_sel2).select(addr_0, IntSegment_sel3).select((addr_1 + 0), IntSegment_sel4).select(addr_0, IntSegment_sel5).select((addr_1 + (0 + 1)), IntSegment_sel6).select(addr_0, IntSegment_sel7).select((addr_1 + (0 + 1)), IntSegment_sel8).select((addr_1 + (0 + 1)), IntSegment_sel9).select((ArraySegment_sel0 + (0 + 1)), IntSegment_sel10)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment.select((addr_2 + (0 + 1)), ArraySegment_sel0)
}

UNSAT
---- State #9 ----
Interpreter{
  PC: ((((((((((((((((true & (!(IntSegment_sel0 < 3))) & (0 < IntSegment_sel1)) & (!(IntSegment_sel2 != IntSegment_sel3))) & (!(0 < IntSegment_sel4))) & ((0 + 1) < IntSegment_sel5)) & (!(IntSegment_sel6 != IntSegment_sel7))) & (0 < IntSegment_sel8)) & (!((0 + 1) == 0))) & (!((0 + 1) != 0))) & ((0 + 1) < IntSegment_sel9)) & (!((0 + 1) == (0 + 1)))) & ((0 + 1) != (0 + 1))) & (!(IntSegment_sel10 != 0))) & (((0 + 1) + 1) < IntSegment_sel11)) & ((0 + 1) == ((0 + 1) + 1))) & (IntSegment_sel12 != 1))
  ReturnResults:
    [0] false
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 3, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
      P[2]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=13, store=0, select=13)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=(addr_1 + 0), res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=(addr_1 + 0), res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=(addr_1 + (0 + 1)), res=IntSegment_sel6
        op[7]: select idx=addr_0, res=IntSegment_sel7
        op[8]: select idx=(addr_1 + (0 + 1)), res=IntSegment_sel8
        op[9]: select idx=(addr_1 + (0 + 1)), res=IntSegment_sel9
        op[10]: select idx=(ArraySegment_sel0 + (0 + 1)), res=IntSegment_sel10
        op[11]: select idx=(addr_1 + (0 + 1)), res=IntSegment_sel11
        op[12]: select idx=(ArraySegment_sel1 + ((0 + 1) + 1)), res=IntSegment_sel12
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=2, store=0, select=2)
        op[0]: select idx=(addr_2 + (0 + 1)), res=ArraySegment_sel0
        op[1]: select idx=(addr_2 + (0 + 1)), res=ArraySegment_sel1
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select((addr_1 + 0), IntSegment_sel2).select(addr_0, IntSegment_sel3).select((addr_1 + 0), IntSegment_sel4).select(addr_0, IntSegment_sel5).select((addr_1 + (0 + 1)), IntSegment_sel6).select(addr_0, IntSegment_sel7).select((addr_1 + (0 + 1)), IntSegment_sel8).select((addr_1 + (0 + 1)), IntSegment_sel9).select((ArraySegment_sel0 + (0 + 1)), IntSegment_sel10).select((addr_1 + (0 + 1)), IntSegment_sel11).select((ArraySegment_sel1 + ((0 + 1) + 1)), IntSegment_sel12)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment.select((addr_2 + (0 + 1)), ArraySegment_sel0).select((addr_2 + (0 + 1)), ArraySegment_sel1)
}

UNSAT

==================== ReallyMultiDimensionalArray ====================
CFG построен для функции с 4 блоками
interpretDynamically: &array[1:int]
interpretDynamically: *t0
interpretDynamically: &t1[2:int]
interpretDynamically: *t2
interpretDynamically: &t3[3:int]
interpretDynamically: *t4
interpretDynamically: t5 != 12345:int
interpretDynamically: if t6 goto 1 else 3
resolveExpression: t5 != 12345:int
resolveExpression: *t4
resolveExpression: *t2
resolveExpression: *t0
resolveExpression: parameter array : [][][]int
resolveExpression: 1:int
resolveExpression: 2:int
resolveExpression: 3:int
resolveExpression: 12345:int
interpretDynamically: &array[1:int]
interpretDynamically: &array[1:int]
interpretDynamically: *t7
interpretDynamically: *t12
interpretDynamically: &t13[2:int]
interpretDynamically: &t8[2:int]
interpretDynamically: *t9
interpretDynamically: *t14
interpretDynamically: &t15[3:int]
interpretDynamically: *t16
interpretDynamically: t17 - 24690:int
interpretDynamically: &t15[3:int]
interpretDynamically: &t10[3:int]
interpretDynamically: *t11 = 12345:int
resolveExpression: 12345:int
resolveExpression: *t9
resolveExpression: *t7
resolveExpression: parameter array : [][][]int
resolveExpression: 1:int
resolveExpression: 2:int
resolveExpression: 3:int
interpretDynamically: jump 2
interpretDynamically: return array
resolveExpression: parameter array : [][][]int
interpretDynamically: *t19 = t18
resolveExpression: t17 - 24690:int
resolveExpression: *t16
resolveExpression: *t14
resolveExpression: *t12
resolveExpression: parameter array : [][][]int
resolveExpression: 1:int
resolveExpression: 2:int
resolveExpression: 3:int
resolveExpression: 24690:int
resolveExpression: *t14
resolveExpression: *t12
resolveExpression: parameter array : [][][]int
resolveExpression: 1:int
resolveExpression: 2:int
resolveExpression: 3:int
interpretDynamically: jump 2
interpretDynamically: return array
resolveExpression: parameter array : [][][]int
States: 2
---- State #0 ----
Interpreter{
  PC: (true & (IntSegment_sel0 != 12345))
  ReturnResults:
    [0] &(addr_0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=array, dim=2, base=int)
    IntSegment:    IntSegment (elem=int, ops=2, store=1, select=1)
        op[0]: select idx=(ArraySegment_sel1 + 3), res=IntSegment_sel0
        op[1]: store idx=(ArraySegment_sel3 + 3), val=12345
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=(addr_0 + 1), res=ArraySegment_sel0
        op[1]: select idx=(ArraySegment_sel0 + 2), res=ArraySegment_sel1
        op[2]: select idx=(addr_0 + 1), res=ArraySegment_sel2
        op[3]: select idx=(ArraySegment_sel2 + 2), res=ArraySegment_sel3
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select((ArraySegment_sel1 + 3), IntSegment_sel0).store((ArraySegment_sel3 + 3), 12345)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment.select((addr_0 + 1), ArraySegment_sel0).select((ArraySegment_sel0 + 2), ArraySegment_sel1).select((addr_0 + 1), ArraySegment_sel2).select((ArraySegment_sel2 + 2), ArraySegment_sel3)
}

SAT
Model:
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #xffffffffffffcfc6)
addr_0 -> #xffffffffffffffff
ArraySegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
ArraySegment_sel3 -> #x0000000000000000
ArraySegment_sel0 -> #x0000000000000000
ArraySegment_sel2 -> #x0000000000000000
IntSegment_sel0 -> #xffffffffffffcfc6
ArraySegment_sel1 -> #x0000000000000000

---- State #1 ----
Interpreter{
  PC: (true & (!(IntSegment_sel0 != 12345)))
  ReturnResults:
    [0] &(addr_0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=array, dim=2, base=int)
    IntSegment:    IntSegment (elem=int, ops=3, store=1, select=2)
        op[0]: select idx=(ArraySegment_sel1 + 3), res=IntSegment_sel0
        op[1]: select idx=(ArraySegment_sel3 + 3), res=IntSegment_sel1
        op[2]: store idx=(ArraySegment_sel5 + 3), val=(IntSegment_sel1 - 24690)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=6, store=0, select=6)
        op[0]: select idx=(addr_0 + 1), res=ArraySegment_sel0
        op[1]: select idx=(ArraySegment_sel0 + 2), res=ArraySegment_sel1
        op[2]: select idx=(addr_0 + 1), res=ArraySegment_sel2
        op[3]: select idx=(ArraySegment_sel2 + 2), res=ArraySegment_sel3
        op[4]: select idx=(addr_0 + 1), res=ArraySegment_sel4
        op[5]: select idx=(ArraySegment_sel4 + 2), res=ArraySegment_sel5
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select((ArraySegment_sel1 + 3), IntSegment_sel0).select((ArraySegment_sel3 + 3), IntSegment_sel1).store((ArraySegment_sel5 + 3), (IntSegment_sel1 - 24690))
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment.select((addr_0 + 1), ArraySegment_sel0).select((ArraySegment_sel0 + 2), ArraySegment_sel1).select((addr_0 + 1), ArraySegment_sel2).select((ArraySegment_sel2 + 2), ArraySegment_sel3).select((addr_0 + 1), ArraySegment_sel4).select((ArraySegment_sel4 + 2), ArraySegment_sel5)
}

SAT
Model:
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000003039)
addr_0 -> #xffffffffffffffff
ArraySegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
ArraySegment_sel3 -> #x0000000000000000
ArraySegment_sel0 -> #x0000000000000000
ArraySegment_sel4 -> #x0000000000000000
IntSegment_sel1 -> #x0000000000003039
ArraySegment_sel2 -> #x0000000000000000
ArraySegment_sel5 -> #x0000000000000000
ArraySegment_sel1 -> #x0000000000000000
IntSegment_sel0 -> #x0000000000003039


==================== FillMultiArrayWithArray ====================
CFG построен для функции с 9 блоками
interpretDynamically: len_value < 2:int
interpretDynamically: if t0 goto 1 else 2
resolveExpression: len_value < 2:int
resolveExpression: parameter len_value : int
resolveExpression: 2:int
interpretDynamically: new [0][]int (makeslice)
interpretDynamically: slice t1[:0:int]
interpretDynamically: len(value)
resolveExpression: len(value)
interpretDynamically: jump 3
interpretDynamically: phi [2: -1:int, 4: t5]
resolveExpression: -1:int
interpretDynamically: t4 + 1:int
interpretDynamically: return t2
resolveExpression: slice t1[:0:int]
resolveExpression: new [0][]int (makeslice)
interpretDynamically: t5 < t3
interpretDynamically: if t6 goto 4 else 5
resolveExpression: t5 < t3
resolveExpression: t4 + 1:int
resolveExpression: phi [2: -1:int, 4: t5]
resolveExpression: 1:int
resolveExpression: len(value)
interpretDynamically: make [][]int 3:int 3:int
interpretDynamically: &value[t5]
interpretDynamically: *t7
interpretDynamically: t8 + t5
interpretDynamically: jump 8
interpretDynamically: phi [5: 0:int, 6: t13] #i
resolveExpression: 0:int
interpretDynamically: t14 < 3:int
interpretDynamically: if t15 goto 6 else 7
resolveExpression: t14 < 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 3:int
interpretDynamically: &t11[t14]
interpretDynamically: *t12 = value
resolveExpression: parameter value : []int
resolveExpression: make [][]int 3:int 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
interpretDynamically: t14 + 1:int
interpretDynamically: return t11
resolveExpression: make [][]int 3:int 3:int
interpretDynamically: &value[t5]
interpretDynamically: *t10 = t9
resolveExpression: t8 + t5
resolveExpression: *t7
resolveExpression: parameter value : []int
resolveExpression: t4 + 1:int
resolveExpression: phi [2: -1:int, 4: t5]
resolveExpression: 1:int
resolveExpression: t4 + 1:int
resolveExpression: phi [2: -1:int, 4: t5]
resolveExpression: 1:int
resolveExpression: parameter value : []int
resolveExpression: t4 + 1:int
resolveExpression: phi [2: -1:int, 4: t5]
resolveExpression: 1:int
interpretDynamically: jump 3
interpretDynamically: phi [2: -1:int, 4: t5]
resolveExpression: t4 + 1:int
resolveExpression: phi [2: -1:int, 4: t5]
resolveExpression: 1:int
interpretDynamically: t4 + 1:int
interpretDynamically: t5 < t3
interpretDynamically: if t6 goto 4 else 5
resolveExpression: t5 < t3
resolveExpression: t4 + 1:int
resolveExpression: phi [2: -1:int, 4: t5]
resolveExpression: 1:int
resolveExpression: len(value)
interpretDynamically: &value[t5]
interpretDynamically: make [][]int 3:int 3:int
interpretDynamically: jump 8
interpretDynamically: *t7
interpretDynamically: t8 + t5
interpretDynamically: &value[t5]
interpretDynamically: *t10 = t9
resolveExpression: t8 + t5
resolveExpression: *t7
resolveExpression: parameter value : []int
resolveExpression: t4 + 1:int
resolveExpression: phi [2: -1:int, 4: t5]
resolveExpression: 1:int
resolveExpression: t4 + 1:int
resolveExpression: phi [2: -1:int, 4: t5]
resolveExpression: 1:int
resolveExpression: parameter value : []int
resolveExpression: t4 + 1:int
resolveExpression: phi [2: -1:int, 4: t5]
resolveExpression: 1:int
interpretDynamically: jump 3
interpretDynamically: jump 8
interpretDynamically: phi [5: 0:int, 6: t13] #i
resolveExpression: t14 + 1:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 1:int
interpretDynamically: t14 < 3:int
interpretDynamically: if t15 goto 6 else 7
resolveExpression: t14 < 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 3:int
interpretDynamically: &t11[t14]
interpretDynamically: *t12 = value
resolveExpression: parameter value : []int
resolveExpression: make [][]int 3:int 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
interpretDynamically: t14 + 1:int
interpretDynamically: jump 8
interpretDynamically: phi [2: -1:int, 4: t5]
resolveExpression: t4 + 1:int
resolveExpression: phi [2: -1:int, 4: t5]
resolveExpression: 1:int
interpretDynamically: t4 + 1:int
interpretDynamically: t5 < t3
interpretDynamically: if t6 goto 4 else 5
resolveExpression: t5 < t3
resolveExpression: t4 + 1:int
resolveExpression: phi [2: -1:int, 4: t5]
resolveExpression: 1:int
resolveExpression: len(value)
interpretDynamically: &value[t5]
interpretDynamically: *t7
interpretDynamically: make [][]int 3:int 3:int
interpretDynamically: jump 8
interpretDynamically: phi [5: 0:int, 6: t13] #i
resolveExpression: 0:int
interpretDynamically: phi [5: 0:int, 6: t13] #i
resolveExpression: 0:int
interpretDynamically: t14 < 3:int
interpretDynamically: if t15 goto 6 else 7
resolveExpression: t14 < 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 3:int
interpretDynamically: &t11[t14]
interpretDynamically: *t12 = value
resolveExpression: parameter value : []int
resolveExpression: make [][]int 3:int 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
interpretDynamically: t14 + 1:int
interpretDynamically: jump 8
interpretDynamically: phi [5: 0:int, 6: t13] #i
resolveExpression: t14 + 1:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 1:int
interpretDynamically: t14 < 3:int
interpretDynamically: if t15 goto 6 else 7
resolveExpression: t14 < 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 3:int
interpretDynamically: return t11
resolveExpression: make [][]int 3:int 3:int
interpretDynamically: return t11
resolveExpression: make [][]int 3:int 3:int
interpretDynamically: phi [5: 0:int, 6: t13] #i
resolveExpression: t14 + 1:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 1:int
interpretDynamically: t14 < 3:int
interpretDynamically: if t15 goto 6 else 7
resolveExpression: t14 < 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 3:int
interpretDynamically: &t11[t14]
interpretDynamically: return t11
resolveExpression: make [][]int 3:int 3:int
interpretDynamically: *t12 = value
resolveExpression: parameter value : []int
resolveExpression: make [][]int 3:int 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
interpretDynamically: t14 + 1:int
interpretDynamically: jump 8
interpretDynamically: phi [5: 0:int, 6: t13] #i
resolveExpression: t14 + 1:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 1:int
interpretDynamically: t14 < 3:int
interpretDynamically: if t15 goto 6 else 7
resolveExpression: t14 < 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 3:int
interpretDynamically: return t11
resolveExpression: make [][]int 3:int 3:int
interpretDynamically: return t11
resolveExpression: make [][]int 3:int 3:int
interpretDynamically: t8 + t5
interpretDynamically: &value[t5]
interpretDynamically: *t10 = t9
resolveExpression: t8 + t5
resolveExpression: *t7
resolveExpression: parameter value : []int
resolveExpression: t4 + 1:int
resolveExpression: phi [2: -1:int, 4: t5]
resolveExpression: 1:int
resolveExpression: t4 + 1:int
resolveExpression: phi [2: -1:int, 4: t5]
resolveExpression: 1:int
resolveExpression: parameter value : []int
resolveExpression: t4 + 1:int
resolveExpression: phi [2: -1:int, 4: t5]
resolveExpression: 1:int
interpretDynamically: jump 3
interpretDynamically: phi [2: -1:int, 4: t5]
resolveExpression: t4 + 1:int
resolveExpression: phi [2: -1:int, 4: t5]
resolveExpression: 1:int
interpretDynamically: t4 + 1:int
interpretDynamically: t5 < t3
interpretDynamically: if t6 goto 4 else 5
resolveExpression: t5 < t3
resolveExpression: t4 + 1:int
resolveExpression: phi [2: -1:int, 4: t5]
resolveExpression: 1:int
resolveExpression: len(value)
interpretDynamically: &value[t5]
interpretDynamically: *t7
interpretDynamically: t8 + t5
interpretDynamically: make [][]int 3:int 3:int
interpretDynamically: jump 8
interpretDynamically: phi [5: 0:int, 6: t13] #i
resolveExpression: 0:int
interpretDynamically: &value[t5]
interpretDynamically: *t10 = t9
resolveExpression: t8 + t5
resolveExpression: *t7
resolveExpression: parameter value : []int
resolveExpression: t4 + 1:int
resolveExpression: phi [2: -1:int, 4: t5]
resolveExpression: 1:int
resolveExpression: t4 + 1:int
resolveExpression: phi [2: -1:int, 4: t5]
resolveExpression: 1:int
resolveExpression: parameter value : []int
resolveExpression: t4 + 1:int
resolveExpression: phi [2: -1:int, 4: t5]
resolveExpression: 1:int
interpretDynamically: jump 3
interpretDynamically: phi [2: -1:int, 4: t5]
resolveExpression: t4 + 1:int
resolveExpression: phi [2: -1:int, 4: t5]
resolveExpression: 1:int
interpretDynamically: t4 + 1:int
interpretDynamically: t5 < t3
interpretDynamically: if t6 goto 4 else 5
resolveExpression: t5 < t3
resolveExpression: t4 + 1:int
resolveExpression: phi [2: -1:int, 4: t5]
resolveExpression: 1:int
resolveExpression: len(value)
interpretDynamically: &value[t5]
interpretDynamically: *t7
interpretDynamically: make [][]int 3:int 3:int
interpretDynamically: jump 8
interpretDynamically: phi [5: 0:int, 6: t13] #i
resolveExpression: 0:int
interpretDynamically: t14 < 3:int
interpretDynamically: t8 + t5
interpretDynamically: &value[t5]
interpretDynamically: *t10 = t9
resolveExpression: t8 + t5
resolveExpression: *t7
resolveExpression: parameter value : []int
resolveExpression: t4 + 1:int
resolveExpression: phi [2: -1:int, 4: t5]
resolveExpression: 1:int
resolveExpression: t4 + 1:int
resolveExpression: phi [2: -1:int, 4: t5]
resolveExpression: 1:int
resolveExpression: parameter value : []int
resolveExpression: t4 + 1:int
resolveExpression: phi [2: -1:int, 4: t5]
resolveExpression: 1:int
interpretDynamically: jump 3
interpretDynamically: if t15 goto 6 else 7
resolveExpression: t14 < 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 3:int
interpretDynamically: return t11
resolveExpression: make [][]int 3:int 3:int
interpretDynamically: &t11[t14]
interpretDynamically: *t12 = value
resolveExpression: parameter value : []int
resolveExpression: make [][]int 3:int 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
interpretDynamically: t14 + 1:int
interpretDynamically: jump 8
interpretDynamically: phi [5: 0:int, 6: t13] #i
resolveExpression: t14 + 1:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 1:int
interpretDynamically: t14 < 3:int
interpretDynamically: if t15 goto 6 else 7
resolveExpression: t14 < 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 3:int
interpretDynamically: &t11[t14]
interpretDynamically: return t11
resolveExpression: make [][]int 3:int 3:int
interpretDynamically: *t12 = value
resolveExpression: parameter value : []int
resolveExpression: make [][]int 3:int 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
interpretDynamically: t14 + 1:int
interpretDynamically: jump 8
interpretDynamically: phi [5: 0:int, 6: t13] #i
resolveExpression: t14 + 1:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 1:int
interpretDynamically: t14 < 3:int
interpretDynamically: if t15 goto 6 else 7
resolveExpression: t14 < 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 3:int
interpretDynamically: &t11[t14]
interpretDynamically: *t12 = value
resolveExpression: parameter value : []int
resolveExpression: make [][]int 3:int 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
interpretDynamically: t14 + 1:int
interpretDynamically: jump 8
interpretDynamically: phi [5: 0:int, 6: t13] #i
resolveExpression: t14 + 1:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 1:int
interpretDynamically: t14 < 3:int
interpretDynamically: if t15 goto 6 else 7
resolveExpression: t14 < 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 3:int
interpretDynamically: &t11[t14]
interpretDynamically: return t11
resolveExpression: make [][]int 3:int 3:int
interpretDynamically: *t12 = value
resolveExpression: parameter value : []int
resolveExpression: make [][]int 3:int 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
interpretDynamically: t14 + 1:int
interpretDynamically: jump 8
interpretDynamically: phi [5: 0:int, 6: t13] #i
resolveExpression: t14 + 1:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 1:int
interpretDynamically: t14 < 3:int
interpretDynamically: if t15 goto 6 else 7
resolveExpression: t14 < 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 3:int
interpretDynamically: return t11
resolveExpression: make [][]int 3:int 3:int
interpretDynamically: &t11[t14]
interpretDynamically: *t12 = value
resolveExpression: parameter value : []int
resolveExpression: make [][]int 3:int 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
interpretDynamically: t14 + 1:int
interpretDynamically: jump 8
interpretDynamically: return t11
resolveExpression: make [][]int 3:int 3:int
interpretDynamically: t14 < 3:int
interpretDynamically: if t15 goto 6 else 7
resolveExpression: t14 < 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 3:int
interpretDynamically: &t11[t14]
interpretDynamically: *t12 = value
resolveExpression: parameter value : []int
resolveExpression: make [][]int 3:int 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
interpretDynamically: t14 + 1:int
interpretDynamically: jump 8
interpretDynamically: phi [5: 0:int, 6: t13] #i
resolveExpression: t14 + 1:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 1:int
interpretDynamically: t14 < 3:int
interpretDynamically: if t15 goto 6 else 7
resolveExpression: t14 < 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 3:int
interpretDynamically: return t11
resolveExpression: make [][]int 3:int 3:int
interpretDynamically: &t11[t14]
interpretDynamically: *t12 = value
resolveExpression: parameter value : []int
resolveExpression: make [][]int 3:int 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
interpretDynamically: return t11
resolveExpression: make [][]int 3:int 3:int
interpretDynamically: t14 + 1:int
interpretDynamically: jump 8
interpretDynamically: phi [5: 0:int, 6: t13] #i
resolveExpression: t14 + 1:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 1:int
interpretDynamically: t14 < 3:int
interpretDynamically: if t15 goto 6 else 7
resolveExpression: t14 < 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 3:int
interpretDynamically: &t11[t14]
interpretDynamically: *t12 = value
resolveExpression: parameter value : []int
resolveExpression: make [][]int 3:int 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
interpretDynamically: t14 + 1:int
interpretDynamically: return t11
resolveExpression: make [][]int 3:int 3:int
interpretDynamically: jump 8
interpretDynamically: phi [5: 0:int, 6: t13] #i
resolveExpression: t14 + 1:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 1:int
interpretDynamically: t14 < 3:int
interpretDynamically: if t15 goto 6 else 7
resolveExpression: t14 < 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 3:int
interpretDynamically: return t11
resolveExpression: make [][]int 3:int 3:int
interpretDynamically: &t11[t14]
interpretDynamically: *t12 = value
resolveExpression: parameter value : []int
resolveExpression: make [][]int 3:int 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
interpretDynamically: t14 + 1:int
interpretDynamically: jump 8
interpretDynamically: phi [5: 0:int, 6: t13] #i
resolveExpression: t14 + 1:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 1:int
interpretDynamically: t14 < 3:int
interpretDynamically: if t15 goto 6 else 7
resolveExpression: t14 < 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 3:int
interpretDynamically: return t11
resolveExpression: make [][]int 3:int 3:int
interpretDynamically: &t11[t14]
interpretDynamically: *t12 = value
resolveExpression: parameter value : []int
resolveExpression: make [][]int 3:int 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
interpretDynamically: t14 + 1:int
interpretDynamically: jump 8
interpretDynamically: t14 < 3:int
interpretDynamically: if t15 goto 6 else 7
resolveExpression: t14 < 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 3:int
interpretDynamically: return t11
resolveExpression: make [][]int 3:int 3:int
interpretDynamically: &t11[t14]
interpretDynamically: &t11[t14]
interpretDynamically: *t12 = value
resolveExpression: parameter value : []int
resolveExpression: make [][]int 3:int 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
interpretDynamically: t14 + 1:int
interpretDynamically: jump 8
interpretDynamically: phi [5: 0:int, 6: t13] #i
resolveExpression: t14 + 1:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 1:int
interpretDynamically: t14 < 3:int
interpretDynamically: if t15 goto 6 else 7
resolveExpression: t14 < 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 3:int
interpretDynamically: &t11[t14]
interpretDynamically: *t12 = value
resolveExpression: parameter value : []int
resolveExpression: make [][]int 3:int 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
interpretDynamically: return t11
resolveExpression: make [][]int 3:int 3:int
interpretDynamically: t14 + 1:int
interpretDynamically: jump 8
interpretDynamically: *t12 = value
resolveExpression: parameter value : []int
resolveExpression: make [][]int 3:int 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
interpretDynamically: t14 + 1:int
interpretDynamically: jump 8
interpretDynamically: phi [5: 0:int, 6: t13] #i
resolveExpression: t14 + 1:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 1:int
interpretDynamically: t14 < 3:int
interpretDynamically: if t15 goto 6 else 7
resolveExpression: t14 < 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 3:int
interpretDynamically: return t11
resolveExpression: make [][]int 3:int 3:int
interpretDynamically: &t11[t14]
interpretDynamically: *t12 = value
resolveExpression: parameter value : []int
resolveExpression: make [][]int 3:int 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
interpretDynamically: t14 + 1:int
interpretDynamically: jump 8
interpretDynamically: phi [5: 0:int, 6: t13] #i
resolveExpression: t14 + 1:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 1:int
interpretDynamically: t14 < 3:int
interpretDynamically: if t15 goto 6 else 7
resolveExpression: t14 < 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 3:int
interpretDynamically: &t11[t14]
interpretDynamically: return t11
resolveExpression: make [][]int 3:int 3:int
interpretDynamically: *t12 = value
resolveExpression: parameter value : []int
resolveExpression: make [][]int 3:int 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
interpretDynamically: t14 + 1:int
interpretDynamically: jump 8
interpretDynamically: phi [5: 0:int, 6: t13] #i
resolveExpression: t14 + 1:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 1:int
interpretDynamically: t14 < 3:int
interpretDynamically: if t15 goto 6 else 7
resolveExpression: t14 < 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 3:int
interpretDynamically: &t11[t14]
interpretDynamically: *t12 = value
resolveExpression: parameter value : []int
resolveExpression: make [][]int 3:int 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
interpretDynamically: t14 + 1:int
interpretDynamically: return t11
resolveExpression: make [][]int 3:int 3:int
interpretDynamically: jump 8
interpretDynamically: phi [5: 0:int, 6: t13] #i
resolveExpression: t14 + 1:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 1:int
interpretDynamically: t14 < 3:int
interpretDynamically: if t15 goto 6 else 7
resolveExpression: t14 < 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 3:int
interpretDynamically: &t11[t14]
interpretDynamically: return t11
resolveExpression: make [][]int 3:int 3:int
interpretDynamically: *t12 = value
resolveExpression: parameter value : []int
resolveExpression: make [][]int 3:int 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
interpretDynamically: t14 + 1:int
interpretDynamically: jump 8
interpretDynamically: &t11[t14]
interpretDynamically: *t12 = value
resolveExpression: parameter value : []int
resolveExpression: make [][]int 3:int 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
interpretDynamically: t14 + 1:int
interpretDynamically: jump 8
interpretDynamically: phi [5: 0:int, 6: t13] #i
resolveExpression: t14 + 1:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 1:int
interpretDynamically: t14 < 3:int
interpretDynamically: if t15 goto 6 else 7
resolveExpression: t14 < 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 3:int
interpretDynamically: &t11[t14]
interpretDynamically: *t12 = value
resolveExpression: parameter value : []int
resolveExpression: make [][]int 3:int 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
interpretDynamically: return t11
resolveExpression: make [][]int 3:int 3:int
interpretDynamically: t14 + 1:int
interpretDynamically: jump 8
interpretDynamically: phi [5: 0:int, 6: t13] #i
resolveExpression: t14 + 1:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 1:int
interpretDynamically: t14 < 3:int
interpretDynamically: if t15 goto 6 else 7
resolveExpression: t14 < 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 3:int
interpretDynamically: &t11[t14]
interpretDynamically: *t12 = value
resolveExpression: parameter value : []int
resolveExpression: make [][]int 3:int 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
interpretDynamically: return t11
resolveExpression: make [][]int 3:int 3:int
interpretDynamically: t14 + 1:int
interpretDynamically: jump 8
interpretDynamically: phi [5: 0:int, 6: t13] #i
resolveExpression: t14 + 1:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 1:int
interpretDynamically: t14 < 3:int
interpretDynamically: if t15 goto 6 else 7
resolveExpression: t14 < 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
resolveExpression: 3:int
interpretDynamically: return t11
resolveExpression: make [][]int 3:int 3:int
interpretDynamically: &t11[t14]
interpretDynamically: *t12 = value
resolveExpression: parameter value : []int
resolveExpression: make [][]int 3:int 3:int
resolveExpression: phi [5: 0:int, 6: t13] #i
interpretDynamically: t14 + 1:int
interpretDynamically: jump 8
States: 26
---- State #0 ----
Interpreter{
  PC: (true & (IntSegment_sel0 < 2))
  ReturnResults:
    [0] &(addr_1)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 1, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    AllocationsList:
      A[0]: &(addr_1) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=1, store=0, select=1)
        op[0]: select idx=addr_0, res=IntSegment_sel0
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 6
    C[0]: (addr_0 != addr_1)
    C[1]: IntSegment.select(addr_0, IntSegment_sel0)
    C[2]: BoolSegment
    C[3]: StringSegment
    C[4]: FloatSegment
    C[5]: ArraySegment
}

SAT
Model:
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000001)
addr_1 -> #xffffffffffffffff
addr_0 -> #x0000000000000000
IntSegment_sel0 -> #x0000000000000001

---- State #1 ----
Interpreter{
  PC: (((true & (!(IntSegment_sel0 < 2))) & (!((-1 + 1) < len_parameter_value_____int))) & (!(0 < 3)))
  ReturnResults:
    [0] &(addr_1)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 1, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    AllocationsList:
      A[0]: &(addr_1) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=1, store=0, select=1)
        op[0]: select idx=addr_0, res=IntSegment_sel0
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 6
    C[0]: (addr_0 != addr_1)
    C[1]: IntSegment.select(addr_0, IntSegment_sel0)
    C[2]: BoolSegment
    C[3]: StringSegment
    C[4]: FloatSegment
    C[5]: ArraySegment
}

UNSAT
---- State #2 ----
Interpreter{
  PC: (((((true & (!(IntSegment_sel0 < 2))) & ((-1 + 1) < len_parameter_value_____int)) & (!(((-1 + 1) + 1) < len_parameter_value_____int))) & (0 < 3)) & (!((0 + 1) < 3)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    AllocationsList:
      A[0]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=3, store=1, select=2)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=(addr_1 + (-1 + 1)), res=IntSegment_sel1
        op[2]: store idx=(addr_1 + (-1 + 1)), val=(IntSegment_sel1 + (-1 + 1))
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=1, store=1, select=0)
        op[0]: store idx=(addr_2 + 0), val=addr_1
    StructSegments: <empty>
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_2)
    C[1]: (addr_1 != addr_2)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select((addr_1 + (-1 + 1)), IntSegment_sel1).store((addr_1 + (-1 + 1)), (IntSegment_sel1 + (-1 + 1)))
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment.store((addr_2 + 0), addr_1)
}

UNSAT
---- State #3 ----
Interpreter{
  PC: (((((true & (!(IntSegment_sel0 < 2))) & (!((-1 + 1) < len_parameter_value_____int))) & (0 < 3)) & ((0 + 1) < 3)) & (!(((0 + 1) + 1) < 3)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    AllocationsList:
      A[0]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=1, store=0, select=1)
        op[0]: select idx=addr_0, res=IntSegment_sel0
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=2, store=2, select=0)
        op[0]: store idx=(addr_2 + 0), val=addr_1
        op[1]: store idx=(addr_2 + (0 + 1)), val=addr_1
    StructSegments: <empty>
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_2)
    C[1]: (addr_1 != addr_2)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0)
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment.store((addr_2 + 0), addr_1).store((addr_2 + (0 + 1)), addr_1)
}

UNSAT
---- State #4 ----
Interpreter{
  PC: ((((((true & (!(IntSegment_sel0 < 2))) & (!((-1 + 1) < len_parameter_value_____int))) & (0 < 3)) & ((0 + 1) < 3)) & (((0 + 1) + 1) < 3)) & (!((((0 + 1) + 1) + 1) < 3)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    AllocationsList:
      A[0]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=1, store=0, select=1)
        op[0]: select idx=addr_0, res=IntSegment_sel0
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=3, store=3, select=0)
        op[0]: store idx=(addr_2 + 0), val=addr_1
        op[1]: store idx=(addr_2 + (0 + 1)), val=addr_1
        op[2]: store idx=(addr_2 + ((0 + 1) + 1)), val=addr_1
    StructSegments: <empty>
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_2)
    C[1]: (addr_1 != addr_2)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0)
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment.store((addr_2 + 0), addr_1).store((addr_2 + (0 + 1)), addr_1).store((addr_2 + ((0 + 1) + 1)), addr_1)
}

SAT
Model:
addr_0 -> #xffffffffffffffff
addr_1 -> #xffffffffffffffff
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000002)
len_parameter_value_____int -> #x0000000000000000
addr_2 -> #x0000000000000000
IntSegment_sel0 -> #x0000000000000002

---- State #5 ----
Interpreter{
  PC: (((((((true & (!(IntSegment_sel0 < 2))) & ((-1 + 1) < len_parameter_value_____int)) & (((-1 + 1) + 1) < len_parameter_value_____int)) & ((((-1 + 1) + 1) + 1) < len_parameter_value_____int)) & (((((-1 + 1) + 1) + 1) + 1) < len_parameter_value_____int)) & (!((((((-1 + 1) + 1) + 1) + 1) + 1) < len_parameter_value_____int))) & (!(0 < 3)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    AllocationsList:
      A[0]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=9, store=4, select=5)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=(addr_1 + (-1 + 1)), res=IntSegment_sel1
        op[2]: store idx=(addr_1 + (-1 + 1)), val=(IntSegment_sel1 + (-1 + 1))
        op[3]: select idx=(addr_1 + ((-1 + 1) + 1)), res=IntSegment_sel3
        op[4]: store idx=(addr_1 + ((-1 + 1) + 1)), val=(IntSegment_sel3 + ((-1 + 1) + 1))
        op[5]: select idx=(addr_1 + (((-1 + 1) + 1) + 1)), res=IntSegment_sel5
        op[6]: store idx=(addr_1 + (((-1 + 1) + 1) + 1)), val=(IntSegment_sel5 + (((-1 + 1) + 1) + 1))
        op[7]: select idx=(addr_1 + ((((-1 + 1) + 1) + 1) + 1)), res=IntSegment_sel7
        op[8]: store idx=(addr_1 + ((((-1 + 1) + 1) + 1) + 1)), val=(IntSegment_sel7 + ((((-1 + 1) + 1) + 1) + 1))
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_2)
    C[1]: (addr_1 != addr_2)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select((addr_1 + (-1 + 1)), IntSegment_sel1).store((addr_1 + (-1 + 1)), (IntSegment_sel1 + (-1 + 1))).select((addr_1 + ((-1 + 1) + 1)), IntSegment_sel3).store((addr_1 + ((-1 + 1) + 1)), (IntSegment_sel3 + ((-1 + 1) + 1))).select((addr_1 + (((-1 + 1) + 1) + 1)), IntSegment_sel5).store((addr_1 + (((-1 + 1) + 1) + 1)), (IntSegment_sel5 + (((-1 + 1) + 1) + 1))).select((addr_1 + ((((-1 + 1) + 1) + 1) + 1)), IntSegment_sel7).store((addr_1 + ((((-1 + 1) + 1) + 1) + 1)), (IntSegment_sel7 + ((((-1 + 1) + 1) + 1) + 1)))
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment
}

UNSAT
---- State #6 ----
Interpreter{
  PC: ((((((((true & (!(IntSegment_sel0 < 2))) & ((-1 + 1) < len_parameter_value_____int)) & (((-1 + 1) + 1) < len_parameter_value_____int)) & ((((-1 + 1) + 1) + 1) < len_parameter_value_____int)) & (((((-1 + 1) + 1) + 1) + 1) < len_parameter_value_____int)) & (!((((((-1 + 1) + 1) + 1) + 1) + 1) < len_parameter_value_____int))) & (0 < 3)) & (!((0 + 1) < 3)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    AllocationsList:
      A[0]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=9, store=4, select=5)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=(addr_1 + (-1 + 1)), res=IntSegment_sel1
        op[2]: store idx=(addr_1 + (-1 + 1)), val=(IntSegment_sel1 + (-1 + 1))
        op[3]: select idx=(addr_1 + ((-1 + 1) + 1)), res=IntSegment_sel3
        op[4]: store idx=(addr_1 + ((-1 + 1) + 1)), val=(IntSegment_sel3 + ((-1 + 1) + 1))
        op[5]: select idx=(addr_1 + (((-1 + 1) + 1) + 1)), res=IntSegment_sel5
        op[6]: store idx=(addr_1 + (((-1 + 1) + 1) + 1)), val=(IntSegment_sel5 + (((-1 + 1) + 1) + 1))
        op[7]: select idx=(addr_1 + ((((-1 + 1) + 1) + 1) + 1)), res=IntSegment_sel7
        op[8]: store idx=(addr_1 + ((((-1 + 1) + 1) + 1) + 1)), val=(IntSegment_sel7 + ((((-1 + 1) + 1) + 1) + 1))
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=1, store=1, select=0)
        op[0]: store idx=(addr_2 + 0), val=addr_1
    StructSegments: <empty>
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_2)
    C[1]: (addr_1 != addr_2)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select((addr_1 + (-1 + 1)), IntSegment_sel1).store((addr_1 + (-1 + 1)), (IntSegment_sel1 + (-1 + 1))).select((addr_1 + ((-1 + 1) + 1)), IntSegment_sel3).store((addr_1 + ((-1 + 1) + 1)), (IntSegment_sel3 + ((-1 + 1) + 1))).select((addr_1 + (((-1 + 1) + 1) + 1)), IntSegment_sel5).store((addr_1 + (((-1 + 1) + 1) + 1)), (IntSegment_sel5 + (((-1 + 1) + 1) + 1))).select((addr_1 + ((((-1 + 1) + 1) + 1) + 1)), IntSegment_sel7).store((addr_1 + ((((-1 + 1) + 1) + 1) + 1)), (IntSegment_sel7 + ((((-1 + 1) + 1) + 1) + 1)))
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment.store((addr_2 + 0), addr_1)
}

UNSAT
---- State #7 ----
Interpreter{
  PC: (((((((((((true & (!(IntSegment_sel0 < 2))) & ((-1 + 1) < len_parameter_value_____int)) & (((-1 + 1) + 1) < len_parameter_value_____int)) & ((((-1 + 1) + 1) + 1) < len_parameter_value_____int)) & (((((-1 + 1) + 1) + 1) + 1) < len_parameter_value_____int)) & (!((((((-1 + 1) + 1) + 1) + 1) + 1) < len_parameter_value_____int))) & (0 < 3)) & ((0 + 1) < 3)) & (((0 + 1) + 1) < 3)) & ((((0 + 1) + 1) + 1) < 3)) & (!(((((0 + 1) + 1) + 1) + 1) < 3)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    AllocationsList:
      A[0]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=9, store=4, select=5)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=(addr_1 + (-1 + 1)), res=IntSegment_sel1
        op[2]: store idx=(addr_1 + (-1 + 1)), val=(IntSegment_sel1 + (-1 + 1))
        op[3]: select idx=(addr_1 + ((-1 + 1) + 1)), res=IntSegment_sel3
        op[4]: store idx=(addr_1 + ((-1 + 1) + 1)), val=(IntSegment_sel3 + ((-1 + 1) + 1))
        op[5]: select idx=(addr_1 + (((-1 + 1) + 1) + 1)), res=IntSegment_sel5
        op[6]: store idx=(addr_1 + (((-1 + 1) + 1) + 1)), val=(IntSegment_sel5 + (((-1 + 1) + 1) + 1))
        op[7]: select idx=(addr_1 + ((((-1 + 1) + 1) + 1) + 1)), res=IntSegment_sel7
        op[8]: store idx=(addr_1 + ((((-1 + 1) + 1) + 1) + 1)), val=(IntSegment_sel7 + ((((-1 + 1) + 1) + 1) + 1))
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=4, store=4, select=0)
        op[0]: store idx=(addr_2 + 0), val=addr_1
        op[1]: store idx=(addr_2 + (0 + 1)), val=addr_1
        op[2]: store idx=(addr_2 + ((0 + 1) + 1)), val=addr_1
        op[3]: store idx=(addr_2 + (((0 + 1) + 1) + 1)), val=addr_1
    StructSegments: <empty>
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_2)
    C[1]: (addr_1 != addr_2)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select((addr_1 + (-1 + 1)), IntSegment_sel1).store((addr_1 + (-1 + 1)), (IntSegment_sel1 + (-1 + 1))).select((addr_1 + ((-1 + 1) + 1)), IntSegment_sel3).store((addr_1 + ((-1 + 1) + 1)), (IntSegment_sel3 + ((-1 + 1) + 1))).select((addr_1 + (((-1 + 1) + 1) + 1)), IntSegment_sel5).store((addr_1 + (((-1 + 1) + 1) + 1)), (IntSegment_sel5 + (((-1 + 1) + 1) + 1))).select((addr_1 + ((((-1 + 1) + 1) + 1) + 1)), IntSegment_sel7).store((addr_1 + ((((-1 + 1) + 1) + 1) + 1)), (IntSegment_sel7 + ((((-1 + 1) + 1) + 1) + 1)))
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment.store((addr_2 + 0), addr_1).store((addr_2 + (0 + 1)), addr_1).store((addr_2 + ((0 + 1) + 1)), addr_1).store((addr_2 + (((0 + 1) + 1) + 1)), addr_1)
}

UNSAT
---- State #8 ----
Interpreter{
  PC: (((((((((true & (!(IntSegment_sel0 < 2))) & ((-1 + 1) < len_parameter_value_____int)) & (((-1 + 1) + 1) < len_parameter_value_____int)) & ((((-1 + 1) + 1) + 1) < len_parameter_value_____int)) & (((((-1 + 1) + 1) + 1) + 1) < len_parameter_value_____int)) & (!((((((-1 + 1) + 1) + 1) + 1) + 1) < len_parameter_value_____int))) & (0 < 3)) & ((0 + 1) < 3)) & (!(((0 + 1) + 1) < 3)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    AllocationsList:
      A[0]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=9, store=4, select=5)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=(addr_1 + (-1 + 1)), res=IntSegment_sel1
        op[2]: store idx=(addr_1 + (-1 + 1)), val=(IntSegment_sel1 + (-1 + 1))
        op[3]: select idx=(addr_1 + ((-1 + 1) + 1)), res=IntSegment_sel3
        op[4]: store idx=(addr_1 + ((-1 + 1) + 1)), val=(IntSegment_sel3 + ((-1 + 1) + 1))
        op[5]: select idx=(addr_1 + (((-1 + 1) + 1) + 1)), res=IntSegment_sel5
        op[6]: store idx=(addr_1 + (((-1 + 1) + 1) + 1)), val=(IntSegment_sel5 + (((-1 + 1) + 1) + 1))
        op[7]: select idx=(addr_1 + ((((-1 + 1) + 1) + 1) + 1)), res=IntSegment_sel7
        op[8]: store idx=(addr_1 + ((((-1 + 1) + 1) + 1) + 1)), val=(IntSegment_sel7 + ((((-1 + 1) + 1) + 1) + 1))
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=2, store=2, select=0)
        op[0]: store idx=(addr_2 + 0), val=addr_1
        op[1]: store idx=(addr_2 + (0 + 1)), val=addr_1
    StructSegments: <empty>
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_2)
    C[1]: (addr_1 != addr_2)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select((addr_1 + (-1 + 1)), IntSegment_sel1).store((addr_1 + (-1 + 1)), (IntSegment_sel1 + (-1 + 1))).select((addr_1 + ((-1 + 1) + 1)), IntSegment_sel3).store((addr_1 + ((-1 + 1) + 1)), (IntSegment_sel3 + ((-1 + 1) + 1))).select((addr_1 + (((-1 + 1) + 1) + 1)), IntSegment_sel5).store((addr_1 + (((-1 + 1) + 1) + 1)), (IntSegment_sel5 + (((-1 + 1) + 1) + 1))).select((addr_1 + ((((-1 + 1) + 1) + 1) + 1)), IntSegment_sel7).store((addr_1 + ((((-1 + 1) + 1) + 1) + 1)), (IntSegment_sel7 + ((((-1 + 1) + 1) + 1) + 1)))
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment.store((addr_2 + 0), addr_1).store((addr_2 + (0 + 1)), addr_1)
}

UNSAT
---- State #9 ----
Interpreter{
  PC: ((((((((((true & (!(IntSegment_sel0 < 2))) & ((-1 + 1) < len_parameter_value_____int)) & (((-1 + 1) + 1) < len_parameter_value_____int)) & ((((-1 + 1) + 1) + 1) < len_parameter_value_____int)) & (((((-1 + 1) + 1) + 1) + 1) < len_parameter_value_____int)) & (!((((((-1 + 1) + 1) + 1) + 1) + 1) < len_parameter_value_____int))) & (0 < 3)) & ((0 + 1) < 3)) & (((0 + 1) + 1) < 3)) & (!((((0 + 1) + 1) + 1) < 3)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    AllocationsList:
      A[0]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=9, store=4, select=5)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=(addr_1 + (-1 + 1)), res=IntSegment_sel1
        op[2]: store idx=(addr_1 + (-1 + 1)), val=(IntSegment_sel1 + (-1 + 1))
        op[3]: select idx=(addr_1 + ((-1 + 1) + 1)), res=IntSegment_sel3
        op[4]: store idx=(addr_1 + ((-1 + 1) + 1)), val=(IntSegment_sel3 + ((-1 + 1) + 1))
        op[5]: select idx=(addr_1 + (((-1 + 1) + 1) + 1)), res=IntSegment_sel5
        op[6]: store idx=(addr_1 + (((-1 + 1) + 1) + 1)), val=(IntSegment_sel5 + (((-1 + 1) + 1) + 1))
        op[7]: select idx=(addr_1 + ((((-1 + 1) + 1) + 1) + 1)), res=IntSegment_sel7
        op[8]: store idx=(addr_1 + ((((-1 + 1) + 1) + 1) + 1)), val=(IntSegment_sel7 + ((((-1 + 1) + 1) + 1) + 1))
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=3, store=3, select=0)
        op[0]: store idx=(addr_2 + 0), val=addr_1
        op[1]: store idx=(addr_2 + (0 + 1)), val=addr_1
        op[2]: store idx=(addr_2 + ((0 + 1) + 1)), val=addr_1
    StructSegments: <empty>
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_2)
    C[1]: (addr_1 != addr_2)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select((addr_1 + (-1 + 1)), IntSegment_sel1).store((addr_1 + (-1 + 1)), (IntSegment_sel1 + (-1 + 1))).select((addr_1 + ((-1 + 1) + 1)), IntSegment_sel3).store((addr_1 + ((-1 + 1) + 1)), (IntSegment_sel3 + ((-1 + 1) + 1))).select((addr_1 + (((-1 + 1) + 1) + 1)), IntSegment_sel5).store((addr_1 + (((-1 + 1) + 1) + 1)), (IntSegment_sel5 + (((-1 + 1) + 1) + 1))).select((addr_1 + ((((-1 + 1) + 1) + 1) + 1)), IntSegment_sel7).store((addr_1 + ((((-1 + 1) + 1) + 1) + 1)), (IntSegment_sel7 + ((((-1 + 1) + 1) + 1) + 1)))
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment.store((addr_2 + 0), addr_1).store((addr_2 + (0 + 1)), addr_1).store((addr_2 + ((0 + 1) + 1)), addr_1)
}

SAT
Model:
addr_0 -> #xffffffffffffffff
addr_1 -> #xffffffffffffffff
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000002)
len_parameter_value_____int -> #x0000000000000004
addr_2 -> #x0000000000000000
IntSegment_sel7 -> #x0000000000000002
IntSegment_sel5 -> #x0000000000000002
IntSegment_sel1 -> #x0000000000000002
IntSegment_sel3 -> #x0000000000000002
IntSegment_sel0 -> #x0000000000000002

---- State #10 ----
Interpreter{
  PC: (((((((true & (!(IntSegment_sel0 < 2))) & ((-1 + 1) < len_parameter_value_____int)) & (((-1 + 1) + 1) < len_parameter_value_____int)) & ((((-1 + 1) + 1) + 1) < len_parameter_value_____int)) & (!(((((-1 + 1) + 1) + 1) + 1) < len_parameter_value_____int))) & (0 < 3)) & (!((0 + 1) < 3)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    AllocationsList:
      A[0]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=7, store=3, select=4)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=(addr_1 + (-1 + 1)), res=IntSegment_sel1
        op[2]: store idx=(addr_1 + (-1 + 1)), val=(IntSegment_sel1 + (-1 + 1))
        op[3]: select idx=(addr_1 + ((-1 + 1) + 1)), res=IntSegment_sel3
        op[4]: store idx=(addr_1 + ((-1 + 1) + 1)), val=(IntSegment_sel3 + ((-1 + 1) + 1))
        op[5]: select idx=(addr_1 + (((-1 + 1) + 1) + 1)), res=IntSegment_sel5
        op[6]: store idx=(addr_1 + (((-1 + 1) + 1) + 1)), val=(IntSegment_sel5 + (((-1 + 1) + 1) + 1))
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=1, store=1, select=0)
        op[0]: store idx=(addr_2 + 0), val=addr_1
    StructSegments: <empty>
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_2)
    C[1]: (addr_1 != addr_2)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select((addr_1 + (-1 + 1)), IntSegment_sel1).store((addr_1 + (-1 + 1)), (IntSegment_sel1 + (-1 + 1))).select((addr_1 + ((-1 + 1) + 1)), IntSegment_sel3).store((addr_1 + ((-1 + 1) + 1)), (IntSegment_sel3 + ((-1 + 1) + 1))).select((addr_1 + (((-1 + 1) + 1) + 1)), IntSegment_sel5).store((addr_1 + (((-1 + 1) + 1) + 1)), (IntSegment_sel5 + (((-1 + 1) + 1) + 1)))
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment.store((addr_2 + 0), addr_1)
}

UNSAT
---- State #11 ----
Interpreter{
  PC: ((((((true & (!(IntSegment_sel0 < 2))) & ((-1 + 1) < len_parameter_value_____int)) & (((-1 + 1) + 1) < len_parameter_value_____int)) & ((((-1 + 1) + 1) + 1) < len_parameter_value_____int)) & (!(((((-1 + 1) + 1) + 1) + 1) < len_parameter_value_____int))) & (!(0 < 3)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    AllocationsList:
      A[0]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=7, store=3, select=4)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=(addr_1 + (-1 + 1)), res=IntSegment_sel1
        op[2]: store idx=(addr_1 + (-1 + 1)), val=(IntSegment_sel1 + (-1 + 1))
        op[3]: select idx=(addr_1 + ((-1 + 1) + 1)), res=IntSegment_sel3
        op[4]: store idx=(addr_1 + ((-1 + 1) + 1)), val=(IntSegment_sel3 + ((-1 + 1) + 1))
        op[5]: select idx=(addr_1 + (((-1 + 1) + 1) + 1)), res=IntSegment_sel5
        op[6]: store idx=(addr_1 + (((-1 + 1) + 1) + 1)), val=(IntSegment_sel5 + (((-1 + 1) + 1) + 1))
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_2)
    C[1]: (addr_1 != addr_2)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select((addr_1 + (-1 + 1)), IntSegment_sel1).store((addr_1 + (-1 + 1)), (IntSegment_sel1 + (-1 + 1))).select((addr_1 + ((-1 + 1) + 1)), IntSegment_sel3).store((addr_1 + ((-1 + 1) + 1)), (IntSegment_sel3 + ((-1 + 1) + 1))).select((addr_1 + (((-1 + 1) + 1) + 1)), IntSegment_sel5).store((addr_1 + (((-1 + 1) + 1) + 1)), (IntSegment_sel5 + (((-1 + 1) + 1) + 1)))
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment
}

UNSAT
---- State #12 ----
Interpreter{
  PC: ((((((((true & (!(IntSegment_sel0 < 2))) & ((-1 + 1) < len_parameter_value_____int)) & (((-1 + 1) + 1) < len_parameter_value_____int)) & ((((-1 + 1) + 1) + 1) < len_parameter_value_____int)) & (!(((((-1 + 1) + 1) + 1) + 1) < len_parameter_value_____int))) & (0 < 3)) & ((0 + 1) < 3)) & (!(((0 + 1) + 1) < 3)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    AllocationsList:
      A[0]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=7, store=3, select=4)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=(addr_1 + (-1 + 1)), res=IntSegment_sel1
        op[2]: store idx=(addr_1 + (-1 + 1)), val=(IntSegment_sel1 + (-1 + 1))
        op[3]: select idx=(addr_1 + ((-1 + 1) + 1)), res=IntSegment_sel3
        op[4]: store idx=(addr_1 + ((-1 + 1) + 1)), val=(IntSegment_sel3 + ((-1 + 1) + 1))
        op[5]: select idx=(addr_1 + (((-1 + 1) + 1) + 1)), res=IntSegment_sel5
        op[6]: store idx=(addr_1 + (((-1 + 1) + 1) + 1)), val=(IntSegment_sel5 + (((-1 + 1) + 1) + 1))
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=2, store=2, select=0)
        op[0]: store idx=(addr_2 + 0), val=addr_1
        op[1]: store idx=(addr_2 + (0 + 1)), val=addr_1
    StructSegments: <empty>
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_2)
    C[1]: (addr_1 != addr_2)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select((addr_1 + (-1 + 1)), IntSegment_sel1).store((addr_1 + (-1 + 1)), (IntSegment_sel1 + (-1 + 1))).select((addr_1 + ((-1 + 1) + 1)), IntSegment_sel3).store((addr_1 + ((-1 + 1) + 1)), (IntSegment_sel3 + ((-1 + 1) + 1))).select((addr_1 + (((-1 + 1) + 1) + 1)), IntSegment_sel5).store((addr_1 + (((-1 + 1) + 1) + 1)), (IntSegment_sel5 + (((-1 + 1) + 1) + 1)))
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment.store((addr_2 + 0), addr_1).store((addr_2 + (0 + 1)), addr_1)
}

UNSAT
---- State #13 ----
Interpreter{
  PC: (((((((((true & (!(IntSegment_sel0 < 2))) & ((-1 + 1) < len_parameter_value_____int)) & (((-1 + 1) + 1) < len_parameter_value_____int)) & ((((-1 + 1) + 1) + 1) < len_parameter_value_____int)) & (!(((((-1 + 1) + 1) + 1) + 1) < len_parameter_value_____int))) & (0 < 3)) & ((0 + 1) < 3)) & (((0 + 1) + 1) < 3)) & (!((((0 + 1) + 1) + 1) < 3)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    AllocationsList:
      A[0]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=7, store=3, select=4)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=(addr_1 + (-1 + 1)), res=IntSegment_sel1
        op[2]: store idx=(addr_1 + (-1 + 1)), val=(IntSegment_sel1 + (-1 + 1))
        op[3]: select idx=(addr_1 + ((-1 + 1) + 1)), res=IntSegment_sel3
        op[4]: store idx=(addr_1 + ((-1 + 1) + 1)), val=(IntSegment_sel3 + ((-1 + 1) + 1))
        op[5]: select idx=(addr_1 + (((-1 + 1) + 1) + 1)), res=IntSegment_sel5
        op[6]: store idx=(addr_1 + (((-1 + 1) + 1) + 1)), val=(IntSegment_sel5 + (((-1 + 1) + 1) + 1))
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=3, store=3, select=0)
        op[0]: store idx=(addr_2 + 0), val=addr_1
        op[1]: store idx=(addr_2 + (0 + 1)), val=addr_1
        op[2]: store idx=(addr_2 + ((0 + 1) + 1)), val=addr_1
    StructSegments: <empty>
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_2)
    C[1]: (addr_1 != addr_2)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select((addr_1 + (-1 + 1)), IntSegment_sel1).store((addr_1 + (-1 + 1)), (IntSegment_sel1 + (-1 + 1))).select((addr_1 + ((-1 + 1) + 1)), IntSegment_sel3).store((addr_1 + ((-1 + 1) + 1)), (IntSegment_sel3 + ((-1 + 1) + 1))).select((addr_1 + (((-1 + 1) + 1) + 1)), IntSegment_sel5).store((addr_1 + (((-1 + 1) + 1) + 1)), (IntSegment_sel5 + (((-1 + 1) + 1) + 1)))
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment.store((addr_2 + 0), addr_1).store((addr_2 + (0 + 1)), addr_1).store((addr_2 + ((0 + 1) + 1)), addr_1)
}

SAT
Model:
addr_0 -> #xffffffffffffffff
addr_1 -> #xffffffffffffffff
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000002)
len_parameter_value_____int -> #x0000000000000003
addr_2 -> #x0000000000000000
IntSegment_sel5 -> #x0000000000000002
IntSegment_sel1 -> #x0000000000000002
IntSegment_sel3 -> #x0000000000000002
IntSegment_sel0 -> #x0000000000000002

---- State #14 ----
Interpreter{
  PC: ((((((((((true & (!(IntSegment_sel0 < 2))) & ((-1 + 1) < len_parameter_value_____int)) & (((-1 + 1) + 1) < len_parameter_value_____int)) & ((((-1 + 1) + 1) + 1) < len_parameter_value_____int)) & (!(((((-1 + 1) + 1) + 1) + 1) < len_parameter_value_____int))) & (0 < 3)) & ((0 + 1) < 3)) & (((0 + 1) + 1) < 3)) & ((((0 + 1) + 1) + 1) < 3)) & (!(((((0 + 1) + 1) + 1) + 1) < 3)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    AllocationsList:
      A[0]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=7, store=3, select=4)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=(addr_1 + (-1 + 1)), res=IntSegment_sel1
        op[2]: store idx=(addr_1 + (-1 + 1)), val=(IntSegment_sel1 + (-1 + 1))
        op[3]: select idx=(addr_1 + ((-1 + 1) + 1)), res=IntSegment_sel3
        op[4]: store idx=(addr_1 + ((-1 + 1) + 1)), val=(IntSegment_sel3 + ((-1 + 1) + 1))
        op[5]: select idx=(addr_1 + (((-1 + 1) + 1) + 1)), res=IntSegment_sel5
        op[6]: store idx=(addr_1 + (((-1 + 1) + 1) + 1)), val=(IntSegment_sel5 + (((-1 + 1) + 1) + 1))
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=4, store=4, select=0)
        op[0]: store idx=(addr_2 + 0), val=addr_1
        op[1]: store idx=(addr_2 + (0 + 1)), val=addr_1
        op[2]: store idx=(addr_2 + ((0 + 1) + 1)), val=addr_1
        op[3]: store idx=(addr_2 + (((0 + 1) + 1) + 1)), val=addr_1
    StructSegments: <empty>
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_2)
    C[1]: (addr_1 != addr_2)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select((addr_1 + (-1 + 1)), IntSegment_sel1).store((addr_1 + (-1 + 1)), (IntSegment_sel1 + (-1 + 1))).select((addr_1 + ((-1 + 1) + 1)), IntSegment_sel3).store((addr_1 + ((-1 + 1) + 1)), (IntSegment_sel3 + ((-1 + 1) + 1))).select((addr_1 + (((-1 + 1) + 1) + 1)), IntSegment_sel5).store((addr_1 + (((-1 + 1) + 1) + 1)), (IntSegment_sel5 + (((-1 + 1) + 1) + 1)))
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment.store((addr_2 + 0), addr_1).store((addr_2 + (0 + 1)), addr_1).store((addr_2 + ((0 + 1) + 1)), addr_1).store((addr_2 + (((0 + 1) + 1) + 1)), addr_1)
}

UNSAT
---- State #15 ----
Interpreter{
  PC: (((((true & (!(IntSegment_sel0 < 2))) & ((-1 + 1) < len_parameter_value_____int)) & (((-1 + 1) + 1) < len_parameter_value_____int)) & (!((((-1 + 1) + 1) + 1) < len_parameter_value_____int))) & (!(0 < 3)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    AllocationsList:
      A[0]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=5, store=2, select=3)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=(addr_1 + (-1 + 1)), res=IntSegment_sel1
        op[2]: store idx=(addr_1 + (-1 + 1)), val=(IntSegment_sel1 + (-1 + 1))
        op[3]: select idx=(addr_1 + ((-1 + 1) + 1)), res=IntSegment_sel3
        op[4]: store idx=(addr_1 + ((-1 + 1) + 1)), val=(IntSegment_sel3 + ((-1 + 1) + 1))
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_2)
    C[1]: (addr_1 != addr_2)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select((addr_1 + (-1 + 1)), IntSegment_sel1).store((addr_1 + (-1 + 1)), (IntSegment_sel1 + (-1 + 1))).select((addr_1 + ((-1 + 1) + 1)), IntSegment_sel3).store((addr_1 + ((-1 + 1) + 1)), (IntSegment_sel3 + ((-1 + 1) + 1)))
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment
}

UNSAT
---- State #16 ----
Interpreter{
  PC: ((((true & (!(IntSegment_sel0 < 2))) & (!((-1 + 1) < len_parameter_value_____int))) & (0 < 3)) & (!((0 + 1) < 3)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    AllocationsList:
      A[0]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=1, store=0, select=1)
        op[0]: select idx=addr_0, res=IntSegment_sel0
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=1, store=1, select=0)
        op[0]: store idx=(addr_2 + 0), val=addr_1
    StructSegments: <empty>
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_2)
    C[1]: (addr_1 != addr_2)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0)
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment.store((addr_2 + 0), addr_1)
}

UNSAT
---- State #17 ----
Interpreter{
  PC: (((((((true & (!(IntSegment_sel0 < 2))) & (!((-1 + 1) < len_parameter_value_____int))) & (0 < 3)) & ((0 + 1) < 3)) & (((0 + 1) + 1) < 3)) & ((((0 + 1) + 1) + 1) < 3)) & (!(((((0 + 1) + 1) + 1) + 1) < 3)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    AllocationsList:
      A[0]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=1, store=0, select=1)
        op[0]: select idx=addr_0, res=IntSegment_sel0
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=4, store=4, select=0)
        op[0]: store idx=(addr_2 + 0), val=addr_1
        op[1]: store idx=(addr_2 + (0 + 1)), val=addr_1
        op[2]: store idx=(addr_2 + ((0 + 1) + 1)), val=addr_1
        op[3]: store idx=(addr_2 + (((0 + 1) + 1) + 1)), val=addr_1
    StructSegments: <empty>
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_2)
    C[1]: (addr_1 != addr_2)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0)
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment.store((addr_2 + 0), addr_1).store((addr_2 + (0 + 1)), addr_1).store((addr_2 + ((0 + 1) + 1)), addr_1).store((addr_2 + (((0 + 1) + 1) + 1)), addr_1)
}

UNSAT
---- State #18 ----
Interpreter{
  PC: ((((true & (!(IntSegment_sel0 < 2))) & ((-1 + 1) < len_parameter_value_____int)) & (!(((-1 + 1) + 1) < len_parameter_value_____int))) & (!(0 < 3)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    AllocationsList:
      A[0]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=3, store=1, select=2)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=(addr_1 + (-1 + 1)), res=IntSegment_sel1
        op[2]: store idx=(addr_1 + (-1 + 1)), val=(IntSegment_sel1 + (-1 + 1))
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_2)
    C[1]: (addr_1 != addr_2)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select((addr_1 + (-1 + 1)), IntSegment_sel1).store((addr_1 + (-1 + 1)), (IntSegment_sel1 + (-1 + 1)))
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment
}

UNSAT
---- State #19 ----
Interpreter{
  PC: ((((((true & (!(IntSegment_sel0 < 2))) & ((-1 + 1) < len_parameter_value_____int)) & (((-1 + 1) + 1) < len_parameter_value_____int)) & (!((((-1 + 1) + 1) + 1) < len_parameter_value_____int))) & (0 < 3)) & (!((0 + 1) < 3)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    AllocationsList:
      A[0]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=5, store=2, select=3)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=(addr_1 + (-1 + 1)), res=IntSegment_sel1
        op[2]: store idx=(addr_1 + (-1 + 1)), val=(IntSegment_sel1 + (-1 + 1))
        op[3]: select idx=(addr_1 + ((-1 + 1) + 1)), res=IntSegment_sel3
        op[4]: store idx=(addr_1 + ((-1 + 1) + 1)), val=(IntSegment_sel3 + ((-1 + 1) + 1))
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=1, store=1, select=0)
        op[0]: store idx=(addr_2 + 0), val=addr_1
    StructSegments: <empty>
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_2)
    C[1]: (addr_1 != addr_2)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select((addr_1 + (-1 + 1)), IntSegment_sel1).store((addr_1 + (-1 + 1)), (IntSegment_sel1 + (-1 + 1))).select((addr_1 + ((-1 + 1) + 1)), IntSegment_sel3).store((addr_1 + ((-1 + 1) + 1)), (IntSegment_sel3 + ((-1 + 1) + 1)))
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment.store((addr_2 + 0), addr_1)
}

UNSAT
---- State #20 ----
Interpreter{
  PC: (((((((true & (!(IntSegment_sel0 < 2))) & ((-1 + 1) < len_parameter_value_____int)) & (((-1 + 1) + 1) < len_parameter_value_____int)) & (!((((-1 + 1) + 1) + 1) < len_parameter_value_____int))) & (0 < 3)) & ((0 + 1) < 3)) & (!(((0 + 1) + 1) < 3)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    AllocationsList:
      A[0]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=5, store=2, select=3)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=(addr_1 + (-1 + 1)), res=IntSegment_sel1
        op[2]: store idx=(addr_1 + (-1 + 1)), val=(IntSegment_sel1 + (-1 + 1))
        op[3]: select idx=(addr_1 + ((-1 + 1) + 1)), res=IntSegment_sel3
        op[4]: store idx=(addr_1 + ((-1 + 1) + 1)), val=(IntSegment_sel3 + ((-1 + 1) + 1))
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=2, store=2, select=0)
        op[0]: store idx=(addr_2 + 0), val=addr_1
        op[1]: store idx=(addr_2 + (0 + 1)), val=addr_1
    StructSegments: <empty>
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_2)
    C[1]: (addr_1 != addr_2)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select((addr_1 + (-1 + 1)), IntSegment_sel1).store((addr_1 + (-1 + 1)), (IntSegment_sel1 + (-1 + 1))).select((addr_1 + ((-1 + 1) + 1)), IntSegment_sel3).store((addr_1 + ((-1 + 1) + 1)), (IntSegment_sel3 + ((-1 + 1) + 1)))
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment.store((addr_2 + 0), addr_1).store((addr_2 + (0 + 1)), addr_1)
}

UNSAT
---- State #21 ----
Interpreter{
  PC: ((((((((true & (!(IntSegment_sel0 < 2))) & ((-1 + 1) < len_parameter_value_____int)) & (((-1 + 1) + 1) < len_parameter_value_____int)) & (!((((-1 + 1) + 1) + 1) < len_parameter_value_____int))) & (0 < 3)) & ((0 + 1) < 3)) & (((0 + 1) + 1) < 3)) & (!((((0 + 1) + 1) + 1) < 3)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    AllocationsList:
      A[0]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=5, store=2, select=3)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=(addr_1 + (-1 + 1)), res=IntSegment_sel1
        op[2]: store idx=(addr_1 + (-1 + 1)), val=(IntSegment_sel1 + (-1 + 1))
        op[3]: select idx=(addr_1 + ((-1 + 1) + 1)), res=IntSegment_sel3
        op[4]: store idx=(addr_1 + ((-1 + 1) + 1)), val=(IntSegment_sel3 + ((-1 + 1) + 1))
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=3, store=3, select=0)
        op[0]: store idx=(addr_2 + 0), val=addr_1
        op[1]: store idx=(addr_2 + (0 + 1)), val=addr_1
        op[2]: store idx=(addr_2 + ((0 + 1) + 1)), val=addr_1
    StructSegments: <empty>
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_2)
    C[1]: (addr_1 != addr_2)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select((addr_1 + (-1 + 1)), IntSegment_sel1).store((addr_1 + (-1 + 1)), (IntSegment_sel1 + (-1 + 1))).select((addr_1 + ((-1 + 1) + 1)), IntSegment_sel3).store((addr_1 + ((-1 + 1) + 1)), (IntSegment_sel3 + ((-1 + 1) + 1)))
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment.store((addr_2 + 0), addr_1).store((addr_2 + (0 + 1)), addr_1).store((addr_2 + ((0 + 1) + 1)), addr_1)
}

SAT
Model:
addr_0 -> #xffffffffffffffff
addr_1 -> #xffffffffffffffff
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000002)
len_parameter_value_____int -> #x0000000000000002
addr_2 -> #x0000000000000000
IntSegment_sel1 -> #x0000000000000002
IntSegment_sel3 -> #x0000000000000002
IntSegment_sel0 -> #x0000000000000002

---- State #22 ----
Interpreter{
  PC: (((((((((true & (!(IntSegment_sel0 < 2))) & ((-1 + 1) < len_parameter_value_____int)) & (((-1 + 1) + 1) < len_parameter_value_____int)) & (!((((-1 + 1) + 1) + 1) < len_parameter_value_____int))) & (0 < 3)) & ((0 + 1) < 3)) & (((0 + 1) + 1) < 3)) & ((((0 + 1) + 1) + 1) < 3)) & (!(((((0 + 1) + 1) + 1) + 1) < 3)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    AllocationsList:
      A[0]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=5, store=2, select=3)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=(addr_1 + (-1 + 1)), res=IntSegment_sel1
        op[2]: store idx=(addr_1 + (-1 + 1)), val=(IntSegment_sel1 + (-1 + 1))
        op[3]: select idx=(addr_1 + ((-1 + 1) + 1)), res=IntSegment_sel3
        op[4]: store idx=(addr_1 + ((-1 + 1) + 1)), val=(IntSegment_sel3 + ((-1 + 1) + 1))
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=4, store=4, select=0)
        op[0]: store idx=(addr_2 + 0), val=addr_1
        op[1]: store idx=(addr_2 + (0 + 1)), val=addr_1
        op[2]: store idx=(addr_2 + ((0 + 1) + 1)), val=addr_1
        op[3]: store idx=(addr_2 + (((0 + 1) + 1) + 1)), val=addr_1
    StructSegments: <empty>
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_2)
    C[1]: (addr_1 != addr_2)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select((addr_1 + (-1 + 1)), IntSegment_sel1).store((addr_1 + (-1 + 1)), (IntSegment_sel1 + (-1 + 1))).select((addr_1 + ((-1 + 1) + 1)), IntSegment_sel3).store((addr_1 + ((-1 + 1) + 1)), (IntSegment_sel3 + ((-1 + 1) + 1)))
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment.store((addr_2 + 0), addr_1).store((addr_2 + (0 + 1)), addr_1).store((addr_2 + ((0 + 1) + 1)), addr_1).store((addr_2 + (((0 + 1) + 1) + 1)), addr_1)
}

UNSAT
---- State #23 ----
Interpreter{
  PC: ((((((true & (!(IntSegment_sel0 < 2))) & ((-1 + 1) < len_parameter_value_____int)) & (!(((-1 + 1) + 1) < len_parameter_value_____int))) & (0 < 3)) & ((0 + 1) < 3)) & (!(((0 + 1) + 1) < 3)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    AllocationsList:
      A[0]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=3, store=1, select=2)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=(addr_1 + (-1 + 1)), res=IntSegment_sel1
        op[2]: store idx=(addr_1 + (-1 + 1)), val=(IntSegment_sel1 + (-1 + 1))
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=2, store=2, select=0)
        op[0]: store idx=(addr_2 + 0), val=addr_1
        op[1]: store idx=(addr_2 + (0 + 1)), val=addr_1
    StructSegments: <empty>
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_2)
    C[1]: (addr_1 != addr_2)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select((addr_1 + (-1 + 1)), IntSegment_sel1).store((addr_1 + (-1 + 1)), (IntSegment_sel1 + (-1 + 1)))
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment.store((addr_2 + 0), addr_1).store((addr_2 + (0 + 1)), addr_1)
}

UNSAT
---- State #24 ----
Interpreter{
  PC: (((((((true & (!(IntSegment_sel0 < 2))) & ((-1 + 1) < len_parameter_value_____int)) & (!(((-1 + 1) + 1) < len_parameter_value_____int))) & (0 < 3)) & ((0 + 1) < 3)) & (((0 + 1) + 1) < 3)) & (!((((0 + 1) + 1) + 1) < 3)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    AllocationsList:
      A[0]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=3, store=1, select=2)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=(addr_1 + (-1 + 1)), res=IntSegment_sel1
        op[2]: store idx=(addr_1 + (-1 + 1)), val=(IntSegment_sel1 + (-1 + 1))
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=3, store=3, select=0)
        op[0]: store idx=(addr_2 + 0), val=addr_1
        op[1]: store idx=(addr_2 + (0 + 1)), val=addr_1
        op[2]: store idx=(addr_2 + ((0 + 1) + 1)), val=addr_1
    StructSegments: <empty>
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_2)
    C[1]: (addr_1 != addr_2)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select((addr_1 + (-1 + 1)), IntSegment_sel1).store((addr_1 + (-1 + 1)), (IntSegment_sel1 + (-1 + 1)))
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment.store((addr_2 + 0), addr_1).store((addr_2 + (0 + 1)), addr_1).store((addr_2 + ((0 + 1) + 1)), addr_1)
}

SAT
Model:
addr_0 -> #xffffffffffffffff
addr_1 -> #xffffffffffffffff
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000002)
len_parameter_value_____int -> #x0000000000000001
addr_2 -> #x0000000000000000
IntSegment_sel1 -> #x0000000000000002
IntSegment_sel0 -> #x0000000000000002

---- State #25 ----
Interpreter{
  PC: ((((((((true & (!(IntSegment_sel0 < 2))) & ((-1 + 1) < len_parameter_value_____int)) & (!(((-1 + 1) + 1) < len_parameter_value_____int))) & (0 < 3)) & ((0 + 1) < 3)) & (((0 + 1) + 1) < 3)) & ((((0 + 1) + 1) + 1) < 3)) & (!(((((0 + 1) + 1) + 1) + 1) < 3)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=array, dim=0, base=int)
    AllocationsList:
      A[0]: &(addr_2) (refType=array, dim=1, base=int)
    IntSegment:    IntSegment (elem=int, ops=3, store=1, select=2)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=(addr_1 + (-1 + 1)), res=IntSegment_sel1
        op[2]: store idx=(addr_1 + (-1 + 1)), val=(IntSegment_sel1 + (-1 + 1))
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=4, store=4, select=0)
        op[0]: store idx=(addr_2 + 0), val=addr_1
        op[1]: store idx=(addr_2 + (0 + 1)), val=addr_1
        op[2]: store idx=(addr_2 + ((0 + 1) + 1)), val=addr_1
        op[3]: store idx=(addr_2 + (((0 + 1) + 1) + 1)), val=addr_1
    StructSegments: <empty>
  }
  HeapConstraints: 7
    C[0]: (addr_0 != addr_2)
    C[1]: (addr_1 != addr_2)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select((addr_1 + (-1 + 1)), IntSegment_sel1).store((addr_1 + (-1 + 1)), (IntSegment_sel1 + (-1 + 1)))
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment.store((addr_2 + 0), addr_1).store((addr_2 + (0 + 1)), addr_1).store((addr_2 + ((0 + 1) + 1)), addr_1).store((addr_2 + (((0 + 1) + 1) + 1)), addr_1)
}

UNSAT

==================== Complement ====================
CFG построен для функции с 1 блоками
interpretDynamically: ^x
interpretDynamically: t0 == 1:int
interpretDynamically: return t1
resolveExpression: t0 == 1:int
resolveExpression: ^x
resolveExpression: parameter x : int
resolveExpression: 1:int
States: 1
---- State #0 ----
Interpreter{
  PC: true
  ReturnResults:
    [0] ((^IntSegment_sel0) == 1)
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
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
addr_0 -> #x0000000000000000
IntSegment_sel0 -> #x0000000000000000


==================== Xor ====================
CFG построен для функции с 1 блоками
interpretDynamically: x ^ y
interpretDynamically: t0 == 0:int
interpretDynamically: return t1
resolveExpression: t0 == 0:int
resolveExpression: x ^ y
resolveExpression: parameter x : int
resolveExpression: parameter y : int
resolveExpression: 0:int
States: 1
---- State #0 ----
Interpreter{
  PC: true
  ReturnResults:
    [0] ((IntSegment_sel0 ^ IntSegment_sel1) == 0)
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
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
addr_1 -> #x0000000000000000
IntSegment_sel1 -> #x0000000000000000
addr_0 -> #x0000000000000000
IntSegment_sel0 -> #x0000000000000000


==================== Or ====================
CFG построен для функции с 1 блоками
interpretDynamically: val | 7:int
interpretDynamically: t0 == 15:int
interpretDynamically: return t1
resolveExpression: t0 == 15:int
resolveExpression: val | 7:int
resolveExpression: parameter val : int
resolveExpression: 7:int
resolveExpression: 15:int
States: 1
---- State #0 ----
Interpreter{
  PC: true
  ReturnResults:
    [0] ((IntSegment_sel0 | 7) == 15)
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
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
addr_0 -> #x0000000000000000
IntSegment_sel0 -> #x0000000000000000


==================== And ====================
CFG построен для функции с 1 блоками
interpretDynamically: value - 1:int
interpretDynamically: value & t0
interpretDynamically: t1 == 0:int
interpretDynamically: return t2
resolveExpression: t1 == 0:int
resolveExpression: value & t0
resolveExpression: parameter value : int
resolveExpression: value - 1:int
resolveExpression: parameter value : int
resolveExpression: 1:int
resolveExpression: 0:int
States: 1
---- State #0 ----
Interpreter{
  PC: true
  ReturnResults:
    [0] ((IntSegment_sel0 & (IntSegment_sel1 - 1)) == 0)
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
addr_0 -> #x0000000000000000
IntSegment_sel1 -> #x0000000000000000
IntSegment_sel0 -> #x0000000000000000


==================== BooleanNot ====================
CFG построен для функции с 8 блоками
interpretDynamically: if boolA goto 1 else 2
resolveExpression: parameter boolA : bool
interpretDynamically: phi [0: false:bool, 1: boolB] #&&
resolveExpression: false:bool
interpretDynamically: if boolA goto 3 else 4
resolveExpression: parameter boolA : bool
interpretDynamically: jump 4
interpretDynamically: phi [2: true:bool, 3: boolB] #||
resolveExpression: parameter boolB : bool
interpretDynamically: phi [2: true:bool, 3: boolB] #||
resolveExpression: true:bool
interpretDynamically: if t0 goto 7 else 6
resolveExpression: phi [0: false:bool, 1: boolB] #&&
interpretDynamically: if t1 goto 5 else 6
resolveExpression: phi [2: true:bool, 3: boolB] #||
interpretDynamically: return 200:int
resolveExpression: 200:int
interpretDynamically: return 100:int
resolveExpression: 100:int
interpretDynamically: if t0 goto 7 else 6
resolveExpression: phi [0: false:bool, 1: boolB] #&&
interpretDynamically: return 200:int
resolveExpression: 200:int
interpretDynamically: return 200:int
resolveExpression: 200:int
interpretDynamically: jump 2
interpretDynamically: phi [0: false:bool, 1: boolB] #&&
resolveExpression: parameter boolB : bool
interpretDynamically: if boolA goto 3 else 4
resolveExpression: parameter boolA : bool
interpretDynamically: phi [2: true:bool, 3: boolB] #||
resolveExpression: true:bool
interpretDynamically: if t0 goto 7 else 6
resolveExpression: phi [0: false:bool, 1: boolB] #&&
interpretDynamically: if t1 goto 5 else 6
resolveExpression: phi [2: true:bool, 3: boolB] #||
interpretDynamically: return 200:int
resolveExpression: 200:int
interpretDynamically: return 200:int
resolveExpression: 200:int
interpretDynamically: return 100:int
resolveExpression: 100:int
interpretDynamically: jump 4
interpretDynamically: phi [2: true:bool, 3: boolB] #||
resolveExpression: parameter boolB : bool
interpretDynamically: if t0 goto 7 else 6
resolveExpression: phi [0: false:bool, 1: boolB] #&&
interpretDynamically: return 200:int
resolveExpression: 200:int
interpretDynamically: if t1 goto 5 else 6
resolveExpression: phi [2: true:bool, 3: boolB] #||
interpretDynamically: return 100:int
resolveExpression: 100:int
interpretDynamically: return 200:int
resolveExpression: 200:int
interpretDynamically: if t1 goto 5 else 6
resolveExpression: phi [2: true:bool, 3: boolB] #||
interpretDynamically: return 200:int
resolveExpression: 200:int
interpretDynamically: return 100:int
resolveExpression: 100:int
States: 12
---- State #0 ----
Interpreter{
  PC: ((((true & (!BoolSegment_sel0)) & (!BoolSegment_sel1)) & false) & true)
  ReturnResults:
    [0] 100
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=2, store=0, select=2)
        op[0]: select idx=addr_0, res=BoolSegment_sel0
        op[1]: select idx=addr_0, res=BoolSegment_sel1
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment.select(addr_0, BoolSegment_sel0).select(addr_0, BoolSegment_sel1)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #1 ----
Interpreter{
  PC: (((true & (!BoolSegment_sel0)) & BoolSegment_sel1) & (!false))
  ReturnResults:
    [0] 200
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=bool)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=3, store=0, select=3)
        op[0]: select idx=addr_0, res=BoolSegment_sel0
        op[1]: select idx=addr_0, res=BoolSegment_sel1
        op[2]: select idx=addr_1, res=BoolSegment_sel2
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment.select(addr_0, BoolSegment_sel0).select(addr_0, BoolSegment_sel1).select(addr_1, BoolSegment_sel2)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #2 ----
Interpreter{
  PC: (((true & (!BoolSegment_sel0)) & (!BoolSegment_sel1)) & (!false))
  ReturnResults:
    [0] 200
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=2, store=0, select=2)
        op[0]: select idx=addr_0, res=BoolSegment_sel0
        op[1]: select idx=addr_0, res=BoolSegment_sel1
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment.select(addr_0, BoolSegment_sel0).select(addr_0, BoolSegment_sel1)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
BoolSegment -> ((as const (Array (_ BitVec 64) Bool)) false)
BoolSegment_sel1 -> false
BoolSegment_sel0 -> false

---- State #3 ----
Interpreter{
  PC: ((((true & BoolSegment_sel0) & (!BoolSegment_sel2)) & BoolSegment_sel1) & (!true))
  ReturnResults:
    [0] 200
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=bool)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=3, store=0, select=3)
        op[0]: select idx=addr_0, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
        op[2]: select idx=addr_0, res=BoolSegment_sel2
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment.select(addr_0, BoolSegment_sel0).select(addr_1, BoolSegment_sel1).select(addr_0, BoolSegment_sel2)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #4 ----
Interpreter{
  PC: (((true & BoolSegment_sel0) & (!BoolSegment_sel2)) & (!BoolSegment_sel1))
  ReturnResults:
    [0] 200
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=bool)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=3, store=0, select=3)
        op[0]: select idx=addr_0, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
        op[2]: select idx=addr_0, res=BoolSegment_sel2
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment.select(addr_0, BoolSegment_sel0).select(addr_1, BoolSegment_sel1).select(addr_0, BoolSegment_sel2)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #5 ----
Interpreter{
  PC: ((((true & BoolSegment_sel0) & (!BoolSegment_sel2)) & BoolSegment_sel1) & true)
  ReturnResults:
    [0] 100
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=bool)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=3, store=0, select=3)
        op[0]: select idx=addr_0, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
        op[2]: select idx=addr_0, res=BoolSegment_sel2
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment.select(addr_0, BoolSegment_sel0).select(addr_1, BoolSegment_sel1).select(addr_0, BoolSegment_sel2)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #6 ----
Interpreter{
  PC: ((((true & BoolSegment_sel0) & BoolSegment_sel2) & BoolSegment_sel1) & BoolSegment_sel3)
  ReturnResults:
    [0] 100
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=bool)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=4, store=0, select=4)
        op[0]: select idx=addr_0, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
        op[2]: select idx=addr_0, res=BoolSegment_sel2
        op[3]: select idx=addr_1, res=BoolSegment_sel3
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment.select(addr_0, BoolSegment_sel0).select(addr_1, BoolSegment_sel1).select(addr_0, BoolSegment_sel2).select(addr_1, BoolSegment_sel3)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
BoolSegment -> ((as const (Array (_ BitVec 64) Bool)) true)
BoolSegment_sel2 -> true
BoolSegment_sel1 -> true
BoolSegment_sel0 -> true
BoolSegment_sel3 -> true

---- State #7 ----
Interpreter{
  PC: ((((true & BoolSegment_sel0) & BoolSegment_sel2) & BoolSegment_sel1) & (!BoolSegment_sel3))
  ReturnResults:
    [0] 200
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=bool)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=4, store=0, select=4)
        op[0]: select idx=addr_0, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
        op[2]: select idx=addr_0, res=BoolSegment_sel2
        op[3]: select idx=addr_1, res=BoolSegment_sel3
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment.select(addr_0, BoolSegment_sel0).select(addr_1, BoolSegment_sel1).select(addr_0, BoolSegment_sel2).select(addr_1, BoolSegment_sel3)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #8 ----
Interpreter{
  PC: (((true & BoolSegment_sel0) & BoolSegment_sel2) & (!BoolSegment_sel1))
  ReturnResults:
    [0] 200
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=bool)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=4, store=0, select=4)
        op[0]: select idx=addr_0, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
        op[2]: select idx=addr_0, res=BoolSegment_sel2
        op[3]: select idx=addr_1, res=BoolSegment_sel3
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment.select(addr_0, BoolSegment_sel0).select(addr_1, BoolSegment_sel1).select(addr_0, BoolSegment_sel2).select(addr_1, BoolSegment_sel3)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
addr_1 -> #xffffffffffffffff
addr_0 -> #x0000000000000000
BoolSegment -> (store ((as const (Array (_ BitVec 64) Bool)) false) #x0000000000000000 true)
BoolSegment_sel2 -> true
BoolSegment_sel1 -> false
BoolSegment_sel0 -> true
BoolSegment_sel3 -> false

---- State #9 ----
Interpreter{
  PC: ((((true & (!BoolSegment_sel0)) & (!BoolSegment_sel1)) & false) & (!true))
  ReturnResults:
    [0] 200
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=2, store=0, select=2)
        op[0]: select idx=addr_0, res=BoolSegment_sel0
        op[1]: select idx=addr_0, res=BoolSegment_sel1
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment.select(addr_0, BoolSegment_sel0).select(addr_0, BoolSegment_sel1)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #10 ----
Interpreter{
  PC: ((((true & (!BoolSegment_sel0)) & BoolSegment_sel1) & false) & BoolSegment_sel2)
  ReturnResults:
    [0] 100
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=bool)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=3, store=0, select=3)
        op[0]: select idx=addr_0, res=BoolSegment_sel0
        op[1]: select idx=addr_0, res=BoolSegment_sel1
        op[2]: select idx=addr_1, res=BoolSegment_sel2
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment.select(addr_0, BoolSegment_sel0).select(addr_0, BoolSegment_sel1).select(addr_1, BoolSegment_sel2)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #11 ----
Interpreter{
  PC: ((((true & (!BoolSegment_sel0)) & BoolSegment_sel1) & false) & (!BoolSegment_sel2))
  ReturnResults:
    [0] 200
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=bool)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=3, store=0, select=3)
        op[0]: select idx=addr_0, res=BoolSegment_sel0
        op[1]: select idx=addr_0, res=BoolSegment_sel1
        op[2]: select idx=addr_1, res=BoolSegment_sel2
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment.select(addr_0, BoolSegment_sel0).select(addr_0, BoolSegment_sel1).select(addr_1, BoolSegment_sel2)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT

==================== BooleanXorCompare ====================
CFG построен для функции с 3 блоками
interpretDynamically: aBool != bBool
interpretDynamically: if t0 goto 1 else 2
resolveExpression: aBool != bBool
resolveExpression: parameter aBool : bool
resolveExpression: parameter bBool : bool
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: return 0:int
resolveExpression: 0:int
States: 2
---- State #0 ----
Interpreter{
  PC: (true & (!(BoolSegment_sel0 != BoolSegment_sel1)))
  ReturnResults:
    [0] 0
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=bool)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=2, store=0, select=2)
        op[0]: select idx=addr_0, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment.select(addr_0, BoolSegment_sel0).select(addr_1, BoolSegment_sel1)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
BoolSegment -> ((as const (Array (_ BitVec 64) Bool)) false)
addr_1 -> #x0000000000000000
BoolSegment_sel0 -> false
BoolSegment_sel1 -> false

---- State #1 ----
Interpreter{
  PC: (true & (BoolSegment_sel0 != BoolSegment_sel1))
  ReturnResults:
    [0] 1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=bool)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=2, store=0, select=2)
        op[0]: select idx=addr_0, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment.select(addr_0, BoolSegment_sel0).select(addr_1, BoolSegment_sel1)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
addr_1 -> #xffffffffffffffff
addr_0 -> #x0000000000000000
BoolSegment -> (store ((as const (Array (_ BitVec 64) Bool)) false) #xffffffffffffffff true)
BoolSegment_sel1 -> true
BoolSegment_sel0 -> false


==================== ShlWithBigLongShift ====================
CFG построен для функции с 5 блоками
interpretDynamically: shift < 40:int64
interpretDynamically: if t0 goto 1 else 2
resolveExpression: shift < 40:int64
resolveExpression: parameter shift : int64
resolveExpression: 40:int64
interpretDynamically: 2004318071:int << shift
interpretDynamically: t1 == 2004318064:int
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: if t2 goto 3 else 4
resolveExpression: t1 == 2004318064:int
resolveExpression: 2004318071:int << shift
resolveExpression: 2004318071:int
resolveExpression: parameter shift : int64
resolveExpression: 2004318064:int
interpretDynamically: return 3:int
resolveExpression: 3:int
interpretDynamically: return 2:int
resolveExpression: 2:int
States: 3
---- State #0 ----
Interpreter{
  PC: (true & (IntSegment_sel0 < 40))
  ReturnResults:
    [0] 1
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
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000027)
addr_0 -> #x0000000000000000
IntSegment_sel0 -> #x0000000000000027

---- State #1 ----
Interpreter{
  PC: ((true & (!(IntSegment_sel0 < 40))) & ((2004318071 << IntSegment_sel1) == 2004318064))
  ReturnResults:
    [0] 2
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

UNSAT
---- State #2 ----
Interpreter{
  PC: ((true & (!(IntSegment_sel0 < 40))) & (!((2004318071 << IntSegment_sel1) == 2004318064)))
  ReturnResults:
    [0] 3
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
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x4000000000000023)
addr_0 -> #x0000000000000000
IntSegment_sel1 -> #x4000000000000023
IntSegment_sel0 -> #x4000000000000023


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
IntSegment -> (store ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x4000000000000030)
                     #xffffffffffffffff
                     #x8000000000000054)
IntSegment_sel1 -> #x8000000000000054
IntSegment_sel0 -> #x4000000000000030

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
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x4000000000000030)
IntSegment_sel2 -> #x4000000000000030
IntSegment_sel1 -> #x4000000000000030
IntSegment_sel3 -> #x4000000000000030
IntSegment_sel0 -> #x4000000000000030


==================== CreateObjectFromValue ====================
CFG построен для функции с 3 блоками
interpretDynamically: value == 0:int
interpretDynamically: if t0 goto 1 else 2
resolveExpression: value == 0:int
resolveExpression: parameter value : int
resolveExpression: 0:int
interpretDynamically: phi [0: value, 1: 1:int] #value
resolveExpression: parameter value : int
interpretDynamically: initialize(t1)
resolveExpression: phi [0: value, 1: 1:int] #value
interpretDynamically: jump 2
interpretDynamically: phi [0: value, 1: 1:int] #value
resolveExpression: 1:int
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

---- State #1 ----
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
interpretDynamically: return nil:*InvokeClass
resolveExpression: nil:*main.InvokeClass
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
States: 2
---- State #0 ----
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

---- State #1 ----
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
interpretDynamically: getNull()
interpretDynamically: getFive()
interpretDynamically: return 5:int
resolveExpression: 5:int
interpretDynamically: return nil:*InvokeClass
resolveExpression: nil:*main.InvokeClass
interpretDynamically: return t3
resolveExpression: getNull()
interpretDynamically: &invokeObject.Value [#0]
interpretDynamically: *t5 = t4
resolveExpression: getFive()
resolveExpression: parameter invokeObject : *InvokeClass
interpretDynamically: return invokeObject
resolveExpression: parameter invokeObject : *InvokeClass
States: 2
---- State #0 ----
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

---- State #1 ----
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


==================== CompareWithDiv ====================
CFG построен для функции с 3 блоками
interpretDynamically: a + 0.5:float64
interpretDynamically: a / t0
interpretDynamically: t1 > b
interpretDynamically: if t2 goto 1 else 2
resolveExpression: t1 > b
resolveExpression: a / t0
resolveExpression: parameter a : float64
resolveExpression: a + 0.5:float64
resolveExpression: parameter a : float64
resolveExpression: 0.5:float64
resolveExpression: parameter b : float64
interpretDynamically: return 0:float64
resolveExpression: 0:float64
interpretDynamically: return 1:float64
resolveExpression: 1:float64
States: 2
---- State #0 ----
Interpreter{
  PC: (true & (!((FloatSegment_sel0 / (FloatSegment_sel1 + 0.5)) > FloatSegment_sel2)))
  ReturnResults:
    [0] 0
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=float)
      P[1]: &(addr_1) (refType=float)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment.select(addr_0, FloatSegment_sel0).select(addr_0, FloatSegment_sel1).select(addr_1, FloatSegment_sel2)
    C[4]: ArraySegment
}

SAT
Model:
addr_0 -> #x0000000000000000
FloatSegment -> ((as const (Array (_ BitVec 64) Real)) (- (/ 1.0 2.0)))
addr_1 -> #x0000000000000000
FloatSegment_sel0 -> (- (/ 1.0 2.0))
FloatSegment_sel1 -> (- (/ 1.0 2.0))
FloatSegment_sel2 -> (- (/ 1.0 2.0))
/0 -> {
  (- (/ 1.0 2.0))
}

---- State #1 ----
Interpreter{
  PC: (true & ((FloatSegment_sel0 / (FloatSegment_sel1 + 0.5)) > FloatSegment_sel2))
  ReturnResults:
    [0] 1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=float)
      P[1]: &(addr_1) (refType=float)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment.select(addr_0, FloatSegment_sel0).select(addr_0, FloatSegment_sel1).select(addr_1, FloatSegment_sel2)
    C[4]: ArraySegment
}

SAT
Model:
addr_0 -> #x0000000000000000
FloatSegment -> ((as const (Array (_ BitVec 64) Real)) (- (/ 1.0 2.0)))
addr_1 -> #x0000000000000000
FloatSegment_sel0 -> (- (/ 1.0 2.0))
FloatSegment_sel1 -> (- (/ 1.0 2.0))
FloatSegment_sel2 -> (- (/ 1.0 2.0))
/0 -> {
  0.0
}


==================== Mul ====================
CFG построен для функции с 7 блоками
interpretDynamically: a * b
interpretDynamically: t0 > 33.32:float64
interpretDynamically: if t1 goto 3 else 2
resolveExpression: t0 > 33.32:float64
resolveExpression: a * b
resolveExpression: parameter a : float64
resolveExpression: parameter b : float64
resolveExpression: 33.32:float64
interpretDynamically: a * b
interpretDynamically: t2 > 33.333:float64
interpretDynamically: if t3 goto 6 else 5
resolveExpression: t2 > 33.333:float64
resolveExpression: a * b
resolveExpression: parameter a : float64
resolveExpression: parameter b : float64
resolveExpression: 33.333:float64
interpretDynamically: return 1.3:float64
resolveExpression: 1.3:float64
interpretDynamically: a * b
interpretDynamically: t6 < 33.7592:float64
interpretDynamically: if t7 goto 4 else 5
resolveExpression: t6 < 33.7592:float64
resolveExpression: a * b
resolveExpression: parameter a : float64
resolveExpression: parameter b : float64
resolveExpression: 33.7592:float64
interpretDynamically: return 1.3:float64
resolveExpression: 1.3:float64
interpretDynamically: return 1.2:float64
resolveExpression: 1.2:float64
interpretDynamically: a * b
interpretDynamically: t4 < 33.333:float64
interpretDynamically: if t5 goto 1 else 2
resolveExpression: t4 < 33.333:float64
resolveExpression: a * b
resolveExpression: parameter a : float64
resolveExpression: parameter b : float64
resolveExpression: 33.333:float64
interpretDynamically: return 1.1:float64
resolveExpression: 1.1:float64
interpretDynamically: a * b
interpretDynamically: t2 > 33.333:float64
interpretDynamically: if t3 goto 6 else 5
resolveExpression: t2 > 33.333:float64
resolveExpression: a * b
resolveExpression: parameter a : float64
resolveExpression: parameter b : float64
resolveExpression: 33.333:float64
interpretDynamically: a * b
interpretDynamically: t6 < 33.7592:float64
interpretDynamically: if t7 goto 4 else 5
resolveExpression: t6 < 33.7592:float64
resolveExpression: a * b
resolveExpression: parameter a : float64
resolveExpression: parameter b : float64
resolveExpression: 33.7592:float64
interpretDynamically: return 1.3:float64
resolveExpression: 1.3:float64
interpretDynamically: return 1.3:float64
resolveExpression: 1.3:float64
interpretDynamically: return 1.2:float64
resolveExpression: 1.2:float64
States: 7
---- State #0 ----
Interpreter{
  PC: ((true & (!((FloatSegment_sel0 * FloatSegment_sel1) > 33.32))) & (!((FloatSegment_sel2 * FloatSegment_sel3) > 33.333)))
  ReturnResults:
    [0] 1.3
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=float)
      P[1]: &(addr_1) (refType=float)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment.select(addr_0, FloatSegment_sel0).select(addr_1, FloatSegment_sel1).select(addr_0, FloatSegment_sel2).select(addr_1, FloatSegment_sel3)
    C[4]: ArraySegment
}

SAT
Model:
addr_0 -> #x0000000000000000
FloatSegment -> ((as const (Array (_ BitVec 64) Real)) 0.0)
addr_1 -> #x0000000000000000
FloatSegment_sel3 -> 0.0
FloatSegment_sel0 -> 0.0
FloatSegment_sel1 -> 0.0
FloatSegment_sel2 -> 0.0

---- State #1 ----
Interpreter{
  PC: (((true & (!((FloatSegment_sel0 * FloatSegment_sel1) > 33.32))) & ((FloatSegment_sel2 * FloatSegment_sel3) > 33.333)) & (!((FloatSegment_sel4 * FloatSegment_sel5) < 33.7592)))
  ReturnResults:
    [0] 1.3
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=float)
      P[1]: &(addr_1) (refType=float)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment.select(addr_0, FloatSegment_sel0).select(addr_1, FloatSegment_sel1).select(addr_0, FloatSegment_sel2).select(addr_1, FloatSegment_sel3).select(addr_0, FloatSegment_sel4).select(addr_1, FloatSegment_sel5)
    C[4]: ArraySegment
}

UNSAT
---- State #2 ----
Interpreter{
  PC: (((true & (!((FloatSegment_sel0 * FloatSegment_sel1) > 33.32))) & ((FloatSegment_sel2 * FloatSegment_sel3) > 33.333)) & ((FloatSegment_sel4 * FloatSegment_sel5) < 33.7592))
  ReturnResults:
    [0] 1.2
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=float)
      P[1]: &(addr_1) (refType=float)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment.select(addr_0, FloatSegment_sel0).select(addr_1, FloatSegment_sel1).select(addr_0, FloatSegment_sel2).select(addr_1, FloatSegment_sel3).select(addr_0, FloatSegment_sel4).select(addr_1, FloatSegment_sel5)
    C[4]: ArraySegment
}

UNSAT
---- State #3 ----
Interpreter{
  PC: ((true & ((FloatSegment_sel0 * FloatSegment_sel1) > 33.32)) & ((FloatSegment_sel2 * FloatSegment_sel3) < 33.333))
  ReturnResults:
    [0] 1.1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=float)
      P[1]: &(addr_1) (refType=float)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment.select(addr_0, FloatSegment_sel0).select(addr_1, FloatSegment_sel1).select(addr_0, FloatSegment_sel2).select(addr_1, FloatSegment_sel3)
    C[4]: ArraySegment
}

SAT
Model:
addr_0 -> #x0000000000000000
FloatSegment -> ((as const (Array (_ BitVec 64) Real)) (/ 739.0 128.0))
addr_1 -> #x0000000000000000
FloatSegment_sel3 -> (/ 739.0 128.0)
FloatSegment_sel0 -> (/ 739.0 128.0)
FloatSegment_sel1 -> (/ 739.0 128.0)
FloatSegment_sel2 -> (/ 739.0 128.0)

---- State #4 ----
Interpreter{
  PC: ((((true & ((FloatSegment_sel0 * FloatSegment_sel1) > 33.32)) & (!((FloatSegment_sel2 * FloatSegment_sel3) < 33.333))) & ((FloatSegment_sel4 * FloatSegment_sel5) > 33.333)) & (!((FloatSegment_sel6 * FloatSegment_sel7) < 33.7592)))
  ReturnResults:
    [0] 1.3
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=float)
      P[1]: &(addr_1) (refType=float)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment.select(addr_0, FloatSegment_sel0).select(addr_1, FloatSegment_sel1).select(addr_0, FloatSegment_sel2).select(addr_1, FloatSegment_sel3).select(addr_0, FloatSegment_sel4).select(addr_1, FloatSegment_sel5).select(addr_0, FloatSegment_sel6).select(addr_1, FloatSegment_sel7)
    C[4]: ArraySegment
}

SAT
Model:
addr_0 -> #x0000000000000000
FloatSegment -> ((as const (Array (_ BitVec 64) Real)) 6.0)
addr_1 -> #x0000000000000000
FloatSegment_sel4 -> 6.0
FloatSegment_sel6 -> 6.0
FloatSegment_sel3 -> 6.0
FloatSegment_sel7 -> 6.0
FloatSegment_sel0 -> 6.0
FloatSegment_sel5 -> 6.0
FloatSegment_sel1 -> 6.0
FloatSegment_sel2 -> 6.0

---- State #5 ----
Interpreter{
  PC: (((true & ((FloatSegment_sel0 * FloatSegment_sel1) > 33.32)) & (!((FloatSegment_sel2 * FloatSegment_sel3) < 33.333))) & (!((FloatSegment_sel4 * FloatSegment_sel5) > 33.333)))
  ReturnResults:
    [0] 1.3
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=float)
      P[1]: &(addr_1) (refType=float)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment.select(addr_0, FloatSegment_sel0).select(addr_1, FloatSegment_sel1).select(addr_0, FloatSegment_sel2).select(addr_1, FloatSegment_sel3).select(addr_0, FloatSegment_sel4).select(addr_1, FloatSegment_sel5)
    C[4]: ArraySegment
}

SAT
Model:
addr_0 -> #x0000000000000000
FloatSegment -> ((as const (Array (_ BitVec 64) Real))
                  (root-obj (+ (* 1000 (^ x 2)) (- 33333)) 2))
addr_1 -> #x0000000000000000
FloatSegment_sel3 -> (root-obj (+ (* 1000 (^ x 2)) (- 33333)) 2)
FloatSegment_sel4 -> (root-obj (+ (* 1000 (^ x 2)) (- 33333)) 2)
FloatSegment_sel0 -> (root-obj (+ (* 1000 (^ x 2)) (- 33333)) 2)
FloatSegment_sel5 -> (root-obj (+ (* 1000 (^ x 2)) (- 33333)) 2)
FloatSegment_sel1 -> (root-obj (+ (* 1000 (^ x 2)) (- 33333)) 2)
FloatSegment_sel2 -> (root-obj (+ (* 1000 (^ x 2)) (- 33333)) 2)

---- State #6 ----
Interpreter{
  PC: ((((true & ((FloatSegment_sel0 * FloatSegment_sel1) > 33.32)) & (!((FloatSegment_sel2 * FloatSegment_sel3) < 33.333))) & ((FloatSegment_sel4 * FloatSegment_sel5) > 33.333)) & ((FloatSegment_sel6 * FloatSegment_sel7) < 33.7592))
  ReturnResults:
    [0] 1.2
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=float)
      P[1]: &(addr_1) (refType=float)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment.select(addr_0, FloatSegment_sel0).select(addr_1, FloatSegment_sel1).select(addr_0, FloatSegment_sel2).select(addr_1, FloatSegment_sel3).select(addr_0, FloatSegment_sel4).select(addr_1, FloatSegment_sel5).select(addr_0, FloatSegment_sel6).select(addr_1, FloatSegment_sel7)
    C[4]: ArraySegment
}

SAT
Model:
addr_0 -> #x0000000000000000
FloatSegment -> ((as const (Array (_ BitVec 64) Real)) (/ 185.0 32.0))
addr_1 -> #x0000000000000000
FloatSegment_sel4 -> (/ 185.0 32.0)
FloatSegment_sel6 -> (/ 185.0 32.0)
FloatSegment_sel3 -> (/ 185.0 32.0)
FloatSegment_sel7 -> (/ 185.0 32.0)
FloatSegment_sel0 -> (/ 185.0 32.0)
FloatSegment_sel5 -> (/ 185.0 32.0)
FloatSegment_sel1 -> (/ 185.0 32.0)
FloatSegment_sel2 -> (/ 185.0 32.0)


==================== LoopWithConcreteBound ====================
CFG построен для функции с 4 блоками
interpretDynamically: jump 3
interpretDynamically: phi [0: 0:int, 1: t0] #result
resolveExpression: 0:int
interpretDynamically: phi [0: 0:int, 1: t1] #i
resolveExpression: 0:int
interpretDynamically: t3 < 10:int
interpretDynamically: if t4 goto 1 else 2
resolveExpression: t3 < 10:int
resolveExpression: phi [0: 0:int, 1: t1] #i
resolveExpression: 10:int
interpretDynamically: t2 + t3
interpretDynamically: t3 + 1:int
interpretDynamically: return t2
resolveExpression: phi [0: 0:int, 1: t0] #result
interpretDynamically: jump 3
interpretDynamically: phi [0: 0:int, 1: t0] #result
resolveExpression: t2 + t3
resolveExpression: phi [0: 0:int, 1: t0] #result
resolveExpression: phi [0: 0:int, 1: t1] #i
interpretDynamically: phi [0: 0:int, 1: t1] #i
resolveExpression: t3 + 1:int
resolveExpression: phi [0: 0:int, 1: t1] #i
resolveExpression: 1:int
interpretDynamically: t3 < 10:int
interpretDynamically: if t4 goto 1 else 2
resolveExpression: t3 < 10:int
resolveExpression: phi [0: 0:int, 1: t1] #i
resolveExpression: 10:int
interpretDynamically: t2 + t3
interpretDynamically: t3 + 1:int
interpretDynamically: jump 3
interpretDynamically: phi [0: 0:int, 1: t0] #result
resolveExpression: t2 + t3
resolveExpression: phi [0: 0:int, 1: t0] #result
resolveExpression: phi [0: 0:int, 1: t1] #i
interpretDynamically: phi [0: 0:int, 1: t1] #i
resolveExpression: t3 + 1:int
resolveExpression: phi [0: 0:int, 1: t1] #i
resolveExpression: 1:int
interpretDynamically: t3 < 10:int
interpretDynamically: if t4 goto 1 else 2
resolveExpression: t3 < 10:int
resolveExpression: phi [0: 0:int, 1: t1] #i
resolveExpression: 10:int
interpretDynamically: return t2
resolveExpression: phi [0: 0:int, 1: t0] #result
interpretDynamically: t2 + t3
interpretDynamically: t3 + 1:int
interpretDynamically: return t2
resolveExpression: phi [0: 0:int, 1: t0] #result
interpretDynamically: jump 3
interpretDynamically: phi [0: 0:int, 1: t0] #result
resolveExpression: t2 + t3
resolveExpression: phi [0: 0:int, 1: t0] #result
resolveExpression: phi [0: 0:int, 1: t1] #i
interpretDynamically: phi [0: 0:int, 1: t1] #i
resolveExpression: t3 + 1:int
resolveExpression: phi [0: 0:int, 1: t1] #i
resolveExpression: 1:int
interpretDynamically: t3 < 10:int
interpretDynamically: if t4 goto 1 else 2
resolveExpression: t3 < 10:int
resolveExpression: phi [0: 0:int, 1: t1] #i
resolveExpression: 10:int
interpretDynamically: t2 + t3
interpretDynamically: return t2
resolveExpression: phi [0: 0:int, 1: t0] #result
interpretDynamically: t3 + 1:int
interpretDynamically: jump 3
interpretDynamically: phi [0: 0:int, 1: t0] #result
resolveExpression: t2 + t3
resolveExpression: phi [0: 0:int, 1: t0] #result
resolveExpression: phi [0: 0:int, 1: t1] #i
interpretDynamically: phi [0: 0:int, 1: t1] #i
resolveExpression: t3 + 1:int
resolveExpression: phi [0: 0:int, 1: t1] #i
resolveExpression: 1:int
interpretDynamically: t3 < 10:int
interpretDynamically: if t4 goto 1 else 2
resolveExpression: t3 < 10:int
resolveExpression: phi [0: 0:int, 1: t1] #i
resolveExpression: 10:int
interpretDynamically: t2 + t3
interpretDynamically: t3 + 1:int
interpretDynamically: jump 3
interpretDynamically: return t2
resolveExpression: phi [0: 0:int, 1: t0] #result
States: 5
---- State #0 ----
Interpreter{
  PC: (true & (!(0 < 10)))
  ReturnResults:
    [0] 0
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 0
    Params: 0, Allocations: 0
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #1 ----
Interpreter{
  PC: (((true & (0 < 10)) & ((0 + 1) < 10)) & (!(((0 + 1) + 1) < 10)))
  ReturnResults:
    [0] ((0 + 0) + (0 + 1))
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 0
    Params: 0, Allocations: 0
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #2 ----
Interpreter{
  PC: ((true & (0 < 10)) & (!((0 + 1) < 10)))
  ReturnResults:
    [0] (0 + 0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 0
    Params: 0, Allocations: 0
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #3 ----
Interpreter{
  PC: ((((true & (0 < 10)) & ((0 + 1) < 10)) & (((0 + 1) + 1) < 10)) & (!((((0 + 1) + 1) + 1) < 10)))
  ReturnResults:
    [0] (((0 + 0) + (0 + 1)) + ((0 + 1) + 1))
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 0
    Params: 0, Allocations: 0
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #4 ----
Interpreter{
  PC: (((((true & (0 < 10)) & ((0 + 1) < 10)) & (((0 + 1) + 1) < 10)) & ((((0 + 1) + 1) + 1) < 10)) & (!(((((0 + 1) + 1) + 1) + 1) < 10)))
  ReturnResults:
    [0] ((((0 + 0) + (0 + 1)) + ((0 + 1) + 1)) + (((0 + 1) + 1) + 1))
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 0
    Params: 0, Allocations: 0
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT

==================== LoopWithSymbolicBound ====================
CFG построен для функции с 6 блоками
interpretDynamically: n > 10:int
interpretDynamically: if t0 goto 1 else 2
resolveExpression: n > 10:int
resolveExpression: parameter n : int
resolveExpression: 10:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 3: t1] #result
resolveExpression: 0:int
interpretDynamically: phi [2: 0:int, 3: t2] #i
resolveExpression: 0:int
interpretDynamically: t4 < n
interpretDynamically: if t5 goto 3 else 4
resolveExpression: t4 < n
resolveExpression: phi [2: 0:int, 3: t2] #i
resolveExpression: parameter n : int
interpretDynamically: return t3
resolveExpression: phi [2: 0:int, 3: t1] #result
interpretDynamically: return -1:int
resolveExpression: -1:int
interpretDynamically: t3 + t4
interpretDynamically: t4 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 3: t1] #result
resolveExpression: t3 + t4
resolveExpression: phi [2: 0:int, 3: t1] #result
resolveExpression: phi [2: 0:int, 3: t2] #i
interpretDynamically: phi [2: 0:int, 3: t2] #i
resolveExpression: t4 + 1:int
resolveExpression: phi [2: 0:int, 3: t2] #i
resolveExpression: 1:int
interpretDynamically: t4 < n
interpretDynamically: if t5 goto 3 else 4
resolveExpression: t4 < n
resolveExpression: phi [2: 0:int, 3: t2] #i
resolveExpression: parameter n : int
interpretDynamically: return t3
resolveExpression: phi [2: 0:int, 3: t1] #result
interpretDynamically: t3 + t4
interpretDynamically: t4 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 3: t1] #result
resolveExpression: t3 + t4
resolveExpression: phi [2: 0:int, 3: t1] #result
resolveExpression: phi [2: 0:int, 3: t2] #i
interpretDynamically: phi [2: 0:int, 3: t2] #i
resolveExpression: t4 + 1:int
resolveExpression: phi [2: 0:int, 3: t2] #i
resolveExpression: 1:int
interpretDynamically: t4 < n
interpretDynamically: if t5 goto 3 else 4
resolveExpression: t4 < n
resolveExpression: phi [2: 0:int, 3: t2] #i
resolveExpression: parameter n : int
interpretDynamically: return t3
resolveExpression: phi [2: 0:int, 3: t1] #result
interpretDynamically: t3 + t4
interpretDynamically: t4 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 3: t1] #result
resolveExpression: t3 + t4
resolveExpression: phi [2: 0:int, 3: t1] #result
resolveExpression: phi [2: 0:int, 3: t2] #i
interpretDynamically: phi [2: 0:int, 3: t2] #i
resolveExpression: t4 + 1:int
resolveExpression: phi [2: 0:int, 3: t2] #i
resolveExpression: 1:int
interpretDynamically: t4 < n
interpretDynamically: if t5 goto 3 else 4
resolveExpression: t4 < n
resolveExpression: phi [2: 0:int, 3: t2] #i
resolveExpression: parameter n : int
interpretDynamically: t3 + t4
interpretDynamically: t4 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 3: t1] #result
resolveExpression: t3 + t4
resolveExpression: phi [2: 0:int, 3: t1] #result
resolveExpression: phi [2: 0:int, 3: t2] #i
interpretDynamically: phi [2: 0:int, 3: t2] #i
resolveExpression: t4 + 1:int
resolveExpression: phi [2: 0:int, 3: t2] #i
resolveExpression: 1:int
interpretDynamically: t4 < n
interpretDynamically: if t5 goto 3 else 4
resolveExpression: t4 < n
resolveExpression: phi [2: 0:int, 3: t2] #i
resolveExpression: parameter n : int
interpretDynamically: t3 + t4
interpretDynamically: t4 + 1:int
interpretDynamically: return t3
resolveExpression: phi [2: 0:int, 3: t1] #result
interpretDynamically: jump 5
interpretDynamically: return t3
resolveExpression: phi [2: 0:int, 3: t1] #result
States: 6
---- State #0 ----
Interpreter{
  PC: ((true & (!(IntSegment_sel0 > 10))) & (!(0 < IntSegment_sel1)))
  ReturnResults:
    [0] 0
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
addr_0 -> #x0000000000000000
IntSegment_sel1 -> #x0000000000000000
IntSegment_sel0 -> #x0000000000000000

---- State #1 ----
Interpreter{
  PC: (true & (IntSegment_sel0 > 10))
  ReturnResults:
    [0] -1
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
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x000000000000000b)
addr_0 -> #x0000000000000000
IntSegment_sel0 -> #x000000000000000b

---- State #2 ----
Interpreter{
  PC: (((true & (!(IntSegment_sel0 > 10))) & (0 < IntSegment_sel1)) & (!((0 + 1) < IntSegment_sel2)))
  ReturnResults:
    [0] (0 + 0)
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
addr_0 -> #x0000000000000000
IntSegment_sel2 -> #x0000000000000001
IntSegment_sel1 -> #x0000000000000001
IntSegment_sel0 -> #x0000000000000001

---- State #3 ----
Interpreter{
  PC: ((((((true & (!(IntSegment_sel0 > 10))) & (0 < IntSegment_sel1)) & ((0 + 1) < IntSegment_sel2)) & (((0 + 1) + 1) < IntSegment_sel3)) & ((((0 + 1) + 1) + 1) < IntSegment_sel4)) & (!(((((0 + 1) + 1) + 1) + 1) < IntSegment_sel5)))
  ReturnResults:
    [0] ((((0 + 0) + (0 + 1)) + ((0 + 1) + 1)) + (((0 + 1) + 1) + 1))
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=6, store=0, select=6)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000004)
addr_0 -> #x0000000000000000
IntSegment_sel2 -> #x0000000000000004
IntSegment_sel5 -> #x0000000000000004
IntSegment_sel1 -> #x0000000000000004
IntSegment_sel4 -> #x0000000000000004
IntSegment_sel3 -> #x0000000000000004
IntSegment_sel0 -> #x0000000000000004

---- State #4 ----
Interpreter{
  PC: (((((true & (!(IntSegment_sel0 > 10))) & (0 < IntSegment_sel1)) & ((0 + 1) < IntSegment_sel2)) & (((0 + 1) + 1) < IntSegment_sel3)) & (!((((0 + 1) + 1) + 1) < IntSegment_sel4)))
  ReturnResults:
    [0] (((0 + 0) + (0 + 1)) + ((0 + 1) + 1))
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=5, store=0, select=5)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000003)
addr_0 -> #x0000000000000000
IntSegment_sel2 -> #x0000000000000003
IntSegment_sel1 -> #x0000000000000003
IntSegment_sel4 -> #x0000000000000003
IntSegment_sel3 -> #x0000000000000003
IntSegment_sel0 -> #x0000000000000003

---- State #5 ----
Interpreter{
  PC: ((((true & (!(IntSegment_sel0 > 10))) & (0 < IntSegment_sel1)) & ((0 + 1) < IntSegment_sel2)) & (!(((0 + 1) + 1) < IntSegment_sel3)))
  ReturnResults:
    [0] ((0 + 0) + (0 + 1))
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000002)
addr_0 -> #x0000000000000000
IntSegment_sel2 -> #x0000000000000002
IntSegment_sel1 -> #x0000000000000002
IntSegment_sel3 -> #x0000000000000002
IntSegment_sel0 -> #x0000000000000002


==================== LoopWithSymbolicBoundAndSymbolicBranching ====================
CFG построен для функции с 9 блоками
interpretDynamically: n > 10:int
interpretDynamically: if t0 goto 1 else 2
resolveExpression: n > 10:int
resolveExpression: parameter n : int
resolveExpression: 10:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 7: t5] #result
resolveExpression: 0:int
interpretDynamically: phi [2: 0:int, 7: t6] #i
resolveExpression: 0:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: parameter n : int
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
interpretDynamically: t2 % 2:int
interpretDynamically: t7 == 0:int
interpretDynamically: if t8 goto 6 else 7
resolveExpression: t7 == 0:int
resolveExpression: t2 % 2:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 2:int
resolveExpression: 0:int
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: t1 + t2
interpretDynamically: jump 7
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: t1 + t2
resolveExpression: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [2: 0:int, 7: t6] #i
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [3: t1, 8: t1, 6: t4] #result
interpretDynamically: phi [2: 0:int, 7: t6] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: parameter n : int
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [3: t1, 8: t1, 6: t4] #result
interpretDynamically: phi [2: 0:int, 7: t6] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: parameter n : int
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
interpretDynamically: t2 % 2:int
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: t7 == 0:int
interpretDynamically: if t8 goto 6 else 7
resolveExpression: t7 == 0:int
resolveExpression: t2 % 2:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 2:int
resolveExpression: 0:int
interpretDynamically: t1 + t2
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: jump 7
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [3: t1, 8: t1, 6: t4] #result
interpretDynamically: phi [2: 0:int, 7: t6] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: parameter n : int
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [3: t1, 8: t1, 6: t4] #result
interpretDynamically: t2 % 2:int
interpretDynamically: t7 == 0:int
interpretDynamically: if t8 goto 6 else 7
resolveExpression: t7 == 0:int
resolveExpression: t2 % 2:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 2:int
resolveExpression: 0:int
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: t1 + t2
interpretDynamically: jump 7
interpretDynamically: phi [2: 0:int, 7: t6] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: parameter n : int
interpretDynamically: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [3: t1, 8: t1, 6: t4] #result
interpretDynamically: phi [2: 0:int, 7: t6] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: parameter n : int
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: t2 + 1:int
interpretDynamically: t2 % 2:int
interpretDynamically: t7 == 0:int
interpretDynamically: jump 5
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: t2 + 1:int
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [3: t1, 8: t1, 6: t4] #result
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: t1 + t2
resolveExpression: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [2: 0:int, 7: t6] #i
interpretDynamically: phi [2: 0:int, 7: t6] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: parameter n : int
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
interpretDynamically: t2 % 2:int
interpretDynamically: t7 == 0:int
interpretDynamically: if t8 goto 6 else 7
resolveExpression: t7 == 0:int
resolveExpression: t2 % 2:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 2:int
resolveExpression: 0:int
interpretDynamically: t1 + t2
interpretDynamically: jump 7
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [3: t1, 8: t1, 6: t4] #result
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: phi [2: 0:int, 7: t6] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: parameter n : int
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: t2 % 2:int
interpretDynamically: t7 == 0:int
interpretDynamically: if t8 goto 6 else 7
resolveExpression: t7 == 0:int
resolveExpression: t2 % 2:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 2:int
resolveExpression: 0:int
interpretDynamically: t1 + t2
interpretDynamically: jump 7
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: t2 + 1:int
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: t1 + t2
resolveExpression: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [2: 0:int, 7: t6] #i
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: jump 5
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [3: t1, 8: t1, 6: t4] #result
interpretDynamically: phi [2: 0:int, 7: t6] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: parameter n : int
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [3: t1, 8: t1, 6: t4] #result
interpretDynamically: phi [2: 0:int, 7: t6] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: parameter n : int
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
interpretDynamically: t2 % 2:int
interpretDynamically: t2 % 2:int
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: t1 + t2
resolveExpression: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [2: 0:int, 7: t6] #i
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [3: t1, 8: t1, 6: t4] #result
interpretDynamically: phi [2: 0:int, 7: t6] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [3: t1, 8: t1, 6: t4] #result
interpretDynamically: phi [2: 0:int, 7: t6] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: parameter n : int
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
interpretDynamically: t2 % 2:int
interpretDynamically: t7 == 0:int
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: if t8 goto 6 else 7
resolveExpression: t7 == 0:int
resolveExpression: t2 % 2:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 2:int
resolveExpression: 0:int
interpretDynamically: t1 + t2
interpretDynamically: jump 7
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: t1 + t2
resolveExpression: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [2: 0:int, 7: t6] #i
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [3: t1, 8: t1, 6: t4] #result
interpretDynamically: phi [2: 0:int, 7: t6] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: parameter n : int
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: t2 % 2:int
interpretDynamically: t7 == 0:int
interpretDynamically: if t8 goto 6 else 7
resolveExpression: t7 == 0:int
resolveExpression: t2 % 2:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 2:int
resolveExpression: 0:int
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [3: t1, 8: t1, 6: t4] #result
interpretDynamically: phi [2: 0:int, 7: t6] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: parameter n : int
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: t1 + t2
interpretDynamically: jump 7
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: t1 + t2
resolveExpression: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [2: 0:int, 7: t6] #i
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [3: t1, 8: t1, 6: t4] #result
interpretDynamically: phi [2: 0:int, 7: t6] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: parameter n : int
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
interpretDynamically: t2 % 2:int
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: t7 == 0:int
interpretDynamically: if t8 goto 6 else 7
resolveExpression: t7 == 0:int
resolveExpression: t2 % 2:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 2:int
resolveExpression: 0:int
interpretDynamically: t1 + t2
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [3: t1, 8: t1, 6: t4] #result
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
interpretDynamically: t2 % 2:int
interpretDynamically: t7 == 0:int
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [3: t1, 8: t1, 6: t4] #result
interpretDynamically: phi [2: 0:int, 7: t6] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t8 goto 6 else 7
resolveExpression: t7 == 0:int
resolveExpression: t2 % 2:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 2:int
resolveExpression: 0:int
interpretDynamically: t1 + t2
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: jump 7
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: t1 + t2
resolveExpression: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [2: 0:int, 7: t6] #i
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [3: t1, 8: t1, 6: t4] #result
interpretDynamically: phi [2: 0:int, 7: t6] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: parameter n : int
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: jump 7
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: t1 + t2
resolveExpression: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [2: 0:int, 7: t6] #i
interpretDynamically: t2 + 1:int
interpretDynamically: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [3: t1, 8: t1, 6: t4] #result
interpretDynamically: phi [2: 0:int, 7: t6] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: parameter n : int
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
interpretDynamically: t2 % 2:int
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 7: t5] #result
resolveExpression: phi [3: t1, 8: t1, 6: t4] #result
interpretDynamically: phi [2: 0:int, 7: t6] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: parameter n : int
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
interpretDynamically: t2 % 2:int
interpretDynamically: phi [3: t1, 8: t1, 6: t4] #result
resolveExpression: phi [2: 0:int, 7: t5] #result
interpretDynamically: t2 + 1:int
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 7: t6] #i
resolveExpression: parameter n : int
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
Analyse stopped: context deadline exceeded
States: 11
---- State #0 ----
Interpreter{
  PC: ((true & (!(IntSegment_sel0 > 10))) & (!(0 < IntSegment_sel1)))
  ReturnResults:
    [0] 0
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
addr_0 -> #x0000000000000000
IntSegment_sel1 -> #x0000000000000000
IntSegment_sel0 -> #x0000000000000000

---- State #1 ----
Interpreter{
  PC: (((((((((((((true & (!(IntSegment_sel0 > 10))) & (0 < IntSegment_sel1)) & BoolSegment_sel0) & ((0 % 2) == 0)) & ((0 + 1) < IntSegment_sel2)) & (!BoolSegment_sel1)) & (((0 + 1) + 1) < IntSegment_sel3)) & BoolSegment_sel2) & (!((((0 + 1) + 1) % 2) == 0))) & ((((0 + 1) + 1) + 1) < IntSegment_sel4)) & BoolSegment_sel3) & (!(((((0 + 1) + 1) + 1) % 2) == 0))) & (!(((((0 + 1) + 1) + 1) + 1) < IntSegment_sel5)))
  ReturnResults:
    [0] (0 + 0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=6, store=0, select=6)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
    BoolSegment:   BoolSegment (elem=bool, ops=4, store=0, select=4)
        op[0]: select idx=addr_1, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
        op[2]: select idx=addr_1, res=BoolSegment_sel2
        op[3]: select idx=addr_1, res=BoolSegment_sel3
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5)
    C[1]: BoolSegment.select(addr_1, BoolSegment_sel0).select(addr_1, BoolSegment_sel1).select(addr_1, BoolSegment_sel2).select(addr_1, BoolSegment_sel3)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #2 ----
Interpreter{
  PC: ((((((((((true & (!(IntSegment_sel0 > 10))) & (0 < IntSegment_sel1)) & BoolSegment_sel0) & ((0 % 2) == 0)) & ((0 + 1) < IntSegment_sel2)) & (!BoolSegment_sel1)) & (((0 + 1) + 1) < IntSegment_sel3)) & BoolSegment_sel2) & (!((((0 + 1) + 1) % 2) == 0))) & (!((((0 + 1) + 1) + 1) < IntSegment_sel4)))
  ReturnResults:
    [0] (0 + 0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=5, store=0, select=5)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
    BoolSegment:   BoolSegment (elem=bool, ops=3, store=0, select=3)
        op[0]: select idx=addr_1, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
        op[2]: select idx=addr_1, res=BoolSegment_sel2
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4)
    C[1]: BoolSegment.select(addr_1, BoolSegment_sel0).select(addr_1, BoolSegment_sel1).select(addr_1, BoolSegment_sel2)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #3 ----
Interpreter{
  PC: (((((((((true & (!(IntSegment_sel0 > 10))) & (0 < IntSegment_sel1)) & BoolSegment_sel0) & ((0 % 2) == 0)) & ((0 + 1) < IntSegment_sel2)) & (!BoolSegment_sel1)) & (((0 + 1) + 1) < IntSegment_sel3)) & (!BoolSegment_sel2)) & (!((((0 + 1) + 1) + 1) < IntSegment_sel4)))
  ReturnResults:
    [0] (0 + 0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=5, store=0, select=5)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
    BoolSegment:   BoolSegment (elem=bool, ops=3, store=0, select=3)
        op[0]: select idx=addr_1, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
        op[2]: select idx=addr_1, res=BoolSegment_sel2
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4)
    C[1]: BoolSegment.select(addr_1, BoolSegment_sel0).select(addr_1, BoolSegment_sel1).select(addr_1, BoolSegment_sel2)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #4 ----
Interpreter{
  PC: (((((((((((true & (!(IntSegment_sel0 > 10))) & (0 < IntSegment_sel1)) & BoolSegment_sel0) & ((0 % 2) == 0)) & ((0 + 1) < IntSegment_sel2)) & (!BoolSegment_sel1)) & (((0 + 1) + 1) < IntSegment_sel3)) & (!BoolSegment_sel2)) & ((((0 + 1) + 1) + 1) < IntSegment_sel4)) & (!BoolSegment_sel3)) & (!(((((0 + 1) + 1) + 1) + 1) < IntSegment_sel5)))
  ReturnResults:
    [0] (0 + 0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=6, store=0, select=6)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
    BoolSegment:   BoolSegment (elem=bool, ops=4, store=0, select=4)
        op[0]: select idx=addr_1, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
        op[2]: select idx=addr_1, res=BoolSegment_sel2
        op[3]: select idx=addr_1, res=BoolSegment_sel3
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5)
    C[1]: BoolSegment.select(addr_1, BoolSegment_sel0).select(addr_1, BoolSegment_sel1).select(addr_1, BoolSegment_sel2).select(addr_1, BoolSegment_sel3)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #5 ----
Interpreter{
  PC: ((((((((((((true & (!(IntSegment_sel0 > 10))) & (0 < IntSegment_sel1)) & BoolSegment_sel0) & ((0 % 2) == 0)) & ((0 + 1) < IntSegment_sel2)) & (!BoolSegment_sel1)) & (((0 + 1) + 1) < IntSegment_sel3)) & BoolSegment_sel2) & ((((0 + 1) + 1) % 2) == 0)) & ((((0 + 1) + 1) + 1) < IntSegment_sel4)) & (!BoolSegment_sel3)) & (!(((((0 + 1) + 1) + 1) + 1) < IntSegment_sel5)))
  ReturnResults:
    [0] ((0 + 0) + ((0 + 1) + 1))
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=6, store=0, select=6)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
    BoolSegment:   BoolSegment (elem=bool, ops=4, store=0, select=4)
        op[0]: select idx=addr_1, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
        op[2]: select idx=addr_1, res=BoolSegment_sel2
        op[3]: select idx=addr_1, res=BoolSegment_sel3
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5)
    C[1]: BoolSegment.select(addr_1, BoolSegment_sel0).select(addr_1, BoolSegment_sel1).select(addr_1, BoolSegment_sel2).select(addr_1, BoolSegment_sel3)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #6 ----
Interpreter{
  PC: (((((((true & (!(IntSegment_sel0 > 10))) & (0 < IntSegment_sel1)) & BoolSegment_sel0) & ((0 % 2) == 0)) & ((0 + 1) < IntSegment_sel2)) & (!BoolSegment_sel1)) & (!(((0 + 1) + 1) < IntSegment_sel3)))
  ReturnResults:
    [0] (0 + 0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
    BoolSegment:   BoolSegment (elem=bool, ops=2, store=0, select=2)
        op[0]: select idx=addr_1, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3)
    C[1]: BoolSegment.select(addr_1, BoolSegment_sel0).select(addr_1, BoolSegment_sel1)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #7 ----
Interpreter{
  PC: ((((((((((((true & (!(IntSegment_sel0 > 10))) & (0 < IntSegment_sel1)) & BoolSegment_sel0) & ((0 % 2) == 0)) & ((0 + 1) < IntSegment_sel2)) & (!BoolSegment_sel1)) & (((0 + 1) + 1) < IntSegment_sel3)) & (!BoolSegment_sel2)) & ((((0 + 1) + 1) + 1) < IntSegment_sel4)) & BoolSegment_sel3) & (!(((((0 + 1) + 1) + 1) % 2) == 0))) & (!(((((0 + 1) + 1) + 1) + 1) < IntSegment_sel5)))
  ReturnResults:
    [0] (0 + 0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=6, store=0, select=6)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
    BoolSegment:   BoolSegment (elem=bool, ops=4, store=0, select=4)
        op[0]: select idx=addr_1, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
        op[2]: select idx=addr_1, res=BoolSegment_sel2
        op[3]: select idx=addr_1, res=BoolSegment_sel3
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5)
    C[1]: BoolSegment.select(addr_1, BoolSegment_sel0).select(addr_1, BoolSegment_sel1).select(addr_1, BoolSegment_sel2).select(addr_1, BoolSegment_sel3)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #8 ----
Interpreter{
  PC: ((((true & (!(IntSegment_sel0 > 10))) & (0 < IntSegment_sel1)) & (!BoolSegment_sel0)) & (!((0 + 1) < IntSegment_sel2)))
  ReturnResults:
    [0] 0
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=3, store=0, select=3)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
    BoolSegment:   BoolSegment (elem=bool, ops=1, store=0, select=1)
        op[0]: select idx=addr_1, res=BoolSegment_sel0
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2)
    C[1]: BoolSegment.select(addr_1, BoolSegment_sel0)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
BoolSegment -> ((as const (Array (_ BitVec 64) Bool)) false)
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000001)
addr_0 -> #x0000000000000000
IntSegment_sel2 -> #x0000000000000001
BoolSegment_sel0 -> false
IntSegment_sel1 -> #x0000000000000001
IntSegment_sel0 -> #x0000000000000001

---- State #9 ----
Interpreter{
  PC: (((((((true & (!(IntSegment_sel0 > 10))) & (0 < IntSegment_sel1)) & (!BoolSegment_sel0)) & ((0 + 1) < IntSegment_sel2)) & BoolSegment_sel1) & (((0 + 1) % 2) == 0)) & (!(((0 + 1) + 1) < IntSegment_sel3)))
  ReturnResults:
    [0] (0 + (0 + 1))
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
    BoolSegment:   BoolSegment (elem=bool, ops=2, store=0, select=2)
        op[0]: select idx=addr_1, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3)
    C[1]: BoolSegment.select(addr_1, BoolSegment_sel0).select(addr_1, BoolSegment_sel1)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #10 ----
Interpreter{
  PC: ((((((((((true & (!(IntSegment_sel0 > 10))) & (0 < IntSegment_sel1)) & (!BoolSegment_sel0)) & ((0 + 1) < IntSegment_sel2)) & BoolSegment_sel1) & (!(((0 + 1) % 2) == 0))) & (((0 + 1) + 1) < IntSegment_sel3)) & BoolSegment_sel2) & (!((((0 + 1) + 1) % 2) == 0))) & (!((((0 + 1) + 1) + 1) < IntSegment_sel4)))
  ReturnResults:
    [0] 0
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=5, store=0, select=5)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
    BoolSegment:   BoolSegment (elem=bool, ops=3, store=0, select=3)
        op[0]: select idx=addr_1, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
        op[2]: select idx=addr_1, res=BoolSegment_sel2
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4)
    C[1]: BoolSegment.select(addr_1, BoolSegment_sel0).select(addr_1, BoolSegment_sel1).select(addr_1, BoolSegment_sel2)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT

==================== LoopWithSymbolicBoundAndComplexControlFlow ====================
CFG построен для функции с 10 блоками
interpretDynamically: n > 10:int
interpretDynamically: if t0 goto 1 else 2
resolveExpression: n > 10:int
resolveExpression: parameter n : int
resolveExpression: 10:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 6: t4] #result
resolveExpression: 0:int
interpretDynamically: phi [2: 0:int, 6: t5] #i
resolveExpression: 0:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: parameter n : int
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
interpretDynamically: t2 == 3:int
interpretDynamically: t2 % 2:int
interpretDynamically: t6 != 0:int
interpretDynamically: if t7 goto 6 else 9
resolveExpression: t6 != 0:int
resolveExpression: t2 % 2:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 2:int
resolveExpression: 0:int
interpretDynamically: t1 + t2
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [7: t1, 9: t9] #result
interpretDynamically: phi [2: 0:int, 6: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 1:int
interpretDynamically: jump 6
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: t1 + t2
resolveExpression: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [2: 0:int, 6: t5] #i
interpretDynamically: t2 + 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: parameter n : int
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
interpretDynamically: t2 % 2:int
interpretDynamically: t2 == 3:int
interpretDynamically: if t8 goto 4 else 7
resolveExpression: t2 == 3:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 3:int
interpretDynamically: t2 % 2:int
interpretDynamically: t6 != 0:int
interpretDynamically: if t7 goto 6 else 9
resolveExpression: t6 != 0:int
resolveExpression: t2 % 2:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 2:int
resolveExpression: 0:int
interpretDynamically: t1 + t2
interpretDynamically: jump 6
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: t6 != 0:int
interpretDynamically: if t7 goto 6 else 9
resolveExpression: t6 != 0:int
resolveExpression: t2 % 2:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 2:int
resolveExpression: 0:int
interpretDynamically: t1 + t2
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: jump 6
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: t1 + t2
resolveExpression: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [2: 0:int, 6: t5] #i
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: t1 + t2
resolveExpression: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [2: 0:int, 6: t5] #i
interpretDynamically: t2 + 1:int
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [7: t1, 9: t9] #result
interpretDynamically: phi [2: 0:int, 6: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: parameter n : int
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: t2 % 2:int
interpretDynamically: t6 != 0:int
interpretDynamically: if t7 goto 6 else 9
resolveExpression: t6 != 0:int
resolveExpression: t2 % 2:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 2:int
resolveExpression: 0:int
interpretDynamically: t1 + t2
interpretDynamically: jump 6
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [7: t1, 9: t9] #result
interpretDynamically: phi [2: 0:int, 6: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: parameter n : int
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
interpretDynamically: t2 == 3:int
interpretDynamically: t2 % 2:int
interpretDynamically: t6 != 0:int
interpretDynamically: if t7 goto 6 else 9
resolveExpression: t6 != 0:int
resolveExpression: t2 % 2:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 2:int
resolveExpression: 0:int
interpretDynamically: t1 + t2
interpretDynamically: jump 6
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: t1 + t2
resolveExpression: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [2: 0:int, 6: t5] #i
interpretDynamically: t2 + 1:int
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [7: t1, 9: t9] #result
interpretDynamically: phi [2: 0:int, 6: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: parameter n : int
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
interpretDynamically: t2 % 2:int
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: t1 + t2
resolveExpression: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [2: 0:int, 6: t5] #i
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [7: t1, 9: t9] #result
interpretDynamically: phi [2: 0:int, 6: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 1:int
interpretDynamically: t6 != 0:int
interpretDynamically: if t7 goto 6 else 9
resolveExpression: t6 != 0:int
resolveExpression: t2 % 2:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 2:int
resolveExpression: 0:int
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: t2 + 1:int
interpretDynamically: t1 + t2
interpretDynamically: jump 5
interpretDynamically: jump 6
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: t1 + t2
resolveExpression: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [2: 0:int, 6: t5] #i
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: parameter n : int
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: if t8 goto 4 else 7
resolveExpression: t2 == 3:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 3:int
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: t2 % 2:int
interpretDynamically: t6 != 0:int
interpretDynamically: if t7 goto 6 else 9
resolveExpression: t6 != 0:int
resolveExpression: t2 % 2:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 2:int
resolveExpression: 0:int
interpretDynamically: t1 + t2
interpretDynamically: jump 6
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: t1 + t2
resolveExpression: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [2: 0:int, 6: t5] #i
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [7: t1, 9: t9] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [7: t1, 9: t9] #result
interpretDynamically: phi [2: 0:int, 6: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: parameter n : int
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
interpretDynamically: t2 % 2:int
interpretDynamically: t6 != 0:int
interpretDynamically: if t7 goto 6 else 9
resolveExpression: t6 != 0:int
resolveExpression: t2 % 2:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 2:int
resolveExpression: 0:int
interpretDynamically: t1 + t2
interpretDynamically: t2 == 3:int
interpretDynamically: if t8 goto 4 else 7
resolveExpression: t2 == 3:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 3:int
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: jump 6
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: t1 + t2
resolveExpression: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [2: 0:int, 6: t5] #i
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [7: t1, 9: t9] #result
interpretDynamically: phi [2: 0:int, 6: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 1:int
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [7: t1, 9: t9] #result
interpretDynamically: phi [2: 0:int, 6: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: parameter n : int
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
interpretDynamically: t2 == 3:int
interpretDynamically: t2 % 2:int
interpretDynamically: if t8 goto 4 else 7
resolveExpression: t2 == 3:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 3:int
interpretDynamically: t2 % 2:int
interpretDynamically: t6 != 0:int
interpretDynamically: if t7 goto 6 else 9
resolveExpression: t6 != 0:int
resolveExpression: t2 % 2:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 2:int
resolveExpression: 0:int
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [7: t1, 9: t9] #result
interpretDynamically: t1 + t2
interpretDynamically: jump 6
interpretDynamically: t2 % 2:int
interpretDynamically: t6 != 0:int
interpretDynamically: if t7 goto 6 else 9
resolveExpression: t6 != 0:int
resolveExpression: t2 % 2:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 2:int
resolveExpression: 0:int
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [7: t1, 9: t9] #result
interpretDynamically: phi [2: 0:int, 6: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: t1 + t2
interpretDynamically: jump 6
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: t1 + t2
resolveExpression: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [2: 0:int, 6: t5] #i
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [7: t1, 9: t9] #result
interpretDynamically: phi [2: 0:int, 6: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: parameter n : int
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
interpretDynamically: t2 == 3:int
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: t2 % 2:int
interpretDynamically: t6 != 0:int
interpretDynamically: if t7 goto 6 else 9
resolveExpression: t6 != 0:int
resolveExpression: t2 % 2:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 2:int
resolveExpression: 0:int
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: t1 + t2
interpretDynamically: if t8 goto 4 else 7
resolveExpression: t2 == 3:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 3:int
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: t2 % 2:int
interpretDynamically: t6 != 0:int
interpretDynamically: if t7 goto 6 else 9
resolveExpression: t6 != 0:int
resolveExpression: t2 % 2:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 2:int
resolveExpression: 0:int
interpretDynamically: t1 + t2
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [7: t1, 9: t9] #result
interpretDynamically: jump 6
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: t1 + t2
resolveExpression: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [2: 0:int, 6: t5] #i
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [7: t1, 9: t9] #result
interpretDynamically: phi [2: 0:int, 6: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: parameter n : int
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
interpretDynamically: t2 == 3:int
interpretDynamically: if t8 goto 4 else 7
resolveExpression: t2 == 3:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 3:int
interpretDynamically: t2 % 2:int
interpretDynamically: t6 != 0:int
interpretDynamically: if t7 goto 6 else 9
resolveExpression: t6 != 0:int
resolveExpression: t2 % 2:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 2:int
resolveExpression: 0:int
interpretDynamically: t1 + t2
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: t2 + 1:int
interpretDynamically: jump 6
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: t1 + t2
resolveExpression: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [2: 0:int, 6: t5] #i
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [7: t1, 9: t9] #result
interpretDynamically: phi [2: 0:int, 6: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: t1 + t2
resolveExpression: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [2: 0:int, 6: t5] #i
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [7: t1, 9: t9] #result
interpretDynamically: phi [2: 0:int, 6: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: parameter n : int
interpretDynamically: return t1
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: if condition goto 8 else 7
resolveExpression: parameter condition : bool
interpretDynamically: t2 % 2:int
interpretDynamically: t2 == 3:int
interpretDynamically: if t8 goto 4 else 7
resolveExpression: t2 == 3:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 3:int
interpretDynamically: t2 % 2:int
interpretDynamically: t6 != 0:int
interpretDynamically: t6 != 0:int
interpretDynamically: if t7 goto 6 else 9
resolveExpression: t6 != 0:int
resolveExpression: t2 % 2:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 2:int
resolveExpression: 0:int
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: phi [2: 0:int, 6: t4] #result
interpretDynamically: t1 + t2
interpretDynamically: jump 6
interpretDynamically: phi [7: t1, 9: t9] #result
resolveExpression: t1 + t2
resolveExpression: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [2: 0:int, 6: t5] #i
interpretDynamically: t2 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 6: t4] #result
resolveExpression: phi [7: t1, 9: t9] #result
interpretDynamically: phi [2: 0:int, 6: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: 1:int
interpretDynamically: t2 < n
interpretDynamically: if t3 goto 3 else 4
resolveExpression: t2 < n
resolveExpression: phi [2: 0:int, 6: t5] #i
resolveExpression: parameter n : int
Analyse stopped: context deadline exceeded
States: 11
---- State #0 ----
Interpreter{
  PC: (((((((true & (!(IntSegment_sel0 > 10))) & (0 < IntSegment_sel1)) & (!BoolSegment_sel0)) & ((0 % 2) != 0)) & ((0 + 1) < IntSegment_sel2)) & BoolSegment_sel1) & ((0 + 1) == 3))
  ReturnResults:
    [0] 0
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=3, store=0, select=3)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
    BoolSegment:   BoolSegment (elem=bool, ops=2, store=0, select=2)
        op[0]: select idx=addr_1, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2)
    C[1]: BoolSegment.select(addr_1, BoolSegment_sel0).select(addr_1, BoolSegment_sel1)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #1 ----
Interpreter{
  PC: ((((((((true & (!(IntSegment_sel0 > 10))) & (0 < IntSegment_sel1)) & (!BoolSegment_sel0)) & ((0 % 2) != 0)) & ((0 + 1) < IntSegment_sel2)) & (!BoolSegment_sel1)) & (((0 + 1) % 2) != 0)) & (!(((0 + 1) + 1) < IntSegment_sel3)))
  ReturnResults:
    [0] 0
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
    BoolSegment:   BoolSegment (elem=bool, ops=2, store=0, select=2)
        op[0]: select idx=addr_1, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3)
    C[1]: BoolSegment.select(addr_1, BoolSegment_sel0).select(addr_1, BoolSegment_sel1)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #2 ----
Interpreter{
  PC: (((((((((((true & (!(IntSegment_sel0 > 10))) & (0 < IntSegment_sel1)) & (!BoolSegment_sel0)) & ((0 % 2) != 0)) & ((0 + 1) < IntSegment_sel2)) & (!BoolSegment_sel1)) & (((0 + 1) % 2) != 0)) & (((0 + 1) + 1) < IntSegment_sel3)) & (!BoolSegment_sel2)) & ((((0 + 1) + 1) % 2) != 0)) & (!((((0 + 1) + 1) + 1) < IntSegment_sel4)))
  ReturnResults:
    [0] 0
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=5, store=0, select=5)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
    BoolSegment:   BoolSegment (elem=bool, ops=3, store=0, select=3)
        op[0]: select idx=addr_1, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
        op[2]: select idx=addr_1, res=BoolSegment_sel2
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4)
    C[1]: BoolSegment.select(addr_1, BoolSegment_sel0).select(addr_1, BoolSegment_sel1).select(addr_1, BoolSegment_sel2)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #3 ----
Interpreter{
  PC: ((((((((((((((true & (!(IntSegment_sel0 > 10))) & (0 < IntSegment_sel1)) & (!BoolSegment_sel0)) & ((0 % 2) != 0)) & ((0 + 1) < IntSegment_sel2)) & (!BoolSegment_sel1)) & (((0 + 1) % 2) != 0)) & (((0 + 1) + 1) < IntSegment_sel3)) & (!BoolSegment_sel2)) & ((((0 + 1) + 1) % 2) != 0)) & ((((0 + 1) + 1) + 1) < IntSegment_sel4)) & (!BoolSegment_sel3)) & (!(((((0 + 1) + 1) + 1) % 2) != 0))) & (!(((((0 + 1) + 1) + 1) + 1) < IntSegment_sel5)))
  ReturnResults:
    [0] (0 + (((0 + 1) + 1) + 1))
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=6, store=0, select=6)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
    BoolSegment:   BoolSegment (elem=bool, ops=4, store=0, select=4)
        op[0]: select idx=addr_1, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
        op[2]: select idx=addr_1, res=BoolSegment_sel2
        op[3]: select idx=addr_1, res=BoolSegment_sel3
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5)
    C[1]: BoolSegment.select(addr_1, BoolSegment_sel0).select(addr_1, BoolSegment_sel1).select(addr_1, BoolSegment_sel2).select(addr_1, BoolSegment_sel3)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #4 ----
Interpreter{
  PC: (((((((((((true & (!(IntSegment_sel0 > 10))) & (0 < IntSegment_sel1)) & (!BoolSegment_sel0)) & ((0 % 2) != 0)) & ((0 + 1) < IntSegment_sel2)) & (!BoolSegment_sel1)) & (((0 + 1) % 2) != 0)) & (((0 + 1) + 1) < IntSegment_sel3)) & (!BoolSegment_sel2)) & (!((((0 + 1) + 1) % 2) != 0))) & (!((((0 + 1) + 1) + 1) < IntSegment_sel4)))
  ReturnResults:
    [0] (0 + ((0 + 1) + 1))
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=5, store=0, select=5)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
    BoolSegment:   BoolSegment (elem=bool, ops=3, store=0, select=3)
        op[0]: select idx=addr_1, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
        op[2]: select idx=addr_1, res=BoolSegment_sel2
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4)
    C[1]: BoolSegment.select(addr_1, BoolSegment_sel0).select(addr_1, BoolSegment_sel1).select(addr_1, BoolSegment_sel2)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #5 ----
Interpreter{
  PC: (((((((((true & (!(IntSegment_sel0 > 10))) & (0 < IntSegment_sel1)) & BoolSegment_sel0) & (!(0 == 3))) & (!((0 % 2) != 0))) & ((0 + 1) < IntSegment_sel2)) & (!BoolSegment_sel1)) & (((0 + 1) % 2) != 0)) & (!(((0 + 1) + 1) < IntSegment_sel3)))
  ReturnResults:
    [0] (0 + 0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
    BoolSegment:   BoolSegment (elem=bool, ops=2, store=0, select=2)
        op[0]: select idx=addr_1, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3)
    C[1]: BoolSegment.select(addr_1, BoolSegment_sel0).select(addr_1, BoolSegment_sel1)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #6 ----
Interpreter{
  PC: (((((((((((true & (!(IntSegment_sel0 > 10))) & (0 < IntSegment_sel1)) & BoolSegment_sel0) & (!(0 == 3))) & (!((0 % 2) != 0))) & ((0 + 1) < IntSegment_sel2)) & (!BoolSegment_sel1)) & (((0 + 1) % 2) != 0)) & (((0 + 1) + 1) < IntSegment_sel3)) & BoolSegment_sel2) & (((0 + 1) + 1) == 3))
  ReturnResults:
    [0] (0 + 0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
    BoolSegment:   BoolSegment (elem=bool, ops=3, store=0, select=3)
        op[0]: select idx=addr_1, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
        op[2]: select idx=addr_1, res=BoolSegment_sel2
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3)
    C[1]: BoolSegment.select(addr_1, BoolSegment_sel0).select(addr_1, BoolSegment_sel1).select(addr_1, BoolSegment_sel2)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #7 ----
Interpreter{
  PC: ((((((((((true & (!(IntSegment_sel0 > 10))) & (0 < IntSegment_sel1)) & BoolSegment_sel0) & (!(0 == 3))) & (!((0 % 2) != 0))) & ((0 + 1) < IntSegment_sel2)) & BoolSegment_sel1) & (!((0 + 1) == 3))) & (!(((0 + 1) % 2) != 0))) & (!(((0 + 1) + 1) < IntSegment_sel3)))
  ReturnResults:
    [0] ((0 + 0) + (0 + 1))
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
    BoolSegment:   BoolSegment (elem=bool, ops=2, store=0, select=2)
        op[0]: select idx=addr_1, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3)
    C[1]: BoolSegment.select(addr_1, BoolSegment_sel0).select(addr_1, BoolSegment_sel1)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #8 ----
Interpreter{
  PC: ((((((((((((true & (!(IntSegment_sel0 > 10))) & (0 < IntSegment_sel1)) & BoolSegment_sel0) & (!(0 == 3))) & (!((0 % 2) != 0))) & ((0 + 1) < IntSegment_sel2)) & BoolSegment_sel1) & (!((0 + 1) == 3))) & (!(((0 + 1) % 2) != 0))) & (((0 + 1) + 1) < IntSegment_sel3)) & BoolSegment_sel2) & (((0 + 1) + 1) == 3))
  ReturnResults:
    [0] ((0 + 0) + (0 + 1))
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
    BoolSegment:   BoolSegment (elem=bool, ops=3, store=0, select=3)
        op[0]: select idx=addr_1, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
        op[2]: select idx=addr_1, res=BoolSegment_sel2
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3)
    C[1]: BoolSegment.select(addr_1, BoolSegment_sel0).select(addr_1, BoolSegment_sel1).select(addr_1, BoolSegment_sel2)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #9 ----
Interpreter{
  PC: ((((((((((((((true & (!(IntSegment_sel0 > 10))) & (0 < IntSegment_sel1)) & BoolSegment_sel0) & (!(0 == 3))) & (!((0 % 2) != 0))) & ((0 + 1) < IntSegment_sel2)) & BoolSegment_sel1) & (!((0 + 1) == 3))) & (!(((0 + 1) % 2) != 0))) & (((0 + 1) + 1) < IntSegment_sel3)) & BoolSegment_sel2) & (!(((0 + 1) + 1) == 3))) & (!((((0 + 1) + 1) % 2) != 0))) & (!((((0 + 1) + 1) + 1) < IntSegment_sel4)))
  ReturnResults:
    [0] (((0 + 0) + (0 + 1)) + ((0 + 1) + 1))
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=5, store=0, select=5)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
    BoolSegment:   BoolSegment (elem=bool, ops=3, store=0, select=3)
        op[0]: select idx=addr_1, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
        op[2]: select idx=addr_1, res=BoolSegment_sel2
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4)
    C[1]: BoolSegment.select(addr_1, BoolSegment_sel0).select(addr_1, BoolSegment_sel1).select(addr_1, BoolSegment_sel2)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #10 ----
Interpreter{
  PC: (((((((((((((true & (!(IntSegment_sel0 > 10))) & (0 < IntSegment_sel1)) & BoolSegment_sel0) & (!(0 == 3))) & (!((0 % 2) != 0))) & ((0 + 1) < IntSegment_sel2)) & (!BoolSegment_sel1)) & (((0 + 1) % 2) != 0)) & (((0 + 1) + 1) < IntSegment_sel3)) & BoolSegment_sel2) & (!(((0 + 1) + 1) == 3))) & (!((((0 + 1) + 1) % 2) != 0))) & (!((((0 + 1) + 1) + 1) < IntSegment_sel4)))
  ReturnResults:
    [0] ((0 + 0) + ((0 + 1) + 1))
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=bool)
    IntSegment:    IntSegment (elem=int, ops=5, store=0, select=5)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
    BoolSegment:   BoolSegment (elem=bool, ops=3, store=0, select=3)
        op[0]: select idx=addr_1, res=BoolSegment_sel0
        op[1]: select idx=addr_1, res=BoolSegment_sel1
        op[2]: select idx=addr_1, res=BoolSegment_sel2
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4)
    C[1]: BoolSegment.select(addr_1, BoolSegment_sel0).select(addr_1, BoolSegment_sel1).select(addr_1, BoolSegment_sel2)
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT

==================== WhileCycle ====================
CFG построен для функции с 4 блоками
interpretDynamically: jump 3
interpretDynamically: phi [0: 0:int, 1: t1] #i
resolveExpression: 0:int
interpretDynamically: phi [0: 0:int, 1: t0] #sum
resolveExpression: 0:int
interpretDynamically: t2 < x
interpretDynamically: if t4 goto 1 else 2
resolveExpression: t2 < x
resolveExpression: phi [0: 0:int, 1: t1] #i
resolveExpression: parameter x : int
interpretDynamically: return t3
resolveExpression: phi [0: 0:int, 1: t0] #sum
interpretDynamically: t3 + t2
interpretDynamically: t2 + 1:int
interpretDynamically: jump 3
interpretDynamically: phi [0: 0:int, 1: t1] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [0: 0:int, 1: t1] #i
resolveExpression: 1:int
interpretDynamically: phi [0: 0:int, 1: t0] #sum
resolveExpression: t3 + t2
resolveExpression: phi [0: 0:int, 1: t0] #sum
resolveExpression: phi [0: 0:int, 1: t1] #i
interpretDynamically: t2 < x
interpretDynamically: if t4 goto 1 else 2
resolveExpression: t2 < x
resolveExpression: phi [0: 0:int, 1: t1] #i
resolveExpression: parameter x : int
interpretDynamically: return t3
resolveExpression: phi [0: 0:int, 1: t0] #sum
interpretDynamically: t3 + t2
interpretDynamically: t2 + 1:int
interpretDynamically: jump 3
interpretDynamically: phi [0: 0:int, 1: t1] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [0: 0:int, 1: t1] #i
resolveExpression: 1:int
interpretDynamically: phi [0: 0:int, 1: t0] #sum
resolveExpression: t3 + t2
resolveExpression: phi [0: 0:int, 1: t0] #sum
resolveExpression: phi [0: 0:int, 1: t1] #i
interpretDynamically: t2 < x
interpretDynamically: if t4 goto 1 else 2
resolveExpression: t2 < x
resolveExpression: phi [0: 0:int, 1: t1] #i
resolveExpression: parameter x : int
interpretDynamically: return t3
resolveExpression: phi [0: 0:int, 1: t0] #sum
interpretDynamically: t3 + t2
interpretDynamically: t2 + 1:int
interpretDynamically: jump 3
interpretDynamically: phi [0: 0:int, 1: t1] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [0: 0:int, 1: t1] #i
resolveExpression: 1:int
interpretDynamically: phi [0: 0:int, 1: t0] #sum
resolveExpression: t3 + t2
resolveExpression: phi [0: 0:int, 1: t0] #sum
resolveExpression: phi [0: 0:int, 1: t1] #i
interpretDynamically: t2 < x
interpretDynamically: if t4 goto 1 else 2
resolveExpression: t2 < x
resolveExpression: phi [0: 0:int, 1: t1] #i
resolveExpression: parameter x : int
interpretDynamically: return t3
resolveExpression: phi [0: 0:int, 1: t0] #sum
interpretDynamically: t3 + t2
interpretDynamically: t2 + 1:int
interpretDynamically: jump 3
interpretDynamically: phi [0: 0:int, 1: t1] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [0: 0:int, 1: t1] #i
resolveExpression: 1:int
interpretDynamically: phi [0: 0:int, 1: t0] #sum
resolveExpression: t3 + t2
resolveExpression: phi [0: 0:int, 1: t0] #sum
resolveExpression: phi [0: 0:int, 1: t1] #i
interpretDynamically: t2 < x
interpretDynamically: if t4 goto 1 else 2
resolveExpression: t2 < x
resolveExpression: phi [0: 0:int, 1: t1] #i
resolveExpression: parameter x : int
interpretDynamically: return t3
resolveExpression: phi [0: 0:int, 1: t0] #sum
interpretDynamically: t3 + t2
interpretDynamically: t2 + 1:int
interpretDynamically: jump 3
States: 5
---- State #0 ----
Interpreter{
  PC: (((true & (0 < IntSegment_sel0)) & ((0 + 1) < IntSegment_sel1)) & (!(((0 + 1) + 1) < IntSegment_sel2)))
  ReturnResults:
    [0] ((0 + (0 + 1)) + ((0 + 1) + 1))
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
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000002)
addr_0 -> #x0000000000000000
IntSegment_sel2 -> #x0000000000000002
IntSegment_sel1 -> #x0000000000000002
IntSegment_sel0 -> #x0000000000000002

---- State #1 ----
Interpreter{
  PC: (true & (!(0 < IntSegment_sel0)))
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
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
addr_0 -> #x0000000000000000
IntSegment_sel0 -> #x0000000000000000

---- State #2 ----
Interpreter{
  PC: ((((true & (0 < IntSegment_sel0)) & ((0 + 1) < IntSegment_sel1)) & (((0 + 1) + 1) < IntSegment_sel2)) & (!((((0 + 1) + 1) + 1) < IntSegment_sel3)))
  ReturnResults:
    [0] (((0 + (0 + 1)) + ((0 + 1) + 1)) + (((0 + 1) + 1) + 1))
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000003)
addr_0 -> #x0000000000000000
IntSegment_sel2 -> #x0000000000000003
IntSegment_sel1 -> #x0000000000000003
IntSegment_sel3 -> #x0000000000000003
IntSegment_sel0 -> #x0000000000000003

---- State #3 ----
Interpreter{
  PC: ((true & (0 < IntSegment_sel0)) & (!((0 + 1) < IntSegment_sel1)))
  ReturnResults:
    [0] (0 + (0 + 1))
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
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000001)
addr_0 -> #x0000000000000000
IntSegment_sel1 -> #x0000000000000001
IntSegment_sel0 -> #x0000000000000001

---- State #4 ----
Interpreter{
  PC: (((((true & (0 < IntSegment_sel0)) & ((0 + 1) < IntSegment_sel1)) & (((0 + 1) + 1) < IntSegment_sel2)) & ((((0 + 1) + 1) + 1) < IntSegment_sel3)) & (!(((((0 + 1) + 1) + 1) + 1) < IntSegment_sel4)))
  ReturnResults:
    [0] ((((0 + (0 + 1)) + ((0 + 1) + 1)) + (((0 + 1) + 1) + 1)) + ((((0 + 1) + 1) + 1) + 1))
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=5, store=0, select=5)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000004)
addr_0 -> #x0000000000000000
IntSegment_sel2 -> #x0000000000000004
IntSegment_sel1 -> #x0000000000000004
IntSegment_sel4 -> #x0000000000000004
IntSegment_sel3 -> #x0000000000000004
IntSegment_sel0 -> #x0000000000000004


==================== LoopInsideLoop ====================
CFG построен для функции с 11 блоками
interpretDynamically: x - 5:int
interpretDynamically: jump 3
interpretDynamically: phi [0: t0, 7: t5] #i
resolveExpression: x - 5:int
resolveExpression: parameter x : int
resolveExpression: 5:int
interpretDynamically: t2 < x
interpretDynamically: if t3 goto 1 else 2
resolveExpression: t2 < x
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: parameter x : int
interpretDynamically: t2 < 0:int
interpretDynamically: if t1 goto 4 else 5
resolveExpression: t2 < 0:int
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: 0:int
interpretDynamically: jump 8
interpretDynamically: return 2:int
resolveExpression: 2:int
interpretDynamically: phi [5: t2, 10: t9] #j
resolveExpression: phi [0: t0, 7: t5] #i
interpretDynamically: x + t2
interpretDynamically: t6 < t7
interpretDynamically: if t8 goto 6 else 7
resolveExpression: t6 < t7
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: x + t2
resolveExpression: parameter x : int
resolveExpression: phi [0: t0, 7: t5] #i
interpretDynamically: t6 == 7:int
interpretDynamically: if t4 goto 9 else 10
resolveExpression: t6 == 7:int
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: 7:int
interpretDynamically: t6 + 1:int
interpretDynamically: jump 8
interpretDynamically: phi [5: t2, 10: t9] #j
resolveExpression: t6 + 1:int
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: 1:int
interpretDynamically: x + t2
interpretDynamically: t6 < t7
interpretDynamically: t2 + 1:int
interpretDynamically: jump 3
interpretDynamically: phi [0: t0, 7: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: 1:int
interpretDynamically: if t8 goto 6 else 7
resolveExpression: t6 < t7
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: x + t2
resolveExpression: parameter x : int
resolveExpression: phi [0: t0, 7: t5] #i
interpretDynamically: t6 == 7:int
interpretDynamically: if t4 goto 9 else 10
resolveExpression: t6 == 7:int
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: 7:int
interpretDynamically: t2 + 1:int
interpretDynamically: jump 3
interpretDynamically: phi [0: t0, 7: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: 1:int
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: t6 + 1:int
interpretDynamically: jump 8
interpretDynamically: phi [5: t2, 10: t9] #j
resolveExpression: t6 + 1:int
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: 1:int
interpretDynamically: x + t2
interpretDynamically: t6 < t7
interpretDynamically: t2 < x
interpretDynamically: return -1:int
resolveExpression: -1:int
interpretDynamically: t2 < x
interpretDynamically: if t3 goto 1 else 2
resolveExpression: t2 < x
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: parameter x : int
interpretDynamically: t2 < 0:int
interpretDynamically: if t1 goto 4 else 5
resolveExpression: t2 < 0:int
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: 0:int
interpretDynamically: jump 8
interpretDynamically: phi [5: t2, 10: t9] #j
resolveExpression: phi [0: t0, 7: t5] #i
interpretDynamically: return 2:int
resolveExpression: 2:int
interpretDynamically: x + t2
interpretDynamically: t6 < t7
interpretDynamically: if t8 goto 6 else 7
resolveExpression: t6 < t7
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: x + t2
resolveExpression: parameter x : int
resolveExpression: phi [0: t0, 7: t5] #i
interpretDynamically: t6 == 7:int
interpretDynamically: if t4 goto 9 else 10
resolveExpression: t6 == 7:int
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: 7:int
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: t6 + 1:int
interpretDynamically: jump 8
interpretDynamically: phi [5: t2, 10: t9] #j
resolveExpression: t6 + 1:int
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: 1:int
interpretDynamically: x + t2
interpretDynamically: t6 < t7
interpretDynamically: if t8 goto 6 else 7
resolveExpression: t6 < t7
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: x + t2
resolveExpression: parameter x : int
resolveExpression: phi [0: t0, 7: t5] #i
interpretDynamically: t2 + 1:int
interpretDynamically: jump 3
interpretDynamically: t6 == 7:int
interpretDynamically: phi [0: t0, 7: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: 1:int
interpretDynamically: if t4 goto 9 else 10
resolveExpression: t6 == 7:int
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: 7:int
interpretDynamically: t6 + 1:int
interpretDynamically: jump 8
interpretDynamically: phi [5: t2, 10: t9] #j
resolveExpression: t6 + 1:int
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: 1:int
interpretDynamically: x + t2
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: t6 < t7
interpretDynamically: if t8 goto 6 else 7
resolveExpression: t6 < t7
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: x + t2
resolveExpression: parameter x : int
resolveExpression: phi [0: t0, 7: t5] #i
interpretDynamically: t6 == 7:int
interpretDynamically: t2 + 1:int
interpretDynamically: jump 3
interpretDynamically: t2 + 1:int
interpretDynamically: jump 3
interpretDynamically: phi [0: t0, 7: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: 1:int
interpretDynamically: t2 < x
interpretDynamically: if t3 goto 1 else 2
resolveExpression: t2 < x
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: parameter x : int
interpretDynamically: return -1:int
resolveExpression: -1:int
interpretDynamically: t2 < 0:int
interpretDynamically: if t1 goto 4 else 5
resolveExpression: t2 < 0:int
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: 0:int
interpretDynamically: jump 8
interpretDynamically: return 2:int
resolveExpression: 2:int
interpretDynamically: if t4 goto 9 else 10
resolveExpression: t6 == 7:int
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: 7:int
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: t6 + 1:int
interpretDynamically: jump 8
interpretDynamically: if t3 goto 1 else 2
resolveExpression: t2 < x
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: parameter x : int
interpretDynamically: t2 < 0:int
interpretDynamically: return -1:int
resolveExpression: -1:int
interpretDynamically: if t1 goto 4 else 5
resolveExpression: t2 < 0:int
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: 0:int
interpretDynamically: jump 8
interpretDynamically: phi [5: t2, 10: t9] #j
resolveExpression: phi [0: t0, 7: t5] #i
interpretDynamically: x + t2
interpretDynamically: t6 < t7
interpretDynamically: if t8 goto 6 else 7
resolveExpression: t6 < t7
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: x + t2
resolveExpression: parameter x : int
resolveExpression: phi [0: t0, 7: t5] #i
interpretDynamically: t6 == 7:int
interpretDynamically: if t4 goto 9 else 10
resolveExpression: t6 == 7:int
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: 7:int
interpretDynamically: t2 + 1:int
interpretDynamically: t6 + 1:int
interpretDynamically: jump 3
interpretDynamically: phi [0: t0, 7: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: 1:int
interpretDynamically: jump 8
interpretDynamically: phi [5: t2, 10: t9] #j
resolveExpression: t6 + 1:int
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: 1:int
interpretDynamically: x + t2
interpretDynamically: t6 < t7
interpretDynamically: if t8 goto 6 else 7
resolveExpression: t6 < t7
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: x + t2
resolveExpression: parameter x : int
resolveExpression: phi [0: t0, 7: t5] #i
interpretDynamically: t2 + 1:int
interpretDynamically: jump 3
interpretDynamically: phi [0: t0, 7: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: 1:int
interpretDynamically: t2 < x
interpretDynamically: if t3 goto 1 else 2
resolveExpression: t2 < x
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: parameter x : int
interpretDynamically: t2 < 0:int
interpretDynamically: if t1 goto 4 else 5
resolveExpression: t2 < 0:int
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: 0:int
interpretDynamically: jump 8
interpretDynamically: phi [5: t2, 10: t9] #j
resolveExpression: phi [0: t0, 7: t5] #i
interpretDynamically: x + t2
interpretDynamically: return 2:int
resolveExpression: 2:int
interpretDynamically: return -1:int
resolveExpression: -1:int
interpretDynamically: t6 < t7
interpretDynamically: if t8 goto 6 else 7
resolveExpression: t6 < t7
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: x + t2
resolveExpression: parameter x : int
resolveExpression: phi [0: t0, 7: t5] #i
interpretDynamically: t6 == 7:int
interpretDynamically: if t4 goto 9 else 10
resolveExpression: t6 == 7:int
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: 7:int
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: t6 + 1:int
interpretDynamically: jump 8
interpretDynamically: t2 + 1:int
interpretDynamically: jump 3
interpretDynamically: return 2:int
resolveExpression: 2:int
interpretDynamically: t6 == 7:int
interpretDynamically: if t4 goto 9 else 10
resolveExpression: t6 == 7:int
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: 7:int
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: phi [0: t0, 7: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: 1:int
interpretDynamically: t2 < x
interpretDynamically: if t3 goto 1 else 2
resolveExpression: t2 < x
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: parameter x : int
interpretDynamically: t2 < 0:int
interpretDynamically: return -1:int
resolveExpression: -1:int
interpretDynamically: t6 + 1:int
interpretDynamically: jump 8
interpretDynamically: phi [5: t2, 10: t9] #j
resolveExpression: t6 + 1:int
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: 1:int
interpretDynamically: x + t2
interpretDynamically: t6 < t7
interpretDynamically: if t8 goto 6 else 7
resolveExpression: t6 < t7
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: x + t2
resolveExpression: parameter x : int
resolveExpression: phi [0: t0, 7: t5] #i
interpretDynamically: t6 == 7:int
interpretDynamically: if t4 goto 9 else 10
resolveExpression: t6 == 7:int
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: 7:int
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: t6 + 1:int
interpretDynamically: jump 8
interpretDynamically: phi [5: t2, 10: t9] #j
resolveExpression: t6 + 1:int
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: 1:int
interpretDynamically: t2 + 1:int
interpretDynamically: jump 3
interpretDynamically: phi [0: t0, 7: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: 1:int
interpretDynamically: t2 < x
interpretDynamically: if t3 goto 1 else 2
resolveExpression: t2 < x
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: parameter x : int
interpretDynamically: return -1:int
resolveExpression: -1:int
interpretDynamically: t2 < 0:int
interpretDynamically: if t1 goto 4 else 5
resolveExpression: t2 < 0:int
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: 0:int
interpretDynamically: return 2:int
resolveExpression: 2:int
interpretDynamically: jump 8
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: t2 < x
interpretDynamically: if t3 goto 1 else 2
resolveExpression: t2 < x
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: parameter x : int
interpretDynamically: return -1:int
resolveExpression: -1:int
interpretDynamically: t2 < 0:int
interpretDynamically: if t1 goto 4 else 5
resolveExpression: t2 < 0:int
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: 0:int
interpretDynamically: return 2:int
resolveExpression: 2:int
interpretDynamically: jump 8
interpretDynamically: phi [5: t2, 10: t9] #j
resolveExpression: phi [0: t0, 7: t5] #i
interpretDynamically: x + t2
interpretDynamically: t6 < t7
interpretDynamically: if t8 goto 6 else 7
resolveExpression: t6 < t7
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: x + t2
resolveExpression: parameter x : int
resolveExpression: phi [0: t0, 7: t5] #i
interpretDynamically: t6 == 7:int
interpretDynamically: if t4 goto 9 else 10
resolveExpression: t6 == 7:int
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: 7:int
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: t6 + 1:int
interpretDynamically: t2 + 1:int
interpretDynamically: jump 3
interpretDynamically: jump 8
interpretDynamically: phi [5: t2, 10: t9] #j
resolveExpression: t6 + 1:int
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: 1:int
interpretDynamically: x + t2
interpretDynamically: t6 < t7
interpretDynamically: if t8 goto 6 else 7
resolveExpression: t6 < t7
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: x + t2
resolveExpression: parameter x : int
resolveExpression: phi [0: t0, 7: t5] #i
interpretDynamically: t2 + 1:int
interpretDynamically: jump 3
interpretDynamically: phi [0: t0, 7: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: 1:int
interpretDynamically: t6 == 7:int
interpretDynamically: if t4 goto 9 else 10
resolveExpression: t6 == 7:int
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: 7:int
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: t6 + 1:int
interpretDynamically: jump 8
interpretDynamically: t2 < x
interpretDynamically: if t3 goto 1 else 2
resolveExpression: t2 < x
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: parameter x : int
interpretDynamically: return -1:int
resolveExpression: -1:int
interpretDynamically: t2 < 0:int
interpretDynamically: if t1 goto 4 else 5
resolveExpression: t2 < 0:int
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: 0:int
interpretDynamically: jump 8
interpretDynamically: return 2:int
resolveExpression: 2:int
interpretDynamically: phi [0: t0, 7: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: 1:int
interpretDynamically: t2 < x
interpretDynamically: if t3 goto 1 else 2
resolveExpression: t2 < x
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: parameter x : int
interpretDynamically: return -1:int
resolveExpression: -1:int
interpretDynamically: t2 < 0:int
interpretDynamically: if t1 goto 4 else 5
resolveExpression: t2 < 0:int
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: 0:int
interpretDynamically: jump 8
interpretDynamically: phi [5: t2, 10: t9] #j
resolveExpression: phi [0: t0, 7: t5] #i
interpretDynamically: x + t2
interpretDynamically: t6 < t7
interpretDynamically: if t8 goto 6 else 7
resolveExpression: t6 < t7
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: x + t2
resolveExpression: parameter x : int
resolveExpression: phi [0: t0, 7: t5] #i
interpretDynamically: t2 + 1:int
interpretDynamically: jump 3
interpretDynamically: phi [0: t0, 7: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: 1:int
interpretDynamically: t2 < x
interpretDynamically: t6 == 7:int
interpretDynamically: if t4 goto 9 else 10
resolveExpression: t6 == 7:int
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: 7:int
interpretDynamically: t6 + 1:int
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: jump 8
interpretDynamically: if t3 goto 1 else 2
resolveExpression: t2 < x
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: parameter x : int
interpretDynamically: t2 < 0:int
interpretDynamically: if t1 goto 4 else 5
resolveExpression: t2 < 0:int
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: 0:int
interpretDynamically: jump 8
interpretDynamically: return 2:int
resolveExpression: 2:int
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: return 2:int
resolveExpression: 2:int
interpretDynamically: t2 < x
interpretDynamically: if t3 goto 1 else 2
resolveExpression: t2 < x
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: parameter x : int
interpretDynamically: return -1:int
resolveExpression: -1:int
interpretDynamically: t2 < 0:int
interpretDynamically: if t1 goto 4 else 5
resolveExpression: t2 < 0:int
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: 0:int
interpretDynamically: return 2:int
resolveExpression: 2:int
interpretDynamically: jump 8
interpretDynamically: phi [5: t2, 10: t9] #j
resolveExpression: phi [0: t0, 7: t5] #i
interpretDynamically: x + t2
interpretDynamically: t6 < t7
interpretDynamically: if t8 goto 6 else 7
resolveExpression: t6 < t7
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: x + t2
resolveExpression: parameter x : int
resolveExpression: phi [0: t0, 7: t5] #i
interpretDynamically: t2 + 1:int
interpretDynamically: jump 3
interpretDynamically: phi [0: t0, 7: t5] #i
resolveExpression: t2 + 1:int
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: 1:int
interpretDynamically: t2 < x
interpretDynamically: if t3 goto 1 else 2
resolveExpression: t2 < x
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: parameter x : int
interpretDynamically: return -1:int
resolveExpression: -1:int
interpretDynamically: t6 == 7:int
interpretDynamically: if t4 goto 9 else 10
resolveExpression: t6 == 7:int
resolveExpression: phi [5: t2, 10: t9] #j
resolveExpression: 7:int
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: t6 + 1:int
interpretDynamically: jump 8
interpretDynamically: t2 < 0:int
interpretDynamically: if t1 goto 4 else 5
resolveExpression: t2 < 0:int
resolveExpression: phi [0: t0, 7: t5] #i
resolveExpression: 0:int
Analyse stopped: context deadline exceeded
States: 34
---- State #0 ----
Interpreter{
  PC: ((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & ((IntSegment_sel0 - 5) < 0))
  ReturnResults:
    [0] 2
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
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x8000020000000004)
addr_0 -> #x0000000000000000
IntSegment_sel1 -> #x8000020000000004
IntSegment_sel0 -> #x8000020000000004

---- State #1 ----
Interpreter{
  PC: ((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & ((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5)))) & (!((IntSegment_sel0 - 5) == 7))) & (((IntSegment_sel0 - 5) + 1) < (IntSegment_sel3 + (IntSegment_sel0 - 5)))) & (((IntSegment_sel0 - 5) + 1) == 7))
  ReturnResults:
    [0] 1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x000000000000000b)
IntSegment_sel2 -> #x000000000000000b
IntSegment_sel1 -> #x000000000000000b
IntSegment_sel3 -> #x000000000000000b
IntSegment_sel0 -> #x000000000000000b

---- State #2 ----
Interpreter{
  PC: (true & (!((IntSegment_sel0 - 5) < IntSegment_sel1)))
  ReturnResults:
    [0] -1
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
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x8000000000000004)
addr_0 -> #x0000000000000000
IntSegment_sel1 -> #x8000000000000004
IntSegment_sel0 -> #x8000000000000004

---- State #3 ----
Interpreter{
  PC: (((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & ((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5)))) & (!((IntSegment_sel0 - 5) == 7))) & (!(((IntSegment_sel0 - 5) + 1) < (IntSegment_sel3 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel4)) & (((IntSegment_sel0 - 5) + 1) < 0))
  ReturnResults:
    [0] 2
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=5, store=0, select=5)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #4 ----
Interpreter{
  PC: (((((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & ((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5)))) & (!((IntSegment_sel0 - 5) == 7))) & (!(((IntSegment_sel0 - 5) + 1) < (IntSegment_sel3 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel4)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (((IntSegment_sel0 - 5) + 1) < (IntSegment_sel5 + ((IntSegment_sel0 - 5) + 1)))) & (!(((IntSegment_sel0 - 5) + 1) == 7))) & ((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel6 + ((IntSegment_sel0 - 5) + 1)))) & ((((IntSegment_sel0 - 5) + 1) + 1) == 7))
  ReturnResults:
    [0] 1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=7, store=0, select=7)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=addr_0, res=IntSegment_sel6
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5).select(addr_0, IntSegment_sel6)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #5 ----
Interpreter{
  PC: (((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & ((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5)))) & (!((IntSegment_sel0 - 5) == 7))) & (!(((IntSegment_sel0 - 5) + 1) < (IntSegment_sel3 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel4)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (((IntSegment_sel0 - 5) + 1) < (IntSegment_sel5 + ((IntSegment_sel0 - 5) + 1)))) & (((IntSegment_sel0 - 5) + 1) == 7))
  ReturnResults:
    [0] 1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=6, store=0, select=6)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #6 ----
Interpreter{
  PC: (((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & ((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5)))) & (!((IntSegment_sel0 - 5) == 7))) & (!(((IntSegment_sel0 - 5) + 1) < (IntSegment_sel3 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel4)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (!(((IntSegment_sel0 - 5) + 1) < (IntSegment_sel5 + ((IntSegment_sel0 - 5) + 1))))) & (!((((IntSegment_sel0 - 5) + 1) + 1) < IntSegment_sel6)))
  ReturnResults:
    [0] -1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=7, store=0, select=7)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=addr_0, res=IntSegment_sel6
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5).select(addr_0, IntSegment_sel6)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #7 ----
Interpreter{
  PC: ((((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & ((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5)))) & (!((IntSegment_sel0 - 5) == 7))) & (!(((IntSegment_sel0 - 5) + 1) < (IntSegment_sel3 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel4)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (!(((IntSegment_sel0 - 5) + 1) < (IntSegment_sel5 + ((IntSegment_sel0 - 5) + 1))))) & ((((IntSegment_sel0 - 5) + 1) + 1) < IntSegment_sel6)) & ((((IntSegment_sel0 - 5) + 1) + 1) < 0))
  ReturnResults:
    [0] 2
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=7, store=0, select=7)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=addr_0, res=IntSegment_sel6
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5).select(addr_0, IntSegment_sel6)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #8 ----
Interpreter{
  PC: (((((((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & ((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5)))) & (!((IntSegment_sel0 - 5) == 7))) & (!(((IntSegment_sel0 - 5) + 1) < (IntSegment_sel3 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel4)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (((IntSegment_sel0 - 5) + 1) < (IntSegment_sel5 + ((IntSegment_sel0 - 5) + 1)))) & (!(((IntSegment_sel0 - 5) + 1) == 7))) & ((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel6 + ((IntSegment_sel0 - 5) + 1)))) & (!((((IntSegment_sel0 - 5) + 1) + 1) == 7))) & (((((IntSegment_sel0 - 5) + 1) + 1) + 1) < (IntSegment_sel7 + ((IntSegment_sel0 - 5) + 1)))) & (((((IntSegment_sel0 - 5) + 1) + 1) + 1) == 7))
  ReturnResults:
    [0] 1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=8, store=0, select=8)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=addr_0, res=IntSegment_sel6
        op[7]: select idx=addr_0, res=IntSegment_sel7
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5).select(addr_0, IntSegment_sel6).select(addr_0, IntSegment_sel7)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #9 ----
Interpreter{
  PC: ((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & (!((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel3)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (!(((IntSegment_sel0 - 5) + 1) < (IntSegment_sel4 + ((IntSegment_sel0 - 5) + 1))))) & ((((IntSegment_sel0 - 5) + 1) + 1) < IntSegment_sel5)) & ((((IntSegment_sel0 - 5) + 1) + 1) < 0))
  ReturnResults:
    [0] 2
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=6, store=0, select=6)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #10 ----
Interpreter{
  PC: (((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & (!((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel3)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (!(((IntSegment_sel0 - 5) + 1) < (IntSegment_sel4 + ((IntSegment_sel0 - 5) + 1))))) & (!((((IntSegment_sel0 - 5) + 1) + 1) < IntSegment_sel5)))
  ReturnResults:
    [0] -1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=6, store=0, select=6)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #11 ----
Interpreter{
  PC: ((((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & (!((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel3)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (!(((IntSegment_sel0 - 5) + 1) < (IntSegment_sel4 + ((IntSegment_sel0 - 5) + 1))))) & ((((IntSegment_sel0 - 5) + 1) + 1) < IntSegment_sel5)) & (!((((IntSegment_sel0 - 5) + 1) + 1) < 0))) & ((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel6 + (((IntSegment_sel0 - 5) + 1) + 1)))) & ((((IntSegment_sel0 - 5) + 1) + 1) == 7))
  ReturnResults:
    [0] 1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=7, store=0, select=7)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=addr_0, res=IntSegment_sel6
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5).select(addr_0, IntSegment_sel6)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #12 ----
Interpreter{
  PC: (((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & (!((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel3)) & (((IntSegment_sel0 - 5) + 1) < 0))
  ReturnResults:
    [0] 2
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #13 ----
Interpreter{
  PC: (((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & (!((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel3)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (((IntSegment_sel0 - 5) + 1) < (IntSegment_sel4 + ((IntSegment_sel0 - 5) + 1)))) & (!(((IntSegment_sel0 - 5) + 1) == 7))) & ((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel5 + ((IntSegment_sel0 - 5) + 1)))) & ((((IntSegment_sel0 - 5) + 1) + 1) == 7))
  ReturnResults:
    [0] 1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=6, store=0, select=6)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #14 ----
Interpreter{
  PC: ((((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & (!((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel3)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (!(((IntSegment_sel0 - 5) + 1) < (IntSegment_sel4 + ((IntSegment_sel0 - 5) + 1))))) & ((((IntSegment_sel0 - 5) + 1) + 1) < IntSegment_sel5)) & (!((((IntSegment_sel0 - 5) + 1) + 1) < 0))) & (!((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel6 + (((IntSegment_sel0 - 5) + 1) + 1))))) & (!(((((IntSegment_sel0 - 5) + 1) + 1) + 1) < IntSegment_sel7)))
  ReturnResults:
    [0] -1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=8, store=0, select=8)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=addr_0, res=IntSegment_sel6
        op[7]: select idx=addr_0, res=IntSegment_sel7
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5).select(addr_0, IntSegment_sel6).select(addr_0, IntSegment_sel7)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #15 ----
Interpreter{
  PC: (((((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & (!((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel3)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (((IntSegment_sel0 - 5) + 1) < (IntSegment_sel4 + ((IntSegment_sel0 - 5) + 1)))) & (!(((IntSegment_sel0 - 5) + 1) == 7))) & ((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel5 + ((IntSegment_sel0 - 5) + 1)))) & (!((((IntSegment_sel0 - 5) + 1) + 1) == 7))) & (((((IntSegment_sel0 - 5) + 1) + 1) + 1) < (IntSegment_sel6 + ((IntSegment_sel0 - 5) + 1)))) & (((((IntSegment_sel0 - 5) + 1) + 1) + 1) == 7))
  ReturnResults:
    [0] 1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=7, store=0, select=7)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=addr_0, res=IntSegment_sel6
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5).select(addr_0, IntSegment_sel6)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #16 ----
Interpreter{
  PC: ((((((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & (!((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel3)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (((IntSegment_sel0 - 5) + 1) < (IntSegment_sel4 + ((IntSegment_sel0 - 5) + 1)))) & (!(((IntSegment_sel0 - 5) + 1) == 7))) & ((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel5 + ((IntSegment_sel0 - 5) + 1)))) & (!((((IntSegment_sel0 - 5) + 1) + 1) == 7))) & (!(((((IntSegment_sel0 - 5) + 1) + 1) + 1) < (IntSegment_sel6 + ((IntSegment_sel0 - 5) + 1))))) & ((((IntSegment_sel0 - 5) + 1) + 1) < IntSegment_sel7)) & ((((IntSegment_sel0 - 5) + 1) + 1) < 0))
  ReturnResults:
    [0] 2
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=8, store=0, select=8)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=addr_0, res=IntSegment_sel6
        op[7]: select idx=addr_0, res=IntSegment_sel7
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5).select(addr_0, IntSegment_sel6).select(addr_0, IntSegment_sel7)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #17 ----
Interpreter{
  PC: ((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & ((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5)))) & ((IntSegment_sel0 - 5) == 7))
  ReturnResults:
    [0] 1
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
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x000000000000000c)
IntSegment_sel2 -> #x000000000000000c
IntSegment_sel1 -> #x000000000000000c
IntSegment_sel0 -> #x000000000000000c

---- State #18 ----
Interpreter{
  PC: ((((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & (!((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel3)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (((IntSegment_sel0 - 5) + 1) < (IntSegment_sel4 + ((IntSegment_sel0 - 5) + 1)))) & (!(((IntSegment_sel0 - 5) + 1) == 7))) & (!((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel5 + ((IntSegment_sel0 - 5) + 1))))) & ((((IntSegment_sel0 - 5) + 1) + 1) < IntSegment_sel6)) & ((((IntSegment_sel0 - 5) + 1) + 1) < 0))
  ReturnResults:
    [0] 2
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=7, store=0, select=7)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=addr_0, res=IntSegment_sel6
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5).select(addr_0, IntSegment_sel6)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #19 ----
Interpreter{
  PC: (((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & (!((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel3)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (((IntSegment_sel0 - 5) + 1) < (IntSegment_sel4 + ((IntSegment_sel0 - 5) + 1)))) & (!(((IntSegment_sel0 - 5) + 1) == 7))) & (!((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel5 + ((IntSegment_sel0 - 5) + 1))))) & (!((((IntSegment_sel0 - 5) + 1) + 1) < IntSegment_sel6)))
  ReturnResults:
    [0] -1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=7, store=0, select=7)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=addr_0, res=IntSegment_sel6
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5).select(addr_0, IntSegment_sel6)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #20 ----
Interpreter{
  PC: ((((((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & (!((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel3)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (((IntSegment_sel0 - 5) + 1) < (IntSegment_sel4 + ((IntSegment_sel0 - 5) + 1)))) & (!(((IntSegment_sel0 - 5) + 1) == 7))) & (!((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel5 + ((IntSegment_sel0 - 5) + 1))))) & ((((IntSegment_sel0 - 5) + 1) + 1) < IntSegment_sel6)) & (!((((IntSegment_sel0 - 5) + 1) + 1) < 0))) & ((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel7 + (((IntSegment_sel0 - 5) + 1) + 1)))) & ((((IntSegment_sel0 - 5) + 1) + 1) == 7))
  ReturnResults:
    [0] 1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=8, store=0, select=8)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=addr_0, res=IntSegment_sel6
        op[7]: select idx=addr_0, res=IntSegment_sel7
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5).select(addr_0, IntSegment_sel6).select(addr_0, IntSegment_sel7)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #21 ----
Interpreter{
  PC: (((((((((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & (!((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel3)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (((IntSegment_sel0 - 5) + 1) < (IntSegment_sel4 + ((IntSegment_sel0 - 5) + 1)))) & (!(((IntSegment_sel0 - 5) + 1) == 7))) & (!((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel5 + ((IntSegment_sel0 - 5) + 1))))) & ((((IntSegment_sel0 - 5) + 1) + 1) < IntSegment_sel6)) & (!((((IntSegment_sel0 - 5) + 1) + 1) < 0))) & ((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel7 + (((IntSegment_sel0 - 5) + 1) + 1)))) & (!((((IntSegment_sel0 - 5) + 1) + 1) == 7))) & (!(((((IntSegment_sel0 - 5) + 1) + 1) + 1) < (IntSegment_sel8 + (((IntSegment_sel0 - 5) + 1) + 1))))) & (((((IntSegment_sel0 - 5) + 1) + 1) + 1) < IntSegment_sel9)) & (((((IntSegment_sel0 - 5) + 1) + 1) + 1) < 0))
  ReturnResults:
    [0] 2
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=10, store=0, select=10)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=addr_0, res=IntSegment_sel6
        op[7]: select idx=addr_0, res=IntSegment_sel7
        op[8]: select idx=addr_0, res=IntSegment_sel8
        op[9]: select idx=addr_0, res=IntSegment_sel9
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5).select(addr_0, IntSegment_sel6).select(addr_0, IntSegment_sel7).select(addr_0, IntSegment_sel8).select(addr_0, IntSegment_sel9)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #22 ----
Interpreter{
  PC: ((((((((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & (!((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel3)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (((IntSegment_sel0 - 5) + 1) < (IntSegment_sel4 + ((IntSegment_sel0 - 5) + 1)))) & (!(((IntSegment_sel0 - 5) + 1) == 7))) & (!((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel5 + ((IntSegment_sel0 - 5) + 1))))) & ((((IntSegment_sel0 - 5) + 1) + 1) < IntSegment_sel6)) & (!((((IntSegment_sel0 - 5) + 1) + 1) < 0))) & ((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel7 + (((IntSegment_sel0 - 5) + 1) + 1)))) & (!((((IntSegment_sel0 - 5) + 1) + 1) == 7))) & (!(((((IntSegment_sel0 - 5) + 1) + 1) + 1) < (IntSegment_sel8 + (((IntSegment_sel0 - 5) + 1) + 1))))) & (!(((((IntSegment_sel0 - 5) + 1) + 1) + 1) < IntSegment_sel9)))
  ReturnResults:
    [0] -1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=10, store=0, select=10)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=addr_0, res=IntSegment_sel6
        op[7]: select idx=addr_0, res=IntSegment_sel7
        op[8]: select idx=addr_0, res=IntSegment_sel8
        op[9]: select idx=addr_0, res=IntSegment_sel9
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5).select(addr_0, IntSegment_sel6).select(addr_0, IntSegment_sel7).select(addr_0, IntSegment_sel8).select(addr_0, IntSegment_sel9)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #23 ----
Interpreter{
  PC: ((((((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & (!((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel3)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (((IntSegment_sel0 - 5) + 1) < (IntSegment_sel4 + ((IntSegment_sel0 - 5) + 1)))) & (!(((IntSegment_sel0 - 5) + 1) == 7))) & (!((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel5 + ((IntSegment_sel0 - 5) + 1))))) & ((((IntSegment_sel0 - 5) + 1) + 1) < IntSegment_sel6)) & (!((((IntSegment_sel0 - 5) + 1) + 1) < 0))) & (!((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel7 + (((IntSegment_sel0 - 5) + 1) + 1))))) & (!(((((IntSegment_sel0 - 5) + 1) + 1) + 1) < IntSegment_sel8)))
  ReturnResults:
    [0] -1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=9, store=0, select=9)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=addr_0, res=IntSegment_sel6
        op[7]: select idx=addr_0, res=IntSegment_sel7
        op[8]: select idx=addr_0, res=IntSegment_sel8
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5).select(addr_0, IntSegment_sel6).select(addr_0, IntSegment_sel7).select(addr_0, IntSegment_sel8)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #24 ----
Interpreter{
  PC: (((((((((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & (!((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel3)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (((IntSegment_sel0 - 5) + 1) < (IntSegment_sel4 + ((IntSegment_sel0 - 5) + 1)))) & (!(((IntSegment_sel0 - 5) + 1) == 7))) & (!((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel5 + ((IntSegment_sel0 - 5) + 1))))) & ((((IntSegment_sel0 - 5) + 1) + 1) < IntSegment_sel6)) & (!((((IntSegment_sel0 - 5) + 1) + 1) < 0))) & (!((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel7 + (((IntSegment_sel0 - 5) + 1) + 1))))) & (((((IntSegment_sel0 - 5) + 1) + 1) + 1) < IntSegment_sel8)) & (!(((((IntSegment_sel0 - 5) + 1) + 1) + 1) < 0))) & (((((IntSegment_sel0 - 5) + 1) + 1) + 1) < (IntSegment_sel9 + ((((IntSegment_sel0 - 5) + 1) + 1) + 1)))) & (((((IntSegment_sel0 - 5) + 1) + 1) + 1) == 7))
  ReturnResults:
    [0] 1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=10, store=0, select=10)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=addr_0, res=IntSegment_sel6
        op[7]: select idx=addr_0, res=IntSegment_sel7
        op[8]: select idx=addr_0, res=IntSegment_sel8
        op[9]: select idx=addr_0, res=IntSegment_sel9
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5).select(addr_0, IntSegment_sel6).select(addr_0, IntSegment_sel7).select(addr_0, IntSegment_sel8).select(addr_0, IntSegment_sel9)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #25 ----
Interpreter{
  PC: ((((((((((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & (!((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel3)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (((IntSegment_sel0 - 5) + 1) < (IntSegment_sel4 + ((IntSegment_sel0 - 5) + 1)))) & (!(((IntSegment_sel0 - 5) + 1) == 7))) & (!((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel5 + ((IntSegment_sel0 - 5) + 1))))) & ((((IntSegment_sel0 - 5) + 1) + 1) < IntSegment_sel6)) & (!((((IntSegment_sel0 - 5) + 1) + 1) < 0))) & (!((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel7 + (((IntSegment_sel0 - 5) + 1) + 1))))) & (((((IntSegment_sel0 - 5) + 1) + 1) + 1) < IntSegment_sel8)) & (!(((((IntSegment_sel0 - 5) + 1) + 1) + 1) < 0))) & (!(((((IntSegment_sel0 - 5) + 1) + 1) + 1) < (IntSegment_sel9 + ((((IntSegment_sel0 - 5) + 1) + 1) + 1))))) & ((((((IntSegment_sel0 - 5) + 1) + 1) + 1) + 1) < IntSegment_sel10)) & ((((((IntSegment_sel0 - 5) + 1) + 1) + 1) + 1) < 0))
  ReturnResults:
    [0] 2
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=11, store=0, select=11)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=addr_0, res=IntSegment_sel6
        op[7]: select idx=addr_0, res=IntSegment_sel7
        op[8]: select idx=addr_0, res=IntSegment_sel8
        op[9]: select idx=addr_0, res=IntSegment_sel9
        op[10]: select idx=addr_0, res=IntSegment_sel10
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5).select(addr_0, IntSegment_sel6).select(addr_0, IntSegment_sel7).select(addr_0, IntSegment_sel8).select(addr_0, IntSegment_sel9).select(addr_0, IntSegment_sel10)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #26 ----
Interpreter{
  PC: ((((((((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & (!((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel3)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (((IntSegment_sel0 - 5) + 1) < (IntSegment_sel4 + ((IntSegment_sel0 - 5) + 1)))) & (!(((IntSegment_sel0 - 5) + 1) == 7))) & (!((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel5 + ((IntSegment_sel0 - 5) + 1))))) & ((((IntSegment_sel0 - 5) + 1) + 1) < IntSegment_sel6)) & (!((((IntSegment_sel0 - 5) + 1) + 1) < 0))) & ((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel7 + (((IntSegment_sel0 - 5) + 1) + 1)))) & (!((((IntSegment_sel0 - 5) + 1) + 1) == 7))) & (((((IntSegment_sel0 - 5) + 1) + 1) + 1) < (IntSegment_sel8 + (((IntSegment_sel0 - 5) + 1) + 1)))) & (((((IntSegment_sel0 - 5) + 1) + 1) + 1) == 7))
  ReturnResults:
    [0] 1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=9, store=0, select=9)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=addr_0, res=IntSegment_sel6
        op[7]: select idx=addr_0, res=IntSegment_sel7
        op[8]: select idx=addr_0, res=IntSegment_sel8
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5).select(addr_0, IntSegment_sel6).select(addr_0, IntSegment_sel7).select(addr_0, IntSegment_sel8)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #27 ----
Interpreter{
  PC: (((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & (!((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel3)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (((IntSegment_sel0 - 5) + 1) < (IntSegment_sel4 + ((IntSegment_sel0 - 5) + 1)))) & (((IntSegment_sel0 - 5) + 1) == 7))
  ReturnResults:
    [0] 1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=5, store=0, select=5)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #28 ----
Interpreter{
  PC: (((((((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & (!((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel3)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (((IntSegment_sel0 - 5) + 1) < (IntSegment_sel4 + ((IntSegment_sel0 - 5) + 1)))) & (!(((IntSegment_sel0 - 5) + 1) == 7))) & (!((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel5 + ((IntSegment_sel0 - 5) + 1))))) & ((((IntSegment_sel0 - 5) + 1) + 1) < IntSegment_sel6)) & (!((((IntSegment_sel0 - 5) + 1) + 1) < 0))) & (!((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel7 + (((IntSegment_sel0 - 5) + 1) + 1))))) & (((((IntSegment_sel0 - 5) + 1) + 1) + 1) < IntSegment_sel8)) & (((((IntSegment_sel0 - 5) + 1) + 1) + 1) < 0))
  ReturnResults:
    [0] 2
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=9, store=0, select=9)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=addr_0, res=IntSegment_sel6
        op[7]: select idx=addr_0, res=IntSegment_sel7
        op[8]: select idx=addr_0, res=IntSegment_sel8
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5).select(addr_0, IntSegment_sel6).select(addr_0, IntSegment_sel7).select(addr_0, IntSegment_sel8)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #29 ----
Interpreter{
  PC: (((((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & ((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5)))) & (!((IntSegment_sel0 - 5) == 7))) & (!(((IntSegment_sel0 - 5) + 1) < (IntSegment_sel3 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel4)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (((IntSegment_sel0 - 5) + 1) < (IntSegment_sel5 + ((IntSegment_sel0 - 5) + 1)))) & (!(((IntSegment_sel0 - 5) + 1) == 7))) & (!((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel6 + ((IntSegment_sel0 - 5) + 1))))) & (!((((IntSegment_sel0 - 5) + 1) + 1) < IntSegment_sel7)))
  ReturnResults:
    [0] -1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=8, store=0, select=8)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=addr_0, res=IntSegment_sel6
        op[7]: select idx=addr_0, res=IntSegment_sel7
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5).select(addr_0, IntSegment_sel6).select(addr_0, IntSegment_sel7)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #30 ----
Interpreter{
  PC: (((((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & (!((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel3)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (((IntSegment_sel0 - 5) + 1) < (IntSegment_sel4 + ((IntSegment_sel0 - 5) + 1)))) & (!(((IntSegment_sel0 - 5) + 1) == 7))) & ((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel5 + ((IntSegment_sel0 - 5) + 1)))) & (!((((IntSegment_sel0 - 5) + 1) + 1) == 7))) & (!(((((IntSegment_sel0 - 5) + 1) + 1) + 1) < (IntSegment_sel6 + ((IntSegment_sel0 - 5) + 1))))) & (!((((IntSegment_sel0 - 5) + 1) + 1) < IntSegment_sel7)))
  ReturnResults:
    [0] -1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=8, store=0, select=8)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=addr_0, res=IntSegment_sel6
        op[7]: select idx=addr_0, res=IntSegment_sel7
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5).select(addr_0, IntSegment_sel6).select(addr_0, IntSegment_sel7)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #31 ----
Interpreter{
  PC: ((((((((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & ((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5)))) & (!((IntSegment_sel0 - 5) == 7))) & (!(((IntSegment_sel0 - 5) + 1) < (IntSegment_sel3 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel4)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (((IntSegment_sel0 - 5) + 1) < (IntSegment_sel5 + ((IntSegment_sel0 - 5) + 1)))) & (!(((IntSegment_sel0 - 5) + 1) == 7))) & (!((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel6 + ((IntSegment_sel0 - 5) + 1))))) & ((((IntSegment_sel0 - 5) + 1) + 1) < IntSegment_sel7)) & (!((((IntSegment_sel0 - 5) + 1) + 1) < 0))) & (!((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel8 + (((IntSegment_sel0 - 5) + 1) + 1))))) & (!(((((IntSegment_sel0 - 5) + 1) + 1) + 1) < IntSegment_sel9)))
  ReturnResults:
    [0] -1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=10, store=0, select=10)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=addr_0, res=IntSegment_sel6
        op[7]: select idx=addr_0, res=IntSegment_sel7
        op[8]: select idx=addr_0, res=IntSegment_sel8
        op[9]: select idx=addr_0, res=IntSegment_sel9
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5).select(addr_0, IntSegment_sel6).select(addr_0, IntSegment_sel7).select(addr_0, IntSegment_sel8).select(addr_0, IntSegment_sel9)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #32 ----
Interpreter{
  PC: ((((((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & ((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5)))) & (!((IntSegment_sel0 - 5) == 7))) & (!(((IntSegment_sel0 - 5) + 1) < (IntSegment_sel3 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel4)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (((IntSegment_sel0 - 5) + 1) < (IntSegment_sel5 + ((IntSegment_sel0 - 5) + 1)))) & (!(((IntSegment_sel0 - 5) + 1) == 7))) & (!((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel6 + ((IntSegment_sel0 - 5) + 1))))) & ((((IntSegment_sel0 - 5) + 1) + 1) < IntSegment_sel7)) & ((((IntSegment_sel0 - 5) + 1) + 1) < 0))
  ReturnResults:
    [0] 2
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=8, store=0, select=8)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=addr_0, res=IntSegment_sel6
        op[7]: select idx=addr_0, res=IntSegment_sel7
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5).select(addr_0, IntSegment_sel6).select(addr_0, IntSegment_sel7)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT
---- State #33 ----
Interpreter{
  PC: ((((((((((((((true & ((IntSegment_sel0 - 5) < IntSegment_sel1)) & (!((IntSegment_sel0 - 5) < 0))) & ((IntSegment_sel0 - 5) < (IntSegment_sel2 + (IntSegment_sel0 - 5)))) & (!((IntSegment_sel0 - 5) == 7))) & (!(((IntSegment_sel0 - 5) + 1) < (IntSegment_sel3 + (IntSegment_sel0 - 5))))) & (((IntSegment_sel0 - 5) + 1) < IntSegment_sel4)) & (!(((IntSegment_sel0 - 5) + 1) < 0))) & (((IntSegment_sel0 - 5) + 1) < (IntSegment_sel5 + ((IntSegment_sel0 - 5) + 1)))) & (!(((IntSegment_sel0 - 5) + 1) == 7))) & (!((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel6 + ((IntSegment_sel0 - 5) + 1))))) & ((((IntSegment_sel0 - 5) + 1) + 1) < IntSegment_sel7)) & (!((((IntSegment_sel0 - 5) + 1) + 1) < 0))) & ((((IntSegment_sel0 - 5) + 1) + 1) < (IntSegment_sel8 + (((IntSegment_sel0 - 5) + 1) + 1)))) & ((((IntSegment_sel0 - 5) + 1) + 1) == 7))
  ReturnResults:
    [0] 1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=9, store=0, select=9)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=addr_0, res=IntSegment_sel6
        op[7]: select idx=addr_0, res=IntSegment_sel7
        op[8]: select idx=addr_0, res=IntSegment_sel8
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).select(addr_0, IntSegment_sel5).select(addr_0, IntSegment_sel6).select(addr_0, IntSegment_sel7).select(addr_0, IntSegment_sel8)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

UNSAT

==================== Factorial ====================
CFG построен для функции с 5 блоками
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
interpretDynamically: n < 0:int
interpretDynamically: if t0 goto 1 else 2
resolveExpression: n < 0:int
resolveExpression: parameter n : int
resolveExpression: 0:int
interpretDynamically: return 0:int
resolveExpression: 0:int
interpretDynamically: return 0:int
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
interpretDynamically: return 0:int
resolveExpression: 0:int
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
interpretDynamically: n - 1:int
interpretDynamically: Factorial(t2)
resolveExpression: n - 1:int
resolveExpression: parameter n : int
resolveExpression: 1:int
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: return 0:int
resolveExpression: 0:int
States: 10
---- State #0 ----
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

---- State #1 ----
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
---- State #2 ----
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

---- State #3 ----
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

---- State #4 ----
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
---- State #5 ----
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
---- State #6 ----
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
---- State #7 ----
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

==================== NewObjectWithPrimitivesClass ====================
CFG построен для функции с 1 блоками
interpretDynamically: new ObjectWithPrimitivesClass (complit)
interpretDynamically: &t0.ValueByDefault [#0]
interpretDynamically: *t1 = 5:int
resolveExpression: 5:int
resolveExpression: new ObjectWithPrimitivesClass (complit)
interpretDynamically: return t0
resolveExpression: new ObjectWithPrimitivesClass (complit)
States: 1
---- State #0 ----
Interpreter{
  PC: true
  ReturnResults:
    [0] &(addr_0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 0, Allocations: 1
    AllocationsList:
      A[0]: &(addr_0) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=1, store=1, select=0)
        op[0]: store idx=addr_0, val=5
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=0, store=0, select=0)
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=0, store=0, select=0)
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=0, store=0, select=0)
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=0, store=0, select=0)
  }
  HeapConstraints: 10
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: ObjectWithPrimitivesClass_0.store(addr_0, 5)
    C[6]: ObjectWithPrimitivesClass_1
    C[7]: ObjectWithPrimitivesClass_2
    C[8]: ObjectWithPrimitivesClass_3
    C[9]: ObjectWithPrimitivesClass_4
}

SAT
Model:


==================== Max ====================
CFG построен для функции с 7 блоками
interpretDynamically: &fst.x [#1]
interpretDynamically: *t0
interpretDynamically: &snd.x [#1]
interpretDynamically: *t2
interpretDynamically: t1 > t3
interpretDynamically: if t4 goto 3 else 2
resolveExpression: t1 > t3
resolveExpression: *t0
resolveExpression: parameter fst : *ObjectWithPrimitivesClass
resolveExpression: *t2
resolveExpression: parameter snd : *ObjectWithPrimitivesClass
interpretDynamically: &fst.y [#2]
interpretDynamically: *t10
interpretDynamically: &snd.y [#2]
interpretDynamically: *t12
interpretDynamically: t11 > t13
interpretDynamically: if t14 goto 1 else 2
resolveExpression: t11 > t13
resolveExpression: *t10
resolveExpression: parameter fst : *ObjectWithPrimitivesClass
resolveExpression: *t12
resolveExpression: parameter snd : *ObjectWithPrimitivesClass
interpretDynamically: return fst
resolveExpression: parameter fst : *ObjectWithPrimitivesClass
interpretDynamically: &fst.x [#1]
interpretDynamically: *t5
interpretDynamically: &snd.x [#1]
interpretDynamically: *t7
interpretDynamically: t6 < t8
interpretDynamically: if t9 goto 6 else 5
resolveExpression: t6 < t8
resolveExpression: *t5
resolveExpression: parameter fst : *ObjectWithPrimitivesClass
resolveExpression: *t7
resolveExpression: parameter snd : *ObjectWithPrimitivesClass
interpretDynamically: &fst.y [#2]
interpretDynamically: return fst
resolveExpression: parameter fst : *ObjectWithPrimitivesClass
interpretDynamically: *t15
interpretDynamically: &snd.y [#2]
interpretDynamically: *t17
interpretDynamically: t16 < t18
interpretDynamically: if t19 goto 4 else 5
resolveExpression: t16 < t18
resolveExpression: *t15
resolveExpression: parameter fst : *ObjectWithPrimitivesClass
resolveExpression: *t17
resolveExpression: parameter snd : *ObjectWithPrimitivesClass
interpretDynamically: return snd
resolveExpression: parameter snd : *ObjectWithPrimitivesClass
interpretDynamically: return fst
resolveExpression: parameter fst : *ObjectWithPrimitivesClass
interpretDynamically: &fst.x [#1]
interpretDynamically: *t5
interpretDynamically: &snd.x [#1]
interpretDynamically: *t7
interpretDynamically: t6 < t8
interpretDynamically: if t9 goto 6 else 5
resolveExpression: t6 < t8
resolveExpression: *t5
resolveExpression: parameter fst : *ObjectWithPrimitivesClass
resolveExpression: *t7
resolveExpression: parameter snd : *ObjectWithPrimitivesClass
interpretDynamically: &fst.y [#2]
interpretDynamically: return fst
resolveExpression: parameter fst : *ObjectWithPrimitivesClass
interpretDynamically: *t15
interpretDynamically: &snd.y [#2]
interpretDynamically: *t17
interpretDynamically: t16 < t18
interpretDynamically: if t19 goto 4 else 5
resolveExpression: t16 < t18
resolveExpression: *t15
resolveExpression: parameter fst : *ObjectWithPrimitivesClass
resolveExpression: *t17
resolveExpression: parameter snd : *ObjectWithPrimitivesClass
interpretDynamically: return snd
resolveExpression: parameter snd : *ObjectWithPrimitivesClass
interpretDynamically: return fst
resolveExpression: parameter fst : *ObjectWithPrimitivesClass
States: 7
---- State #0 ----
Interpreter{
  PC: ((true & (ObjectWithPrimitivesClass_1_sel0 > ObjectWithPrimitivesClass_1_sel1)) & (ObjectWithPrimitivesClass_2_sel0 > ObjectWithPrimitivesClass_2_sel1))
  ReturnResults:
    [0] &(addr_0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
      P[1]: &(addr_1) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=0, store=0, select=0)
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=2, store=0, select=2)
        op[0]: select idx=addr_0, res=ObjectWithPrimitivesClass_1_sel0
        op[1]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel1
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=2, store=0, select=2)
        op[0]: select idx=addr_0, res=ObjectWithPrimitivesClass_2_sel0
        op[1]: select idx=addr_1, res=ObjectWithPrimitivesClass_2_sel1
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=0, store=0, select=0)
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=0, store=0, select=0)
  }
  HeapConstraints: 10
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: ObjectWithPrimitivesClass_2.select(addr_0, ObjectWithPrimitivesClass_2_sel0).select(addr_1, ObjectWithPrimitivesClass_2_sel1)
    C[6]: ObjectWithPrimitivesClass_3
    C[7]: ObjectWithPrimitivesClass_4
    C[8]: ObjectWithPrimitivesClass_0
    C[9]: ObjectWithPrimitivesClass_1.select(addr_0, ObjectWithPrimitivesClass_1_sel0).select(addr_1, ObjectWithPrimitivesClass_1_sel1)
}

SAT
Model:
addr_1 -> #x0000000501fe02d2
addr_0 -> #xfffffffafe01fd2d
ObjectWithPrimitivesClass_1 -> (store ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x8000ffffffff8000)
                                      #xfffffffafe01fd2d
                                      #x7fff000000004fff)
ObjectWithPrimitivesClass_2 -> (store ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x8000000000000000)
                                      #xfffffffafe01fd2d
                                      #x7fffffffffffffff)
ObjectWithPrimitivesClass_2_sel0 -> #x7fffffffffffffff
ObjectWithPrimitivesClass_1_sel1 -> #x8000ffffffff8000
ObjectWithPrimitivesClass_2_sel1 -> #x8000000000000000
ObjectWithPrimitivesClass_1_sel0 -> #x7fff000000004fff

---- State #1 ----
Interpreter{
  PC: (((true & (ObjectWithPrimitivesClass_1_sel0 > ObjectWithPrimitivesClass_1_sel1)) & (!(ObjectWithPrimitivesClass_2_sel0 > ObjectWithPrimitivesClass_2_sel1))) & (!(ObjectWithPrimitivesClass_1_sel2 < ObjectWithPrimitivesClass_1_sel3)))
  ReturnResults:
    [0] &(addr_0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
      P[1]: &(addr_1) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=0, store=0, select=0)
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_0, res=ObjectWithPrimitivesClass_1_sel0
        op[1]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel1
        op[2]: select idx=addr_0, res=ObjectWithPrimitivesClass_1_sel2
        op[3]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel3
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=2, store=0, select=2)
        op[0]: select idx=addr_0, res=ObjectWithPrimitivesClass_2_sel0
        op[1]: select idx=addr_1, res=ObjectWithPrimitivesClass_2_sel1
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=0, store=0, select=0)
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=0, store=0, select=0)
  }
  HeapConstraints: 10
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: ObjectWithPrimitivesClass_4
    C[6]: ObjectWithPrimitivesClass_0
    C[7]: ObjectWithPrimitivesClass_1.select(addr_0, ObjectWithPrimitivesClass_1_sel0).select(addr_1, ObjectWithPrimitivesClass_1_sel1).select(addr_0, ObjectWithPrimitivesClass_1_sel2).select(addr_1, ObjectWithPrimitivesClass_1_sel3)
    C[8]: ObjectWithPrimitivesClass_2.select(addr_0, ObjectWithPrimitivesClass_2_sel0).select(addr_1, ObjectWithPrimitivesClass_2_sel1)
    C[9]: ObjectWithPrimitivesClass_3
}

SAT
Model:
addr_1 -> #xffffffffffffffff
addr_0 -> #x0000000000000000
ObjectWithPrimitivesClass_1 -> (store ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x7fffffffffffffff)
                                      #xffffffffffffffff
                                      #x8000000000000000)
ObjectWithPrimitivesClass_2 -> (store ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x7fffffffffffffff)
                                      #x0000000000000000
                                      #x8000000000000000)
ObjectWithPrimitivesClass_2_sel0 -> #x8000000000000000
ObjectWithPrimitivesClass_1_sel3 -> #x8000000000000000
ObjectWithPrimitivesClass_1_sel1 -> #x8000000000000000
ObjectWithPrimitivesClass_2_sel1 -> #x7fffffffffffffff
ObjectWithPrimitivesClass_1_sel0 -> #x7fffffffffffffff
ObjectWithPrimitivesClass_1_sel2 -> #x7fffffffffffffff

---- State #2 ----
Interpreter{
  PC: ((((true & (ObjectWithPrimitivesClass_1_sel0 > ObjectWithPrimitivesClass_1_sel1)) & (!(ObjectWithPrimitivesClass_2_sel0 > ObjectWithPrimitivesClass_2_sel1))) & (ObjectWithPrimitivesClass_1_sel2 < ObjectWithPrimitivesClass_1_sel3)) & (ObjectWithPrimitivesClass_2_sel2 < ObjectWithPrimitivesClass_2_sel3))
  ReturnResults:
    [0] &(addr_1)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
      P[1]: &(addr_1) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=0, store=0, select=0)
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_0, res=ObjectWithPrimitivesClass_1_sel0
        op[1]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel1
        op[2]: select idx=addr_0, res=ObjectWithPrimitivesClass_1_sel2
        op[3]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel3
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_0, res=ObjectWithPrimitivesClass_2_sel0
        op[1]: select idx=addr_1, res=ObjectWithPrimitivesClass_2_sel1
        op[2]: select idx=addr_0, res=ObjectWithPrimitivesClass_2_sel2
        op[3]: select idx=addr_1, res=ObjectWithPrimitivesClass_2_sel3
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=0, store=0, select=0)
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=0, store=0, select=0)
  }
  HeapConstraints: 10
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: ObjectWithPrimitivesClass_4
    C[6]: ObjectWithPrimitivesClass_0
    C[7]: ObjectWithPrimitivesClass_1.select(addr_0, ObjectWithPrimitivesClass_1_sel0).select(addr_1, ObjectWithPrimitivesClass_1_sel1).select(addr_0, ObjectWithPrimitivesClass_1_sel2).select(addr_1, ObjectWithPrimitivesClass_1_sel3)
    C[8]: ObjectWithPrimitivesClass_2.select(addr_0, ObjectWithPrimitivesClass_2_sel0).select(addr_1, ObjectWithPrimitivesClass_2_sel1).select(addr_0, ObjectWithPrimitivesClass_2_sel2).select(addr_1, ObjectWithPrimitivesClass_2_sel3)
    C[9]: ObjectWithPrimitivesClass_3
}

UNSAT
---- State #3 ----
Interpreter{
  PC: ((((true & (ObjectWithPrimitivesClass_1_sel0 > ObjectWithPrimitivesClass_1_sel1)) & (!(ObjectWithPrimitivesClass_2_sel0 > ObjectWithPrimitivesClass_2_sel1))) & (ObjectWithPrimitivesClass_1_sel2 < ObjectWithPrimitivesClass_1_sel3)) & (!(ObjectWithPrimitivesClass_2_sel2 < ObjectWithPrimitivesClass_2_sel3)))
  ReturnResults:
    [0] &(addr_0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
      P[1]: &(addr_1) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=0, store=0, select=0)
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_0, res=ObjectWithPrimitivesClass_1_sel0
        op[1]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel1
        op[2]: select idx=addr_0, res=ObjectWithPrimitivesClass_1_sel2
        op[3]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel3
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_0, res=ObjectWithPrimitivesClass_2_sel0
        op[1]: select idx=addr_1, res=ObjectWithPrimitivesClass_2_sel1
        op[2]: select idx=addr_0, res=ObjectWithPrimitivesClass_2_sel2
        op[3]: select idx=addr_1, res=ObjectWithPrimitivesClass_2_sel3
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=0, store=0, select=0)
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=0, store=0, select=0)
  }
  HeapConstraints: 10
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: ObjectWithPrimitivesClass_1.select(addr_0, ObjectWithPrimitivesClass_1_sel0).select(addr_1, ObjectWithPrimitivesClass_1_sel1).select(addr_0, ObjectWithPrimitivesClass_1_sel2).select(addr_1, ObjectWithPrimitivesClass_1_sel3)
    C[6]: ObjectWithPrimitivesClass_2.select(addr_0, ObjectWithPrimitivesClass_2_sel0).select(addr_1, ObjectWithPrimitivesClass_2_sel1).select(addr_0, ObjectWithPrimitivesClass_2_sel2).select(addr_1, ObjectWithPrimitivesClass_2_sel3)
    C[7]: ObjectWithPrimitivesClass_3
    C[8]: ObjectWithPrimitivesClass_4
    C[9]: ObjectWithPrimitivesClass_0
}

UNSAT
---- State #4 ----
Interpreter{
  PC: (((true & (!(ObjectWithPrimitivesClass_1_sel0 > ObjectWithPrimitivesClass_1_sel1))) & (ObjectWithPrimitivesClass_1_sel2 < ObjectWithPrimitivesClass_1_sel3)) & (!(ObjectWithPrimitivesClass_2_sel0 < ObjectWithPrimitivesClass_2_sel1)))
  ReturnResults:
    [0] &(addr_0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
      P[1]: &(addr_1) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=0, store=0, select=0)
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_0, res=ObjectWithPrimitivesClass_1_sel0
        op[1]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel1
        op[2]: select idx=addr_0, res=ObjectWithPrimitivesClass_1_sel2
        op[3]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel3
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=2, store=0, select=2)
        op[0]: select idx=addr_0, res=ObjectWithPrimitivesClass_2_sel0
        op[1]: select idx=addr_1, res=ObjectWithPrimitivesClass_2_sel1
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=0, store=0, select=0)
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=0, store=0, select=0)
  }
  HeapConstraints: 10
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: ObjectWithPrimitivesClass_1.select(addr_0, ObjectWithPrimitivesClass_1_sel0).select(addr_1, ObjectWithPrimitivesClass_1_sel1).select(addr_0, ObjectWithPrimitivesClass_1_sel2).select(addr_1, ObjectWithPrimitivesClass_1_sel3)
    C[6]: ObjectWithPrimitivesClass_2.select(addr_0, ObjectWithPrimitivesClass_2_sel0).select(addr_1, ObjectWithPrimitivesClass_2_sel1)
    C[7]: ObjectWithPrimitivesClass_3
    C[8]: ObjectWithPrimitivesClass_4
    C[9]: ObjectWithPrimitivesClass_0
}

SAT
Model:
addr_0 -> #x0400101012000000
addr_1 -> #x3176f5fff7fffddf
ObjectWithPrimitivesClass_1 -> (store ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x8000000490000000)
                                      #x3176f5fff7fffddf
                                      #x1ffff6fb6fb7ffff)
ObjectWithPrimitivesClass_2 -> (store ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0fa8620ae5851df2)
                                      #x3176f5fff7fffddf
                                      #xc1f5d00bd2051844)
ObjectWithPrimitivesClass_2_sel0 -> #x0fa8620ae5851df2
ObjectWithPrimitivesClass_1_sel3 -> #x1ffff6fb6fb7ffff
ObjectWithPrimitivesClass_1_sel1 -> #x1ffff6fb6fb7ffff
ObjectWithPrimitivesClass_2_sel1 -> #xc1f5d00bd2051844
ObjectWithPrimitivesClass_1_sel0 -> #x8000000490000000
ObjectWithPrimitivesClass_1_sel2 -> #x8000000490000000

---- State #5 ----
Interpreter{
  PC: (((true & (!(ObjectWithPrimitivesClass_1_sel0 > ObjectWithPrimitivesClass_1_sel1))) & (ObjectWithPrimitivesClass_1_sel2 < ObjectWithPrimitivesClass_1_sel3)) & (ObjectWithPrimitivesClass_2_sel0 < ObjectWithPrimitivesClass_2_sel1))
  ReturnResults:
    [0] &(addr_1)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
      P[1]: &(addr_1) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=0, store=0, select=0)
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_0, res=ObjectWithPrimitivesClass_1_sel0
        op[1]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel1
        op[2]: select idx=addr_0, res=ObjectWithPrimitivesClass_1_sel2
        op[3]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel3
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=2, store=0, select=2)
        op[0]: select idx=addr_0, res=ObjectWithPrimitivesClass_2_sel0
        op[1]: select idx=addr_1, res=ObjectWithPrimitivesClass_2_sel1
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=0, store=0, select=0)
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=0, store=0, select=0)
  }
  HeapConstraints: 10
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: ObjectWithPrimitivesClass_2.select(addr_0, ObjectWithPrimitivesClass_2_sel0).select(addr_1, ObjectWithPrimitivesClass_2_sel1)
    C[6]: ObjectWithPrimitivesClass_3
    C[7]: ObjectWithPrimitivesClass_4
    C[8]: ObjectWithPrimitivesClass_0
    C[9]: ObjectWithPrimitivesClass_1.select(addr_0, ObjectWithPrimitivesClass_1_sel0).select(addr_1, ObjectWithPrimitivesClass_1_sel1).select(addr_0, ObjectWithPrimitivesClass_1_sel2).select(addr_1, ObjectWithPrimitivesClass_1_sel3)
}

SAT
Model:
addr_0 -> #x0000000000020000
addr_1 -> #xfffffffffff9b3da
ObjectWithPrimitivesClass_1 -> (store ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0ba0030000000140)
                                      #xfffffffffff9b3da
                                      #x26bffdfff7fff953)
ObjectWithPrimitivesClass_2 -> (store ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000094)
                                      #xfffffffffff9b3da
                                      #x7fffffffffffff6b)
ObjectWithPrimitivesClass_2_sel0 -> #x0000000000000094
ObjectWithPrimitivesClass_1_sel3 -> #x26bffdfff7fff953
ObjectWithPrimitivesClass_1_sel1 -> #x26bffdfff7fff953
ObjectWithPrimitivesClass_2_sel1 -> #x7fffffffffffff6b
ObjectWithPrimitivesClass_1_sel0 -> #x0ba0030000000140
ObjectWithPrimitivesClass_1_sel2 -> #x0ba0030000000140

---- State #6 ----
Interpreter{
  PC: ((true & (!(ObjectWithPrimitivesClass_1_sel0 > ObjectWithPrimitivesClass_1_sel1))) & (!(ObjectWithPrimitivesClass_1_sel2 < ObjectWithPrimitivesClass_1_sel3)))
  ReturnResults:
    [0] &(addr_0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
      P[1]: &(addr_1) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=0, store=0, select=0)
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_0, res=ObjectWithPrimitivesClass_1_sel0
        op[1]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel1
        op[2]: select idx=addr_0, res=ObjectWithPrimitivesClass_1_sel2
        op[3]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel3
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=0, store=0, select=0)
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=0, store=0, select=0)
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=0, store=0, select=0)
  }
  HeapConstraints: 10
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: ObjectWithPrimitivesClass_1.select(addr_0, ObjectWithPrimitivesClass_1_sel0).select(addr_1, ObjectWithPrimitivesClass_1_sel1).select(addr_0, ObjectWithPrimitivesClass_1_sel2).select(addr_1, ObjectWithPrimitivesClass_1_sel3)
    C[6]: ObjectWithPrimitivesClass_2
    C[7]: ObjectWithPrimitivesClass_3
    C[8]: ObjectWithPrimitivesClass_4
    C[9]: ObjectWithPrimitivesClass_0
}

SAT
Model:
addr_1 -> #xffffffffffffffff
addr_0 -> #x0000000000000000
ObjectWithPrimitivesClass_1 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x6f01ea1c06b3866e)
ObjectWithPrimitivesClass_1_sel3 -> #x6f01ea1c06b3866e
ObjectWithPrimitivesClass_1_sel1 -> #x6f01ea1c06b3866e
ObjectWithPrimitivesClass_1_sel0 -> #x6f01ea1c06b3866e
ObjectWithPrimitivesClass_1_sel2 -> #x6f01ea1c06b3866e


==================== Example ====================
CFG построен для функции с 3 блоками
interpretDynamically: &value.x [#1]
interpretDynamically: *t0
interpretDynamically: t1 == 1:int
interpretDynamically: if t2 goto 1 else 2
resolveExpression: t1 == 1:int
resolveExpression: *t0
resolveExpression: parameter value : *ObjectWithPrimitivesClass
resolveExpression: 1:int
interpretDynamically: &value.x [#1]
interpretDynamically: *t3 = 1:int
resolveExpression: 1:int
resolveExpression: parameter value : *ObjectWithPrimitivesClass
interpretDynamically: return value
resolveExpression: parameter value : *ObjectWithPrimitivesClass
interpretDynamically: return value
resolveExpression: parameter value : *ObjectWithPrimitivesClass
States: 2
---- State #0 ----
Interpreter{
  PC: (true & (!(ObjectWithPrimitivesClass_1_sel0 == 1)))
  ReturnResults:
    [0] &(addr_0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=0, store=0, select=0)
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=2, store=1, select=1)
        op[0]: select idx=addr_0, res=ObjectWithPrimitivesClass_1_sel0
        op[1]: store idx=addr_0, val=1
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=0, store=0, select=0)
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=0, store=0, select=0)
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=0, store=0, select=0)
  }
  HeapConstraints: 10
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: ObjectWithPrimitivesClass_0
    C[6]: ObjectWithPrimitivesClass_1.select(addr_0, ObjectWithPrimitivesClass_1_sel0).store(addr_0, 1)
    C[7]: ObjectWithPrimitivesClass_2
    C[8]: ObjectWithPrimitivesClass_3
    C[9]: ObjectWithPrimitivesClass_4
}

SAT
Model:
ObjectWithPrimitivesClass_1 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #xfffffffffffffffe)
addr_0 -> #x0000000000000000
ObjectWithPrimitivesClass_1_sel0 -> #xfffffffffffffffe

---- State #1 ----
Interpreter{
  PC: (true & (ObjectWithPrimitivesClass_1_sel0 == 1))
  ReturnResults:
    [0] &(addr_0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=0, store=0, select=0)
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=1, store=0, select=1)
        op[0]: select idx=addr_0, res=ObjectWithPrimitivesClass_1_sel0
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=0, store=0, select=0)
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=0, store=0, select=0)
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=0, store=0, select=0)
  }
  HeapConstraints: 10
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: ObjectWithPrimitivesClass_0
    C[6]: ObjectWithPrimitivesClass_1.select(addr_0, ObjectWithPrimitivesClass_1_sel0)
    C[7]: ObjectWithPrimitivesClass_2
    C[8]: ObjectWithPrimitivesClass_3
    C[9]: ObjectWithPrimitivesClass_4
}

SAT
Model:
ObjectWithPrimitivesClass_1 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000001)
ObjectWithPrimitivesClass_1_sel0 -> #x0000000000000001


==================== CreateObject ====================
CFG построен для функции с 3 блоками
interpretDynamically: new ObjectWithPrimitivesClass (complit)
interpretDynamically: &t0.ValueByDefault [#0]
interpretDynamically: *t1 = 5:int
resolveExpression: 5:int
resolveExpression: new ObjectWithPrimitivesClass (complit)
interpretDynamically: a + 5:int
interpretDynamically: &t0.x [#1]
interpretDynamically: *t3 = t2
resolveExpression: a + 5:int
resolveExpression: parameter a : int
resolveExpression: 5:int
resolveExpression: new ObjectWithPrimitivesClass (complit)
interpretDynamically: b + 6:int
interpretDynamically: &t0.y [#2]
interpretDynamically: *t5 = t4
resolveExpression: b + 6:int
resolveExpression: parameter b : int
resolveExpression: 6:int
resolveExpression: new ObjectWithPrimitivesClass (complit)
interpretDynamically: &objectExample.Weight [#4]
interpretDynamically: *t6
interpretDynamically: &t0.Weight [#4]
interpretDynamically: *t8 = t7
resolveExpression: *t6
resolveExpression: parameter objectExample : *ObjectWithPrimitivesClass
resolveExpression: new ObjectWithPrimitivesClass (complit)
interpretDynamically: &t0.Weight [#4]
interpretDynamically: *t9
interpretDynamically: t10 < 0:float64
interpretDynamically: if t11 goto 1 else 2
resolveExpression: t10 < 0:float64
resolveExpression: *t9
resolveExpression: new ObjectWithPrimitivesClass (complit)
resolveExpression: 0:float64
interpretDynamically: return t0
resolveExpression: new ObjectWithPrimitivesClass (complit)
interpretDynamically: return nil:*ObjectWithPrimitivesClass
resolveExpression: nil:*main.ObjectWithPrimitivesClass
States: 2
---- State #0 ----
Interpreter{
  PC: (true & (!(ObjectWithPrimitivesClass_4_sel2 < 0)))
  ReturnResults:
    [0] &(addr_0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 4
    Params: 3, Allocations: 1
    ParamsList:
      P[0]: &(addr_1) (refType=int)
      P[1]: &(addr_2) (refType=int)
      P[2]: &(addr_3) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
    AllocationsList:
      A[0]: &(addr_0) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=2, store=0, select=2)
        op[0]: select idx=addr_1, res=IntSegment_sel0
        op[1]: select idx=addr_2, res=IntSegment_sel1
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=1, store=1, select=0)
        op[0]: store idx=addr_0, val=5
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=1, store=1, select=0)
        op[0]: store idx=addr_0, val=(IntSegment_sel0 + 5)
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=1, store=1, select=0)
        op[0]: store idx=addr_0, val=(IntSegment_sel1 + 6)
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=0, store=0, select=0)
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=3, store=1, select=2)
        op[0]: select idx=addr_3, res=ObjectWithPrimitivesClass_4_sel0
        op[1]: store idx=addr_0, val=ObjectWithPrimitivesClass_4_sel0
        op[2]: select idx=addr_0, res=ObjectWithPrimitivesClass_4_sel2
  }
  HeapConstraints: 13
    C[0]: (addr_1 != addr_0)
    C[1]: (addr_2 != addr_0)
    C[2]: (addr_3 != addr_0)
    C[3]: IntSegment.select(addr_1, IntSegment_sel0).select(addr_2, IntSegment_sel1)
    C[4]: BoolSegment
    C[5]: StringSegment
    C[6]: FloatSegment
    C[7]: ArraySegment
    C[8]: ObjectWithPrimitivesClass_4.select(addr_3, ObjectWithPrimitivesClass_4_sel0).store(addr_0, ObjectWithPrimitivesClass_4_sel0).select(addr_0, ObjectWithPrimitivesClass_4_sel2)
    C[9]: ObjectWithPrimitivesClass_0.store(addr_0, 5)
    C[10]: ObjectWithPrimitivesClass_1.store(addr_0, (IntSegment_sel0 + 5))
    C[11]: ObjectWithPrimitivesClass_2.store(addr_0, (IntSegment_sel1 + 6))
    C[12]: ObjectWithPrimitivesClass_3
}

SAT
Model:
ObjectWithPrimitivesClass_4 -> ((as const (Array (_ BitVec 64) Real)) 0.0)
addr_1 -> #xffffffffffffffff
addr_2 -> #xffffffffffffffff
addr_3 -> #xffffffffffffffff
addr_0 -> #x0000000000000000
ObjectWithPrimitivesClass_4_sel2 -> 0.0
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
IntSegment_sel1 -> #x0000000000000000
ObjectWithPrimitivesClass_4_sel0 -> 0.0
IntSegment_sel0 -> #x0000000000000000

---- State #1 ----
Interpreter{
  PC: (true & (ObjectWithPrimitivesClass_4_sel2 < 0))
  ReturnResults:
    [0] nil
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 4
    Params: 3, Allocations: 1
    ParamsList:
      P[0]: &(addr_1) (refType=int)
      P[1]: &(addr_2) (refType=int)
      P[2]: &(addr_3) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
    AllocationsList:
      A[0]: &(addr_0) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=2, store=0, select=2)
        op[0]: select idx=addr_1, res=IntSegment_sel0
        op[1]: select idx=addr_2, res=IntSegment_sel1
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=1, store=1, select=0)
        op[0]: store idx=addr_0, val=5
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=1, store=1, select=0)
        op[0]: store idx=addr_0, val=(IntSegment_sel0 + 5)
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=1, store=1, select=0)
        op[0]: store idx=addr_0, val=(IntSegment_sel1 + 6)
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=0, store=0, select=0)
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=3, store=1, select=2)
        op[0]: select idx=addr_3, res=ObjectWithPrimitivesClass_4_sel0
        op[1]: store idx=addr_0, val=ObjectWithPrimitivesClass_4_sel0
        op[2]: select idx=addr_0, res=ObjectWithPrimitivesClass_4_sel2
  }
  HeapConstraints: 13
    C[0]: (addr_1 != addr_0)
    C[1]: (addr_2 != addr_0)
    C[2]: (addr_3 != addr_0)
    C[3]: IntSegment.select(addr_1, IntSegment_sel0).select(addr_2, IntSegment_sel1)
    C[4]: BoolSegment
    C[5]: StringSegment
    C[6]: FloatSegment
    C[7]: ArraySegment
    C[8]: ObjectWithPrimitivesClass_0.store(addr_0, 5)
    C[9]: ObjectWithPrimitivesClass_1.store(addr_0, (IntSegment_sel0 + 5))
    C[10]: ObjectWithPrimitivesClass_2.store(addr_0, (IntSegment_sel1 + 6))
    C[11]: ObjectWithPrimitivesClass_3
    C[12]: ObjectWithPrimitivesClass_4.select(addr_3, ObjectWithPrimitivesClass_4_sel0).store(addr_0, ObjectWithPrimitivesClass_4_sel0).select(addr_0, ObjectWithPrimitivesClass_4_sel2)
}

SAT
Model:
addr_3 -> #xffffffffffffffff
addr_1 -> #xffffffffffffffff
addr_2 -> #xffffffffffffffff
ObjectWithPrimitivesClass_4 -> ((as const (Array (_ BitVec 64) Real)) (- 1.0))
addr_0 -> #x0000000000000000
ObjectWithPrimitivesClass_4_sel2 -> (- 1.0)
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
IntSegment_sel1 -> #x0000000000000000
ObjectWithPrimitivesClass_4_sel0 -> (- 1.0)
IntSegment_sel0 -> #x0000000000000000


==================== Memory ====================
CFG построен для функции с 4 блоками
interpretDynamically: value > 0:int
interpretDynamically: if t0 goto 1 else 3
resolveExpression: value > 0:int
resolveExpression: parameter value : int
resolveExpression: 0:int
interpretDynamically: &objectExample.x [#1]
interpretDynamically: *t4 = -1:int
resolveExpression: -1:int
resolveExpression: parameter objectExample : *ObjectWithPrimitivesClass
interpretDynamically: &objectExample.y [#2]
interpretDynamically: *t5 = -2:int
resolveExpression: -2:int
resolveExpression: parameter objectExample : *ObjectWithPrimitivesClass
interpretDynamically: &objectExample.Weight [#4]
interpretDynamically: *t6 = -1.2:float64
resolveExpression: -1.2:float64
resolveExpression: parameter objectExample : *ObjectWithPrimitivesClass
interpretDynamically: jump 2
interpretDynamically: return objectExample
resolveExpression: parameter objectExample : *ObjectWithPrimitivesClass
interpretDynamically: &objectExample.x [#1]
interpretDynamically: *t1 = 1:int
resolveExpression: 1:int
resolveExpression: parameter objectExample : *ObjectWithPrimitivesClass
interpretDynamically: &objectExample.y [#2]
interpretDynamically: *t2 = 2:int
resolveExpression: 2:int
resolveExpression: parameter objectExample : *ObjectWithPrimitivesClass
interpretDynamically: &objectExample.Weight [#4]
interpretDynamically: *t3 = 1.2:float64
resolveExpression: 1.2:float64
resolveExpression: parameter objectExample : *ObjectWithPrimitivesClass
interpretDynamically: jump 2
interpretDynamically: return objectExample
resolveExpression: parameter objectExample : *ObjectWithPrimitivesClass
States: 2
---- State #0 ----
Interpreter{
  PC: (true & (!(IntSegment_sel0 > 0)))
  ReturnResults:
    [0] &(addr_1)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=1, store=0, select=1)
        op[0]: select idx=addr_0, res=IntSegment_sel0
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=0, store=0, select=0)
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=1, store=1, select=0)
        op[0]: store idx=addr_1, val=-1
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=1, store=1, select=0)
        op[0]: store idx=addr_1, val=-2
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=0, store=0, select=0)
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=1, store=1, select=0)
        op[0]: store idx=addr_1, val=-1.2
  }
  HeapConstraints: 10
    C[0]: IntSegment.select(addr_0, IntSegment_sel0)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: ObjectWithPrimitivesClass_3
    C[6]: ObjectWithPrimitivesClass_4.store(addr_1, -1.2)
    C[7]: ObjectWithPrimitivesClass_0
    C[8]: ObjectWithPrimitivesClass_1.store(addr_1, -1)
    C[9]: ObjectWithPrimitivesClass_2.store(addr_1, -2)
}

SAT
Model:
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
addr_0 -> #x0000000000000000
IntSegment_sel0 -> #x0000000000000000

---- State #1 ----
Interpreter{
  PC: (true & (IntSegment_sel0 > 0))
  ReturnResults:
    [0] &(addr_1)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=1, store=0, select=1)
        op[0]: select idx=addr_0, res=IntSegment_sel0
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=0, store=0, select=0)
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=1, store=1, select=0)
        op[0]: store idx=addr_1, val=1
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=1, store=1, select=0)
        op[0]: store idx=addr_1, val=2
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=0, store=0, select=0)
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=1, store=1, select=0)
        op[0]: store idx=addr_1, val=1.2
  }
  HeapConstraints: 10
    C[0]: IntSegment.select(addr_0, IntSegment_sel0)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: ObjectWithPrimitivesClass_2.store(addr_1, 2)
    C[6]: ObjectWithPrimitivesClass_3
    C[7]: ObjectWithPrimitivesClass_4.store(addr_1, 1.2)
    C[8]: ObjectWithPrimitivesClass_0
    C[9]: ObjectWithPrimitivesClass_1.store(addr_1, 1)
}

SAT
Model:
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000001)
addr_0 -> #x0000000000000000
IntSegment_sel0 -> #x0000000000000001


==================== CompareTwoNullObjects ====================
CFG построен для функции с 7 блоками
interpretDynamically: new ObjectWithPrimitivesClass (complit)
interpretDynamically: &t0.ValueByDefault [#0]
interpretDynamically: *t1 = 5:int
resolveExpression: 5:int
resolveExpression: new ObjectWithPrimitivesClass (complit)
interpretDynamically: new ObjectWithPrimitivesClass (complit)
interpretDynamically: &t2.ValueByDefault [#0]
interpretDynamically: *t3 = 5:int
resolveExpression: 5:int
resolveExpression: new ObjectWithPrimitivesClass (complit)
interpretDynamically: value + 1:int
interpretDynamically: &t0.x [#1]
interpretDynamically: *t5 = t4
resolveExpression: value + 1:int
resolveExpression: parameter value : int
resolveExpression: 1:int
resolveExpression: new ObjectWithPrimitivesClass (complit)
interpretDynamically: value + 2:int
interpretDynamically: &t2.x [#1]
interpretDynamically: *t7 = t6
resolveExpression: value + 2:int
resolveExpression: parameter value : int
resolveExpression: 2:int
resolveExpression: new ObjectWithPrimitivesClass (complit)
interpretDynamically: &t0.x [#1]
interpretDynamically: *t8
interpretDynamically: value + 1:int
interpretDynamically: t9 == t10
interpretDynamically: if t11 goto 1 else 2
resolveExpression: t9 == t10
resolveExpression: *t8
resolveExpression: new ObjectWithPrimitivesClass (complit)
resolveExpression: value + 1:int
resolveExpression: parameter value : int
resolveExpression: 1:int
interpretDynamically: jump 2
interpretDynamically: phi [0: t0, 1: nil:*ObjectWithPrimitivesClass] #fst
resolveExpression: nil:*main.ObjectWithPrimitivesClass
interpretDynamically: phi [0: t0, 1: nil:*ObjectWithPrimitivesClass] #fst
resolveExpression: new ObjectWithPrimitivesClass (complit)
interpretDynamically: &t2.x [#1]
interpretDynamically: *t13
interpretDynamically: value + 2:int
interpretDynamically: t14 == t15
interpretDynamically: if t16 goto 3 else 4
resolveExpression: t14 == t15
resolveExpression: *t13
resolveExpression: new ObjectWithPrimitivesClass (complit)
resolveExpression: value + 2:int
resolveExpression: parameter value : int
resolveExpression: 2:int
interpretDynamically: jump 4
interpretDynamically: phi [2: t2, 3: nil:*ObjectWithPrimitivesClass] #snd
resolveExpression: nil:*main.ObjectWithPrimitivesClass
interpretDynamically: t12 == t17
interpretDynamically: if t18 goto 5 else 6
resolveExpression: t12 == t17
resolveExpression: phi [0: t0, 1: nil:*ObjectWithPrimitivesClass] #fst
resolveExpression: phi [2: t2, 3: nil:*ObjectWithPrimitivesClass] #snd
interpretDynamically: return 0:int
resolveExpression: 0:int
interpretDynamically: phi [2: t2, 3: nil:*ObjectWithPrimitivesClass] #snd
resolveExpression: new ObjectWithPrimitivesClass (complit)
interpretDynamically: t12 == t17
interpretDynamically: if t18 goto 5 else 6
resolveExpression: t12 == t17
resolveExpression: phi [0: t0, 1: nil:*ObjectWithPrimitivesClass] #fst
resolveExpression: phi [2: t2, 3: nil:*ObjectWithPrimitivesClass] #snd
interpretDynamically: return 0:int
resolveExpression: 0:int
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: &t2.x [#1]
interpretDynamically: *t13
interpretDynamically: value + 2:int
interpretDynamically: t14 == t15
interpretDynamically: if t16 goto 3 else 4
resolveExpression: t14 == t15
resolveExpression: *t13
resolveExpression: new ObjectWithPrimitivesClass (complit)
resolveExpression: value + 2:int
resolveExpression: parameter value : int
resolveExpression: 2:int
interpretDynamically: jump 4
interpretDynamically: phi [2: t2, 3: nil:*ObjectWithPrimitivesClass] #snd
resolveExpression: new ObjectWithPrimitivesClass (complit)
interpretDynamically: t12 == t17
interpretDynamically: phi [2: t2, 3: nil:*ObjectWithPrimitivesClass] #snd
resolveExpression: nil:*main.ObjectWithPrimitivesClass
interpretDynamically: t12 == t17
interpretDynamically: if t18 goto 5 else 6
resolveExpression: t12 == t17
resolveExpression: phi [0: t0, 1: nil:*ObjectWithPrimitivesClass] #fst
resolveExpression: phi [2: t2, 3: nil:*ObjectWithPrimitivesClass] #snd
interpretDynamically: return 0:int
resolveExpression: 0:int
interpretDynamically: if t18 goto 5 else 6
resolveExpression: t12 == t17
resolveExpression: phi [0: t0, 1: nil:*ObjectWithPrimitivesClass] #fst
resolveExpression: phi [2: t2, 3: nil:*ObjectWithPrimitivesClass] #snd
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: return 0:int
resolveExpression: 0:int
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: return 1:int
resolveExpression: 1:int
States: 8
---- State #0 ----
Interpreter{
  PC: (((true & (ObjectWithPrimitivesClass_1_sel2 == (IntSegment_sel2 + 1))) & (ObjectWithPrimitivesClass_1_sel3 == (IntSegment_sel3 + 2))) & (!(0 == 0)))
  ReturnResults:
    [0] 0
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 1, Allocations: 2
    ParamsList:
      P[0]: &(addr_2) (refType=int)
    AllocationsList:
      A[0]: &(addr_0) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
      A[1]: &(addr_1) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_2, res=IntSegment_sel0
        op[1]: select idx=addr_2, res=IntSegment_sel1
        op[2]: select idx=addr_2, res=IntSegment_sel2
        op[3]: select idx=addr_2, res=IntSegment_sel3
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=2, store=2, select=0)
        op[0]: store idx=addr_0, val=5
        op[1]: store idx=addr_1, val=5
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=4, store=2, select=2)
        op[0]: store idx=addr_0, val=(IntSegment_sel0 + 1)
        op[1]: store idx=addr_1, val=(IntSegment_sel1 + 2)
        op[2]: select idx=addr_0, res=ObjectWithPrimitivesClass_1_sel2
        op[3]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel3
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=0, store=0, select=0)
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=0, store=0, select=0)
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=0, store=0, select=0)
  }
  HeapConstraints: 13
    C[0]: (addr_2 != addr_0)
    C[1]: (addr_2 != addr_1)
    C[2]: (addr_0 != addr_1)
    C[3]: IntSegment.select(addr_2, IntSegment_sel0).select(addr_2, IntSegment_sel1).select(addr_2, IntSegment_sel2).select(addr_2, IntSegment_sel3)
    C[4]: BoolSegment
    C[5]: StringSegment
    C[6]: FloatSegment
    C[7]: ArraySegment
    C[8]: ObjectWithPrimitivesClass_0.store(addr_0, 5).store(addr_1, 5)
    C[9]: ObjectWithPrimitivesClass_1.store(addr_0, (IntSegment_sel0 + 1)).store(addr_1, (IntSegment_sel1 + 2)).select(addr_0, ObjectWithPrimitivesClass_1_sel2).select(addr_1, ObjectWithPrimitivesClass_1_sel3)
    C[10]: ObjectWithPrimitivesClass_2
    C[11]: ObjectWithPrimitivesClass_3
    C[12]: ObjectWithPrimitivesClass_4
}

UNSAT
---- State #1 ----
Interpreter{
  PC: (((true & (ObjectWithPrimitivesClass_1_sel2 == (IntSegment_sel2 + 1))) & (!(ObjectWithPrimitivesClass_1_sel3 == (IntSegment_sel3 + 2)))) & (!(&(addr_0) == addr_1)))
  ReturnResults:
    [0] 0
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 1, Allocations: 2
    ParamsList:
      P[0]: &(addr_2) (refType=int)
    AllocationsList:
      A[0]: &(addr_0) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
      A[1]: &(addr_1) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_2, res=IntSegment_sel0
        op[1]: select idx=addr_2, res=IntSegment_sel1
        op[2]: select idx=addr_2, res=IntSegment_sel2
        op[3]: select idx=addr_2, res=IntSegment_sel3
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=2, store=2, select=0)
        op[0]: store idx=addr_0, val=5
        op[1]: store idx=addr_1, val=5
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=4, store=2, select=2)
        op[0]: store idx=addr_0, val=(IntSegment_sel0 + 1)
        op[1]: store idx=addr_1, val=(IntSegment_sel1 + 2)
        op[2]: select idx=addr_0, res=ObjectWithPrimitivesClass_1_sel2
        op[3]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel3
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=0, store=0, select=0)
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=0, store=0, select=0)
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=0, store=0, select=0)
  }
  HeapConstraints: 13
    C[0]: (addr_2 != addr_0)
    C[1]: (addr_2 != addr_1)
    C[2]: (addr_0 != addr_1)
    C[3]: IntSegment.select(addr_2, IntSegment_sel0).select(addr_2, IntSegment_sel1).select(addr_2, IntSegment_sel2).select(addr_2, IntSegment_sel3)
    C[4]: BoolSegment
    C[5]: StringSegment
    C[6]: FloatSegment
    C[7]: ArraySegment
    C[8]: ObjectWithPrimitivesClass_2
    C[9]: ObjectWithPrimitivesClass_3
    C[10]: ObjectWithPrimitivesClass_4
    C[11]: ObjectWithPrimitivesClass_0.store(addr_0, 5).store(addr_1, 5)
    C[12]: ObjectWithPrimitivesClass_1.store(addr_0, (IntSegment_sel0 + 1)).store(addr_1, (IntSegment_sel1 + 2)).select(addr_0, ObjectWithPrimitivesClass_1_sel2).select(addr_1, ObjectWithPrimitivesClass_1_sel3)
}

UNSAT
---- State #2 ----
Interpreter{
  PC: (((true & (!(ObjectWithPrimitivesClass_1_sel2 == (IntSegment_sel2 + 1)))) & (ObjectWithPrimitivesClass_1_sel3 == (IntSegment_sel3 + 2))) & (!(addr_0 == &(addr_0))))
  ReturnResults:
    [0] 0
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 1, Allocations: 2
    ParamsList:
      P[0]: &(addr_2) (refType=int)
    AllocationsList:
      A[0]: &(addr_0) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
      A[1]: &(addr_1) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_2, res=IntSegment_sel0
        op[1]: select idx=addr_2, res=IntSegment_sel1
        op[2]: select idx=addr_2, res=IntSegment_sel2
        op[3]: select idx=addr_2, res=IntSegment_sel3
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=2, store=2, select=0)
        op[0]: store idx=addr_0, val=5
        op[1]: store idx=addr_1, val=5
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=4, store=2, select=2)
        op[0]: store idx=addr_0, val=(IntSegment_sel0 + 1)
        op[1]: store idx=addr_1, val=(IntSegment_sel1 + 2)
        op[2]: select idx=addr_0, res=ObjectWithPrimitivesClass_1_sel2
        op[3]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel3
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=0, store=0, select=0)
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=0, store=0, select=0)
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=0, store=0, select=0)
  }
  HeapConstraints: 13
    C[0]: (addr_2 != addr_0)
    C[1]: (addr_2 != addr_1)
    C[2]: (addr_0 != addr_1)
    C[3]: IntSegment.select(addr_2, IntSegment_sel0).select(addr_2, IntSegment_sel1).select(addr_2, IntSegment_sel2).select(addr_2, IntSegment_sel3)
    C[4]: BoolSegment
    C[5]: StringSegment
    C[6]: FloatSegment
    C[7]: ArraySegment
    C[8]: ObjectWithPrimitivesClass_0.store(addr_0, 5).store(addr_1, 5)
    C[9]: ObjectWithPrimitivesClass_1.store(addr_0, (IntSegment_sel0 + 1)).store(addr_1, (IntSegment_sel1 + 2)).select(addr_0, ObjectWithPrimitivesClass_1_sel2).select(addr_1, ObjectWithPrimitivesClass_1_sel3)
    C[10]: ObjectWithPrimitivesClass_2
    C[11]: ObjectWithPrimitivesClass_3
    C[12]: ObjectWithPrimitivesClass_4
}

UNSAT
---- State #3 ----
Interpreter{
  PC: (((true & (!(ObjectWithPrimitivesClass_1_sel2 == (IntSegment_sel2 + 1)))) & (!(ObjectWithPrimitivesClass_1_sel3 == (IntSegment_sel3 + 2)))) & (addr_0 == addr_1))
  ReturnResults:
    [0] 1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 1, Allocations: 2
    ParamsList:
      P[0]: &(addr_2) (refType=int)
    AllocationsList:
      A[0]: &(addr_0) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
      A[1]: &(addr_1) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_2, res=IntSegment_sel0
        op[1]: select idx=addr_2, res=IntSegment_sel1
        op[2]: select idx=addr_2, res=IntSegment_sel2
        op[3]: select idx=addr_2, res=IntSegment_sel3
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=2, store=2, select=0)
        op[0]: store idx=addr_0, val=5
        op[1]: store idx=addr_1, val=5
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=4, store=2, select=2)
        op[0]: store idx=addr_0, val=(IntSegment_sel0 + 1)
        op[1]: store idx=addr_1, val=(IntSegment_sel1 + 2)
        op[2]: select idx=addr_0, res=ObjectWithPrimitivesClass_1_sel2
        op[3]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel3
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=0, store=0, select=0)
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=0, store=0, select=0)
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=0, store=0, select=0)
  }
  HeapConstraints: 13
    C[0]: (addr_2 != addr_0)
    C[1]: (addr_2 != addr_1)
    C[2]: (addr_0 != addr_1)
    C[3]: IntSegment.select(addr_2, IntSegment_sel0).select(addr_2, IntSegment_sel1).select(addr_2, IntSegment_sel2).select(addr_2, IntSegment_sel3)
    C[4]: BoolSegment
    C[5]: StringSegment
    C[6]: FloatSegment
    C[7]: ArraySegment
    C[8]: ObjectWithPrimitivesClass_0.store(addr_0, 5).store(addr_1, 5)
    C[9]: ObjectWithPrimitivesClass_1.store(addr_0, (IntSegment_sel0 + 1)).store(addr_1, (IntSegment_sel1 + 2)).select(addr_0, ObjectWithPrimitivesClass_1_sel2).select(addr_1, ObjectWithPrimitivesClass_1_sel3)
    C[10]: ObjectWithPrimitivesClass_2
    C[11]: ObjectWithPrimitivesClass_3
    C[12]: ObjectWithPrimitivesClass_4
}

UNSAT
---- State #4 ----
Interpreter{
  PC: (((true & (!(ObjectWithPrimitivesClass_1_sel2 == (IntSegment_sel2 + 1)))) & (!(ObjectWithPrimitivesClass_1_sel3 == (IntSegment_sel3 + 2)))) & (!(addr_0 == addr_1)))
  ReturnResults:
    [0] 0
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 1, Allocations: 2
    ParamsList:
      P[0]: &(addr_2) (refType=int)
    AllocationsList:
      A[0]: &(addr_0) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
      A[1]: &(addr_1) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_2, res=IntSegment_sel0
        op[1]: select idx=addr_2, res=IntSegment_sel1
        op[2]: select idx=addr_2, res=IntSegment_sel2
        op[3]: select idx=addr_2, res=IntSegment_sel3
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=2, store=2, select=0)
        op[0]: store idx=addr_0, val=5
        op[1]: store idx=addr_1, val=5
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=4, store=2, select=2)
        op[0]: store idx=addr_0, val=(IntSegment_sel0 + 1)
        op[1]: store idx=addr_1, val=(IntSegment_sel1 + 2)
        op[2]: select idx=addr_0, res=ObjectWithPrimitivesClass_1_sel2
        op[3]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel3
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=0, store=0, select=0)
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=0, store=0, select=0)
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=0, store=0, select=0)
  }
  HeapConstraints: 13
    C[0]: (addr_2 != addr_0)
    C[1]: (addr_2 != addr_1)
    C[2]: (addr_0 != addr_1)
    C[3]: IntSegment.select(addr_2, IntSegment_sel0).select(addr_2, IntSegment_sel1).select(addr_2, IntSegment_sel2).select(addr_2, IntSegment_sel3)
    C[4]: BoolSegment
    C[5]: StringSegment
    C[6]: FloatSegment
    C[7]: ArraySegment
    C[8]: ObjectWithPrimitivesClass_3
    C[9]: ObjectWithPrimitivesClass_4
    C[10]: ObjectWithPrimitivesClass_0.store(addr_0, 5).store(addr_1, 5)
    C[11]: ObjectWithPrimitivesClass_1.store(addr_0, (IntSegment_sel0 + 1)).store(addr_1, (IntSegment_sel1 + 2)).select(addr_0, ObjectWithPrimitivesClass_1_sel2).select(addr_1, ObjectWithPrimitivesClass_1_sel3)
    C[12]: ObjectWithPrimitivesClass_2
}

UNSAT
---- State #5 ----
Interpreter{
  PC: (((true & (ObjectWithPrimitivesClass_1_sel2 == (IntSegment_sel2 + 1))) & (ObjectWithPrimitivesClass_1_sel3 == (IntSegment_sel3 + 2))) & (0 == 0))
  ReturnResults:
    [0] 1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 1, Allocations: 2
    ParamsList:
      P[0]: &(addr_2) (refType=int)
    AllocationsList:
      A[0]: &(addr_0) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
      A[1]: &(addr_1) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_2, res=IntSegment_sel0
        op[1]: select idx=addr_2, res=IntSegment_sel1
        op[2]: select idx=addr_2, res=IntSegment_sel2
        op[3]: select idx=addr_2, res=IntSegment_sel3
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=2, store=2, select=0)
        op[0]: store idx=addr_0, val=5
        op[1]: store idx=addr_1, val=5
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=4, store=2, select=2)
        op[0]: store idx=addr_0, val=(IntSegment_sel0 + 1)
        op[1]: store idx=addr_1, val=(IntSegment_sel1 + 2)
        op[2]: select idx=addr_0, res=ObjectWithPrimitivesClass_1_sel2
        op[3]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel3
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=0, store=0, select=0)
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=0, store=0, select=0)
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=0, store=0, select=0)
  }
  HeapConstraints: 13
    C[0]: (addr_2 != addr_0)
    C[1]: (addr_2 != addr_1)
    C[2]: (addr_0 != addr_1)
    C[3]: IntSegment.select(addr_2, IntSegment_sel0).select(addr_2, IntSegment_sel1).select(addr_2, IntSegment_sel2).select(addr_2, IntSegment_sel3)
    C[4]: BoolSegment
    C[5]: StringSegment
    C[6]: FloatSegment
    C[7]: ArraySegment
    C[8]: ObjectWithPrimitivesClass_1.store(addr_0, (IntSegment_sel0 + 1)).store(addr_1, (IntSegment_sel1 + 2)).select(addr_0, ObjectWithPrimitivesClass_1_sel2).select(addr_1, ObjectWithPrimitivesClass_1_sel3)
    C[9]: ObjectWithPrimitivesClass_2
    C[10]: ObjectWithPrimitivesClass_3
    C[11]: ObjectWithPrimitivesClass_4
    C[12]: ObjectWithPrimitivesClass_0.store(addr_0, 5).store(addr_1, 5)
}

SAT
Model:
addr_2 -> #x0000000000000002
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x7fffffffffffffff)
addr_0 -> #x0000000000000001
addr_1 -> #x0000000000000000
ObjectWithPrimitivesClass_1_sel3 -> #x8000000000000001
IntSegment_sel2 -> #x7fffffffffffffff
IntSegment_sel1 -> #x7fffffffffffffff
IntSegment_sel0 -> #x7fffffffffffffff
IntSegment_sel3 -> #x7fffffffffffffff
ObjectWithPrimitivesClass_1_sel2 -> #x8000000000000000

---- State #6 ----
Interpreter{
  PC: (((true & (!(ObjectWithPrimitivesClass_1_sel2 == (IntSegment_sel2 + 1)))) & (ObjectWithPrimitivesClass_1_sel3 == (IntSegment_sel3 + 2))) & (addr_0 == &(addr_0)))
  ReturnResults:
    [0] 1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 1, Allocations: 2
    ParamsList:
      P[0]: &(addr_2) (refType=int)
    AllocationsList:
      A[0]: &(addr_0) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
      A[1]: &(addr_1) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_2, res=IntSegment_sel0
        op[1]: select idx=addr_2, res=IntSegment_sel1
        op[2]: select idx=addr_2, res=IntSegment_sel2
        op[3]: select idx=addr_2, res=IntSegment_sel3
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=2, store=2, select=0)
        op[0]: store idx=addr_0, val=5
        op[1]: store idx=addr_1, val=5
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=4, store=2, select=2)
        op[0]: store idx=addr_0, val=(IntSegment_sel0 + 1)
        op[1]: store idx=addr_1, val=(IntSegment_sel1 + 2)
        op[2]: select idx=addr_0, res=ObjectWithPrimitivesClass_1_sel2
        op[3]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel3
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=0, store=0, select=0)
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=0, store=0, select=0)
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=0, store=0, select=0)
  }
  HeapConstraints: 13
    C[0]: (addr_2 != addr_0)
    C[1]: (addr_2 != addr_1)
    C[2]: (addr_0 != addr_1)
    C[3]: IntSegment.select(addr_2, IntSegment_sel0).select(addr_2, IntSegment_sel1).select(addr_2, IntSegment_sel2).select(addr_2, IntSegment_sel3)
    C[4]: BoolSegment
    C[5]: StringSegment
    C[6]: FloatSegment
    C[7]: ArraySegment
    C[8]: ObjectWithPrimitivesClass_0.store(addr_0, 5).store(addr_1, 5)
    C[9]: ObjectWithPrimitivesClass_1.store(addr_0, (IntSegment_sel0 + 1)).store(addr_1, (IntSegment_sel1 + 2)).select(addr_0, ObjectWithPrimitivesClass_1_sel2).select(addr_1, ObjectWithPrimitivesClass_1_sel3)
    C[10]: ObjectWithPrimitivesClass_2
    C[11]: ObjectWithPrimitivesClass_3
    C[12]: ObjectWithPrimitivesClass_4
}

UNSAT
---- State #7 ----
Interpreter{
  PC: (((true & (ObjectWithPrimitivesClass_1_sel2 == (IntSegment_sel2 + 1))) & (!(ObjectWithPrimitivesClass_1_sel3 == (IntSegment_sel3 + 2)))) & (&(addr_0) == addr_1))
  ReturnResults:
    [0] 1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 1, Allocations: 2
    ParamsList:
      P[0]: &(addr_2) (refType=int)
    AllocationsList:
      A[0]: &(addr_0) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
      A[1]: &(addr_1) (refType=struct, type=ObjectWithPrimitivesClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=4, store=0, select=4)
        op[0]: select idx=addr_2, res=IntSegment_sel0
        op[1]: select idx=addr_2, res=IntSegment_sel1
        op[2]: select idx=addr_2, res=IntSegment_sel2
        op[3]: select idx=addr_2, res=IntSegment_sel3
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithPrimitivesClass:
        field[0]: ObjectWithPrimitivesClass_0 (elem=int, ops=2, store=2, select=0)
        op[0]: store idx=addr_0, val=5
        op[1]: store idx=addr_1, val=5
        field[1]: ObjectWithPrimitivesClass_1 (elem=int, ops=4, store=2, select=2)
        op[0]: store idx=addr_0, val=(IntSegment_sel0 + 1)
        op[1]: store idx=addr_1, val=(IntSegment_sel1 + 2)
        op[2]: select idx=addr_0, res=ObjectWithPrimitivesClass_1_sel2
        op[3]: select idx=addr_1, res=ObjectWithPrimitivesClass_1_sel3
        field[2]: ObjectWithPrimitivesClass_2 (elem=int, ops=0, store=0, select=0)
        field[3]: ObjectWithPrimitivesClass_3 (elem=int, ops=0, store=0, select=0)
        field[4]: ObjectWithPrimitivesClass_4 (elem=float, ops=0, store=0, select=0)
  }
  HeapConstraints: 13
    C[0]: (addr_2 != addr_0)
    C[1]: (addr_2 != addr_1)
    C[2]: (addr_0 != addr_1)
    C[3]: IntSegment.select(addr_2, IntSegment_sel0).select(addr_2, IntSegment_sel1).select(addr_2, IntSegment_sel2).select(addr_2, IntSegment_sel3)
    C[4]: BoolSegment
    C[5]: StringSegment
    C[6]: FloatSegment
    C[7]: ArraySegment
    C[8]: ObjectWithPrimitivesClass_0.store(addr_0, 5).store(addr_1, 5)
    C[9]: ObjectWithPrimitivesClass_1.store(addr_0, (IntSegment_sel0 + 1)).store(addr_1, (IntSegment_sel1 + 2)).select(addr_0, ObjectWithPrimitivesClass_1_sel2).select(addr_1, ObjectWithPrimitivesClass_1_sel3)
    C[10]: ObjectWithPrimitivesClass_2
    C[11]: ObjectWithPrimitivesClass_3
    C[12]: ObjectWithPrimitivesClass_4
}

UNSAT

==================== WriteToRefTypeField ====================
CFG построен для функции с 3 блоками
interpretDynamically: value != 42:int
interpretDynamically: if t0 goto 1 else 2
resolveExpression: value != 42:int
resolveExpression: parameter value : int
resolveExpression: 42:int
interpretDynamically: return nil:*ObjectWithRefFieldClass
resolveExpression: nil:*main.ObjectWithRefFieldClass
interpretDynamically: new SimpleDataClass (complit)
interpretDynamically: &t1.a [#0]
interpretDynamically: &t1.b [#1]
interpretDynamically: value * 2:int
interpretDynamically: *t2 = value
resolveExpression: parameter value : int
resolveExpression: new SimpleDataClass (complit)
interpretDynamically: *t3 = t4
resolveExpression: value * 2:int
resolveExpression: parameter value : int
resolveExpression: 2:int
resolveExpression: new SimpleDataClass (complit)
interpretDynamically: &objectExample.RefField [#3]
interpretDynamically: *t5 = t1
resolveExpression: new SimpleDataClass (complit)
resolveExpression: parameter objectExample : *ObjectWithRefFieldClass
interpretDynamically: return objectExample
resolveExpression: parameter objectExample : *ObjectWithRefFieldClass
States: 2
---- State #0 ----
Interpreter{
  PC: (true & (IntSegment_sel0 != 42))
  ReturnResults:
    [0] nil
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
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #xffffffffffffffd5)
addr_0 -> #x0000000000000000
IntSegment_sel0 -> #xffffffffffffffd5

---- State #1 ----
Interpreter{
  PC: (true & (!(IntSegment_sel0 != 42)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_2) (refType=struct, type=ObjectWithRefFieldClass, field=-1)
    AllocationsList:
      A[0]: &(addr_1) (refType=struct, type=SimpleDataClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=3, store=0, select=3)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 2 type(s)
      ObjectWithRefFieldClass:
        field[0]: ObjectWithRefFieldClass_0 (elem=int, ops=0, store=0, select=0)
        field[1]: ObjectWithRefFieldClass_1 (elem=int, ops=0, store=0, select=0)
        field[2]: ObjectWithRefFieldClass_2 (elem=float, ops=0, store=0, select=0)
        field[3]: ObjectWithRefFieldClass_3 (elem=struct, ops=1, store=1, select=0)
        op[0]: store idx=addr_2, val=&(addr_1)
        field[4]: ObjectWithRefFieldClass_4 (elem=int, ops=0, store=0, select=0)
      SimpleDataClass:
        field[0]: SimpleDataClass_0 (elem=int, ops=1, store=1, select=0)
        op[0]: store idx=addr_1, val=IntSegment_sel1
        field[1]: SimpleDataClass_1 (elem=int, ops=1, store=1, select=0)
        op[0]: store idx=addr_1, val=(IntSegment_sel2 * 2)
  }
  HeapConstraints: 14
    C[0]: (addr_0 != addr_1)
    C[1]: (addr_2 != addr_1)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2)
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment
    C[7]: SimpleDataClass_0.store(addr_1, IntSegment_sel1)
    C[8]: SimpleDataClass_1.store(addr_1, (IntSegment_sel2 * 2))
    C[9]: ObjectWithRefFieldClass_3.store(addr_2, &(addr_1))
    C[10]: ObjectWithRefFieldClass_4
    C[11]: ObjectWithRefFieldClass_0
    C[12]: ObjectWithRefFieldClass_1
    C[13]: ObjectWithRefFieldClass_2
}

SAT
Model:
addr_0 -> #xffffffffffffffff
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x000000000000002a)
addr_2 -> #xffffffffffffffff
addr_1 -> #x0000000000000000
IntSegment_sel2 -> #x000000000000002a
IntSegment_sel1 -> #x000000000000002a
IntSegment_sel0 -> #x000000000000002a


==================== WriteToArrayField ====================
CFG построен для функции с 6 блоками
interpretDynamically: length < 3:int
interpretDynamically: if t0 goto 1 else 2
resolveExpression: length < 3:int
resolveExpression: parameter length : int
resolveExpression: 3:int
interpretDynamically: make []int length length
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 3: t4] #i
resolveExpression: 0:int
interpretDynamically: t10 < length
interpretDynamically: return nil:*ObjectWithRefFieldClass
resolveExpression: nil:*main.ObjectWithRefFieldClass
interpretDynamically: if t11 goto 3 else 4
resolveExpression: t10 < length
resolveExpression: phi [2: 0:int, 3: t4] #i
resolveExpression: parameter length : int
interpretDynamically: &objectExample.ArrayField [#4]
interpretDynamically: t10 + 1:int
interpretDynamically: &t1[t10]
interpretDynamically: *t3 = t2
resolveExpression: t10 + 1:int
resolveExpression: phi [2: 0:int, 3: t4] #i
resolveExpression: 1:int
resolveExpression: make []int length length
resolveExpression: phi [2: 0:int, 3: t4] #i
interpretDynamically: t10 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 3: t4] #i
resolveExpression: t10 + 1:int
resolveExpression: phi [2: 0:int, 3: t4] #i
resolveExpression: 1:int
interpretDynamically: t10 < length
interpretDynamically: if t11 goto 3 else 4
resolveExpression: t10 < length
resolveExpression: phi [2: 0:int, 3: t4] #i
resolveExpression: parameter length : int
interpretDynamically: &objectExample.ArrayField [#4]
interpretDynamically: *t5 = t1
resolveExpression: make []int length length
resolveExpression: parameter objectExample : *ObjectWithRefFieldClass
interpretDynamically: &objectExample.ArrayField [#4]
interpretDynamically: t10 + 1:int
interpretDynamically: *t6
interpretDynamically: length - 1:int
interpretDynamically: &t7[t8]
interpretDynamically: *t9 = 100:int
resolveExpression: 100:int
resolveExpression: *t6
resolveExpression: parameter objectExample : *ObjectWithRefFieldClass
resolveExpression: length - 1:int
resolveExpression: parameter length : int
resolveExpression: 1:int
interpretDynamically: return objectExample
resolveExpression: parameter objectExample : *ObjectWithRefFieldClass
interpretDynamically: *t5 = t1
resolveExpression: make []int length length
resolveExpression: parameter objectExample : *ObjectWithRefFieldClass
interpretDynamically: &objectExample.ArrayField [#4]
interpretDynamically: *t6
interpretDynamically: length - 1:int
interpretDynamically: &t7[t8]
interpretDynamically: *t9 = 100:int
resolveExpression: 100:int
resolveExpression: *t6
resolveExpression: parameter objectExample : *ObjectWithRefFieldClass
resolveExpression: length - 1:int
resolveExpression: parameter length : int
resolveExpression: 1:int
interpretDynamically: return objectExample
resolveExpression: parameter objectExample : *ObjectWithRefFieldClass
interpretDynamically: &t1[t10]
interpretDynamically: *t3 = t2
resolveExpression: t10 + 1:int
resolveExpression: phi [2: 0:int, 3: t4] #i
resolveExpression: 1:int
resolveExpression: make []int length length
resolveExpression: phi [2: 0:int, 3: t4] #i
interpretDynamically: t10 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 3: t4] #i
resolveExpression: t10 + 1:int
resolveExpression: phi [2: 0:int, 3: t4] #i
resolveExpression: 1:int
interpretDynamically: t10 < length
interpretDynamically: if t11 goto 3 else 4
resolveExpression: t10 < length
resolveExpression: phi [2: 0:int, 3: t4] #i
resolveExpression: parameter length : int
interpretDynamically: t10 + 1:int
interpretDynamically: &objectExample.ArrayField [#4]
interpretDynamically: &t1[t10]
interpretDynamically: *t3 = t2
resolveExpression: t10 + 1:int
resolveExpression: phi [2: 0:int, 3: t4] #i
resolveExpression: 1:int
resolveExpression: make []int length length
resolveExpression: phi [2: 0:int, 3: t4] #i
interpretDynamically: t10 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 3: t4] #i
resolveExpression: t10 + 1:int
resolveExpression: phi [2: 0:int, 3: t4] #i
resolveExpression: 1:int
interpretDynamically: t10 < length
interpretDynamically: if t11 goto 3 else 4
resolveExpression: t10 < length
resolveExpression: phi [2: 0:int, 3: t4] #i
resolveExpression: parameter length : int
interpretDynamically: &objectExample.ArrayField [#4]
interpretDynamically: *t5 = t1
resolveExpression: make []int length length
resolveExpression: parameter objectExample : *ObjectWithRefFieldClass
interpretDynamically: &objectExample.ArrayField [#4]
interpretDynamically: t10 + 1:int
interpretDynamically: &t1[t10]
interpretDynamically: *t3 = t2
resolveExpression: t10 + 1:int
resolveExpression: phi [2: 0:int, 3: t4] #i
resolveExpression: 1:int
resolveExpression: make []int length length
resolveExpression: phi [2: 0:int, 3: t4] #i
interpretDynamically: t10 + 1:int
interpretDynamically: jump 5
interpretDynamically: phi [2: 0:int, 3: t4] #i
resolveExpression: t10 + 1:int
resolveExpression: phi [2: 0:int, 3: t4] #i
resolveExpression: 1:int
interpretDynamically: t10 < length
interpretDynamically: if t11 goto 3 else 4
resolveExpression: t10 < length
resolveExpression: phi [2: 0:int, 3: t4] #i
resolveExpression: parameter length : int
interpretDynamically: &objectExample.ArrayField [#4]
interpretDynamically: t10 + 1:int
interpretDynamically: &t1[t10]
interpretDynamically: *t3 = t2
resolveExpression: t10 + 1:int
resolveExpression: phi [2: 0:int, 3: t4] #i
resolveExpression: 1:int
resolveExpression: make []int length length
resolveExpression: phi [2: 0:int, 3: t4] #i
interpretDynamically: *t5 = t1
resolveExpression: make []int length length
resolveExpression: parameter objectExample : *ObjectWithRefFieldClass
interpretDynamically: *t6
interpretDynamically: length - 1:int
interpretDynamically: t10 + 1:int
interpretDynamically: jump 5
interpretDynamically: &objectExample.ArrayField [#4]
interpretDynamically: *t6
interpretDynamically: length - 1:int
interpretDynamically: &t7[t8]
interpretDynamically: *t9 = 100:int
resolveExpression: 100:int
resolveExpression: *t6
resolveExpression: parameter objectExample : *ObjectWithRefFieldClass
resolveExpression: length - 1:int
resolveExpression: parameter length : int
resolveExpression: 1:int
interpretDynamically: return objectExample
resolveExpression: parameter objectExample : *ObjectWithRefFieldClass
interpretDynamically: *t5 = t1
resolveExpression: make []int length length
resolveExpression: parameter objectExample : *ObjectWithRefFieldClass
interpretDynamically: &objectExample.ArrayField [#4]
interpretDynamically: *t6
interpretDynamically: length - 1:int
interpretDynamically: &t7[t8]
interpretDynamically: *t9 = 100:int
resolveExpression: 100:int
resolveExpression: *t6
resolveExpression: parameter objectExample : *ObjectWithRefFieldClass
resolveExpression: length - 1:int
resolveExpression: parameter length : int
resolveExpression: 1:int
interpretDynamically: return objectExample
resolveExpression: parameter objectExample : *ObjectWithRefFieldClass
interpretDynamically: &t7[t8]
interpretDynamically: *t9 = 100:int
resolveExpression: 100:int
resolveExpression: *t6
resolveExpression: parameter objectExample : *ObjectWithRefFieldClass
resolveExpression: length - 1:int
resolveExpression: parameter length : int
resolveExpression: 1:int
interpretDynamically: return objectExample
resolveExpression: parameter objectExample : *ObjectWithRefFieldClass
States: 6
---- State #0 ----
Interpreter{
  PC: (true & (IntSegment_sel0 < 3))
  ReturnResults:
    [0] nil
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
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000002)
addr_0 -> #x0000000000000000
IntSegment_sel0 -> #x0000000000000002

---- State #1 ----
Interpreter{
  PC: (((true & (!(IntSegment_sel0 < 3))) & (0 < IntSegment_sel1)) & (!((0 + 1) < IntSegment_sel3)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_2) (refType=struct, type=ObjectWithRefFieldClass, field=-1)
    AllocationsList:
      A[0]: &(addr_1) (refType=array, dim=0, base=int)
    IntSegment:    IntSegment (elem=int, ops=6, store=2, select=4)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: store idx=(addr_1 + 0), val=(0 + 1)
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: select idx=addr_0, res=IntSegment_sel4
        op[5]: store idx=(ObjectWithRefFieldClass_4_sel1 + (IntSegment_sel4 - 1)), val=100
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithRefFieldClass:
        field[0]: ObjectWithRefFieldClass_0 (elem=int, ops=0, store=0, select=0)
        field[1]: ObjectWithRefFieldClass_1 (elem=int, ops=0, store=0, select=0)
        field[2]: ObjectWithRefFieldClass_2 (elem=float, ops=0, store=0, select=0)
        field[3]: ObjectWithRefFieldClass_3 (elem=struct, ops=0, store=0, select=0)
        field[4]: ObjectWithRefFieldClass_4 (elem=int, ops=2, store=1, select=1)
        op[0]: store idx=addr_2, val=&(addr_1)
        op[1]: select idx=addr_2, res=ObjectWithRefFieldClass_4_sel1
  }
  HeapConstraints: 12
    C[0]: (addr_0 != addr_1)
    C[1]: (addr_2 != addr_1)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).store((addr_1 + 0), (0 + 1)).select(addr_0, IntSegment_sel3).select(addr_0, IntSegment_sel4).store((ObjectWithRefFieldClass_4_sel1 + (IntSegment_sel4 - 1)), 100)
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment
    C[7]: ObjectWithRefFieldClass_0
    C[8]: ObjectWithRefFieldClass_1
    C[9]: ObjectWithRefFieldClass_2
    C[10]: ObjectWithRefFieldClass_3
    C[11]: ObjectWithRefFieldClass_4.store(addr_2, &(addr_1)).select(addr_2, ObjectWithRefFieldClass_4_sel1)
}

UNSAT
---- State #2 ----
Interpreter{
  PC: ((true & (!(IntSegment_sel0 < 3))) & (!(0 < IntSegment_sel1)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_2) (refType=struct, type=ObjectWithRefFieldClass, field=-1)
    AllocationsList:
      A[0]: &(addr_1) (refType=array, dim=0, base=int)
    IntSegment:    IntSegment (elem=int, ops=4, store=1, select=3)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: select idx=addr_0, res=IntSegment_sel2
        op[3]: store idx=(ObjectWithRefFieldClass_4_sel1 + (IntSegment_sel2 - 1)), val=100
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithRefFieldClass:
        field[0]: ObjectWithRefFieldClass_0 (elem=int, ops=0, store=0, select=0)
        field[1]: ObjectWithRefFieldClass_1 (elem=int, ops=0, store=0, select=0)
        field[2]: ObjectWithRefFieldClass_2 (elem=float, ops=0, store=0, select=0)
        field[3]: ObjectWithRefFieldClass_3 (elem=struct, ops=0, store=0, select=0)
        field[4]: ObjectWithRefFieldClass_4 (elem=int, ops=2, store=1, select=1)
        op[0]: store idx=addr_2, val=&(addr_1)
        op[1]: select idx=addr_2, res=ObjectWithRefFieldClass_4_sel1
  }
  HeapConstraints: 12
    C[0]: (addr_0 != addr_1)
    C[1]: (addr_2 != addr_1)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).select(addr_0, IntSegment_sel2).store((ObjectWithRefFieldClass_4_sel1 + (IntSegment_sel2 - 1)), 100)
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment
    C[7]: ObjectWithRefFieldClass_2
    C[8]: ObjectWithRefFieldClass_3
    C[9]: ObjectWithRefFieldClass_4.store(addr_2, &(addr_1)).select(addr_2, ObjectWithRefFieldClass_4_sel1)
    C[10]: ObjectWithRefFieldClass_0
    C[11]: ObjectWithRefFieldClass_1
}

UNSAT
---- State #3 ----
Interpreter{
  PC: ((((((true & (!(IntSegment_sel0 < 3))) & (0 < IntSegment_sel1)) & ((0 + 1) < IntSegment_sel3)) & (((0 + 1) + 1) < IntSegment_sel5)) & ((((0 + 1) + 1) + 1) < IntSegment_sel7)) & (!(((((0 + 1) + 1) + 1) + 1) < IntSegment_sel9)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_2) (refType=struct, type=ObjectWithRefFieldClass, field=-1)
    AllocationsList:
      A[0]: &(addr_1) (refType=array, dim=0, base=int)
    IntSegment:    IntSegment (elem=int, ops=12, store=5, select=7)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: store idx=(addr_1 + 0), val=(0 + 1)
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: store idx=(addr_1 + (0 + 1)), val=((0 + 1) + 1)
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: store idx=(addr_1 + ((0 + 1) + 1)), val=(((0 + 1) + 1) + 1)
        op[7]: select idx=addr_0, res=IntSegment_sel7
        op[8]: store idx=(addr_1 + (((0 + 1) + 1) + 1)), val=((((0 + 1) + 1) + 1) + 1)
        op[9]: select idx=addr_0, res=IntSegment_sel9
        op[10]: select idx=addr_0, res=IntSegment_sel10
        op[11]: store idx=(ObjectWithRefFieldClass_4_sel1 + (IntSegment_sel10 - 1)), val=100
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithRefFieldClass:
        field[0]: ObjectWithRefFieldClass_0 (elem=int, ops=0, store=0, select=0)
        field[1]: ObjectWithRefFieldClass_1 (elem=int, ops=0, store=0, select=0)
        field[2]: ObjectWithRefFieldClass_2 (elem=float, ops=0, store=0, select=0)
        field[3]: ObjectWithRefFieldClass_3 (elem=struct, ops=0, store=0, select=0)
        field[4]: ObjectWithRefFieldClass_4 (elem=int, ops=2, store=1, select=1)
        op[0]: store idx=addr_2, val=&(addr_1)
        op[1]: select idx=addr_2, res=ObjectWithRefFieldClass_4_sel1
  }
  HeapConstraints: 12
    C[0]: (addr_0 != addr_1)
    C[1]: (addr_2 != addr_1)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).store((addr_1 + 0), (0 + 1)).select(addr_0, IntSegment_sel3).store((addr_1 + (0 + 1)), ((0 + 1) + 1)).select(addr_0, IntSegment_sel5).store((addr_1 + ((0 + 1) + 1)), (((0 + 1) + 1) + 1)).select(addr_0, IntSegment_sel7).store((addr_1 + (((0 + 1) + 1) + 1)), ((((0 + 1) + 1) + 1) + 1)).select(addr_0, IntSegment_sel9).select(addr_0, IntSegment_sel10).store((ObjectWithRefFieldClass_4_sel1 + (IntSegment_sel10 - 1)), 100)
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment
    C[7]: ObjectWithRefFieldClass_2
    C[8]: ObjectWithRefFieldClass_3
    C[9]: ObjectWithRefFieldClass_4.store(addr_2, &(addr_1)).select(addr_2, ObjectWithRefFieldClass_4_sel1)
    C[10]: ObjectWithRefFieldClass_0
    C[11]: ObjectWithRefFieldClass_1
}

SAT
Model:
addr_2 -> #x0000000000000003
addr_0 -> #xffffffffffffffff
ObjectWithRefFieldClass_4_sel1 -> #xfffffffffffffffc
IntSegment -> (store ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000003)
                     #xffffffffffffffff
                     #x0000000000000005)
addr_1 -> #xfffffffffffffffc
IntSegment_sel10 -> #x0000000000000004
IntSegment_sel1 -> #x0000000000000005
IntSegment_sel7 -> #x0000000000000005
IntSegment_sel5 -> #x0000000000000005
IntSegment_sel9 -> #x0000000000000004
IntSegment_sel3 -> #x0000000000000005
IntSegment_sel0 -> #x0000000000000005

---- State #4 ----
Interpreter{
  PC: ((((true & (!(IntSegment_sel0 < 3))) & (0 < IntSegment_sel1)) & ((0 + 1) < IntSegment_sel3)) & (!(((0 + 1) + 1) < IntSegment_sel5)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_2) (refType=struct, type=ObjectWithRefFieldClass, field=-1)
    AllocationsList:
      A[0]: &(addr_1) (refType=array, dim=0, base=int)
    IntSegment:    IntSegment (elem=int, ops=8, store=3, select=5)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: store idx=(addr_1 + 0), val=(0 + 1)
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: store idx=(addr_1 + (0 + 1)), val=((0 + 1) + 1)
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: select idx=addr_0, res=IntSegment_sel6
        op[7]: store idx=(ObjectWithRefFieldClass_4_sel1 + (IntSegment_sel6 - 1)), val=100
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithRefFieldClass:
        field[0]: ObjectWithRefFieldClass_0 (elem=int, ops=0, store=0, select=0)
        field[1]: ObjectWithRefFieldClass_1 (elem=int, ops=0, store=0, select=0)
        field[2]: ObjectWithRefFieldClass_2 (elem=float, ops=0, store=0, select=0)
        field[3]: ObjectWithRefFieldClass_3 (elem=struct, ops=0, store=0, select=0)
        field[4]: ObjectWithRefFieldClass_4 (elem=int, ops=2, store=1, select=1)
        op[0]: store idx=addr_2, val=&(addr_1)
        op[1]: select idx=addr_2, res=ObjectWithRefFieldClass_4_sel1
  }
  HeapConstraints: 12
    C[0]: (addr_0 != addr_1)
    C[1]: (addr_2 != addr_1)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).store((addr_1 + 0), (0 + 1)).select(addr_0, IntSegment_sel3).store((addr_1 + (0 + 1)), ((0 + 1) + 1)).select(addr_0, IntSegment_sel5).select(addr_0, IntSegment_sel6).store((ObjectWithRefFieldClass_4_sel1 + (IntSegment_sel6 - 1)), 100)
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment
    C[7]: ObjectWithRefFieldClass_0
    C[8]: ObjectWithRefFieldClass_1
    C[9]: ObjectWithRefFieldClass_2
    C[10]: ObjectWithRefFieldClass_3
    C[11]: ObjectWithRefFieldClass_4.store(addr_2, &(addr_1)).select(addr_2, ObjectWithRefFieldClass_4_sel1)
}

SAT
Model:
addr_2 -> #x00002228aaa00000
addr_0 -> #xffffddd755600000
ObjectWithRefFieldClass_4_sel1 -> #xffffddd7555fffff
IntSegment -> (store ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000003)
                     #xffffddd755600000
                     #x0200000000000002)
IntSegment_sel6 -> #x0000000000000002
addr_1 -> #xffffddd7555fffff
IntSegment_sel5 -> #x0000000000000002
IntSegment_sel1 -> #x0200000000000002
IntSegment_sel3 -> #x0200000000000002
IntSegment_sel0 -> #x0200000000000002

---- State #5 ----
Interpreter{
  PC: (((((true & (!(IntSegment_sel0 < 3))) & (0 < IntSegment_sel1)) & ((0 + 1) < IntSegment_sel3)) & (((0 + 1) + 1) < IntSegment_sel5)) & (!((((0 + 1) + 1) + 1) < IntSegment_sel7)))
  ReturnResults:
    [0] &(addr_2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_2) (refType=struct, type=ObjectWithRefFieldClass, field=-1)
    AllocationsList:
      A[0]: &(addr_1) (refType=array, dim=0, base=int)
    IntSegment:    IntSegment (elem=int, ops=10, store=4, select=6)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
        op[2]: store idx=(addr_1 + 0), val=(0 + 1)
        op[3]: select idx=addr_0, res=IntSegment_sel3
        op[4]: store idx=(addr_1 + (0 + 1)), val=((0 + 1) + 1)
        op[5]: select idx=addr_0, res=IntSegment_sel5
        op[6]: store idx=(addr_1 + ((0 + 1) + 1)), val=(((0 + 1) + 1) + 1)
        op[7]: select idx=addr_0, res=IntSegment_sel7
        op[8]: select idx=addr_0, res=IntSegment_sel8
        op[9]: store idx=(ObjectWithRefFieldClass_4_sel1 + (IntSegment_sel8 - 1)), val=100
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithRefFieldClass:
        field[0]: ObjectWithRefFieldClass_0 (elem=int, ops=0, store=0, select=0)
        field[1]: ObjectWithRefFieldClass_1 (elem=int, ops=0, store=0, select=0)
        field[2]: ObjectWithRefFieldClass_2 (elem=float, ops=0, store=0, select=0)
        field[3]: ObjectWithRefFieldClass_3 (elem=struct, ops=0, store=0, select=0)
        field[4]: ObjectWithRefFieldClass_4 (elem=int, ops=2, store=1, select=1)
        op[0]: store idx=addr_2, val=&(addr_1)
        op[1]: select idx=addr_2, res=ObjectWithRefFieldClass_4_sel1
  }
  HeapConstraints: 12
    C[0]: (addr_0 != addr_1)
    C[1]: (addr_2 != addr_1)
    C[2]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1).store((addr_1 + 0), (0 + 1)).select(addr_0, IntSegment_sel3).store((addr_1 + (0 + 1)), ((0 + 1) + 1)).select(addr_0, IntSegment_sel5).store((addr_1 + ((0 + 1) + 1)), (((0 + 1) + 1) + 1)).select(addr_0, IntSegment_sel7).select(addr_0, IntSegment_sel8).store((ObjectWithRefFieldClass_4_sel1 + (IntSegment_sel8 - 1)), 100)
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment
    C[7]: ObjectWithRefFieldClass_0
    C[8]: ObjectWithRefFieldClass_1
    C[9]: ObjectWithRefFieldClass_2
    C[10]: ObjectWithRefFieldClass_3
    C[11]: ObjectWithRefFieldClass_4.store(addr_2, &(addr_1)).select(addr_2, ObjectWithRefFieldClass_4_sel1)
}

SAT
Model:
addr_2 -> #x0000a2a0a0a90001
addr_0 -> #xffff5d5f5f570000
ObjectWithRefFieldClass_4_sel1 -> #xffff5d5f5f56fffe
IntSegment -> (store ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000003)
                     #xffff5d5f5f570000
                     #x2000000000000000)
IntSegment_sel7 -> #x0000000000000003
IntSegment_sel8 -> #x0000000000000003
addr_1 -> #xffff5d5f5f56fffe
IntSegment_sel5 -> #x2000000000000000
IntSegment_sel1 -> #x2000000000000000
IntSegment_sel3 -> #x2000000000000000
IntSegment_sel0 -> #x2000000000000000


==================== ReadFromArrayField ====================
CFG построен для функции с 4 блоками
interpretDynamically: &objectExample.ArrayField [#4]
interpretDynamically: *t0
interpretDynamically: len(t1)
resolveExpression: len(t1)
interpretDynamically: t2 > 2:int
interpretDynamically: if t3 goto 3 else 2
resolveExpression: t2 > 2:int
resolveExpression: len(t1)
resolveExpression: 2:int
interpretDynamically: &objectExample.ArrayField [#4]
interpretDynamically: *t4
interpretDynamically: return 2:int
resolveExpression: 2:int
interpretDynamically: &t5[2:int]
interpretDynamically: *t6
interpretDynamically: t7 == value
interpretDynamically: if t8 goto 1 else 2
resolveExpression: t7 == value
resolveExpression: *t6
resolveExpression: *t4
resolveExpression: parameter objectExample : *ObjectWithRefFieldClass
resolveExpression: 2:int
resolveExpression: parameter value : int
interpretDynamically: return 2:int
resolveExpression: 2:int
interpretDynamically: return 1:int
resolveExpression: 1:int
States: 3
---- State #0 ----
Interpreter{
  PC: (true & (!(len__t0 > 2)))
  ReturnResults:
    [0] 2
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 0
    Params: 0, Allocations: 0
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: <empty>
  }
  HeapConstraints: 5
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
}

SAT
Model:
len__t0 -> #x0000000000000002

---- State #1 ----
Interpreter{
  PC: ((true & (len__t0 > 2)) & (IntSegment_sel0 == IntSegment_sel1))
  ReturnResults:
    [0] 1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=ObjectWithRefFieldClass, field=-1)
      P[1]: &(addr_1) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=2, store=0, select=2)
        op[0]: select idx=(ObjectWithRefFieldClass_4_sel0 + 2), res=IntSegment_sel0
        op[1]: select idx=addr_1, res=IntSegment_sel1
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithRefFieldClass:
        field[0]: ObjectWithRefFieldClass_0 (elem=int, ops=0, store=0, select=0)
        field[1]: ObjectWithRefFieldClass_1 (elem=int, ops=0, store=0, select=0)
        field[2]: ObjectWithRefFieldClass_2 (elem=float, ops=0, store=0, select=0)
        field[3]: ObjectWithRefFieldClass_3 (elem=struct, ops=0, store=0, select=0)
        field[4]: ObjectWithRefFieldClass_4 (elem=int, ops=1, store=0, select=1)
        op[0]: select idx=addr_0, res=ObjectWithRefFieldClass_4_sel0
  }
  HeapConstraints: 10
    C[0]: IntSegment.select((ObjectWithRefFieldClass_4_sel0 + 2), IntSegment_sel0).select(addr_1, IntSegment_sel1)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: ObjectWithRefFieldClass_3
    C[6]: ObjectWithRefFieldClass_4.select(addr_0, ObjectWithRefFieldClass_4_sel0)
    C[7]: ObjectWithRefFieldClass_0
    C[8]: ObjectWithRefFieldClass_1
    C[9]: ObjectWithRefFieldClass_2
}

SAT
Model:
len__t0 -> #x0000000000000003
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
ObjectWithRefFieldClass_4 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #xfffffffffffffffe)
addr_1 -> #x0000000000000000
IntSegment_sel1 -> #x0000000000000000
addr_0 -> #x0000000000000000
ObjectWithRefFieldClass_4_sel0 -> #xfffffffffffffffe
IntSegment_sel0 -> #x0000000000000000

---- State #2 ----
Interpreter{
  PC: ((true & (len__t0 > 2)) & (!(IntSegment_sel0 == IntSegment_sel1)))
  ReturnResults:
    [0] 2
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=ObjectWithRefFieldClass, field=-1)
      P[1]: &(addr_1) (refType=int)
    IntSegment:    IntSegment (elem=int, ops=2, store=0, select=2)
        op[0]: select idx=(ObjectWithRefFieldClass_4_sel0 + 2), res=IntSegment_sel0
        op[1]: select idx=addr_1, res=IntSegment_sel1
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithRefFieldClass:
        field[0]: ObjectWithRefFieldClass_0 (elem=int, ops=0, store=0, select=0)
        field[1]: ObjectWithRefFieldClass_1 (elem=int, ops=0, store=0, select=0)
        field[2]: ObjectWithRefFieldClass_2 (elem=float, ops=0, store=0, select=0)
        field[3]: ObjectWithRefFieldClass_3 (elem=struct, ops=0, store=0, select=0)
        field[4]: ObjectWithRefFieldClass_4 (elem=int, ops=1, store=0, select=1)
        op[0]: select idx=addr_0, res=ObjectWithRefFieldClass_4_sel0
  }
  HeapConstraints: 10
    C[0]: IntSegment.select((ObjectWithRefFieldClass_4_sel0 + 2), IntSegment_sel0).select(addr_1, IntSegment_sel1)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: ObjectWithRefFieldClass_4.select(addr_0, ObjectWithRefFieldClass_4_sel0)
    C[6]: ObjectWithRefFieldClass_0
    C[7]: ObjectWithRefFieldClass_1
    C[8]: ObjectWithRefFieldClass_2
    C[9]: ObjectWithRefFieldClass_3
}

SAT
Model:
addr_1 -> #xffffffffffffffff
ObjectWithRefFieldClass_4 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #xfffffffffffffffe)
len__t0 -> #x0000000000000003
IntSegment -> (store ((as const (Array (_ BitVec 64) (_ BitVec 64))) #xffffffffffffffff)
                     #xffffffffffffffff
                     #x0000000000000000)
IntSegment_sel1 -> #x0000000000000000
addr_0 -> #x0000000000000000
ObjectWithRefFieldClass_4_sel0 -> #xfffffffffffffffe
IntSegment_sel0 -> #xffffffffffffffff


==================== CompareTwoDifferentObjectsFromArguments ====================
CFG построен для функции с 8 блоками
interpretDynamically: &fst.x [#0]
interpretDynamically: *t0
interpretDynamically: t1 > 0:int
interpretDynamically: if t2 goto 3 else 2
resolveExpression: t1 > 0:int
resolveExpression: *t0
resolveExpression: parameter fst : *ObjectWithRefFieldClass
resolveExpression: 0:int
interpretDynamically: &snd.x [#0]
interpretDynamically: *t4
interpretDynamically: &fst.x [#0]
interpretDynamically: &snd.x [#0]
interpretDynamically: *t6 = t5
resolveExpression: *t4
resolveExpression: parameter snd : *ObjectWithRefFieldClass
resolveExpression: parameter fst : *ObjectWithRefFieldClass
interpretDynamically: &snd.y [#1]
interpretDynamically: *t7
interpretDynamically: &fst.y [#1]
interpretDynamically: *t9 = t8
resolveExpression: *t7
resolveExpression: parameter snd : *ObjectWithRefFieldClass
resolveExpression: parameter fst : *ObjectWithRefFieldClass
interpretDynamically: &snd.Weight [#2]
interpretDynamically: *t10
interpretDynamically: &fst.Weight [#2]
interpretDynamically: *t12 = t11
resolveExpression: *t10
resolveExpression: parameter snd : *ObjectWithRefFieldClass
resolveExpression: parameter fst : *ObjectWithRefFieldClass
interpretDynamically: fst == snd
interpretDynamically: if t13 goto 6 else 7
resolveExpression: fst == snd
resolveExpression: parameter fst : *ObjectWithRefFieldClass
resolveExpression: parameter snd : *ObjectWithRefFieldClass
interpretDynamically: return 3:int
resolveExpression: 3:int
interpretDynamically: return 2:int
resolveExpression: 2:int
interpretDynamically: *t14
interpretDynamically: t15 < 0:int
interpretDynamically: if t16 goto 1 else 2
resolveExpression: t15 < 0:int
resolveExpression: *t14
resolveExpression: parameter snd : *ObjectWithRefFieldClass
resolveExpression: 0:int
interpretDynamically: fst == snd
interpretDynamically: if t3 goto 4 else 5
resolveExpression: fst == snd
resolveExpression: parameter fst : *ObjectWithRefFieldClass
resolveExpression: parameter snd : *ObjectWithRefFieldClass
interpretDynamically: return 0:int
resolveExpression: 0:int
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: &snd.x [#0]
interpretDynamically: *t4
interpretDynamically: &fst.x [#0]
interpretDynamically: *t6 = t5
resolveExpression: *t4
resolveExpression: parameter snd : *ObjectWithRefFieldClass
resolveExpression: parameter fst : *ObjectWithRefFieldClass
interpretDynamically: &snd.y [#1]
interpretDynamically: *t7
interpretDynamically: &fst.y [#1]
interpretDynamically: *t9 = t8
resolveExpression: *t7
resolveExpression: parameter snd : *ObjectWithRefFieldClass
resolveExpression: parameter fst : *ObjectWithRefFieldClass
interpretDynamically: &snd.Weight [#2]
interpretDynamically: *t10
interpretDynamically: &fst.Weight [#2]
interpretDynamically: *t12 = t11
resolveExpression: *t10
resolveExpression: parameter snd : *ObjectWithRefFieldClass
resolveExpression: parameter fst : *ObjectWithRefFieldClass
interpretDynamically: fst == snd
interpretDynamically: if t13 goto 6 else 7
resolveExpression: fst == snd
resolveExpression: parameter fst : *ObjectWithRefFieldClass
resolveExpression: parameter snd : *ObjectWithRefFieldClass
interpretDynamically: return 3:int
resolveExpression: 3:int
interpretDynamically: return 2:int
resolveExpression: 2:int
States: 6
---- State #0 ----
Interpreter{
  PC: ((true & (!(ObjectWithRefFieldClass_0_sel0 > 0))) & (!(addr_0 == addr_1)))
  ReturnResults:
    [0] 3
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=ObjectWithRefFieldClass, field=-1)
      P[1]: &(addr_1) (refType=struct, type=ObjectWithRefFieldClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithRefFieldClass:
        field[0]: ObjectWithRefFieldClass_0 (elem=int, ops=3, store=1, select=2)
        op[0]: select idx=addr_0, res=ObjectWithRefFieldClass_0_sel0
        op[1]: select idx=addr_1, res=ObjectWithRefFieldClass_0_sel1
        op[2]: store idx=addr_0, val=ObjectWithRefFieldClass_0_sel1
        field[1]: ObjectWithRefFieldClass_1 (elem=int, ops=2, store=1, select=1)
        op[0]: select idx=addr_1, res=ObjectWithRefFieldClass_1_sel0
        op[1]: store idx=addr_0, val=ObjectWithRefFieldClass_1_sel0
        field[2]: ObjectWithRefFieldClass_2 (elem=float, ops=2, store=1, select=1)
        op[0]: select idx=addr_1, res=ObjectWithRefFieldClass_2_sel0
        op[1]: store idx=addr_0, val=ObjectWithRefFieldClass_2_sel0
        field[3]: ObjectWithRefFieldClass_3 (elem=struct, ops=0, store=0, select=0)
        field[4]: ObjectWithRefFieldClass_4 (elem=int, ops=0, store=0, select=0)
  }
  HeapConstraints: 10
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: ObjectWithRefFieldClass_0.select(addr_0, ObjectWithRefFieldClass_0_sel0).select(addr_1, ObjectWithRefFieldClass_0_sel1).store(addr_0, ObjectWithRefFieldClass_0_sel1)
    C[6]: ObjectWithRefFieldClass_1.select(addr_1, ObjectWithRefFieldClass_1_sel0).store(addr_0, ObjectWithRefFieldClass_1_sel0)
    C[7]: ObjectWithRefFieldClass_2.select(addr_1, ObjectWithRefFieldClass_2_sel0).store(addr_0, ObjectWithRefFieldClass_2_sel0)
    C[8]: ObjectWithRefFieldClass_3
    C[9]: ObjectWithRefFieldClass_4
}

SAT
Model:
ObjectWithRefFieldClass_0 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
addr_1 -> #xffffffffffffffff
addr_0 -> #x0000000000000000
ObjectWithRefFieldClass_1 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
ObjectWithRefFieldClass_1_sel0 -> #x0000000000000000
ObjectWithRefFieldClass_0_sel1 -> #x0000000000000000
ObjectWithRefFieldClass_0_sel0 -> #x0000000000000000
ObjectWithRefFieldClass_2 -> ((as const (Array (_ BitVec 64) Real)) 0.0)
ObjectWithRefFieldClass_2_sel0 -> 0.0

---- State #1 ----
Interpreter{
  PC: ((true & (!(ObjectWithRefFieldClass_0_sel0 > 0))) & (addr_0 == addr_1))
  ReturnResults:
    [0] 2
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=ObjectWithRefFieldClass, field=-1)
      P[1]: &(addr_1) (refType=struct, type=ObjectWithRefFieldClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithRefFieldClass:
        field[0]: ObjectWithRefFieldClass_0 (elem=int, ops=3, store=1, select=2)
        op[0]: select idx=addr_0, res=ObjectWithRefFieldClass_0_sel0
        op[1]: select idx=addr_1, res=ObjectWithRefFieldClass_0_sel1
        op[2]: store idx=addr_0, val=ObjectWithRefFieldClass_0_sel1
        field[1]: ObjectWithRefFieldClass_1 (elem=int, ops=2, store=1, select=1)
        op[0]: select idx=addr_1, res=ObjectWithRefFieldClass_1_sel0
        op[1]: store idx=addr_0, val=ObjectWithRefFieldClass_1_sel0
        field[2]: ObjectWithRefFieldClass_2 (elem=float, ops=2, store=1, select=1)
        op[0]: select idx=addr_1, res=ObjectWithRefFieldClass_2_sel0
        op[1]: store idx=addr_0, val=ObjectWithRefFieldClass_2_sel0
        field[3]: ObjectWithRefFieldClass_3 (elem=struct, ops=0, store=0, select=0)
        field[4]: ObjectWithRefFieldClass_4 (elem=int, ops=0, store=0, select=0)
  }
  HeapConstraints: 10
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: ObjectWithRefFieldClass_0.select(addr_0, ObjectWithRefFieldClass_0_sel0).select(addr_1, ObjectWithRefFieldClass_0_sel1).store(addr_0, ObjectWithRefFieldClass_0_sel1)
    C[6]: ObjectWithRefFieldClass_1.select(addr_1, ObjectWithRefFieldClass_1_sel0).store(addr_0, ObjectWithRefFieldClass_1_sel0)
    C[7]: ObjectWithRefFieldClass_2.select(addr_1, ObjectWithRefFieldClass_2_sel0).store(addr_0, ObjectWithRefFieldClass_2_sel0)
    C[8]: ObjectWithRefFieldClass_3
    C[9]: ObjectWithRefFieldClass_4
}

SAT
Model:
ObjectWithRefFieldClass_0 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
ObjectWithRefFieldClass_1 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
addr_1 -> #x0000000000000000
ObjectWithRefFieldClass_1_sel0 -> #x0000000000000000
addr_0 -> #x0000000000000000
ObjectWithRefFieldClass_0_sel1 -> #x0000000000000000
ObjectWithRefFieldClass_0_sel0 -> #x0000000000000000
ObjectWithRefFieldClass_2 -> ((as const (Array (_ BitVec 64) Real)) 0.0)
ObjectWithRefFieldClass_2_sel0 -> 0.0

---- State #2 ----
Interpreter{
  PC: (((true & (ObjectWithRefFieldClass_0_sel0 > 0)) & (ObjectWithRefFieldClass_0_sel1 < 0)) & (addr_0 == addr_1))
  ReturnResults:
    [0] 0
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=ObjectWithRefFieldClass, field=-1)
      P[1]: &(addr_1) (refType=struct, type=ObjectWithRefFieldClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithRefFieldClass:
        field[0]: ObjectWithRefFieldClass_0 (elem=int, ops=2, store=0, select=2)
        op[0]: select idx=addr_0, res=ObjectWithRefFieldClass_0_sel0
        op[1]: select idx=addr_1, res=ObjectWithRefFieldClass_0_sel1
        field[1]: ObjectWithRefFieldClass_1 (elem=int, ops=0, store=0, select=0)
        field[2]: ObjectWithRefFieldClass_2 (elem=float, ops=0, store=0, select=0)
        field[3]: ObjectWithRefFieldClass_3 (elem=struct, ops=0, store=0, select=0)
        field[4]: ObjectWithRefFieldClass_4 (elem=int, ops=0, store=0, select=0)
  }
  HeapConstraints: 10
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: ObjectWithRefFieldClass_0.select(addr_0, ObjectWithRefFieldClass_0_sel0).select(addr_1, ObjectWithRefFieldClass_0_sel1)
    C[6]: ObjectWithRefFieldClass_1
    C[7]: ObjectWithRefFieldClass_2
    C[8]: ObjectWithRefFieldClass_3
    C[9]: ObjectWithRefFieldClass_4
}

UNSAT
---- State #3 ----
Interpreter{
  PC: (((true & (ObjectWithRefFieldClass_0_sel0 > 0)) & (ObjectWithRefFieldClass_0_sel1 < 0)) & (!(addr_0 == addr_1)))
  ReturnResults:
    [0] 1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=ObjectWithRefFieldClass, field=-1)
      P[1]: &(addr_1) (refType=struct, type=ObjectWithRefFieldClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithRefFieldClass:
        field[0]: ObjectWithRefFieldClass_0 (elem=int, ops=2, store=0, select=2)
        op[0]: select idx=addr_0, res=ObjectWithRefFieldClass_0_sel0
        op[1]: select idx=addr_1, res=ObjectWithRefFieldClass_0_sel1
        field[1]: ObjectWithRefFieldClass_1 (elem=int, ops=0, store=0, select=0)
        field[2]: ObjectWithRefFieldClass_2 (elem=float, ops=0, store=0, select=0)
        field[3]: ObjectWithRefFieldClass_3 (elem=struct, ops=0, store=0, select=0)
        field[4]: ObjectWithRefFieldClass_4 (elem=int, ops=0, store=0, select=0)
  }
  HeapConstraints: 10
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: ObjectWithRefFieldClass_0.select(addr_0, ObjectWithRefFieldClass_0_sel0).select(addr_1, ObjectWithRefFieldClass_0_sel1)
    C[6]: ObjectWithRefFieldClass_1
    C[7]: ObjectWithRefFieldClass_2
    C[8]: ObjectWithRefFieldClass_3
    C[9]: ObjectWithRefFieldClass_4
}

SAT
Model:
addr_0 -> #xffffffffffffff7f
addr_1 -> #x0000000000000080
ObjectWithRefFieldClass_0 -> (store ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x4000000000000000)
                                    #x0000000000000080
                                    #xbfffffffffffffff)
ObjectWithRefFieldClass_0_sel1 -> #xbfffffffffffffff
ObjectWithRefFieldClass_0_sel0 -> #x4000000000000000

---- State #4 ----
Interpreter{
  PC: (((true & (ObjectWithRefFieldClass_0_sel0 > 0)) & (!(ObjectWithRefFieldClass_0_sel1 < 0))) & (!(addr_0 == addr_1)))
  ReturnResults:
    [0] 3
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=ObjectWithRefFieldClass, field=-1)
      P[1]: &(addr_1) (refType=struct, type=ObjectWithRefFieldClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithRefFieldClass:
        field[0]: ObjectWithRefFieldClass_0 (elem=int, ops=4, store=1, select=3)
        op[0]: select idx=addr_0, res=ObjectWithRefFieldClass_0_sel0
        op[1]: select idx=addr_1, res=ObjectWithRefFieldClass_0_sel1
        op[2]: select idx=addr_1, res=ObjectWithRefFieldClass_0_sel2
        op[3]: store idx=addr_0, val=ObjectWithRefFieldClass_0_sel2
        field[1]: ObjectWithRefFieldClass_1 (elem=int, ops=2, store=1, select=1)
        op[0]: select idx=addr_1, res=ObjectWithRefFieldClass_1_sel0
        op[1]: store idx=addr_0, val=ObjectWithRefFieldClass_1_sel0
        field[2]: ObjectWithRefFieldClass_2 (elem=float, ops=2, store=1, select=1)
        op[0]: select idx=addr_1, res=ObjectWithRefFieldClass_2_sel0
        op[1]: store idx=addr_0, val=ObjectWithRefFieldClass_2_sel0
        field[3]: ObjectWithRefFieldClass_3 (elem=struct, ops=0, store=0, select=0)
        field[4]: ObjectWithRefFieldClass_4 (elem=int, ops=0, store=0, select=0)
  }
  HeapConstraints: 10
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: ObjectWithRefFieldClass_4
    C[6]: ObjectWithRefFieldClass_0.select(addr_0, ObjectWithRefFieldClass_0_sel0).select(addr_1, ObjectWithRefFieldClass_0_sel1).select(addr_1, ObjectWithRefFieldClass_0_sel2).store(addr_0, ObjectWithRefFieldClass_0_sel2)
    C[7]: ObjectWithRefFieldClass_1.select(addr_1, ObjectWithRefFieldClass_1_sel0).store(addr_0, ObjectWithRefFieldClass_1_sel0)
    C[8]: ObjectWithRefFieldClass_2.select(addr_1, ObjectWithRefFieldClass_2_sel0).store(addr_0, ObjectWithRefFieldClass_2_sel0)
    C[9]: ObjectWithRefFieldClass_3
}

SAT
Model:
ObjectWithRefFieldClass_0 -> (store ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000001)
                                    #x0000000000000000
                                    #x0000000000000000)
addr_0 -> #x0000000000000001
addr_1 -> #x0000000000000000
ObjectWithRefFieldClass_0_sel2 -> #x0000000000000000
ObjectWithRefFieldClass_1 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
ObjectWithRefFieldClass_1_sel0 -> #x0000000000000000
ObjectWithRefFieldClass_0_sel1 -> #x0000000000000000
ObjectWithRefFieldClass_0_sel0 -> #x0000000000000001
ObjectWithRefFieldClass_2 -> ((as const (Array (_ BitVec 64) Real)) 0.0)
ObjectWithRefFieldClass_2_sel0 -> 0.0

---- State #5 ----
Interpreter{
  PC: (((true & (ObjectWithRefFieldClass_0_sel0 > 0)) & (!(ObjectWithRefFieldClass_0_sel1 < 0))) & (addr_0 == addr_1))
  ReturnResults:
    [0] 2
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=ObjectWithRefFieldClass, field=-1)
      P[1]: &(addr_1) (refType=struct, type=ObjectWithRefFieldClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      ObjectWithRefFieldClass:
        field[0]: ObjectWithRefFieldClass_0 (elem=int, ops=4, store=1, select=3)
        op[0]: select idx=addr_0, res=ObjectWithRefFieldClass_0_sel0
        op[1]: select idx=addr_1, res=ObjectWithRefFieldClass_0_sel1
        op[2]: select idx=addr_1, res=ObjectWithRefFieldClass_0_sel2
        op[3]: store idx=addr_0, val=ObjectWithRefFieldClass_0_sel2
        field[1]: ObjectWithRefFieldClass_1 (elem=int, ops=2, store=1, select=1)
        op[0]: select idx=addr_1, res=ObjectWithRefFieldClass_1_sel0
        op[1]: store idx=addr_0, val=ObjectWithRefFieldClass_1_sel0
        field[2]: ObjectWithRefFieldClass_2 (elem=float, ops=2, store=1, select=1)
        op[0]: select idx=addr_1, res=ObjectWithRefFieldClass_2_sel0
        op[1]: store idx=addr_0, val=ObjectWithRefFieldClass_2_sel0
        field[3]: ObjectWithRefFieldClass_3 (elem=struct, ops=0, store=0, select=0)
        field[4]: ObjectWithRefFieldClass_4 (elem=int, ops=0, store=0, select=0)
  }
  HeapConstraints: 10
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: ObjectWithRefFieldClass_0.select(addr_0, ObjectWithRefFieldClass_0_sel0).select(addr_1, ObjectWithRefFieldClass_0_sel1).select(addr_1, ObjectWithRefFieldClass_0_sel2).store(addr_0, ObjectWithRefFieldClass_0_sel2)
    C[6]: ObjectWithRefFieldClass_1.select(addr_1, ObjectWithRefFieldClass_1_sel0).store(addr_0, ObjectWithRefFieldClass_1_sel0)
    C[7]: ObjectWithRefFieldClass_2.select(addr_1, ObjectWithRefFieldClass_2_sel0).store(addr_0, ObjectWithRefFieldClass_2_sel0)
    C[8]: ObjectWithRefFieldClass_3
    C[9]: ObjectWithRefFieldClass_4
}

SAT
Model:
ObjectWithRefFieldClass_0 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000001)
addr_1 -> #x0000000000000000
ObjectWithRefFieldClass_0_sel2 -> #x0000000000000001
ObjectWithRefFieldClass_1 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
ObjectWithRefFieldClass_1_sel0 -> #x0000000000000000
addr_0 -> #x0000000000000000
ObjectWithRefFieldClass_0_sel1 -> #x0000000000000001
ObjectWithRefFieldClass_0_sel0 -> #x0000000000000001
ObjectWithRefFieldClass_2 -> ((as const (Array (_ BitVec 64) Real)) 0.0)
ObjectWithRefFieldClass_2_sel0 -> 0.0


==================== CompareTwoObjectsWithTheSameRefField ====================
CFG построен для функции с 3 блоками
interpretDynamically: new SimpleDataClass (complit)
interpretDynamically: &fst.RefField [#3]
interpretDynamically: *t1 = t0
resolveExpression: new SimpleDataClass (complit)
resolveExpression: parameter fst : *ObjectWithRefFieldClass
interpretDynamically: &snd.RefField [#3]
interpretDynamically: *t2 = t0
resolveExpression: new SimpleDataClass (complit)
resolveExpression: parameter snd : *ObjectWithRefFieldClass
interpretDynamically: &snd.x [#0]
interpretDynamically: *t3
interpretDynamically: &fst.x [#0]
interpretDynamically: *t5 = t4
resolveExpression: *t3
resolveExpression: parameter snd : *ObjectWithRefFieldClass
resolveExpression: parameter fst : *ObjectWithRefFieldClass
interpretDynamically: &snd.y [#1]
interpretDynamically: *t6
interpretDynamically: &fst.y [#1]
interpretDynamically: *t8 = t7
resolveExpression: *t6
resolveExpression: parameter snd : *ObjectWithRefFieldClass
resolveExpression: parameter fst : *ObjectWithRefFieldClass
interpretDynamically: &snd.Weight [#2]
interpretDynamically: *t9
interpretDynamically: &fst.Weight [#2]
interpretDynamically: *t11 = t10
resolveExpression: *t9
resolveExpression: parameter snd : *ObjectWithRefFieldClass
resolveExpression: parameter fst : *ObjectWithRefFieldClass
interpretDynamically: fst == snd
interpretDynamically: if t12 goto 1 else 2
resolveExpression: fst == snd
resolveExpression: parameter fst : *ObjectWithRefFieldClass
resolveExpression: parameter snd : *ObjectWithRefFieldClass
interpretDynamically: return 2:int
resolveExpression: 2:int
interpretDynamically: return 1:int
resolveExpression: 1:int
States: 2
---- State #0 ----
Interpreter{
  PC: (true & (!(addr_1 == addr_2)))
  ReturnResults:
    [0] 2
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_1) (refType=struct, type=ObjectWithRefFieldClass, field=-1)
      P[1]: &(addr_2) (refType=struct, type=ObjectWithRefFieldClass, field=-1)
    AllocationsList:
      A[0]: &(addr_0) (refType=struct, type=SimpleDataClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 2 type(s)
      ObjectWithRefFieldClass:
        field[0]: ObjectWithRefFieldClass_0 (elem=int, ops=2, store=1, select=1)
        op[0]: select idx=addr_2, res=ObjectWithRefFieldClass_0_sel0
        op[1]: store idx=addr_1, val=ObjectWithRefFieldClass_0_sel0
        field[1]: ObjectWithRefFieldClass_1 (elem=int, ops=2, store=1, select=1)
        op[0]: select idx=addr_2, res=ObjectWithRefFieldClass_1_sel0
        op[1]: store idx=addr_1, val=ObjectWithRefFieldClass_1_sel0
        field[2]: ObjectWithRefFieldClass_2 (elem=float, ops=2, store=1, select=1)
        op[0]: select idx=addr_2, res=ObjectWithRefFieldClass_2_sel0
        op[1]: store idx=addr_1, val=ObjectWithRefFieldClass_2_sel0
        field[3]: ObjectWithRefFieldClass_3 (elem=struct, ops=2, store=2, select=0)
        op[0]: store idx=addr_1, val=&(addr_0)
        op[1]: store idx=addr_2, val=&(addr_0)
        field[4]: ObjectWithRefFieldClass_4 (elem=int, ops=0, store=0, select=0)
      SimpleDataClass:
        field[0]: SimpleDataClass_0 (elem=int, ops=0, store=0, select=0)
        field[1]: SimpleDataClass_1 (elem=int, ops=0, store=0, select=0)
  }
  HeapConstraints: 14
    C[0]: (addr_1 != addr_0)
    C[1]: (addr_2 != addr_0)
    C[2]: IntSegment
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment
    C[7]: SimpleDataClass_0
    C[8]: SimpleDataClass_1
    C[9]: ObjectWithRefFieldClass_2.select(addr_2, ObjectWithRefFieldClass_2_sel0).store(addr_1, ObjectWithRefFieldClass_2_sel0)
    C[10]: ObjectWithRefFieldClass_3.store(addr_1, &(addr_0)).store(addr_2, &(addr_0))
    C[11]: ObjectWithRefFieldClass_4
    C[12]: ObjectWithRefFieldClass_0.select(addr_2, ObjectWithRefFieldClass_0_sel0).store(addr_1, ObjectWithRefFieldClass_0_sel0)
    C[13]: ObjectWithRefFieldClass_1.select(addr_2, ObjectWithRefFieldClass_1_sel0).store(addr_1, ObjectWithRefFieldClass_1_sel0)
}

SAT
Model:
addr_2 -> #x0000000000000001
addr_1 -> #x0000000000000002
addr_0 -> #x0000000000000000
ObjectWithRefFieldClass_1 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
ObjectWithRefFieldClass_1_sel0 -> #x0000000000000000
ObjectWithRefFieldClass_0 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
ObjectWithRefFieldClass_0_sel0 -> #x0000000000000000
ObjectWithRefFieldClass_2 -> ((as const (Array (_ BitVec 64) Real)) 0.0)
ObjectWithRefFieldClass_2_sel0 -> 0.0

---- State #1 ----
Interpreter{
  PC: (true & (addr_1 == addr_2))
  ReturnResults:
    [0] 1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 3
    Params: 2, Allocations: 1
    ParamsList:
      P[0]: &(addr_1) (refType=struct, type=ObjectWithRefFieldClass, field=-1)
      P[1]: &(addr_2) (refType=struct, type=ObjectWithRefFieldClass, field=-1)
    AllocationsList:
      A[0]: &(addr_0) (refType=struct, type=SimpleDataClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 2 type(s)
      ObjectWithRefFieldClass:
        field[0]: ObjectWithRefFieldClass_0 (elem=int, ops=2, store=1, select=1)
        op[0]: select idx=addr_2, res=ObjectWithRefFieldClass_0_sel0
        op[1]: store idx=addr_1, val=ObjectWithRefFieldClass_0_sel0
        field[1]: ObjectWithRefFieldClass_1 (elem=int, ops=2, store=1, select=1)
        op[0]: select idx=addr_2, res=ObjectWithRefFieldClass_1_sel0
        op[1]: store idx=addr_1, val=ObjectWithRefFieldClass_1_sel0
        field[2]: ObjectWithRefFieldClass_2 (elem=float, ops=2, store=1, select=1)
        op[0]: select idx=addr_2, res=ObjectWithRefFieldClass_2_sel0
        op[1]: store idx=addr_1, val=ObjectWithRefFieldClass_2_sel0
        field[3]: ObjectWithRefFieldClass_3 (elem=struct, ops=2, store=2, select=0)
        op[0]: store idx=addr_1, val=&(addr_0)
        op[1]: store idx=addr_2, val=&(addr_0)
        field[4]: ObjectWithRefFieldClass_4 (elem=int, ops=0, store=0, select=0)
      SimpleDataClass:
        field[0]: SimpleDataClass_0 (elem=int, ops=0, store=0, select=0)
        field[1]: SimpleDataClass_1 (elem=int, ops=0, store=0, select=0)
  }
  HeapConstraints: 14
    C[0]: (addr_1 != addr_0)
    C[1]: (addr_2 != addr_0)
    C[2]: IntSegment
    C[3]: BoolSegment
    C[4]: StringSegment
    C[5]: FloatSegment
    C[6]: ArraySegment
    C[7]: SimpleDataClass_0
    C[8]: SimpleDataClass_1
    C[9]: ObjectWithRefFieldClass_3.store(addr_1, &(addr_0)).store(addr_2, &(addr_0))
    C[10]: ObjectWithRefFieldClass_4
    C[11]: ObjectWithRefFieldClass_0.select(addr_2, ObjectWithRefFieldClass_0_sel0).store(addr_1, ObjectWithRefFieldClass_0_sel0)
    C[12]: ObjectWithRefFieldClass_1.select(addr_2, ObjectWithRefFieldClass_1_sel0).store(addr_1, ObjectWithRefFieldClass_1_sel0)
    C[13]: ObjectWithRefFieldClass_2.select(addr_2, ObjectWithRefFieldClass_2_sel0).store(addr_1, ObjectWithRefFieldClass_2_sel0)
}

SAT
Model:
addr_2 -> #xffffffffffffffff
addr_0 -> #x0000000000000000
ObjectWithRefFieldClass_1 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
ObjectWithRefFieldClass_1_sel0 -> #x0000000000000000
addr_1 -> #xffffffffffffffff
ObjectWithRefFieldClass_0 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
ObjectWithRefFieldClass_0_sel0 -> #x0000000000000000
ObjectWithRefFieldClass_2 -> ((as const (Array (_ BitVec 64) Real)) 0.0)
ObjectWithRefFieldClass_2_sel0 -> 0.0


==================== NextValue ====================
CFG построен для функции с 6 блоками
interpretDynamically: value == 0:int
interpretDynamically: if t0 goto 1 else 2
resolveExpression: value == 0:int
resolveExpression: parameter value : int
resolveExpression: 0:int
interpretDynamically: &node.Next [#0]
interpretDynamically: *t1
interpretDynamically: t2 != nil:*RecursiveTypeClass
interpretDynamically: if t3 goto 5 else 4
resolveExpression: t2 != nil:*RecursiveTypeClass
resolveExpression: *t1
resolveExpression: parameter node : *RecursiveTypeClass
resolveExpression: nil:*main.RecursiveTypeClass
interpretDynamically: &node.Next [#0]
interpretDynamically: *t6
interpretDynamically: &t7.Value [#1]
interpretDynamically: return nil:*RecursiveTypeClass
resolveExpression: nil:*main.RecursiveTypeClass
interpretDynamically: *t8
interpretDynamically: t9 == value
interpretDynamically: if t10 goto 3 else 4
resolveExpression: t9 == value
resolveExpression: *t8
resolveExpression: *t6
resolveExpression: parameter node : *RecursiveTypeClass
resolveExpression: parameter value : int
interpretDynamically: &node.Next [#0]
interpretDynamically: *t4
interpretDynamically: return t5
resolveExpression: *t4
resolveExpression: parameter node : *RecursiveTypeClass
interpretDynamically: return nil:*RecursiveTypeClass
resolveExpression: nil:*main.RecursiveTypeClass
interpretDynamically: return nil:*RecursiveTypeClass
resolveExpression: nil:*main.RecursiveTypeClass
States: 4
---- State #0 ----
Interpreter{
  PC: (true & (IntSegment_sel0 == 0))
  ReturnResults:
    [0] nil
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
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
IntSegment_sel0 -> #x0000000000000000

---- State #1 ----
Interpreter{
  PC: (((true & (!(IntSegment_sel0 == 0))) & (RecursiveTypeClass_0_sel0 != &(addr_0))) & (!(RecursiveTypeClass_1_sel0 == IntSegment_sel1)))
  ReturnResults:
    [0] nil
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=struct, type=RecursiveTypeClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=2, store=0, select=2)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      RecursiveTypeClass:
        field[0]: RecursiveTypeClass_0 (elem=struct, ops=2, store=0, select=2)
        op[0]: select idx=addr_1, res=RecursiveTypeClass_0_sel0
        op[1]: select idx=addr_1, res=RecursiveTypeClass_0_sel1
        field[1]: RecursiveTypeClass_1 (elem=int, ops=1, store=0, select=1)
        op[0]: select idx=RecursiveTypeClass_0_sel1, res=RecursiveTypeClass_1_sel0
  }
  HeapConstraints: 7
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: RecursiveTypeClass_0.select(addr_1, RecursiveTypeClass_0_sel0).select(addr_1, RecursiveTypeClass_0_sel1)
    C[6]: RecursiveTypeClass_1.select(RecursiveTypeClass_0_sel1, RecursiveTypeClass_1_sel0)
}

SAT
Model:
RecursiveTypeClass_1 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
RecursiveTypeClass_0 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #xffffffffffffffff)
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #xffffffffffffffff)
addr_0 -> #x0000000000000000
addr_1 -> #x0000000000000000
RecursiveTypeClass_0_sel1 -> #xffffffffffffffff
RecursiveTypeClass_0_sel0 -> #xffffffffffffffff
IntSegment_sel1 -> #xffffffffffffffff
RecursiveTypeClass_1_sel0 -> #x0000000000000000
IntSegment_sel0 -> #xffffffffffffffff

---- State #2 ----
Interpreter{
  PC: (((true & (!(IntSegment_sel0 == 0))) & (RecursiveTypeClass_0_sel0 != &(addr_0))) & (RecursiveTypeClass_1_sel0 == IntSegment_sel1))
  ReturnResults:
    [0] &(RecursiveTypeClass_0_sel2)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=struct, type=RecursiveTypeClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=2, store=0, select=2)
        op[0]: select idx=addr_0, res=IntSegment_sel0
        op[1]: select idx=addr_0, res=IntSegment_sel1
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      RecursiveTypeClass:
        field[0]: RecursiveTypeClass_0 (elem=struct, ops=3, store=0, select=3)
        op[0]: select idx=addr_1, res=RecursiveTypeClass_0_sel0
        op[1]: select idx=addr_1, res=RecursiveTypeClass_0_sel1
        op[2]: select idx=addr_1, res=RecursiveTypeClass_0_sel2
        field[1]: RecursiveTypeClass_1 (elem=int, ops=1, store=0, select=1)
        op[0]: select idx=RecursiveTypeClass_0_sel1, res=RecursiveTypeClass_1_sel0
  }
  HeapConstraints: 7
    C[0]: IntSegment.select(addr_0, IntSegment_sel0).select(addr_0, IntSegment_sel1)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: RecursiveTypeClass_0.select(addr_1, RecursiveTypeClass_0_sel0).select(addr_1, RecursiveTypeClass_0_sel1).select(addr_1, RecursiveTypeClass_0_sel2)
    C[6]: RecursiveTypeClass_1.select(RecursiveTypeClass_0_sel1, RecursiveTypeClass_1_sel0)
}

SAT
Model:
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #xffffffffffffffff)
RecursiveTypeClass_0 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #xffffffffffffffff)
RecursiveTypeClass_1 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #xffffffffffffffff)
addr_0 -> #x0000000000000000
addr_1 -> #x0000000000000000
RecursiveTypeClass_0_sel2 -> #xffffffffffffffff
RecursiveTypeClass_0_sel1 -> #xffffffffffffffff
RecursiveTypeClass_0_sel0 -> #xffffffffffffffff
IntSegment_sel1 -> #xffffffffffffffff
RecursiveTypeClass_1_sel0 -> #xffffffffffffffff
IntSegment_sel0 -> #xffffffffffffffff

---- State #3 ----
Interpreter{
  PC: ((true & (!(IntSegment_sel0 == 0))) & (!(RecursiveTypeClass_0_sel0 != &(addr_0))))
  ReturnResults:
    [0] nil
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 2, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=int)
      P[1]: &(addr_1) (refType=struct, type=RecursiveTypeClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=1, store=0, select=1)
        op[0]: select idx=addr_0, res=IntSegment_sel0
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      RecursiveTypeClass:
        field[0]: RecursiveTypeClass_0 (elem=struct, ops=1, store=0, select=1)
        op[0]: select idx=addr_1, res=RecursiveTypeClass_0_sel0
        field[1]: RecursiveTypeClass_1 (elem=int, ops=0, store=0, select=0)
  }
  HeapConstraints: 7
    C[0]: IntSegment.select(addr_0, IntSegment_sel0)
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: RecursiveTypeClass_0.select(addr_1, RecursiveTypeClass_0_sel0)
    C[6]: RecursiveTypeClass_1
}

SAT
Model:
IntSegment -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #xffffffffffffffff)
RecursiveTypeClass_0 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
addr_1 -> #x0000000000000000
addr_0 -> #x0000000000000000
RecursiveTypeClass_0_sel0 -> #x0000000000000000
IntSegment_sel0 -> #xffffffffffffffff


==================== WriteObjectField ====================
CFG построен для функции с 3 блоками
interpretDynamically: &node.Next [#0]
interpretDynamically: *t0
interpretDynamically: t1 == nil:*RecursiveTypeClass
interpretDynamically: if t2 goto 1 else 2
resolveExpression: t1 == nil:*RecursiveTypeClass
resolveExpression: *t0
resolveExpression: parameter node : *RecursiveTypeClass
resolveExpression: nil:*main.RecursiveTypeClass
interpretDynamically: new RecursiveTypeClass (complit)
interpretDynamically: &node.Next [#0]
interpretDynamically: *t4 = t3
resolveExpression: new RecursiveTypeClass (complit)
resolveExpression: parameter node : *RecursiveTypeClass
interpretDynamically: jump 2
interpretDynamically: &node.Next [#0]
interpretDynamically: *t5
interpretDynamically: &node.Next [#0]
interpretDynamically: &node.Next [#0]
interpretDynamically: *t5
interpretDynamically: &node.Next [#0]
interpretDynamically: *t7
interpretDynamically: *t7
interpretDynamically: &t8.Value [#1]
interpretDynamically: *t9
interpretDynamically: t10 + 1:int
interpretDynamically: &t6.Value [#1]
interpretDynamically: *t12 = t11
resolveExpression: t10 + 1:int
resolveExpression: *t9
resolveExpression: *t7
resolveExpression: parameter node : *RecursiveTypeClass
resolveExpression: 1:int
resolveExpression: *t5
resolveExpression: parameter node : *RecursiveTypeClass
interpretDynamically: return node
resolveExpression: parameter node : *RecursiveTypeClass
interpretDynamically: &t8.Value [#1]
interpretDynamically: *t9
interpretDynamically: t10 + 1:int
interpretDynamically: &t6.Value [#1]
interpretDynamically: *t12 = t11
resolveExpression: t10 + 1:int
resolveExpression: *t9
resolveExpression: *t7
resolveExpression: parameter node : *RecursiveTypeClass
resolveExpression: 1:int
resolveExpression: *t5
resolveExpression: parameter node : *RecursiveTypeClass
interpretDynamically: return node
resolveExpression: parameter node : *RecursiveTypeClass
States: 2
---- State #0 ----
Interpreter{
  PC: (true & (RecursiveTypeClass_0_sel0 == &(addr_0)))
  ReturnResults:
    [0] &(addr_0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 1, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=RecursiveTypeClass, field=-1)
    AllocationsList:
      A[0]: &(addr_1) (refType=struct, type=RecursiveTypeClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      RecursiveTypeClass:
        field[0]: RecursiveTypeClass_0 (elem=struct, ops=4, store=1, select=3)
        op[0]: select idx=addr_0, res=RecursiveTypeClass_0_sel0
        op[1]: store idx=addr_0, val=&(addr_1)
        op[2]: select idx=addr_0, res=RecursiveTypeClass_0_sel2
        op[3]: select idx=addr_0, res=RecursiveTypeClass_0_sel3
        field[1]: RecursiveTypeClass_1 (elem=int, ops=2, store=1, select=1)
        op[0]: select idx=RecursiveTypeClass_0_sel2, res=RecursiveTypeClass_1_sel0
        op[1]: store idx=RecursiveTypeClass_0_sel3, val=(RecursiveTypeClass_1_sel0 + 1)
  }
  HeapConstraints: 8
    C[0]: (addr_0 != addr_1)
    C[1]: IntSegment
    C[2]: BoolSegment
    C[3]: StringSegment
    C[4]: FloatSegment
    C[5]: ArraySegment
    C[6]: RecursiveTypeClass_0.select(addr_0, RecursiveTypeClass_0_sel0).store(addr_0, &(addr_1)).select(addr_0, RecursiveTypeClass_0_sel2).select(addr_0, RecursiveTypeClass_0_sel3)
    C[7]: RecursiveTypeClass_1.select(RecursiveTypeClass_0_sel2, RecursiveTypeClass_1_sel0).store(RecursiveTypeClass_0_sel3, (RecursiveTypeClass_1_sel0 + 1))
}

SAT
Model:
addr_0 -> #x0000000000000000
RecursiveTypeClass_0_sel3 -> #xffffffffffffffff
RecursiveTypeClass_0 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
addr_1 -> #xffffffffffffffff
RecursiveTypeClass_0_sel2 -> #xffffffffffffffff
RecursiveTypeClass_0_sel0 -> #x0000000000000000
RecursiveTypeClass_1 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
RecursiveTypeClass_1_sel0 -> #x0000000000000000

---- State #1 ----
Interpreter{
  PC: (true & (!(RecursiveTypeClass_0_sel0 == &(addr_0))))
  ReturnResults:
    [0] &(addr_0)
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 1
    Params: 1, Allocations: 0
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=RecursiveTypeClass, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      RecursiveTypeClass:
        field[0]: RecursiveTypeClass_0 (elem=struct, ops=3, store=0, select=3)
        op[0]: select idx=addr_0, res=RecursiveTypeClass_0_sel0
        op[1]: select idx=addr_0, res=RecursiveTypeClass_0_sel1
        op[2]: select idx=addr_0, res=RecursiveTypeClass_0_sel2
        field[1]: RecursiveTypeClass_1 (elem=int, ops=2, store=1, select=1)
        op[0]: select idx=RecursiveTypeClass_0_sel1, res=RecursiveTypeClass_1_sel0
        op[1]: store idx=RecursiveTypeClass_0_sel2, val=(RecursiveTypeClass_1_sel0 + 1)
  }
  HeapConstraints: 7
    C[0]: IntSegment
    C[1]: BoolSegment
    C[2]: StringSegment
    C[3]: FloatSegment
    C[4]: ArraySegment
    C[5]: RecursiveTypeClass_0.select(addr_0, RecursiveTypeClass_0_sel0).select(addr_0, RecursiveTypeClass_0_sel1).select(addr_0, RecursiveTypeClass_0_sel2)
    C[6]: RecursiveTypeClass_1.select(RecursiveTypeClass_0_sel1, RecursiveTypeClass_1_sel0).store(RecursiveTypeClass_0_sel2, (RecursiveTypeClass_1_sel0 + 1))
}

SAT
Model:
RecursiveTypeClass_0 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #xffffffffffffffff)
addr_0 -> #x0000000000000000
RecursiveTypeClass_0_sel2 -> #xffffffffffffffff
RecursiveTypeClass_0_sel0 -> #xffffffffffffffff
RecursiveTypeClass_0_sel1 -> #xffffffffffffffff
RecursiveTypeClass_1 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
RecursiveTypeClass_1_sel0 -> #x0000000000000000


==================== TestPathConstraintMutability ====================
CFG построен для функции с 5 блоками
interpretDynamically: local Person (p)
interpretDynamically: *t0 = p
resolveExpression: parameter p : Person
resolveExpression: local Person (p)
interpretDynamically: &t0.Age [#1]
interpretDynamically: *t1
interpretDynamically: t2 != 18:int
interpretDynamically: if t3 goto 1 else 2
resolveExpression: t2 != 18:int
resolveExpression: *t1
resolveExpression: local Person (p)
resolveExpression: 18:int
interpretDynamically: &t0.Age [#1]
interpretDynamically: *t4 = 18:int
resolveExpression: 18:int
resolveExpression: local Person (p)
interpretDynamically: &t0.Age [#1]
interpretDynamically: *t5
interpretDynamically: t6 != 18:int
interpretDynamically: if t7 goto 3 else 4
resolveExpression: t6 != 18:int
resolveExpression: *t5
resolveExpression: local Person (p)
resolveExpression: 18:int
interpretDynamically: return 0:int
resolveExpression: 0:int
interpretDynamically: return 1:int
resolveExpression: 1:int
interpretDynamically: return 2:int
resolveExpression: 2:int
States: 3
---- State #0 ----
Interpreter{
  PC: ((true & (Person_1_sel2 != 18)) & (Person_1_sel4 != 18))
  ReturnResults:
    [0] 0
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 1, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=Person, field=-1)
    AllocationsList:
      A[0]: &(addr_1) (refType=struct, type=Person, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      Person:
        field[0]: Person_0 (elem=string, ops=2, store=1, select=1)
        op[0]: select idx=addr_0, res=Person_0_sel0
        op[1]: store idx=addr_1, val=Person_0_sel0
        field[1]: Person_1 (elem=int, ops=5, store=2, select=3)
        op[0]: select idx=addr_0, res=Person_1_sel0
        op[1]: store idx=addr_1, val=Person_1_sel0
        op[2]: select idx=addr_1, res=Person_1_sel2
        op[3]: store idx=addr_1, val=18
        op[4]: select idx=addr_1, res=Person_1_sel4
        field[2]: Person_2 (elem=int, ops=2, store=1, select=1)
        op[0]: select idx=addr_0, res=Person_2_sel0
        op[1]: store idx=addr_1, val=Person_2_sel0
  }
  HeapConstraints: 9
    C[0]: (addr_0 != addr_1)
    C[1]: IntSegment
    C[2]: BoolSegment
    C[3]: StringSegment
    C[4]: FloatSegment
    C[5]: ArraySegment
    C[6]: Person_0.select(addr_0, Person_0_sel0).store(addr_1, Person_0_sel0)
    C[7]: Person_1.select(addr_0, Person_1_sel0).store(addr_1, Person_1_sel0).select(addr_1, Person_1_sel2).store(addr_1, 18).select(addr_1, Person_1_sel4)
    C[8]: Person_2.select(addr_0, Person_2_sel0).store(addr_1, Person_2_sel0)
}

UNSAT
---- State #1 ----
Interpreter{
  PC: ((true & (Person_1_sel2 != 18)) & (!(Person_1_sel4 != 18)))
  ReturnResults:
    [0] 1
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 1, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=Person, field=-1)
    AllocationsList:
      A[0]: &(addr_1) (refType=struct, type=Person, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      Person:
        field[0]: Person_0 (elem=string, ops=2, store=1, select=1)
        op[0]: select idx=addr_0, res=Person_0_sel0
        op[1]: store idx=addr_1, val=Person_0_sel0
        field[1]: Person_1 (elem=int, ops=5, store=2, select=3)
        op[0]: select idx=addr_0, res=Person_1_sel0
        op[1]: store idx=addr_1, val=Person_1_sel0
        op[2]: select idx=addr_1, res=Person_1_sel2
        op[3]: store idx=addr_1, val=18
        op[4]: select idx=addr_1, res=Person_1_sel4
        field[2]: Person_2 (elem=int, ops=2, store=1, select=1)
        op[0]: select idx=addr_0, res=Person_2_sel0
        op[1]: store idx=addr_1, val=Person_2_sel0
  }
  HeapConstraints: 9
    C[0]: (addr_0 != addr_1)
    C[1]: IntSegment
    C[2]: BoolSegment
    C[3]: StringSegment
    C[4]: FloatSegment
    C[5]: ArraySegment
    C[6]: Person_0.select(addr_0, Person_0_sel0).store(addr_1, Person_0_sel0)
    C[7]: Person_1.select(addr_0, Person_1_sel0).store(addr_1, Person_1_sel0).select(addr_1, Person_1_sel2).store(addr_1, 18).select(addr_1, Person_1_sel4)
    C[8]: Person_2.select(addr_0, Person_2_sel0).store(addr_1, Person_2_sel0)
}

SAT
Model:
addr_1 -> #xffffffffffffffff
Person_1 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #xffffffffffffffed)
addr_0 -> #x0000000000000000
Person_1_sel0 -> #xffffffffffffffed
Person_0 -> ((as const (Array (_ BitVec 64) String)) String!val!0)
Person_0_sel0 -> String!val!0
Person_1_sel2 -> #xffffffffffffffed
Person_2 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
Person_2_sel0 -> #x0000000000000000
Person_1_sel4 -> #x0000000000000012

---- State #2 ----
Interpreter{
  PC: (true & (!(Person_1_sel2 != 18)))
  ReturnResults:
    [0] 2
  CallStack: 0 frame(s)
  Heap: SymbolicMemory{
    LastRefIdx: 2
    Params: 1, Allocations: 1
    ParamsList:
      P[0]: &(addr_0) (refType=struct, type=Person, field=-1)
    AllocationsList:
      A[0]: &(addr_1) (refType=struct, type=Person, field=-1)
    IntSegment:    IntSegment (elem=int, ops=0, store=0, select=0)
    BoolSegment:   BoolSegment (elem=bool, ops=0, store=0, select=0)
    StringSegment: StringSegment (elem=string, ops=0, store=0, select=0)
    ArraySegment:  ArraySegment (elem=int, ops=0, store=0, select=0)
    StructSegments: 1 type(s)
      Person:
        field[0]: Person_0 (elem=string, ops=2, store=1, select=1)
        op[0]: select idx=addr_0, res=Person_0_sel0
        op[1]: store idx=addr_1, val=Person_0_sel0
        field[1]: Person_1 (elem=int, ops=3, store=1, select=2)
        op[0]: select idx=addr_0, res=Person_1_sel0
        op[1]: store idx=addr_1, val=Person_1_sel0
        op[2]: select idx=addr_1, res=Person_1_sel2
        field[2]: Person_2 (elem=int, ops=2, store=1, select=1)
        op[0]: select idx=addr_0, res=Person_2_sel0
        op[1]: store idx=addr_1, val=Person_2_sel0
  }
  HeapConstraints: 9
    C[0]: (addr_0 != addr_1)
    C[1]: IntSegment
    C[2]: BoolSegment
    C[3]: StringSegment
    C[4]: FloatSegment
    C[5]: ArraySegment
    C[6]: Person_1.select(addr_0, Person_1_sel0).store(addr_1, Person_1_sel0).select(addr_1, Person_1_sel2)
    C[7]: Person_2.select(addr_0, Person_2_sel0).store(addr_1, Person_2_sel0)
    C[8]: Person_0.select(addr_0, Person_0_sel0).store(addr_1, Person_0_sel0)
}

SAT
Model:
addr_1 -> #xffffffffffffffff
Person_1 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000012)
addr_0 -> #x0000000000000000
Person_1_sel0 -> #x0000000000000012
Person_0 -> ((as const (Array (_ BitVec 64) String)) String!val!0)
Person_0_sel0 -> String!val!0
Person_1_sel2 -> #x0000000000000012
Person_2 -> ((as const (Array (_ BitVec 64) (_ BitVec 64))) #x0000000000000000)
Person_2_sel0 -> #x0000000000000000

