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
	nodes [2]Node
	endpoints [2]int
}

func NewDefaultEdge(A Node, B Node) Edge{
	return Edge{
		nodes : [2]Node{A,B},
		endpoints : [2]int{Undefined,Undefined},
	}
}

func NewEdge(A Node, B Node, end1 int, end2 int) Edge{
	return Edge{
		nodes : [2]Node{A,B},
		endpoints : [2]int{end1,end2},
	}
}

func (e Edge) String() string{
	var s [2]string
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
	return fmt.Sprintf("%v %v-%v %v",e.nodes[0].name,s[0],s[1],e.nodes[1].name)
}