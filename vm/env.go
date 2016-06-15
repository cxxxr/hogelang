package vm

type Env struct {
	parent *Env
	sp int
	length int
	values []Object
}

func NewStackEnv(parent *Env, sp int, length int) *Env {
	return &Env{parent: parent, sp: sp, length: length}
}

func NewHeapEnv(parent *Env, values []Object) *Env {
	return &Env{parent: parent, values: values, sp: -1}
}
