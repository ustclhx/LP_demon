package graph

import(
//	"fmt"
)

// use dfs method to find all the paths between two nodes
func (g *Graph)DFSpath(from Node, to Node) []*Path{
	path := make([]Node,0)
	path = append(path,from)
	paths := make([]*Path,0)
	g.dfs(from,from,to,path,&paths)
	return paths
}

//recursive fuction of dfs
func (g *Graph) dfs(from Node, now Node, to Node,path []Node, paths *[]*Path){
	if now == to {
		p := NewPath(from,path,to)
		*paths = append(*paths,p)
	}else{
		for n,_ :=range g.edges[now]{
			state := false
			for _,innode := range path{
				if n == innode {
					state = true
				}
			}
			if !state{
				newp := make([]Node,len(path))
				copy(newp,path)
				newp =append(newp,n)
				g.dfs(from,n,to,newp,paths)	
			}
		} 
	}
}