package opencagelocation

type OpencageEndpoint struct {
	Results []ResultProperty `json:"results"`
}

type ResultProperty struct {
	Annotations AnnotationsProperty `json:"annotations"`
	Components  ComponentsProperty  `json:"components"`
	Geometry    GeometryProperty    `json:"geometry"`
}

type AnnotationsProperty struct {
	Maidenhead string `json:"maidenhead"`
}

type ComponentsProperty struct {
	City    string `json:"city"`
	Stadium string `json:"stadium"`
	State   string `json:"state"`
}

type GeometryProperty struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lng"`
}
