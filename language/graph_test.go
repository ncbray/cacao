package language

import (
	"fmt"
	"testing"
)

func createNode() *Node {
	return &Node{}
}

/*
func checkLinks(name string, op OperationID, links []*Link, nodes []*Node, t *testing.T) {
	if len(nodes) != len(links) {
		t.Errorf("expected %d %s nodes, got %d", len(nodes), name, len(links))
	} else {
		for i := 0; i < len(nodes); i++ {
			full := fmt.Sprintf("%s.%d", name, i)
			l := links[i]
			if l.Op != op {
				t.Errorf("%s: link does not point to operation", full)
			}
			n := nodes[i]
			if l.Node != n {
				t.Errorf("%s: expected %v, got %v", full, n, l.Node)
			}
		}
	}
}
*/

func checkOp(g *Graph, name string, op OperationID, cin []ControlID, ein []EffectID, vin []ValueID, cout []ControlID, eout []EffectID, vout []ValueID, t *testing.T) {
	if len(cin) == g.NumControlUses(op) {
		for i, expected := range cin {
			full := fmt.Sprintf("%s.use.%d", name, i)
			actual, ok := g.GetControlUse(op, i)
			if ok {
				if actual != expected {
					t.Errorf("%s: expected %d, got %d", full, expected, actual)
				}
			} else {
				t.Errorf("%s: not linked", full)
			}
		}
	} else {
		t.Errorf("%s: expected %d control uses, got %d", name, len(cin), g.NumControlUses(op))
	}

	if len(cout) == g.NumControlDefs(op) {
		for i, expected := range cout {
			full := fmt.Sprintf("%s.def.%d", name, i)
			actual, ok := g.GetControlDef(op, i)
			if ok {
				if actual != expected {
					t.Errorf("%s: expected %d, got %d", full, expected, actual)
				}
			} else {
				t.Errorf("%s: not linked", full)
			}
		}
	} else {
		t.Errorf("%s: expected %d control defs, got %d", name, len(cin), g.NumControlDefs(op))
	}

	if len(ein) == g.NumEffectUses(op) {
		for i, expected := range ein {
			full := fmt.Sprintf("%s.use.%d", name, i)
			actual, ok := g.GetEffectUse(op, i)
			if ok {
				if actual != expected {
					t.Errorf("%s: expected %d, got %d", full, expected, actual)
				}
			} else {
				t.Errorf("%s: not linked", full)
			}
		}
	} else {
		t.Errorf("%s: expected %d effect uses, got %d", name, len(cin), g.NumEffectUses(op))
	}

	if len(eout) == g.NumEffectDefs(op) {
		for i, expected := range eout {
			full := fmt.Sprintf("%s.def.%d", name, i)
			actual, ok := g.GetEffectDef(op, i)
			if ok {
				if actual != expected {
					t.Errorf("%s: expected %d, got %d", full, expected, actual)
				}
			} else {
				t.Errorf("%s: not linked", full)
			}
		}
	} else {
		t.Errorf("%s: expected %d effect defs, got %d", name, len(cin), g.NumEffectDefs(op))
	}

	if len(vin) == g.NumValueUses(op) {
		for i, expected := range vin {
			full := fmt.Sprintf("%s.use.%d", name, i)
			actual, ok := g.GetValueUse(op, i)
			if ok {
				if actual != expected {
					t.Errorf("%s: expected %d, got %d", full, expected, actual)
				}
			} else {
				t.Errorf("%s: not linked", full)
			}
		}
	} else {
		t.Errorf("%s: expected %d value uses, got %d", name, len(cin), g.NumValueUses(op))
	}

	if len(vout) == g.NumValueDefs(op) {
		for i, expected := range vout {
			full := fmt.Sprintf("%s.def.%d", name, i)
			actual, ok := g.GetValueDef(op, i)
			if ok {
				if actual != expected {
					t.Errorf("%s: expected %d, got %d", full, expected, actual)
				}
			} else {
				t.Errorf("%s: not linked", full)
			}
		}
	} else {
		t.Errorf("%s: expected %d value defs, got %d", name, len(cin), g.NumValueDefs(op))
	}

}

