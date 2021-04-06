package stats
import (
	"github.com/beki0312/bank/v2/pkg/types"
	"reflect"
	"testing"
	"fmt"
)
func ExampleAvg(){ payments := [] types.Payment{
	{ 		ID: 1,		Amount:110,		Category:"salon",	},
	{		ID: 2,		Amount:60,		Category:"avto",	},
	{		ID: 3,		Amount:-50,		Category:"taxi",	},
	{		ID: 4,		Amount:100,		Category:"salon",	},
	}
	avg:=Avg(payments)
	fmt.Println(avg)
total:=TotalInCategory(payments,"salon")
fmt.Println(total)
//Output:
//90
//210
}


func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID:1, Category: "auto", Amount: 100},
		{ID:2, Category: "food", Amount: 80},
		{ID:3, Category: "auto", Amount: 100},
		{ID:4, Category: "auto", Amount: 100},
		{ID:5, Category: "fun", Amount: 50},
	}
   expected :=map[types.Category]types.Money{
	   "auto": 100,
	   "food": 80,
	   "fun": 50,
   }
   result :=CategoriesAvg(payments)
   if !reflect.DeepEqual(expected,result){
	   t.Errorf("invalid result, expected: %v, actal: %v", expected,result)
   }
}	
func TestCategoriesAvg_nil(t *testing.T) {
var payments []types.Payment
result := CategoriesAvg(payments)
 if len(result) !=0 {
	 t.Errorf("\n got >% v want>nil", result)
 }
}

func TestCategoriesAvg_one(t *testing.T) {
 payments := []types.Payment{
	{ID:1, Category: "auto", Amount: 100},
	{ID:2, Category: "auto", Amount: 100},
 }
expected :=map[types.Category]types.Money{
	"auto": 100,
}
result :=CategoriesAvg(payments)
if !reflect.DeepEqual(expected,result){
	t.Errorf("\n got>% v want>nil",result)
}
}