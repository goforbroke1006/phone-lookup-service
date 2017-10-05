package api_description

const (
	STATUS_VALID_CONFIRMED             string = "VALID_CONFIRMED"
	STATUS_VALID_UNCONFIRMED           string = "VALID_UNCONFIRMED"
	STATUS_INVALID                     string = "INVALID"
	STATUS_DELAYED                     string = "DELAYED"
	STATUS_RATE_LIMIT_EXCEEDED         string = "RATE_LIMIT_EXCEEDED"
	STATUS_API_KEY_INVALID_OR_DEPLETED string = "API_KEY_INVALID_OR_DEPLETED"
)

const (
	LT_FIXED_LINE      string = "FIXED_LINE"
	LT_MOBILE          string = "MOBILE"
	LT_VOIP            string = "VOIP"
	LT_TOLL_FREE       string = "TOLL_FREE"
	LT_PREMIUM_RATE    string = "PREMIUM_RATE"
	LT_SHARED_COST     string = "SHARED_COST"
	LT_PERSONAL_NUMBER string = "PERSONAL_NUMBER"
	LT_PAGER           string = "PAGER"
	LT_UAN             string = "UAN"
	LT_VOICEMAIL       string = "VOICEMAIL"
)

type ResponseModel struct {
	Status              string `json:"status"`
	LineType            string `json:"linetype"`
	Location            string `json:"location"`
	CountryCode         string `json:"countrycode"`
	FormatNational      string `json:"formatnational"`
	FormatInternational string `json:"formatinternational"`
	Mcc                 string `json:"mcc"`
	Mnc                 string `json:"mnc"`
}

func (m *ResponseModel) GetPhone() string {
	return m.FormatNational
}
func (m *ResponseModel) IsValid() bool {
	return m.Status == STATUS_VALID_CONFIRMED
}
func (m *ResponseModel) GetLocation() string {
	return m.Location
}
func (m *ResponseModel) GetProviderName() string {
	return m.Mnc
}

type ApiPhoneValidatorNetAPI struct{}

func (d *ApiPhoneValidatorNetAPI) GetBaseUrl() string {
	return "http://api.phone-validator.net"
}
func (d *ApiPhoneValidatorNetAPI) GetResource() string {
	return "/api/v2/verify"
}
func (d *ApiPhoneValidatorNetAPI) GetModelInstance() Model {
	return new(ResponseModel)
}
func (d *ApiPhoneValidatorNetAPI) GetHttpMethod() string {
	return "GET"
}
func (d *ApiPhoneValidatorNetAPI) GetParams(phone string) map[string]interface{} {
	return map[string]interface{}{
		"PhoneNumber": phone,
	}
}
