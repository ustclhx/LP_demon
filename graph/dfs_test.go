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
	if d,err := NewDag(nodes,edges);err != nil{
		fmt.Println(err.Error())
	}else{
		paths := d.DFSpath(*nodes[6],*nodes[0],true)
		for _,p:= range paths{
			fmt.Println(p)
		}
		paths = d.DFSpath(*nodes[3],*nodes[6],false)
		for _,pi:= range paths{
			fmt.Println(pi)
		}
		desc := d.AllDescendant(*nodes[8])
		for _,n:= range desc{
			fmt.Println(n.name)
		}
	}

     
}
