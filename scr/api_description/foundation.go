package api_description

type Model interface {
	GetPhone() string
	IsValid() bool
	GetLocation() string
	GetProviderName() string
}

type ApiDescription interface {
	GetBaseUrl() string
	GetResource() string
	GetModelInstance() Model
	GetHttpMethod() string
	GetParams(phone string) map[string]interface{}
}