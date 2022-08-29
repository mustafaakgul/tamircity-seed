package progress

type DeviceTypes struct {
	DeviceTypes []DeviceTypeRequest `json:"device_types"`
}

type DeviceTypeRequest struct {
	ServiceName    string `json:"service_name"`
	IdentityNumber string `json:"identity_number"`
	PhoneNumber    string `json:"phone_number"`
	Email          string `json:"email"`
	Iban           string `json:"iban"`
}

func DeviceTypeProgress() {
	println("DeviceTypeProgress")
}