func checkIterator(g *Graph, name string, iter OperationIterator, ops []OperationID, t *testing.T) {
	i := 0
	for iter.HasNext() {
		full := fmt.Sprintf("%s.%d", name, i)
		if i >= len(ops) {
			t.Errorf("%s: iterator is longer than expected", full)
			break
		}

		actual := iter.GetNext()
		expected := ops[i]
		i += 1

		if expected != actual {
			t.Errorf("%s: expected %v, got %v", full, expected, actual)
		}
	}
	if i < len(ops) {
		t.Errorf("%s: iterator is shorter than expected", name)
	}
}

func checkControl(g *Graph, name string, node ControlID, defs []OperationID, uses []OperationID, t *testing.T) {
	checkIterator(g, name+".def", g.IterControlDefs(node), defs, t)
	checkIterator(g, name+".use", g.IterControlUses(node), uses, t)
}

func checkEffect(g *Graph, name string, node EffectID, defs []OperationID, uses []OperationID, t *testing.T) {
	checkIterator(g, name+".def", g.IterEffectDefs(node), defs, t)
	checkIterator(g, name+".use", g.IterEffectUses(node), uses, t)
}

func checkValue(g *Graph, name string, node ValueID, defs []OperationID, uses []OperationID, t *testing.T) {
	checkIterator(g, name+".def", g.IterValueDefs(node), defs, t)
	checkIterator(g, name+".use", g.IterValueUses(node), uses, t)
}

func TestOpPassthrough(t *testing.T) {
	g := CreateGraph()

	op := g.CreateOperation(OperationConfig{1, 1, 1, 1, 1, 1})
	c0 := g.CreateControl()
	c1 := g.CreateControl()
	e0 := g.CreateEffect()
	e1 := g.CreateEffect()
	v0 := g.CreateValue()
	v1 := g.CreateValue()

	g.LinkControlUse(c0, op, 0)
	g.LinkControlDef(op, 0, c1)

	g.LinkEffectUse(e0, op, 0)
	g.LinkEffectDef(op, 0, e1)

	g.LinkValueUse(v0, op, 0)
	g.LinkValueDef(op, 0, v1)

	checkOp(g, "op", op, []ControlID{c0}, []EffectID{e0}, []ValueID{v0}, []ControlID{c1}, []EffectID{e1}, []ValueID{v1}, t)

	checkControl(g, "c0", c0, nil, []OperationID{op}, t)
	checkControl(g, "c1", c1, []OperationID{op}, nil, t)
	checkEffect(g, "e0", e0, nil, []OperationID{op}, t)
	checkEffect(g, "e1", e1, []OperationID{op}, nil, t)
	checkValue(g, "v0", v0, nil, []OperationID{op}, t)
	checkValue(g, "v1", v1, []OperationID{op}, nil, t)
}

func TestOpControlReplaceHead(t *testing.T) {
	g := CreateGraph()

	op := g.CreateOperation(OperationConfig{2, 0, 0, 2, 0, 0})
	c0 := g.CreateControl()
	c1 := g.CreateControl()
	c2 := g.CreateControl()

	c3 := g.CreateControl()
	c4 := g.CreateControl()
	c5 := g.CreateControl()

	g.LinkControlUse(c0, op, 0)
	g.LinkControlUse(c1, op, 1)
	g.LinkControlUse(c2, op, 0)

	g.LinkControlDef(op, 0, c3)
	g.LinkControlDef(op, 1, c4)
	g.LinkControlDef(op, 0, c5)

	checkOp(g, "op", op, []ControlID{c2, c1}, nil, nil, []ControlID{c5, c4}, nil, nil, t)

	checkControl(g, "c0", c0, nil, nil, t)
	checkControl(g, "c1", c1, nil, []OperationID{op}, t)
	checkControl(g, "c2", c2, nil, []OperationID{op}, t)

	checkControl(g, "c3", c3, nil, nil, t)
	checkControl(g, "c4", c4, []OperationID{op}, nil, t)
	checkControl(g, "c5", c5, []OperationID{op}, nil, t)
}

