package identify

import(
	"LP_demon/graph"
	"github.com/crillab/gophersat/solver"
	// "fmt"
)
/*
It's a np-complete problem to find all the sets of nodes 
to satisify the backdoor criterion between X and Y. So we may also need to 
provide approximate algotithms or heruistic algorithms
*/

// return all the backpaths from treatment to outcome, and all descendants of the treatment in a DAG
func Backpath_dag_o2o(d *graph.Dag,treatment graph.Node, outcome graph.Node) ([]*graph.Path,[]graph.Node){
	desc := d.AllDescendant(treatment)
	backdoor :=  make([]*graph.Path,0)
	paths := d.DFSpath(treatment,outcome,false)
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
func Backverify_dag(d *graph.Dag, t []graph.Node, o []graph.Node, z []graph.Node) bool{
	mapz := make(map[graph.Node]bool)
	for _,n := range z{
		if !n.Isob(){
			return false
		}
		mapz[n] = true
	}
	for _,ti := range t{
		for _,oi := range o{
			backpath,desc := Backpath_dag_o2o(d,ti,oi)
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
						for _,des := range d.AllDescendant(n){
							if mapz[des]{
								ok_1 = true
							}
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
		}
	}
	return true 
}

//for a pair of treatment and outcome, determine whether there is a set of nodes can satisfy
//the backdoor criterion, if the answer is yes , also return a feasible solution
func Backsearch_dag(d *graph.Dag,t []graph.Node,o []graph.Node)(bool,[]graph.Node){
	clauses,nodes := backclauses(d,t,o)
	z := make([]graph.Node,0)
	pb := solver.ParseSlice(clauses)
	s := solver.New(pb)
	stat :=s.Solve()
	if stat != solver.Sat{
		return false,nil
	}
	m:=s.Model()
	for i,b := range m{
		if b{
			z = append(z,nodes[i])
		}
	}
	return true,z
}

//search all the sets of nodes can satisfy the backdoor criterion between a pair of 
//treatment and outcome.
//Attention: for a high dimensional graph, it may cost a lot of time to find all sets
func Backallsearch_dag(d *graph.Dag,t []graph.Node,o []graph.Node)(bool,[][]graph.Node){
	clauses,nodes := backclauses(d,t,o)
	zs := make([][]graph.Node,0)
	pb := solver.ParseSlice(clauses)
	s := solver.New(pb)
	stat :=s.Solve()
	if stat != solver.Sat{
		return false,nil
	}
	models := make(chan []bool)
	stop := make(chan struct{})
	go s.Enumerate(models,stop)
	for m := range models{
		z := make([]graph.Node,0)
		for i,b := range m{
			if b{
				z = append(z,nodes[i])
			}
		}
		zs = append(zs,z)
	}
	return true,zs
}

//turn the backdoor search problem to a sat problem
func backclauses(d *graph.Dag,t []graph.Node,o []graph.Node)([][]int,[]graph.Node){
	backnodes := make([]graph.Node,0) //record the nodes appear in the backpath
	nodeindex := make(map[graph.Node]int)//record the index of nodes in backnodes
	clauses := make([][]int,0) 
	for _,ti := range t{
		for _,oi := range o{
			backpath,desc := Backpath_dag_o2o(d,ti,oi)
			for _,p := range backpath{
				clause := make([]int,0)
				ty,_ := d.IdentifyPath(p)
				descs := make(map[int][]int)
				for s,ns := range ty{
					for _,n := range ns{
						if _,ok := nodeindex[n];!ok{
							backnodes = append(backnodes,n)
							nodeindex[n] = len(backnodes)
						}
						if s == "collider"{
							clause = append(clause,-nodeindex[n])
							loop:
							for _,des := range d.AllDescendant(n){
								for _,tj := range t{
									if des == tj{
										continue loop
									}
								}
								for _,oj := range o{
									if des == oj{
										continue loop
									}
								}
								if _,ok := nodeindex[des];!ok{
									backnodes = append(backnodes,des)
									nodeindex[des] = len(backnodes)
								}
								descs[nodeindex[n]] = append(descs[nodeindex[n]],nodeindex[des])
							}
						}else{
							clause = append(clause,nodeindex[n])
						}
					}
				}
				clauses = append(clauses,clause)
				for i,di := range descs{
					for _,d := range di{
						newclause := make([]int,len(clause))
						copy(newclause,clause)
						for index,j := range newclause{
							if i == -j{
								newclause[index] = -d
							}
						}
						clauses = append(clauses,newclause)
					}
				}
			}	
			for _,n := range desc{
				if _,ok := nodeindex[n]; ok{
					i := -nodeindex[n]
					clauses = append(clauses,[]int{i})
				}
			}		
		}
	}	
	for i,n := range backnodes{
		if !n.Isob(){
			clauses = append(clauses,[]int{-(i+1)})
		}
	}
	for _,n := range t{
		if _,ok := nodeindex[n]; ok{
			i := -nodeindex[n]
			clauses = append(clauses,[]int{i})
	
		}
	}
	return clauses,backnodes
}

//search all minimal sets of nodes satisfy the backdoor criterion
//minimal:any set such, if you removed any one of the variables from the set,it would no
//longer meet the criterion
func Backminimal_dag(d *graph.Dag,t []graph.Node,o []graph.Node)(bool,[][]graph.Node){
	b,z := Backallsearch_dag(d,t,o)
	minimal := make([][]graph.Node,0)
	if !b{
		return b,nil
	} 
	for _,ns := range z{
		flag := false
		// remove one node from z, judge if it still satisify the backdoor criterion
		for i,_ := range ns{
			nt := make([]graph.Node,len(ns))
			copy(nt,ns)
			zt := append(nt[:i],nt[i+1:]...)
			if Backverify_dag(d,t,o,zt){
				flag = true
				break
			}
		}
		if !flag{
			minimal = append(minimal,ns)
		}
	}
	return true,minimal
}