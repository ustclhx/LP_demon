package graph

import(
	"testing"
	"fmt" 
)

func TestNode(t *testing.T){
	var n Node
	n.Setname("x")
	str:=n.Getname()
	fmt.Println(str)
}

func TestEdge(t *testing.T){
	var A,B Node
	A.Setname("A")
	B.Setname("B")
	e:=NewDefaultEdge(A,B)
	fmt.Println(e)
	e=NewEdge(A,B,Arrow,Tail)
	fmt.Println(e)
	e=NewEdge(A,B,Undefined,Arrow)
	fmt.Println(e)
}