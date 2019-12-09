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
	endpoints map[Node]int
}

//return a point to a new edge, both the endpoints are not defined
func NewDefaultEdge(A *Node, B *Node) *Edge{	
	return &Edge{
		nodes : [2]*Node{A,B},
		endpoints : map[Node]int{
			*A : Undefined,
			*B : Undefined,
		},
	}
}

//return a point to a new edge
func NewEdge(A *Node,end1 int, B *Node, end2 int) *Edge{
	return &Edge{
		nodes : [2]*Node{A,B},
		endpoints : map[Node]int{
			*A : end1,
			*B : end2,
		},
	}
}

func (e Edge) Isob() bool{
	return e.nodes[0].observable && e.nodes[1].observable
}

func (e Edge) IsDirected()bool{
	end_0 := e.endpoints[*e.nodes[0]]
	end_1 := e.endpoints[*e.nodes[1]]
	return (end_0 == Arrow && end_1 == Tail)||(end_0 == Tail && end_1 == Arrow)
}

func (e Edge) Endpoint(N Node) int{
	return e.endpoints[N]
}

func (e Edge) String() string{
	var s [2]string
	var ob string
	for i,n:=range e.nodes{
		v := e.endpoints[*n]
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