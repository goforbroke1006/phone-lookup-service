package main

import (
	"fmt"
	"os"
	"regexp"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

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

func main() {
	if len(os.Args) < 3 {
		panic("You should specify access_key and phone which needs to check")
	}

	accessKey := os.Args[1]
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
	fmt.Println(responseApi.Carrier)
}

func doPanic(err error) {
	if err != nil {
		panic(err.Error())
	}
}
