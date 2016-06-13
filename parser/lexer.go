package parser

import (
	"github.com/cxxxr/hogelang/ast"
	"github.com/cxxxr/lex"
	"io"
)

type Error struct {
	Message string
}

func (err *Error) Error() string {
	return err.Message
}

type Parser struct {
	line   int
	offset int
	text   string
	lex    func() int
	result *ast.Block
	err    error
}

const (
	IGNORE = iota
	EOF    = -1
)

var keywords = map[string]int{
	"var":      VAR,
	"for":      FOR,
	"while":    WHILE,
	"if":       IF,
	"elif":     ELIF,
	"else":     ELSE,
	"end":      END,
	"return":   RETURN,
	"break":    BREAK,
	"continue": CONTINUE,
	"nil":      NIL,
	"true":     TRUE,
	"false":    FALSE,
}

func NewParser(r io.Reader) *Parser {
	parser := new(Parser)
	parser.line = 1
	parser.offset = 0
	ld := lex.NewLexerDef()
	ld.SetIgnoreValue(IGNORE)
	ld.SetEOF(func(sc *lex.Scanner) int {
		return EOF
	})
	ld.Add("$", func(sc *lex.Scanner) int {
		parser.line++
		return NL
	})
	ld.Add("[ \t]+", func(sc *lex.Scanner) int {
		parser.offset += len(sc.Text())
		return IGNORE
	})
	ld.Add("[_a-zA-Z][_a-zA-Z0-9]*", func(sc *lex.Scanner) int {
		text := sc.Text()
		parser.offset += len(text)
		val, ok := keywords[text]
		if ok {
			return val
		} else {
			parser.text = text
			return ID
		}
	})
	ld.Add("[0-9]+\\.[0-9]+", func(sc *lex.Scanner) int {
		parser.text = sc.Text()
		parser.offset += len(parser.text)
		return FLOATNUM
	})
	ld.Add("[0-9]+", func(sc *lex.Scanner) int {
		parser.text = sc.Text()
		parser.offset += len(parser.text)
		return FIXNUM
	})
	ld.Add(`"(?:[\\].|[^\\"])*"`, func(sc *lex.Scanner) int {
		text := sc.Text()
		parser.offset += len(text)
		parser.text = text[1 : len(text)-1]
		return STRING
	})
	ld.Add("&&", func(sc *lex.Scanner) int {
		parser.offset += 2
		return AND
	})
	ld.Add("\\|\\|", func(sc *lex.Scanner) int {
		parser.offset += 2
		return OR
	})
	ld.Add("<=", func(sc *lex.Scanner) int {
		parser.offset += 2
		return LE
	})
	ld.Add(">=", func(sc *lex.Scanner) int {
		parser.offset += 2
		return GE
	})
	ld.Add("<", func(sc *lex.Scanner) int {
		parser.offset += 2
		return LT
	})
	ld.Add(">", func(sc *lex.Scanner) int {
		parser.offset += 2
		return GT
	})
	ld.Add("==", func(sc *lex.Scanner) int {
		parser.offset += 2
		return EQ
	})
	ld.Add("!=", func(sc *lex.Scanner) int {
		parser.offset += 2
		return NE
	})
	ld.Add("\\[", func(sc *lex.Scanner) int {
		parser.offset++
		return LIST_DISPATCH
	})
	ld.Add("@\\[", func(sc *lex.Scanner) int {
		parser.offset++
		return TABLE_DISPATCH
	})
	ld.Add(":", func(sc *lex.Scanner) int {
		parser.offset++
		return COLON
	})
	ld.Add("\\.\\.", func(sc *lex.Scanner) int {
		parser.offset += 2
		return DOTDOT
	})
	ld.Add(".", func(sc *lex.Scanner) int {
		parser.offset++
		return int(sc.Text()[0])
	})
	parser.lex = ld.GenerateLexer(r)
	return parser
}

func (parser *Parser) Parse() (*ast.Block, error) {
	yyErrorVerbose = true
	if yyParse(parser) != 0 {
		return nil, parser.err
	} else {
		return parser.result, nil
	}
}

func (parser *Parser) Lex(lval *yySymType) int {
	t := parser.lex()
	lval.token = Token{tok: t, literal: parser.text}
	lval.token.SetPos(ast.NewPos(parser.line, parser.offset))
	return t
}

func (parser *Parser) Error(msg string) {
	parser.err = &Error{Message: msg}
}
