package identify

import(
	"testing"
	"fmt"
)

func TestBackpath(t *testing.T){
	d,nodes := back_example_1()
	backpath,desc := backpath_dag_o2o(d,*nodes[0],*nodes[7])
	for _,path := range backpath{
		fmt.Println(path)
	}
	for _,descnode := range desc{
		fmt.Println(descnode.Getname())
	}
}