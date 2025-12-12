/*
Package that:

  - Calls the Team Summary endpoint (espnHiddenTeamSummaryBaseURL),
    for a given ESPN TeamID, from the ESPN Hidden API:
    https://gist.github.com/akeaswaran/b48b02f1c94f873c6655e7129910fc3b
  - Parses and returns the JSON encoded response into `map`
*/
package espn

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/sharedtypes"
	"have-a-nice-pickem-etl/etl/utils"
	"log"
)

type EspnCfbTeam struct {
	TeamID string
}

type EspnNflTeam struct {
	TeamID string
}

func decodeEspnTeamResponse(body []byte) (sharedtypes.ESPNTeamSummaryResponse, error) {
	return utils.DecodeJSON[sharedtypes.ESPNTeamSummaryResponse](body)
}

// Call ESPN CFB Team Summary API Endpoint for a given team ID
func (cfb EspnCfbTeam) GetTeamSummary() sharedtypes.ESPNTeamSummaryResponse {
	var espnTeamEndpoint string = fmt.Sprintf("%s%s", utils.ESPN_CFB_TEAM_ENDPOINT_URL, cfb.TeamID)
	log.Printf("\nCalling ESPN endpoint for CFB Team: %s\n", espnTeamEndpoint)

	body, err := utils.CallEndpoint(espnTeamEndpoint)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	teamDetails, err := decodeEspnTeamResponse(body)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	return teamDetails
}

// Call ESPN CFB Team Summary API Endpoint for a given team ID
func (nfl EspnNflTeam) GetTeamSummary() sharedtypes.ESPNTeamSummaryResponse {
	var espnTeamEndpoint string = fmt.Sprintf("%s%s", utils.ESPN_NFL_TEAM_ENDPOINT_URL, nfl.TeamID)
	log.Printf("\nCalling ESPN endpoint for NFL Team: %s\n", espnTeamEndpoint)

	body, err := utils.CallEndpoint(espnTeamEndpoint)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	teamDetails, err := decodeEspnTeamResponse(body)
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	return teamDetails
}
