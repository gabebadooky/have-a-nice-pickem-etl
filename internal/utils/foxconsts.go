// Package utils provides URL constants for Fox Sports endpoints used in the ETL pipeline.
// These constants define the base URLs and URL patterns for accessing Fox Sports data
// including schedules, game pages, and statistics for both college football (CFB) and NFL.
package utils

// Fox Base URL
const FOX_BASE_URL string = "https://www.foxsports.com"

// Regular Season Schedule URL Base
const FOX_CFB_REGULAR_SEASON_SCHEDULE_URL string = "https://www.foxsports.com/college-football/schedule?groupId=2&seasonType=reg&week="
const FOX_NFL_REGULAR_SEASON_SCHEDULE_URL string = "https://www.foxsports.com/nfl/schedule?groupId=2&seasonType=reg&week="

// Postseason Schedule URL Base
const FOX_CFB_POST_SEASON_SCHEDULE_URL string = "https://www.foxsports.com/college-football/schedule?groupId=2&seasonType=post&week="
const FOX_NFL_POST_SEASON_SCHEDULE_URL string = "https://www.foxsports.com/nfl/schedule?groupId=2&seasonType=post&week="

// Games URL Base
const FOX_CFB_GAME_URL string = "https://www.foxsports.com/college-football"
const FOX_NFL_GAME_URL string = "https://www.foxsports.com/nfl"

// Game Tabs URL Suffix
const FOX_GAME_ODDS_URL_SUFFIX string = "?tab=odds"
const FOX_GAME_STATS_URL_SUFFIX string = "?tab=gamestats"
const FOX_GAME_BOXSCORE_URL_SUFFIX string = "?tab=boxscore"
