package main

import(
	"LP_demon/identify"
	"LP_demon/model"
	"LP_demon/estimate"
	"golang.org/x/exp/rand"
)

func main(){
	var sample int = 1000
	//import the graph and dataset
	//you can load your causal graph and data based real world and hypothesis
	d,nodes := identify.Back_example_1()
	src := rand.NewSource(24)
	dataset,_ := estimate.Dag_linear_dataset(d,sample,src,nodes[0],nodes[7])
	//initialises a new causal model
	cm := model.NewModel(*d,*dataset,"xi","xj")
	//use backdoor and frontdoor method to judge whether the causal effect
	//between treatment and outcome can be identified
	cm.Identify()
	//use propensity score and strata method way to estimate the ATE
	cm.Propensity_strata(sample/50,[]float64{0.2, 0.8})
}
