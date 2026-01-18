package main

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/extract"
	"have-a-nice-pickem-etl/etl/transform/gamedetails"
)

func main() {
	games := extract.ExtractGames(extract.CfbGamesExtract{Week: 16})

	for i := range 4 {
		// fmt.Printf("Game: %v", games[i])
		gamedetails := gamedetails.InstantiateGameDetails(gamedetails.New{Game: games[i]})
		fmt.Println()
		fmt.Printf(`
			GameID: %s,
			ESPN Code: %s, 
			CBS Code: %s, 
			FOX Code: %s, 
			League: %s,
			Week: %d,
			Year: %d,
			AwayTeamID: %s,
			HomeTeamID: %s,
			ZuluTimestamp: %s,
			Broadcast: %s,
			Location: %s,
			Finished: %v,
			`,
			gamedetails.GameID,
			gamedetails.EspnCode,
			gamedetails.CbsCode,
			gamedetails.FoxCode,
			gamedetails.League,
			gamedetails.Week,
			gamedetails.Year,
			gamedetails.AwayTeamID,
			gamedetails.HomeTeamID,
			gamedetails.ZuluTimestamp,
			gamedetails.Broadcast,
			gamedetails.Location,
			gamedetails.Finished,
		)
		fmt.Println()
		fmt.Println()
	}
}
