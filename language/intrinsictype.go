package language

/* Generated from regenerate_cacao, do not edit by hand. */

type IntrinsicType uint8

const (
	INTRINSIC_TYPE_BOOL IntrinsicType = 0
	INTRINSIC_TYPE_I8   IntrinsicType = 1
	INTRINSIC_TYPE_I16  IntrinsicType = 2
	INTRINSIC_TYPE_I32  IntrinsicType = 3
	INTRINSIC_TYPE_I64  IntrinsicType = 4
	INTRINSIC_TYPE_U8   IntrinsicType = 5
	INTRINSIC_TYPE_U16  IntrinsicType = 6
	INTRINSIC_TYPE_U32  IntrinsicType = 7
	INTRINSIC_TYPE_U64  IntrinsicType = 8
	INTRINSIC_TYPE_F32  IntrinsicType = 9
	INTRINSIC_TYPE_F64  IntrinsicType = 10
	NUM_INTRINSIC_TYPE  IntrinsicType = 11
)

type IntrinsicTypeInfo struct {
	Enum     IntrinsicType
	Name     string
	Text     string
	Bits     int
	Unsigned bool
	Float    bool
}

var IntrinsicTypeInfos = []*IntrinsicTypeInfo{
	{
		Enum:     INTRINSIC_TYPE_BOOL,
		Name:     "INTRINSIC_TYPE_BOOL",
		Text:     "bool",
		Bits:     1,
		Unsigned: true,
	},
	{
		Enum: INTRINSIC_TYPE_I8,
		Name: "INTRINSIC_TYPE_I8",
		Text: "i8",
		Bits: 8,
	},
	{
		Enum: INTRINSIC_TYPE_I16,
		Name: "INTRINSIC_TYPE_I16",
		Text: "i16",
		Bits: 16,
	},
	{
		Enum: INTRINSIC_TYPE_I32,
		Name: "INTRINSIC_TYPE_I32",
		Text: "i32",
		Bits: 32,
	},
	{
		Enum: INTRINSIC_TYPE_I64,
		Name: "INTRINSIC_TYPE_I64",
		Text: "i64",
		Bits: 64,
	},
	{
		Enum:     INTRINSIC_TYPE_U8,
		Name:     "INTRINSIC_TYPE_U8",
		Text:     "u8",
		Bits:     8,
		Unsigned: true,
	},
	{
		Enum:     INTRINSIC_TYPE_U16,
		Name:     "INTRINSIC_TYPE_U16",
		Text:     "u16",
		Bits:     16,
		Unsigned: true,
	},
	{
		Enum:     INTRINSIC_TYPE_U32,
		Name:     "INTRINSIC_TYPE_U32",
		Text:     "u32",
		Bits:     32,
		Unsigned: true,
	},
	{
		Enum:     INTRINSIC_TYPE_U64,
		Name:     "INTRINSIC_TYPE_U64",
		Text:     "u64",
		Bits:     64,
		Unsigned: true,
	},
	{
		Enum:  INTRINSIC_TYPE_F32,
		Name:  "INTRINSIC_TYPE_F32",
		Text:  "f32",
		Bits:  32,
		Float: true,
	},
	{
		Enum:  INTRINSIC_TYPE_F64,
		Name:  "INTRINSIC_TYPE_F64",
		Text:  "f64",
		Bits:  64,
		Float: true,
	},
}

func (value IntrinsicType) String() string {
	if value < 0 || value >= 11 {
		return "INVALID_INTRINSIC_TYPE"
	}
	return IntrinsicTypeInfos[value].Name
}
