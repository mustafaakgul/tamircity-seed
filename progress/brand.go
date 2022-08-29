package progress

type Brands struct {
	Brands []BrandRequest `json:"brands"`
}

type BrandRequest struct {
	ServiceName    string `json:"service_name"`
	IdentityNumber string `json:"identity_number"`
	PhoneNumber    string `json:"phone_number"`
	Email          string `json:"email"`
	Iban           string `json:"iban"`
}

func BrandProgress() {
	println("BrandProgress")
}
