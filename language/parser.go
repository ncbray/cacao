package language

import (
	"github.com/ncbray/cacao/framework"
)

func ConsumeWhiteSpace(state *framework.ParserState) bool {
	return state.ConsumeRune(' ') || state.ConsumeRune('\t')
}

func LineTerminator(state *framework.ParserState) bool {
	if state.ConsumeRune('\n') {
		return true
	} else if state.ConsumeRune('\r') {
		state.ConsumeRune('\n')
		return true
	}
	return false
}

func SingleLineComment(state *framework.ParserState) bool {

	if !state.ConsumeRune('/') {
		return false
	}
	if !state.ConsumeRune('/') {
		return false
	}
	for {
		current := state.Current()
		if current == '\n' || current == '\r' || !state.Consume() {
			break
		}
	}
	return true
}

func S(state *framework.ParserState) {
	state.Repeat(0, func() bool {
		return state.Choose(func() bool {
			return ConsumeWhiteSpace(state)
		}, func() bool {
			return LineTerminator(state)
		}, func() bool {
			return SingleLineComment(state)
		})
	})
}

func Cautious(state *framework.ParserState) {
	state.Repeat(0, func() bool {
		return ConsumeWhiteSpace(state)
	})
}

func EndOfStatement(state *framework.ParserState) bool {
	return state.Choose(func() bool {
		S(state)
		if state.AtEndOfStream() {
			return true
		}
		switch state.Current() {
		case ';':
			state.Consume()
			return true
		case '}', ')':
			// Lookahead
			return true
		default:
			return false
		}
	}, func() bool {
		state.Optional(func() bool {
			return SingleLineComment(state)
		})
		return LineTerminator(state)
	})
}

func ParseInteger(state *framework.ParserState) (*IntegerLiteral, bool) {
	begin := state.Pos()
	ok := state.Repeat(1, func() bool {
		return state.ConsumeRange('0', '9')
	})
	if ok {
		return &IntegerLiteral{Pos: begin, Text: state.Slice(begin)}, true
	} else {
		return nil, false
	}
}

func IsDecimalDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func IsIdentifierStart(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r == '_'
}

func IsIdentifierRest(r rune) bool {
	return IsIdentifierStart(r) || IsDecimalDigit(r)
}

func ParseKeyword(state *framework.ParserState, text string) bool {
	runes := []rune(text)
	for _, r := range runes {
		if !state.ConsumeRune(r) {
			return false
		}
	}
	return !IsIdentifierRest(state.Current())
}

func ParseIdentifier(state *framework.ParserState) (*Token, bool) {
	if !IsIdentifierStart(state.Current()) {
		return nil, false
	}
	begin := state.Pos()
	state.Consume()
	for IsIdentifierRest(state.Current()) {
		state.Consume()
	}
	return &Token{Pos: begin, Text: state.Slice(begin)}, true
}

func ParseAtom(state *framework.ParserState) (Expr, bool) {
	var result Expr
	ok := state.Choose(func() bool {
		expr, ok := ParseInteger(state)
		if !ok {
			return false
		}
		result = expr
		return true
	}, func() bool {
		if !state.ConsumeRune('(') {
			return false
		}
		S(state)
		expr, ok := ParseExpr(state)
		if !ok {
			return false
		}
		S(state)
		if !state.ConsumeRune(')') {
			return false
		}
		result = expr
		return true
	}, func() bool {
		expr, ok := ParseBlock(state)
		if !ok {
			return false
		}
		result = expr
		return true
	}, func() bool {
		expr, ok := ParseFunctionDecl(state)
		if !ok {
			return false
		}
		result = expr
		return true
	}, func() bool {
		pos := state.Pos()
		if !ParseKeyword(state, "if") {
			return false
		}
		S(state)
		cond, ok := ParseExpr(state)
		if !ok {
			return false
		}
		S(state)
		t, ok := ParseBlock(state)
		if !ok {
			return false
		}

		var f *Block = nil
		state.Optional(func() bool {
			S(state)
			if !ParseKeyword(state, "else") {
				return false
			}
			S(state)

			b, ok := ParseBlock(state)
			if !ok {
				return false
			}
			f = b
			return true
		})
		result = &If{
			Pos:  pos,
			Cond: cond,
			T:    t,
			F:    f,
		}
		return true
	}, func() bool {
		pos := state.Pos()
		if !ParseKeyword(state, "return") {
			return false
		}
		var expr Expr = nil
		state.Optional(func() bool {
			Cautious(state)
			child, ok := ParseExpr(state)
			if !ok {
				return false
			}
			expr = child
			return true
		})
		result = &Return{
			Pos:  pos,
			Expr: expr,
		}
		return true
	}, func() bool {
		name, ok := ParseIdentifier(state)
		if !ok {
			return false
		}
		result = &GetName{
			Name: name,
		}
		return true
	})
	return result, ok
}

