package progress

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type TechnicalServices struct {
	TechnicalServices []TechnicalServiceRequest `json:"technical_services"`
}

type TechnicalServiceRequest struct {
	ServiceName    string `json:"service_name"`
	IdentityNumber string `json:"identity_number"`
	PhoneNumber    string `json:"phone_number"`
	Email          string `json:"email"`
	Iban           string `json:"iban"`
}

func TechnicalServiceProgress() {
	jsonFile, err := os.Open("data/technical_services.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened tech.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var technicalServices TechnicalServices
	json.Unmarshal(byteValue, &technicalServices)

	for i := 0; i < len(technicalServices.TechnicalServices); i++ {
		technicalServiceRequest := TechnicalServiceRequest{
			technicalServices.TechnicalServices[i].ServiceName,
			technicalServices.TechnicalServices[i].IdentityNumber,
			technicalServices.TechnicalServices[i].PhoneNumber,
			technicalServices.TechnicalServices[i].Email,
			technicalServices.TechnicalServices[i].Iban}
		jsonTechSer, _ := json.Marshal(technicalServiceRequest)

		response, err := http.Post("http://localhost:8080/api/v1/technical-services",
			"application/json;charset=utf-8",
			bytes.NewBuffer(jsonTechSer))

		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()
	}
}
