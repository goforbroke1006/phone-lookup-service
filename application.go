package main

import (
	"fmt"
	"os"
	"net/http"
	"regexp"
	"./scr/config"
	"./scr/api_description"
	"io/ioutil"
	"encoding/json"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		panic("You should specify phone which needs to check")
	}

	var configuration config.Configuration;
	configuration.LoadConfiguration("config.json")

	var descriptions = map[string]api_description.ApiDescription{
		"apilayer_net":            new(api_description.ApiLayerAPI),
		"api_phone_validator_net": new(api_description.ApiPhoneValidatorNetAPI),
		"neutrinoapi_com":         new(api_description.NeutrinoApiAPI),
	}

	phone := os.Args[1]
	re := regexp.MustCompile("[^0-9]+")
	phone = re.ReplaceAllString(phone, "")

	httpClient := &http.Client{}

	for sysName, descr := range descriptions {
		req, err := http.NewRequest(descr.GetHttpMethod(), descr.GetBaseUrl()+descr.GetResource(), nil)
		doPanic(err)

		var provider *config.Provider = configuration.GetProvider(sysName)
		loadParams(req, provider, descr.GetParams(phone))

		resp, err := httpClient.Do(req)
		doPanic(err)

		var apiResponse api_description.Model = descr.GetModelInstance()

		body, err := ioutil.ReadAll(resp.Body)
		doPanic(err)

		json.Unmarshal(body, &apiResponse)

		resp.Body.Close()

		if apiResponse.IsValid() {
			fmt.Println(descr.GetBaseUrl() + descr.GetResource())
			Display(apiResponse)
			//break
		}
	}
}

func loadParams(req *http.Request, provider *config.Provider, options map[string]interface{}) {
	if strings.ToLower(req.Method) == "get" {
		values := req.URL.Query()
		for pn, pv := range provider.Params {
			values.Add(pn, pv)
		}
		for pn, pv := range options {
			values.Add(pn, pv.(string))
		}
		req.URL.RawQuery = values.Encode()
	}
	if strings.ToLower(req.Method) == "post" {
		panic("POST params loading is not implemented!")
	}
}

func Display(m api_description.Model) {
	fmt.Println(m.GetPhone())
	fmt.Println(m.GetLocation())
	fmt.Println(m.GetProviderName())
}

func doPanic(err error) {
	if err != nil {
		panic(err.Error())
	}
}
