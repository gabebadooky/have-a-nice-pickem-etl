// Package record provides conference record parsing functionality.
// It extracts conference win-loss-tie records from ESPN Team Summary API responses.
package record

import (
	"have-a-nice-pickem-etl/internal/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// scrapeConferenceRecordText extracts the conference record W-L-T text from the CBS team subnav.
func scrapeConferenceRecordText(cbsTeamStatsPage *goquery.Selection) string {
	recordText := cbsTeamStatsPage.Find("div.GlobalSubnav_overviewWrap__kQuwK").Find("ul.GlobalSubnav_overview__7CYn6").Find("li.GlobalSubnav_overviewItem__KtQsg").Last().Text()
	return recordText
}

// parseWins returns conference wins from the CBS conference record text.
func (c ConferenceRecord) parseWins() uint {
	if c.CBS == nil {
		return 0
	}

	conferenceRecordText := scrapeConferenceRecordText(c.CBS)
	if conferenceRecordText == "" {
		return 0
	}

	conferenceWinsString := strings.Split(conferenceRecordText, "-")[0]
	var conferenceWins int = utils.ConvertStringToInt(conferenceWinsString)
	return uint(conferenceWins)
}

// parseLosses returns conference losses from the CBS conference record text.
func (c ConferenceRecord) parseLosses() uint {
	if c.CBS == nil {
		return 0
	}

	conferenceRecordText := scrapeConferenceRecordText(c.CBS)
	if conferenceRecordText == "" {
		return 0
	}

	conferenceLossesString := strings.Split(conferenceRecordText, "-")[1]
	var conferenceLosses int = utils.ConvertStringToInt(conferenceLossesString)
	return uint(conferenceLosses)
}

// parseTies returns conference ties from the CBS conference record text.
func (c ConferenceRecord) parseTies() uint {
	var conferenceTiesString string
	if c.CBS == nil {
		return 0
	}

	conferenceRecordText := scrapeConferenceRecordText(c.CBS)
	if conferenceRecordText == "" {
		return 0
	}

	conferenceTiesSlice := strings.Split(conferenceRecordText, "-")

	if len(conferenceTiesSlice) == 3 {
		conferenceTiesString = conferenceTiesSlice[2]
	} else {
		conferenceTiesString = "0"
	}

	var conferenceTies int = utils.ConvertStringToInt(conferenceTiesString)
	return uint(conferenceTies)
}
