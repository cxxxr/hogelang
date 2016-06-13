package vm

type Env struct {
	parent *Env
	values []Object
}

func NewEnv(parent *Env, values []Object) *Env {
	return &Env{parent: parent, values: values}
}

func (env *Env) get(i, j int) Object {
	for n := 0; n < i; n++ {
		env = env.parent
	}
	return env.values[j]
}

func (env *Env) set(i int, j int, v Object) {
	for n := 0; n < i; n++ {
		env = env.parent
	}
	env.values[j] = v
}
