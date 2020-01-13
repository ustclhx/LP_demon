package model

import(
	"LP_demon/graph"
	"LP_demon/estimate"
)

type casual_model struct {
	graph graph.Graph
	dataset  estimate.Dataset
	starta []estimate.Dataset
//	propensity map[string]
}
func NewModel(g graph.Graph, data estimate.Dataset) (cm *casual_model){
	return &casual_model{
		graph : g,
		dataset : data,
	}
}

func (*casual_model) identify() ([][]graph.Node){
	
}