package functions

import "fmt"

func MapExample() {
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

	newMapCovidByCountry["UK"] = 100
	delete(newMapCovidByCountry, "Mexico")

	fmt.Println(newMapCovidByCountry)

	fmt.Println()
	for country, infected := range map[string]int{
		"Canada": 50,
		"EUA":    200,
		"Mexico": 80,
		"Brazil": 100,
	} {
		fmt.Printf("the count infected was %d on %s\n", infected, country)
		fmt.Println("------")
	}
}
