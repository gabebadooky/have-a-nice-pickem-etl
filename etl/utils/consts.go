package utils

const OPENCAGE_GEOCODE_ENDPOINT_URL string = "https://api.opencagedata.com/geocode/v1/json"

const SEASON_YEAR string = "2025"
const CFB_REG_SEASON_WEEKS uint = 14
const NFL_REG_SEASON_WEEKS uint = 18

const FOX_CFB_REGULER_SEASON_SCHEDULE_URL string = "https://www.foxsports.com/college-football/schedule?groupId=2&seasonType=reg&week="
const FOX_CFB_POST_SEASON_SCHEDULE_URL string = "https://www.foxsports.com/college-football/schedule?groupId=2&seasonType=post&week="
const FOX_NFL_REGULAR_SEASON_SCHEDULE_URL string = "https://www.foxsports.com/nfl/schedule?groupId=2&seasonType=reg&week="
const FOX_NFL_POST_SEASON_SCHEDULE_URL string = "https://www.foxsports.com/nfl/schedule?groupId=2&seasonType=post&week="

const FOX_GAME_BASE_URL string = "https://www.foxsports.com"
const FOX_CFB_GAME_URL string = "https://www.foxsports.com/college-football"
const FOX_NFL_GAME_URL string = "https://www.foxsports.com/nfl"
const FOX_GAME_STATS_URL_SUFFIX string = "?tab=gamestats"

const ESPN_CFB_REGULAR_SEASON_SCHEDULE_URL string = "https://site.api.espn.com/apis/site/v2/sports/football/college-football/scoreboard?groups=80&year=" + SEASON_YEAR + "&seasonType=2&week="
const ESPN_CFB_POST_SEASON_SCHEDULE_URL string = "https://site.api.espn.com/apis/site/v2/sports/football/college-football/scoreboard?groups=80&year=" + SEASON_YEAR + "&seasonType=3"
const ESPN_NFL_REGULAR_SEASON_SCHEDULE_URL string = "https://site.api.espn.com/apis/site/v2/sports/football/nfl/scoreboard?seasontype=2&dates=" + SEASON_YEAR + "&week="
const ESPN_NFL_POST_SEASON_SCHEDULE_URL string = "https://site.api.espn.com/apis/site/v2/sports/football/nfl/scoreboard?seasontype=2&dates=" + SEASON_YEAR + "&week="

const ESPN_CFB_GAME_ENDPOINT_URL string = "https://site.api.espn.com/apis/site/v2/sports/football/college-football/summary?event="
const ESPN_NFL_GAME_ENDPOINT_URL string = "https://site.api.espn.com/apis/site/v2/sports/football/nfl/summary?event="
const ESPN_CFB_TEAM_ENDPOINT_URL string = "https://site.api.espn.com/apis/site/v2/sports/football/college-football/teams/"
const ESPN_NFL_TEAM_ENDPOINT_URL string = "https://site.api.espn.com/apis/site/v2/sports/football/nfl/teams/"

const CBS_BASE_URL string = "https://www.cbssports.com"
const CBS_CFB_SCHEDULE_URL string = "https://www.cbssports.com/college-football/odds/FBS"
const CBS_NFL_SCHEDULE_URL string = ""

const CBS_CFB_ALL_TEAMS_PAGE_URL string = "https://www.cbssports.com/college-football/teams/"
const CBS_NFL_ALL_TEAMS_PAGE_URL string = "https://www.cbssports.com/nfl/teams/"

const CBS_CFB_REGULAR_SEASON_SCHEDULE_URL string = "https://www.cbssports.com/college-football/odds/FBS/" + SEASON_YEAR + "/regular/week-"
const CBS_CFB_POST_SEASON_SCHEDULE_URL string = "https://www.cbssports.com/college-football/odds/FBS/" + SEASON_YEAR + "/postseason/bowls-"

const CBS_NFL_REGULAR_SEASON_SCHEDULE_URL string = "https://www.cbssports.com/nfl/odds/" + SEASON_YEAR + "/regular/week-"
const CBS_NFL_POST_SEASON_SCHEDULE_URL string = "https://www.cbssports.com/nfl/odds/" + SEASON_YEAR + "/postseason/"