func ParseTypeRef(state *framework.ParserState) (TypeRef, bool) {
	name, ok := ParseIdentifier(state)
	if !ok {
		return nil, false
	}
	return &TypeName{Name: name}, true
}

func ParseParam(state *framework.ParserState, named bool) (*Param, bool) {
	var name *Token = nil
	var ok bool
	if named {
		name, ok = ParseIdentifier(state)
		if !ok {
			return nil, false
		}
		S(state)
	}
	t, ok := ParseTypeRef(state)
	if !ok {
		return nil, false
	}
	return &Param{Name: name, Type: t}, true
}

func ParseParamList(state *framework.ParserState, named bool) ([]*Param, bool) {
	params := []*Param{}
	state.Optional(func() bool {
		first, ok := ParseParam(state, named)
		if !ok {
			return false
		}
		params = append(params, first)
		return state.Repeat(0, func() bool {
			S(state)
			if !state.ConsumeRune(',') {
				return false
			}
			S(state)
			param, ok := ParseParam(state, named)
			if !ok {
				return false
			}
			params = append(params, param)
			return true
		})
	})
	return params, true
}

func ParseReturnList(state *framework.ParserState) ([]*Param, bool) {
	var result []*Param
	ok := state.Choose(func() bool {
		if !state.ConsumeRune('(') {
			return false
		}
		S(state)
		var params []*Param
		ok := state.Choose(func() bool {
			temp, ok := ParseParamList(state, true)
			if !ok {
				return false
			}
			S(state)
			if !state.ConsumeRune(')') {
				return false
			}
			params = temp
			return true
		}, func() bool {
			temp, ok := ParseParamList(state, false)
			if !ok {
				return false
			}
			S(state)
			if !state.ConsumeRune(')') {
				return false
			}
			params = temp
			return true
		})
		if !ok {
			return false
		}
		result = params
		return true
	}, func() bool {
		param, ok := ParseParam(state, false)
		if !ok {
			return false
		}
		result = []*Param{param}
		return true
	}, func() bool {
		result = []*Param{}
		return true
	})
	return result, ok
}

func ParseFunctionDecl(state *framework.ParserState) (*FunctionDecl, bool) {
	if !ParseKeyword(state, "func") {
		return nil, false
	}
	S(state)
	name, ok := ParseIdentifier(state)
	if !ok {
		return nil, false
	}
	S(state)
	if !state.ConsumeRune('(') {
		return nil, false
	}
	S(state)
	params, ok := ParseParamList(state, true)
	if !ok {
		return nil, false
	}
	S(state)
	if !state.ConsumeRune(')') {
		return nil, false
	}
	S(state)
	results, ok := ParseReturnList(state)
	if !ok {
		return nil, false
	}
	S(state)
	body, ok := ParseBlock(state)
	if !ok {
		return nil, false
	}
	result := &FunctionDecl{
		Name:    name,
		Params:  params,
		Results: results,
		Body:    body,
	}
	return result, true
}

