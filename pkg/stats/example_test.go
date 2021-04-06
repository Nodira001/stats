package stats

import ("github.com/akhrorov/bank/v2/pkg/types"
"fmt"
)

func ExampleAvg() {
	payments := [] types.Payment {
		{ID: 1, Amount: 10_000_00, Category: "cars", Status: types.StatusOk },
		{ID: 2, Amount: 20_000_00, Category: "book", Status: types.StatusOk },
		{ID: 3, Amount: 20_000_00, Category: "book", Status: types.StatusFail },
		} 
	
	 fmt.Println(Avg(payments))

	//Output:
	//1500000
}

func ExampleTotalInCategory () {
	payments := [] types.Payment {
		{ID: 1, Amount: 10_000_00, Category: "cars", Status: types.StatusOk },
		{ID: 2, Amount: 20_000_00, Category: "book", Status: types.StatusOk },
		{ID: 3, Amount: 20_000_00, Category: "book", Status: types.StatusFail },
		{ID: 4, Amount: 20_000_00, Category: "book", Status: types.StatusOk },
		}
		
	fmt.Println(TotalInCategory(payments,"book"))
	//Output:
	//4000000


}
  