package graph

import(
	"testing"
	"fmt" 
	"strconv"
)
func TestDfs(t *testing.T){
	nodes := make([]*Node,0)
	for i :=0;i<=8;i++{
		node := NewDefaultNode(strconv.Itoa(i))
		nodes = append(nodes,node)
	}
	edges := make([]*Edge,0)
	edges = append(edges,NewEdge(nodes[0],Arrow,nodes[1],Tail))
	edges = append(edges,NewEdge(nodes[0],Arrow,nodes[2],Tail))
	edges = append(edges,NewEdge(nodes[3],Arrow,nodes[1],Tail))
	edges = append(edges,NewEdge(nodes[4],Arrow,nodes[1],Tail))
	edges = append(edges,NewEdge(nodes[2],Arrow,nodes[5],Tail))
	edges = append(edges,NewEdge(nodes[2],Arrow,nodes[6],Tail))
	edges = append(edges,NewEdge(nodes[3],Arrow,nodes[7],Tail))
	edges = append(edges,NewEdge(nodes[4],Arrow,nodes[5],Tail))
	edges = append(edges,NewEdge(nodes[4],Arrow,nodes[7],Tail))
	edges = append(edges,NewEdge(nodes[5],Arrow,nodes[6],Tail))
	edges = append(edges,NewEdge(nodes[6],Arrow,nodes[8],Tail))
	if g,err := NewGraph(nodes,edges);err != nil{
		fmt.Println(err.Error())
	}else{
		paths := g.DFSpath(*nodes[3],*nodes[6])
		for _,p:= range paths{
			fmt.Println(p)
		}
	}

     
}
