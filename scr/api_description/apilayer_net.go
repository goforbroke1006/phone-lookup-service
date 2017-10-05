package api_description

type ApiLayerResponse struct {
	Valid               bool   `json:"valid"`
	Number              string `json:"number"`
	LocalFormat         string `json:"local_format"`
	InternationalFormat string `json:"international_format"`
	CountryPrefix       string `json:"country_prefix"`
	CountryCode         string `json:"country_code"`
	CountryName         string `json:"country_name"`
	Location            string `json:"location"`
	Carrier             string `json:"carrier"`
	LineType            string `json:"line_type"`
}

func (m *ApiLayerResponse) GetPhone() string {
	return m.Number
}
func (m *ApiLayerResponse) IsValid() bool {
	return m.Valid
}
func (m *ApiLayerResponse) GetLocation() string {
	return m.Location
}
func (m *ApiLayerResponse) GetProviderName() string {
	return m.Carrier
}

type ApiLayerAPI struct{}

func (d *ApiLayerAPI) GetBaseUrl() string {
	return "http://apilayer.net"
}
func (d *ApiLayerAPI) GetResource() string {
	return "/api/validate"
}
func (d *ApiLayerAPI) GetModelInstance() interface{} {
	return ApiLayerResponse{}
}
func (d *ApiLayerAPI) GetHttpMethod() string {
	return "GET"
}
func (d *ApiLayerAPI) GetParams(phone string) map[string]interface{} {
	return map[string]interface{}{
		"number": phone,
	}
}
