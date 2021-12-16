package country

import (
	"encoding/json"
	"fmt"
)

type Country struct {
	Name   string
	Locale string
}

type CountryWithMethod struct {
	Id int
	Country
}

func (country Country) GetName() string {
	return country.Name
}

func (country CountryWithMethod) GetId() int {
	return country.Id
}

func (country Country) ToJson() string {
	jsonBuffer, err := json.Marshal(country)

	if err != nil {
		fmt.Println("Fail to convert country to JSON")
	}

	return string(jsonBuffer)
}