func TestOpControlReplaceTail(t *testing.T) {
	g := CreateGraph()

	op := g.CreateOperation(OperationConfig{2, 0, 0, 2, 0, 0})
	c0 := g.CreateControl()
	c1 := g.CreateControl()
	c2 := g.CreateControl()

	c3 := g.CreateControl()
	c4 := g.CreateControl()
	c5 := g.CreateControl()

	g.LinkControlUse(c0, op, 0)
	g.LinkControlUse(c1, op, 1)
	g.LinkControlUse(c2, op, 1)

	g.LinkControlDef(op, 0, c3)
	g.LinkControlDef(op, 1, c4)
	g.LinkControlDef(op, 1, c5)

	checkOp(g, "op", op, []ControlID{c0, c2}, nil, nil, []ControlID{c3, c5}, nil, nil, t)

	checkControl(g, "c0", c0, nil, []OperationID{op}, t)
	checkControl(g, "c1", c1, nil, nil, t)
	checkControl(g, "c2", c2, nil, []OperationID{op}, t)

	checkControl(g, "c3", c3, []OperationID{op}, nil, t)
	checkControl(g, "c4", c4, nil, nil, t)
	checkControl(g, "c5", c5, []OperationID{op}, nil, t)
}

func TestOpControl3(t *testing.T) {
	g := CreateGraph()

	op := g.CreateOperation(OperationConfig{3, 0, 0, 3, 0, 0})
	c0 := g.CreateControl()
	c1 := g.CreateControl()
	c2 := g.CreateControl()

	c3 := g.CreateControl()
	c4 := g.CreateControl()
	c5 := g.CreateControl()

	g.LinkControlUse(c0, op, 0)
	g.LinkControlUse(c1, op, 1)
	g.LinkControlUse(c2, op, 2)

	g.LinkControlDef(op, 0, c3)
	g.LinkControlDef(op, 1, c4)
	g.LinkControlDef(op, 2, c5)

	checkOp(g, "op", op, []ControlID{c0, c1, c2}, nil, nil, []ControlID{c3, c4, c5}, nil, nil, t)

	checkControl(g, "c0", c0, nil, []OperationID{op}, t)
	checkControl(g, "c1", c1, nil, []OperationID{op}, t)
	checkControl(g, "c2", c2, nil, []OperationID{op}, t)

	checkControl(g, "c3", c3, []OperationID{op}, nil, t)
	checkControl(g, "c4", c4, []OperationID{op}, nil, t)
	checkControl(g, "c5", c5, []OperationID{op}, nil, t)
}

func TestNodeControl3(t *testing.T) {
	g := CreateGraph()

	c := g.CreateControl()

	op0 := g.CreateOperation(OperationConfig{0, 0, 0, 1, 0, 0})
	op1 := g.CreateOperation(OperationConfig{0, 0, 0, 1, 0, 0})
	op2 := g.CreateOperation(OperationConfig{0, 0, 0, 1, 0, 0})

	op3 := g.CreateOperation(OperationConfig{1, 0, 0, 0, 0, 0})
	op4 := g.CreateOperation(OperationConfig{1, 0, 0, 0, 0, 0})
	op5 := g.CreateOperation(OperationConfig{1, 0, 0, 0, 0, 0})

	g.LinkControlDef(op0, 0, c)
	g.LinkControlDef(op1, 0, c)
	g.LinkControlDef(op2, 0, c)

	g.LinkControlUse(c, op3, 0)
	g.LinkControlUse(c, op4, 0)
	g.LinkControlUse(c, op5, 0)

	checkControl(g, "c", c, []OperationID{op0, op1, op2}, []OperationID{op3, op4, op5}, t)

	checkOp(g, "op0", op0, nil, nil, nil, []ControlID{c}, nil, nil, t)
	checkOp(g, "op1", op1, nil, nil, nil, []ControlID{c}, nil, nil, t)
	checkOp(g, "op2", op2, nil, nil, nil, []ControlID{c}, nil, nil, t)
	checkOp(g, "op3", op3, []ControlID{c}, nil, nil, nil, nil, nil, t)
	checkOp(g, "op4", op4, []ControlID{c}, nil, nil, nil, nil, nil, t)
	checkOp(g, "op5", op5, []ControlID{c}, nil, nil, nil, nil, nil, t)
}
