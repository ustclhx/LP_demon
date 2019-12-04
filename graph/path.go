package graph

import(
	"fmt"
)

type Path struct{
	from Node
	to Node
	nodes []Node
}

func NewEmptyPath(from Node,to Node) *Path{
	nodes := make([]Node,0)
	return &Path{
		from : from,
		to : to,
		nodes : nodes,
	}
}

func NewPath(from Node, nodes []Node, to Node) *Path{
	return &Path{
		from : from,
		to : to,
		nodes : nodes,
	}
}

func (p *Path)PushNode(n Node){
	p.nodes = append (p.nodes, n)
}

func (p *Path)PopNode() (n Node , ok bool){
	if len(p.nodes)== 0{
		ok = false
		return
	}
	n = p.nodes[len(p.nodes)-1]
	p.nodes=p.nodes[:len(p.nodes)-1]
	return n,true
}

func (p Path) String() string{
	var s string
	for i,n := range p.nodes{
		if i == 0{
			s = n.name
		}else{
			s = s + "-" + n.name
		}
	}
	return fmt.Sprintf(s)
}