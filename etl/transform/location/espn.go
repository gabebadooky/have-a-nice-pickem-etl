package location

import "have-a-nice-pickem-etl/etl/pickemstructs"

func ParseStadium(consolidatedGameProperties pickemstructs.ConsolidatedGameProperties) string {
	var stadium string = consolidatedGameProperties.EspnDetails.GameInfo.Venue.FullName
	return stadium
}

func ParseCity(consolidatedGameProperties pickemstructs.ConsolidatedGameProperties) string {
	var city string = consolidatedGameProperties.EspnDetails.GameInfo.Venue.Address.City
	return city
}

func ParseState(consolidatedGameProperties pickemstructs.ConsolidatedGameProperties) string {
	var state = consolidatedGameProperties.EspnDetails.GameInfo.Venue.Address.State
	return state
}

func ParseZipcode(consolidatedGameProperties pickemstructs.ConsolidatedGameProperties) string {
	var zipcode = consolidatedGameProperties.EspnDetails.GameInfo.Venue.Address.Zipcode
	return zipcode
}

func ParseCountry(consolidatedGameProperties pickemstructs.ConsolidatedGameProperties) string {
	var country = consolidatedGameProperties.EspnDetails.GameInfo.Venue.Address.Country
	return country
}
