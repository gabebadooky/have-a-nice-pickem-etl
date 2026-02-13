// Package utils provides URL constants for CBS Sports endpoints used in the ETL pipeline.
// These constants define the base URLs and URL patterns for accessing CBS Sports data
// including schedules, team information, and statistics for both college football (CFB)
// and NFL.
package utils

// CBS Base URL
const CBS_BASE_URL string = "https://www.cbssports.com"

// Regular Season Schedule URL Base
const CBS_CFB_REGULAR_SEASON_SCHEDULE_URL string = "https://www.cbssports.com/college-football/odds/FBS/" + SEASON_YEAR + "/regular/week-"
const CBS_NFL_REGULAR_SEASON_SCHEDULE_URL string = "https://www.cbssports.com/nfl/odds/" + SEASON_YEAR + "/regular/week-"

// Posteason Schedule URL Base
const CBS_CFB_POST_SEASON_SCHEDULE_URL string = "https://www.cbssports.com/college-football/odds/FBS/" + SEASON_YEAR + "/postseason/bowls-"
const CBS_NFL_POST_SEASON_SCHEDULE_URL string = "https://www.cbssports.com/nfl/odds/" + SEASON_YEAR + "/postseason/"

// Teams URL Base
const CBS_CFB_ALL_TEAMS_PAGE_URL string = "https://www.cbssports.com/college-football/teams/"
const CBS_NFL_ALL_TEAMS_PAGE_URL string = "https://www.cbssports.com/nfl/teams/"

// Team Stats URL Suffix
const CBS_TEAM_STATS_URL_SUFFIX string = "stats/"
