package language

/* Generated from ../data/dsl/enum/cacao_intrinsic_operation.json, do not edit by hand. */

type IntrinsicOperation int

const (
	INVALID_INTRINSIC_OPERATION    IntrinsicOperation = -1
	INTRINSIC_OPERATION_I32_ADD    IntrinsicOperation = 0
	INTRINSIC_OPERATION_I32_SUB    IntrinsicOperation = 1
	INTRINSIC_OPERATION_I32_MUL    IntrinsicOperation = 2
	INTRINSIC_OPERATION_I32_DIV    IntrinsicOperation = 3
	INTRINSIC_OPERATION_I32_MOD    IntrinsicOperation = 4
	INTRINSIC_OPERATION_I32_CMP_EQ IntrinsicOperation = 5
	INTRINSIC_OPERATION_I32_CMP_NE IntrinsicOperation = 6
	INTRINSIC_OPERATION_I32_CMP_LT IntrinsicOperation = 7
	INTRINSIC_OPERATION_I32_CMP_LE IntrinsicOperation = 8
	INTRINSIC_OPERATION_I32_CMP_GT IntrinsicOperation = 9
	INTRINSIC_OPERATION_I32_CMP_GE IntrinsicOperation = 10
	INTRINSIC_OPERATION_F32_ADD    IntrinsicOperation = 11
	INTRINSIC_OPERATION_F32_SUB    IntrinsicOperation = 12
	INTRINSIC_OPERATION_F32_MUL    IntrinsicOperation = 13
	INTRINSIC_OPERATION_F32_DIV    IntrinsicOperation = 14
	INTRINSIC_OPERATION_F32_CMP_EQ IntrinsicOperation = 15
	INTRINSIC_OPERATION_F32_CMP_NE IntrinsicOperation = 16
	INTRINSIC_OPERATION_F32_CMP_LT IntrinsicOperation = 17
	INTRINSIC_OPERATION_F32_CMP_LE IntrinsicOperation = 18
	INTRINSIC_OPERATION_F32_CMP_GT IntrinsicOperation = 19
	INTRINSIC_OPERATION_F32_CMP_GE IntrinsicOperation = 20
	NUM_INTRINSIC_OPERATION        IntrinsicOperation = 21
)

type IntrinsicOperationInfo struct {
	Enum          IntrinsicOperation
	Name          string
	InfixOperator InfixOperator
	InfixLeft     IntrinsicType
	InfixRight    IntrinsicType
	InfixResult   IntrinsicType
}

