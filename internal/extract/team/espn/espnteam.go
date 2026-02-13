package espnteam

import (
	"fmt"
	"have-a-nice-pickem-etl/internal/utils"
	"log"
)

type EspnCfbTeam struct {
	TeamCode string
}

type EspnNflTeam struct {
	TeamCode string
}
type espnTeamInstantiator interface {
	getTeamSummary() TeamSummaryEndpoint
}

func GetTeamSummary(t espnTeamInstantiator) TeamSummaryEndpoint {
	return t.getTeamSummary()
}

// Make and handle ESPN Team Summary endpoint call
func fetchTeamEndpointCall(espnTeamEndpoint string) TeamSummaryEndpoint {
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
func (cfb EspnCfbTeam) getTeamSummary() TeamSummaryEndpoint {
	espnTeamEndpoint := fmt.Sprintf("%s%s", utils.ESPN_CFB_TEAM_ENDPOINT_URL, cfb.TeamCode)
	espnTeamSummary := fetchTeamEndpointCall(espnTeamEndpoint)
	return espnTeamSummary
}

// Call ESPN CFB Team Summary API Endpoint for a given team ID
func (nfl EspnNflTeam) getTeamSummary() TeamSummaryEndpoint {
	espnTeamEndpoint := fmt.Sprintf("%s%s", utils.ESPN_NFL_TEAM_ENDPOINT_URL, nfl.TeamCode)
	espnTeamSummary := fetchTeamEndpointCall(espnTeamEndpoint)
	return espnTeamSummary
}
