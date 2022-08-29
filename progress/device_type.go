package progress

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type DeviceTypes struct {
	DeviceTypes []DeviceTypeRequest `json:"device_types"`
}

type DeviceTypeRequest struct {
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	IsActive         bool   `json:"is_active"`
}

func DeviceTypeProgress() {
	jsonFile, err := os.Open("data/device_types.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened device_types.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var deviceTypes DeviceTypes
	json.Unmarshal(byteValue, &deviceTypes)

	for i := 0; i < len(deviceTypes.DeviceTypes); i++ {
		deviceTypeRequest := DeviceTypeRequest{
			deviceTypes.DeviceTypes[i].Name,
			deviceTypes.DeviceTypes[i].ShortDescription,
			deviceTypes.DeviceTypes[i].IsActive}
		jsonDeviceType, _ := json.Marshal(deviceTypeRequest)

		response, err := http.Post("http://localhost:8080/api/v1/device-types",
			"application/json;charset=utf-8",
			bytes.NewBuffer(jsonDeviceType))

		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()
	}
}
