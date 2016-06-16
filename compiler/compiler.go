package compiler

import (
	"fmt"
	"github.com/cxxxr/hogelang/ast"
	"github.com/cxxxr/hogelang/vm"
	"strconv"
)

type SyntaxError struct {
	msg string
}

func (e *SyntaxError) Error() string {
	return e.msg
}

var labelCounter int

func initLabelCounter() {
	labelCounter = -1
}

func genLabel() int {
	labelCounter++
	return labelCounter
}

type loopType int

const (
	WHILE loopType = iota
	FOR
)

type loopInfo struct {
	typ        loopType
	beginLabel int
	endLabel   int
}

func syntaxError(msg string) {
	panic(&SyntaxError{msg: fmt.Sprintf("syntax error: %s", msg)})
}

func genSetVar(exe *vm.Executable, env *Env, name string, pos *ast.Pos) {
	i, j := env.Find(name)
	if i == -1 {
		i, _ = exe.AddGlobal(name)
		exe.Gen(pos, vm.GSET, i, 0)
	} else {
		exe.Gen(pos, vm.LSET, i, j)
	}
}

func compStmt(stmt ast.Stmt, exe *vm.Executable, env *Env, loopinfo *loopInfo) {
	switch stmt := stmt.(type) {
	case *ast.Block:
		compBlock(stmt, exe, env, loopinfo)
	case *ast.VarStmt:
		compVarStmt(stmt, exe, env, loopinfo)
	case *ast.AssignStmt:
		compAssignStmt(stmt, exe, env, loopinfo)
	case *ast.IfStmt:
		compIfStmt(stmt, exe, env, loopinfo)
	case *ast.ForStmt:
		compForStmt(stmt, exe, env, loopinfo)
	case *ast.WhileStmt:
		compWhileStmt(stmt, exe, env, loopinfo)
	case *ast.ReturnStmt:
		compReturnStmt(stmt, exe, env, loopinfo)
	case *ast.BreakStmt:
		compBreakStmt(stmt, exe, env, loopinfo)
	case *ast.ContinueStmt:
		compContinueStmt(stmt, exe, env, loopinfo)
	case *ast.ExprStmt:
		compExprStmt(stmt, exe, env, loopinfo)
	}
}

func compBlock(block *ast.Block, exe *vm.Executable, env *Env, loopinfo *loopInfo) {
	for _, stmt := range block.Stmts {
		compStmt(stmt, exe, env, loopinfo)
	}
}

func compVarStmt(stmt *ast.VarStmt, exe *vm.Executable, env *Env, loopinfo *loopInfo) {
	for _, decl := range stmt.Vars {
		if decl.Value != nil {
			compExpr(decl.Value, exe, env)
		} else {
			exe.Gen(stmt.Pos(), vm.PUSH_NIL, 0, 0)
		}
		env.Put(decl.Name)
		genSetVar(exe, env, decl.Name, stmt.Pos())
	}
}

func compAssignStmt(stmt *ast.AssignStmt, exe *vm.Executable, env *Env, loopinfo *loopInfo) {
	compExpr(stmt.Rhs, exe, env)
	compLhs(stmt.Lhs, exe, env)
}

func compLhs(expr ast.Expr, exe *vm.Executable, env *Env) {
	switch expr := expr.(type) {
	case *ast.RefvarExpr:
		genSetVar(exe, env, expr.Name, expr.Pos())
	case *ast.IndexExpr:
		compExpr(expr.Prefix, exe, env)
		compExpr(expr.Index, exe, env)
		exe.Gen(expr.Pos(), vm.SET_INDEX, 0, 0)
	default:
		syntaxError("can't assign")
	}
}

func compIfStmt(stmt *ast.IfStmt, exe *vm.Executable, env *Env, loopinfo *loopInfo) {
	label1 := genLabel()
	label2 := genLabel()
	compExpr(stmt.Test, exe, env)
	exe.Gen(stmt.Pos(), vm.JUMP_IF_FALSE, label1, 0)
	compStmt(stmt.Then, exe, env, loopinfo)
	exe.Gen(stmt.Pos(), vm.JUMP, label2, 0)
	exe.Gen(stmt.Pos(), vm.LABEL, label1, 0)
	compStmt(stmt.Else, exe, env, loopinfo)
	exe.Gen(stmt.Pos(), vm.LABEL, label2, 0)
}

