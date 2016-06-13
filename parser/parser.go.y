%{
package parser

import (
	"github.com/cxxxr/hogelang/ast"
)

type Token struct {
	ast.Ast
	tok int
	literal string
}

%}

%union {
	block      *ast.Block
	stmts      []ast.Stmt
	stmt       ast.Stmt
	var_decls  []*ast.VarDeclStmt
	var_decl   *ast.VarDeclStmt
	if_stmt    *ast.IfStmt
	expr       ast.Expr
	expr_pair  ast.Expr
	expr_pairs []ast.Expr
	namelist   []string
	exprlist   []ast.Expr
	token      Token
}

%type<block> program
%type<block> block
%type<stmts> stmts
%type<stmt> stmt
%type<var_decls> var_decls
%type<var_decl> var_decl
%type<if_stmt> if_stmt
%type<if_stmt> elif_stmt
%type<expr> expr
%type<expr_pair> expr_pair
%type<expr_pairs> expr_pairs
%type<namelist> namelist
%type<exprlist> exprlist
%type<exprlist> exprlist_

%token<token> VAR FOR WHILE IF ELIF ELSE END RETURN BREAK CONTINUE NIL TRUE FALSE ID FIXNUM FLOATNUM STRING AND OR LT LE EQ NE GT GE NL LIST_DISPATCH TABLE_DISPATCH COLON DOTDOT

%left OR
%left AND
%left GT GE LT LE
%left '+' '-'
%left '*' '/' '%'
%right UNARY

%%

program
	: block
	{
		$$ = $1
		yylex.(*Parser).result = $$
	}

block
	: opt_decls
	{
		$$ = nil
	}
	| stmts opt_decls
	{
		$$ = &ast.Block{Stmts: $1}
	}

stmts
	:
	{
		$$ = nil
	}
	| opt_decls stmt
	{
		$$ = []ast.Stmt{$2}
	}
	| stmts decls stmt
	{
		$$ = append($1, $3)
	}

stmt
	: VAR var_decls
	{
		$$ = &ast.VarStmt{Vars: $2}
		$$.SetPos($1.Pos())
	}
	| expr '=' expr
	{
		$$ = &ast.AssignStmt{Lhs: $1, Rhs: $3}
		$$.SetPos($1.Pos())
	}
	| if_stmt
	{
		$$ = $1
		$$.SetPos($1.Pos())
	}
	| FOR '(' ID COLON expr ')' block END
	{
		$$ = &ast.ForStmt{Name: $3.literal, Exp: $5, Body: $7}
		$$.SetPos($1.Pos())
	}
	| WHILE '(' expr ')' block END
	{
		$$ = &ast.WhileStmt{Test: $3, Body: $5}
		$$.SetPos($1.Pos())
	}
	| RETURN
	{
		$$ = &ast.ReturnStmt{Value: nil}
		$$.SetPos($1.Pos())
	}
	| RETURN expr
	{
		$$ = &ast.ReturnStmt{Value: $2}
		$$.SetPos($1.Pos())
	}
	| BREAK
	{
		$$ = &ast.BreakStmt{}
		$$.SetPos($1.Pos())
	}
	| CONTINUE
	{
		$$ = &ast.ContinueStmt{}
		$$.SetPos($1.Pos())
	}
	| expr
	{
		$$ = &ast.ExprStmt{Exp: $1}
		$$.SetPos($1.Pos())
	}

var_decls
	: var_decl
	{
		$$ = []*ast.VarDeclStmt{$1}
	}
	| var_decls ',' var_decl
	{
		$$ = append($1, $3)
	}

var_decl
	: ID
	{
		$$ = &ast.VarDeclStmt{Name: $1.literal, Value: nil}
		$$.SetPos($1.Pos())
	}
	| ID '=' expr
	{
		$$ = &ast.VarDeclStmt{Name: $1.literal, Value: $3}
		$$.SetPos($1.Pos())
	}

