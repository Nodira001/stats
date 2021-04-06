package stats

import "github.com/akhrorov/bank/v2/pkg/types"

func Avg(payment []types.Payment) types.Money {
	var sum types.Money
	i := 0
	for _, v := range payment {
		if v.Status == types.StatusFail {
			continue

		}
		sum += v.Amount
		i++
	}
	return sum / types.Money(i)

}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}

		if payment.Category == category {
			sum += payment.Amount

		}

	}
	return sum

}
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	counts := map[types.Category]types.Money{}
	for _, payment := range payments {

		categories[payment.Category] += payment.Amount
		counts[payment.Category] += 1
	}

	for category := range categories {
		categories[category] = categories[category] / counts[category]

	}
	return categories

}

func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	result := map[types.Category]types.Money{}
	if len(first) >= len(second) {
		for key := range first {
			result[key] = second[key] - first[key]
		}
		return result

	}
	for key := range second {
		result[key] = second[key] - first[key]
	}
	return result

}
