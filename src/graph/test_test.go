package graph

import(
	"testing"
	"fmt" 
	"strconv"
)

func TestNode(t *testing.T){
	n:=NewDefaultNode("x")
	fmt.Println(n.Getname(),n.Isob())
	n.Setname("y")
	n.Setob(false)
	fmt.Println(n.Getname(),n.Isob())
}

func TestEdge(t *testing.T){
	A := NewDefaultNode("x")
	B := NewNode("y",false)
	e := NewDefaultEdge(A,B)
	fmt.Println(e)
	e = NewEdge(A,Arrow,B,Tail)
	fmt.Println(e)
	B.Setob(true)
	fmt.Println(A.Isob(),B.Isob(),e.Isob())
	fmt.Println(e)
	e=NewEdge(A,Undefined,B,Arrow)
	fmt.Println(e)
}

func TestNewGraph(t *testing.T){
	testnum := 4
	nodes := make([]*Node,0) 
	edges := make([]*Edge,0)
	for i :=0;i<testnum;i++{
		node := NewDefaultNode(strconv.Itoa(i))
		nodes = append(nodes,node)
	}
	for i :=0;i+1<testnum;i++{
		edge := NewEdge(nodes[i],Tail,nodes[i+1],Arrow)
		edges = append(edges,edge)
	}
	if g,err := NewGraph(nodes,edges);err != nil{
		fmt.Println(err.Error())
	}else{
		 sort,_ :=g.Toposort()
		 for _,n := range sort{
		 	fmt.Println(n.name)
		 }
		if _,ok := g.Toposort(); !ok{
			fmt.Println("has a cycle")
		}else{
			fmt.Println("has no cycle")
		}
		g.AddEdge(NewEdge(nodes[0],Tail,nodes[testnum-1],Arrow))
		// for _,m := range g.edges{
		// 	for _,e :=range m{
		// 		fmt.Println(e)
		// 	}
		// }
		for _,n := range g.nodes{
			fmt.Println(n.name,"'s indegree is",g.in_degree[*n])
		}
		g.AddEdge(NewEdge(nodes[testnum-1],Tail,nodes[testnum-2],Arrow))
		sort2,_ :=g.Toposort()
		 for _,n := range sort2{
		 	fmt.Println(n.name)
		 }
		if _,ok := g.Toposort(); !ok{
			fmt.Println("has a cycle")
		}else{
			fmt.Println("has no cycle")
		}
	}
	errnode :=NewDefaultNode(strconv.Itoa(testnum))
	edges = append(edges,NewEdge(nodes[0],Arrow,errnode,Tail))
	if _,err := NewGraph(nodes,edges);err !=nil{
		fmt.Println(err.Error())
	}
}