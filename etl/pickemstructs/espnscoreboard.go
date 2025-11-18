package pickemstructs

type ESPNScheduleResponse struct {
	Events []Event `json:"events"`
}

type Event struct {
	ID string `json:"id"`
}
