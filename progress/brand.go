package progress

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Brands struct {
	Brands []BrandRequest `json:"brands"`
}

type BrandRequest struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
	//DeviceTypes []*DeviceTypeRequest `json:"device_types"`
}

type Todo struct {
	ID                int         `json:"ID"`
	CreatedAt         time.Time   `json:"CreatedAt"`
	UpdatedAt         time.Time   `json:"UpdatedAt"`
	DeletedAt         interface{} `json:"DeletedAt"`
	Name              string      `json:"Name"`
	ShortDescription  string      `json:"ShortDescription"`
	IsActive          bool        `json:"IsActive"`
	TechnicalServices interface{} `json:"TechnicalServices"`
	Brands            interface{} `json:"Brands"`
	FixTypes          interface{} `json:"FixTypes"`
	Models            interface{} `json:"Models"`
	Reservation       interface{} `json:"Reservation"`
}

func BrandProgress() {

	response, err := http.Get("http://localhost:8080/api/v1/device-types/1")
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	var todo Todo
	json.Unmarshal(bodyBytes, &todo)
	fmt.Println(todo)

	jsonFile, err := os.Open("data/brands.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened brands.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var brands Brands
	json.Unmarshal(byteValue, &brands)

	for i := 0; i < len(brands.Brands); i++ {
		brandRequest := BrandRequest{
			brands.Brands[i].Name,
			brands.Brands[i].IsActive}
		jsonBrand, _ := json.Marshal(brandRequest)

		response, err := http.Post("http://localhost:8080/api/v1/brands",
			"application/json;charset=utf-8",
			bytes.NewBuffer(jsonBrand))

		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()
	}
}
