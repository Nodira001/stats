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
