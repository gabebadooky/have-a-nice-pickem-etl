package location

import "have-a-nice-pickem-etl/etl/pickemstructs"

func ParseStadium(espnGameDetails pickemstructs.ESPNGameDetailsResponse) string {
	var stadium string = espnGameDetails.GameInfo.Venue.FullName
	return stadium
}

func ParseCity(espnGameDetails pickemstructs.ESPNGameDetailsResponse) string {
	var city string = espnGameDetails.GameInfo.Venue.Address.City
	return city
}

func ParseState(espnGameDetails pickemstructs.ESPNGameDetailsResponse) string {
	var state = espnGameDetails.GameInfo.Venue.Address.State
	return state
}

func ParseZipcode(espnGameDetails pickemstructs.ESPNGameDetailsResponse) string {
	var zipcode = espnGameDetails.GameInfo.Venue.Address.Zipcode
	return zipcode
}

func ParseCountry(espnGameDetails pickemstructs.ESPNGameDetailsResponse) string {
	var country = espnGameDetails.GameInfo.Venue.Address.Country
	return country
}
