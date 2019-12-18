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

func TestFrontsearch(t *testing.T){
	d,nodes := front_example_2()
	fmt.Println("search:")
	status, z := Frontsearch_dag(d,[]graph.Node{*nodes[0]},[]graph.Node{*nodes[7]})
	if status{
		for _,n := range z{
			fmt.Println(n.Getname())
		}
	}else{
		fmt.Println("no invalid frontdoor")
	}
	fmt.Println("all:")
	status, zs := Frontallsearch_dag(d,[]graph.Node{*nodes[0]},[]graph.Node{*nodes[7]})
	if status{
		for _,ns := range zs{
			for _,n := range ns{
				fmt.Printf("%s ",n.Getname())
			}
			fmt.Printf("\n")
		}
	}else{
		fmt.Println("no invalid frontdoor")
	}
	fmt.Println("minimal:")
	status, zs = Frontminimal_dag(d,[]graph.Node{*nodes[0]},[]graph.Node{*nodes[7]})
	if status{
		for _,ns := range zs{
			for _,n := range ns{
				fmt.Printf("%s ",n.Getname())
			}
			fmt.Printf("\n")
		}
	}else{
		fmt.Println("no invalid frontdoor")
	}
}

func TestFrontsearch_2(t *testing.T){
	d,nodes := front_example_3()
	fmt.Println("search:")
	status, z := Frontsearch_dag(d,[]graph.Node{*nodes[0]},[]graph.Node{*nodes[7]})
	if status{
		for _,n := range z{
			fmt.Println(n.Getname())
		}
	}else{
		fmt.Println("no invalid frontdoor")
	}
	fmt.Println("all:")
	status, zs := Frontallsearch_dag(d,[]graph.Node{*nodes[0]},[]graph.Node{*nodes[7]})
	if status{
		for _,ns := range zs{
			for _,n := range ns{
				fmt.Printf("%s ",n.Getname())
			}
			fmt.Printf("\n")
		}
	}else{
		fmt.Println("no invalid frontdoor")
	}
	fmt.Println("minimal:")
	status, zs = Frontminimal_dag(d,[]graph.Node{*nodes[0]},[]graph.Node{*nodes[7]})
	if status{
		for _,ns := range zs{
			for _,n := range ns{
				fmt.Printf("%s ",n.Getname())
			}
			fmt.Printf("\n")
		}
	}else{
		fmt.Println("no invalid frontdoor")
	}
}