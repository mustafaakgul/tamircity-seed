package progress

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Brands struct {
	Brands []BrandRequest `json:"brands"`
}

type BrandRequest struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

func BrandProgress() {
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
