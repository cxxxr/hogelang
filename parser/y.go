//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"github.com/cxxxr/hogelang/ast"
)

type Token struct {
	ast.Ast
	tok     int
	literal string
}

//line parser.go.y:16
type yySymType struct {
	yys        int
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

const VAR = 57346
const FOR = 57347
const WHILE = 57348
const IF = 57349
const ELIF = 57350
const ELSE = 57351
const END = 57352
const RETURN = 57353
const BREAK = 57354
const CONTINUE = 57355
const NIL = 57356
const TRUE = 57357
const FALSE = 57358
const ID = 57359
const FIXNUM = 57360
const FLOATNUM = 57361
const STRING = 57362
const AND = 57363
const OR = 57364
const LT = 57365
const LE = 57366
const EQ = 57367
const NE = 57368
const GT = 57369
const GE = 57370
const NL = 57371
const LIST_DISPATCH = 57372
const TABLE_DISPATCH = 57373
const COLON = 57374
const DOTDOT = 57375
const UNARY = 57376

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"VAR",
	"FOR",
	"WHILE",
	"IF",
	"ELIF",
	"ELSE",
	"END",
	"RETURN",
	"BREAK",
	"CONTINUE",
	"NIL",
	"TRUE",
	"FALSE",
	"ID",
	"FIXNUM",
	"FLOATNUM",
	"STRING",
	"AND",
	"OR",
	"LT",
	"LE",
	"EQ",
	"NE",
	"GT",
	"GE",
	"NL",
	"LIST_DISPATCH",
	"TABLE_DISPATCH",
	"COLON",
	"DOTDOT",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"UNARY",
	"'='",
	"'('",
	"')'",
	"','",
	"'{'",
	"'}'",
	"'!'",
	"'['",
	"']'",
	"';'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:445

//line yacctab:1
var yyExca = [...]int{
	-1, 0,
	1, 4,
	-2, 74,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 26,
	45, 4,
	-2, 74,
	-1, 94,
	45, 4,
	-2, 74,
	-1, 112,
	43, 68,
	44, 68,
	-2, 27,
	-1, 114,
	10, 4,
	-2, 74,
	-1, 122,
	8, 4,
	9, 4,
	10, 4,
	-2, 74,
	-1, 137,
	10, 4,
	-2, 74,
	-1, 141,
	10, 4,
	-2, 74,
	-1, 152,
	8, 4,
	9, 4,
	10, 4,
	-2, 74,
	-1, 154,
	10, 4,
	-2, 74,
}

const yyNprod = 80
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 795

var yyAct = [...]int{

	11, 142, 37, 68, 119, 3, 98, 7, 2, 128,
	33, 41, 42, 43, 44, 45, 57, 7, 40, 96,
	40, 102, 148, 100, 54, 72, 54, 8, 61, 62,
	63, 65, 69, 95, 94, 60, 107, 8, 70, 56,
	74, 76, 77, 78, 79, 80, 81, 82, 83, 84,
	85, 86, 87, 88, 89, 90, 55, 93, 64, 73,
	9, 113, 129, 46, 47, 48, 49, 50, 51, 52,
	53, 104, 38, 101, 106, 105, 41, 42, 43, 44,
	45, 92, 59, 40, 150, 131, 43, 44, 45, 54,
	130, 40, 110, 143, 154, 71, 149, 54, 147, 75,
	117, 116, 138, 115, 121, 66, 118, 58, 120, 67,
	123, 143, 141, 140, 126, 5, 6, 12, 36, 132,
	34, 69, 35, 127, 134, 4, 1, 0, 0, 48,
	49, 135, 139, 52, 53, 0, 0, 0, 0, 0,
	41, 42, 43, 44, 45, 0, 144, 40, 0, 151,
	146, 35, 0, 54, 0, 155, 0, 10, 13, 14,
	32, 153, 0, 156, 15, 16, 17, 19, 20, 21,
	18, 22, 23, 24, 0, 0, 0, 0, 0, 0,
	0, 0, 7, 30, 31, 25, 0, 0, 27, 0,
	0, 0, 0, 0, 29, 0, 0, 26, 0, 28,
	0, 0, 8, 10, 13, 14, 32, 0, 0, 0,
	15, 16, 17, 19, 20, 21, 18, 22, 23, 24,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 30,
	31, 25, 0, 0, 27, 0, 0, 0, 0, 0,
	29, 0, 0, 26, 0, 28, 46, 47, 48, 49,
	50, 51, 52, 53, 0, 0, 0, 109, 0, 41,
	42, 43, 44, 45, 0, 0, 40, 0, 0, 0,
	0, 0, 54, 108, 46, 47, 48, 49, 50, 51,
	52, 53, 0, 0, 0, 0, 0, 41, 42, 43,
	44, 45, 0, 0, 40, 0, 0, 0, 0, 0,
	54, 145, 46, 47, 48, 49, 50, 51, 52, 53,
	0, 0, 0, 0, 0, 41, 42, 43, 44, 45,
	0, 0, 40, 0, 0, 0, 0, 0, 54, 136,
	46, 47, 48, 49, 50, 51, 52, 53, 0, 0,
	0, 0, 0, 41, 42, 43, 44, 45, 0, 0,
	40, 0, 0, 0, 0, 0, 54, 125, 46, 47,
	48, 49, 50, 51, 52, 53, 0, 0, 0, 0,
	0, 41, 42, 43, 44, 45, 0, 0, 40, 152,
	0, 0, 0, 0, 54, 46, 47, 48, 49, 50,
	51, 52, 53, 0, 0, 0, 0, 0, 41, 42,
	43, 44, 45, 0, 0, 40, 137, 0, 0, 0,
	0, 54, 19, 20, 21, 18, 22, 23, 24, 0,
	19, 20, 21, 18, 22, 23, 24, 0, 30, 31,
	25, 0, 0, 27, 0, 0, 30, 31, 25, 29,
	0, 27, 26, 0, 28, 0, 133, 29, 0, 0,
	26, 0, 28, 0, 124, 46, 47, 48, 49, 50,
	51, 52, 53, 0, 0, 0, 0, 0, 41, 42,
	43, 44, 45, 0, 0, 40, 122, 0, 0, 0,
	0, 54, 46, 47, 48, 49, 50, 51, 52, 53,
	0, 0, 0, 0, 0, 41, 42, 43, 44, 45,
	0, 0, 40, 114, 0, 0, 0, 0, 54, 19,
	20, 21, 112, 22, 23, 24, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 30, 31, 25, 0, 0,
	27, 0, 0, 0, 0, 0, 29, 0, 0, 26,
	0, 28, 0, 111, 46, 47, 48, 49, 50, 51,
	52, 53, 0, 0, 0, 0, 0, 41, 42, 43,
	44, 45, 0, 103, 40, 0, 0, 0, 0, 0,
	54, 46, 47, 48, 49, 50, 51, 52, 53, 0,
	0, 0, 0, 99, 41, 42, 43, 44, 45, 0,
	0, 40, 0, 0, 0, 0, 0, 54, 46, 47,
	48, 49, 50, 51, 52, 53, 0, 0, 0, 0,
	0, 41, 42, 43, 44, 45, 0, 0, 40, 97,
	0, 0, 0, 0, 54, 46, 47, 48, 49, 50,
	51, 52, 53, 0, 0, 0, 0, 0, 41, 42,
	43, 44, 45, 0, 39, 40, 0, 0, 0, 0,
	0, 54, 46, 47, 48, 49, 50, 51, 52, 53,
	0, 0, 0, 0, 0, 41, 42, 43, 44, 45,
	0, 0, 40, 0, 0, 0, 0, 46, 54, 48,
	49, 50, 51, 52, 53, 0, 0, 0, 0, 0,
	41, 42, 43, 44, 45, 0, 0, 40, 0, 0,
	0, 0, 0, 54, 19, 20, 21, 18, 22, 23,
	24, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	30, 31, 25, 0, 0, 27, 0, 0, 0, 0,
	0, 29, 0, 0, 26, 0, 28, 48, 49, 50,
	51, 52, 53, 0, 0, 0, 0, 0, 41, 42,
	43, 44, 45, 0, 0, 40, 0, 0, 0, 0,
	0, 54, 19, 20, 21, 18, 22, 23, 24, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 30, 31,
	91, 0, 0, 27, 0, 0, 0, 0, 0, 29,
	0, 0, 26, 0, 28,
}
var yyPact = [...]int{

	-12, -1000, -1000, 199, -12, -12, -1000, -1000, -1000, -1000,
	55, 604, -1000, 15, -2, 690, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 65, -12, 690, 690, 690,
	690, 690, -3, -1000, 153, -1000, -18, -1000, 19, 690,
	690, 690, 690, 690, 690, 690, 690, 690, 690, 690,
	690, 690, 690, 690, 748, 64, 690, 631, -10, -1000,
	-26, -21, -21, 577, -42, 550, -20, -22, -1000, 523,
	690, -1000, 55, 690, 631, -6, 631, 50, 50, -21,
	-21, -21, 714, 656, -23, -23, 106, 106, -23, -23,
	225, 495, 29, 461, -12, -12, -1000, -1000, -1000, 690,
	-12, -44, -12, 690, 434, -1000, 631, -1000, -1000, 406,
	309, -1000, -1000, 690, -12, -36, 45, 42, 690, -1000,
	398, 631, -12, 281, -1000, -1000, 364, 92, -1000, -1000,
	-1000, 690, 631, -1000, -1000, 103, -1000, -12, -1000, 253,
	-1000, -12, 88, -19, 86, -1000, 74, -1000, 690, -1000,
	-1000, 337, -12, 85, -12, -1000, -1000,
}
var yyPgo = [...]int{

	0, 126, 8, 125, 60, 118, 2, 117, 1, 0,
	3, 109, 107, 58, 105, 5, 115, 116,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 3, 3, 3, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 5, 5, 6,
	6, 7, 7, 7, 8, 8, 8, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 11, 11, 11, 10, 12, 12, 12,
	13, 13, 14, 14, 15, 15, 16, 16, 17, 17,
}
var yyR2 = [...]int{

	0, 1, 1, 2, 0, 2, 3, 2, 3, 1,
	8, 6, 1, 2, 1, 1, 1, 1, 3, 1,
	3, 6, 8, 7, 5, 7, 6, 1, 1, 1,
	1, 1, 1, 1, 4, 5, 3, 2, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 6, 5, 5, 4, 3, 4,
	5, 5, 7, 0, 1, 4, 3, 0, 1, 4,
	0, 1, 1, 4, 0, 1, 1, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -2, -15, -3, -16, -17, 29, 49, -4,
	4, -9, -7, 5, 6, 11, 12, 13, 17, 14,
	15, 16, 18, 19, 20, 32, 44, 35, 46, 41,
	30, 31, 7, -15, -16, -17, -5, -6, 17, 40,
	41, 34, 35, 36, 37, 38, 21, 22, 23, 24,
	25, 26, 27, 28, 47, 41, 41, -9, -12, 17,
	-2, -9, -9, -9, -13, -9, -14, -11, -10, -9,
	41, -4, 43, 40, -9, -13, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, -9, -9, -9,
	-9, 32, 17, -9, 44, 43, 45, 42, 48, 33,
	43, -15, 43, 40, -9, -6, -9, 42, 48, 32,
	-9, 48, 17, 32, 42, -2, -15, -9, -15, 48,
	-15, -9, 42, -9, 48, 48, -9, -2, 45, 17,
	48, 43, -9, 48, -10, -2, 48, 42, 10, -9,
	10, 9, -8, 8, -2, 48, -2, 10, 41, 10,
	10, -9, 42, -2, 9, -8, -2,
}
var yyDef = [...]int{

	-2, -2, 1, 2, 74, 75, 76, 78, 79, 5,
	0, 16, 9, 0, 0, 12, 14, 15, 27, 28,
	29, 30, 31, 32, 33, 67, -2, 0, 0, 0,
	70, 63, 0, 3, 75, 77, 7, 17, 19, 0,
	70, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 13, 0, 68,
	0, 37, 38, 0, 0, 72, 71, 74, 64, 0,
	0, 6, 0, 0, 8, 0, 72, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	0, 67, 0, 0, -2, 74, 36, 52, 58, 0,
	74, 0, 74, 0, 0, 18, 20, 34, 53, 0,
	0, 57, -2, 0, -2, 0, 0, 0, 0, 59,
	0, 66, -2, 0, 55, 56, 0, 0, 35, 69,
	61, 0, 73, 60, 65, 0, 54, -2, 11, 0,
	21, -2, 0, 0, 0, 62, 0, 23, 0, 10,
	22, 0, -2, 24, -2, 26, 25,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 46, 3, 3, 3, 38, 3, 3,
	41, 42, 36, 34, 43, 35, 3, 37, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 49,
	3, 40, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 47, 3, 48, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 44, 3, 45,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 39,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:60
		{
			yyVAL.block = yyDollar[1].block
			yylex.(*Parser).result = yyVAL.block
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:67
		{
			yyVAL.block = nil
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:71
		{
			yyVAL.block = &ast.Block{Stmts: yyDollar[1].stmts}
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:77
		{
			yyVAL.stmts = nil
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:81
		{
			yyVAL.stmts = []ast.Stmt{yyDollar[2].stmt}
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:85
		{
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[3].stmt)
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:91
		{
			yyVAL.stmt = &ast.VarStmt{Vars: yyDollar[2].var_decls}
			yyVAL.stmt.SetPos(yyDollar[1].token.Pos())
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:96
		{
			yyVAL.stmt = &ast.AssignStmt{Lhs: yyDollar[1].expr, Rhs: yyDollar[3].expr}
			yyVAL.stmt.SetPos(yyDollar[1].expr.Pos())
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:101
		{
			yyVAL.stmt = yyDollar[1].if_stmt
			yyVAL.stmt.SetPos(yyDollar[1].if_stmt.Pos())
		}
	case 10:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:106
		{
			yyVAL.stmt = &ast.ForStmt{Name: yyDollar[3].token.literal, Exp: yyDollar[5].expr, Body: yyDollar[7].block}
			yyVAL.stmt.SetPos(yyDollar[1].token.Pos())
		}
	case 11:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:111
		{
			yyVAL.stmt = &ast.WhileStmt{Test: yyDollar[3].expr, Body: yyDollar[5].block}
			yyVAL.stmt.SetPos(yyDollar[1].token.Pos())
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:116
		{
			yyVAL.stmt = &ast.ReturnStmt{Value: nil}
			yyVAL.stmt.SetPos(yyDollar[1].token.Pos())
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:121
		{
			yyVAL.stmt = &ast.ReturnStmt{Value: yyDollar[2].expr}
			yyVAL.stmt.SetPos(yyDollar[1].token.Pos())
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:126
		{
			yyVAL.stmt = &ast.BreakStmt{}
			yyVAL.stmt.SetPos(yyDollar[1].token.Pos())
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:131
		{
			yyVAL.stmt = &ast.ContinueStmt{}
			yyVAL.stmt.SetPos(yyDollar[1].token.Pos())
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:136
		{
			yyVAL.stmt = &ast.ExprStmt{Exp: yyDollar[1].expr}
			yyVAL.stmt.SetPos(yyDollar[1].expr.Pos())
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:143
		{
			yyVAL.var_decls = []*ast.VarDeclStmt{yyDollar[1].var_decl}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:147
		{
			yyVAL.var_decls = append(yyDollar[1].var_decls, yyDollar[3].var_decl)
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:153
		{
			yyVAL.var_decl = &ast.VarDeclStmt{Name: yyDollar[1].token.literal, Value: nil}
			yyVAL.var_decl.SetPos(yyDollar[1].token.Pos())
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:158
		{
			yyVAL.var_decl = &ast.VarDeclStmt{Name: yyDollar[1].token.literal, Value: yyDollar[3].expr}
			yyVAL.var_decl.SetPos(yyDollar[1].token.Pos())
		}
	case 21:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:165
		{
			yyVAL.if_stmt = &ast.IfStmt{Test: yyDollar[3].expr, Then: yyDollar[5].block, Else: nil}
			yyVAL.if_stmt.SetPos(yyDollar[1].token.Pos())
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:170
		{
			yyVAL.if_stmt = &ast.IfStmt{Test: yyDollar[3].expr, Then: yyDollar[5].block, Else: yyDollar[7].block}
			yyVAL.if_stmt.SetPos(yyDollar[1].token.Pos())
		}
	case 23:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:175
		{
			yyVAL.if_stmt = &ast.IfStmt{Test: yyDollar[3].expr, Then: yyDollar[5].block, Else: yyDollar[6].if_stmt}
			yyVAL.if_stmt.SetPos(yyDollar[1].token.Pos())
		}
	case 24:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:182
		{
			yyVAL.if_stmt = &ast.IfStmt{Test: yyDollar[3].expr, Then: yyDollar[5].block, Else: nil}
			yyVAL.if_stmt.SetPos(yyDollar[1].token.Pos())
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:187
		{
			yyVAL.if_stmt = &ast.IfStmt{Test: yyDollar[3].expr, Then: yyDollar[5].block, Else: yyDollar[7].block}
			yyVAL.if_stmt.SetPos(yyDollar[1].token.Pos())
		}
	case 26:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:192
		{
			yyVAL.if_stmt = &ast.IfStmt{Test: yyDollar[3].expr, Then: yyDollar[5].block, Else: yyDollar[6].if_stmt}
			yyVAL.if_stmt.SetPos(yyDollar[1].token.Pos())
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:199
		{
			yyVAL.expr = &ast.RefvarExpr{Name: yyDollar[1].token.literal}
			yyVAL.expr.SetPos(yyDollar[1].token.Pos())
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:204
		{
			yyVAL.expr = &ast.NilExpr{}
			yyVAL.expr.SetPos(yyDollar[1].token.Pos())
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:209
		{
			yyVAL.expr = &ast.TrueExpr{}
			yyVAL.expr.SetPos(yyDollar[1].token.Pos())
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:214
		{
			yyVAL.expr = &ast.FalseExpr{}
			yyVAL.expr.SetPos(yyDollar[1].token.Pos())
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:219
		{
			yyVAL.expr = &ast.FixnumExpr{Value: yyDollar[1].token.literal}
			yyVAL.expr.SetPos(yyDollar[1].token.Pos())
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:224
		{
			yyVAL.expr = &ast.FloatnumExpr{Value: yyDollar[1].token.literal}
			yyVAL.expr.SetPos(yyDollar[1].token.Pos())
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:229
		{
			yyVAL.expr = &ast.StringExpr{Value: yyDollar[1].token.literal}
			yyVAL.expr.SetPos(yyDollar[1].token.Pos())
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:234
		{
			yyVAL.expr = &ast.FuncallExpr{Fun: yyDollar[1].expr, Args: yyDollar[3].exprlist}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:239
		{
			yyVAL.expr = &ast.FunctionExpr{Parameters: yyDollar[2].namelist, Body: yyDollar[4].block}
			yyVAL.expr.SetPos(yyDollar[1].token.Pos())
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:244
		{
			yyVAL.expr = &ast.FunctionExpr{Parameters: []string{}, Body: yyDollar[2].block}
			yyVAL.expr.SetPos(yyDollar[2].block.Pos())
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:249
		{
			yyVAL.expr = &ast.UnaryExpr{Op: "-", Value: yyDollar[2].expr}
			yyVAL.expr.SetPos(yyDollar[2].expr.Pos())
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:254
		{
			yyVAL.expr = &ast.UnaryExpr{Op: "!", Value: yyDollar[2].expr}
			yyVAL.expr.SetPos(yyDollar[2].expr.Pos())
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:259
		{
			yyVAL.expr = &ast.BinaryExpr{Op: "+", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:264
		{
			yyVAL.expr = &ast.BinaryExpr{Op: "-", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:269
		{
			yyVAL.expr = &ast.BinaryExpr{Op: "*", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:274
		{
			yyVAL.expr = &ast.BinaryExpr{Op: "/", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:279
		{
			yyVAL.expr = &ast.BinaryExpr{Op: "%", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:284
		{
			yyVAL.expr = &ast.AndExpr{Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:289
		{
			yyVAL.expr = &ast.OrExpr{Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:294
		{
			yyVAL.expr = &ast.BinaryExpr{Op: "<", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:299
		{
			yyVAL.expr = &ast.BinaryExpr{Op: "<=", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:304
		{
			yyVAL.expr = &ast.BinaryExpr{Op: "==", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:309
		{
			yyVAL.expr = &ast.BinaryExpr{Op: "!=", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:314
		{
			yyVAL.expr = &ast.BinaryExpr{Op: ">", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:319
		{
			yyVAL.expr = &ast.BinaryExpr{Op: ">=", Left: yyDollar[1].expr, Right: yyDollar[3].expr}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:324
		{
			yyVAL.expr = yyDollar[2].expr
			yyVAL.expr.SetPos(yyDollar[2].expr.Pos())
		}
	case 53:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:329
		{
			yyVAL.expr = &ast.IndexExpr{Prefix: yyDollar[1].expr, Index: yyDollar[3].expr}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 54:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:334
		{
			yyVAL.expr = &ast.SliceExpr{Prefix: yyDollar[1].expr, Begin: yyDollar[3].expr, End: yyDollar[5].expr}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 55:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:339
		{
			yyVAL.expr = &ast.SliceExpr{Prefix: yyDollar[1].expr, Begin: yyDollar[3].expr, End: nil}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 56:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:344
		{
			yyVAL.expr = &ast.SliceExpr{Prefix: yyDollar[1].expr, Begin: nil, End: yyDollar[4].expr}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 57:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:349
		{
			yyVAL.expr = &ast.SliceExpr{Prefix: yyDollar[1].expr, Begin: nil, End: nil}
			yyVAL.expr.SetPos(yyDollar[1].expr.Pos())
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:354
		{
			yyVAL.expr = &ast.ListExpr{Elements: yyDollar[2].exprlist}
			yyVAL.expr.SetPos(yyDollar[1].token.Pos())
		}
	case 59:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:359
		{
			yyVAL.expr = &ast.TableExpr{Elements: yyDollar[2].expr_pairs}
			yyVAL.expr.SetPos(yyDollar[1].token.Pos())
		}
	case 60:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:364
		{
			yyVAL.expr = &ast.TableExpr{Elements: yyDollar[2].expr_pairs}
			yyVAL.expr.SetPos(yyDollar[1].token.Pos())
		}
	case 61:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:369
		{
			yyVAL.expr = &ast.RangeExpr{Start: yyDollar[2].expr, End: yyDollar[4].expr}
			yyVAL.expr.SetPos(yyDollar[1].token.Pos())
		}
	case 62:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:374
		{
			yyVAL.expr = &ast.RangeExpr{Start: yyDollar[2].expr, End: yyDollar[4].expr, Step: yyDollar[6].expr}
			yyVAL.expr.SetPos(yyDollar[1].token.Pos())
		}
	case 63:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:381
		{
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:384
		{
			yyVAL.expr_pairs = []ast.Expr{yyDollar[1].expr_pair}
		}
	case 65:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:388
		{
			yyVAL.expr_pairs = append(yyDollar[1].expr_pairs, yyDollar[4].expr_pair)
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:394
		{
			yyVAL.expr_pair = &ast.PairExpr{Key: yyDollar[1].expr, Value: yyDollar[3].expr}
			yyVAL.expr_pair.SetPos(yyDollar[1].expr.Pos())
		}
	case 67:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:401
		{
			yyVAL.namelist = []string{}
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:405
		{
			yyVAL.namelist = []string{yyDollar[1].token.literal}
		}
	case 69:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:409
		{
			yyVAL.namelist = append(yyDollar[1].namelist, yyDollar[4].token.literal)
		}
	case 70:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:415
		{
			yyVAL.exprlist = []ast.Expr{}
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:419
		{
			yyVAL.exprlist = yyDollar[1].exprlist
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:425
		{
			yyVAL.exprlist = []ast.Expr{yyDollar[1].expr}
		}
	case 73:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:429
		{
			yyVAL.exprlist = append(yyDollar[1].exprlist, yyDollar[4].expr)
		}
	}
	goto yystack /* stack new state and value */
}
