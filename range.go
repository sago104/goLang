package main

import "fmt"

func main() {
	// Create a slice
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Print the numbers
	for i := range numbers {
		fmt.Println("Slice item", i, "is: ", numbers[i])
	}

	// Create a map
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}

	// Print map using keys
	for country := range countryCapitalMap {
		fmt.Println("Capital of ", country, "is ", countryCapitalMap[country])
	}

	// Print map using key-value
	for country, capital := range countryCapitalMap {
		fmt.Println("Capital of ", country, "is ", capital)
	}

}
