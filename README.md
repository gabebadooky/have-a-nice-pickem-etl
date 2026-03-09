# have-a-nice-pickem-etl
Rewrite of the Pickem Application ETL scripts in GoLang

## ETL Pipeline Design and Development approach
1. Develop `Extract` packages that retrieve desired data fields from various web pages and API endpoints

    i. `schedule`<br>
        - [ESPN Schedule Endpoint](https://site.api.espn.com/apis/site/v2/sports/football/college-football/scoreboard?groups=80&year=2026&seasonType=2&week=1)<br>
        - [CBS Schedule](https://www.cbssports.com/college-football/odds/FBS/2026/regular/week-1)<br>
        - [Fox Schedule](https://www.foxsports.com/college-football/schedule?groupId=2&seasonType=reg&week=1)<br>

    ii. `game`<br>
        - [ESPN Game Endpoint](https://site.api.espn.com/apis/site/v2/sports/football/college-football/scoreboard?event=401754528)<br>
        - [Fox Game](https://www.foxsports.com/college-football/iowa-state-cyclones-vs-kansas-state-wildcats-aug-23-2025-game-boxscore-42830)<br>
    
    iii. `team`<br>
        - [ESPN Team Endpoint](https://site.api.espn.com/apis/site/v2/sports/football/college-football/teams/158)<br>
        - [CBS Team](https://www.cbssports.com/college-football/teams/MIAMI/miami-fla-hurricanes/stats/)<br>
    
    iv. `location`<br>
        - [Opencage Forward Geocode API](https://www.cbssports.com/college-football/teams/MIAMI/miami-fla-hurricanes/stats/)<br>

2. Develop `Transform` packages that manipulate extracted data into required structures