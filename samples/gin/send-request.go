package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
type IPLocation struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}
func main() {
	ip := "101.229.179.144"
	ipstring := "http://ip-api.com/json/" + ip + "?lang=zh-CN"
	response, err := http.Get(ipstring)
	if err != nil || response.StatusCode != http.StatusOK {
		panic(err)
		return
	}
	body, _ := ioutil.ReadAll(response.Body)

	var _IPLocation IPLocation
	if err = json.Unmarshal(body, &_IPLocation); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return
	}

	fmt.Println(_IPLocation)
}
