package main

import(
	"LP_demon/graph"
	"fmt"
	"github.com/crillab/gophersat/solver"
)

func main(){
	n1:= graph.NewNode("1",true)
	n2:=graph.NewDefaultNode("2")
	e:= graph.NewEdge(n1,graph.Arrow,n2,graph.Tail)
	fmt.Println(e)
	clauses := [][]int{
		[]int{-6},
		[]int{4},
		[]int{2,4,5},
		[]int{1,3,4},
		[]int{1,2,3,5,-4},
	}
	pb := solver.ParseSlice(clauses)
	s := solver.New(pb)
	status := s.Solve()
	fmt.Println(status)
	s.OutputModel()
}