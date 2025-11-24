package pickemstructs

type TeamSummaryResponse struct {
	Team Team `json:"team"`
}

type Team struct {
	ID             string        `json:"slug"`
	Code           string        `json:"id"`
	Location       string        `json:"location"`
	Name           string        `json:"name"`
	PrimaryColor   string        `json:"color"`
	AlternateColor string        `json:"alternateColor"`
	Logos          []Logo        `json:"logos"`
	OverallRecord  OverallRecord `json:"record"`
	Groups         Groups        `json:"groups"`
	Ranking        uint8         `json:"rank"`
}

type Groups struct {
	ID string `json:"id"`
}

type Logo struct {
	HREF string `json:"href"`
}

type OverallRecord struct {
	RecordItems []RecordItem `json:"items"`
}

type RecordItem struct {
	Summary string `json:"summary"`
}
