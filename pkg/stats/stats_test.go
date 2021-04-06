package stats

import (
	"reflect"
	"testing"

	"github.com/akhrorov/bank/v2/pkg/types"
)

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Amount: 1_000_00, Category: "mobile", Status: types.StatusOk},
		{ID: 2, Amount: 2_000_00, Category: "cars", Status: types.StatusOk},
		{ID: 3, Amount: 3_000_00, Category: "cars", Status: types.StatusFail},
	}

	expected := map[types.Category]types.Money{
		"mobile": 1_000_00,
		"cars":   2_500_00,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Error! result: %v, expected: %v", result, expected)

	}

}

func TestPeriodsDynamic(t *testing.T) {
	first := map[types.Category]types.Money{
		"foods": 10, "cars": 20,
	}

	second := map[types.Category]types.Money{
		"foods": 5, "cars": 3,
	}

	expected := map[types.Category]types.Money{
		"foods": -5, "cars": -17,
	}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {

		t.Errorf("Error! expexred %v, result: %v", expected, result)

	}

}
