package utils

var ConferenceMapping map[string]map[string]any = map[string]map[string]any{
	"1":   acc,
	"151": aac,
	"4":   big12,
	"5":   b1G,
	"12":  cusa,
	"18":  independent,
	"15":  mac,
	"17":  mwc,
	"9":   pac12,
	"8":   sec,
	"37":  sbc,
}

// ACC
var acc map[string]any = map[string]any{
	"Name":            "Atlantic Coast Conference",
	"Abbreviation":    "ACC",
	"PowerConference": true,
}

// American
var aac map[string]any = map[string]any{
	"Name":            "The American Conference",
	"Abbreviation":    "AAC",
	"PowerConference": false,
}

// Big 12
var big12 map[string]any = map[string]any{
	"Name":            "Big 12 Conference",
	"Abbreviation":    "Big 12",
	"PowerConference": true,
}

// Big Ten
var b1G map[string]any = map[string]any{
	"Name":            "Big Ten Conference",
	"Abbreviation":    "B1G",
	"PowerConference": true,
}

// CUSA
var cusa map[string]any = map[string]any{
	"Name":            "Conference USA",
	"Abbreviation":    "CUSA",
	"PowerConference": false,
}

// Independent
var independent map[string]any = map[string]any{
	"Name":            "Independent",
	"Abbreviation":    "Indep.",
	"PowerConference": false,
}

// MAC
var mac map[string]any = map[string]any{
	"Name":            "Mid-American Conference",
	"Abbreviation":    "MAC",
	"PowerConference": false,
}

// Mountain West
var mwc map[string]any = map[string]any{
	"Name":            "Mountain West Conference",
	"Abbreviation":    "MWC",
	"PowerConference": false,
}

// Pac-12
var pac12 map[string]any = map[string]any{
	"Name":            "Pacific Athletic Conference",
	"Abbreviation":    "PAC 12",
	"PowerConference": false,
}

// SEC
var sec map[string]any = map[string]any{
	"Name":            "Southeastern Conference",
	"Abbreviation":    "SEC",
	"PowerConference": true,
}

// Sun Belt
var sbc map[string]any = map[string]any{
	"Name":            "Sun Belt Conference",
	"Abbreviation":    "SBC",
	"PowerConference": false,
}
