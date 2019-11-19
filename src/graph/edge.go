package graph

import(
	"fmt"
)

const(
	Undefined = 0
	Arrow = 1
	Tail = 2
)

type Edge struct{
	nodes [2]*Node
	endpoints [2]int
}

//return a point to a new edge, both the endpoints are not defined
func NewDefaultEdge(A *Node, B *Node) *Edge{
	return &Edge{
		nodes : [2]*Node{A,B},
		endpoints : [2]int{Undefined,Undefined},
	}
}

//return a point to a new edge
func NewEdge(A *Node,end1 int, B *Node, end2 int) *Edge{
	return &Edge{
		nodes : [2]*Node{A,B},
		endpoints : [2]int{end1,end2},
	}
}

func (e Edge) Isob() bool{
	return e.nodes[0].observable && e.nodes[1].observable
}

func (e Edge) String() string{
	var s [2]string
	var ob string
	for i,v:=range e.endpoints{
		if v == Undefined{
			s[i]="o"
		}else if v == Tail{
			s[i]=""
		}else{
			if i==0{
				s[i]="<"
			}else{
				s[i]=">"
			}
		}
	}
	if e.nodes[0].observable && e.nodes[1].observable{
		ob = "—"
	}else{
		ob = "--"
	}
	return fmt.Sprintf("%v %v %v %v %v",e.nodes[0].name,s[0],ob,s[1],e.nodes[1].name)
}