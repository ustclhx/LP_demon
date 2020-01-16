package estimate

import(
	"fmt"
	"encoding/csv"
	"os"
	"io"
	"strconv"
) 

type Dataset struct{
	data map[string][]float64 
	propensity_score []float64
	sample int //the number of samples 
}

func ReadfromCSV(fliepath string) (*Dataset,error){
	if _,err := os.Stat(fliepath); err != nil{
		return nil,err
	}
	file,err := os.Open(fliepath)
	if err != nil{
		return nil,err
	} 
	defer file.Close()
	reader := csv.NewReader(file)
	var variables []string
	data := make(map[string][]float64)
	var sample int 
	record,err := reader.Read()
	if err != nil{
		return nil,err
	}else {
		variables = record
	}
	record,err = reader.Read()
	for err != io.EOF{
		for i,v := range record{
			fl,er := strconv.ParseFloat(v,64)
			if er != nil{
				return nil, er
			} 
			data[variables[i]] = append(data[variables[i]],fl)	
		}
     	sample ++
		record,err = reader.Read()	
	}
	return &Dataset{
		data:data,
		sample:sample,
	},nil
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

//calculate the propensity score of dataset,using the method defined in f
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

//return the subset of the dataset from the ith element to the jth element
func (ds *Dataset) subset(i,j int) Dataset{
	if !(i<j)&&!(j<ds.sample){
		panic("the arguments of Dataset.subject are not suitable")
	}
	data := make(map[string][]float64)
	var propensity_score []float64
	for s,v := range ds.data{
		data[s] = v[i-1:j]
	}
	if ds.propensity_score != nil{
		propensity_score = ds.propensity_score[i-1:j]
	}else{
		propensity_score = nil
	}
	return Dataset{
		data : data,
		propensity_score : propensity_score,
		sample : j-i+1,
	}
}

//calculate the average treat effect, ie: E(Y|X=1)-E(Y|X=0)
func (ds *Dataset) ATE(treatment,outcome string) float64{
	var t_count, c_count int
	var t_sum, c_sum float64
	for i := 0; i<ds.sample; i++{
		if ds.data[treatment][i] == 0{
			c_count ++
			c_sum = c_sum + ds.data[outcome][i]
		}else if ds.data[treatment][i] == 1{
			t_count ++
			t_sum = t_sum + ds.data[outcome][i]
		}else{
			panic("the treatment variable of ATE method should be a binary variable")
		}
	}
    return t_sum/float64(t_count) - c_sum/float64(c_count)
}
