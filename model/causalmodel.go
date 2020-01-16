package model

import(
	"LP_demon/graph"
	"LP_demon/identify"
	"LP_demon/estimate"
	"fmt"
	"sort"
)

type casual_model struct {
	dag graph.Dag
	dataset  estimate.Dataset
//	starta []estimate.Dataset
	treatment graph.Node 
	outcome graph.Node
	identify_flag bool
	backdoor_node [][]string
	frontdoor_node [][]string
}
func NewModel(d graph.Dag, data estimate.Dataset, treatment, outcome string) (cm *casual_model){
	if !d.IsStringIn(treatment){
		panic("the treatment variable is not in the graph ")
	}
	if !d.IsStringIn(outcome){
		panic("the outcome variable is not in the graph ")
	}
	return &casual_model{
		dag : d,
		dataset : data,
		treatment : *d.GetNode(treatment),
		outcome :  *d.GetNode(outcome),
	}
}

//use backdoor and frontdoor method to identify
func (cm *casual_model) Identify() {
	b,back_nodes := identify.Backminimal_dag(&cm.dag,[]graph.Node{cm.treatment},[]graph.Node{cm.outcome})	 
	if !b{
		fmt.Println("there is no back-door path from treatment to outcome")
	}else{
		fmt.Println("the back-door paths are:")
		var backstring [][]string
		for i,ns := range back_nodes{
			var s []string
			fmt.Printf("%v :",i+1)
			for _,n := range ns{
				fmt.Printf(" %s," ,n.Getname())
				s = append(s,n.Getname())
			}
			fmt.Printf("\n")
			backstring = append(backstring,s)
		}
		cm.backdoor_node = backstring
	}
	b,front_nodes := identify.Frontminimal_dag(&cm.dag,[]graph.Node{cm.treatment},[]graph.Node{cm.outcome})	 
	if !b{
		fmt.Println("there is no front-door path from treatment to outcome")
	}else{
		fmt.Println("the front-door paths are:")
		var frontstring [][]string
		for i,ns := range front_nodes{
			var s []string
			fmt.Printf("%v :",i+1)
			for _,n := range ns{
				fmt.Printf(" %s," ,n.Getname())
				s = append(s,n.Getname())
			}
			fmt.Printf("\n")
			frontstring = append(frontstring,s)
		}
		cm.frontdoor_node = frontstring
	}
	cm.identify_flag = true
}

//use propensity score and strata method to calculate ATE
func (cm *casual_model) Propensity_strata(i int ,boundry []float64){
	if !cm.identify_flag{
		panic("the identify method must be excuted before estimate method")
	}
	if cm.backdoor_node == nil{
		panic("there is no back-door path, can't use propensity_strata method to estimate ATE")
	}
	cm.dataset.Propensity(estimate.Propensity_logistic,cm.treatment.Getname(),cm.backdoor_node[0])
	sort.Sort(&cm.dataset)
	if strata,err := cm.dataset.Propensity_stratify(i,boundry);err !=nil{
		panic(err)
	}else{
		var sum float64
		for _,d := range strata{
			ate := d.ATE("xi","xj")
			// fmt.Println("ATE of the strata ",i,":",ate )
			sum = sum + ate
		}
		fmt.Println("propensity strata ATE: from",cm.treatment,"to",cm.outcome,sum/float64(i))
	}
}