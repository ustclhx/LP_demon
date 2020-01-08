package estimate

import(
	// "LP_demon/graph"
	"LP_demon/identify"
	"testing"
	"fmt"
	"golang.org/x/exp/rand"
)

func TestDataset(t *testing.T){
	d,nodes := identify.Back_example_1()
	src := rand.NewSource(2)
	dataset,eff := dag_linear_dataset(d,10,src,nodes[0],nodes[7])
	for i:=0; i< len(nodes); i++{
		fmt.Printf("%s ,",nodes[i].Getname())
	}
	fmt.Printf("\n")
	for i:= 0; i< 10; i++{
		if sample,err := dataset.GetSample(i); err == nil{
			for j:=0; j< len(nodes);j++{
				fmt.Printf("%v ,",sample[nodes[j].Getname()])
			}
			fmt.Printf("\n")
		}
	}
	fmt.Println(eff)
}

// func TestPropensity(t *testing.T){
// 	d,nodes :=
// }
