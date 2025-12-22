/*
Package that:

  - Calls the Team Summary endpoint (espnHiddenTeamSummaryBaseURL),
    for a given ESPN TeamID, from the ESPN Hidden API:
    https://gist.github.com/akeaswaran/b48b02f1c94f873c6655e7129910fc3b

  - Parses and returns the JSON encoded response
*/
package espn

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/utils"
	"log"
)

type EspnTeam interface {
	GetTeamSummary() TeamSummaryEndpoint
}

type EspnCfbTeam struct {
	TeamCode string
}

type EspnNflTeam struct {
	TeamCode string
}

func makeAndHandleTeamEndpointCall(espnTeamEndpoint string) TeamSummaryEndpoint {
	log.Printf("\nCalling ESPN endpoint for Team: %s\n", espnTeamEndpoint)

	body, err := utils.CallEndpoint(espnTeamEndpoint)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	teamDetails, err := utils.DecodeJSON[TeamSummaryEndpoint](body)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	return teamDetails
}

// Call ESPN CFB Team Summary API Endpoint for a given team ID
func (cfb EspnCfbTeam) GetTeamSummary() TeamSummaryEndpoint {
	var espnTeamEndpoint string = fmt.Sprintf("%s%s", utils.ESPN_CFB_TEAM_ENDPOINT_URL, cfb.TeamCode)
	return makeAndHandleTeamEndpointCall(espnTeamEndpoint)
}

// Call ESPN CFB Team Summary API Endpoint for a given team ID
func (nfl EspnNflTeam) GetTeamSummary() TeamSummaryEndpoint {
	var espnTeamEndpoint string = fmt.Sprintf("%s%s", utils.ESPN_NFL_TEAM_ENDPOINT_URL, nfl.TeamCode)
	return makeAndHandleTeamEndpointCall(espnTeamEndpoint)
}
