package vm

import (
	"fmt"
	"github.com/cxxxr/hogelang/ast"
)

type Opcode int

const (
	ADD Opcode = iota + 1
	CALL
	CONST
	DIV
	EQ
	NEW_LIST
	NEW_TABLE
	SET_TABLE
	FOR
	END
	BREAK
	CONTINUE
	GE
	GREF
	GSET
	GT
	HALT
	INDEX
	JUMP
	JUMP_IF_FALSE
	JUMP_IF_FALSE_OR_POP
	JUMP_IF_TRUE_OR_POP
	LABEL
	LE
	LREF
	LSET
	LT
	MAKE_FUNCTION
	MINUS
	MOD
	MUL
	NE
	NOT
	POP
	PUSH_NIL
	PUSH_FALSE
	PUSH_TRUE
	RETURN
	SET_INDEX
	SLICE
	SLICE0
	SLICE1
	SLICE2
	SUB
	RANGE
)

func (o Opcode) String() string {
	switch o {
	case ADD:
		return "ADD"
	case CALL:
		return "CALL"
	case CONST:
		return "CONST"
	case DIV:
		return "DIV"
	case EQ:
		return "EQ"
	case NEW_LIST:
		return "NEW_LIST"
	case NEW_TABLE:
		return "NEW_TABLE"
	case SET_TABLE:
		return "SET_TABLE"
	case FOR:
		return "FOR"
	case END:
		return "END"
	case BREAK:
		return "BREAK"
	case CONTINUE:
		return "CONTINUE"
	case GE:
		return "GE"
	case GREF:
		return "GREF"
	case GSET:
		return "GSET"
	case GT:
		return "GT"
	case HALT:
		return "HALT"
	case INDEX:
		return "INDEX"
	case JUMP:
		return "JUMP"
	case JUMP_IF_FALSE:
		return "JUMP_IF_FALSE"
	case JUMP_IF_TRUE_OR_POP:
		return "JUMP_IF_TRUE_OR_POP"
	case JUMP_IF_FALSE_OR_POP:
		return "JUMP_IF_FALSE_OR_POP"
	case LABEL:
		return "LABEL"
	case LE:
		return "LE"
	case LREF:
		return "LREF"
	case LSET:
		return "LSET"
	case LT:
		return "LT"
	case MAKE_FUNCTION:
		return "MAKE_FUNCTION"
	case MINUS:
		return "MINUS"
	case MOD:
		return "MOD"
	case MUL:
		return "MUL"
	case NE:
		return "NE"
	case NOT:
		return "NOT"
	case POP:
		return "POP"
	case PUSH_NIL:
		return "PUSH_NIL"
	case PUSH_FALSE:
		return "PUSH_FALSE"
	case PUSH_TRUE:
		return "PUSH_TRUE"
	case RETURN:
		return "RETURN"
	case SET_INDEX:
		return "SET_INDEX"
	case SLICE:
		return "SLICE"
	case SLICE0:
		return "SLICE0"
	case SLICE1:
		return "SLICE1"
	case SLICE2:
		return "SLICE2"
	case SUB:
		return "SUB"
	case RANGE:
		return "RANGE"
	default:
		return "???"
	}
}

