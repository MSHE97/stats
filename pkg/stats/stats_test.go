package stats

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/MSHE97/bank/v2/pkg/types"
)

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   100_00,
			Category: "cafe",
		},
		{
			ID:       2,
			Amount:   50_00,
			Category: "market",
		}, {
			ID:       3,
			Amount:   125_00,
			Category: "market",
		},
		{
			ID:       4,
			Amount:   25_00,
			Category: "cafe",
		},
	}

	result := CategoriesAvg(payments)
	expected := map[types.Category]types.Money{
		"cafe":   62_50,
		"market": 87_50,
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("got\n%v \nwant:\n %v ", result, expected)
	}
}

func TestFilterbyCategory(t *testing.T) {
	var payments []types.Payment
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len !=0")
	}
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   100_00,
			Category: "cafe",
		},
		{
			ID:       2,
			Amount:   50_00,
			Category: "market",
		}, {
			ID:       3,
			Amount:   125_00,
			Category: "transport",
		},
		{
			ID:       4,
			Amount:   25_00,
			Category: "cafe",
		},
	}

	fmt.Println(TotalInCategory(payments, "cafe"))

	// Output: 12500
}
func ExampleAvg() {

	payments := []types.Payment{
		{
			ID:       1,
			Amount:   100_00,
			Category: "cafe",
		},
		{
			ID:       2,
			Amount:   50_00,
			Category: "market",
		}, {
			ID:       3,
			Amount:   125_00,
			Category: "transport",
		},
		{
			ID:       4,
			Amount:   25_00,
			Category: "cafe",
		},
	}

	fmt.Println(Avg(payments))

	// Output: 7500
}
