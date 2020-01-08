package estimate

import(
	"github.com/cdipaolo/goml/linear"
	"github.com/cdipaolo/goml/base"
)

//use logistic regression to calculate propensity score P(x|z)
func propensity_logistic(ds dataset,t string,co []string )(model *linear.Logistic,err error){
	covariate := make([][]float64,len(co))
	treatment := ds.GetVariate(t)
	for _,v := range co{
		covariate = append(covariate,ds.GetVariate(v))
	}
	//Learning Rate & Maximum Iterations(0.00001&1000), Regularization(0)
	model = linear.NewLogistic(base.BatchGA,0.00001,0,1000,covariate,treatment)
	if err := model.Learn(); err != nil{
		return nil,err
	}
	return model,nil
}