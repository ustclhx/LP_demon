package identify

import(
	"LP_demon/graph"
	"testing"
	"fmt"
)

func TestFrontverify(t *testing.T){
	d,nodes := front_example_1()
	if Frontverify_dag(d,[]graph.Node{*nodes[0]},[]graph.Node{*nodes[3]},[]graph.Node{*nodes[1]}){
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}
	if Frontverify_dag(d,[]graph.Node{*nodes[0]},[]graph.Node{*nodes[3]},[]graph.Node{*nodes[2]}){
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}
	if Backverify_dag(d,[]graph.Node{*nodes[0]},[]graph.Node{*nodes[3]},[]graph.Node{*nodes[2]}){
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}
	d.AddEdge(graph.NewEdge(nodes[0],graph.Tail,nodes[3],graph.Arrow))
	if Frontverify_dag(d,[]graph.Node{*nodes[0]},[]graph.Node{*nodes[3]},[]graph.Node{*nodes[1]}){
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}
}