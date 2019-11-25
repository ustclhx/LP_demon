package graph

import(
	"fmt"
	//"errors"
)

type Graph struct{
	nodes []*Node 
	edges map[Node]map[Node]*Edge
}

func (g Graph) IsNodeIn(n Node) bool{
	_,ok := g.edges[n]
	return ok
}

func NewGraph(ns []*Node, es []*Edge) (g *Graph,err error){
	edges := make(map[Node]map[Node]*Edge)
	for  _, node := range ns{
		edges[*node]= make(map[Node]*Edge)
	}
	for _, e := range es{
		_,ok_0 := edges[*e.nodes[0]]
		_,ok_1 := edges[*e.nodes[1]]
		if !(ok_0 && ok_1){
			err =  fmt.Errorf("the edge between node %s and node %s has a node not in the graph",e.nodes[0].name,e.nodes[1].name)
			return 
		}
		edges[*e.nodes[0]][*e.nodes[1]]= e
		edges[*e.nodes[1]][*e.nodes[0]]= e	
	}
	g = &Graph{
		nodes : ns,
		edges : edges,
	}
	return 
}

func (g *Graph) AddNode(n *Node){
	g.nodes = append(g.nodes,n)
}

/*func (g *Graph) AddEdge(e *Edge)(err error){
	if g,edges == nil{
		g.edges = make(map[Node]map[Node]*Edge)
	}
}*/

