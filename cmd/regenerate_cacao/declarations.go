package main

import (
	"github.com/ncbray/cacao/regenerate"
)

var declarations = &regenerate.Declarations{
	Trees: []*regenerate.TreeDecl{
		{
			DataSource: "regenerate_cacao",
			Package:    "github.com/ncbray/cacao/language",
			File:       "ast.go",
			Dump:       "dump.go",
			Nodes: []*regenerate.StructDecl{
				{
					Name: "Token",
					Fields: []*regenerate.FieldDecl{
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
					Fields: []*regenerate.FieldDecl{
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
					Fields: []*regenerate.FieldDecl{
						{
							Name: "Decls",
							Type: "FunctionDecl",
							Ref:  true,
							List: true,
						},
					},
				},
			},
			Groups: []*regenerate.GroupDecl{
				{
					Name: "TypeRef",
					Nodes: []*regenerate.StructDecl{
						{
							Name: "TypeName",
							Fields: []*regenerate.FieldDecl{
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
					Nodes: []*regenerate.StructDecl{
						{
							Name: "GetName",
							Fields: []*regenerate.FieldDecl{
								{
									Name: "Name",
									Type: "Token",
									Ref:  true,
								},
							},
						},
						{
							Name: "IntegerLiteral",
							Fields: []*regenerate.FieldDecl{
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
							Fields: []*regenerate.FieldDecl{
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
							Fields: []*regenerate.FieldDecl{
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
							Fields: []*regenerate.FieldDecl{
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
							Fields: []*regenerate.FieldDecl{
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
							Fields: []*regenerate.FieldDecl{
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
							Fields: []*regenerate.FieldDecl{
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
							Fields: []*regenerate.FieldDecl{
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
