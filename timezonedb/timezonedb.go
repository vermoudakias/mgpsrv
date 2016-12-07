package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type TimezoneDBResponse struct {
	Status       string `json:"status"`
	Message      string `json:"message"`
	CountryCode  string `json:"countryCode"`
	ZoneName     string `json:"zoneName"`
	Abbreviation string `json:"abbreviation"`
	GMTOffset    string `json:"gmtOffset"`
	DST          string `json:"dst"`
	Timestamp    int64  `json:"timestamp"`
}

func main() {
	flag.Parse()
	if len(flag.Args()) < 3 {
		log.Fatal("Args: <key> <latitude> <longitude>")
	}
	var key, lat, lng string
	key = flag.Args()[0]
	lat = flag.Args()[1]
	lng = flag.Args()[2]
	req, err := http.NewRequest("GET", "http://api.timezonedb.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	qparams := req.URL.Query()
	qparams.Add("key", key)
	qparams.Add("format", "json")
	qparams.Add("lat", lat)
	qparams.Add("lng", lng)
	req.URL.RawQuery = qparams.Encode()
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", data)
	timezoneDBres := TimezoneDBResponse{}
	err = json.Unmarshal([]byte(data), &timezoneDBres)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("status=%s\n", timezoneDBres.Status)
	fmt.Printf("message=%s\n", timezoneDBres.Message)
	fmt.Printf("CC=%s\n", timezoneDBres.CountryCode)
	fmt.Printf("zone=%s\n", timezoneDBres.ZoneName)
	fmt.Printf("abbreviation=%s\n", timezoneDBres.Abbreviation)
	fmt.Printf("gmtOffset=%s\n", timezoneDBres.GMTOffset)
	fmt.Printf("dst=%s\n", timezoneDBres.DST)
	fmt.Printf("timestamp=%d\n", timezoneDBres.Timestamp)
}
