package country

import (
	"encoding/json"
	"fmt"
)

type Country struct {
	Name   string `json:"name"`
	Locale string `json:"locale"`
}

type CountryWithId struct {
	Id int `json:"id"`
	Country
}

func (country Country) GetName() string {
	return country.Name
}

func (country CountryWithId) GetId() int {
	return country.Id
}

func ToJson(country interface{}) string {
	jsonBuffer, err := json.Marshal(country)

	if err != nil {
		fmt.Println("Fail to convert country with id to JSON")
	}

	return string(jsonBuffer)
}
