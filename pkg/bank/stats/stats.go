package stats

import "github.com/harryfinder/average/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	var length int
	var sum int64
	for i, operations := range payments {
		sum += int64(operations.Amount)
		length += i
	}

	average := types.Money(sum / int64(length))

	return average
}
