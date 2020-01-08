package estimate

import(
	"fmt"
)

type dataset struct{
	data map[string][]float64
	sample int 
}

func (ds dataset) GetVariate(s string) []float64{
	return ds.data[s]
}

func (ds dataset) GetSample(i int) (map[string]float64,error){
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
// func ReadfromCSV()*dataset{
	
// }