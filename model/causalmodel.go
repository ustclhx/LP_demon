package model

import(
	"LP_demon/graph"
)

type casual_model struct {
	graph graph.Graph
	dataset  map[string][]float64
//	propensity map[string]
}
func NewModel(g graph.Graph, data map[string][]float64) (cm *casual_model){
	return &casual_model{
		graph : g,
		dataset : data,
	}
}

func (*casual_model) identify() ([][]graph.Node){
	
}