var IntrinsicOperationInfos = []*IntrinsicOperationInfo{
	{
		Enum:          INTRINSIC_OPERATION_I32_ADD,
		Name:          "INTRINSIC_OPERATION_I32_ADD",
		InfixOperator: INFIX_OPERATOR_ADD,
		InfixLeft:     INTRINSIC_TYPE_I32,
		InfixRight:    INTRINSIC_TYPE_I32,
		InfixResult:   INTRINSIC_TYPE_I32,
	},
	{
		Enum:          INTRINSIC_OPERATION_I32_SUB,
		Name:          "INTRINSIC_OPERATION_I32_SUB",
		InfixOperator: INFIX_OPERATOR_SUB,
		InfixLeft:     INTRINSIC_TYPE_I32,
		InfixRight:    INTRINSIC_TYPE_I32,
		InfixResult:   INTRINSIC_TYPE_I32,
	},
	{
		Enum:          INTRINSIC_OPERATION_I32_MUL,
		Name:          "INTRINSIC_OPERATION_I32_MUL",
		InfixOperator: INFIX_OPERATOR_MUL,
		InfixLeft:     INTRINSIC_TYPE_I32,
		InfixRight:    INTRINSIC_TYPE_I32,
		InfixResult:   INTRINSIC_TYPE_I32,
	},
	{
		Enum:          INTRINSIC_OPERATION_I32_DIV,
		Name:          "INTRINSIC_OPERATION_I32_DIV",
		InfixOperator: INFIX_OPERATOR_DIV,
		InfixLeft:     INTRINSIC_TYPE_I32,
		InfixRight:    INTRINSIC_TYPE_I32,
		InfixResult:   INTRINSIC_TYPE_I32,
	},
	{
		Enum:          INTRINSIC_OPERATION_I32_MOD,
		Name:          "INTRINSIC_OPERATION_I32_MOD",
		InfixOperator: INFIX_OPERATOR_MOD,
		InfixLeft:     INTRINSIC_TYPE_I32,
		InfixRight:    INTRINSIC_TYPE_I32,
		InfixResult:   INTRINSIC_TYPE_I32,
	},
	{
		Enum:          INTRINSIC_OPERATION_I32_CMP_EQ,
		Name:          "INTRINSIC_OPERATION_I32_CMP_EQ",
		InfixOperator: INFIX_OPERATOR_CMP_EQ,
		InfixLeft:     INTRINSIC_TYPE_I32,
		InfixRight:    INTRINSIC_TYPE_I32,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_I32_CMP_NE,
		Name:          "INTRINSIC_OPERATION_I32_CMP_NE",
		InfixOperator: INFIX_OPERATOR_CMP_NE,
		InfixLeft:     INTRINSIC_TYPE_I32,
		InfixRight:    INTRINSIC_TYPE_I32,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_I32_CMP_LT,
		Name:          "INTRINSIC_OPERATION_I32_CMP_LT",
		InfixOperator: INFIX_OPERATOR_CMP_LT,
		InfixLeft:     INTRINSIC_TYPE_I32,
		InfixRight:    INTRINSIC_TYPE_I32,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_I32_CMP_LE,
		Name:          "INTRINSIC_OPERATION_I32_CMP_LE",
		InfixOperator: INFIX_OPERATOR_CMP_LE,
		InfixLeft:     INTRINSIC_TYPE_I32,
		InfixRight:    INTRINSIC_TYPE_I32,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_I32_CMP_GT,
		Name:          "INTRINSIC_OPERATION_I32_CMP_GT",
		InfixOperator: INFIX_OPERATOR_CMP_GT,
		InfixLeft:     INTRINSIC_TYPE_I32,
		InfixRight:    INTRINSIC_TYPE_I32,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_I32_CMP_GE,
		Name:          "INTRINSIC_OPERATION_I32_CMP_GE",
		InfixOperator: INFIX_OPERATOR_CMP_GE,
		InfixLeft:     INTRINSIC_TYPE_I32,
		InfixRight:    INTRINSIC_TYPE_I32,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_F32_ADD,
		Name:          "INTRINSIC_OPERATION_F32_ADD",
		InfixOperator: INFIX_OPERATOR_ADD,
		InfixLeft:     INTRINSIC_TYPE_F32,
		InfixRight:    INTRINSIC_TYPE_F32,
		InfixResult:   INTRINSIC_TYPE_F32,
	},
	{
		Enum:          INTRINSIC_OPERATION_F32_SUB,
		Name:          "INTRINSIC_OPERATION_F32_SUB",
		InfixOperator: INFIX_OPERATOR_SUB,
		InfixLeft:     INTRINSIC_TYPE_F32,
		InfixRight:    INTRINSIC_TYPE_F32,
		InfixResult:   INTRINSIC_TYPE_F32,
	},
	{
		Enum:          INTRINSIC_OPERATION_F32_MUL,
		Name:          "INTRINSIC_OPERATION_F32_MUL",
		InfixOperator: INFIX_OPERATOR_MUL,
		InfixLeft:     INTRINSIC_TYPE_F32,
		InfixRight:    INTRINSIC_TYPE_F32,
		InfixResult:   INTRINSIC_TYPE_F32,
	},
	{
		Enum:          INTRINSIC_OPERATION_F32_DIV,
		Name:          "INTRINSIC_OPERATION_F32_DIV",
		InfixOperator: INFIX_OPERATOR_DIV,
		InfixLeft:     INTRINSIC_TYPE_F32,
		InfixRight:    INTRINSIC_TYPE_F32,
		InfixResult:   INTRINSIC_TYPE_F32,
	},
	{
		Enum:          INTRINSIC_OPERATION_F32_CMP_EQ,
		Name:          "INTRINSIC_OPERATION_F32_CMP_EQ",
		InfixOperator: INFIX_OPERATOR_CMP_EQ,
		InfixLeft:     INTRINSIC_TYPE_F32,
		InfixRight:    INTRINSIC_TYPE_F32,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_F32_CMP_NE,
		Name:          "INTRINSIC_OPERATION_F32_CMP_NE",
		InfixOperator: INFIX_OPERATOR_CMP_NE,
		InfixLeft:     INTRINSIC_TYPE_F32,
		InfixRight:    INTRINSIC_TYPE_F32,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_F32_CMP_LT,
		Name:          "INTRINSIC_OPERATION_F32_CMP_LT",
		InfixOperator: INFIX_OPERATOR_CMP_LT,
		InfixLeft:     INTRINSIC_TYPE_F32,
		InfixRight:    INTRINSIC_TYPE_F32,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_F32_CMP_LE,
		Name:          "INTRINSIC_OPERATION_F32_CMP_LE",
		InfixOperator: INFIX_OPERATOR_CMP_LE,
		InfixLeft:     INTRINSIC_TYPE_F32,
		InfixRight:    INTRINSIC_TYPE_F32,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_F32_CMP_GT,
		Name:          "INTRINSIC_OPERATION_F32_CMP_GT",
		InfixOperator: INFIX_OPERATOR_CMP_GT,
		InfixLeft:     INTRINSIC_TYPE_F32,
		InfixRight:    INTRINSIC_TYPE_F32,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_F32_CMP_GE,
		Name:          "INTRINSIC_OPERATION_F32_CMP_GE",
		InfixOperator: INFIX_OPERATOR_CMP_GE,
		InfixLeft:     INTRINSIC_TYPE_F32,
		InfixRight:    INTRINSIC_TYPE_F32,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
}

func (value IntrinsicOperation) String() string {
	if value < 0 || value >= 21 {
		return "INVALID_INTRINSIC_OPERATION"
	}
	return IntrinsicOperationInfos[value].Name
}

var InfixToIntrinsic = map[InfixOperator]map[IntrinsicType]map[IntrinsicType]*IntrinsicOperationInfo{
	INFIX_OPERATOR_ADD: {
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_ADD],
		},
		INTRINSIC_TYPE_F32: {
			INTRINSIC_TYPE_F32: IntrinsicOperationInfos[INTRINSIC_OPERATION_F32_ADD],
		},
	},
	INFIX_OPERATOR_SUB: {
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_SUB],
		},
		INTRINSIC_TYPE_F32: {
			INTRINSIC_TYPE_F32: IntrinsicOperationInfos[INTRINSIC_OPERATION_F32_SUB],
		},
	},
	INFIX_OPERATOR_MUL: {
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_MUL],
		},
		INTRINSIC_TYPE_F32: {
			INTRINSIC_TYPE_F32: IntrinsicOperationInfos[INTRINSIC_OPERATION_F32_MUL],
		},
	},
	INFIX_OPERATOR_DIV: {
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_DIV],
		},
		INTRINSIC_TYPE_F32: {
			INTRINSIC_TYPE_F32: IntrinsicOperationInfos[INTRINSIC_OPERATION_F32_DIV],
		},
	},
	INFIX_OPERATOR_MOD: {
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_MOD],
		},
	},
	INFIX_OPERATOR_CMP_EQ: {
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_CMP_EQ],
		},
		INTRINSIC_TYPE_F32: {
			INTRINSIC_TYPE_F32: IntrinsicOperationInfos[INTRINSIC_OPERATION_F32_CMP_EQ],
		},
	},
	INFIX_OPERATOR_CMP_NE: {
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_CMP_NE],
		},
		INTRINSIC_TYPE_F32: {
			INTRINSIC_TYPE_F32: IntrinsicOperationInfos[INTRINSIC_OPERATION_F32_CMP_NE],
		},
	},
	INFIX_OPERATOR_CMP_LT: {
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_CMP_LT],
		},
		INTRINSIC_TYPE_F32: {
			INTRINSIC_TYPE_F32: IntrinsicOperationInfos[INTRINSIC_OPERATION_F32_CMP_LT],
		},
	},
	INFIX_OPERATOR_CMP_LE: {
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_CMP_LE],
		},
		INTRINSIC_TYPE_F32: {
			INTRINSIC_TYPE_F32: IntrinsicOperationInfos[INTRINSIC_OPERATION_F32_CMP_LE],
		},
	},
	INFIX_OPERATOR_CMP_GT: {
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_CMP_GT],
		},
		INTRINSIC_TYPE_F32: {
			INTRINSIC_TYPE_F32: IntrinsicOperationInfos[INTRINSIC_OPERATION_F32_CMP_GT],
		},
	},
	INFIX_OPERATOR_CMP_GE: {
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_CMP_GE],
		},
		INTRINSIC_TYPE_F32: {
			INTRINSIC_TYPE_F32: IntrinsicOperationInfos[INTRINSIC_OPERATION_F32_CMP_GE],
		},
	},
}
