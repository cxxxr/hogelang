package vm

import (
	"errors"
	"fmt"
	"github.com/cxxxr/hogelang/ast"
)

type Machine struct {
	exe         *Executable
	code        []Instr
	constTable  []Object
	env         *Env
	globalTable []Object
	stack       []Object
	stackTop    int
	pc          int
	err         error
}

func NewMachine(exe *Executable) *Machine {
	return &Machine{
		exe:         exe,
		globalTable: make([]Object, 0),
	}
}

func (m *Machine) SetExecutable(exe *Executable) {
	m.exe = exe
}

func (m *Machine) setProgram(code []Instr, pc int, env *Env) {
	m.code = code
	m.pc = pc
	m.env = env
}

func (m *Machine) pos() *ast.Pos {
	return m.code[m.pc].Pos
}

func (m *Machine) throw(err error) {
	m.err = err
	panic(err)
}

func (m *Machine) jump(pc int) {
	m.pc = pc
}

func (m *Machine) getConst(i int) Object {
	return m.constTable[i]
}

func (m *Machine) setGlobal(i int, v Object) {
	m.globalTable[i] = v
}

func (m *Machine) addGlobal(v Object) {
	m.globalTable = append(m.globalTable, v)
}

func (m *Machine) getGlobal(i int) (Object, bool) {
	v := m.globalTable[i]
	return v, v != nil
}

func (m *Machine) copyToHeapEnv() *Env {
	var bottomEnv *Env = nil
	var env2 *Env = nil
	var prev *Env = nil
	for env := m.env; env != nil; env = env.parent {
		values := make([]Object, env.length)
		copy(values, m.stack[env.sp:env.sp+env.length])
		env2 = NewHeapEnv(nil, values)
		if prev != nil {
			prev.parent = env2
		} else {
			bottomEnv = env2
		}
		prev = env2
	}

	return bottomEnv
}

func (m *Machine) getLocal(i, j int) Object {
	env := m.env
	for n := 0; n < i; n++ {
		env = env.parent
	}

	if env.sp < 0 {
		return env.values[j]
	} else {
		return m.stack[env.sp+j]
	}
}

func (m *Machine) setLocal(i int, j int, v Object) {
	env := m.env
	for n := 0; n < i; n++ {
		env = env.parent
	}
	if env.sp < 0 {
		env.values[j] = v
	} else {
		m.stack[env.sp+j] = v
	}
}

func (m *Machine) codeAhead() Opcode {
	return m.code[m.pc].Opcode
}

func (m *Machine) arg1() int {
	return m.code[m.pc].Arg1
}

func (m *Machine) arg2() int {
	return m.code[m.pc].Arg2
}

func (m *Machine) tos() Object {
	return m.stack[m.stackTop-1]
}

func (m *Machine) tos1() Object {
	return m.stack[m.stackTop-2]
}

func (m *Machine) tos2() Object {
	return m.stack[m.stackTop-3]
}

func (m *Machine) setTos(v Object) {
	m.stack[m.stackTop-1] = v
}

func (m *Machine) setTos1(v Object) {
	m.stack[m.stackTop-2] = v
	m.stackTop--
}

func (m *Machine) setTos2(v Object) {
	m.stack[m.stackTop-3] = v
	m.stackTop -= 2
}

func (m *Machine) pop() Object {
	v := m.tos()
	m.stackTop--
	return v
}

func (m *Machine) popN(n int) {
	m.stackTop -= n
}

func (m *Machine) push(v Object) {
	m.stack[m.stackTop] = v
	m.stackTop++
}

func (m *Machine) step() {
	m.pc++
}

func assertInt(m *Machine, v Object, types int) int {
	i, ok := v.(Fixnum)
	if ok {
		return int(i)
	} else {
		typeError(m, v, types)
		return 0
	}
}

func assertFixnum(m *Machine, v Object, types int) Fixnum {
	i, ok := v.(Fixnum)
	if ok {
		return i
	} else {
		typeError(m, v, types)
		return 0
	}
}

func raise(m *Machine, msg string) {
	pos := m.pos()
	m.throw(errors.New(fmt.Sprintf("error:%d:%d: %s", pos.Line, pos.Offset, msg)))
}

func typeError(m *Machine, v Object, types int) {
	raise(m, fmt.Sprintf("%s not in (%s)", v, typesToString(types, ", ")))
}

