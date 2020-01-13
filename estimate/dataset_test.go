package estimate

import(
	"LP_demon/graph"
	"LP_demon/identify"
	"testing"
	"fmt"
	"sort"
	"golang.org/x/exp/rand"
)

func TestDataset(t *testing.T){
	d,nodes := identify.Back_example_1()
	src := rand.NewSource(2)
	dataset,eff := Dag_linear_dataset(d,1000,src,nodes[0],nodes[7])
	dataset.Head(10)
	fmt.Println(eff)
}

func TestPropensity(t *testing.T){
	d,nodes := identify.Back_example_1()
	_,z := identify.Backminimal_dag(d,[]graph.Node{*d.GetNode("xi")},[]graph.Node{*d.GetNode("xj")})
	src := rand.NewSource(5)
	ds,_ := Dag_linear_dataset(d,100,src,nodes[0],nodes[7])
	zs := make([]string,0)
	for _,zn := range z[0] {
		zs = append(zs,zn.Getname())
	}
	ds.Propensity(Propensity_logistic,"xi",zs)
	sort.Sort(ds)
	ds.Head(20)
}
