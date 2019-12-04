package graph

import(
	"testing"
	"fmt" 
	"strconv"
)

func TestNode(t *testing.T){
	n:=NewDefaultNode("x")
	fmt.Println(n.Getname(),n.Isob())
	n.Setname("y")
	n.Setob(false)
	fmt.Println(n.Getname(),n.Isob())
}

func TestEdge(t *testing.T){
	A := NewDefaultNode("x")
	B := NewNode("y",false)
	e := NewDefaultEdge(A,B)
	fmt.Println(e)
	e = NewEdge(A,Arrow,B,Tail)
	fmt.Println(e)
	B.Setob(true)
	fmt.Println(A.Isob(),B.Isob(),e.Isob())
	fmt.Println(e)
	e=NewEdge(A,Undefined,B,Arrow)
	fmt.Println(e)
}

func TestNewGraph(t *testing.T){
	testnum := 4
	nodes := make([]*Node,0) 
	edges := make([]*Edge,0)
	for i :=0;i<testnum;i++{
		node := NewDefaultNode(strconv.Itoa(i))
		nodes = append(nodes,node)
	}
	for i :=0;i+1<testnum;i++{
		edge := NewEdge(nodes[i],Tail,nodes[i+1],Arrow)
		edges = append(edges,edge)
	}
	if g,err := NewGraph(nodes,edges);err != nil{
		fmt.Println(err.Error())
	}else{
		if _,ok := g.Toposort(); !ok{
			fmt.Println("has a cycle")
		}else{
			fmt.Println("has no cycle")
		}
		if err := g.AddEdge(NewEdge(nodes[0],Tail,nodes[testnum-1],Arrow));err != nil{
			fmt.Println(err.Error())
		}
		if err :=g.AddEdge(NewEdge(nodes[testnum-1],Tail,nodes[testnum-2],Arrow)); err!= nil{
			fmt.Println(err.Error())
		}
		if err :=g.AddEdge(NewEdge(nodes[testnum-1],Tail,nodes[testnum-3],Arrow)); err!= nil{
			fmt.Println(err.Error())
		}
		// sort2,_ :=g.Toposort()
		//  for _,n := range sort2{
		//  	fmt.Println(n.name)
		//  }
		if _,ok := g.Toposort(); !ok{
			fmt.Println("has a cycle")
		}else{
			fmt.Println("has no cycle")
		}
		if err := g.RemoveEdge(edges[testnum-2]);err !=nil{
			fmt.Println(err.Error())
		}
		if _,ok := g.Toposort(); !ok{
			fmt.Println("has a cycle")
		}else{
			fmt.Println("has no cycle")
		}
	}
	errnode :=NewDefaultNode(strconv.Itoa(testnum))
	edges = append(edges,NewEdge(nodes[0],Arrow,errnode,Tail))
	if _,err := NewGraph(nodes,edges);err !=nil{
		fmt.Println(err.Error())
	}
}

func TestDag(t *testing.T){
	testnum := 4
	nodes := make([]*Node,0) 
	edges := make([]*Edge,0)
	for i :=0;i<testnum;i++{
		node := NewDefaultNode(strconv.Itoa(i))
		nodes = append(nodes,node)
	}
	for i :=0;i+1<testnum;i++{
		edge := NewEdge(nodes[i],Tail,nodes[i+1],Arrow)
		edges = append(edges,edge)
	}
	edges = append(edges,NewEdge(nodes[testnum-1],Tail,nodes[testnum-3],Arrow)) 
	if _,err := NewDag(nodes,edges);err != nil{
		fmt.Println(err.Error())
	}
}

func TestTriple(t *testing.T){
	nodes := make([]*Node,0) 
	edges := make([]*Edge,0)
	triples :=make([]*Triple,0)
	for i :=0;i<5;i++{
		node := NewDefaultNode(strconv.Itoa(i))
		nodes = append(nodes,node)
	}
	for i:=0;i<3;i++{
		triple := NewTriple(nodes[i],nodes[i+1],nodes[i+2])
		triples = append(triples,triple)
	}
	edges = append(edges,NewEdge(nodes[0],Tail,nodes[1],Arrow))
	edges = append(edges,NewEdge(nodes[1],Tail,nodes[2],Arrow))
	edges = append(edges,NewEdge(nodes[3],Tail,nodes[2],Arrow))
	edges = append(edges,NewEdge(nodes[3],Tail,nodes[4],Arrow))
	if d,err :=NewDag(nodes,edges);err!= nil{
		fmt.Println(err.Error())
	}else{
		for _,t := range triples{
			_,s,_ := d.Identify(t)
			fmt.Println("the triple",t,"is a",s)
		}
	}
}

func TestPath(t *testing.T){
	nodes := make([]*Node,0)
	for i :=0;i<5;i++{
		node := NewDefaultNode(strconv.Itoa(i))
		nodes = append(nodes,node)
	}
	p := NewEmptyPath(*nodes[0],*nodes[4])
	p.PushNode(*nodes[1])
	p.PushNode(*nodes[3])
	fmt.Println(p)
	for i :=0;i<3;i++{
		if n,ok := p.PopNode();!ok{
			fmt.Println("the path is empty now!")
		}else{
			fmt.Println(n.name)
		}
	}
}