package language

import (
	"fmt"
	"github.com/ncbray/cacao/framework"
	"strconv"
)

type abstractValue interface {
	Pos() int
	Type() Type
}

type typedValue struct {
	P int
	T Type
}

func (value *typedValue) Pos() int {
	return value.P
}

func (value *typedValue) Type() Type {
	return value.T
}

type functionValue struct {
	P int
	F *Function
}

func (value *functionValue) Pos() int {
	return value.P
}

func (value *functionValue) Type() Type {
	return value.F.Type
}

type poisonValue struct {
}

func (value *poisonValue) Pos() int {
	return -1
}

func (value *poisonValue) Type() Type {
	return nil
}

var poison abstractValue = &poisonValue{}

func PrintableTypeName(t Type) string {
	switch t := t.(type) {
	case *IntrinsicTypeX:
		return IntrinsicTypeInfos[t.UID].Text
	default:
		panic(t)
	}
}

type semanticPass struct {
	namespace      *Namespace
	locals         *Namespace
	returnType     Type
	intrinsicTypes []*IntrinsicTypeX
	status         framework.CompileStatus
}

func (semantic *semanticPass) registerName(name *Token, obj Nameable) {
	if semantic.namespace.IsDefined(name.Text) {
		semantic.status.LocationError(name.Pos, fmt.Sprintf("tried to redefined name %#v", name.Text))
	} else {
		semantic.namespace.Set(name.Text, obj)
	}
}

func (semantic *semanticPass) indexModuleNamespace(file *File) {
	for _, node := range file.Decls {
		f := &Function{
			Name: node.Name.Text,
		}
		semantic.registerName(node.Name, f)
	}
}

func (semantic *semanticPass) typeTypeRef(node TypeRef) Type {
	switch node := node.(type) {
	case *TypeName:
		name := node.Name.Text
		if !semantic.namespace.IsDefined(name) {
			semantic.status.LocationError(node.Name.Pos, fmt.Sprintf("cannot resolve name %#v", name))
			return nil
		}
		o := semantic.namespace.index[name]
		t, ok := o.(Type)
		if !ok {
			semantic.status.LocationError(node.Name.Pos, fmt.Sprintf("%#v is not a type", name))
			return nil
		}
		return t
	default:
		panic(node)
	}
}

func (semantic *semanticPass) typeDecl(node *FunctionDecl) {
	f := semantic.recoverFunction(node)

	ptypes := make([]Type, len(node.Params))
	for i, p := range node.Params {
		t := semantic.typeTypeRef(p.Type)
		ptypes[i] = t
		slot := &LocalSlot{
			Name: p.Name.Text,
			Type: t,
		}
		f.Locals = append(f.Locals, slot)
		f.NumParams += 1
	}

	rtypes := make([]Type, len(node.Results))
	for i, p := range node.Results {
		t := semantic.typeTypeRef(p.Type)
		rtypes[i] = t
		// TODO named locals
		if p.Name != nil {
			panic(p)
		}
	}

	var rtype Type
	if len(rtypes) == 0 {
		rtype = &VoidType{}
	} else if len(rtypes) == 1 {
		rtype = rtypes[0]
	} else {
		panic(rtypes)
	}

	f.Type = &FunctionType{Params: ptypes, Result: rtype}
}

func (semantic *semanticPass) typeFile(file *File) {
	for _, d := range file.Decls {
		semantic.typeDecl(d)
	}
}

func (semantic *semanticPass) opBlock(node *Block) abstractValue {
	for _, child := range node.Exprs {
		semantic.opExpr(child)
	}
	return &typedValue{P: node.Pos, T: &VoidType{}}
}

func (semantic *semanticPass) extractIntrinsicType(value abstractValue) (IntrinsicType, bool) {
	switch value := value.(type) {
	case *poisonValue:
		return INVALID_INTRINSIC_TYPE, false
	default:
		switch t := value.Type().(type) {
		case *IntrinsicTypeX:
			return t.UID, true
		default:
			semantic.status.LocationError(value.Pos(), "expected intrinsic type")
			return INVALID_INTRINSIC_TYPE, false
		}
	}
}

func (semantic *semanticPass) checkValueMatchesType(v abstractValue, t Type) {
	switch v := v.(type) {
	case *poisonValue:
		// Don't check.
	default:
		vt := v.Type()
		if vt != t {
			semantic.status.LocationError(v.Pos(), fmt.Sprintf("expected %s value but got %s value instead", PrintableTypeName(t), PrintableTypeName(vt)))
		}
	}
}

