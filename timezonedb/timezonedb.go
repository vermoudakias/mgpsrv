package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 2 {
		log.Fatal("Args: <latitude> <longitude>")
	}
	var lat, lng string
	lat = flag.Args()[0]
	lng = flag.Args()[1]
	req, err := http.NewRequest("GET", "http://api.timezonedb.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	qparams := req.URL.Query()
	qparams.Add("key", "XXXXXXXX")
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
	fmt.Printf("%s", data)
}
