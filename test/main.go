package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type TechnicalServiceRequest struct {
	ServiceName    string `json:"service_name"`
	IdentityNumber string `json:"identity_number"`
	PhoneNumber    string `json:"phone_number"`
	Email          string `json:"email"`
	Iban           string `json:"iban"`
}

func main() {
	//technicalServiceRequest := TechnicalServiceRequest{"hello1seed", "hello2seed", "hello3seed", "hello4seed", "hello5seed"}
	//jsonTodo, err := json.Marshal(technicalServiceRequest)
	//println(jsonTodo, err)
	var jsonTodo2, _ = ReadToJson2Struct("technical_services.json")
	println(jsonTodo2.ServiceName)
	technicalServiceRequest := TechnicalServiceRequest{
		jsonTodo2.ServiceName,
		jsonTodo2.IdentityNumber,
		jsonTodo2.PhoneNumber,
		jsonTodo2.Email,
		jsonTodo2.Iban}
	jsonTodo, _ := json.Marshal(technicalServiceRequest)
	println(jsonTodo)

	var jsonTodo3, _ = ReadToJson("technical_services.json")
	println(jsonTodo3)

	//response, err := http.Post("http://localhost:8080/api/v1/technical-services",
	//	"application/json;charset=utf-8",
	//	bytes.NewBuffer(jsonTodo))
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer response.Body.Close()
	//
	//bodyBytes, _ := ioutil.ReadAll(response.Body)
	//
	//bodyString := string(bodyBytes)
	//fmt.Println(bodyString)
}

func ReadToJson(path string) (result map[string]interface{}, err error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	//var anyJson map[string]interface{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &result)
	//println(result)
	defer jsonFile.Close()
	return result, nil
}

func ReadToJson2Struct(path string) (result TechnicalServiceRequest, err error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return
	}

	var t TechnicalServiceRequest
	byteValue2, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue2, &t)

	//var anyJson map[string]interface{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &result)
	//println(result)
	defer jsonFile.Close()
	return t, nil
}

/*
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type TechnicalServiceRequest struct {
	ServiceName    string `json:"service_name"`
	IdentityNumber string `json:"identity_number"`
	PhoneNumber    string `json:"phone_number"`
	Email          string `json:"email"`
	Iban           string `json:"iban"`
}

func main() {
	//technicalServiceRequest := TechnicalServiceRequest{"hello1seed", "hello2seed", "hello3seed", "hello4seed", "hello5seed"}
	//jsonTodo, err := json.Marshal(technicalServiceRequest)
	//println(jsonTodo, err)
	var jsonTodo2, _ = ReadToJson2Struct("technical_services.json")
	println(jsonTodo2.ServiceName)
	technicalServiceRequest := TechnicalServiceRequest{
		jsonTodo2.ServiceName,
		jsonTodo2.IdentityNumber,
		jsonTodo2.PhoneNumber,
		jsonTodo2.Email,
		jsonTodo2.Iban}
	jsonTodo, err := json.Marshal(technicalServiceRequest)

	var jsonTodo3, _ = ReadToJson("technical_services.json")
	println(jsonTodo3)

	response, err := http.Post("http://localhost:8080/api/v1/technical-services",
		"application/json;charset=utf-8",
		bytes.NewBuffer(jsonTodo))

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

func ReadToJson(path string) (result map[string]interface{}, err error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	//var anyJson map[string]interface{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &result)
	//println(result)
	defer jsonFile.Close()
	return result, nil
}

func ReadToJson2Struct(path string) (result TechnicalServiceRequest, err error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return
	}

	var t TechnicalServiceRequest
	byteValue2, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue2, &t)

	//var anyJson map[string]interface{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &result)
	//println(result)
	defer jsonFile.Close()
	return t, nil
}

*/