func (semantic *semanticPass) opExpr(node Expr) abstractValue {
	switch node := node.(type) {
	case *IntegerLiteral:
		_, err := strconv.ParseInt(node.Text, 10, 32)
		if err != nil {
			semantic.status.LocationError(node.Pos, "invalid i32")
		}
		return &typedValue{P: node.Pos, T: semantic.intrinsicTypes[INTRINSIC_TYPE_I32]}
	case *GetName:
		name := node.Name.Text
		o, ok := semantic.locals.Get(name)
		if !ok {
			o, ok = semantic.namespace.Get(name)
			if !ok {
				semantic.status.LocationError(node.Name.Pos, fmt.Sprintf("cannot resolve name %#v", name))
				return poison
			}
		}

		switch o := o.(type) {
		case *LocalSlot:
			return &typedValue{P: node.Name.Pos, T: o.Type}
		case *Function:
			return &functionValue{P: node.Name.Pos, F: o}
		default:
			panic(o)
		}
	case *InfixExpr:
		l := semantic.opExpr(node.Left)
		r := semantic.opExpr(node.Right)

		lt, lok := semantic.extractIntrinsicType(l)
		rt, rok := semantic.extractIntrinsicType(r)
		if !lok || !rok {
			return poison
		}

		fullOverloads, ok := InfixToIntrinsic[node.Op]
		if ok {
			rightOverloads, ok := fullOverloads[lt]
			if ok {
				info, ok := rightOverloads[rt]
				if ok {
					return &typedValue{P: node.Pos, T: semantic.intrinsicTypes[info.InfixResult]}
				}
			}
		}

		semantic.status.LocationError(node.Pos, fmt.Sprintf("unknown infix op (%s %s %s)", PrintableTypeName(l.Type()), InfixOperatorInfos[node.Op].Text, PrintableTypeName(r.Type())))
		return poison
	case *Call:
		expr := semantic.opExpr(node.Expr)
		args := []abstractValue{}
		for _, arg := range node.Args {
			args = append(args, semantic.opExpr(arg))
		}
		switch expr := expr.(type) {
		case *functionValue:
			ft := expr.F.Type
			if len(ft.Params) == len(args) {
				for i := 0; i < len(ft.Params); i++ {
					semantic.checkValueMatchesType(args[i], ft.Params[i])
				}
			} else {
				semantic.status.LocationError(node.Pos, fmt.Sprintf("expected %d arguments, got %d", len(ft.Params), len(args)))
			}
			return &typedValue{P: node.Pos, T: ft.Result}
		case *poisonValue:
			return poison
		default:
			semantic.status.LocationError(node.Pos, "callee is not a function")
			return poison
		}
	case *If:
		cond := semantic.opExpr(node.Cond)
		semantic.checkValueMatchesType(cond, semantic.intrinsicTypes[INTRINSIC_TYPE_BOOL])
		semantic.opBlock(node.T)
		if node.F != nil {
			semantic.opBlock(node.F)
		}
		// TODO return value for if.
		return &typedValue{P: node.Pos, T: &VoidType{}}
	case *Block:
		return semantic.opBlock(node)
	case *Return:
		if node.Expr != nil {
			value := semantic.opExpr(node.Expr)
			semantic.checkValueMatchesType(value, semantic.returnType)
		} else {
			panic(node)
		}
		return &typedValue{P: node.Pos, T: &VoidType{}}
	default:
		panic(node)
	}
}

func (semantic *semanticPass) opDecl(node *FunctionDecl) {
	f := semantic.recoverFunction(node)
	locals := MakeNamespace()

	for i := 0; i < f.NumParams; i++ {
		slot := f.Locals[i]
		locals.Set(slot.Name, slot)
	}
	semantic.locals = locals
	semantic.returnType = f.Type.Result

	semantic.opBlock(node.Body)

	semantic.locals = nil
	semantic.returnType = nil
}

func (semantic *semanticPass) opFile(file *File) {
	for _, d := range file.Decls {
		semantic.opDecl(d)
	}
}

func (semantic *semanticPass) recoverFunction(node *FunctionDecl) *Function {
	o, ok := semantic.namespace.Get(node.Name.Text)
	if !ok {
		panic(node.Name.Text)
	}
	f, ok := o.(*Function)
	if !ok {
		panic(node.Name.Text)
	}
	return f
}

func SemanticPass(node *File, status framework.CompileStatus) {
	namespace := MakeNamespace()
	intrinsicTypes := make([]*IntrinsicTypeX, len(IntrinsicTypeInfos))

	for i, info := range IntrinsicTypeInfos {
		t := &IntrinsicTypeX{
			UID:  info.Enum,
			Info: info,
		}
		namespace.Set(info.Text, t)
		intrinsicTypes[i] = t
	}

	semantic := &semanticPass{
		namespace:      namespace,
		intrinsicTypes: intrinsicTypes,
		status:         status,
	}

	// Resolve names.
	semantic.indexModuleNamespace(node)
	if status.ErrorOccured() {
		return
	}

	// Resolve types outside of function bodies.
	semantic.typeFile(node)
	if status.ErrorOccured() {
		return
	}

	// Process function bodies.
	semantic.opFile(node)
}
