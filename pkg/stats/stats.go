package stats

import (
	"github.com/MSHE97/bank/v2/pkg/types"
)

// Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	sum := 0
	for _, pay := range payments {
		if pay.Status != types.StatusFail {
			sum += int(pay.Amount)
		}
	}

	return types.Money( sum / len(payments) )
}

// TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory( payments []types.Payment, category types.Category) types.Money {
	sum := 0
	for _, pay := range payments {
		if pay.Status != types.StatusFail {
			if (pay.Category == category){
				sum += int(pay.Amount)
			}
		}
	}
	return types.Money(sum)
}