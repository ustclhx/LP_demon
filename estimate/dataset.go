package estimate

import(
	"fmt"
) 

type Dataset struct{
	data map[string][]float64 
	propensity_score []float64
	sample int //the number of samples 
}

func (ds *Dataset) GetVariate(s string) []float64{
	return ds.data[s]
}

func (ds *Dataset) GetSample(i int) (map[string]float64,error){
	if i>ds.sample{
		err := fmt.Errorf("the argument %v exceed the amount of sample %v",i,ds.sample)
		return nil,err
	}
	si := make(map[string]float64)
	for s,d := range ds.data{
		si[s] = d[i] 
	}
	return si,nil
}

func (ds *Dataset) Head(i int){
	variable := make([]string,0)
	for v,_ := range ds.data{
		fmt.Printf("%s ,",v)
		variable = append(variable,v)
	}
	if ds.propensity_score != nil{
		fmt.Printf("propensity score ")
	}
	fmt.Printf("\n")
	if ds.sample <= i{
		i = ds.sample
	}
	for it := 0; it <i; it ++{
		si,_ := ds.GetSample(it)
		for j := 0; j<len(variable);j++{
			fmt.Printf("%v ,",si[variable[j]])
		}
		if ds.propensity_score != nil{
			fmt.Printf("%v ", ds.propensity_score[it])
		}
		fmt.Printf("\n")
	}
}

func (ds *Dataset) Propensity(f Propensity_function ,t string, co []string){
	if err := f(ds,t,co); err != nil{
		panic(err)
	}
}

func (ds *Dataset) Len() int{
	return ds.sample
}

func (ds *Dataset) Swap(i,j int) {
	ds.propensity_score[i],ds.propensity_score[j] = ds.propensity_score[j],ds.propensity_score[i]
	for _,v := range ds.data{
		v[i],v[j] = v[j],v[i]
	}
}

func (ds *Dataset) Less(i,j int) bool{
	return ds.propensity_score[i] < ds.propensity_score[j]
}

// func ReadfromCSV()*dataset{
	
// }