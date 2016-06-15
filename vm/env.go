package vm

type Env struct {
	parent *Env
	sp int
}

func NewEnv(parent *Env, sp int) *Env {
	return &Env{parent: parent, sp: sp}
}
