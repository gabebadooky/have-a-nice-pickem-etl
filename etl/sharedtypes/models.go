package sharedtypes

type Location struct {
	LocationID string
	Stadium    string
	City       string
	State      string
	Latitude   float64
	Longitude  float64
}

type Record struct {
	TeamID     string
	RecordType string
	Wins       uint
	Losses     uint
	Ties       uint
}
