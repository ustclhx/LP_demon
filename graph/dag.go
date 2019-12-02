package graph

import(
	"fmt"
)
/* a class for directed acyclic graph, 
which contains only directed edge and no cycles
*/
type Dag struct{
	Graph
}

//you can only add directed edge into a dag
func (d *Dag)AddEdge(e *Edge)(err error){
	if err = d.AddDirectedEdge(e); err!=nil{
		return
	}
	_,ok := d.Toposort()
	if !ok{
		err = fmt.Errorf("after add the edge %s, the Dag will has a cycle",e)
	}
	return
}

func NewDag(ns []*Node, es []*Edge) (d *Dag, err error){
	for _,edge := range es{
		if !edge.IsDirected(){
			err =  fmt.Errorf("the edge %s,is not a directed edge",edge)
			return
		}
	}
	g,er := NewGraph(ns,es)
	if er != nil{
		return nil,er 
	}
	d = &Dag{
		Graph : *g, 
	}
	if  _,ok := d.Toposort(); !ok{
		err = fmt.Errorf("the dag want to creat has a cycle")
		return nil,err
	}
    return d,nil
}
                 


