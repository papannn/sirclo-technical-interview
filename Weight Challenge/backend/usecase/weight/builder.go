package ucweight

import (
	"github.com/backend/domain/weight"
)

func BuildWeightResp(weightData []weight.WeightData) WeightResp {
	var resp WeightResp

	totalMax := float64(0)
	totalMin := float64(0)
	totalDiff := float64(0)
	for idx, data := range weightData {
		diff := data.Max - data.Min
		weightData[idx].Diff = diff
		totalMax += float64(data.Max)
		totalMin += float64(data.Min)
		totalDiff += float64(diff)
	}

	resp.Data = weightData
	if len(weightData) > 0 {
		resp.Mean.Max = totalMax / float64(len(weightData))
		resp.Mean.Min = totalMin / float64(len(weightData))
		resp.Mean.Diff = totalDiff / float64(len(weightData))
	}

	return resp
}
