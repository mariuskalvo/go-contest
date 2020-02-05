package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	url := "http://10.200.233.89:8080/contest"

	me := map[string]string{
		"name":        "Marius Kalvø",
		"email":       "mariuskalvoe@gmail.com",
		"source_code": "github.com/mariuskalvo/go-contest",
	}

	var value, err = json.Marshal(me)
	if err != nil {
		panic("Could not marshall json object")
	}

	request, err := http.NewRequest("POST", url, strings.NewReader(string(value)))
	if err != nil {
		panic(err)
	}
	request.Header.Add("Content-Type", "application/json")
	httpClient := &http.Client{}
	resp, err := httpClient.Do(request)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)

	resp.Body.Close()
}
