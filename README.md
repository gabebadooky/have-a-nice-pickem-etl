# have-a-nice-pickem-etl
Rewrite of the Pickem Application ETL scripts in GoLang

## Project Structure
```
+-- go.mod
+-- go.sum
+-- LICENSE
+-- main.go
+-- etl
|   +-- extract
|   |   +-- cbs
|   |   |   +-- schedule.go
|   |   +-- espn
|   |   |   +-- game.go
|   |   |   +-- team.go
|   |   +-- fox
|   +-- transform
|   |   +-- common.go
|   |   +-- game.go
|   +-- load
|   +-- types
|   |   +-- espngame.go
```