func indexError(m *Machine, seq Object, i int) {
	raise(m, fmt.Sprintf("index error %s %d", seq, i))
}

func unboundVariableError(m *Machine, i int) {
	raise(m, fmt.Sprintf("unbound varialbe '%s'", m.exe.globalVarName(i)))
}

func tobool(v Object) bool {
	switch v := v.(type) {
	case Fixnum:
		return v != 0
	case Floatnum:
		return v != 0.0
	case String:
		return v != ""
	case Bool:
		return bool(v)
	case Table:
		return len(v) != 0
	case List:
		return len(v) != 0
	default:
		panic(fmt.Sprintf("invalid: %s", v))
	}
}

func lt(m *Machine, a, b Object) Object {
	switch v1 := a.(type) {
	case Fixnum:
		switch v2 := b.(type) {
		case Fixnum:
			return Bool(v1 < v2)
		case Floatnum:
			return Bool(Floatnum(v1) < v2)
		default:
			typeError(m, b, T_FIXNUM|T_FLOATNUM)
			return nil
		}
	case Floatnum:
		switch v2 := b.(type) {
		case Fixnum:
			return Bool(v1 < Floatnum(v2))
		case Floatnum:
			return Bool(v1 < v2)
		default:
			typeError(m, b, T_FIXNUM|T_FLOATNUM)
			return nil
		}
	case string:
		switch v2 := b.(type) {
		case string:
			return Bool(v1 < v2)
		default:
			typeError(m, b, T_STRING)
			return nil
		}
	default:
		typeError(m, a, T_FIXNUM|T_FLOATNUM|T_STRING)
		return nil
	}
}

func le(m *Machine, a, b Object) Object {
	switch v1 := a.(type) {
	case Fixnum:
		switch v2 := b.(type) {
		case Fixnum:
			return Bool(v1 <= v2)
		case Floatnum:
			return Bool(Floatnum(v1) <= v2)
		default:
			typeError(m, b, T_FIXNUM|T_FLOATNUM)
			return nil
		}
	case Floatnum:
		switch v2 := b.(type) {
		case Fixnum:
			return Bool(v1 <= Floatnum(v2))
		case Floatnum:
			return Bool(v1 <= v2)
		default:
			typeError(m, b, T_FIXNUM|T_FLOATNUM)
			return nil
		}
	case string:
		switch v2 := b.(type) {
		case string:
			return Bool(v1 <= v2)
		default:
			typeError(m, b, T_STRING)
			return nil
		}
	default:
		typeError(m, a, T_FIXNUM|T_FLOATNUM|T_STRING)
		return nil
	}
}

func equal(a, b Object) bool {
	switch v1 := a.(type) {
	case Fixnum:
		switch v2 := b.(type) {
		case Fixnum:
			return v1 == v2
		case Floatnum:
			return Floatnum(v1) == v2
		}
	case Floatnum:
		switch v2 := b.(type) {
		case Fixnum:
			return v1 == Floatnum(v2)
		case Floatnum:
			return v1 == v2
		}
	case string:
		switch v2 := b.(type) {
		case string:
			return v1 == v2
		}
	case bool:
		switch v2 := b.(type) {
		case bool:
			return v1 == v2
		}
	case Table:
		switch b.(type) {
		case Table:
			return false
		}
	case List:
		switch v2 := b.(type) {
		case List:
			if len(v1) != len(v2) {
				return false
			}
			for i := 0; i < len(v1); i++ {
				if !equal(v1[i], v2[i]) {
					return false
				}
			}
			return true
		}
	}
	return false
}

func index(m *Machine, x, y Object) Object {
	switch seq := x.(type) {
	case Table:
		v, ok := seq[y]
		if ok {
			return v
		} else {
			return NilValue
		}
	case List:
		i := assertInt(m, y, T_FIXNUM)
		if 0 <= i && i < len(seq) {
			return seq[i]
		} else {
			indexError(m, seq, i)
		}
	case String:
		i := assertInt(m, y, T_FIXNUM)
		if 0 <= i && i < len(seq) {
			return String(seq[i])
		} else {
			indexError(m, seq, i)
		}
	default:
		typeError(m, x, T_TABEL|T_LIST)
	}
	return nil
}

func setIndex(m *Machine, x, y, z Object) {
	switch seq := x.(type) {
	case Table:
		seq[y] = z
	case List:
		i := assertInt(m, y, T_FIXNUM)
		if 0 <= i && i < len(seq) {
			seq[i] = z
		} else {
			indexError(m, seq, i)
		}
	default:
		typeError(m, x, T_TABEL|T_LIST)
	}
}

