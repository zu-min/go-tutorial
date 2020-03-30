package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i := range numbers {
		fmt.Println("slice item", i, "is", numbers[i])
	}

	countryCapitalMap := map[string]string{"flance": "paris", "italy": "rome", "japan": "tokyo"}
	for country := range countryCapitalMap {
		fmt.Println("capital of", country, "is", countryCapitalMap[country])
	}

	for country, capital := range countryCapitalMap {
		fmt.Println("capital of", country, "is", capital)
	}
}