func compForStmt(stmt *ast.ForStmt, exe *vm.Executable, env *Env, loopinfo *loopInfo) {
	compExpr(stmt.Exp, exe, env)
	env.Put(stmt.Name)
	exe2 := vm.NewExecutable(exe)
	genSetVar(exe2, env, stmt.Name, stmt.Pos())
	compStmt(stmt.Body, exe2, env, &loopInfo{typ: FOR})
	exe2.Gen(stmt.Pos(), vm.END, 0, 0)
	exe2 = assemble(exe2)
	exe.Gen(stmt.Pos(), vm.FOR, exe.SetConst(vm.NewBlock(exe2.GetCode())), 0)
}

func compWhileStmt(stmt *ast.WhileStmt, exe *vm.Executable, env *Env, loopinfo *loopInfo) {
	beginLabel := genLabel()
	endLabel := genLabel()
	exe.Gen(stmt.Pos(), vm.LABEL, beginLabel, 0)
	compExpr(stmt.Test, exe, env)
	exe.Gen(stmt.Pos(), vm.JUMP_IF_FALSE, endLabel, 0)
	compStmt(stmt.Body, exe, env, &loopInfo{typ: WHILE, beginLabel: beginLabel, endLabel: endLabel})
	exe.Gen(stmt.Pos(), vm.JUMP, beginLabel, 0)
	exe.Gen(stmt.Pos(), vm.LABEL, endLabel, 0)
}

func compReturnStmt(stmt *ast.ReturnStmt, exe *vm.Executable, env *Env, loopinfo *loopInfo) {
	if stmt.Value == nil {
		exe.Gen(stmt.Pos(), vm.PUSH_NIL, 0, 0)
	} else {
		compExpr(stmt.Value, exe, env)
	}
	exe.Gen(stmt.Pos(), vm.RETURN, 0, 0)
}

func compBreakStmt(stmt *ast.BreakStmt, exe *vm.Executable, env *Env, loopinfo *loopInfo) {
	if loopinfo == nil {
		syntaxError("break outside loop")
	}
	switch loopinfo.typ {
	case FOR:
		exe.Gen(stmt.Pos(), vm.BREAK, 0, 0)
	case WHILE:
		exe.Gen(stmt.Pos(), vm.JUMP, loopinfo.endLabel, 0)
	}
}

func compContinueStmt(stmt *ast.ContinueStmt, exe *vm.Executable, env *Env, loopinfo *loopInfo) {
	if loopinfo == nil {
		syntaxError("continue outside loop")
	}
	switch loopinfo.typ {
	case FOR:
		exe.Gen(stmt.Pos(), vm.CONTINUE, 0, 0)
	case WHILE:
		exe.Gen(stmt.Pos(), vm.JUMP, loopinfo.beginLabel, 0)
	}
}

func compExprStmt(stmt *ast.ExprStmt, exe *vm.Executable, env *Env, loopinfo *loopInfo) {
	compExpr(stmt.Exp, exe, env)
	exe.Gen(stmt.Pos(), vm.POP, 0, 0)
}

