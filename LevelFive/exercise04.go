package main

import "fmt"

func main() {
	a := struct {
		customers []string
		balances  map[string]int
	}{
		customers: []string{"Bill", "Nancy", "Laura"},
		balances: map[string]int{
			"Bill":  100,
			"Nancy": 9000,
			"Laura": 100},
	}

	fmt.Println("Customers:", a.customers)
	fmt.Println("Balances:")
	for name, bal := range a.balances {
		fmt.Printf("\t%s: %d\n", name, bal)
	}

}
