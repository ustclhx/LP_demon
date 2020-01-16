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
	if dataset,err := ReadfromCSV("test1.csv");err != nil{
		fmt.Println(err.Error())
	}else{
		dataset.Head(10)
	}
}

func TestPropensity(t *testing.T){
	d,nodes := identify.Back_example_1()
	_,z := identify.Backminimal_dag(d,[]graph.Node{*d.GetNode("xi")},[]graph.Node{*d.GetNode("xj")})
	src := rand.NewSource(5)
	ds,_ := Dag_linear_dataset(d,20,src,nodes[0],nodes[7])
	zs := make([]string,0)
	for _,zn := range z[0] {
		zs = append(zs,zn.Getname())
	}
	ds.Propensity(Propensity_logistic,"xi",zs)
	ds.Head(20)
	fmt.Println(sort.IsSorted(ds))
	sort.Sort(ds)
	ds.Head(20)
	fmt.Println(sort.IsSorted(ds))
	if strata,err := ds.Propensity_stratify(5,[]float64{0,1});err != nil{
		fmt.Println(err.Error())
	}else{
		for _, d := range strata{
			d.Head(4)
		}
	}
}

func TestATE(t *testing.T){
	d,nodes := identify.Back_example_1()
	_,z := identify.Backminimal_dag(d,[]graph.Node{*d.GetNode("xi")},[]graph.Node{*d.GetNode("xj")})
	src := rand.NewSource(10)
	ds,true_effect := Dag_linear_dataset(d,10000,src,nodes[0],nodes[7])
	fmt.Println("true effect:",true_effect)
	zs := make([]string,0)
	for _,zn := range z[0] {
		zs = append(zs,zn.Getname())
	}
	fmt.Println("ATE of whole dataset :", ds.ATE("xi","xj"))
	ds.Propensity(Propensity_logistic,"xi",zs)	
	//strata_ate := make([]float64,0)
	if strata,err := ds.Propensity_stratify(20,[]float64{0.2,0.8});err != nil{
		fmt.Println(err.Error())
	}else{
		var sum float64
		for i,d := range strata{
			ate := d.ATE("xi","xj")
			fmt.Println("ATE of the strata ",i,":",ate )
			sum = sum + ate
		}
		fmt.Println("strata ATE:",sum/20)
	}
}