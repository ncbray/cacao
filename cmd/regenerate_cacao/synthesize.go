package main

import (
	"github.com/ncbray/cacao/regenerate"
	"github.com/ncbray/compilerutil/fs"
	"github.com/ncbray/compilerutil/goutil"
)

type IntrinsicTypeInfo struct {
	Name     string
	Text     string
	Bits     int
	Unsigned bool
	Float    bool
}

func getIntrinsicTypes(fsys fs.FileSystem) []*IntrinsicTypeInfo {
	inf := fsys.InputFile("dsl/enum/intrinsic_type.csv")
	inp, err := inf.GetReader()
	if err != nil {
		panic(err)
	}
	data := []*IntrinsicTypeInfo{}
	err = goutil.ReadCSV(inp, &data)
	if err != nil {
		panic(err)
	}
	return data
}

type PrefixOperatorInfo struct {
	Name       string
	Text       string
	Logical    bool
	Bitwise    bool
	IntArith   bool
	FloatArith bool
}

func getPrefixOperators(fsys fs.FileSystem) []*PrefixOperatorInfo {
	inf := fsys.InputFile("dsl/enum/prefix_operator.csv")
	inp, err := inf.GetReader()
	if err != nil {
		panic(err)
	}
	data := []*PrefixOperatorInfo{}
	err = goutil.ReadCSV(inp, &data)
	if err != nil {
		panic(err)
	}
	return data
}

type InfixOperatorInfo struct {
	Name        string
	Text        string
	Precedence  int
	IntArith    bool
	FloatArith  bool
	Shift       bool
	Bitwise     bool
	Equality    bool
	Compare     bool
	Shortcircut bool
}