func roundIndex(i, length int) int {
	if i < 0 {
		i = 0
	} else if i > length {
		i = length
	}
	return i
}

func sliceList(m *Machine, seq List, begin, end int) Object {
	length := len(seq)
	begin = roundIndex(begin, length)
	end = roundIndex(end, length)
	if begin > end {
		begin = end
	}
	return seq[begin:end]
}

func sliceString(m *Machine, seq String, begin, end int) Object {
	length := len(seq)
	begin = roundIndex(begin, length)
	end = roundIndex(end, length)
	if begin > end {
		begin = end
	}
	return seq[begin:end]
}

func slice0(m *Machine, x Object) Object {
	switch seq := x.(type) {
	case List:
		return sliceList(m, seq, 0, len(seq))
	case String:
		return sliceString(m, seq, 0, len(seq))
	default:
		typeError(m, x, T_LIST|T_STRING)
		return nil
	}
}

func slice1(m *Machine, x, y Object) Object {
	switch seq := x.(type) {
	case List:
		return sliceList(m, seq, assertInt(m, y, T_FIXNUM), len(seq))
	case String:
		return sliceString(m, seq, assertInt(m, y, T_FIXNUM), len(seq))
	default:
		typeError(m, x, T_LIST|T_STRING)
		return nil
	}
}

func slice2(m *Machine, x, y Object) Object {
	switch seq := x.(type) {
	case List:
		return sliceList(m, seq, 0, assertInt(m, y, T_FIXNUM))
	case String:
		return sliceString(m, seq, 0, assertInt(m, y, T_FIXNUM))
	default:
		typeError(m, x, T_LIST|T_STRING)
		return nil
	}
}

func slice(m *Machine, x, y, z Object) Object {
	switch seq := x.(type) {
	case List:
		return sliceList(m, seq, assertInt(m, y, T_FIXNUM), assertInt(m, z, T_FIXNUM))
	case String:
		return sliceString(m, seq, assertInt(m, y, T_FIXNUM), assertInt(m, z, T_FIXNUM))
	default:
		typeError(m, x, T_LIST|T_STRING)
		return nil
	}
}

