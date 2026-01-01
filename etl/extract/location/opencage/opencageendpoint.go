package opencagelocation

type OpencageEndpoint struct {
	Results []ResultProperty `json:"results"`
}

type ResultProperty struct {
	Geometry GeometryProperty `json:"geometry"`
}

type GeometryProperty struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lng"`
}
