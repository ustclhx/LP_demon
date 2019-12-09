package identify

import(
	"LP_demon/graph"
)
/*
It's a np-complete problem to find all the sets of nodes 
to satisify the backdoor criterion between X and Y. So we may also need to 
provide approximate algotithms or heruistic algorithms
*/

// return all the backpaths from treatment to outcome, and all descendants
// of the treatment in a DAG
func backpath_dag_o2o(d *graph.Dag,treatment graph.Node, outcome graph.Node) ([]*graph.Path,[]graph.Node){
	desc := d.AllDescendant(treatment)
	backdoor :=  make([]*graph.Path,0)
	paths := d.DFSpath(treatment,outcome)
	for _,path := range paths{
		node := path.Nodes()[1]
		e := d.Edges(treatment)[node]
		if e.Endpoint(treatment) == graph.Arrow{
			backdoor = append(backdoor,path)
		}
	}
	return backdoor,desc
}

// determine whether a set of nodes meet the backdoor criterion between 
//treatment and outcome
func backverify_dag_o2o(d *graph.Dag, t graph.Node, o graph.Node, z []graph.Node) bool{
	backpath,desc := backpath_dag_o2o(d,t,o)
	for _,zi := range z{
		for _,di := range desc{
			if zi == di{
				return false
			}
		}
	}
	for _,path := range backpath{

	}

}