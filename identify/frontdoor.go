package identify

import(
	"LP_demon/graph"
	"github.com/crillab/gophersat/solver"
	// "fmt"
)


func Frontverify_dag(d *graph.Dag, t []graph.Node, o []graph.Node, z []graph.Node) bool{
	mapz := make(map[graph.Node]bool)
	for _,n := range z {
		if !n.Isob(){
			return false
		}
		mapz[n] = true
	}
	if !Backverify_dag(d,t,z,nil){
		return false
	}
	if !Backverify_dag(d,z,o,t){
		return false
	}
	for _,ti := range t{
		for _,oi := range o{
			frontpath := d.DFSpath(ti,oi,true)
			for _,p := range frontpath{
				var flag bool
				ns := p.Nodes()
				for i,n := range ns{
					if(i!= 0 || i != len(ns)-1){
						if mapz[n]{
							flag = true
						}
					}
				}
				if !flag{
					return false
				}
			}
		}
	}
	return true
}

func Frontsearch_dag(d *graph.Dag,t []graph.Node,o []graph.Node)(bool,[]graph.Node){
	clauses,nodes := frontclauses(d,t,o)
	z := make([]graph.Node,0)
	pb := solver.ParseSlice(clauses)
	s := solver.New(pb)
	stat := s.Solve()
	if stat != solver.Sat{
		return false,nil
	}
	m := s.Model()
	for i,b := range m {
		if b{
			z = append(z,nodes[i])
		}
	}
	return true,z
}

func Frontallsearch_dag(d *graph.Dag,t []graph.Node,o []graph.Node)(bool,[][]graph.Node){
	clauses,nodes := frontclauses(d,t,o)
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

func Frontminimal_dag(d *graph.Dag,t []graph.Node,o []graph.Node)(bool,[][]graph.Node){
	b,z := Frontallsearch_dag(d,t,o)
	if !b{
		return b,nil
	} 
	minimal := make([][]graph.Node,0)
	for _,ns := range z{
		flag := false
		for i,_ := range ns{
			nt := make([]graph.Node,len(ns))
			copy(nt,ns)
			zt := append(nt[:i],nt[i+1:]...)
			if Frontverify_dag(d,t,o,zt){
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


func frontclauses(d *graph.Dag,t []graph.Node,o []graph.Node)([][]int,[]graph.Node){
	frontnodes := make([]graph.Node,0)
	nodeindex := make(map[graph.Node]int)
	clauses := make([][]int,0)
	for _,ti := range t{
		for _,oi := range o{
			frontpath := d.DFSpath(ti,oi,true)
			for _,p := range frontpath{
				clause := make([]int,0)
				ns := p.Nodes()
				for i,n := range ns{
					if(i !=0 && i != len(ns)-1){
						if _,ok := nodeindex[n];!ok{
							if Backverify_dag(d,t,[]graph.Node{n},nil) && Backverify_dag(d,[]graph.Node{n},o,t) && n.Isob(){
								frontnodes = append(frontnodes,n)
								nodeindex[n] = len(frontnodes)
							}else{
								nodeindex[n] = -1
							}
						} 
						if nodeindex[n] != -1{
							clause = append(clause,nodeindex[n])
						}
					}
				}
				if len(clause)>0{
					clauses = append(clauses,clause)
				}
			}
		}
	}
	return clauses,frontnodes
}