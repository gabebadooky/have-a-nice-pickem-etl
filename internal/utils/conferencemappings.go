package utils

type conferenceProperties struct {
	Name            string
	Abbreviation    string
	PowerConference bool
}

var ConferenceMapping map[string]conferenceProperties = map[string]conferenceProperties{
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
var acc conferenceProperties = conferenceProperties{
	Name:            "Atlantic Coast Conference",
	Abbreviation:    "ACC",
	PowerConference: true,
}

// American
var aac conferenceProperties = conferenceProperties{
	Name:            "The American Conference",
	Abbreviation:    "AAC",
	PowerConference: false,
}

// Big 12
var big12 conferenceProperties = conferenceProperties{
	Name:            "Big 12 Conference",
	Abbreviation:    "Big 12",
	PowerConference: true,
}

// Big Ten
var b1G conferenceProperties = conferenceProperties{
	Name:            "Big Ten Conference",
	Abbreviation:    "B1G",
	PowerConference: true,
}

// CUSA
var cusa conferenceProperties = conferenceProperties{
	Name:            "Conference USA",
	Abbreviation:    "CUSA",
	PowerConference: false,
}

// Independent
var independent conferenceProperties = conferenceProperties{
	Name:            "Independent",
	Abbreviation:    "Indep.",
	PowerConference: false,
}

// MAC
var mac conferenceProperties = conferenceProperties{
	Name:            "Mid-American Conference",
	Abbreviation:    "MAC",
	PowerConference: false,
}

// Mountain West
var mwc conferenceProperties = conferenceProperties{
	Name:            "Mountain West Conference",
	Abbreviation:    "MWC",
	PowerConference: false,
}

// Pac-12
var pac12 conferenceProperties = conferenceProperties{
	Name:            "Pacific Athletic Conference",
	Abbreviation:    "PAC 12",
	PowerConference: false,
}

// SEC
var sec conferenceProperties = conferenceProperties{
	Name:            "Southeastern Conference",
	Abbreviation:    "SEC",
	PowerConference: true,
}

// Sun Belt
var sbc conferenceProperties = conferenceProperties{
	Name:            "Sun Belt Conference",
	Abbreviation:    "SBC",
	PowerConference: false,
}
