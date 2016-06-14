package language

/* Generated from regenerate_cacao, do not edit by hand. */

import (
	"github.com/ncbray/cacao/framework"
)

type InfixOperator int

const (
	INVALID_INFIX_OPERATOR InfixOperator = -1
	INFIX_OPERATOR_MUL     InfixOperator = 0
	INFIX_OPERATOR_DIV     InfixOperator = 1
	INFIX_OPERATOR_MOD     InfixOperator = 2
	INFIX_OPERATOR_ADD     InfixOperator = 3
	INFIX_OPERATOR_SUB     InfixOperator = 4
	INFIX_OPERATOR_BIT_LSH InfixOperator = 5
	INFIX_OPERATOR_BIT_RSH InfixOperator = 6
	INFIX_OPERATOR_BIT_AND InfixOperator = 7
	INFIX_OPERATOR_BIT_XOR InfixOperator = 8
	INFIX_OPERATOR_BIT_IOR InfixOperator = 9
	INFIX_OPERATOR_CMP_EQ  InfixOperator = 10
	INFIX_OPERATOR_CMP_NE  InfixOperator = 11
	INFIX_OPERATOR_CMP_LT  InfixOperator = 12
	INFIX_OPERATOR_CMP_LE  InfixOperator = 13
	INFIX_OPERATOR_CMP_GT  InfixOperator = 14
	INFIX_OPERATOR_CMP_GE  InfixOperator = 15
	INFIX_OPERATOR_AND     InfixOperator = 16
	INFIX_OPERATOR_OR      InfixOperator = 17
	NUM_INFIX_OPERATOR     InfixOperator = 18
)

type InfixOperatorInfo struct {
	Enum       InfixOperator
	Name       string
	Text       string
	Precedence int
}

var InfixOperatorInfos = []*InfixOperatorInfo{
	{
		Enum:       INFIX_OPERATOR_MUL,
		Name:       "INFIX_OPERATOR_MUL",
		Text:       "*",
		Precedence: 10,
	},
	{
		Enum:       INFIX_OPERATOR_DIV,
		Name:       "INFIX_OPERATOR_DIV",
		Text:       "/",
		Precedence: 10,
	},
	{
		Enum:       INFIX_OPERATOR_MOD,
		Name:       "INFIX_OPERATOR_MOD",
		Text:       "%",
		Precedence: 10,
	},
	{
		Enum:       INFIX_OPERATOR_ADD,
		Name:       "INFIX_OPERATOR_ADD",
		Text:       "+",
		Precedence: 9,
	},
	{
		Enum:       INFIX_OPERATOR_SUB,
		Name:       "INFIX_OPERATOR_SUB",
		Text:       "-",
		Precedence: 9,
	},
	{
		Enum:       INFIX_OPERATOR_BIT_LSH,
		Name:       "INFIX_OPERATOR_BIT_LSH",
		Text:       "<<",
		Precedence: 8,
	},
	{
		Enum:       INFIX_OPERATOR_BIT_RSH,
		Name:       "INFIX_OPERATOR_BIT_RSH",
		Text:       ">>",
		Precedence: 8,
	},
	{
		Enum:       INFIX_OPERATOR_BIT_AND,
		Name:       "INFIX_OPERATOR_BIT_AND",
		Text:       "&",
		Precedence: 7,
	},
	{
		Enum:       INFIX_OPERATOR_BIT_XOR,
		Name:       "INFIX_OPERATOR_BIT_XOR",
		Text:       "^",
		Precedence: 6,
	},
	{
		Enum:       INFIX_OPERATOR_BIT_IOR,
		Name:       "INFIX_OPERATOR_BIT_IOR",
		Text:       "|",
		Precedence: 5,
	},
	{
		Enum:       INFIX_OPERATOR_CMP_EQ,
		Name:       "INFIX_OPERATOR_CMP_EQ",
		Text:       "==",
		Precedence: 4,
	},
	{
		Enum:       INFIX_OPERATOR_CMP_NE,
		Name:       "INFIX_OPERATOR_CMP_NE",
		Text:       "!=",
		Precedence: 4,
	},
	{
		Enum:       INFIX_OPERATOR_CMP_LT,
		Name:       "INFIX_OPERATOR_CMP_LT",
		Text:       "<",
		Precedence: 4,
	},
	{
		Enum:       INFIX_OPERATOR_CMP_LE,
		Name:       "INFIX_OPERATOR_CMP_LE",
		Text:       "<=",
		Precedence: 4,
	},
	{
		Enum:       INFIX_OPERATOR_CMP_GT,
		Name:       "INFIX_OPERATOR_CMP_GT",
		Text:       ">",
		Precedence: 4,
	},
	{
		Enum:       INFIX_OPERATOR_CMP_GE,
		Name:       "INFIX_OPERATOR_CMP_GE",
		Text:       ">=",
		Precedence: 4,
	},
	{
		Enum:       INFIX_OPERATOR_AND,
		Name:       "INFIX_OPERATOR_AND",
		Text:       "&&",
		Precedence: 3,
	},
	{
		Enum:       INFIX_OPERATOR_OR,
		Name:       "INFIX_OPERATOR_OR",
		Text:       "||",
		Precedence: 2,
	},
}

func (value InfixOperator) String() string {
	if value < 0 || value >= 18 {
		return "INVALID_INFIX_OPERATOR"
	}
	return InfixOperatorInfos[value].Name
}

func ParseInfixOperator(state *framework.ParserState) (InfixOperator, bool) {
	switch state.Current() {
	case '!':
		state.Consume()
		switch state.Current() {
		case '=':
			state.Consume()
			return INFIX_OPERATOR_CMP_NE, true
		default:
			return INVALID_INFIX_OPERATOR, false
		}
	case '%':
		state.Consume()
		return INFIX_OPERATOR_MOD, true
	case '&':
		state.Consume()
		switch state.Current() {
		case '&':
			state.Consume()
			return INFIX_OPERATOR_AND, true
		default:
			return INFIX_OPERATOR_BIT_AND, true
		}
	case '*':
		state.Consume()
		return INFIX_OPERATOR_MUL, true
	case '+':
		state.Consume()
		return INFIX_OPERATOR_ADD, true
	case '-':
		state.Consume()
		return INFIX_OPERATOR_SUB, true
	case '/':
		state.Consume()
		return INFIX_OPERATOR_DIV, true
	case '<':
		state.Consume()
		switch state.Current() {
		case '<':
			state.Consume()
			return INFIX_OPERATOR_BIT_LSH, true
		case '=':
			state.Consume()
			return INFIX_OPERATOR_CMP_LE, true
		default:
			return INFIX_OPERATOR_CMP_LT, true
		}
	case '=':
		state.Consume()
		switch state.Current() {
		case '=':
			state.Consume()
			return INFIX_OPERATOR_CMP_EQ, true
		default:
			return INVALID_INFIX_OPERATOR, false
		}
	case '>':
		state.Consume()
		switch state.Current() {
		case '=':
			state.Consume()
			return INFIX_OPERATOR_CMP_GE, true
		case '>':
			state.Consume()
			return INFIX_OPERATOR_BIT_RSH, true
		default:
			return INFIX_OPERATOR_CMP_GT, true
		}
	case '^':
		state.Consume()
		return INFIX_OPERATOR_BIT_XOR, true
	case '|':
		state.Consume()
		switch state.Current() {
		case '|':
			state.Consume()
			return INFIX_OPERATOR_OR, true
		default:
			return INFIX_OPERATOR_BIT_IOR, true
		}
	default:
		return INVALID_INFIX_OPERATOR, false
	}
}
