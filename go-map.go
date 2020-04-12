package main

import "fmt"

func main() {
	covidByCountry := make(map[string]int)
	covidByCountry["EUA"] = 200
	covidByCountry["Canada"] = 50
	covidByCountry["Mexico"] = 80
	covidByCountry["Brazil"] = 100

	fmt.Println(covidByCountry)

	newMapCovidByCountry := map[string]int{
		"Canada": 50,
		"EUA":    200,
		"Mexico": 80,
		"Brazil": 100,
	}
	fmt.Println(newMapCovidByCountry)
}
