package identify

import(
	"LP_demon/graph"
	"github.com/crillab/gophersat/solver"
	"fmt"
)
/*
It's a np-complete problem to find all the sets of nodes 
to satisify the backdoor criterion between X and Y. So we may also need to 
provide approximate algotithms or heruistic algorithms
*/

// return all the backpaths from treatment to outcome, and all descendants of the treatment in a DAG
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

//determine whether a set of nodes meet the backdoor criterion between treatment and outcome
func backverify_dag_o2o(d *graph.Dag, t graph.Node, o graph.Node, z []graph.Node) bool{
	mapz := make(map[graph.Node]bool)
	for _,n := range z{
		if !n.Isob(){
			return false
		}
		mapz[n] = true
	}
	backpath,desc := backpath_dag_o2o(d,t,o)
	for _,di := range desc{
		if mapz[di]{
			return false
		}
	}
	for _,path := range backpath{
		var ok_1,ok_2 bool
		ty,_ := d.IdentifyPath(path)
		// determine whether the backpath is blocked by a unadjusted collider
		if ty["collider"] == nil{
			ok_1 = true
		}else{
			for _,n := range ty["collider"]{
				if mapz[n] {
					ok_1 = true
				}
			}
		}
		// determine whether the backpath is blocked by an adjusted fork or chain
		for _,n := range ty["fork"]{
			if mapz[n]{
				ok_2 = true
			}
		}
		for _,n := range ty["chain"]{
			if mapz[n]{
				ok_2 = true
			}
		}
		//determine whether the backpath is blocked
		if ok_1 && !ok_2{
			return false
		}
	}
	return true 
}

//for a pair of treatment and outcome, determine whether there is a set of nodes can meet
//the backdoor criterion, if the answer is yes , also return a feasible solution
func backsearch_dag_o2o(d *graph.Dag,t graph.Node,o graph.Node)(bool,[]graph.Node){
	backpath,desc := backpath_dag_o2o(d,t,o)
	nodeindex := make(map[graph.Node]int)
	clauses := make([][]int,0) 
	z := make([]graph.Node,0)
	var status bool
	nodes := d.Nodes()
	//turn the backdoor search problem to a sat problem
	for i,n := range nodes{
		nodeindex[*n] = i+1
		if !n.Isob(){
			clauses = append(clauses,[]int{-(i+1)})
		}
	}
	for _,n := range desc{
		i := -nodeindex[n]
		clauses = append(clauses,[]int{i})
	}
	for _,p := range backpath{
		clause := make([]int,0)
		ty,_ := d.IdentifyPath(p)
		for _,n := range ty["collider"]{
			i := -nodeindex[n]
			clause = append(clause,i)
		}
		for _,n := range ty["fork"]{
			i := nodeindex[n]
			clause = append(clause,i)
		}
		for _,n := range ty["chain"]{
			i := nodeindex[n]
			clause = append(clause,i)
		}
		clauses = append(clauses,clause)
	}
	for _,c :=range clauses{
		for _,i:= range c{
			fmt.Printf("%v ",i)	
		}
		fmt.Printf("\n")
	}
	// use the gophersat lib to solve the sat problem
	pb := solver.ParseSlice(clauses)
	s := solver.New(pb)
	stat :=s.Solve()
	if stat == solver.Sat{
		status = true
	}
	m:=s.Model()
	for i,b := range m{
		if b{
			z = append(z,*nodes[i])
		}
	}
	return status,z
} 