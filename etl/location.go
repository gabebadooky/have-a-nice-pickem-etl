func parseStadium(espnGameDetails map[string]any) string {
	stadium, err := espnGameDetails["gameinfo"].(map[string]any)["venue"].(map[string]string)["fullName"]
	if err {
		return ""
	}
	return stadium
}