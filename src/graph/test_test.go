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
	testnum := 3
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
	errnode :=NewDefaultNode(strconv.Itoa(testnum+1))
	edges = append(edges,NewEdge(nodes[0],Arrow,errnode,Tail))
	_,err := NewGraph(nodes,edges)
    if err ==nil{
		fmt.Println("OK")
	}else{
		fmt.Println(err.Error())
	}

}