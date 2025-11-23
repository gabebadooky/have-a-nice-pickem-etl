package pickemstructs

type OpencageResponse struct {
	Results []Result `json:"results"`
}

type Result struct {
	Geometry Geometry `json:"geometry"`
}

type Geometry struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lng"`
}
