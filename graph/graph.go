package graph

import(
	"fmt"
	//"errors"
)

type Graph struct{
	nodes []*Node 
	edges map[Node]map[Node]*Edge
	node_index	map[string]*Node
	in_degree map[Node]int  //just consider the in-degree in directed edge
}

func (g Graph) IsNodeIn(n Node) bool{
	_,ok := g.edges[n]
	return ok
}

func (g Graph) IsEdgeIn(e Edge) bool{
	_,ok := g.edges[*e.nodes[0]][*e.nodes[1]]
	return ok
}

func (g Graph) Nodes() []*Node{
	return g.nodes
}

func (g Graph) IsAdjacentto(A Node, B Node) bool{
	_,ok := g.edges[A][B]
	return ok
}

func (g Graph)IsDirectto(A Node, B Node) bool{
	e,ok := g.edges[A][B]
	if !ok{
		return false
	}else{
		return e.endpoints[A] == Tail && e.endpoints[B]== Arrow
	}
	
}

func (g Graph) Edges(A Node) map[Node]*Edge{
	return g.edges[A]
}

func (g Graph) GetNode(s string) (*Node){
	n,_ := g.node_index[s]
	return n
}

func (g Graph) IsStringIn(s string) bool{
	_,b := g.node_index[s]
	return b 
}

func NewGraph(ns []*Node, es []*Edge) (g *Graph,err error){
	edges := make(map[Node]map[Node]*Edge)
	indegree := make(map[Node]int)
	node_index := make(map[string]*Node)
	for  _, node := range ns{
		edges[*node] = make(map[Node]*Edge)
		if _,ok := node_index[node.Getname()]; ok{
			err = fmt.Errorf("There are two nodes have same names")
			return 
		}else{
			node_index[node.Getname()] = node
		}
		indegree[*node] = 0
	}
	g = &Graph{
		nodes : ns,
		edges : edges,
		node_index : node_index,
		in_degree : indegree,
	}
	for _, e := range es{
		if err = g.AddEdge(e);err != nil{
			return 
		}
	}
	return 
}

func (g *Graph) AddNode(n *Node) (err error){
	if _,ok := g.node_index[n.Getname()]; ok{
		err = fmt.Errorf("There are two nodes have same names")
		return 
	}
	g.node_index[n.Getname()] = n
	g.nodes = append(g.nodes,n)
	g.edges[*n] = make(map[Node]*Edge)
	return nil
}

func (g *Graph) AddEdge(e *Edge)(err error){
	if e.nodes[0].name == e.nodes[1].name {
		err = fmt.Errorf("wrong edge form %s to itself",e.nodes[0].name)
	}
	for i:=0;i<=1;i++{
		ok := g.IsNodeIn(*e.nodes[i])
		if !ok{
			err = fmt.Errorf("the edge %s, has the node %s not in the graph",e,e.nodes[i].name)
			return			
		}
	}
	if ok:= g.IsEdgeIn(*e); ok{
			err = fmt.Errorf("There is an edge between node %s and node %s here, can't add a new one ", e.nodes[0].name, e.nodes[1].name)
			return
		}
	g.edges[*e.nodes[0]][*e.nodes[1]]= e
	g.edges[*e.nodes[1]][*e.nodes[0]]= e
	//increase the degree
	if e.IsDirected(){
		if e.endpoints[*e.nodes[0]] == Arrow{
			g.in_degree[*e.nodes[0]]++
		}else{
			g.in_degree[*e.nodes[1]]++
		}
	}
	return
}

func (g *Graph) AddDirectedEdge(e *Edge) (err error){
	if !e.IsDirected(){
		err = fmt.Errorf("the edge %s is not a directed edge", e)
	}
	err = g.AddEdge(e)
	return
}

func (g *Graph) RemoveEdge(e *Edge) (err error){
	if g.edges[*e.nodes[0]][*e.nodes[1]] == nil || g.edges[*e.nodes[1]][*e.nodes[0]]==nil{
		err = fmt.Errorf("the edge %s is not in the graph",e)
	}
	delete(g.edges[*e.nodes[0]],*e.nodes[1])
	delete(g.edges[*e.nodes[1]],*e.nodes[0])
	if e.IsDirected(){
		if e.endpoints[*e.nodes[0]] == Arrow{
			g.in_degree[*e.nodes[0]]--
		}else{
			g.in_degree[*e.nodes[1]]--
		}
	}
	return
}


func (g *Graph) Toposort() (sort []Node, ok bool){
	queue := make([]Node,0)
	sort = make([]Node,0)
	//change to pass value
	indegree := make(map[Node]int)  
	for k,v := range g.in_degree{
		indegree[k] = v
	}
	//find the node with indegree = 0
	for n,d := range indegree{
		if d == 0{
			queue = append(queue,n)
		}	
	}
	for len(queue)>0{
		n := queue[0]
		sort = append(sort,n)
		queue = queue[1:]
		for neigh,e := range g.edges[n]{
			if e. endpoints[neigh] == Arrow{
				indegree[neigh]--
				if indegree[neigh] == 0{
					queue = append(queue,neigh)
				}
			}
		}

	}
	if len(sort)<len(g.nodes){
		ok = false
		return
	}else{
		ok = true
		return
	}
}