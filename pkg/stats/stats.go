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

	return types.Money(sum / len(payments))
}

// TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := 0
	for _, pay := range payments {
		if pay.Status != types.StatusFail {
			if pay.Category == category {
				sum += int(pay.Amount)
			}
		}
	}
	return types.Money(sum)
}

// FilterByCategory возвращает платежи в указанной категории
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {

	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}

	return filtered
}

// CategoriesAvg возвращает map со средними показателем расходов в каждой категории
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	payCount := map[types.Category]int{}
	for _, payment := range payments {
		payCount[payment.Category] += 1
		categories[payment.Category] += payment.Amount
	}
	for each := range categories {
		categories[each] /= types.Money(payCount[each])
	}
	return categories
}