func runLoop(m *Machine, isLooping bool) (Object, bool) {
	var lastValue Object = nil
	for {
		op := m.codeAhead()
		switch op {
		case MINUS:
			v := m.tos()
			switch v := v.(type) {
			case Fixnum:
				m.setTos(Fixnum(-v))
			case Floatnum:
				m.setTos(Floatnum(-v))
			default:
				typeError(m, m.tos(), T_FIXNUM|T_FLOATNUM)
			}
		case NOT:
			v := m.tos()
			m.setTos(Bool(!tobool(v)))
		case ADD:
			a := m.tos1()
			b := m.tos()
			switch v1 := a.(type) {
			case Fixnum:
				switch v2 := b.(type) {
				case Fixnum:
					m.setTos1(Fixnum(v1 + v2))
				case Floatnum:
					m.setTos1(Floatnum(Floatnum(v1) + v2))
				default:
					typeError(m, b, T_FIXNUM|T_FLOATNUM)
				}
			case Floatnum:
				switch v2 := b.(type) {
				case Fixnum:
					m.setTos1(Floatnum(v1 + Floatnum(v2)))
				case Floatnum:
					m.setTos1(Floatnum(v1 + v2))
				default:
					typeError(m, b, T_FIXNUM|T_FLOATNUM)
				}
			default:
				typeError(m, a, T_FIXNUM|T_FLOATNUM)
			}
		case SUB:
			a := m.tos1()
			b := m.tos()
			switch v1 := a.(type) {
			case Fixnum:
				switch v2 := b.(type) {
				case Fixnum:
					m.setTos1(Fixnum(v1 - v2))
				case Floatnum:
					m.setTos1(Floatnum(Floatnum(v1) - v2))
				default:
					typeError(m, b, T_FIXNUM|T_FLOATNUM)
				}
			case Floatnum:
				switch v2 := b.(type) {
				case Fixnum:
					m.setTos1(Floatnum(v1 - Floatnum(v2)))
				case Floatnum:
					m.setTos1(Floatnum(v1 - v2))
				default:
					typeError(m, b, T_FIXNUM|T_FLOATNUM)
				}
			default:
				typeError(m, a, T_FIXNUM|T_FLOATNUM)
			}
		case MUL:
			a := m.tos1()
			b := m.tos()
			switch v1 := a.(type) {
			case Fixnum:
				switch v2 := b.(type) {
				case Fixnum:
					m.setTos1(Fixnum(v1 * v2))
				case Floatnum:
					m.setTos1(Floatnum(Floatnum(v1) * v2))
				default:
					typeError(m, b, T_FIXNUM|T_FLOATNUM)
				}
			case Floatnum:
				switch v2 := b.(type) {
				case Fixnum:
					m.setTos1(Floatnum(v1 * Floatnum(v2)))
				case Floatnum:
					m.setTos1(Floatnum(v1 * v2))
				default:
					typeError(m, b, T_FIXNUM|T_FLOATNUM)
				}
			default:
				typeError(m, a, T_FIXNUM|T_FLOATNUM)
			}
		case DIV:
			a := m.tos1()
			b := m.tos()
			switch v1 := a.(type) {
			case Fixnum:
				switch v2 := b.(type) {
				case Fixnum:
					m.setTos1(Fixnum(v1 / v2))
				case Floatnum:
					m.setTos1(Floatnum(Floatnum(v1) / v2))
				default:
					typeError(m, b, T_FIXNUM|T_FLOATNUM)
				}
			case Floatnum:
				switch v2 := b.(type) {
				case Fixnum:
					m.setTos1(Floatnum(v1 / Floatnum(v2)))
				case Floatnum:
					m.setTos1(Floatnum(v1 / v2))
				default:
					typeError(m, b, T_FIXNUM|T_FLOATNUM)
				}
			default:
				typeError(m, a, T_FIXNUM|T_FLOATNUM)
			}
		case MOD:
			a := m.tos1()
			b := m.tos()
			v1 := assertInt(m, a, T_FIXNUM|T_FLOATNUM)
			v2 := assertInt(m, b, T_FIXNUM|T_FLOATNUM)
			m.setTos1(Fixnum(v1 % v2))
		case LT:
			m.setTos1(lt(m, m.tos1(), m.tos()))
		case LE:
			m.setTos1(le(m, m.tos1(), m.tos()))
		case GT:
			m.setTos1(lt(m, m.tos(), m.tos1()))
		case GE:
			m.setTos1(le(m, m.tos(), m.tos1()))
		case EQ:
			m.setTos1(Bool(equal(m.tos1(), m.tos())))
		case NE:
			m.setTos1(Bool(!equal(m.tos1(), m.tos())))
		case INDEX:
			m.setTos1(index(m, m.tos1(), m.tos()))
		case SLICE:
			m.setTos2(slice(m, m.tos2(), m.tos1(), m.tos()))
		case SLICE0:
			m.setTos(slice0(m, m.tos()))
		case SLICE1:
			m.setTos(slice1(m, m.tos1(), m.tos()))
		case SLICE2:
			m.setTos(slice2(m, m.tos1(), m.tos()))
		case SET_INDEX:
			setIndex(m, m.tos1(), m.tos(), m.tos2())
			m.popN(3)
		case POP:
			lastValue = m.pop()
		case PUSH_NIL:
			m.push(NilValue)
		case PUSH_TRUE:
			m.push(Bool(true))
		case PUSH_FALSE:
			m.push(Bool(false))
		case CONST:
			m.push(m.getConst(m.arg1()))
		case GREF:
			v, ok := m.getGlobal(m.arg1())
			if !ok {
				unboundVariableError(m, m.arg1())
			} else {
				m.push(v)
			}
		case GSET:
			m.setGlobal(m.arg1(), m.pop())
		case LREF:
			m.push(m.getLocal(m.arg1(), m.arg2()))
		case LSET:
			m.setLocal(m.arg1(), m.arg2(), m.pop())
		case MAKE_FUNCTION:
			v := m.getConst(m.arg1())
			m.push(NewFunction(v.(*Function), m.copyToHeapEnv()))
		case CALL:
			argc := m.arg1()
			switch fn := m.pop().(type) {
			case *Function:
				if fn.validArgc(argc) {
					if fn.isVariadic() {
						size := argc - fn.min
						list := make(List, size)
						for i := size - 1; i >= 0; i-- {
							list[i] = m.pop()
						}
						m.push(list)
					}
					m.stackTop += fn.numLocalVars - argc
					env := NewStackEnv(fn.env,
						m.stackTop-fn.numLocalVars,
						fn.numLocalVars)
					m.push(m.env)
					m.push(m.pc + 1)
					m.push(m.code)
					m.setProgram(fn.code, 0, env)
					continue
				} else {
					raise(m, "invalid number of argument")
				}
			case *BuiltinFunction:
				if fn.min <= argc && argc <= fn.max {
					args := make([]Object, argc)
					for i := argc - 1; i >= 0; i-- {
						args[i] = m.pop()
					}
					m.push(fn.fun(m, args))
				} else {
					raise(m, "invalid number of argument")
				}
			default:
				typeError(m, fn, T_FUNCTION)
			}
		case RETURN:
			sp := m.env.sp
			v := m.pop()
			code := m.pop().([]Instr)
			pc := m.pop().(int)
			env := m.pop().(*Env)
			m.stackTop = sp
			m.setProgram(code, pc, env)
			m.push(v)
			if isLooping {
				return nil, true
			}
			continue
		case NEW_LIST:
			n := m.arg1()
			v := make(List, n)
			for i := n - 1; i >= 0; i-- {
				v[i] = m.pop()
			}
			m.push(v)
		case NEW_TABLE:
			m.push(make(Table))
		case SET_TABLE:
			v := m.pop()
			k := m.pop()
			table := m.tos().(Table)
			table[k] = v
		case RANGE:
			var step Fixnum
			if m.arg1() == 3 {
				step = assertFixnum(m, m.pop(), T_FIXNUM)
			} else {
				step = 1
			}
			end := assertFixnum(m, m.pop(), T_FIXNUM)
			start := assertFixnum(m, m.pop(), T_FIXNUM)
			m.push(&Range{start: start, end: end, step: step})
		case FOR:
			seq := m.pop()
			block := m.getConst(m.arg1()).(*Block)
			saveCode := m.code
			savePc := m.pc
			switch seq := seq.(type) {
			case List:
				for _, v := range seq {
					m.push(v)
					m.setProgram(block.code, 0, m.env)
					_, isBreak := runLoop(m, true)
					if isBreak {
						break
					}
				}
			case Table:
				for v, _ := range seq {
					m.push(v)
					m.setProgram(block.code, 0, m.env)
					_, isBreak := runLoop(m, true)
					if isBreak {
						break
					}
				}
			case String:
				for _, v := range seq {
					m.push(String(v))
					m.setProgram(block.code, 0, m.env)
					_, isBreak := runLoop(m, true)
					if isBreak {
						break
					}
				}
			case *Range:
				for i := seq.start; i < seq.end; i += seq.step {
					m.push(i)
					m.setProgram(block.code, 0, m.env)
					_, isBreak := runLoop(m, true)
					if isBreak {
						break
					}
				}
			default:
				typeError(m, seq, T_LIST|T_TABEL|T_STRING|T_RANGE)
			}
			m.setProgram(saveCode, savePc+1, m.env)
			continue
		case END:
			return nil, false
		case BREAK:
			return nil, true
		case CONTINUE:
			return nil, false
		case JUMP:
			m.jump(m.arg1())
			continue
		case JUMP_IF_FALSE:
			if !tobool(m.pop()) {
				m.jump(m.arg1())
				continue
			}
		case JUMP_IF_TRUE_OR_POP:
			if tobool(m.tos()) {
				m.jump(m.arg1())
				continue
			}
		case JUMP_IF_FALSE_OR_POP:
			if !tobool(m.tos()) {
				m.jump(m.arg1())
				continue
			}
		case HALT:
			return lastValue, false
		default:
			panic(fmt.Sprintf("unknown opcode: %s", op))
		}
		m.step()
	}
}

func Run(m *Machine) Object {
	exe := m.exe
	m.err = nil
	m.code = m.exe.code
	m.constTable = *exe.constTable
	m.env = nil
	m.stack = make([]Object, 1000)
	m.stackTop = 0
	m.pc = 0
	if len(m.globalTable) < exe.numGlobalVars() {
		for i := len(m.globalTable); i < exe.numGlobalVars(); i++ {
			m.globalTable = append(m.globalTable, nil)
		}
	}

	defer func() {
		err := recover()
		if m.err != err {
			panic(err)
		} else if err != nil {
			fmt.Println(err)
		}
	}()

	v, _ := runLoop(m, false)
	return v
}
