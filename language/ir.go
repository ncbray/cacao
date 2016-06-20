package language

type Namespace struct {
	index map[string]Nameable
}

func (ns *Namespace) IsDefined(name string) bool {
	_, ok := ns.index[name]
	return ok
}

func (ns *Namespace) Set(name string, obj Nameable) {
	ns.index[name] = obj
}

func (ns *Namespace) Get(name string) (Nameable, bool) {
	obj, ok := ns.index[name]
	return obj, ok
}

func MakeNamespace() *Namespace {
	return &Namespace{
		index: map[string]Nameable{},
	}
}

type Nameable interface {
	isNameable()
}

type Type interface {
	Nameable
	isType()
}

type Function struct {
	Name      string
	Type      *FunctionType
	NumParams int
	Locals    []*LocalSlot
}

func (obj *Function) isNameable() {
}

// TODO better name
type IntrinsicTypeX struct {
	UID  IntrinsicType
	Info *IntrinsicTypeInfo
}

func (obj *IntrinsicTypeX) isNameable() {
}

func (obj *IntrinsicTypeX) isType() {
}

type FunctionType struct {
	Params []Type
	Result Type
}

func (obj *FunctionType) isNameable() {
}

func (obj *FunctionType) isType() {
}

type VoidType struct {
}

func (obj *VoidType) isNameable() {
}

func (obj *VoidType) isType() {
}

type LocalSlot struct {
	Name string
	Type Type
}

func (obj *LocalSlot) isNameable() {
}
