package vm

import (
	"fmt"
	"strings"
)

const (
	T_NIL = 1 << iota
	T_FIXNUM
	T_FLOATNUM
	T_STRING
	T_BOOL
	T_TABEL
	T_LIST
	T_FUNCTION
	T_RANGE
)

func typesToString(types int, delim string) string {
	ar := make([]string, 0)
	if types&T_NIL != 0 {
		ar = append(ar, "nil")
	}
	if types&T_FIXNUM != 0 {
		ar = append(ar, "fixnum")
	}
	if types&T_FLOATNUM != 0 {
		ar = append(ar, "floatnum")
	}
	if types&T_STRING != 0 {
		ar = append(ar, "string")
	}
	if types&T_BOOL != 0 {
		ar = append(ar, "bool")
	}
	if types&T_TABEL != 0 {
		ar = append(ar, "tabel")
	}
	if types&T_LIST != 0 {
		ar = append(ar, "list")
	}
	if types&T_FUNCTION != 0 {
		ar = append(ar, "function")
	}
	if types&T_RANGE != 0 {
		ar = append(ar, "range")
	}
	var str = ""
	for i, s := range ar {
		str += s
		if i < len(ar)-1 {
			str += delim
		}
	}
	return str
}

type Object interface{}
type Fixnum int
type Floatnum float64
type String string
type Bool bool
type Table map[Object]Object
type List []Object

func (x Fixnum) String() string {
	return fmt.Sprintf("%d", x)
}

func (x Floatnum) String() string {
	return fmt.Sprintf("%f", x)
}

func (x String) String() string {
	return string(x)
}

func (x Table) String() string {
	ar := []string{}
	for k, v := range x {
		ar = append(ar, fmt.Sprintf("%s = %s", k, v))
	}
	return "@[" + strings.Join(ar, ", ") + "]"
}

func (x List) String() string {
	ar := make([]string, len(x))
	for i, v := range x {
		ar[i] = fmt.Sprint(v)
	}
	return "[" + strings.Join(ar, ", ") + "]"
}

var NilValue Object = nil

type Block struct {
	code []Instr
}

func NewBlock(code []Instr) *Block {
	return &Block{code: code}
}

type Function struct {
	parameters   []string
	code         []Instr
	numLocalVars int
	env          *Env
}

func MakeFunction(parms []string, code []Instr, numLocalVars int) *Function {
	return &Function{
		parameters:   parms,
		code:         code,
		numLocalVars: numLocalVars,
	}
}

func NewFunction(fn *Function, env *Env) *Function {
	return &Function{
		parameters:   fn.parameters,
		code:         fn.code,
		numLocalVars: fn.numLocalVars,
		env:          env,
	}
}

func (fn *Function) argc() int {
	return len(fn.parameters)
}

type BuiltinFunction struct {
	min int
	max int
	fun func(*Machine, []Object) Object
}

type Range struct {
	start Fixnum
	end   Fixnum
	step  Fixnum
}

func (r *Range) String() string {
	if r.step == 1 {
		return fmt.Sprintf("[%s..%s]", r.start, r.end)
	} else {
		return fmt.Sprintf("[%s..%s,%s]", r.start, r.end, r.step)
	}
}
