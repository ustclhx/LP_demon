package graph

import(
	"fmt"
	"strings"
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
			err = fmt.Errorf("the edge %s,is not a directed edge",edge)
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
				
func FastNewDag(ns []string,es []string)(d *Dag,err error){
	nodes := make([]*Node,0)
	edges := make([]*Edge,0)
	node_index := make(map[string]*Node)
	for _,n := range ns{
		s := strings.Split(n,",")
		if len(s)== 1{
			node :=NewDefaultNode(s[0])
			nodes = append(nodes,node)
			node_index[s[0]] = node
		}else if len(s) == 2{
			if s[1] == "T" || s[1] ==  "t"{
				node := NewNode(s[0],true)
				nodes = append(nodes,node)
				node_index[s[0]] = node
			}else if s[1] == "F" || s[1] == "f"{
				node := NewNode(s[0],false)
				nodes = append(nodes,node)
				node_index[s[0]] = node
			}else{
				err = fmt.Errorf("observation flag of node should be T of F")
				return nil,err
			}
		}else{
			err = fmt.Errorf("the length of string representing node should be 1 or 2")
			return nil,err
		}
	}
	for _,e := range es{
		s := strings.Split(e,",")
		if len(s)==2{
			edge := NewEdge(node_index[s[0]],Tail,node_index[s[1]],Arrow)
			edges = append(edges,edge)
		}else{
			err = fmt.Errorf("the length of string representing node should be 2")
			return nil,err
		}
	}
	d,err = NewDag(nodes,edges)
	return 
}