var opInfoTable = map[Opcode]struct {
	argnum int
	isJump bool
}{
	ADD:                  {0, false},
	CALL:                 {1, false},
	CONST:                {1, false},
	DIV:                  {0, false},
	EQ:                   {0, false},
	NEW_LIST:             {1, false},
	NEW_TABLE:            {1, false},
	SET_TABLE:            {0, false},
	FOR:                  {1, false},
	END:                  {0, false},
	BREAK:                {0, false},
	CONTINUE:             {0, false},
	GE:                   {0, false},
	GREF:                 {1, false},
	GSET:                 {1, false},
	GT:                   {0, false},
	HALT:                 {0, false},
	INDEX:                {0, false},
	JUMP:                 {1, true},
	JUMP_IF_FALSE:        {1, true},
	JUMP_IF_TRUE_OR_POP:  {1, true},
	JUMP_IF_FALSE_OR_POP: {1, true},
	LABEL:                {1, false},
	LE:                   {0, false},
	LREF:                 {2, false},
	LSET:                 {2, false},
	LT:                   {0, false},
	MAKE_FUNCTION:        {1, false},
	MINUS:                {0, false},
	MOD:                  {0, false},
	MUL:                  {0, false},
	NE:                   {0, false},
	NOT:                  {0, false},
	POP:                  {0, false},
	PUSH_NIL:             {0, false},
	PUSH_TRUE:            {0, false},
	PUSH_FALSE:           {0, false},
	RETURN:               {0, false},
	SET_INDEX:            {0, false},
	SLICE:                {0, false},
	SLICE0:               {0, false},
	SLICE1:               {0, false},
	SLICE2:               {0, false},
	SUB:                  {0, false},
	RANGE:                {1, false},
}

func OpcodeArgnum(opcode Opcode) int {
	return opInfoTable[opcode].argnum
}

type Instr struct {
	Opcode Opcode
	Arg1   int
	Arg2   int
	Pos    *ast.Pos
}

func (instr Instr) IsJump() bool {
	return opInfoTable[instr.Opcode].isJump
}

type Executable struct {
	code       []Instr
	constTable *[]Object
	globalVars *[]string
}

func NewExecutable(orig *Executable) *Executable {
	exe := &Executable{code: []Instr{}}
	if orig == nil {
		exe.constTable = &[]Object{}
		exe.globalVars = &[]string{}
	} else {
		exe.constTable = orig.constTable
		exe.globalVars = orig.globalVars
	}
	return exe
}

func (exe *Executable) Init() {
	exe.code = make([]Instr, 0)
}

func (exe *Executable) GetCode() []Instr {
	return exe.code
}

func (exe *Executable) Gen(pos *ast.Pos, opcode Opcode, arg1 int, arg2 int) {
	exe.code = append(exe.code, Instr{Opcode: opcode, Arg1: arg1, Arg2: arg2, Pos: pos})
}

func (exe *Executable) GenCopy(instr Instr) {
	exe.Gen(instr.Pos, instr.Opcode, instr.Arg1, instr.Arg2)
}

func (exe *Executable) SetConst(val Object) int {
	for i, val2 := range *exe.constTable {
		if val2 == val {
			return i
		}
	}
	*exe.constTable = append(*exe.constTable, val)
	return len(*exe.constTable) - 1
}

func (exe *Executable) GetConst(i int) Object {
	return (*exe.constTable)[i]
}

func (exe *Executable) AddGlobal(name string) (int, bool) {
	for i, name2 := range *exe.globalVars {
		if name == name2 {
			return i, false
		}
	}
	*exe.globalVars = append(*exe.globalVars, name)
	return len(*exe.globalVars) - 1, true
}

func (exe *Executable) globalVarName(i int) string {
	return (*exe.globalVars)[i]
}

func (exe *Executable) numGlobalVars() int {
	return len(*exe.globalVars)
}

func (exe *Executable) Show() {
	fmt.Println()
	fmt.Println(exe.globalVars)
	showConstTable(*exe.constTable)
	showCode(exe.code)
	fmt.Println()
}

func showCode(code []Instr) {
	for i, instr := range code {
		opinfo := opInfoTable[instr.Opcode]
		fmt.Print(i, "\t", instr.Opcode)
		switch opinfo.argnum {
		case 0:
		case 1:
			fmt.Print(" ", instr.Arg1)
		case 2:
			fmt.Print(" ", instr.Arg1)
			fmt.Print(" ", instr.Arg2)
		}
		fmt.Println()
	}
}

func showConstTable(constTable []Object) {
	for i, val := range constTable {
		switch val := val.(type) {
		case *Function:
			fmt.Println(i, ":")
			showCode(val.code)
		case *Block:
			fmt.Println(i, ":")
			showCode(val.code)
		default:
			fmt.Println(i, ":", val)
		}
	}
	fmt.Println()
}
