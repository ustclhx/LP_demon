package graph

import(
//	"fmt"
)

// use dfs method to find all the paths between two nodes
// use the argument isdirected to choose directed path or undirected path
func (g *Graph)DFSpath(from Node, to Node, isdirected bool) []*Path{
	path := make([]Node,0)
	path = append(path,from)
	paths := make([]*Path,0)
	g.dfs(from,from,to,path,&paths,isdirected)
	return paths
}

//recursive fuction of dfs
func (g *Graph) dfs(from Node, now Node, to Node,path []Node, paths *[]*Path, isdirected bool){
	if now == to {
		p := NewPath(from,path,to)
		*paths = append(*paths,p)
	}else{
		for n,e :=range g.edges[now]{
			if !isdirected||(e.Endpoint(now)==Tail && e.Endpoint(n)== Arrow){
				state := false
				//the node in the path will not be considered again
				for _,innode := range path{
					if n == innode {
						state = true
					}
				}
				if !state{
					newp := make([]Node,len(path))
					copy(newp,path)
					newp =append(newp,n)
					g.dfs(from,n,to,newp,paths,isdirected)	
				}
			}
		} 
	}
}

func (d *Dag) AllDescendant(n Node) []Node{
	nodestack := make([]Node,0)
	desc := make(map[Node] bool)
	descendant := make([]Node,0)
	nodestack = append(nodestack,n)
	for len(nodestack)>0{
		node := nodestack [len(nodestack)-1]
		nodestack = nodestack[:len(nodestack)-1]
		for adj,e := range d.edges[node]{
			if e.endpoints[adj] == Arrow && !desc[adj] && adj != n{
				desc[adj] = true
				nodestack = append(nodestack,adj)
			}
		}
	}
	for node,b := range desc{
		if b{
			descendant = append(descendant,node)
		}
	}
	return descendant
}