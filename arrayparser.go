// arrayparser project oddparser.go
package arrayparser

import (
	. "github.com/andyleap/parser"
)

type Node struct {
	Name     string  `json:"name"`
	Parent   *Node   `json:"-"` // Parent pointer - depth iteration
	Children []*Node `json:"children,omitempty"`
}

func (n Node) String() string {
	if n.Name != "" {
		return n.Name
	}
	ret := "["
	sep := ""
	for _, child := range n.Children {
		ret += sep + child.String()
		sep = ","
	}
	return ret + "]"
}

var g = grammar()

func grammar() *Grammar {
	ws := Ignore(Mult(0, 0, Set("\\s")))

	elem := &Grammar{}

	value := Mult(1, 0, Set("^\\s,\\[\\]"))
	value.Node(func(m Match) (Match, error) {
		return &Node{Name: String(m)}, nil
	})

	array := And(
		Lit("["),
		Require(Tag("elem", elem),
			Mult(0, 0, And(Ignore(Lit(",")), Tag("elem", elem))),
			Lit("]"),
		),
	)

	array.Node(func(m Match) (Match, error) {
		elems := GetTags(m, "elem")
		node := &Node{
			Children: make([]*Node, 0, len(elems)),
		}
		for _, v := range elems {
			n := v.(*Node)
			n.Parent = node
			node.Children = append(node.Children, n)
		}
		return node, nil
	})

	elem.Set(And(ws, Tag("val", Or(value, array)), ws))
	elem.Node(func(m Match) (Match, error) {
		return GetTag(m, "val"), nil
	})

	return elem
}

func Parse(str string) (*Node, error) {
	n, err := g.ParseString(str)
	if err != nil {
		return nil, err
	}
	return n.(*Node), nil
}
