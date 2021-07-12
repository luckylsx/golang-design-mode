package prototypepattern

import "testing"

var manager *PrototypeManager

type type1 struct {
	name string
}

func (t *type1) Clone() Cloneable {
	tc1 := *t
	return &tc1
}

type type2 struct {
	name string
}

func init() {
	manager = NewPrototypeManager()
	t1 := &type1{
		name: "type1",
	}
	manager.Set("t1", t1)
}

func TestClone(t *testing.T) {
	t1 := manager.Get("t1")
	t2 := t1.Clone()
	if t1 == t2 {
		t.Errorf("error! clone not working, t1: memory addr(%p), t2: memory addr(%p)", t1, t2)
	}
}

func TestCloneFromManager(t *testing.T) {
	c := manager.Get("t1").Clone()
	if t1, ok := c.(*type1); !ok || t1.name != "type1" {
		t.Fatal("error")
	}
}
