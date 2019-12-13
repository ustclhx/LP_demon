package identify

import(
	"LP_demon/graph"
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
				for _,n := range ns{
					if mapz[n]{
						flag = true
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