func ParsePostfixExpr(state *framework.ParserState) (Expr, bool) {
	expr, ok := ParseAtom(state)
	if !ok {
		return nil, false
	}
	state.Repeat(0, func() bool {
		Cautious(state)
		pos := state.Pos()
		if !state.ConsumeRune('(') {
			return false
		}
		S(state)
		args, ok := ParseExprList(state)
		if !ok {
			return false
		}
		S(state)
		if !state.ConsumeRune(')') {
			return false
		}
		expr = &Call{
			Pos:  pos,
			Expr: expr,
			Args: args,
		}
		return true
	})
	return expr, true
}

func ParsePrefixExpr(state *framework.ParserState) (Expr, bool) {
	var result Expr
	ok := state.Choose(func() bool {
		op, ok := ParsePrefixOperator(state)
		if !ok {
			return false
		}
		S(state)
		expr, ok := ParsePostfixExpr(state)
		if !ok {
			return false
		}
		result = &PrefixExpr{
			Op:   op,
			Expr: expr,
		}
		return true
	}, func() bool {
		expr, ok := ParsePostfixExpr(state)
		result = expr
		return ok
	})
	return result, ok
}

func ParseInfixExpr(state *framework.ParserState, min_prec int) (Expr, bool) {
	expr, ok := ParsePrefixExpr(state)
	if !ok {
		return nil, false
	}
	state.Repeat(0, func() bool {
		Cautious(state)
		pos := state.Pos()
		op, ok := ParseInfixOperator(state)
		if !ok {
			return false
		}
		info := InfixOperatorInfos[op]
		prec := info.Precedence
		if prec < min_prec {
			return false
		}
		S(state)
		other, ok := ParseInfixExpr(state, prec+1)
		if !ok {
			return false
		}
		expr = &InfixExpr{
			Pos:   pos,
			Left:  expr,
			Op:    op,
			Right: other,
		}
		return true
	})
	return expr, true
}

func ParseExpr(state *framework.ParserState) (Expr, bool) {
	return ParseInfixExpr(state, 0)
}

func ParseExprList(state *framework.ParserState) ([]Expr, bool) {
	exprs := []Expr{}
	state.Optional(func() bool {
		first, ok := ParseExpr(state)
		if !ok {
			return false
		}
		exprs = append(exprs, first)
		return state.Repeat(0, func() bool {
			S(state)
			if !state.ConsumeRune(',') {
				return false
			}
			S(state)
			expr, ok := ParseExpr(state)
			if !ok {
				return false
			}
			exprs = append(exprs, expr)
			return true
		})
	})
	return exprs, true
}

func ParseBlockBody(state *framework.ParserState) ([]Expr, bool) {
	body := []Expr{}

	state.Optional(func() bool {
		expr, ok := ParseExpr(state)
		if !ok {
			return false
		}
		body = append(body, expr)
		state.Repeat(0, func() bool {
			ok := EndOfStatement(state)
			if !ok {
				return false
			}
			S(state)
			expr, ok := ParseExpr(state)
			if !ok {
				return false
			}
			body = append(body, expr)
			return true
		})
		state.Optional(func() bool {
			return EndOfStatement(state)
		})
		return true
	})
	return body, true
}

func ParseBlock(state *framework.ParserState) (*Block, bool) {
	pos := state.Pos()
	if !state.ConsumeRune('{') {
		return nil, false
	}
	S(state)
	exprs, ok := ParseBlockBody(state)
	if !ok {
		return nil, false
	}
	S(state)
	if !state.ConsumeRune('}') {
		return nil, false
	}
	return &Block{Pos: pos, Exprs: exprs}, true
}

func ParseFile(state *framework.ParserState) (*File, bool) {
	decls := []*FunctionDecl{}
	state.Repeat(0, func() bool {
		S(state)
		decl, ok := ParseFunctionDecl(state)
		if !ok {
			return false
		}
		decls = append(decls, decl)
		return true
	})
	S(state)
	if !state.AtEndOfStream() {
		return nil, false
	}
	return &File{Decls: decls}, true
}

func Parse(stream []byte) (*File, int, bool) {
	state := framework.MakeParserState(stream)
	file, ok := ParseFile(state)
	return file, state.Deepest(), ok
}
