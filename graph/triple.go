package graph

import(
	"fmt"
)

const(
	Chain = 1
	Fork = 2
	Collider = 3
)

type Triple struct{
	nodes [3]*Node
}

func NewTriple(A *Node,B *Node,C *Node) *Triple{
	return &Triple{
		nodes : [3]*Node{A,B,C},
	}
}

func (g Graph) IsTriplein(t Triple) bool{
	ok_1 := g.IsAdjacentto(t.nodes[0],t.nodes[1])
	ok_2 := g.IsAdjacentto(t.nodes[1],t.nodes[2])
	return ok_1 && ok_2 
}

func (t Triple) String() string{
	return fmt.Sprintf("{%v %v %v}",t.nodes[0].name,t.nodes[1].name,t.nodes[2].name)
}

func (g *Graph) Identify(t *Triple) (i int, s string, err error){
	s = "undefined"
	if !g.IsTriplein(*t){
		err = fmt.Errorf("the triple %s does't in the graph",t)
		return
	}
	edge_1 := g.edges[*t.nodes[0]][*t.nodes[1]]
	edge_2 := g.edges[*t.nodes[1]][*t.nodes[2]]
	if edge_1.endpoints[*t.nodes[0]] == Tail && edge_1.endpoints[*t.nodes[1]] == Arrow{
		if edge_2.endpoints[*t.nodes[1]] == Tail && edge_2.endpoints[*t.nodes[2]] == Arrow{
			return Chain,"chain",nil
		}
		if edge_2.endpoints[*t.nodes[1]] == Arrow && edge_2.endpoints[*t.nodes[2]] == Tail{
			return Collider,"collider",nil
		}
	}
	if edge_1.endpoints[*t.nodes[0]] == Arrow && edge_1.endpoints[*t.nodes[1]] == Tail{
		if edge_2.endpoints[*t.nodes[1]] == Tail && edge_2.endpoints[*t.nodes[2]] == Arrow{
			return Fork,"fork",nil
		}
		if edge_2.endpoints[*t.nodes[1]] == Arrow && edge_2.endpoints[*t.nodes[2]] == Tail{
			return Chain,"chain",nil
		}
	}
	return
}

