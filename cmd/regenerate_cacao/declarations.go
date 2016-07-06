package main

import (
	"github.com/ncbray/cacao/irgen"
)

var ast = &irgen.Declarations{
	DataSource: "regenerate_cacao",
	Package:    "github.com/ncbray/cacao/language",
	File:       "ast.go",
	Trees: []*irgen.TreeDecl{
		{
			Dump: true,
			Structs: []*irgen.StructDecl{
				{
					Name: "Token",
					Fields: []*irgen.FieldDecl{
						{
							Name: "Pos",
							Type: "int",
						},
						{
							Name: "Text",
							Type: "string",
						},
					},
				},
				{
					Name: "Param",
					Fields: []*irgen.FieldDecl{
						{
							Name: "Name",
							Type: "Token",
							Ref:  true,
						},
						{
							Name: "Type",
							Type: "TypeRef",
						},
					},
				},
				{
					Name: "File",
					Fields: []*irgen.FieldDecl{
						{
							Name: "Decls",
							Type: "FunctionDecl",
							Ref:  true,
							List: true,
						},
					},
				},
			},
			Groups: []*irgen.GroupDecl{
				{
					Name: "TypeRef",
					Structs: []*irgen.StructDecl{
						{
							Name: "TypeName",
							Fields: []*irgen.FieldDecl{
								{
									Name: "Name",
									Type: "Token",
									Ref:  true,
								},
							},
						},
					},
				},
				{
					Name: "Expr",
					Structs: []*irgen.StructDecl{
						{
							Name: "GetName",
							Fields: []*irgen.FieldDecl{
								{
									Name: "Name",
									Type: "Token",
									Ref:  true,
								},
							},
						},
						{
							Name: "IntegerLiteral",
							Fields: []*irgen.FieldDecl{
								{
									Name: "Pos",
									Type: "int",
								},
								{
									Name: "Text",
									Type: "string",
								},
							},
						},
						{
							Name: "PrefixExpr",
							Fields: []*irgen.FieldDecl{
								{
									Name: "Op",
									Type: "PrefixOperator",
								},
								{
									Name: "Expr",
									Type: "Expr",
								},
							},
						},
						{
							Name: "InfixExpr",
							Fields: []*irgen.FieldDecl{
								{
									Name: "Pos",
									Type: "int",
								},
								{
									Name: "Left",
									Type: "Expr",
								},
								{
									Name: "Op",
									Type: "InfixOperator",
								},
								{
									Name: "Right",
									Type: "Expr",
								},
							},
						},
						{
							Name: "Return",
							Fields: []*irgen.FieldDecl{
								{
									Name: "Pos",
									Type: "int",
								},
								{
									Name: "Expr",
									Type: "Expr",
								},
							},
						},
						{
							Name: "Call",
							Fields: []*irgen.FieldDecl{
								{
									Name: "Pos",
									Type: "int",
								},
								{
									Name: "Expr",
									Type: "Expr",
								},
								{
									Name: "Args",
									Type: "Expr",
									List: true,
								},
							},
						},
						{
							Name: "If",
							Fields: []*irgen.FieldDecl{
								{
									Name: "Pos",
									Type: "int",
								},
								{
									Name: "Cond",
									Type: "Expr",
								},
								{
									Name: "T",
									Type: "Block",
									Ref:  true,
								},
								{
									Name: "F",
									Type: "Block",
									Ref:  true,
								},
							},
						},
						{
							Name: "Block",
							Fields: []*irgen.FieldDecl{
								{
									Name: "Pos",
									Type: "int",
								},
								{
									Name: "Exprs",
									Type: "Expr",
									List: true,
								},
							},
						},
						{
							Name: "FunctionDecl",
							Fields: []*irgen.FieldDecl{
								{
									Name: "Name",
									Type: "Token",
									Ref:  true,
								},
								{
									Name: "Params",
									Type: "Param",
									Ref:  true,
									List: true,
								},
								{
									Name: "Results",
									Type: "Param",
									Ref:  true,
									List: true,
								},
								{
									Name: "Body",
									Type: "Block",
									Ref:  true,
								},
							},
						},
					},
				},
			},
		},
	},
}

var struct_ir = &irgen.Declarations{
	DataSource: "regenerate_cacao",
	Package:    "github.com/ncbray/cacao/irgen",
	File:       "structir.go",
	Trees: []*irgen.TreeDecl{
		{
			Dump: true,
			Structs: []*irgen.StructDecl{
				{
					Name: "LLField",
					Fields: []*irgen.FieldDecl{
						{
							Name: "Name",
							Type: "string",
						},
						{
							Name: "Type",
							Type: "LLType",
						},
						{
							Name: "Export",
							Type: "bool",
						},
					},
				},
				{
					Name: "LLStruct",
					Fields: []*irgen.FieldDecl{
						{
							Name: "Name",
							Type: "string",
						},
						{
							Name: "Fields",
							Type: "LLField",
							Ref:  true,
							List: true,
						},
						{
							Name: "Export",
							Type: "bool",
						},
					},
				},
			},
			Groups: []*irgen.GroupDecl{
				{
					Name: "LLType",
					Structs: []*irgen.StructDecl{
						{
							Name: "IntrinsicType",
							Fields: []*irgen.FieldDecl{
								{
									Name: "Element",
									Type: "string",
								},
							},
						},
						{
							Name: "StructType",
							Fields: []*irgen.FieldDecl{
								{
									Name: "Element",
									Type: "LLStruct",
									Ref:  true,
								},
							},
						},
						{
							Name: "PointerType",
							Fields: []*irgen.FieldDecl{
								{
									Name: "Element",
									Type: "LLType",
								},
							},
						},
						{
							Name: "ListType",
							Fields: []*irgen.FieldDecl{
								{
									Name: "Element",
									Type: "LLType",
								},
							},
						},
					},
				},
			},
		},
	},
}

var graph = &irgen.Declarations{
	DataSource: "regenerate_cacao",
	Package:    "github.com/ncbray/cacao/language",
	File:       "graph.go",
	Nodes: []*irgen.NodeDecl{
		{
			Name: "Program",
		},
		{
			Name:   "FunctionGraph",
			Region: "Program",
		},
		{
			Name:   "ControlRegion",
			Region: "FunctionGraph",
		},
		{
			Name:   "Operation",
			Region: "FunctionGraph",
		},
		{
			Name:   "Value",
			Region: "FunctionGraph",
		},
	},
	Parents: []*irgen.ParentRelationshipDecl{
		{
			Src:      "FunctionGraph",
			SrcName:  "ControlRegions",
			Dst:      "ControlRegion",
			DstName:  "Function",
			Required: true,
		},
		{
			Src:      "ControlRegion",
			SrcName:  "Ops",
			Dst:      "Operation",
			DstName:  "Control",
			Required: true,
		},
	},
	Counted: []*irgen.CountedRelationshipDecl{
		{
			Src:     "ControlRegion",
			SrcName: "Transfers",
			Dst:     "ControlRegion",
			DstName: "Merges",
		},
		{
			Src:     "Operation",
			SrcName: "Inputs",
			Dst:     "Value",
			DstName: "Uses",
		},
		{
			Src:     "Operation",
			SrcName: "Outputs",
			Dst:     "Value",
			DstName: "Defs",
		},
	},
}
