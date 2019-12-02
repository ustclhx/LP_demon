package graph

import(
	
)
type Node struct{
	name string
	observable bool
}

// return a point to a new Node, can be observed by default
func NewDefaultNode(s string)  *Node{
	return &Node{
		name:s ,
		observable:true,
	}
}

// return a point to a new Node
func NewNode(s string, b bool) *Node{
	return &Node{
		name:s,
		observable:b,
	}
}

func (n *Node) Setname(s string){
	n.name = s
}

func (n Node ) Getname() string{
	return n.name
}

func (n *Node) Setob(b bool){
	n.observable = b 
}

func (n Node) Isob() bool{
	return n.observable
}
