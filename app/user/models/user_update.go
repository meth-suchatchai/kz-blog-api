package usermodels

type UpdateUser struct {
	Name         *string `json:"name"`
	CountryCode  *string `json:"country_code"`
	MobileNumber *string `json:"mobile_number"`
}
