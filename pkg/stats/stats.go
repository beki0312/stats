package stats

import (
	"github.com/beki0312/bank/v2/pkg/types"
)

//Avg рассчитывет средную сума платежа
func Avg(payments []types.Payment)types.Money{
	avg:=0
var sum int
		a:=0
	for _, payment := range payments {
		if payment.Status ==types.StatusFail{
			continue
		} 
		if payment.Amount > 0 {
			sum+=int(payment.Amount)
			a+=1
		}
	} 
			
		avg=sum/a
		return types.Money(avg)
}

//TotalInCategory находит сумму покупок в определённой категории
func TotalInCategory(payments []types.Payment, category types.Category) types.Money{
			var sum int
			for _, payment := range payments {
				
				if payment.Status ==types.StatusFail{
					continue
				}
				if payment.Category == category{
					sum+=int(payment.Amount)
				}
			
			}

				return types.Money(sum)

}
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money{
	categoriesPayment:=map[types.Category]types.Money{}
	avgPayment:=map[types.Category]types.Money{}
   
	   for _, payment := range payments {
		   
			   categoriesPayment[payment.Category]+=payment.Amount
			   avgPayment[payment.Category]++
		   
	   }
	   for key, value := range categoriesPayment {
		   avg:=avgPayment[key]
		   categoriesPayment[key]=value/avg		
	   }
	   return categoriesPayment
   }