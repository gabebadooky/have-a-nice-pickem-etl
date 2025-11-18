package pickemstructs

type Boxscore struct {
	GameID        string
	TeamID        string
	Q1Score       uint8
	Q2Score       uint8
	Q3Score       uint8
	Q4Score       uint8
	OvertimeScore uint8
	TotalScore    uint8
}
