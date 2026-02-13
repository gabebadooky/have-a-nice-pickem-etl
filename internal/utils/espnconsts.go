// Package utils provides URL constants for ESPN API endpoints used in the ETL pipeline.
// These constants define the API endpoint URLs for accessing ESPN data including schedules,
// game summaries, and team information for both college football (CFB) and NFL.
package utils

// Regular Season Schedule URL Base
const ESPN_CFB_REGULAR_SEASON_SCHEDULE_URL string = "https://site.api.espn.com/apis/site/v2/sports/football/college-football/scoreboard?groups=80&year=" + SEASON_YEAR + "&seasonType=2&week="
const ESPN_NFL_REGULAR_SEASON_SCHEDULE_URL string = "https://site.api.espn.com/apis/site/v2/sports/football/nfl/scoreboard?seasontype=2&dates=" + SEASON_YEAR + "&week="

// Postseason Schedule URL Base
const ESPN_CFB_POST_SEASON_SCHEDULE_URL string = "https://site.api.espn.com/apis/site/v2/sports/football/college-football/scoreboard?groups=80&year=" + SEASON_YEAR + "&seasonType=3"
const ESPN_NFL_POST_SEASON_SCHEDULE_URL string = "https://site.api.espn.com/apis/site/v2/sports/football/nfl/scoreboard?seasontype=2&dates=" + SEASON_YEAR + "&week="

// Games URL Base
const ESPN_CFB_GAME_ENDPOINT_URL string = "https://site.api.espn.com/apis/site/v2/sports/football/college-football/summary?event="
const ESPN_NFL_GAME_ENDPOINT_URL string = "https://site.api.espn.com/apis/site/v2/sports/football/nfl/summary?event="

// Teams URL Base
const ESPN_CFB_TEAM_ENDPOINT_URL string = "https://site.api.espn.com/apis/site/v2/sports/football/college-football/teams/"
const ESPN_NFL_TEAM_ENDPOINT_URL string = "https://site.api.espn.com/apis/site/v2/sports/football/nfl/teams/"
