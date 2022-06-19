package weight

type WeightData struct {
	ID   int64  `json:"id,omitempty"`
	Date string `json:"date,omitempty"`
	Min  int64  `json:"min,omitempty"`
	Max  int64  `json:"max,omitempty"`
	Diff int64  `json:"diff"`
}
