package api_description

type NeutrinoApiResponse struct {
	Valid                    bool   `json:"valid"`
	Type                     string `json:"type"`
	IsMobile                 bool   `json:"is-mobile"`
	Location                 string `json:"location"`
	Country                  string `json:"country"`
	CountryCode              string `json:"country-code"`
	InternationalCallingCode int    `json:"international-calling-code"`
	InternationalNumber      string `json:"international-number"`
	LocalNumber              string `json:"local-number"`
}

func (m *NeutrinoApiResponse) GetPhone() string {
	return m.LocalNumber
}
func (m *NeutrinoApiResponse) IsValid() bool {
	return m.Valid
}
func (m *NeutrinoApiResponse) GetLocation() string {
	return m.Location
}
func (m *NeutrinoApiResponse) GetProviderName() string {
	return ""
}

type NeutrinoApiAPI struct{}

func (d *NeutrinoApiAPI) GetBaseUrl() string {
	return "https://neutrinoapi.com"
}
func (d *NeutrinoApiAPI) GetResource() string {
	return "/phone-validate"
}
func (d *NeutrinoApiAPI) GetModelInstance() Model {
	return new(NeutrinoApiResponse)
}
func (d *NeutrinoApiAPI) GetHttpMethod() string {
	return "GET"
}
func (d *NeutrinoApiAPI) GetParams(phone string) map[string]interface{} {
	return map[string]interface{}{
		"number": phone,
	}
}
