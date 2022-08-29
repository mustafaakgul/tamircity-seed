package progress

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Models struct {
	Models []ModelRequest `json:"models"`
}

type ModelRequest struct {
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	IsActive         bool   `json:"is_active"`
}

func ModelProgress() {
	jsonFile, err := os.Open("data/models.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened models.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var models Models
	json.Unmarshal(byteValue, &models)

	for i := 0; i < len(models.Models); i++ {
		modelRequest := ModelRequest{
			models.Models[i].Name,
			models.Models[i].ShortDescription,
			models.Models[i].IsActive}
		jsonModel, _ := json.Marshal(modelRequest)

		response, err := http.Post("http://localhost:8080/api/v1/models",
			"application/json;charset=utf-8",
			bytes.NewBuffer(jsonModel))

		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()
	}
}
