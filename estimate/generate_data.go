package estimate

import(
	"LP_demon/graph"
	"math"
	"gonum.org/v1/gonum/stat/distuv"
	"golang.org/x/exp/rand"
)

//return a dataset generated by the given causal graph, and the true effect(random generate) from treatment to the outcome 
func Dag_linear_dataset(d *graph.Dag, sample int, src rand.Source,treatment *graph.Node, outcome *graph.Node) ( *Dataset, float64){
	nodes,_ := d.Toposort()
	ds := make(map[string][]float64)
	coefficient := make(map[graph.Node]map[graph.Node]float64)
	c_rand := distuv.Uniform{Min:0,Max:5,Src:src}
	for i,n := range nodes{
		parents := make([][]float64,0)
		weights := make([]float64,0)
		for j:=0; j<i; j++{
			if d.IsDirectto(nodes[j],n){
				parents = append(parents,ds[nodes[j].Getname()])
				c := c_rand.Rand()
				weights = append(weights,c)
				if coefficient[nodes[j]] == nil{
					coefficient[nodes[j]] = make(map[graph.Node]float64)
				}
				coefficient[nodes[j]][n] = c
			}
		} 
		if n != *treatment{
			ds[n.Getname()] = linear_model(parents,weights,sample,0,src)
		}else{
			ds[n.Getname()] = logistic_model(parents,weights,sample,0,src)
		}
	}
	// calculate the true effect from treatment to outcome
	var true_effect float64
	path := d.DFSpath(*treatment,*outcome,true)
	for _,p := range path{
		var pe float64 = 1
		p_ns := p.Nodes()
		for i,_ := range p_ns{
			if i != 0{
				pe = coefficient[p_ns[i-1]][p_ns[i]]*pe
			}	
		}
		true_effect = true_effect + pe
	}
	return &Dataset{
		data : ds,
		sample : sample,
	}, true_effect
}

func sigmoid(x float64) float64{
	return 1/(1+math.Exp(-x))
}

func linear_model(data[][]float64,weights []float64, sample int,offset float64,src rand.Source)[]float64{
	output := make([]float64,0)
	normal := distuv.Normal{Mu:offset,Sigma:1,Src:src}
	for i := 0; i<sample; i++{
		x := normal.Rand()
		for j := 0; j<len(data);j++{
			x = x + weights[j]*data[j][i]
		}
		output = append(output,x)
	}
	return output
}

func logistic_model(data [][]float64,weights []float64, sample int, offset float64, src rand.Source) []float64{
	output := make([]float64,0)
	linear := linear_model(data,weights,sample,offset,src)
	for i := 0; i<sample; i++{
		p := sigmoid(linear[i])
		x := distuv.Bernoulli{P:p,Src:src}.Rand()
		output = append(output,x)
	}
	return output
}

