package vm

import (
	"fmt"
)

const MAX_NUM_ARGUMENTS = 256

func define(m *Machine, name string, min int, max int, fun func(*Machine, []Object) Object) {
	_, ok := m.exe.AddGlobal(name)
	if ok {
		if max < 0 {
			max = MAX_NUM_ARGUMENTS
		}
		m.addGlobal(&BuiltinFunction{min: min, max: max, fun: fun})
	}
}

func Init(m *Machine) {
	define(m, "print", 0, -1, func(m *Machine, args []Object) Object {
		for i, v := range args {
			if i < len(args)-1 {
				fmt.Print(v, " ")
			} else {
				fmt.Print(v)
			}
		}
		fmt.Println()
		return NilValue
	})
}
