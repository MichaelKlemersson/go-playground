package main

import (
	"fmt"
	"go-playground/src/github.com/mklemersson/structs/country"
)

func main() {
	brazil := country.Country{
		Name:   "Brazil",
		Locale: "BR",
	}
	germany := country.CountryWithId{
		Country: country.Country{
			Name:   "Germany",
			Locale: "DE",
		},
		Id: 12,
	}

	fmt.Println(brazil.Name, germany.GetName(), germany.GetId())
	fmt.Println(brazil.ToJson())
	fmt.Println(germany.ToJson())
}
