package main

import (
	"fmt"
	"have-a-nice-pickem-etl/etl/extract"
	"have-a-nice-pickem-etl/etl/transform/gamedetails"
)

func main() {
	games := extract.ExtractGames(extract.CfbGamesExtract{Week: 16})

	for i := range 2 {
		// fmt.Printf("Game: %v", games[i])
		gamedetails := gamedetails.InstantiateGameDetails(gamedetails.New{Game: games[i]})
		fmt.Printf("GameDetails: %v", gamedetails)
	}
}
