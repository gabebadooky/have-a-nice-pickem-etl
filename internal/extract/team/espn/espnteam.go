// Package espnteam provides ESPN team summary API client functionality.
// It calls the ESPN Team Summary API endpoint to retrieve team information including
// records, logos, and team details for both college football (CFB) and NFL.
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

// GetTeamSummary runs the given ESPN team instantiator and returns the team summary response.
func GetTeamSummary(t espnTeamInstantiator) TeamSummaryEndpoint {
	return t.getTeamSummary()
}

// fetchTeamEndpointCall calls the ESPN team summary endpoint and decodes the JSON response.
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

// getTeamSummary fetches the ESPN college football team summary for the configured team code.
func (cfb EspnCfbTeam) getTeamSummary() TeamSummaryEndpoint {
	espnTeamEndpoint := fmt.Sprintf("%s%s", utils.ESPN_CFB_TEAM_ENDPOINT_URL, cfb.TeamCode)
	espnTeamSummary := fetchTeamEndpointCall(espnTeamEndpoint)
	return espnTeamSummary
}

// getTeamSummary fetches the ESPN NFL team summary for the configured team code.
func (nfl EspnNflTeam) getTeamSummary() TeamSummaryEndpoint {
	espnTeamEndpoint := fmt.Sprintf("%s%s", utils.ESPN_NFL_TEAM_ENDPOINT_URL, nfl.TeamCode)
	espnTeamSummary := fetchTeamEndpointCall(espnTeamEndpoint)
	return espnTeamSummary
}
