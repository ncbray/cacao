package language

/* Generated from ../data/dsl/enum/cacao_prefix_operator.json, do not edit by hand. */

import (
	"github.com/ncbray/cacao/framework"
)

type PrefixOperator int

const (
	INVALID_PREFIX_OPERATOR PrefixOperator = -1
	PREFIX_OPERATOR_NOT     PrefixOperator = 0
	PREFIX_OPERATOR_BIT_NOT PrefixOperator = 1
	PREFIX_OPERATOR_POS     PrefixOperator = 2
	PREFIX_OPERATOR_NEG     PrefixOperator = 3
	NUM_PREFIX_OPERATOR     PrefixOperator = 4
)

type PrefixOperatorInfo struct {
	Enum PrefixOperator
	Name string
	Text string
}

var PrefixOperatorInfos = []*PrefixOperatorInfo{
	{
		Enum: PREFIX_OPERATOR_NOT,
		Name: "PREFIX_OPERATOR_NOT",
		Text: "!",
	},
	{
		Enum: PREFIX_OPERATOR_BIT_NOT,
		Name: "PREFIX_OPERATOR_BIT_NOT",
		Text: "~",
	},
	{
		Enum: PREFIX_OPERATOR_POS,
		Name: "PREFIX_OPERATOR_POS",
		Text: "+",
	},
	{
		Enum: PREFIX_OPERATOR_NEG,
		Name: "PREFIX_OPERATOR_NEG",
		Text: "-",
	},
}

func (value PrefixOperator) String() string {
	if value < 0 || value >= 4 {
		return "INVALID_PREFIX_OPERATOR"
	}
	return PrefixOperatorInfos[value].Name
}

func ParsePrefixOperator(state *framework.ParserState) (PrefixOperator, bool) {
	switch state.Current() {
	case '!':
		state.Consume()
		return PREFIX_OPERATOR_NOT, true
	case '+':
		state.Consume()
		return PREFIX_OPERATOR_POS, true
	case '-':
		state.Consume()
		return PREFIX_OPERATOR_NEG, true
	case '~':
		state.Consume()
		return PREFIX_OPERATOR_BIT_NOT, true
	default:
		return INVALID_PREFIX_OPERATOR, false
	}
}
