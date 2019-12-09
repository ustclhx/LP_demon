package main

import(
	"LP_demon/graph"
	"fmt"
)

func main(){
	n1:= graph.NewNode("1",true)
	n2:=graph.NewDefaultNode("2")
	e:= graph.NewEdge(n1,graph.Arrow,n2,graph.Tail)
	fmt.Println(e)
}