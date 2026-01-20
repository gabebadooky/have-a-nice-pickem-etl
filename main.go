package main

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/extract"
	"have-a-nice-pickem-etl/etl/transform/record"
)

func main() {
	/*games := extract.ExtractGames(extract.CfbGamesExtract{Week: 16})
	var distinctTeams []string

	for i := range 4 {
		// fmt.Printf("Game: %v", games[i])
		gamedetailsrow := gamedetails.New{Game: games[i]}.InstantiateGameDetails()
		//espnBettingOdds := bettingodds.InstantiateBettingOdds(bettingodds.EspnBettingOdds{Game: games[i]})
		//cbsBettingOdds := bettingodds.InstantiateBettingOdds(bettingodds.CbsBettingOdds{Game: games[i]})
		//awayBoxScore := boxscore.InstantiateBoxscore(boxscore.AwayBoxscore{Game: games[i]})
		//homeBoxScore := boxscore.InstantiateBoxscore(boxscore.HomeBoxscore{Game: games[i]})

		if !slices.Contains(distinctTeams, gamedetailsrow.AwayTeamID) {
			distinctTeams = append(distinctTeams, gamedetailsrow.AwayTeamID)
		}
		if !slices.Contains(distinctTeams, gamedetailsrow.HomeTeamID) {
			distinctTeams = append(distinctTeams, gamedetailsrow.HomeTeamID)
		}
	}*/

	teams := extract.ExtractTeams(extract.CfbTeamsExtract{Week: 16})

	for i := range len(teams) {
		//teamdetailsrow := teamdetails.New{Team: teams[i]}.Instantiate()
		//conferenceRecord := record.InstantiateRecord(record.ConferenceRecord{Team: teams[i]})
		overallRecord := record.InstantiateRecord(record.OverallRecord{Team: teams[i]})

		fmt.Println()
		/*fmt.Printf(`
			TeamID: %s
			League: %s
			ESPNCode: %s
			CBSCode: %s
			FoxCode: %s
			VegasCode: %s
			ConferenceID: %s
			Name: %s
			Mascot: %s
			PrimaryColor: %s
			AlternateColor: %s
			Ranking: %d
		`,
			teamdetailsrow.TeamID,
			teamdetailsrow.League,
			teamdetailsrow.ESPNCode,
			teamdetailsrow.CBSCode,
			teamdetailsrow.FoxCode,
			teamdetailsrow.VegasCode,
			teamdetailsrow.ConferenceID,
			teamdetailsrow.Name,
			teamdetailsrow.Mascot,
			teamdetailsrow.PrimaryColor,
			teamdetailsrow.AlternateColor,
			teamdetailsrow.Ranking,
		)*/
		/*fmt.Printf(`
			TeamID: %s
			RecordType: %s
			Wins: %d
			Losses: %d
			Ties: %d
		`,
			conferenceRecord.TeamID,
			conferenceRecord.RecordType,
			conferenceRecord.Wins,
			conferenceRecord.Losses,
			conferenceRecord.Ties,
		)*/
		fmt.Printf(`
			TeamID: %s
			RecordType: %s
			Wins: %d
			Losses: %d
			Ties: %d
		`,
			overallRecord.TeamID,
			overallRecord.RecordType,
			overallRecord.Wins,
			overallRecord.Losses,
			overallRecord.Ties,
		)
		fmt.Println()
		fmt.Println()
	}
}
