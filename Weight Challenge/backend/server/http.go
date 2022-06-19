package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/backend/domain/weight"
)

func HandleCreateWeightData(w http.ResponseWriter, r *http.Request) {
	AllowCors(&w)

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var param weight.WeightData
	err = json.Unmarshal(body, &param)
	if err != nil {
		panic(err)
	}

	err = WeightUsecase.CreateWeight(param)
	if err != nil {
		panic(err)
	}

	resp, err := json.Marshal(param)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func HandleIndexWeightData(w http.ResponseWriter, r *http.Request) {
	AllowCors(&w)

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	data, err := WeightUsecase.GetAllWeightData()
	if err != nil {
		panic(err)
	}

	resp, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func HandleEditWeightData(w http.ResponseWriter, r *http.Request) {
	AllowCors(&w)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var param weight.WeightData
	err = json.Unmarshal(body, &param)
	if err != nil {
		panic(err)
	}

	err = WeightUsecase.EditWeightData(param)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func HandleReset(w http.ResponseWriter, r *http.Request) {
	AllowCors(&w)

	WeightUsecase.Reset()
}
