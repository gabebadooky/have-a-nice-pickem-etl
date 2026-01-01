package utils

const OPENCAGE_GEOCODE_ENDPOINT_URL string = "https://api.opencagedata.com/geocode/v1/json"

const SEASON_YEAR string = "2025"
const CFB_REG_SEASON_WEEKS uint = 14
const NFL_REG_SEASON_WEEKS uint = 18

const FOX_CFB_REGULER_SEASON_SCHEDULE_URL string = "https://www.foxsports.com/college-football/schedule?groupId=2&seasonType=reg&week="
const FOX_CFB_POST_SEASON_SCHEDULE_URL string = "https://www.foxsports.com/college-football/schedule?groupId=2&seasonType=post&week="
const FOX_NFL_REGULAR_SEASON_SCHEDULE_URL string = "https://www.foxsports.com/nfl/schedule?groupId=2&seasonType=reg&week="
const FOX_NFL_POST_SEASON_SCHEDULE_URL string = "https://www.foxsports.com/nfl/schedule?groupId=2&seasonType=post&week="

const ESPN_CFB_REGULAR_SEASON_SCHEDULE_URL string = "https://site.api.espn.com/apis/site/v2/sports/football/college-football/scoreboard?seasonType=2&week="
const ESPN_CFB_POST_SEASON_SCHEDULE_URL string = "https://site.api.espn.com/apis/site/v2/sports/football/college-football/scoreboard?week="
const ESPN_NFL_REGULAR_SEASON_SCHEDULE_URL string = "https://site.api.espn.com/apis/site/v2/sports/football/nfl/scoreboard?seasontype=2&dates=" + SEASON_YEAR + "&week="
const ESPN_NFL_POST_SEASON_SCHEDULE_URL string = "https://site.api.espn.com/apis/site/v2/sports/football/nfl/scoreboard?seasontype=2&dates=" + SEASON_YEAR + "&week="

const ESPN_CFB_GAME_ENDPOINT_URL string = "https://site.api.espn.com/apis/site/v2/sports/football/college-football/summary?event="
const ESPN_NFL_GAME_ENDPOINT_URL string = "https://site.api.espn.com/apis/site/v2/sports/football/nfl/summary?event="
const ESPN_CFB_TEAM_ENDPOINT_URL string = "https://site.api.espn.com/apis/site/v2/sports/football/college-football/teams/"
const ESPN_NFL_TEAM_ENDPOINT_URL string = "https://site.api.espn.com/apis/site/v2/sports/football/nfl/teams/"

const CBS_CFB_SCHEDULE_URL string = "https://www.cbssports.com/college-football/odds/FBS"
const CBS_CFB_REGULAR_SEASON_SCHEDULE_URL string = "https://www.cbssports.com/college-football/odds/FBS/" + SEASON_YEAR + "/regular/week-"
const CBS_CFB_POST_SEASON_SCHEDULE_URL string = "https://www.cbssports.com/college-football/odds/FBS/" + SEASON_YEAR + "/postseason/bowls-"
const CBS_NFL_REGULAR_SEASON_SCHEDULE_URL string = "https://www.cbssports.com/nfl/odds/" + SEASON_YEAR + "/regular/week-"
const CBS_NFL_POST_SEASON_SCHEDULE_URL string = "https://www.cbssports.com/nfl/odds/" + SEASON_YEAR + "/postseason/"