func compExpr(expr ast.Expr, exe *vm.Executable, env *Env) {
	switch expr := expr.(type) {
	case *ast.RefvarExpr:
		compRefvarExpr(expr, exe, env)
	case *ast.NilExpr:
		exe.Gen(expr.Pos(), vm.PUSH_NIL, 0, 0)
	case *ast.TrueExpr:
		exe.Gen(expr.Pos(), vm.PUSH_TRUE, 0, 0)
	case *ast.FalseExpr:
		exe.Gen(expr.Pos(), vm.PUSH_FALSE, 0, 0)
	case *ast.FixnumExpr:
		v, err := strconv.ParseInt(expr.Value, 10, 64)
		if err != nil {
			panic(err)
		}
		compConstExpr(expr.Pos(), vm.Fixnum(v), exe)
	case *ast.FloatnumExpr:
		v, err := strconv.ParseFloat(expr.Value, 64)
		if err != nil {
			panic(err)
		}
		compConstExpr(expr.Pos(), vm.Floatnum(v), exe)
	case *ast.StringExpr:
		compConstExpr(expr.Pos(), vm.String(expr.Value), exe)
	case *ast.FuncallExpr:
		compFuncallExpr(expr, exe, env)
	case *ast.FunctionExpr:
		compFunctionExpr(expr, exe, env)
	case *ast.UnaryExpr:
		compUnaryExpr(expr, exe, env)
	case *ast.BinaryExpr:
		compBinaryExpr(expr, exe, env)
	case *ast.AndExpr:
		compAndExpr(expr, exe, env)
	case *ast.OrExpr:
		compOrExpr(expr, exe, env)
	case *ast.IndexExpr:
		compIndexExpr(expr, exe, env)
	case *ast.SliceExpr:
		compSliceExpr(expr, exe, env)
	case *ast.ListExpr:
		compListExpr(expr, exe, env)
	case *ast.TableExpr:
		compTableExpr(expr, exe, env)
	case *ast.RangeExpr:
		compRangeExpr(expr, exe, env)
	}
}

func compRefvarExpr(expr *ast.RefvarExpr, exe *vm.Executable, env *Env) {
	i, j := env.Find(expr.Name)
	if i == -1 {
		i, _ = exe.AddGlobal(expr.Name)
		exe.Gen(expr.Pos(), vm.GREF, i, 0)
	} else {
		exe.Gen(expr.Pos(), vm.LREF, i, j)
	}
}

func compConstExpr(pos *ast.Pos, value vm.Object, exe *vm.Executable) {
	exe.Gen(pos, vm.CONST, exe.SetConst(value), 0)
}

func compFuncallExpr(expr *ast.FuncallExpr, exe *vm.Executable, env *Env) {
	for _, x := range expr.Args {
		compExpr(x, exe, env)
	}
	compExpr(expr.Fun, exe, env)
	exe.Gen(expr.Pos(), vm.CALL, len(expr.Args), 0)
}

func compFunctionExpr(expr *ast.FunctionExpr, exe *vm.Executable, env *Env) {
	env = NewEnv(env)
	for _, n := range expr.Parameters.Names {
		env.Put(n)
	}
	if expr.Parameters.Rest != "" {
		env.Put(expr.Parameters.Rest)
	}
	exe2 := vm.NewExecutable(exe)
	compStmt(expr.Body, exe2, env, nil)
	exe2.Gen(expr.Body.Pos(), vm.PUSH_NIL, 0, 0)
	exe2.Gen(expr.Body.Pos(), vm.RETURN, 0, 0)
	exe2 = assemble(exe2)
	fn := vm.MakeFunction(expr.Parameters, exe2.GetCode(), len(env.names))
	i := exe.SetConst(fn)
	exe.Gen(expr.Pos(), vm.MAKE_FUNCTION, i, 0)
}

func compUnaryExpr(expr *ast.UnaryExpr, exe *vm.Executable, env *Env) {
	compExpr(expr.Value, exe, env)
	var op vm.Opcode
	switch expr.Op {
	case "-":
		op = vm.MINUS
	case "!":
		op = vm.NOT
	}
	exe.Gen(expr.Pos(), op, 0, 0)
}

func compBinaryExpr(expr *ast.BinaryExpr, exe *vm.Executable, env *Env) {
	compExpr(expr.Left, exe, env)
	compExpr(expr.Right, exe, env)
	var op vm.Opcode
	switch expr.Op {
	case "+":
		op = vm.ADD
	case "-":
		op = vm.SUB
	case "*":
		op = vm.MUL
	case "/":
		op = vm.DIV
	case "%":
		op = vm.MOD
	case "<":
		op = vm.LT
	case "<=":
		op = vm.LE
	case "==":
		op = vm.EQ
	case "!=":
		op = vm.NE
	case ">":
		op = vm.GT
	case ">=":
		op = vm.GE
	}
	exe.Gen(expr.Pos(), op, 0, 0)
}

