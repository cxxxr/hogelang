package compiler

type Env struct {
	parent *Env
	names  []string
}

func NewEnv(parent *Env) *Env {
	return &Env{parent: parent, names: []string{}}
}

func (env *Env) Put(name string) {
	if env.parent != nil {
		env.names = append(env.names, name)
	}
}

func (env *Env) find(name string, i int) (int, int) {
	for j, n := range env.names {
		if n == name {
			return i, j
		}
	}
	if env.parent == nil {
		return -1, -1
	} else {
		return env.parent.find(name, i+1)
	}
}

func (env *Env) Find(name string) (int, int) {
	return env.find(name, 0)
}