func getInfixOperators(fsys fs.FileSystem) []*InfixOperatorInfo {
	inf := fsys.InputFile("dsl/enum/infix_operator.csv")
	inp, err := inf.GetReader()
	if err != nil {
		panic(err)
	}
	data := []*InfixOperatorInfo{}
	err = goutil.ReadCSV(inp, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func getBoolType(intrinsic_types []*IntrinsicTypeInfo) *IntrinsicTypeInfo {
	for _, intrinsic_type := range intrinsic_types {
		if intrinsic_type.Bits == 1 {
			return intrinsic_type
		}
	}
	return nil
}

func Synthesize(data_dir string, decls []*regenerate.Declarations, fsys fs.FileSystem) []*regenerate.Declarations {
	fsys = fs.MakeRelative(fsys, data_dir)
	intrinsic_types := getIntrinsicTypes(fsys)
	prefix_operators := getPrefixOperators(fsys)
	infix_operators := getInfixOperators(fsys)

	bool_type := getBoolType(intrinsic_types)

	// Generate intrinsic prefix ops.
	for _, intrinsic_type := range intrinsic_types {
		for _, prefix_op := range prefix_operators {
			if intrinsic_type.Float {
				if !prefix_op.FloatArith {
					continue
				}
			} else if intrinsic_type.Bits == 1 {
				if !(prefix_op.Logical || prefix_op.Bitwise) {
					continue
				}
			}

			//result := intrinsic_type
			//if prefix_op.Logical {
			//	result = bool_type
			//}
		}
	}

	instances := make([]*regenerate.EnumInstance, len(intrinsic_types))
	for i, intrinsic_type := range intrinsic_types {
		instances[i] = &regenerate.EnumInstance{
			Name: intrinsic_type.Name,
			Fields: []*regenerate.EnumInstanceField{
				{
					Name:  "Text",
					Value: intrinsic_type.Text,
				},
				{
					Name:  "Bits",
					Value: intrinsic_type.Bits,
				},
				{
					Name:  "Unsigned",
					Value: intrinsic_type.Unsigned,
				},
				{
					Name:  "Float",
					Value: intrinsic_type.Float,
				},
			},
		}
	}

	decls = append(decls, &regenerate.Declarations{
		DataSource: "regenerate_cacao",
		Package:    "github.com/ncbray/cacao/language",
		File:       "intrinsictype.go",
		Enums: []*regenerate.EnumDecl{
			&regenerate.EnumDecl{
				Name:   "IntrinsicType",
				Prefix: "INTRINSIC_TYPE",
				Fields: []*regenerate.EnumDeclField{
					{
						Name: "Text",
						Type: "string",
					},
					{
						Name: "Bits",
						Type: "int",
					},
					{
						Name: "Unsigned",
						Type: "bool",
					},
					{
						Name: "Float",
						Type: "bool",
					},
				},
				Enums: instances,
			},
		},
	})

	instances = make([]*regenerate.EnumInstance, len(infix_operators))
	for i, infix_op := range infix_operators {
		instances[i] = &regenerate.EnumInstance{
			Name: infix_op.Name,
			Fields: []*regenerate.EnumInstanceField{
				{
					Name:  "Text",
					Value: infix_op.Text,
				},
				{
					Name:  "Precedence",
					Value: infix_op.Precedence,
				},
			},
		}
	}

	decls = append(decls, &regenerate.Declarations{
		DataSource: "regenerate_cacao",
		Package:    "github.com/ncbray/cacao/language",
		File:       "infixoperator.go",
		Enums: []*regenerate.EnumDecl{
			&regenerate.EnumDecl{
				Name:           "InfixOperator",
				Prefix:         "INFIX_OPERATOR",
				GenerateParser: "Text",
				Fields: []*regenerate.EnumDeclField{
					{
						Name: "Text",
						Type: "string",
					},
					{
						Name: "Precedence",
						Type: "int",
					},
				},
				Enums: instances,
			},
		},
	})

	instances = []*regenerate.EnumInstance{}

	// Generate intrinsic infix ops.
	for _, intrinsic_type := range intrinsic_types {
		for _, infix_op := range infix_operators {
			if intrinsic_type.Float {
				if !(infix_op.FloatArith || infix_op.Equality || infix_op.Compare) {
					continue
				}
			} else if intrinsic_type.Bits == 1 {
				if !(infix_op.Equality || infix_op.Bitwise || infix_op.Shortcircut) {
					continue
				}
			} else {
				if infix_op.Shortcircut {
					continue
				}
			}
			result := intrinsic_type
			if infix_op.Equality || infix_op.Compare || infix_op.Shortcircut {
				result = bool_type
			}
			instances = append(instances, &regenerate.EnumInstance{
				Name: intrinsic_type.Name + "_" + infix_op.Name,
				Fields: []*regenerate.EnumInstanceField{
					{
						Name:  "InfixOperator",
						Value: "INFIX_OPERATOR_" + infix_op.Name,
					},
					{
						Name:  "InfixLeft",
						Value: "INTRINSIC_TYPE_" + intrinsic_type.Name,
					},
					{
						Name:  "InfixRight",
						Value: "INTRINSIC_TYPE_" + intrinsic_type.Name,
					},
					{
						Name:  "InfixResult",
						Value: "INTRINSIC_TYPE_" + result.Name,
					},
				},
			})
		}
	}

	decls = append(decls, &regenerate.Declarations{
		DataSource: "regenerate_cacao",
		Package:    "github.com/ncbray/cacao/language",
		File:       "intrinsicoperation.go",

		Enums: []*regenerate.EnumDecl{
			&regenerate.EnumDecl{
				Name:   "IntrinsicOperation",
				Prefix: "INTRINSIC_OPERATION",
				Fields: []*regenerate.EnumDeclField{
					{
						Name: "InfixOperator",
						Type: "InfixOperator",
					},
					{
						Name: "InfixLeft",
						Type: "IntrinsicType",
					},
					{
						Name: "InfixRight",
						Type: "IntrinsicType",
					},
					{
						Name: "InfixResult",
						Type: "IntrinsicType",
					},
				},
				Indexes: []*regenerate.EnumIndex{
					{
						Name: "InfixToIntrinsic",
						Path: []string{
							"InfixOperator",
							"InfixLeft",
							"InfixRight",
						},
					},
				},
				Enums: instances,
			},
		},
	})

	return decls
}
