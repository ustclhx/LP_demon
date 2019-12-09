package identify

import(
	"LP_demon/graph"
	"testing"
	"fmt"
)

func TestBackpath(t *testing.T){
	d,nodes := back_example_1()
	backpath,desc := backpath_dag_o2o(d,*nodes[0],*nodes[7])
	for _,path := range backpath{
		fmt.Println(path)
		if nodetype,err := d.IdentifyPath(path); err != nil{
			fmt.Println(err.Error())
		}else{
			for s,ns := range nodetype{
				fmt.Printf("%s : ",s)
				for _,node := range ns{
					fmt.Printf("%s ",node.Getname())
				}
				fmt.Printf("\n")
			}
		}
	}
	for _,descnode := range desc{
		fmt.Println(descnode.Getname())
	}
}

func TestBackverify(t *testing.T){
	d,nodes := back_example_1()
	z := make([]graph.Node,0)
	z = append(z,*nodes[4])
	if backverify_dag_o2o(d,*nodes[0],*nodes[7],z){
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}
	z = append(z,*nodes[3])
	if backverify_dag_o2o(d,*nodes[0],*nodes[7],z){
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}
	z2 := make([]graph.Node,0)
	z = append(z,*nodes[6])
	if backverify_dag_o2o(d,*nodes[0],*nodes[7],z2){
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}
}