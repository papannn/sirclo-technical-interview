package ucweight

import "github.com/backend/domain/weight"

type Mean struct {
	Max  float64 `json:"max"`
	Min  float64 `json:"min"`
	Diff float64 `json:"diff"`
}

type WeightResp struct {
	Data []weight.WeightData `json:"data"`
	Mean Mean                `json:"mean"`
}
