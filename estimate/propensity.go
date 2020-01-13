package estimate

import(
	"github.com/cdipaolo/goml/linear"
	"github.com/cdipaolo/goml/base"
	// "fmt"
)

type Propensity_function func(ds *Dataset,t string,co []string)error 

//use logistic regression to calculate propensity score P(x|z)
func Propensity_logistic(ds *Dataset,t string,co []string )error{
	covariate := make([][]float64,0)
	treatment := ds.GetVariate(t)
	if len(ds.propensity_score)!= 0{
		ds.propensity_score = make([]float64,0)
	}
	// fmt.Println(treatment)
	for i:=0; i<ds.sample;i++{
		value := make([]float64,0)
		for _,v := range co{
			value = append(value,ds.GetVariate(v)[i])
		}
		covariate = append(covariate,value)
	}
	//Learning Rate & Maximum Iterations(0.00001&1000), Regularization(0)
	model := linear.NewLogistic(base.BatchGA,0.00001,0,1000,covariate,treatment)
	if err := model.Learn(); err != nil{
		return err
	} 
	for _,input := range covariate{
		if score,err := model.Predict(input); err != nil{
			return err 
		}else{
			ds.propensity_score = append(ds.propensity_score,score[0]) 
		}
		
	}
	return nil
}


