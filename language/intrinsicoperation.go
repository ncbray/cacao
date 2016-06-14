package language

/* Generated from regenerate_cacao, do not edit by hand. */

type IntrinsicOperation int

const (
	INVALID_INTRINSIC_OPERATION      IntrinsicOperation = -1
	INTRINSIC_OPERATION_BOOL_BIT_AND IntrinsicOperation = 0
	INTRINSIC_OPERATION_BOOL_BIT_XOR IntrinsicOperation = 1
	INTRINSIC_OPERATION_BOOL_BIT_IOR IntrinsicOperation = 2
	INTRINSIC_OPERATION_BOOL_CMP_EQ  IntrinsicOperation = 3
	INTRINSIC_OPERATION_BOOL_CMP_NE  IntrinsicOperation = 4
	INTRINSIC_OPERATION_BOOL_AND     IntrinsicOperation = 5
	INTRINSIC_OPERATION_BOOL_OR      IntrinsicOperation = 6
	INTRINSIC_OPERATION_I8_MUL       IntrinsicOperation = 7
	INTRINSIC_OPERATION_I8_DIV       IntrinsicOperation = 8
	INTRINSIC_OPERATION_I8_MOD       IntrinsicOperation = 9
	INTRINSIC_OPERATION_I8_ADD       IntrinsicOperation = 10
	INTRINSIC_OPERATION_I8_SUB       IntrinsicOperation = 11
	INTRINSIC_OPERATION_I8_BIT_LSH   IntrinsicOperation = 12
	INTRINSIC_OPERATION_I8_BIT_RSH   IntrinsicOperation = 13
	INTRINSIC_OPERATION_I8_BIT_AND   IntrinsicOperation = 14
	INTRINSIC_OPERATION_I8_BIT_XOR   IntrinsicOperation = 15
	INTRINSIC_OPERATION_I8_BIT_IOR   IntrinsicOperation = 16
	INTRINSIC_OPERATION_I8_CMP_EQ    IntrinsicOperation = 17
	INTRINSIC_OPERATION_I8_CMP_NE    IntrinsicOperation = 18
	INTRINSIC_OPERATION_I8_CMP_LT    IntrinsicOperation = 19
	INTRINSIC_OPERATION_I8_CMP_LE    IntrinsicOperation = 20
	INTRINSIC_OPERATION_I8_CMP_GT    IntrinsicOperation = 21
	INTRINSIC_OPERATION_I8_CMP_GE    IntrinsicOperation = 22
	INTRINSIC_OPERATION_I16_MUL      IntrinsicOperation = 23
	INTRINSIC_OPERATION_I16_DIV      IntrinsicOperation = 24
	INTRINSIC_OPERATION_I16_MOD      IntrinsicOperation = 25
	INTRINSIC_OPERATION_I16_ADD      IntrinsicOperation = 26
	INTRINSIC_OPERATION_I16_SUB      IntrinsicOperation = 27
	INTRINSIC_OPERATION_I16_BIT_LSH  IntrinsicOperation = 28
	INTRINSIC_OPERATION_I16_BIT_RSH  IntrinsicOperation = 29
	INTRINSIC_OPERATION_I16_BIT_AND  IntrinsicOperation = 30
	INTRINSIC_OPERATION_I16_BIT_XOR  IntrinsicOperation = 31
	INTRINSIC_OPERATION_I16_BIT_IOR  IntrinsicOperation = 32
	INTRINSIC_OPERATION_I16_CMP_EQ   IntrinsicOperation = 33
	INTRINSIC_OPERATION_I16_CMP_NE   IntrinsicOperation = 34
	INTRINSIC_OPERATION_I16_CMP_LT   IntrinsicOperation = 35
	INTRINSIC_OPERATION_I16_CMP_LE   IntrinsicOperation = 36
	INTRINSIC_OPERATION_I16_CMP_GT   IntrinsicOperation = 37
	INTRINSIC_OPERATION_I16_CMP_GE   IntrinsicOperation = 38
	INTRINSIC_OPERATION_I32_MUL      IntrinsicOperation = 39
	INTRINSIC_OPERATION_I32_DIV      IntrinsicOperation = 40
	INTRINSIC_OPERATION_I32_MOD      IntrinsicOperation = 41
	INTRINSIC_OPERATION_I32_ADD      IntrinsicOperation = 42
	INTRINSIC_OPERATION_I32_SUB      IntrinsicOperation = 43
	INTRINSIC_OPERATION_I32_BIT_LSH  IntrinsicOperation = 44
	INTRINSIC_OPERATION_I32_BIT_RSH  IntrinsicOperation = 45
	INTRINSIC_OPERATION_I32_BIT_AND  IntrinsicOperation = 46
	INTRINSIC_OPERATION_I32_BIT_XOR  IntrinsicOperation = 47
	INTRINSIC_OPERATION_I32_BIT_IOR  IntrinsicOperation = 48
	INTRINSIC_OPERATION_I32_CMP_EQ   IntrinsicOperation = 49
	INTRINSIC_OPERATION_I32_CMP_NE   IntrinsicOperation = 50
	INTRINSIC_OPERATION_I32_CMP_LT   IntrinsicOperation = 51
	INTRINSIC_OPERATION_I32_CMP_LE   IntrinsicOperation = 52
	INTRINSIC_OPERATION_I32_CMP_GT   IntrinsicOperation = 53
	INTRINSIC_OPERATION_I32_CMP_GE   IntrinsicOperation = 54
	INTRINSIC_OPERATION_I64_MUL      IntrinsicOperation = 55
	INTRINSIC_OPERATION_I64_DIV      IntrinsicOperation = 56
	INTRINSIC_OPERATION_I64_MOD      IntrinsicOperation = 57
	INTRINSIC_OPERATION_I64_ADD      IntrinsicOperation = 58
	INTRINSIC_OPERATION_I64_SUB      IntrinsicOperation = 59
	INTRINSIC_OPERATION_I64_BIT_LSH  IntrinsicOperation = 60
	INTRINSIC_OPERATION_I64_BIT_RSH  IntrinsicOperation = 61
	INTRINSIC_OPERATION_I64_BIT_AND  IntrinsicOperation = 62
	INTRINSIC_OPERATION_I64_BIT_XOR  IntrinsicOperation = 63
	INTRINSIC_OPERATION_I64_BIT_IOR  IntrinsicOperation = 64
	INTRINSIC_OPERATION_I64_CMP_EQ   IntrinsicOperation = 65
	INTRINSIC_OPERATION_I64_CMP_NE   IntrinsicOperation = 66
	INTRINSIC_OPERATION_I64_CMP_LT   IntrinsicOperation = 67
	INTRINSIC_OPERATION_I64_CMP_LE   IntrinsicOperation = 68
	INTRINSIC_OPERATION_I64_CMP_GT   IntrinsicOperation = 69
	INTRINSIC_OPERATION_I64_CMP_GE   IntrinsicOperation = 70
	INTRINSIC_OPERATION_U8_MUL       IntrinsicOperation = 71
	INTRINSIC_OPERATION_U8_DIV       IntrinsicOperation = 72
	INTRINSIC_OPERATION_U8_MOD       IntrinsicOperation = 73
	INTRINSIC_OPERATION_U8_ADD       IntrinsicOperation = 74
	INTRINSIC_OPERATION_U8_SUB       IntrinsicOperation = 75
	INTRINSIC_OPERATION_U8_BIT_LSH   IntrinsicOperation = 76
	INTRINSIC_OPERATION_U8_BIT_RSH   IntrinsicOperation = 77
	INTRINSIC_OPERATION_U8_BIT_AND   IntrinsicOperation = 78
	INTRINSIC_OPERATION_U8_BIT_XOR   IntrinsicOperation = 79
	INTRINSIC_OPERATION_U8_BIT_IOR   IntrinsicOperation = 80
	INTRINSIC_OPERATION_U8_CMP_EQ    IntrinsicOperation = 81
	INTRINSIC_OPERATION_U8_CMP_NE    IntrinsicOperation = 82
	INTRINSIC_OPERATION_U8_CMP_LT    IntrinsicOperation = 83
	INTRINSIC_OPERATION_U8_CMP_LE    IntrinsicOperation = 84
	INTRINSIC_OPERATION_U8_CMP_GT    IntrinsicOperation = 85
	INTRINSIC_OPERATION_U8_CMP_GE    IntrinsicOperation = 86
	INTRINSIC_OPERATION_U16_MUL      IntrinsicOperation = 87
	INTRINSIC_OPERATION_U16_DIV      IntrinsicOperation = 88
	INTRINSIC_OPERATION_U16_MOD      IntrinsicOperation = 89
	INTRINSIC_OPERATION_U16_ADD      IntrinsicOperation = 90
	INTRINSIC_OPERATION_U16_SUB      IntrinsicOperation = 91
	INTRINSIC_OPERATION_U16_BIT_LSH  IntrinsicOperation = 92
	INTRINSIC_OPERATION_U16_BIT_RSH  IntrinsicOperation = 93
	INTRINSIC_OPERATION_U16_BIT_AND  IntrinsicOperation = 94
	INTRINSIC_OPERATION_U16_BIT_XOR  IntrinsicOperation = 95
	INTRINSIC_OPERATION_U16_BIT_IOR  IntrinsicOperation = 96
	INTRINSIC_OPERATION_U16_CMP_EQ   IntrinsicOperation = 97
	INTRINSIC_OPERATION_U16_CMP_NE   IntrinsicOperation = 98
	INTRINSIC_OPERATION_U16_CMP_LT   IntrinsicOperation = 99
	INTRINSIC_OPERATION_U16_CMP_LE   IntrinsicOperation = 100
	INTRINSIC_OPERATION_U16_CMP_GT   IntrinsicOperation = 101
	INTRINSIC_OPERATION_U16_CMP_GE   IntrinsicOperation = 102
	INTRINSIC_OPERATION_U32_MUL      IntrinsicOperation = 103
	INTRINSIC_OPERATION_U32_DIV      IntrinsicOperation = 104
	INTRINSIC_OPERATION_U32_MOD      IntrinsicOperation = 105
	INTRINSIC_OPERATION_U32_ADD      IntrinsicOperation = 106
	INTRINSIC_OPERATION_U32_SUB      IntrinsicOperation = 107
	INTRINSIC_OPERATION_U32_BIT_LSH  IntrinsicOperation = 108
	INTRINSIC_OPERATION_U32_BIT_RSH  IntrinsicOperation = 109
	INTRINSIC_OPERATION_U32_BIT_AND  IntrinsicOperation = 110
	INTRINSIC_OPERATION_U32_BIT_XOR  IntrinsicOperation = 111
	INTRINSIC_OPERATION_U32_BIT_IOR  IntrinsicOperation = 112
	INTRINSIC_OPERATION_U32_CMP_EQ   IntrinsicOperation = 113
	INTRINSIC_OPERATION_U32_CMP_NE   IntrinsicOperation = 114
	INTRINSIC_OPERATION_U32_CMP_LT   IntrinsicOperation = 115
	INTRINSIC_OPERATION_U32_CMP_LE   IntrinsicOperation = 116
	INTRINSIC_OPERATION_U32_CMP_GT   IntrinsicOperation = 117
	INTRINSIC_OPERATION_U32_CMP_GE   IntrinsicOperation = 118
	INTRINSIC_OPERATION_U64_MUL      IntrinsicOperation = 119
	INTRINSIC_OPERATION_U64_DIV      IntrinsicOperation = 120
	INTRINSIC_OPERATION_U64_MOD      IntrinsicOperation = 121
	INTRINSIC_OPERATION_U64_ADD      IntrinsicOperation = 122
	INTRINSIC_OPERATION_U64_SUB      IntrinsicOperation = 123
	INTRINSIC_OPERATION_U64_BIT_LSH  IntrinsicOperation = 124
	INTRINSIC_OPERATION_U64_BIT_RSH  IntrinsicOperation = 125
	INTRINSIC_OPERATION_U64_BIT_AND  IntrinsicOperation = 126
	INTRINSIC_OPERATION_U64_BIT_XOR  IntrinsicOperation = 127
	INTRINSIC_OPERATION_U64_BIT_IOR  IntrinsicOperation = 128
	INTRINSIC_OPERATION_U64_CMP_EQ   IntrinsicOperation = 129
	INTRINSIC_OPERATION_U64_CMP_NE   IntrinsicOperation = 130
	INTRINSIC_OPERATION_U64_CMP_LT   IntrinsicOperation = 131
	INTRINSIC_OPERATION_U64_CMP_LE   IntrinsicOperation = 132
	INTRINSIC_OPERATION_U64_CMP_GT   IntrinsicOperation = 133
	INTRINSIC_OPERATION_U64_CMP_GE   IntrinsicOperation = 134
	INTRINSIC_OPERATION_F32_MUL      IntrinsicOperation = 135
	INTRINSIC_OPERATION_F32_DIV      IntrinsicOperation = 136
	INTRINSIC_OPERATION_F32_ADD      IntrinsicOperation = 137
	INTRINSIC_OPERATION_F32_SUB      IntrinsicOperation = 138
	INTRINSIC_OPERATION_F32_CMP_EQ   IntrinsicOperation = 139
	INTRINSIC_OPERATION_F32_CMP_NE   IntrinsicOperation = 140
	INTRINSIC_OPERATION_F32_CMP_LT   IntrinsicOperation = 141
	INTRINSIC_OPERATION_F32_CMP_LE   IntrinsicOperation = 142
	INTRINSIC_OPERATION_F32_CMP_GT   IntrinsicOperation = 143
	INTRINSIC_OPERATION_F32_CMP_GE   IntrinsicOperation = 144
	INTRINSIC_OPERATION_F64_MUL      IntrinsicOperation = 145
	INTRINSIC_OPERATION_F64_DIV      IntrinsicOperation = 146
	INTRINSIC_OPERATION_F64_ADD      IntrinsicOperation = 147
	INTRINSIC_OPERATION_F64_SUB      IntrinsicOperation = 148
	INTRINSIC_OPERATION_F64_CMP_EQ   IntrinsicOperation = 149
	INTRINSIC_OPERATION_F64_CMP_NE   IntrinsicOperation = 150
	INTRINSIC_OPERATION_F64_CMP_LT   IntrinsicOperation = 151
	INTRINSIC_OPERATION_F64_CMP_LE   IntrinsicOperation = 152
	INTRINSIC_OPERATION_F64_CMP_GT   IntrinsicOperation = 153
	INTRINSIC_OPERATION_F64_CMP_GE   IntrinsicOperation = 154
	NUM_INTRINSIC_OPERATION          IntrinsicOperation = 155
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
		Enum:          INTRINSIC_OPERATION_BOOL_BIT_AND,
		Name:          "INTRINSIC_OPERATION_BOOL_BIT_AND",
		InfixOperator: INFIX_OPERATOR_BIT_AND,
		InfixLeft:     INTRINSIC_TYPE_BOOL,
		InfixRight:    INTRINSIC_TYPE_BOOL,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_BOOL_BIT_XOR,
		Name:          "INTRINSIC_OPERATION_BOOL_BIT_XOR",
		InfixOperator: INFIX_OPERATOR_BIT_XOR,
		InfixLeft:     INTRINSIC_TYPE_BOOL,
		InfixRight:    INTRINSIC_TYPE_BOOL,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_BOOL_BIT_IOR,
		Name:          "INTRINSIC_OPERATION_BOOL_BIT_IOR",
		InfixOperator: INFIX_OPERATOR_BIT_IOR,
		InfixLeft:     INTRINSIC_TYPE_BOOL,
		InfixRight:    INTRINSIC_TYPE_BOOL,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_BOOL_CMP_EQ,
		Name:          "INTRINSIC_OPERATION_BOOL_CMP_EQ",
		InfixOperator: INFIX_OPERATOR_CMP_EQ,
		InfixLeft:     INTRINSIC_TYPE_BOOL,
		InfixRight:    INTRINSIC_TYPE_BOOL,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_BOOL_CMP_NE,
		Name:          "INTRINSIC_OPERATION_BOOL_CMP_NE",
		InfixOperator: INFIX_OPERATOR_CMP_NE,
		InfixLeft:     INTRINSIC_TYPE_BOOL,
		InfixRight:    INTRINSIC_TYPE_BOOL,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_BOOL_AND,
		Name:          "INTRINSIC_OPERATION_BOOL_AND",
		InfixOperator: INFIX_OPERATOR_AND,
		InfixLeft:     INTRINSIC_TYPE_BOOL,
		InfixRight:    INTRINSIC_TYPE_BOOL,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_BOOL_OR,
		Name:          "INTRINSIC_OPERATION_BOOL_OR",
		InfixOperator: INFIX_OPERATOR_OR,
		InfixLeft:     INTRINSIC_TYPE_BOOL,
		InfixRight:    INTRINSIC_TYPE_BOOL,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_I8_MUL,
		Name:          "INTRINSIC_OPERATION_I8_MUL",
		InfixOperator: INFIX_OPERATOR_MUL,
		InfixLeft:     INTRINSIC_TYPE_I8,
		InfixRight:    INTRINSIC_TYPE_I8,
		InfixResult:   INTRINSIC_TYPE_I8,
	},
	{
		Enum:          INTRINSIC_OPERATION_I8_DIV,
		Name:          "INTRINSIC_OPERATION_I8_DIV",
		InfixOperator: INFIX_OPERATOR_DIV,
		InfixLeft:     INTRINSIC_TYPE_I8,
		InfixRight:    INTRINSIC_TYPE_I8,
		InfixResult:   INTRINSIC_TYPE_I8,
	},
	{
		Enum:          INTRINSIC_OPERATION_I8_MOD,
		Name:          "INTRINSIC_OPERATION_I8_MOD",
		InfixOperator: INFIX_OPERATOR_MOD,
		InfixLeft:     INTRINSIC_TYPE_I8,
		InfixRight:    INTRINSIC_TYPE_I8,
		InfixResult:   INTRINSIC_TYPE_I8,
	},
	{
		Enum:          INTRINSIC_OPERATION_I8_ADD,
		Name:          "INTRINSIC_OPERATION_I8_ADD",
		InfixOperator: INFIX_OPERATOR_ADD,
		InfixLeft:     INTRINSIC_TYPE_I8,
		InfixRight:    INTRINSIC_TYPE_I8,
		InfixResult:   INTRINSIC_TYPE_I8,
	},
	{
		Enum:          INTRINSIC_OPERATION_I8_SUB,
		Name:          "INTRINSIC_OPERATION_I8_SUB",
		InfixOperator: INFIX_OPERATOR_SUB,
		InfixLeft:     INTRINSIC_TYPE_I8,
		InfixRight:    INTRINSIC_TYPE_I8,
		InfixResult:   INTRINSIC_TYPE_I8,
	},
	{
		Enum:          INTRINSIC_OPERATION_I8_BIT_LSH,
		Name:          "INTRINSIC_OPERATION_I8_BIT_LSH",
		InfixOperator: INFIX_OPERATOR_BIT_LSH,
		InfixLeft:     INTRINSIC_TYPE_I8,
		InfixRight:    INTRINSIC_TYPE_I8,
		InfixResult:   INTRINSIC_TYPE_I8,
	},
	{
		Enum:          INTRINSIC_OPERATION_I8_BIT_RSH,
		Name:          "INTRINSIC_OPERATION_I8_BIT_RSH",
		InfixOperator: INFIX_OPERATOR_BIT_RSH,
		InfixLeft:     INTRINSIC_TYPE_I8,
		InfixRight:    INTRINSIC_TYPE_I8,
		InfixResult:   INTRINSIC_TYPE_I8,
	},
	{
		Enum:          INTRINSIC_OPERATION_I8_BIT_AND,
		Name:          "INTRINSIC_OPERATION_I8_BIT_AND",
		InfixOperator: INFIX_OPERATOR_BIT_AND,
		InfixLeft:     INTRINSIC_TYPE_I8,
		InfixRight:    INTRINSIC_TYPE_I8,
		InfixResult:   INTRINSIC_TYPE_I8,
	},
	{
		Enum:          INTRINSIC_OPERATION_I8_BIT_XOR,
		Name:          "INTRINSIC_OPERATION_I8_BIT_XOR",
		InfixOperator: INFIX_OPERATOR_BIT_XOR,
		InfixLeft:     INTRINSIC_TYPE_I8,
		InfixRight:    INTRINSIC_TYPE_I8,
		InfixResult:   INTRINSIC_TYPE_I8,
	},
	{
		Enum:          INTRINSIC_OPERATION_I8_BIT_IOR,
		Name:          "INTRINSIC_OPERATION_I8_BIT_IOR",
		InfixOperator: INFIX_OPERATOR_BIT_IOR,
		InfixLeft:     INTRINSIC_TYPE_I8,
		InfixRight:    INTRINSIC_TYPE_I8,
		InfixResult:   INTRINSIC_TYPE_I8,
	},
	{
		Enum:          INTRINSIC_OPERATION_I8_CMP_EQ,
		Name:          "INTRINSIC_OPERATION_I8_CMP_EQ",
		InfixOperator: INFIX_OPERATOR_CMP_EQ,
		InfixLeft:     INTRINSIC_TYPE_I8,
		InfixRight:    INTRINSIC_TYPE_I8,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_I8_CMP_NE,
		Name:          "INTRINSIC_OPERATION_I8_CMP_NE",
		InfixOperator: INFIX_OPERATOR_CMP_NE,
		InfixLeft:     INTRINSIC_TYPE_I8,
		InfixRight:    INTRINSIC_TYPE_I8,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_I8_CMP_LT,
		Name:          "INTRINSIC_OPERATION_I8_CMP_LT",
		InfixOperator: INFIX_OPERATOR_CMP_LT,
		InfixLeft:     INTRINSIC_TYPE_I8,
		InfixRight:    INTRINSIC_TYPE_I8,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_I8_CMP_LE,
		Name:          "INTRINSIC_OPERATION_I8_CMP_LE",
		InfixOperator: INFIX_OPERATOR_CMP_LE,
		InfixLeft:     INTRINSIC_TYPE_I8,
		InfixRight:    INTRINSIC_TYPE_I8,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_I8_CMP_GT,
		Name:          "INTRINSIC_OPERATION_I8_CMP_GT",
		InfixOperator: INFIX_OPERATOR_CMP_GT,
		InfixLeft:     INTRINSIC_TYPE_I8,
		InfixRight:    INTRINSIC_TYPE_I8,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_I8_CMP_GE,
		Name:          "INTRINSIC_OPERATION_I8_CMP_GE",
		InfixOperator: INFIX_OPERATOR_CMP_GE,
		InfixLeft:     INTRINSIC_TYPE_I8,
		InfixRight:    INTRINSIC_TYPE_I8,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_I16_MUL,
		Name:          "INTRINSIC_OPERATION_I16_MUL",
		InfixOperator: INFIX_OPERATOR_MUL,
		InfixLeft:     INTRINSIC_TYPE_I16,
		InfixRight:    INTRINSIC_TYPE_I16,
		InfixResult:   INTRINSIC_TYPE_I16,
	},
	{
		Enum:          INTRINSIC_OPERATION_I16_DIV,
		Name:          "INTRINSIC_OPERATION_I16_DIV",
		InfixOperator: INFIX_OPERATOR_DIV,
		InfixLeft:     INTRINSIC_TYPE_I16,
		InfixRight:    INTRINSIC_TYPE_I16,
		InfixResult:   INTRINSIC_TYPE_I16,
	},
	{
		Enum:          INTRINSIC_OPERATION_I16_MOD,
		Name:          "INTRINSIC_OPERATION_I16_MOD",
		InfixOperator: INFIX_OPERATOR_MOD,
		InfixLeft:     INTRINSIC_TYPE_I16,
		InfixRight:    INTRINSIC_TYPE_I16,
		InfixResult:   INTRINSIC_TYPE_I16,
	},
	{
		Enum:          INTRINSIC_OPERATION_I16_ADD,
		Name:          "INTRINSIC_OPERATION_I16_ADD",
		InfixOperator: INFIX_OPERATOR_ADD,
		InfixLeft:     INTRINSIC_TYPE_I16,
		InfixRight:    INTRINSIC_TYPE_I16,
		InfixResult:   INTRINSIC_TYPE_I16,
	},
	{
		Enum:          INTRINSIC_OPERATION_I16_SUB,
		Name:          "INTRINSIC_OPERATION_I16_SUB",
		InfixOperator: INFIX_OPERATOR_SUB,
		InfixLeft:     INTRINSIC_TYPE_I16,
		InfixRight:    INTRINSIC_TYPE_I16,
		InfixResult:   INTRINSIC_TYPE_I16,
	},
	{
		Enum:          INTRINSIC_OPERATION_I16_BIT_LSH,
		Name:          "INTRINSIC_OPERATION_I16_BIT_LSH",
		InfixOperator: INFIX_OPERATOR_BIT_LSH,
		InfixLeft:     INTRINSIC_TYPE_I16,
		InfixRight:    INTRINSIC_TYPE_I16,
		InfixResult:   INTRINSIC_TYPE_I16,
	},
	{
		Enum:          INTRINSIC_OPERATION_I16_BIT_RSH,
		Name:          "INTRINSIC_OPERATION_I16_BIT_RSH",
		InfixOperator: INFIX_OPERATOR_BIT_RSH,
		InfixLeft:     INTRINSIC_TYPE_I16,
		InfixRight:    INTRINSIC_TYPE_I16,
		InfixResult:   INTRINSIC_TYPE_I16,
	},
	{
		Enum:          INTRINSIC_OPERATION_I16_BIT_AND,
		Name:          "INTRINSIC_OPERATION_I16_BIT_AND",
		InfixOperator: INFIX_OPERATOR_BIT_AND,
		InfixLeft:     INTRINSIC_TYPE_I16,
		InfixRight:    INTRINSIC_TYPE_I16,
		InfixResult:   INTRINSIC_TYPE_I16,
	},
	{
		Enum:          INTRINSIC_OPERATION_I16_BIT_XOR,
		Name:          "INTRINSIC_OPERATION_I16_BIT_XOR",
		InfixOperator: INFIX_OPERATOR_BIT_XOR,
		InfixLeft:     INTRINSIC_TYPE_I16,
		InfixRight:    INTRINSIC_TYPE_I16,
		InfixResult:   INTRINSIC_TYPE_I16,
	},
	{
		Enum:          INTRINSIC_OPERATION_I16_BIT_IOR,
		Name:          "INTRINSIC_OPERATION_I16_BIT_IOR",
		InfixOperator: INFIX_OPERATOR_BIT_IOR,
		InfixLeft:     INTRINSIC_TYPE_I16,
		InfixRight:    INTRINSIC_TYPE_I16,
		InfixResult:   INTRINSIC_TYPE_I16,
	},
	{
		Enum:          INTRINSIC_OPERATION_I16_CMP_EQ,
		Name:          "INTRINSIC_OPERATION_I16_CMP_EQ",
		InfixOperator: INFIX_OPERATOR_CMP_EQ,
		InfixLeft:     INTRINSIC_TYPE_I16,
		InfixRight:    INTRINSIC_TYPE_I16,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_I16_CMP_NE,
		Name:          "INTRINSIC_OPERATION_I16_CMP_NE",
		InfixOperator: INFIX_OPERATOR_CMP_NE,
		InfixLeft:     INTRINSIC_TYPE_I16,
		InfixRight:    INTRINSIC_TYPE_I16,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_I16_CMP_LT,
		Name:          "INTRINSIC_OPERATION_I16_CMP_LT",
		InfixOperator: INFIX_OPERATOR_CMP_LT,
		InfixLeft:     INTRINSIC_TYPE_I16,
		InfixRight:    INTRINSIC_TYPE_I16,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_I16_CMP_LE,
		Name:          "INTRINSIC_OPERATION_I16_CMP_LE",
		InfixOperator: INFIX_OPERATOR_CMP_LE,
		InfixLeft:     INTRINSIC_TYPE_I16,
		InfixRight:    INTRINSIC_TYPE_I16,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_I16_CMP_GT,
		Name:          "INTRINSIC_OPERATION_I16_CMP_GT",
		InfixOperator: INFIX_OPERATOR_CMP_GT,
		InfixLeft:     INTRINSIC_TYPE_I16,
		InfixRight:    INTRINSIC_TYPE_I16,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_I16_CMP_GE,
		Name:          "INTRINSIC_OPERATION_I16_CMP_GE",
		InfixOperator: INFIX_OPERATOR_CMP_GE,
		InfixLeft:     INTRINSIC_TYPE_I16,
		InfixRight:    INTRINSIC_TYPE_I16,
		InfixResult:   INTRINSIC_TYPE_BOOL,
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
		Enum:          INTRINSIC_OPERATION_I32_BIT_LSH,
		Name:          "INTRINSIC_OPERATION_I32_BIT_LSH",
		InfixOperator: INFIX_OPERATOR_BIT_LSH,
		InfixLeft:     INTRINSIC_TYPE_I32,
		InfixRight:    INTRINSIC_TYPE_I32,
		InfixResult:   INTRINSIC_TYPE_I32,
	},
	{
		Enum:          INTRINSIC_OPERATION_I32_BIT_RSH,
		Name:          "INTRINSIC_OPERATION_I32_BIT_RSH",
		InfixOperator: INFIX_OPERATOR_BIT_RSH,
		InfixLeft:     INTRINSIC_TYPE_I32,
		InfixRight:    INTRINSIC_TYPE_I32,
		InfixResult:   INTRINSIC_TYPE_I32,
	},
	{
		Enum:          INTRINSIC_OPERATION_I32_BIT_AND,
		Name:          "INTRINSIC_OPERATION_I32_BIT_AND",
		InfixOperator: INFIX_OPERATOR_BIT_AND,
		InfixLeft:     INTRINSIC_TYPE_I32,
		InfixRight:    INTRINSIC_TYPE_I32,
		InfixResult:   INTRINSIC_TYPE_I32,
	},
	{
		Enum:          INTRINSIC_OPERATION_I32_BIT_XOR,
		Name:          "INTRINSIC_OPERATION_I32_BIT_XOR",
		InfixOperator: INFIX_OPERATOR_BIT_XOR,
		InfixLeft:     INTRINSIC_TYPE_I32,
		InfixRight:    INTRINSIC_TYPE_I32,
		InfixResult:   INTRINSIC_TYPE_I32,
	},
	{
		Enum:          INTRINSIC_OPERATION_I32_BIT_IOR,
		Name:          "INTRINSIC_OPERATION_I32_BIT_IOR",
		InfixOperator: INFIX_OPERATOR_BIT_IOR,
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
		Enum:          INTRINSIC_OPERATION_I64_MUL,
		Name:          "INTRINSIC_OPERATION_I64_MUL",
		InfixOperator: INFIX_OPERATOR_MUL,
		InfixLeft:     INTRINSIC_TYPE_I64,
		InfixRight:    INTRINSIC_TYPE_I64,
		InfixResult:   INTRINSIC_TYPE_I64,
	},
	{
		Enum:          INTRINSIC_OPERATION_I64_DIV,
		Name:          "INTRINSIC_OPERATION_I64_DIV",
		InfixOperator: INFIX_OPERATOR_DIV,
		InfixLeft:     INTRINSIC_TYPE_I64,
		InfixRight:    INTRINSIC_TYPE_I64,
		InfixResult:   INTRINSIC_TYPE_I64,
	},
	{
		Enum:          INTRINSIC_OPERATION_I64_MOD,
		Name:          "INTRINSIC_OPERATION_I64_MOD",
		InfixOperator: INFIX_OPERATOR_MOD,
		InfixLeft:     INTRINSIC_TYPE_I64,
		InfixRight:    INTRINSIC_TYPE_I64,
		InfixResult:   INTRINSIC_TYPE_I64,
	},
	{
		Enum:          INTRINSIC_OPERATION_I64_ADD,
		Name:          "INTRINSIC_OPERATION_I64_ADD",
		InfixOperator: INFIX_OPERATOR_ADD,
		InfixLeft:     INTRINSIC_TYPE_I64,
		InfixRight:    INTRINSIC_TYPE_I64,
		InfixResult:   INTRINSIC_TYPE_I64,
	},
	{
		Enum:          INTRINSIC_OPERATION_I64_SUB,
		Name:          "INTRINSIC_OPERATION_I64_SUB",
		InfixOperator: INFIX_OPERATOR_SUB,
		InfixLeft:     INTRINSIC_TYPE_I64,
		InfixRight:    INTRINSIC_TYPE_I64,
		InfixResult:   INTRINSIC_TYPE_I64,
	},
	{
		Enum:          INTRINSIC_OPERATION_I64_BIT_LSH,
		Name:          "INTRINSIC_OPERATION_I64_BIT_LSH",
		InfixOperator: INFIX_OPERATOR_BIT_LSH,
		InfixLeft:     INTRINSIC_TYPE_I64,
		InfixRight:    INTRINSIC_TYPE_I64,
		InfixResult:   INTRINSIC_TYPE_I64,
	},
	{
		Enum:          INTRINSIC_OPERATION_I64_BIT_RSH,
		Name:          "INTRINSIC_OPERATION_I64_BIT_RSH",
		InfixOperator: INFIX_OPERATOR_BIT_RSH,
		InfixLeft:     INTRINSIC_TYPE_I64,
		InfixRight:    INTRINSIC_TYPE_I64,
		InfixResult:   INTRINSIC_TYPE_I64,
	},
	{
		Enum:          INTRINSIC_OPERATION_I64_BIT_AND,
		Name:          "INTRINSIC_OPERATION_I64_BIT_AND",
		InfixOperator: INFIX_OPERATOR_BIT_AND,
		InfixLeft:     INTRINSIC_TYPE_I64,
		InfixRight:    INTRINSIC_TYPE_I64,
		InfixResult:   INTRINSIC_TYPE_I64,
	},
	{
		Enum:          INTRINSIC_OPERATION_I64_BIT_XOR,
		Name:          "INTRINSIC_OPERATION_I64_BIT_XOR",
		InfixOperator: INFIX_OPERATOR_BIT_XOR,
		InfixLeft:     INTRINSIC_TYPE_I64,
		InfixRight:    INTRINSIC_TYPE_I64,
		InfixResult:   INTRINSIC_TYPE_I64,
	},
	{
		Enum:          INTRINSIC_OPERATION_I64_BIT_IOR,
		Name:          "INTRINSIC_OPERATION_I64_BIT_IOR",
		InfixOperator: INFIX_OPERATOR_BIT_IOR,
		InfixLeft:     INTRINSIC_TYPE_I64,
		InfixRight:    INTRINSIC_TYPE_I64,
		InfixResult:   INTRINSIC_TYPE_I64,
	},
	{
		Enum:          INTRINSIC_OPERATION_I64_CMP_EQ,
		Name:          "INTRINSIC_OPERATION_I64_CMP_EQ",
		InfixOperator: INFIX_OPERATOR_CMP_EQ,
		InfixLeft:     INTRINSIC_TYPE_I64,
		InfixRight:    INTRINSIC_TYPE_I64,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_I64_CMP_NE,
		Name:          "INTRINSIC_OPERATION_I64_CMP_NE",
		InfixOperator: INFIX_OPERATOR_CMP_NE,
		InfixLeft:     INTRINSIC_TYPE_I64,
		InfixRight:    INTRINSIC_TYPE_I64,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_I64_CMP_LT,
		Name:          "INTRINSIC_OPERATION_I64_CMP_LT",
		InfixOperator: INFIX_OPERATOR_CMP_LT,
		InfixLeft:     INTRINSIC_TYPE_I64,
		InfixRight:    INTRINSIC_TYPE_I64,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_I64_CMP_LE,
		Name:          "INTRINSIC_OPERATION_I64_CMP_LE",
		InfixOperator: INFIX_OPERATOR_CMP_LE,
		InfixLeft:     INTRINSIC_TYPE_I64,
		InfixRight:    INTRINSIC_TYPE_I64,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_I64_CMP_GT,
		Name:          "INTRINSIC_OPERATION_I64_CMP_GT",
		InfixOperator: INFIX_OPERATOR_CMP_GT,
		InfixLeft:     INTRINSIC_TYPE_I64,
		InfixRight:    INTRINSIC_TYPE_I64,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_I64_CMP_GE,
		Name:          "INTRINSIC_OPERATION_I64_CMP_GE",
		InfixOperator: INFIX_OPERATOR_CMP_GE,
		InfixLeft:     INTRINSIC_TYPE_I64,
		InfixRight:    INTRINSIC_TYPE_I64,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_U8_MUL,
		Name:          "INTRINSIC_OPERATION_U8_MUL",
		InfixOperator: INFIX_OPERATOR_MUL,
		InfixLeft:     INTRINSIC_TYPE_U8,
		InfixRight:    INTRINSIC_TYPE_U8,
		InfixResult:   INTRINSIC_TYPE_U8,
	},
	{
		Enum:          INTRINSIC_OPERATION_U8_DIV,
		Name:          "INTRINSIC_OPERATION_U8_DIV",
		InfixOperator: INFIX_OPERATOR_DIV,
		InfixLeft:     INTRINSIC_TYPE_U8,
		InfixRight:    INTRINSIC_TYPE_U8,
		InfixResult:   INTRINSIC_TYPE_U8,
	},
	{
		Enum:          INTRINSIC_OPERATION_U8_MOD,
		Name:          "INTRINSIC_OPERATION_U8_MOD",
		InfixOperator: INFIX_OPERATOR_MOD,
		InfixLeft:     INTRINSIC_TYPE_U8,
		InfixRight:    INTRINSIC_TYPE_U8,
		InfixResult:   INTRINSIC_TYPE_U8,
	},
	{
		Enum:          INTRINSIC_OPERATION_U8_ADD,
		Name:          "INTRINSIC_OPERATION_U8_ADD",
		InfixOperator: INFIX_OPERATOR_ADD,
		InfixLeft:     INTRINSIC_TYPE_U8,
		InfixRight:    INTRINSIC_TYPE_U8,
		InfixResult:   INTRINSIC_TYPE_U8,
	},
	{
		Enum:          INTRINSIC_OPERATION_U8_SUB,
		Name:          "INTRINSIC_OPERATION_U8_SUB",
		InfixOperator: INFIX_OPERATOR_SUB,
		InfixLeft:     INTRINSIC_TYPE_U8,
		InfixRight:    INTRINSIC_TYPE_U8,
		InfixResult:   INTRINSIC_TYPE_U8,
	},
	{
		Enum:          INTRINSIC_OPERATION_U8_BIT_LSH,
		Name:          "INTRINSIC_OPERATION_U8_BIT_LSH",
		InfixOperator: INFIX_OPERATOR_BIT_LSH,
		InfixLeft:     INTRINSIC_TYPE_U8,
		InfixRight:    INTRINSIC_TYPE_U8,
		InfixResult:   INTRINSIC_TYPE_U8,
	},
	{
		Enum:          INTRINSIC_OPERATION_U8_BIT_RSH,
		Name:          "INTRINSIC_OPERATION_U8_BIT_RSH",
		InfixOperator: INFIX_OPERATOR_BIT_RSH,
		InfixLeft:     INTRINSIC_TYPE_U8,
		InfixRight:    INTRINSIC_TYPE_U8,
		InfixResult:   INTRINSIC_TYPE_U8,
	},
	{
		Enum:          INTRINSIC_OPERATION_U8_BIT_AND,
		Name:          "INTRINSIC_OPERATION_U8_BIT_AND",
		InfixOperator: INFIX_OPERATOR_BIT_AND,
		InfixLeft:     INTRINSIC_TYPE_U8,
		InfixRight:    INTRINSIC_TYPE_U8,
		InfixResult:   INTRINSIC_TYPE_U8,
	},
	{
		Enum:          INTRINSIC_OPERATION_U8_BIT_XOR,
		Name:          "INTRINSIC_OPERATION_U8_BIT_XOR",
		InfixOperator: INFIX_OPERATOR_BIT_XOR,
		InfixLeft:     INTRINSIC_TYPE_U8,
		InfixRight:    INTRINSIC_TYPE_U8,
		InfixResult:   INTRINSIC_TYPE_U8,
	},
	{
		Enum:          INTRINSIC_OPERATION_U8_BIT_IOR,
		Name:          "INTRINSIC_OPERATION_U8_BIT_IOR",
		InfixOperator: INFIX_OPERATOR_BIT_IOR,
		InfixLeft:     INTRINSIC_TYPE_U8,
		InfixRight:    INTRINSIC_TYPE_U8,
		InfixResult:   INTRINSIC_TYPE_U8,
	},
	{
		Enum:          INTRINSIC_OPERATION_U8_CMP_EQ,
		Name:          "INTRINSIC_OPERATION_U8_CMP_EQ",
		InfixOperator: INFIX_OPERATOR_CMP_EQ,
		InfixLeft:     INTRINSIC_TYPE_U8,
		InfixRight:    INTRINSIC_TYPE_U8,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_U8_CMP_NE,
		Name:          "INTRINSIC_OPERATION_U8_CMP_NE",
		InfixOperator: INFIX_OPERATOR_CMP_NE,
		InfixLeft:     INTRINSIC_TYPE_U8,
		InfixRight:    INTRINSIC_TYPE_U8,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_U8_CMP_LT,
		Name:          "INTRINSIC_OPERATION_U8_CMP_LT",
		InfixOperator: INFIX_OPERATOR_CMP_LT,
		InfixLeft:     INTRINSIC_TYPE_U8,
		InfixRight:    INTRINSIC_TYPE_U8,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_U8_CMP_LE,
		Name:          "INTRINSIC_OPERATION_U8_CMP_LE",
		InfixOperator: INFIX_OPERATOR_CMP_LE,
		InfixLeft:     INTRINSIC_TYPE_U8,
		InfixRight:    INTRINSIC_TYPE_U8,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_U8_CMP_GT,
		Name:          "INTRINSIC_OPERATION_U8_CMP_GT",
		InfixOperator: INFIX_OPERATOR_CMP_GT,
		InfixLeft:     INTRINSIC_TYPE_U8,
		InfixRight:    INTRINSIC_TYPE_U8,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_U8_CMP_GE,
		Name:          "INTRINSIC_OPERATION_U8_CMP_GE",
		InfixOperator: INFIX_OPERATOR_CMP_GE,
		InfixLeft:     INTRINSIC_TYPE_U8,
		InfixRight:    INTRINSIC_TYPE_U8,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_U16_MUL,
		Name:          "INTRINSIC_OPERATION_U16_MUL",
		InfixOperator: INFIX_OPERATOR_MUL,
		InfixLeft:     INTRINSIC_TYPE_U16,
		InfixRight:    INTRINSIC_TYPE_U16,
		InfixResult:   INTRINSIC_TYPE_U16,
	},
	{
		Enum:          INTRINSIC_OPERATION_U16_DIV,
		Name:          "INTRINSIC_OPERATION_U16_DIV",
		InfixOperator: INFIX_OPERATOR_DIV,
		InfixLeft:     INTRINSIC_TYPE_U16,
		InfixRight:    INTRINSIC_TYPE_U16,
		InfixResult:   INTRINSIC_TYPE_U16,
	},
	{
		Enum:          INTRINSIC_OPERATION_U16_MOD,
		Name:          "INTRINSIC_OPERATION_U16_MOD",
		InfixOperator: INFIX_OPERATOR_MOD,
		InfixLeft:     INTRINSIC_TYPE_U16,
		InfixRight:    INTRINSIC_TYPE_U16,
		InfixResult:   INTRINSIC_TYPE_U16,
	},
	{
		Enum:          INTRINSIC_OPERATION_U16_ADD,
		Name:          "INTRINSIC_OPERATION_U16_ADD",
		InfixOperator: INFIX_OPERATOR_ADD,
		InfixLeft:     INTRINSIC_TYPE_U16,
		InfixRight:    INTRINSIC_TYPE_U16,
		InfixResult:   INTRINSIC_TYPE_U16,
	},
	{
		Enum:          INTRINSIC_OPERATION_U16_SUB,
		Name:          "INTRINSIC_OPERATION_U16_SUB",
		InfixOperator: INFIX_OPERATOR_SUB,
		InfixLeft:     INTRINSIC_TYPE_U16,
		InfixRight:    INTRINSIC_TYPE_U16,
		InfixResult:   INTRINSIC_TYPE_U16,
	},
	{
		Enum:          INTRINSIC_OPERATION_U16_BIT_LSH,
		Name:          "INTRINSIC_OPERATION_U16_BIT_LSH",
		InfixOperator: INFIX_OPERATOR_BIT_LSH,
		InfixLeft:     INTRINSIC_TYPE_U16,
		InfixRight:    INTRINSIC_TYPE_U16,
		InfixResult:   INTRINSIC_TYPE_U16,
	},
	{
		Enum:          INTRINSIC_OPERATION_U16_BIT_RSH,
		Name:          "INTRINSIC_OPERATION_U16_BIT_RSH",
		InfixOperator: INFIX_OPERATOR_BIT_RSH,
		InfixLeft:     INTRINSIC_TYPE_U16,
		InfixRight:    INTRINSIC_TYPE_U16,
		InfixResult:   INTRINSIC_TYPE_U16,
	},
	{
		Enum:          INTRINSIC_OPERATION_U16_BIT_AND,
		Name:          "INTRINSIC_OPERATION_U16_BIT_AND",
		InfixOperator: INFIX_OPERATOR_BIT_AND,
		InfixLeft:     INTRINSIC_TYPE_U16,
		InfixRight:    INTRINSIC_TYPE_U16,
		InfixResult:   INTRINSIC_TYPE_U16,
	},
	{
		Enum:          INTRINSIC_OPERATION_U16_BIT_XOR,
		Name:          "INTRINSIC_OPERATION_U16_BIT_XOR",
		InfixOperator: INFIX_OPERATOR_BIT_XOR,
		InfixLeft:     INTRINSIC_TYPE_U16,
		InfixRight:    INTRINSIC_TYPE_U16,
		InfixResult:   INTRINSIC_TYPE_U16,
	},
	{
		Enum:          INTRINSIC_OPERATION_U16_BIT_IOR,
		Name:          "INTRINSIC_OPERATION_U16_BIT_IOR",
		InfixOperator: INFIX_OPERATOR_BIT_IOR,
		InfixLeft:     INTRINSIC_TYPE_U16,
		InfixRight:    INTRINSIC_TYPE_U16,
		InfixResult:   INTRINSIC_TYPE_U16,
	},
	{
		Enum:          INTRINSIC_OPERATION_U16_CMP_EQ,
		Name:          "INTRINSIC_OPERATION_U16_CMP_EQ",
		InfixOperator: INFIX_OPERATOR_CMP_EQ,
		InfixLeft:     INTRINSIC_TYPE_U16,
		InfixRight:    INTRINSIC_TYPE_U16,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_U16_CMP_NE,
		Name:          "INTRINSIC_OPERATION_U16_CMP_NE",
		InfixOperator: INFIX_OPERATOR_CMP_NE,
		InfixLeft:     INTRINSIC_TYPE_U16,
		InfixRight:    INTRINSIC_TYPE_U16,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_U16_CMP_LT,
		Name:          "INTRINSIC_OPERATION_U16_CMP_LT",
		InfixOperator: INFIX_OPERATOR_CMP_LT,
		InfixLeft:     INTRINSIC_TYPE_U16,
		InfixRight:    INTRINSIC_TYPE_U16,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_U16_CMP_LE,
		Name:          "INTRINSIC_OPERATION_U16_CMP_LE",
		InfixOperator: INFIX_OPERATOR_CMP_LE,
		InfixLeft:     INTRINSIC_TYPE_U16,
		InfixRight:    INTRINSIC_TYPE_U16,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_U16_CMP_GT,
		Name:          "INTRINSIC_OPERATION_U16_CMP_GT",
		InfixOperator: INFIX_OPERATOR_CMP_GT,
		InfixLeft:     INTRINSIC_TYPE_U16,
		InfixRight:    INTRINSIC_TYPE_U16,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_U16_CMP_GE,
		Name:          "INTRINSIC_OPERATION_U16_CMP_GE",
		InfixOperator: INFIX_OPERATOR_CMP_GE,
		InfixLeft:     INTRINSIC_TYPE_U16,
		InfixRight:    INTRINSIC_TYPE_U16,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_U32_MUL,
		Name:          "INTRINSIC_OPERATION_U32_MUL",
		InfixOperator: INFIX_OPERATOR_MUL,
		InfixLeft:     INTRINSIC_TYPE_U32,
		InfixRight:    INTRINSIC_TYPE_U32,
		InfixResult:   INTRINSIC_TYPE_U32,
	},
	{
		Enum:          INTRINSIC_OPERATION_U32_DIV,
		Name:          "INTRINSIC_OPERATION_U32_DIV",
		InfixOperator: INFIX_OPERATOR_DIV,
		InfixLeft:     INTRINSIC_TYPE_U32,
		InfixRight:    INTRINSIC_TYPE_U32,
		InfixResult:   INTRINSIC_TYPE_U32,
	},
	{
		Enum:          INTRINSIC_OPERATION_U32_MOD,
		Name:          "INTRINSIC_OPERATION_U32_MOD",
		InfixOperator: INFIX_OPERATOR_MOD,
		InfixLeft:     INTRINSIC_TYPE_U32,
		InfixRight:    INTRINSIC_TYPE_U32,
		InfixResult:   INTRINSIC_TYPE_U32,
	},
	{
		Enum:          INTRINSIC_OPERATION_U32_ADD,
		Name:          "INTRINSIC_OPERATION_U32_ADD",
		InfixOperator: INFIX_OPERATOR_ADD,
		InfixLeft:     INTRINSIC_TYPE_U32,
		InfixRight:    INTRINSIC_TYPE_U32,
		InfixResult:   INTRINSIC_TYPE_U32,
	},
	{
		Enum:          INTRINSIC_OPERATION_U32_SUB,
		Name:          "INTRINSIC_OPERATION_U32_SUB",
		InfixOperator: INFIX_OPERATOR_SUB,
		InfixLeft:     INTRINSIC_TYPE_U32,
		InfixRight:    INTRINSIC_TYPE_U32,
		InfixResult:   INTRINSIC_TYPE_U32,
	},
	{
		Enum:          INTRINSIC_OPERATION_U32_BIT_LSH,
		Name:          "INTRINSIC_OPERATION_U32_BIT_LSH",
		InfixOperator: INFIX_OPERATOR_BIT_LSH,
		InfixLeft:     INTRINSIC_TYPE_U32,
		InfixRight:    INTRINSIC_TYPE_U32,
		InfixResult:   INTRINSIC_TYPE_U32,
	},
	{
		Enum:          INTRINSIC_OPERATION_U32_BIT_RSH,
		Name:          "INTRINSIC_OPERATION_U32_BIT_RSH",
		InfixOperator: INFIX_OPERATOR_BIT_RSH,
		InfixLeft:     INTRINSIC_TYPE_U32,
		InfixRight:    INTRINSIC_TYPE_U32,
		InfixResult:   INTRINSIC_TYPE_U32,
	},
	{
		Enum:          INTRINSIC_OPERATION_U32_BIT_AND,
		Name:          "INTRINSIC_OPERATION_U32_BIT_AND",
		InfixOperator: INFIX_OPERATOR_BIT_AND,
		InfixLeft:     INTRINSIC_TYPE_U32,
		InfixRight:    INTRINSIC_TYPE_U32,
		InfixResult:   INTRINSIC_TYPE_U32,
	},
	{
		Enum:          INTRINSIC_OPERATION_U32_BIT_XOR,
		Name:          "INTRINSIC_OPERATION_U32_BIT_XOR",
		InfixOperator: INFIX_OPERATOR_BIT_XOR,
		InfixLeft:     INTRINSIC_TYPE_U32,
		InfixRight:    INTRINSIC_TYPE_U32,
		InfixResult:   INTRINSIC_TYPE_U32,
	},
	{
		Enum:          INTRINSIC_OPERATION_U32_BIT_IOR,
		Name:          "INTRINSIC_OPERATION_U32_BIT_IOR",
		InfixOperator: INFIX_OPERATOR_BIT_IOR,
		InfixLeft:     INTRINSIC_TYPE_U32,
		InfixRight:    INTRINSIC_TYPE_U32,
		InfixResult:   INTRINSIC_TYPE_U32,
	},
	{
		Enum:          INTRINSIC_OPERATION_U32_CMP_EQ,
		Name:          "INTRINSIC_OPERATION_U32_CMP_EQ",
		InfixOperator: INFIX_OPERATOR_CMP_EQ,
		InfixLeft:     INTRINSIC_TYPE_U32,
		InfixRight:    INTRINSIC_TYPE_U32,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_U32_CMP_NE,
		Name:          "INTRINSIC_OPERATION_U32_CMP_NE",
		InfixOperator: INFIX_OPERATOR_CMP_NE,
		InfixLeft:     INTRINSIC_TYPE_U32,
		InfixRight:    INTRINSIC_TYPE_U32,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_U32_CMP_LT,
		Name:          "INTRINSIC_OPERATION_U32_CMP_LT",
		InfixOperator: INFIX_OPERATOR_CMP_LT,
		InfixLeft:     INTRINSIC_TYPE_U32,
		InfixRight:    INTRINSIC_TYPE_U32,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_U32_CMP_LE,
		Name:          "INTRINSIC_OPERATION_U32_CMP_LE",
		InfixOperator: INFIX_OPERATOR_CMP_LE,
		InfixLeft:     INTRINSIC_TYPE_U32,
		InfixRight:    INTRINSIC_TYPE_U32,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_U32_CMP_GT,
		Name:          "INTRINSIC_OPERATION_U32_CMP_GT",
		InfixOperator: INFIX_OPERATOR_CMP_GT,
		InfixLeft:     INTRINSIC_TYPE_U32,
		InfixRight:    INTRINSIC_TYPE_U32,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_U32_CMP_GE,
		Name:          "INTRINSIC_OPERATION_U32_CMP_GE",
		InfixOperator: INFIX_OPERATOR_CMP_GE,
		InfixLeft:     INTRINSIC_TYPE_U32,
		InfixRight:    INTRINSIC_TYPE_U32,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_U64_MUL,
		Name:          "INTRINSIC_OPERATION_U64_MUL",
		InfixOperator: INFIX_OPERATOR_MUL,
		InfixLeft:     INTRINSIC_TYPE_U64,
		InfixRight:    INTRINSIC_TYPE_U64,
		InfixResult:   INTRINSIC_TYPE_U64,
	},
	{
		Enum:          INTRINSIC_OPERATION_U64_DIV,
		Name:          "INTRINSIC_OPERATION_U64_DIV",
		InfixOperator: INFIX_OPERATOR_DIV,
		InfixLeft:     INTRINSIC_TYPE_U64,
		InfixRight:    INTRINSIC_TYPE_U64,
		InfixResult:   INTRINSIC_TYPE_U64,
	},
	{
		Enum:          INTRINSIC_OPERATION_U64_MOD,
		Name:          "INTRINSIC_OPERATION_U64_MOD",
		InfixOperator: INFIX_OPERATOR_MOD,
		InfixLeft:     INTRINSIC_TYPE_U64,
		InfixRight:    INTRINSIC_TYPE_U64,
		InfixResult:   INTRINSIC_TYPE_U64,
	},
	{
		Enum:          INTRINSIC_OPERATION_U64_ADD,
		Name:          "INTRINSIC_OPERATION_U64_ADD",
		InfixOperator: INFIX_OPERATOR_ADD,
		InfixLeft:     INTRINSIC_TYPE_U64,
		InfixRight:    INTRINSIC_TYPE_U64,
		InfixResult:   INTRINSIC_TYPE_U64,
	},
	{
		Enum:          INTRINSIC_OPERATION_U64_SUB,
		Name:          "INTRINSIC_OPERATION_U64_SUB",
		InfixOperator: INFIX_OPERATOR_SUB,
		InfixLeft:     INTRINSIC_TYPE_U64,
		InfixRight:    INTRINSIC_TYPE_U64,
		InfixResult:   INTRINSIC_TYPE_U64,
	},
	{
		Enum:          INTRINSIC_OPERATION_U64_BIT_LSH,
		Name:          "INTRINSIC_OPERATION_U64_BIT_LSH",
		InfixOperator: INFIX_OPERATOR_BIT_LSH,
		InfixLeft:     INTRINSIC_TYPE_U64,
		InfixRight:    INTRINSIC_TYPE_U64,
		InfixResult:   INTRINSIC_TYPE_U64,
	},
	{
		Enum:          INTRINSIC_OPERATION_U64_BIT_RSH,
		Name:          "INTRINSIC_OPERATION_U64_BIT_RSH",
		InfixOperator: INFIX_OPERATOR_BIT_RSH,
		InfixLeft:     INTRINSIC_TYPE_U64,
		InfixRight:    INTRINSIC_TYPE_U64,
		InfixResult:   INTRINSIC_TYPE_U64,
	},
	{
		Enum:          INTRINSIC_OPERATION_U64_BIT_AND,
		Name:          "INTRINSIC_OPERATION_U64_BIT_AND",
		InfixOperator: INFIX_OPERATOR_BIT_AND,
		InfixLeft:     INTRINSIC_TYPE_U64,
		InfixRight:    INTRINSIC_TYPE_U64,
		InfixResult:   INTRINSIC_TYPE_U64,
	},
	{
		Enum:          INTRINSIC_OPERATION_U64_BIT_XOR,
		Name:          "INTRINSIC_OPERATION_U64_BIT_XOR",
		InfixOperator: INFIX_OPERATOR_BIT_XOR,
		InfixLeft:     INTRINSIC_TYPE_U64,
		InfixRight:    INTRINSIC_TYPE_U64,
		InfixResult:   INTRINSIC_TYPE_U64,
	},
	{
		Enum:          INTRINSIC_OPERATION_U64_BIT_IOR,
		Name:          "INTRINSIC_OPERATION_U64_BIT_IOR",
		InfixOperator: INFIX_OPERATOR_BIT_IOR,
		InfixLeft:     INTRINSIC_TYPE_U64,
		InfixRight:    INTRINSIC_TYPE_U64,
		InfixResult:   INTRINSIC_TYPE_U64,
	},
	{
		Enum:          INTRINSIC_OPERATION_U64_CMP_EQ,
		Name:          "INTRINSIC_OPERATION_U64_CMP_EQ",
		InfixOperator: INFIX_OPERATOR_CMP_EQ,
		InfixLeft:     INTRINSIC_TYPE_U64,
		InfixRight:    INTRINSIC_TYPE_U64,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_U64_CMP_NE,
		Name:          "INTRINSIC_OPERATION_U64_CMP_NE",
		InfixOperator: INFIX_OPERATOR_CMP_NE,
		InfixLeft:     INTRINSIC_TYPE_U64,
		InfixRight:    INTRINSIC_TYPE_U64,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_U64_CMP_LT,
		Name:          "INTRINSIC_OPERATION_U64_CMP_LT",
		InfixOperator: INFIX_OPERATOR_CMP_LT,
		InfixLeft:     INTRINSIC_TYPE_U64,
		InfixRight:    INTRINSIC_TYPE_U64,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_U64_CMP_LE,
		Name:          "INTRINSIC_OPERATION_U64_CMP_LE",
		InfixOperator: INFIX_OPERATOR_CMP_LE,
		InfixLeft:     INTRINSIC_TYPE_U64,
		InfixRight:    INTRINSIC_TYPE_U64,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_U64_CMP_GT,
		Name:          "INTRINSIC_OPERATION_U64_CMP_GT",
		InfixOperator: INFIX_OPERATOR_CMP_GT,
		InfixLeft:     INTRINSIC_TYPE_U64,
		InfixRight:    INTRINSIC_TYPE_U64,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_U64_CMP_GE,
		Name:          "INTRINSIC_OPERATION_U64_CMP_GE",
		InfixOperator: INFIX_OPERATOR_CMP_GE,
		InfixLeft:     INTRINSIC_TYPE_U64,
		InfixRight:    INTRINSIC_TYPE_U64,
		InfixResult:   INTRINSIC_TYPE_BOOL,
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
	{
		Enum:          INTRINSIC_OPERATION_F64_MUL,
		Name:          "INTRINSIC_OPERATION_F64_MUL",
		InfixOperator: INFIX_OPERATOR_MUL,
		InfixLeft:     INTRINSIC_TYPE_F64,
		InfixRight:    INTRINSIC_TYPE_F64,
		InfixResult:   INTRINSIC_TYPE_F64,
	},
	{
		Enum:          INTRINSIC_OPERATION_F64_DIV,
		Name:          "INTRINSIC_OPERATION_F64_DIV",
		InfixOperator: INFIX_OPERATOR_DIV,
		InfixLeft:     INTRINSIC_TYPE_F64,
		InfixRight:    INTRINSIC_TYPE_F64,
		InfixResult:   INTRINSIC_TYPE_F64,
	},
	{
		Enum:          INTRINSIC_OPERATION_F64_ADD,
		Name:          "INTRINSIC_OPERATION_F64_ADD",
		InfixOperator: INFIX_OPERATOR_ADD,
		InfixLeft:     INTRINSIC_TYPE_F64,
		InfixRight:    INTRINSIC_TYPE_F64,
		InfixResult:   INTRINSIC_TYPE_F64,
	},
	{
		Enum:          INTRINSIC_OPERATION_F64_SUB,
		Name:          "INTRINSIC_OPERATION_F64_SUB",
		InfixOperator: INFIX_OPERATOR_SUB,
		InfixLeft:     INTRINSIC_TYPE_F64,
		InfixRight:    INTRINSIC_TYPE_F64,
		InfixResult:   INTRINSIC_TYPE_F64,
	},
	{
		Enum:          INTRINSIC_OPERATION_F64_CMP_EQ,
		Name:          "INTRINSIC_OPERATION_F64_CMP_EQ",
		InfixOperator: INFIX_OPERATOR_CMP_EQ,
		InfixLeft:     INTRINSIC_TYPE_F64,
		InfixRight:    INTRINSIC_TYPE_F64,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_F64_CMP_NE,
		Name:          "INTRINSIC_OPERATION_F64_CMP_NE",
		InfixOperator: INFIX_OPERATOR_CMP_NE,
		InfixLeft:     INTRINSIC_TYPE_F64,
		InfixRight:    INTRINSIC_TYPE_F64,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_F64_CMP_LT,
		Name:          "INTRINSIC_OPERATION_F64_CMP_LT",
		InfixOperator: INFIX_OPERATOR_CMP_LT,
		InfixLeft:     INTRINSIC_TYPE_F64,
		InfixRight:    INTRINSIC_TYPE_F64,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_F64_CMP_LE,
		Name:          "INTRINSIC_OPERATION_F64_CMP_LE",
		InfixOperator: INFIX_OPERATOR_CMP_LE,
		InfixLeft:     INTRINSIC_TYPE_F64,
		InfixRight:    INTRINSIC_TYPE_F64,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_F64_CMP_GT,
		Name:          "INTRINSIC_OPERATION_F64_CMP_GT",
		InfixOperator: INFIX_OPERATOR_CMP_GT,
		InfixLeft:     INTRINSIC_TYPE_F64,
		InfixRight:    INTRINSIC_TYPE_F64,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
	{
		Enum:          INTRINSIC_OPERATION_F64_CMP_GE,
		Name:          "INTRINSIC_OPERATION_F64_CMP_GE",
		InfixOperator: INFIX_OPERATOR_CMP_GE,
		InfixLeft:     INTRINSIC_TYPE_F64,
		InfixRight:    INTRINSIC_TYPE_F64,
		InfixResult:   INTRINSIC_TYPE_BOOL,
	},
}

func (value IntrinsicOperation) String() string {
	if value < 0 || value >= 155 {
		return "INVALID_INTRINSIC_OPERATION"
	}
	return IntrinsicOperationInfos[value].Name
}

var InfixToIntrinsic = map[InfixOperator]map[IntrinsicType]map[IntrinsicType]*IntrinsicOperationInfo{
	INFIX_OPERATOR_BIT_AND: {
		INTRINSIC_TYPE_BOOL: {
			INTRINSIC_TYPE_BOOL: IntrinsicOperationInfos[INTRINSIC_OPERATION_BOOL_BIT_AND],
		},
		INTRINSIC_TYPE_I8: {
			INTRINSIC_TYPE_I8: IntrinsicOperationInfos[INTRINSIC_OPERATION_I8_BIT_AND],
		},
		INTRINSIC_TYPE_I16: {
			INTRINSIC_TYPE_I16: IntrinsicOperationInfos[INTRINSIC_OPERATION_I16_BIT_AND],
		},
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_BIT_AND],
		},
		INTRINSIC_TYPE_I64: {
			INTRINSIC_TYPE_I64: IntrinsicOperationInfos[INTRINSIC_OPERATION_I64_BIT_AND],
		},
		INTRINSIC_TYPE_U8: {
			INTRINSIC_TYPE_U8: IntrinsicOperationInfos[INTRINSIC_OPERATION_U8_BIT_AND],
		},
		INTRINSIC_TYPE_U16: {
			INTRINSIC_TYPE_U16: IntrinsicOperationInfos[INTRINSIC_OPERATION_U16_BIT_AND],
		},
		INTRINSIC_TYPE_U32: {
			INTRINSIC_TYPE_U32: IntrinsicOperationInfos[INTRINSIC_OPERATION_U32_BIT_AND],
		},
		INTRINSIC_TYPE_U64: {
			INTRINSIC_TYPE_U64: IntrinsicOperationInfos[INTRINSIC_OPERATION_U64_BIT_AND],
		},
	},
	INFIX_OPERATOR_BIT_XOR: {
		INTRINSIC_TYPE_BOOL: {
			INTRINSIC_TYPE_BOOL: IntrinsicOperationInfos[INTRINSIC_OPERATION_BOOL_BIT_XOR],
		},
		INTRINSIC_TYPE_I8: {
			INTRINSIC_TYPE_I8: IntrinsicOperationInfos[INTRINSIC_OPERATION_I8_BIT_XOR],
		},
		INTRINSIC_TYPE_I16: {
			INTRINSIC_TYPE_I16: IntrinsicOperationInfos[INTRINSIC_OPERATION_I16_BIT_XOR],
		},
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_BIT_XOR],
		},
		INTRINSIC_TYPE_I64: {
			INTRINSIC_TYPE_I64: IntrinsicOperationInfos[INTRINSIC_OPERATION_I64_BIT_XOR],
		},
		INTRINSIC_TYPE_U8: {
			INTRINSIC_TYPE_U8: IntrinsicOperationInfos[INTRINSIC_OPERATION_U8_BIT_XOR],
		},
		INTRINSIC_TYPE_U16: {
			INTRINSIC_TYPE_U16: IntrinsicOperationInfos[INTRINSIC_OPERATION_U16_BIT_XOR],
		},
		INTRINSIC_TYPE_U32: {
			INTRINSIC_TYPE_U32: IntrinsicOperationInfos[INTRINSIC_OPERATION_U32_BIT_XOR],
		},
		INTRINSIC_TYPE_U64: {
			INTRINSIC_TYPE_U64: IntrinsicOperationInfos[INTRINSIC_OPERATION_U64_BIT_XOR],
		},
	},
	INFIX_OPERATOR_BIT_IOR: {
		INTRINSIC_TYPE_BOOL: {
			INTRINSIC_TYPE_BOOL: IntrinsicOperationInfos[INTRINSIC_OPERATION_BOOL_BIT_IOR],
		},
		INTRINSIC_TYPE_I8: {
			INTRINSIC_TYPE_I8: IntrinsicOperationInfos[INTRINSIC_OPERATION_I8_BIT_IOR],
		},
		INTRINSIC_TYPE_I16: {
			INTRINSIC_TYPE_I16: IntrinsicOperationInfos[INTRINSIC_OPERATION_I16_BIT_IOR],
		},
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_BIT_IOR],
		},
		INTRINSIC_TYPE_I64: {
			INTRINSIC_TYPE_I64: IntrinsicOperationInfos[INTRINSIC_OPERATION_I64_BIT_IOR],
		},
		INTRINSIC_TYPE_U8: {
			INTRINSIC_TYPE_U8: IntrinsicOperationInfos[INTRINSIC_OPERATION_U8_BIT_IOR],
		},
		INTRINSIC_TYPE_U16: {
			INTRINSIC_TYPE_U16: IntrinsicOperationInfos[INTRINSIC_OPERATION_U16_BIT_IOR],
		},
		INTRINSIC_TYPE_U32: {
			INTRINSIC_TYPE_U32: IntrinsicOperationInfos[INTRINSIC_OPERATION_U32_BIT_IOR],
		},
		INTRINSIC_TYPE_U64: {
			INTRINSIC_TYPE_U64: IntrinsicOperationInfos[INTRINSIC_OPERATION_U64_BIT_IOR],
		},
	},
	INFIX_OPERATOR_CMP_EQ: {
		INTRINSIC_TYPE_BOOL: {
			INTRINSIC_TYPE_BOOL: IntrinsicOperationInfos[INTRINSIC_OPERATION_BOOL_CMP_EQ],
		},
		INTRINSIC_TYPE_I8: {
			INTRINSIC_TYPE_I8: IntrinsicOperationInfos[INTRINSIC_OPERATION_I8_CMP_EQ],
		},
		INTRINSIC_TYPE_I16: {
			INTRINSIC_TYPE_I16: IntrinsicOperationInfos[INTRINSIC_OPERATION_I16_CMP_EQ],
		},
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_CMP_EQ],
		},
		INTRINSIC_TYPE_I64: {
			INTRINSIC_TYPE_I64: IntrinsicOperationInfos[INTRINSIC_OPERATION_I64_CMP_EQ],
		},
		INTRINSIC_TYPE_U8: {
			INTRINSIC_TYPE_U8: IntrinsicOperationInfos[INTRINSIC_OPERATION_U8_CMP_EQ],
		},
		INTRINSIC_TYPE_U16: {
			INTRINSIC_TYPE_U16: IntrinsicOperationInfos[INTRINSIC_OPERATION_U16_CMP_EQ],
		},
		INTRINSIC_TYPE_U32: {
			INTRINSIC_TYPE_U32: IntrinsicOperationInfos[INTRINSIC_OPERATION_U32_CMP_EQ],
		},
		INTRINSIC_TYPE_U64: {
			INTRINSIC_TYPE_U64: IntrinsicOperationInfos[INTRINSIC_OPERATION_U64_CMP_EQ],
		},
		INTRINSIC_TYPE_F32: {
			INTRINSIC_TYPE_F32: IntrinsicOperationInfos[INTRINSIC_OPERATION_F32_CMP_EQ],
		},
		INTRINSIC_TYPE_F64: {
			INTRINSIC_TYPE_F64: IntrinsicOperationInfos[INTRINSIC_OPERATION_F64_CMP_EQ],
		},
	},
	INFIX_OPERATOR_CMP_NE: {
		INTRINSIC_TYPE_BOOL: {
			INTRINSIC_TYPE_BOOL: IntrinsicOperationInfos[INTRINSIC_OPERATION_BOOL_CMP_NE],
		},
		INTRINSIC_TYPE_I8: {
			INTRINSIC_TYPE_I8: IntrinsicOperationInfos[INTRINSIC_OPERATION_I8_CMP_NE],
		},
		INTRINSIC_TYPE_I16: {
			INTRINSIC_TYPE_I16: IntrinsicOperationInfos[INTRINSIC_OPERATION_I16_CMP_NE],
		},
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_CMP_NE],
		},
		INTRINSIC_TYPE_I64: {
			INTRINSIC_TYPE_I64: IntrinsicOperationInfos[INTRINSIC_OPERATION_I64_CMP_NE],
		},
		INTRINSIC_TYPE_U8: {
			INTRINSIC_TYPE_U8: IntrinsicOperationInfos[INTRINSIC_OPERATION_U8_CMP_NE],
		},
		INTRINSIC_TYPE_U16: {
			INTRINSIC_TYPE_U16: IntrinsicOperationInfos[INTRINSIC_OPERATION_U16_CMP_NE],
		},
		INTRINSIC_TYPE_U32: {
			INTRINSIC_TYPE_U32: IntrinsicOperationInfos[INTRINSIC_OPERATION_U32_CMP_NE],
		},
		INTRINSIC_TYPE_U64: {
			INTRINSIC_TYPE_U64: IntrinsicOperationInfos[INTRINSIC_OPERATION_U64_CMP_NE],
		},
		INTRINSIC_TYPE_F32: {
			INTRINSIC_TYPE_F32: IntrinsicOperationInfos[INTRINSIC_OPERATION_F32_CMP_NE],
		},
		INTRINSIC_TYPE_F64: {
			INTRINSIC_TYPE_F64: IntrinsicOperationInfos[INTRINSIC_OPERATION_F64_CMP_NE],
		},
	},
	INFIX_OPERATOR_AND: {
		INTRINSIC_TYPE_BOOL: {
			INTRINSIC_TYPE_BOOL: IntrinsicOperationInfos[INTRINSIC_OPERATION_BOOL_AND],
		},
	},
	INFIX_OPERATOR_OR: {
		INTRINSIC_TYPE_BOOL: {
			INTRINSIC_TYPE_BOOL: IntrinsicOperationInfos[INTRINSIC_OPERATION_BOOL_OR],
		},
	},
	INFIX_OPERATOR_MUL: {
		INTRINSIC_TYPE_I8: {
			INTRINSIC_TYPE_I8: IntrinsicOperationInfos[INTRINSIC_OPERATION_I8_MUL],
		},
		INTRINSIC_TYPE_I16: {
			INTRINSIC_TYPE_I16: IntrinsicOperationInfos[INTRINSIC_OPERATION_I16_MUL],
		},
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_MUL],
		},
		INTRINSIC_TYPE_I64: {
			INTRINSIC_TYPE_I64: IntrinsicOperationInfos[INTRINSIC_OPERATION_I64_MUL],
		},
		INTRINSIC_TYPE_U8: {
			INTRINSIC_TYPE_U8: IntrinsicOperationInfos[INTRINSIC_OPERATION_U8_MUL],
		},
		INTRINSIC_TYPE_U16: {
			INTRINSIC_TYPE_U16: IntrinsicOperationInfos[INTRINSIC_OPERATION_U16_MUL],
		},
		INTRINSIC_TYPE_U32: {
			INTRINSIC_TYPE_U32: IntrinsicOperationInfos[INTRINSIC_OPERATION_U32_MUL],
		},
		INTRINSIC_TYPE_U64: {
			INTRINSIC_TYPE_U64: IntrinsicOperationInfos[INTRINSIC_OPERATION_U64_MUL],
		},
		INTRINSIC_TYPE_F32: {
			INTRINSIC_TYPE_F32: IntrinsicOperationInfos[INTRINSIC_OPERATION_F32_MUL],
		},
		INTRINSIC_TYPE_F64: {
			INTRINSIC_TYPE_F64: IntrinsicOperationInfos[INTRINSIC_OPERATION_F64_MUL],
		},
	},
	INFIX_OPERATOR_DIV: {
		INTRINSIC_TYPE_I8: {
			INTRINSIC_TYPE_I8: IntrinsicOperationInfos[INTRINSIC_OPERATION_I8_DIV],
		},
		INTRINSIC_TYPE_I16: {
			INTRINSIC_TYPE_I16: IntrinsicOperationInfos[INTRINSIC_OPERATION_I16_DIV],
		},
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_DIV],
		},
		INTRINSIC_TYPE_I64: {
			INTRINSIC_TYPE_I64: IntrinsicOperationInfos[INTRINSIC_OPERATION_I64_DIV],
		},
		INTRINSIC_TYPE_U8: {
			INTRINSIC_TYPE_U8: IntrinsicOperationInfos[INTRINSIC_OPERATION_U8_DIV],
		},
		INTRINSIC_TYPE_U16: {
			INTRINSIC_TYPE_U16: IntrinsicOperationInfos[INTRINSIC_OPERATION_U16_DIV],
		},
		INTRINSIC_TYPE_U32: {
			INTRINSIC_TYPE_U32: IntrinsicOperationInfos[INTRINSIC_OPERATION_U32_DIV],
		},
		INTRINSIC_TYPE_U64: {
			INTRINSIC_TYPE_U64: IntrinsicOperationInfos[INTRINSIC_OPERATION_U64_DIV],
		},
		INTRINSIC_TYPE_F32: {
			INTRINSIC_TYPE_F32: IntrinsicOperationInfos[INTRINSIC_OPERATION_F32_DIV],
		},
		INTRINSIC_TYPE_F64: {
			INTRINSIC_TYPE_F64: IntrinsicOperationInfos[INTRINSIC_OPERATION_F64_DIV],
		},
	},
	INFIX_OPERATOR_MOD: {
		INTRINSIC_TYPE_I8: {
			INTRINSIC_TYPE_I8: IntrinsicOperationInfos[INTRINSIC_OPERATION_I8_MOD],
		},
		INTRINSIC_TYPE_I16: {
			INTRINSIC_TYPE_I16: IntrinsicOperationInfos[INTRINSIC_OPERATION_I16_MOD],
		},
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_MOD],
		},
		INTRINSIC_TYPE_I64: {
			INTRINSIC_TYPE_I64: IntrinsicOperationInfos[INTRINSIC_OPERATION_I64_MOD],
		},
		INTRINSIC_TYPE_U8: {
			INTRINSIC_TYPE_U8: IntrinsicOperationInfos[INTRINSIC_OPERATION_U8_MOD],
		},
		INTRINSIC_TYPE_U16: {
			INTRINSIC_TYPE_U16: IntrinsicOperationInfos[INTRINSIC_OPERATION_U16_MOD],
		},
		INTRINSIC_TYPE_U32: {
			INTRINSIC_TYPE_U32: IntrinsicOperationInfos[INTRINSIC_OPERATION_U32_MOD],
		},
		INTRINSIC_TYPE_U64: {
			INTRINSIC_TYPE_U64: IntrinsicOperationInfos[INTRINSIC_OPERATION_U64_MOD],
		},
	},
	INFIX_OPERATOR_ADD: {
		INTRINSIC_TYPE_I8: {
			INTRINSIC_TYPE_I8: IntrinsicOperationInfos[INTRINSIC_OPERATION_I8_ADD],
		},
		INTRINSIC_TYPE_I16: {
			INTRINSIC_TYPE_I16: IntrinsicOperationInfos[INTRINSIC_OPERATION_I16_ADD],
		},
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_ADD],
		},
		INTRINSIC_TYPE_I64: {
			INTRINSIC_TYPE_I64: IntrinsicOperationInfos[INTRINSIC_OPERATION_I64_ADD],
		},
		INTRINSIC_TYPE_U8: {
			INTRINSIC_TYPE_U8: IntrinsicOperationInfos[INTRINSIC_OPERATION_U8_ADD],
		},
		INTRINSIC_TYPE_U16: {
			INTRINSIC_TYPE_U16: IntrinsicOperationInfos[INTRINSIC_OPERATION_U16_ADD],
		},
		INTRINSIC_TYPE_U32: {
			INTRINSIC_TYPE_U32: IntrinsicOperationInfos[INTRINSIC_OPERATION_U32_ADD],
		},
		INTRINSIC_TYPE_U64: {
			INTRINSIC_TYPE_U64: IntrinsicOperationInfos[INTRINSIC_OPERATION_U64_ADD],
		},
		INTRINSIC_TYPE_F32: {
			INTRINSIC_TYPE_F32: IntrinsicOperationInfos[INTRINSIC_OPERATION_F32_ADD],
		},
		INTRINSIC_TYPE_F64: {
			INTRINSIC_TYPE_F64: IntrinsicOperationInfos[INTRINSIC_OPERATION_F64_ADD],
		},
	},
	INFIX_OPERATOR_SUB: {
		INTRINSIC_TYPE_I8: {
			INTRINSIC_TYPE_I8: IntrinsicOperationInfos[INTRINSIC_OPERATION_I8_SUB],
		},
		INTRINSIC_TYPE_I16: {
			INTRINSIC_TYPE_I16: IntrinsicOperationInfos[INTRINSIC_OPERATION_I16_SUB],
		},
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_SUB],
		},
		INTRINSIC_TYPE_I64: {
			INTRINSIC_TYPE_I64: IntrinsicOperationInfos[INTRINSIC_OPERATION_I64_SUB],
		},
		INTRINSIC_TYPE_U8: {
			INTRINSIC_TYPE_U8: IntrinsicOperationInfos[INTRINSIC_OPERATION_U8_SUB],
		},
		INTRINSIC_TYPE_U16: {
			INTRINSIC_TYPE_U16: IntrinsicOperationInfos[INTRINSIC_OPERATION_U16_SUB],
		},
		INTRINSIC_TYPE_U32: {
			INTRINSIC_TYPE_U32: IntrinsicOperationInfos[INTRINSIC_OPERATION_U32_SUB],
		},
		INTRINSIC_TYPE_U64: {
			INTRINSIC_TYPE_U64: IntrinsicOperationInfos[INTRINSIC_OPERATION_U64_SUB],
		},
		INTRINSIC_TYPE_F32: {
			INTRINSIC_TYPE_F32: IntrinsicOperationInfos[INTRINSIC_OPERATION_F32_SUB],
		},
		INTRINSIC_TYPE_F64: {
			INTRINSIC_TYPE_F64: IntrinsicOperationInfos[INTRINSIC_OPERATION_F64_SUB],
		},
	},
	INFIX_OPERATOR_BIT_LSH: {
		INTRINSIC_TYPE_I8: {
			INTRINSIC_TYPE_I8: IntrinsicOperationInfos[INTRINSIC_OPERATION_I8_BIT_LSH],
		},
		INTRINSIC_TYPE_I16: {
			INTRINSIC_TYPE_I16: IntrinsicOperationInfos[INTRINSIC_OPERATION_I16_BIT_LSH],
		},
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_BIT_LSH],
		},
		INTRINSIC_TYPE_I64: {
			INTRINSIC_TYPE_I64: IntrinsicOperationInfos[INTRINSIC_OPERATION_I64_BIT_LSH],
		},
		INTRINSIC_TYPE_U8: {
			INTRINSIC_TYPE_U8: IntrinsicOperationInfos[INTRINSIC_OPERATION_U8_BIT_LSH],
		},
		INTRINSIC_TYPE_U16: {
			INTRINSIC_TYPE_U16: IntrinsicOperationInfos[INTRINSIC_OPERATION_U16_BIT_LSH],
		},
		INTRINSIC_TYPE_U32: {
			INTRINSIC_TYPE_U32: IntrinsicOperationInfos[INTRINSIC_OPERATION_U32_BIT_LSH],
		},
		INTRINSIC_TYPE_U64: {
			INTRINSIC_TYPE_U64: IntrinsicOperationInfos[INTRINSIC_OPERATION_U64_BIT_LSH],
		},
	},
	INFIX_OPERATOR_BIT_RSH: {
		INTRINSIC_TYPE_I8: {
			INTRINSIC_TYPE_I8: IntrinsicOperationInfos[INTRINSIC_OPERATION_I8_BIT_RSH],
		},
		INTRINSIC_TYPE_I16: {
			INTRINSIC_TYPE_I16: IntrinsicOperationInfos[INTRINSIC_OPERATION_I16_BIT_RSH],
		},
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_BIT_RSH],
		},
		INTRINSIC_TYPE_I64: {
			INTRINSIC_TYPE_I64: IntrinsicOperationInfos[INTRINSIC_OPERATION_I64_BIT_RSH],
		},
		INTRINSIC_TYPE_U8: {
			INTRINSIC_TYPE_U8: IntrinsicOperationInfos[INTRINSIC_OPERATION_U8_BIT_RSH],
		},
		INTRINSIC_TYPE_U16: {
			INTRINSIC_TYPE_U16: IntrinsicOperationInfos[INTRINSIC_OPERATION_U16_BIT_RSH],
		},
		INTRINSIC_TYPE_U32: {
			INTRINSIC_TYPE_U32: IntrinsicOperationInfos[INTRINSIC_OPERATION_U32_BIT_RSH],
		},
		INTRINSIC_TYPE_U64: {
			INTRINSIC_TYPE_U64: IntrinsicOperationInfos[INTRINSIC_OPERATION_U64_BIT_RSH],
		},
	},
	INFIX_OPERATOR_CMP_LT: {
		INTRINSIC_TYPE_I8: {
			INTRINSIC_TYPE_I8: IntrinsicOperationInfos[INTRINSIC_OPERATION_I8_CMP_LT],
		},
		INTRINSIC_TYPE_I16: {
			INTRINSIC_TYPE_I16: IntrinsicOperationInfos[INTRINSIC_OPERATION_I16_CMP_LT],
		},
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_CMP_LT],
		},
		INTRINSIC_TYPE_I64: {
			INTRINSIC_TYPE_I64: IntrinsicOperationInfos[INTRINSIC_OPERATION_I64_CMP_LT],
		},
		INTRINSIC_TYPE_U8: {
			INTRINSIC_TYPE_U8: IntrinsicOperationInfos[INTRINSIC_OPERATION_U8_CMP_LT],
		},
		INTRINSIC_TYPE_U16: {
			INTRINSIC_TYPE_U16: IntrinsicOperationInfos[INTRINSIC_OPERATION_U16_CMP_LT],
		},
		INTRINSIC_TYPE_U32: {
			INTRINSIC_TYPE_U32: IntrinsicOperationInfos[INTRINSIC_OPERATION_U32_CMP_LT],
		},
		INTRINSIC_TYPE_U64: {
			INTRINSIC_TYPE_U64: IntrinsicOperationInfos[INTRINSIC_OPERATION_U64_CMP_LT],
		},
		INTRINSIC_TYPE_F32: {
			INTRINSIC_TYPE_F32: IntrinsicOperationInfos[INTRINSIC_OPERATION_F32_CMP_LT],
		},
		INTRINSIC_TYPE_F64: {
			INTRINSIC_TYPE_F64: IntrinsicOperationInfos[INTRINSIC_OPERATION_F64_CMP_LT],
		},
	},
	INFIX_OPERATOR_CMP_LE: {
		INTRINSIC_TYPE_I8: {
			INTRINSIC_TYPE_I8: IntrinsicOperationInfos[INTRINSIC_OPERATION_I8_CMP_LE],
		},
		INTRINSIC_TYPE_I16: {
			INTRINSIC_TYPE_I16: IntrinsicOperationInfos[INTRINSIC_OPERATION_I16_CMP_LE],
		},
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_CMP_LE],
		},
		INTRINSIC_TYPE_I64: {
			INTRINSIC_TYPE_I64: IntrinsicOperationInfos[INTRINSIC_OPERATION_I64_CMP_LE],
		},
		INTRINSIC_TYPE_U8: {
			INTRINSIC_TYPE_U8: IntrinsicOperationInfos[INTRINSIC_OPERATION_U8_CMP_LE],
		},
		INTRINSIC_TYPE_U16: {
			INTRINSIC_TYPE_U16: IntrinsicOperationInfos[INTRINSIC_OPERATION_U16_CMP_LE],
		},
		INTRINSIC_TYPE_U32: {
			INTRINSIC_TYPE_U32: IntrinsicOperationInfos[INTRINSIC_OPERATION_U32_CMP_LE],
		},
		INTRINSIC_TYPE_U64: {
			INTRINSIC_TYPE_U64: IntrinsicOperationInfos[INTRINSIC_OPERATION_U64_CMP_LE],
		},
		INTRINSIC_TYPE_F32: {
			INTRINSIC_TYPE_F32: IntrinsicOperationInfos[INTRINSIC_OPERATION_F32_CMP_LE],
		},
		INTRINSIC_TYPE_F64: {
			INTRINSIC_TYPE_F64: IntrinsicOperationInfos[INTRINSIC_OPERATION_F64_CMP_LE],
		},
	},
	INFIX_OPERATOR_CMP_GT: {
		INTRINSIC_TYPE_I8: {
			INTRINSIC_TYPE_I8: IntrinsicOperationInfos[INTRINSIC_OPERATION_I8_CMP_GT],
		},
		INTRINSIC_TYPE_I16: {
			INTRINSIC_TYPE_I16: IntrinsicOperationInfos[INTRINSIC_OPERATION_I16_CMP_GT],
		},
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_CMP_GT],
		},
		INTRINSIC_TYPE_I64: {
			INTRINSIC_TYPE_I64: IntrinsicOperationInfos[INTRINSIC_OPERATION_I64_CMP_GT],
		},
		INTRINSIC_TYPE_U8: {
			INTRINSIC_TYPE_U8: IntrinsicOperationInfos[INTRINSIC_OPERATION_U8_CMP_GT],
		},
		INTRINSIC_TYPE_U16: {
			INTRINSIC_TYPE_U16: IntrinsicOperationInfos[INTRINSIC_OPERATION_U16_CMP_GT],
		},
		INTRINSIC_TYPE_U32: {
			INTRINSIC_TYPE_U32: IntrinsicOperationInfos[INTRINSIC_OPERATION_U32_CMP_GT],
		},
		INTRINSIC_TYPE_U64: {
			INTRINSIC_TYPE_U64: IntrinsicOperationInfos[INTRINSIC_OPERATION_U64_CMP_GT],
		},
		INTRINSIC_TYPE_F32: {
			INTRINSIC_TYPE_F32: IntrinsicOperationInfos[INTRINSIC_OPERATION_F32_CMP_GT],
		},
		INTRINSIC_TYPE_F64: {
			INTRINSIC_TYPE_F64: IntrinsicOperationInfos[INTRINSIC_OPERATION_F64_CMP_GT],
		},
	},
	INFIX_OPERATOR_CMP_GE: {
		INTRINSIC_TYPE_I8: {
			INTRINSIC_TYPE_I8: IntrinsicOperationInfos[INTRINSIC_OPERATION_I8_CMP_GE],
		},
		INTRINSIC_TYPE_I16: {
			INTRINSIC_TYPE_I16: IntrinsicOperationInfos[INTRINSIC_OPERATION_I16_CMP_GE],
		},
		INTRINSIC_TYPE_I32: {
			INTRINSIC_TYPE_I32: IntrinsicOperationInfos[INTRINSIC_OPERATION_I32_CMP_GE],
		},
		INTRINSIC_TYPE_I64: {
			INTRINSIC_TYPE_I64: IntrinsicOperationInfos[INTRINSIC_OPERATION_I64_CMP_GE],
		},
		INTRINSIC_TYPE_U8: {
			INTRINSIC_TYPE_U8: IntrinsicOperationInfos[INTRINSIC_OPERATION_U8_CMP_GE],
		},
		INTRINSIC_TYPE_U16: {
			INTRINSIC_TYPE_U16: IntrinsicOperationInfos[INTRINSIC_OPERATION_U16_CMP_GE],
		},
		INTRINSIC_TYPE_U32: {
			INTRINSIC_TYPE_U32: IntrinsicOperationInfos[INTRINSIC_OPERATION_U32_CMP_GE],
		},
		INTRINSIC_TYPE_U64: {
			INTRINSIC_TYPE_U64: IntrinsicOperationInfos[INTRINSIC_OPERATION_U64_CMP_GE],
		},
		INTRINSIC_TYPE_F32: {
			INTRINSIC_TYPE_F32: IntrinsicOperationInfos[INTRINSIC_OPERATION_F32_CMP_GE],
		},
		INTRINSIC_TYPE_F64: {
			INTRINSIC_TYPE_F64: IntrinsicOperationInfos[INTRINSIC_OPERATION_F64_CMP_GE],
		},
	},
}
