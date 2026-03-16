package main

import "fmt"

var budgetCategories =make(map[int]string)

func init(){
	fmt.Println("Initializing our budget categories...")
	budgetCategories[1] = "Housing"
	budgetCategories[2] = "Food"
	budgetCategories[3] = "Transportation"
	budgetCategories[4] = "Utilities"
	budgetCategories[5] = "Entertainment"
	budgetCategories[6] = "Healthcare"
	budgetCategories[7] = "Savings"
	budgetCategories[8] = "Miscellaneous"
	budgetCategories[9] = "Education"
}

func main(){
	for id, category := range budgetCategories {
		fmt.Printf("Category %d: %s\n", id, category)
	}
}