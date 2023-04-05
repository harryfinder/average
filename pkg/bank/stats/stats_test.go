package stats_test

import (
	"fmt"
	"github.com/harryfinder/average/pkg/bank/stats"
	"github.com/harryfinder/average/pkg/bank/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1234,
			Amount:   2_000_00,
			Category: "restaurnt",
		},
		{
			ID:       1111,
			Amount:   3_000_00,
			Category: "chemist's",
		},
		{
			ID:       7000,
			Amount:   7_000_00,
			Category: "auto",
		},
	}
	fmt.Println(stats.Avg(payments))

	//output:
	//400000
}
