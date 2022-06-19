package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WeightData struct {
	ID   int64  `json:"id,omitempty"`
	Date string `json:"date,omitempty"`
	Min  int64  `json:"min,omitempty"`
	Max  int64  `json:"max,omitempty"`
	Diff int64  `json:"diff"`
}

type Mean struct {
	Max  float64 `json:"max"`
	Min  float64 `json:"min"`
	Diff float64 `json:"diff"`
}

type WeightResp struct {
	Data []WeightData `json:"data"`
	Mean Mean         `json:"mean"`
}

func main() {
	baseURL := "http://localhost:3000"
	client := &http.Client{}
	request, err := http.NewRequest("POST", baseURL+"/reset", nil)
	if err != nil {
		panic(err)
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	request, err = http.NewRequest("GET", baseURL+"/", nil)
	if err != nil {
		panic(err)
	}

	response, err = client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var index WeightResp
	err = json.Unmarshal(body, &index)
	if err != nil {
		panic(err)
	}

	payload := WeightData{
		Date: "2021-11-22",
		Min:  10,
		Max:  25,
	}

	payloadJson, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	request, err = http.NewRequest("POST", baseURL+"/create", bytes.NewBuffer(payloadJson))
	if err != nil {
		panic(err)
	}

	response, err = client.Do(request)
	if err != nil {
		panic(err)
	}

	payload = WeightData{
		Date: "2021-11-23",
		Min:  10,
		Max:  50,
	}

	payloadJson, err = json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	request, err = http.NewRequest("POST", baseURL+"/create", bytes.NewBuffer(payloadJson))
	if err != nil {
		panic(err)
	}

	response, err = client.Do(request)
	if err != nil {
		panic(err)
	}

	request, err = http.NewRequest("GET", baseURL+"/", nil)
	if err != nil {
		panic(err)
	}

	response, err = client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &index)
	if err != nil {
		panic(err)
	}

	fmt.Println(index)
	expect := WeightResp{
		Data: []WeightData{
			{
				Date: "2021-11-23",
				Min:  10,
				Max:  50,
				Diff: 40,
			},
			{
				Date: "2021-11-22",
				Min:  10,
				Max:  25,
				Diff: 15,
			},
		},
		Mean: Mean{
			Max:  37.5,
			Min:  10,
			Diff: 27.5,
		},
	}
	fmt.Printf("Expect %v\n", expect)
	fmt.Printf("Got    %v\n", index)
}