func compAndExpr(expr *ast.AndExpr, exe *vm.Executable, env *Env) {
	label := genLabel()
	compExpr(expr.Left, exe, env)
	exe.Gen(expr.Pos(), vm.JUMP_IF_FALSE_OR_POP, label, 0)
	compExpr(expr.Right, exe, env)
	exe.Gen(expr.Pos(), vm.LABEL, label, 0)
}

func compOrExpr(expr *ast.OrExpr, exe *vm.Executable, env *Env) {
	label := genLabel()
	compExpr(expr.Left, exe, env)
	exe.Gen(expr.Pos(), vm.JUMP_IF_TRUE_OR_POP, label, 0)
	compExpr(expr.Right, exe, env)
	exe.Gen(expr.Pos(), vm.LABEL, label, 0)
}

func compIndexExpr(expr *ast.IndexExpr, exe *vm.Executable, env *Env) {
	compExpr(expr.Prefix, exe, env)
	compExpr(expr.Index, exe, env)
	exe.Gen(expr.Pos(), vm.INDEX, 0, 0)
}

func compSliceExpr(expr *ast.SliceExpr, exe *vm.Executable, env *Env) {
	if expr.Begin == nil && expr.End == nil {
		compExpr(expr.Prefix, exe, env)
		exe.Gen(expr.Pos(), vm.SLICE0, 0, 0)
	} else if expr.End == nil {
		compExpr(expr.Prefix, exe, env)
		compExpr(expr.Begin, exe, env)
		exe.Gen(expr.Pos(), vm.SLICE1, 0, 0)
	} else if expr.Begin == nil {
		compExpr(expr.Prefix, exe, env)
		compExpr(expr.End, exe, env)
		exe.Gen(expr.Pos(), vm.SLICE2, 0, 0)
	} else {
		compExpr(expr.Prefix, exe, env)
		compExpr(expr.Begin, exe, env)
		compExpr(expr.End, exe, env)
		exe.Gen(expr.Pos(), vm.SLICE, 0, 0)
	}
}

func compListExpr(expr *ast.ListExpr, exe *vm.Executable, env *Env) {
	for _, e := range expr.Elements {
		compExpr(e, exe, env)
	}
	exe.Gen(expr.Pos(), vm.NEW_LIST, len(expr.Elements), 0)
}

func compTableExpr(expr *ast.TableExpr, exe *vm.Executable, env *Env) {
	exe.Gen(expr.Pos(), vm.NEW_TABLE, len(expr.Elements), 0)
	for _, e := range expr.Elements {
		pair := e.(*ast.PairExpr)
		compExpr(pair.Key, exe, env)
		compExpr(pair.Value, exe, env)
		exe.Gen(expr.Pos(), vm.SET_TABLE, 0, 0)
	}
}

func compRangeExpr(expr *ast.RangeExpr, exe *vm.Executable, env *Env) {
	n := 2
	compExpr(expr.Start, exe, env)
	compExpr(expr.End, exe, env)
	if expr.Step != nil {
		compExpr(expr.Step, exe, env)
		n++
	}
	exe.Gen(expr.Pos(), vm.RANGE, n, 0)
}

func assemble(exe *vm.Executable) *vm.Executable {
	exe = optimize(exe)
	exe2 := vm.NewExecutable(exe)

	labelMap := make([]int, labelCounter+1)
	count := 0

	code := exe.GetCode()

	for _, instr := range code {
		if instr.Opcode == vm.LABEL {
			labelMap[instr.Arg1] = count
		} else {
			count++
		}
	}

	for _, instr := range code {
		if instr.Opcode != vm.LABEL {
			if instr.IsJump() {
				exe2.Gen(instr.Pos, instr.Opcode, labelMap[instr.Arg1], 0)
			} else {
				exe2.GenCopy(instr)
			}
		}
	}

	return exe2
}

func Compile(exe *vm.Executable, x ast.Stmt) (exe2 *vm.Executable, err error) {
	defer func() {
		er := recover()
		if er == nil {
			return
		}
		if e, ok := er.(*SyntaxError); ok {
			err = e
		} else {
			panic(err)
		}
	}()

	initLabelCounter()
	exe.Init()
	compStmt(x, exe, NewEnv(nil), nil)
	exe.Gen(nil, vm.HALT, 0, 0)
	exe2 = assemble(exe)
	//exe2.Show()
	return
}
