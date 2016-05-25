package language

type Link struct {
	Node *Node
	Op   *Operation
	Next *Link
	Prev *Link
}

type Node struct {
	UID     uint32
	DefHead *Link
	DefTail *Link
	UseHead *Link
	UseTail *Link
}

type Edges struct {
	Control []*Link
	Effect  []*Link
	Value   []*Link
}

type Operation struct {
	UID OperationID
	In  Edges
	Out Edges
}

type OperationID uint32
type ControlID uint32
type EffectID uint32
type ValueID uint32

type OperationConfig struct {
	ControlIn  int
	EffectIn   int
	ValueIn    int
	ControlOut int
	EffectOut  int
	ValueOut   int
}

type OperationIterator struct {
	current *Link
}

func (iter *OperationIterator) HasNext() bool {
	return iter.current != nil
}

func (iter *OperationIterator) GetNext() OperationID {
	uid := iter.current.Op.UID
	iter.current = iter.current.Next
	return uid
}

type Graph struct {
	ops     []*Operation
	control []*Node
	effect  []*Node
	value   []*Node
}

func CreateGraph() *Graph {
	return &Graph{}
}

func (g *Graph) CreateOperation(config OperationConfig) OperationID {
	op := createOperation(
		config.ControlIn,
		config.EffectIn,
		config.ValueIn,
		config.ControlOut,
		config.EffectOut,
		config.ValueOut)
	uid := OperationID(len(g.ops))
	op.UID = uid
	g.ops = append(g.ops, op)
	return uid
}

func (g *Graph) CreateControl() ControlID {
	node := &Node{}
	uid := uint32(len(g.control))
	node.UID = uid
	g.control = append(g.control, node)
	return ControlID(uid)
}

func (g *Graph) CreateEffect() EffectID {
	node := &Node{}
	uid := uint32(len(g.effect))
	node.UID = uid
	g.effect = append(g.effect, node)
	return EffectID(uid)
}

func (g *Graph) CreateValue() ValueID {
	node := &Node{}
	uid := uint32(len(g.value))
	node.UID = uid
	g.value = append(g.value, node)
	return ValueID(uid)
}

func (g *Graph) LinkControlUse(nid ControlID, oid OperationID, i int) {
	n := g.control[nid]
	o := g.ops[oid]
	linkControlUse(n, o, i)
}

func (g *Graph) LinkControlDef(oid OperationID, i int, nid ControlID) {
	o := g.ops[oid]
	n := g.control[nid]
	linkControlDef(o, i, n)
}

func (g *Graph) GetControlUse(oid OperationID, i int) (ControlID, bool) {
	o := g.ops[oid]
	l := o.In.Control[i]
	if l.Node != nil {
		return ControlID(l.Node.UID), true
	} else {
		return 0, false
	}
}

func (g *Graph) GetControlDef(oid OperationID, i int) (ControlID, bool) {
	o := g.ops[oid]
	l := o.Out.Control[i]
	if l.Node != nil {
		return ControlID(l.Node.UID), true
	} else {
		return 0, false
	}
}

func (g *Graph) NumControlUses(oid OperationID) int {
	return len(g.ops[oid].In.Control)
}

func (g *Graph) NumControlDefs(oid OperationID) int {
	return len(g.ops[oid].Out.Control)
}

func (g *Graph) IterControlDefs(cid ControlID) OperationIterator {
	n := g.control[cid]
	return OperationIterator{current: n.DefHead}
}

func (g *Graph) IterControlUses(cid ControlID) OperationIterator {
	n := g.control[cid]
	return OperationIterator{current: n.UseHead}
}

func (g *Graph) LinkEffectUse(nid EffectID, oid OperationID, i int) {
	n := g.effect[nid]
	o := g.ops[oid]
	linkEffectUse(n, o, i)
}

func (g *Graph) LinkEffectDef(oid OperationID, i int, nid EffectID) {
	o := g.ops[oid]
	n := g.effect[nid]
	linkEffectDef(o, i, n)
}

func (g *Graph) GetEffectUse(oid OperationID, i int) (EffectID, bool) {
	o := g.ops[oid]
	l := o.In.Effect[i]
	if l.Node != nil {
		return EffectID(l.Node.UID), true
	} else {
		return 0, false
	}
}

func (g *Graph) GetEffectDef(oid OperationID, i int) (EffectID, bool) {
	o := g.ops[oid]
	l := o.Out.Effect[i]
	if l.Node != nil {
		return EffectID(l.Node.UID), true
	} else {
		return 0, false
	}
}