if_stmt
	: IF '(' expr ')' block END
	{
		$$ = &ast.IfStmt{Test: $3, Then: $5, Else: nil}
		$$.SetPos($1.Pos())
	}
	| IF '(' expr ')' block ELSE block END
	{
		$$ = &ast.IfStmt{Test: $3, Then: $5, Else: $7}
		$$.SetPos($1.Pos())
	}
	| IF '(' expr ')' block elif_stmt END
	{
		$$ = &ast.IfStmt{Test: $3, Then: $5, Else: $6}
		$$.SetPos($1.Pos())
	}

elif_stmt
	: ELIF '(' expr ')' block
	{
		$$ = &ast.IfStmt{Test: $3, Then: $5, Else: nil}
		$$.SetPos($1.Pos())
	}
	| ELIF '(' expr ')' block ELSE block
	{
		$$ = &ast.IfStmt{Test: $3, Then: $5, Else: $7}
		$$.SetPos($1.Pos())
	}
	| ELIF '(' expr ')' block elif_stmt
	{
		$$ = &ast.IfStmt{Test: $3, Then: $5, Else: $6}
		$$.SetPos($1.Pos())
	}

expr
	: ID
	{
		$$ = &ast.RefvarExpr{Name: $1.literal}
		$$.SetPos($1.Pos())
	}
	| NIL
	{
		$$ = &ast.NilExpr{}
		$$.SetPos($1.Pos())
	}
	| TRUE
	{
		$$ = &ast.TrueExpr{}
		$$.SetPos($1.Pos())
	}
	| FALSE
	{
		$$ = &ast.FalseExpr{}
		$$.SetPos($1.Pos())
	}
	| FIXNUM
	{
		$$ = &ast.FixnumExpr{Value: $1.literal}
		$$.SetPos($1.Pos())
	}
	| FLOATNUM
	{
		$$ = &ast.FloatnumExpr{Value: $1.literal}
		$$.SetPos($1.Pos())
	}
	| STRING
	{
		$$ = &ast.StringExpr{Value: $1.literal}
		$$.SetPos($1.Pos())
	}
	| expr '(' exprlist ')'
	{
		$$ = &ast.FuncallExpr{Fun: $1, Args: $3}
		$$.SetPos($1.Pos())
	}
	| COLON namelist '{' block '}'
	{
		$$ = &ast.FunctionExpr{Parameters: $2, Body: $4}
		$$.SetPos($1.Pos())
	}
	| '{' block '}'
	{
		$$ = &ast.FunctionExpr{Parameters: []string{}, Body: $2}
		$$.SetPos($2.Pos())
	}
	| '-' expr %prec UNARY
	{
		$$ = &ast.UnaryExpr{Op: "-", Value: $2}
		$$.SetPos($2.Pos())
	}
	| '!' expr %prec UNARY
	{
		$$ = &ast.UnaryExpr{Op: "!", Value: $2}
		$$.SetPos($2.Pos())
	}
	| expr '+' expr
	{
		$$ = &ast.BinaryExpr{Op: "+", Left: $1, Right: $3}
		$$.SetPos($1.Pos())
	}
	| expr '-' expr
	{
		$$ = &ast.BinaryExpr{Op: "-", Left: $1, Right: $3}
		$$.SetPos($1.Pos())
	}
	| expr '*' expr
	{
		$$ = &ast.BinaryExpr{Op: "*", Left: $1, Right: $3}
		$$.SetPos($1.Pos())
	}
	| expr '/' expr
	{
		$$ = &ast.BinaryExpr{Op: "/", Left: $1, Right: $3}
		$$.SetPos($1.Pos())
	}
	| expr '%' expr
	{
		$$ = &ast.BinaryExpr{Op: "%", Left: $1, Right: $3}
		$$.SetPos($1.Pos())
	}
	| expr AND expr
	{
		$$ = &ast.AndExpr{Left: $1, Right: $3}
		$$.SetPos($1.Pos())
	}
	| expr OR expr
	{
		$$ = &ast.OrExpr{Left: $1, Right: $3}
		$$.SetPos($1.Pos())
	}
	| expr LT expr
	{
		$$ = &ast.BinaryExpr{Op: "<", Left: $1, Right: $3}
		$$.SetPos($1.Pos())
	}
	| expr LE expr
	{
		$$ = &ast.BinaryExpr{Op: "<=", Left: $1, Right: $3}
		$$.SetPos($1.Pos())
	}
	| expr EQ expr
	{
		$$ = &ast.BinaryExpr{Op: "==", Left: $1, Right: $3}
		$$.SetPos($1.Pos())
	}
	| expr NE expr
	{
		$$ = &ast.BinaryExpr{Op: "!=", Left: $1, Right: $3}
		$$.SetPos($1.Pos())
	}
	| expr GT expr
	{
		$$ = &ast.BinaryExpr{Op: ">", Left: $1, Right: $3}
		$$.SetPos($1.Pos())
	}
	| expr GE expr
	{
		$$ = &ast.BinaryExpr{Op: ">=", Left: $1, Right: $3}
		$$.SetPos($1.Pos())
	}
	| '(' expr ')'
	{
		$$ = $2
		$$.SetPos($2.Pos())
	}
	| expr '[' expr ']'
	{
		$$ = &ast.IndexExpr{Prefix: $1, Index: $3}
		$$.SetPos($1.Pos())
	}
	| expr '[' expr COLON expr ']'
	{
		$$ = &ast.SliceExpr{Prefix: $1, Begin: $3, End: $5}
		$$.SetPos($1.Pos())
	}
	| expr '[' expr COLON ']'
	{
		$$ = &ast.SliceExpr{Prefix: $1, Begin: $3, End: nil}
		$$.SetPos($1.Pos())
	}
	| expr '[' COLON expr ']'
	{
		$$ = &ast.SliceExpr{Prefix: $1, Begin: nil, End: $4}
		$$.SetPos($1.Pos())
	}
	| expr '[' COLON ']'
	{
		$$ = &ast.SliceExpr{Prefix: $1, Begin: nil, End: nil}
		$$.SetPos($1.Pos())
	}
	| LIST_DISPATCH exprlist ']'
	{
		$$ = &ast.ListExpr{Elements: $2}
		$$.SetPos($1.Pos())
	}
	| TABLE_DISPATCH expr_pairs opt_decls ']'
	{
		$$ = &ast.TableExpr{Elements: $2}
		$$.SetPos($1.Pos())
	}
	| TABLE_DISPATCH expr_pairs ',' opt_decls ']'
	{
		$$ = &ast.TableExpr{Elements: $2}
		$$.SetPos($1.Pos())
	}
	| LIST_DISPATCH expr DOTDOT expr ']'
	{
		$$ = &ast.RangeExpr{Start: $2, End: $4}
		$$.SetPos($1.Pos())
	}
	| LIST_DISPATCH expr DOTDOT expr ',' expr ']'
	{
		$$ = &ast.RangeExpr{Start: $2, End: $4, Step: $6}
		$$.SetPos($1.Pos())
	}

expr_pairs
	:
	{
	}
	| expr_pair
	{
		$$ = []ast.Expr{$1}
	}
	| expr_pairs ',' opt_decls expr_pair
	{
		$$ = append($1, $4)
	}

expr_pair
	: expr '=' expr
	{
		$$ = &ast.PairExpr{Key: $1, Value: $3}
		$$.SetPos($1.Pos())
	}

namelist
	:
	{
		$$ = []string{}
	}
	| ID
	{
		$$ = []string{$1.literal}
	}
	| namelist ',' opt_decls ID
	{
		$$ = append($1, $4.literal)
	}

exprlist
	:
	{
		$$ = []ast.Expr{}
	}
	| exprlist_
	{
		$$ = $1
	}

exprlist_
	: expr
	{
		$$ = []ast.Expr{$1}
	}
	| exprlist_ ',' opt_decls expr
	{
		$$ = append($1, $4)
	}

opt_decls
	:
	| decls

decls
	: decl
	| decls decl

decl
	: NL
	| ';'

%%
