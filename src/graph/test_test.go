package graph

import(
	"testing"
	"fmt" 
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
	e:=NewDefaultEdge(A,B)
	fmt.Println(e)
	e=NewEdge(A,Arrow,B,Tail)
	fmt.Println(e)
	B.Setob(true)
	fmt.Println(A.Isob(),B.Isob(),e.Isob())
	fmt.Println(e)
	e=NewEdge(A,Undefined,B,Arrow)
	fmt.Println(e)
}