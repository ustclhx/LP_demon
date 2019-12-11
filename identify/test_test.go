package identify

import(
	"LP_demon/graph"
	"testing"
	"fmt"
)

func TestBackpath(t *testing.T){
	d,nodes := back_example_1()
	backpath,desc := Backpath_dag_o2o(d,*nodes[5],*nodes[7])
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
	if Backverify_dag_o2o(d,*nodes[0],*nodes[7],z){
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}
	z = append(z,*nodes[3])
	if Backverify_dag_o2o(d,*nodes[0],*nodes[7],z){
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}
	z = append(z,*nodes[2])
	if Backverify_dag_o2o(d,*nodes[0],*nodes[7],z){
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}
	z = z[1:]
	z2 := make([]graph.Node,0)
	if Backverify_dag_o2o(d,*nodes[0],*nodes[7],z){
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}
	z = append(z,*nodes[6])
	if Backverify_dag_o2o(d,*nodes[0],*nodes[7],z2){
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}
}

func TestBackSearch(t *testing.T){
	d,nodes := back_example_1()
	status, z := Backsearch_dag_o2o(d,*nodes[5],*nodes[7])
	if status{
		for _,n := range z{
			fmt.Println(n.Getname())
		}
	}else{
		fmt.Printf("no invalid backdoor")
	}
}

func TestBackAllSearch(t *testing.T){
	d,nodes := back_example_1()
	status, zs := Backallsearch_dag_o2o(d,*nodes[5],*nodes[7])
	if status{
		for _,ns := range zs{
			for _,n := range ns{
				fmt.Printf("%s ",n.Getname())
			}
			fmt.Printf("\n")
		}
	}else{
		fmt.Printf("no invalid backdoor")
	}
}

func TestBackMinial(t *testing.T){
	d,nodes := back_example_1()
	status, zs := Backminimal_dag_o2o(d,*nodes[5],*nodes[7])
	if status{
		for _,ns := range zs{
			for _,n := range ns{
				fmt.Printf("%s ",n.Getname())
			}
			fmt.Printf("\n")
		}
	}else{
		fmt.Printf("no invalid backdoor")
	}
}