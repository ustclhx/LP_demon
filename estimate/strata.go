package estimate

import(
	"fmt"
	"sort"
)

// stratifies the sample based on the propensity score
func (ds *Dataset)Propensity_stratify(num int, boundry []float64) ([]Dataset,error){
	var low,up int 
	strata := make([]Dataset,0)
	if ds.propensity_score == nil{
		err := fmt.Errorf("the propensity score must be calculated before stratifying")
		return nil,err
	}
	if !sort.IsSorted(ds){
		sort.Sort(ds)
	}
	if len(boundry)!= 2 || !(0 <= boundry[0])||!(boundry[0]< boundry[1])||!(boundry[1]<= 1){
		err := fmt.Errorf("the third argument of stratify should be a interval of [0,1]")
		return nil,err
	}
	for i := 0; i< ds.sample; i++{
		if ds.propensity_score[i] >= boundry[0]{
			low = i
			break
		}
		if i == ds.sample{
			err := fmt.Errorf("all the propensity score are smaller than lower boundry, please reset")
			return nil,err
		}
	}
	for j := ds.sample -1; j >= 0; j--{
		if ds.propensity_score[j] <= boundry[1]{
			up = j
			break
		}
		if j == 0{
			err := fmt.Errorf("all the propensity score are bigger than upper boundry, please reset")
			return nil,err
		}
	}
	sample := up-low+1
	bin_num := sample/num
	for i:= 1; i < num; i++{
		strata = append(strata,ds.subset(low+(i-1)*bin_num+1,low+i*bin_num))
	}
	strata = append(strata,ds.subset(low+(num-1)*bin_num+1,up+1))
	return strata,nil
}