package main

import (
	"fmt"
	"os"
	"./scr/api_description"
	"net/http"
)

func main() {
	if len(os.Args) < 3 {
		panic("You should specify access_key and phone which needs to check")
	}

	var providers = map[string]api_description.ApiDescription{
		"apilayer_net":            new(api_description.ApiLayerAPI),
		"api_phone_validator_net": new(api_description.ApiPhoneValidatorNetAPI),
	}

	fmt.Println("Check providers...")
	for pk, pobj := range providers {
		if ! Ping(pobj) {
			delete(providers, pk)
			continue
		}
		fmt.Println("  Load [" + pk + "] -> " + pobj.GetBaseUrl())
	}

	/*accessKey := os.Args[1]
	phone := os.Args[2]

	re := regexp.MustCompile("[^0-9]+")
	phone = re.ReplaceAllString(phone, "")

	url := fmt.Sprintf("http://apilayer.net/api/validate?access_key=%s&number=%s&country_code=RU&format=1",
		accessKey,
		phone)

	resp, err := http.Get(url)
	doPanic(err)
	defer resp.Body.Close()

	var responseApi ApiLayerResponse

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(body, &responseApi)

	if false == responseApi.Valid {
		if len(responseApi.Number) > 0 {
			panic(fmt.Sprintf("Phone %s", phone))
		} else {
			panic("Response body contains unexpected JSON-data")
		}
	}

	fmt.Println(responseApi.InternationalFormat)
	fmt.Println(responseApi.Location)
	fmt.Println(responseApi.Carrier)*/
}

func Ping(d api_description.ApiDescription) bool {
	url := d.GetBaseUrl() + "/"
	_, err := http.Get(url)
	fmt.Println(err)
	return nil == err
}

func Display(rm interface{}) {
	v, _ := rm.(api_description.Model)
	fmt.Println(v.GetLocation())
	fmt.Println(v.GetProviderName())
}

func doPanic(err error) {
	if err != nil {
		panic(err.Error())
	}
}