func (g *Graph) NumEffectUses(oid OperationID) int {
	return len(g.ops[oid].In.Effect)
}

func (g *Graph) NumEffectDefs(oid OperationID) int {
	return len(g.ops[oid].Out.Effect)
}

func (g *Graph) IterEffectDefs(cid EffectID) OperationIterator {
	n := g.effect[cid]
	return OperationIterator{current: n.DefHead}
}

func (g *Graph) IterEffectUses(cid EffectID) OperationIterator {
	n := g.effect[cid]
	return OperationIterator{current: n.UseHead}
}

func (g *Graph) LinkValueUse(nid ValueID, oid OperationID, i int) {
	n := g.value[nid]
	o := g.ops[oid]
	linkValueUse(n, o, i)
}

func (g *Graph) LinkValueDef(oid OperationID, i int, nid ValueID) {
	o := g.ops[oid]
	n := g.value[nid]
	linkValueDef(o, i, n)
}

func (g *Graph) GetValueUse(oid OperationID, i int) (ValueID, bool) {
	o := g.ops[oid]
	l := o.In.Value[i]
	if l.Node != nil {
		return ValueID(l.Node.UID), true
	} else {
		return 0, false
	}
}

func (g *Graph) GetValueDef(oid OperationID, i int) (ValueID, bool) {
	o := g.ops[oid]
	l := o.Out.Value[i]
	if l.Node != nil {
		return ValueID(l.Node.UID), true
	} else {
		return 0, false
	}
}

func (g *Graph) NumValueUses(oid OperationID) int {
	return len(g.ops[oid].In.Value)
}

func (g *Graph) NumValueDefs(oid OperationID) int {
	return len(g.ops[oid].Out.Value)
}

func (g *Graph) IterValueDefs(cid ValueID) OperationIterator {
	n := g.value[cid]
	return OperationIterator{current: n.DefHead}
}

func (g *Graph) IterValueUses(cid ValueID) OperationIterator {
	n := g.value[cid]
	return OperationIterator{current: n.UseHead}
}

func createLinks(op *Operation, count int) []*Link {
	links := make([]*Link, count)
	for i := 0; i < count; i++ {
		links[i] = &Link{Op: op}
	}
	return links
}

func createEdges(op *Operation, c int, e int, v int) Edges {
	return Edges{
		Control: createLinks(op, c),
		Effect:  createLinks(op, e),
		Value:   createLinks(op, v),
	}
}

func createOperation(cin int, ein int, vin int, cout int, eout int, vout int) *Operation {
	op := &Operation{}
	op.In = createEdges(op, cin, ein, vin)
	op.Out = createEdges(op, cout, eout, vout)
	return op
}

func linkUse(n *Node, links []*Link, i int) {
	l := links[i]
	if l.Node != nil {
		p := l.Prev
		n := l.Next
		if p != nil {
			p.Next = n
		} else {
			l.Node.UseHead = n
		}
		if n != nil {
			n.Prev = p
		} else {
			l.Node.UseTail = p
		}
		l.Next = nil
		l.Prev = nil
	}
	l.Node = n
	if n.UseTail == nil {
		n.UseHead = l
		n.UseTail = l
	} else {
		t := n.UseTail
		l.Prev = t
		t.Next = l
		n.UseTail = l
	}
}

func linkDef(links []*Link, i int, n *Node) {
	l := links[i]
	if l.Node != nil {
		p := l.Prev
		n := l.Next
		if p != nil {
			p.Next = n
		} else {
			l.Node.DefHead = n
		}
		if n != nil {
			n.Prev = p
		} else {
			l.Node.DefTail = p
		}
		l.Next = nil
		l.Prev = nil
	}
	l.Node = n
	if n.DefTail == nil {
		n.DefHead = l
		n.DefTail = l
	} else {
		t := n.DefTail
		l.Prev = t
		t.Next = l
		n.DefTail = l
	}
}

func linkControlUse(v *Node, o *Operation, i int) {
	linkUse(v, o.In.Control, i)
}

func linkControlDef(o *Operation, i int, v *Node) {
	linkDef(o.Out.Control, i, v)
}

func linkEffectUse(v *Node, o *Operation, i int) {
	linkUse(v, o.In.Effect, i)
}

func linkEffectDef(o *Operation, i int, v *Node) {
	linkDef(o.Out.Effect, i, v)
}

func linkValueUse(v *Node, o *Operation, i int) {
	linkUse(v, o.In.Value, i)
}

func linkValueDef(o *Operation, i int, v *Node) {
	linkDef(o.Out.Value, i, v)
}
