package identify

import(
	"LP_demon/graph"
	"strconv"
)

//corresponding to Figure3.4 in 《Causality》
func back_example_1() (*graph.Dag,[]*graph.Node){
	nodes := make([]*graph.Node,0) 
	edges := make([]*graph.Edge,0)
	nodes = append(nodes,graph.NewDefaultNode("xi"))
	for i := 1; i<= 6; i++{
		nodes = append(nodes,graph.NewDefaultNode("x"+strconv.Itoa(i)))
	}
	nodes = append(nodes,graph.NewDefaultNode("xj"))
	edges = append(edges,graph.NewEdge(nodes[0],graph.Tail,nodes[6],graph.Arrow))
	edges = append(edges,graph.NewEdge(nodes[6],graph.Tail,nodes[7],graph.Arrow))
	edges = append(edges,graph.NewEdge(nodes[4],graph.Tail,nodes[0],graph.Arrow))
	edges = append(edges,graph.NewEdge(nodes[4],graph.Tail,nodes[7],graph.Arrow))
	edges = append(edges,graph.NewEdge(nodes[3],graph.Tail,nodes[0],graph.Arrow))
	edges = append(edges,graph.NewEdge(nodes[1],graph.Tail,nodes[3],graph.Arrow))
	edges = append(edges,graph.NewEdge(nodes[1],graph.Tail,nodes[4],graph.Arrow))
	edges = append(edges,graph.NewEdge(nodes[2],graph.Tail,nodes[4],graph.Arrow))
	edges = append(edges,graph.NewEdge(nodes[2],graph.Tail,nodes[5],graph.Arrow))
	edges = append(edges,graph.NewEdge(nodes[5],graph.Tail,nodes[7],graph.Arrow))
	d,_ := graph.NewDag(nodes,edges)
	return d,nodes
}

func  front_example_1() (*graph.Dag,[]*graph.Node){
	nodes := make([]*graph.Node,0)
	edges := make([]*graph.Edge,0)
	nodes = append(nodes,graph.NewNode("X",true),graph.NewNode("Z",true),
	graph.NewNode("U",false),graph.NewNode("Y",false))
	edges = append(edges,graph.NewEdge(nodes[0],graph.Tail,nodes[1],graph.Arrow))
	edges = append(edges,graph.NewEdge(nodes[1],graph.Tail,nodes[3],graph.Arrow))
	edges = append(edges,graph.NewEdge(nodes[2],graph.Tail,nodes[0],graph.Arrow))
	edges = append(edges,graph.NewEdge(nodes[2],graph.Tail,nodes[3],graph.Arrow))
	d,_ := graph.NewDag(nodes,edges)
	return d,nodes
}