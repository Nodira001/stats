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

