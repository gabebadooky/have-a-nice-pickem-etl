// Package record provides conference record parsing functionality.
// It extracts conference win-loss-tie records from ESPN Team Summary API responses.
package record

import (
	"have-a-nice-pickem-etl/internal/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func scrapeConferenceRecordText(cbsTeamStatsPage *goquery.Selection) string {
	recordText := cbsTeamStatsPage.Find("div.GlobalSubnav_overviewWrap__kQuwK").Find("ul.GlobalSubnav_overview__7CYn6").Find("li.GlobalSubnav_overviewItem__KtQsg").Last().Text()
	return recordText
}

func (c ConferenceRecord) parseWins() uint {
	conferenceRecordText := scrapeConferenceRecordText(c.CBS)
	conferenceWinsString := strings.Split(conferenceRecordText, "-")[0]
	var conferenceWins int = utils.ConvertStringToInt(conferenceWinsString)
	return uint(conferenceWins)
}

func (c ConferenceRecord) parseLosses() uint {
	conferenceRecordText := scrapeConferenceRecordText(c.CBS)
	conferenceLossesString := strings.Split(conferenceRecordText, "-")[1]
	var conferenceLosses int = utils.ConvertStringToInt(conferenceLossesString)
	return uint(conferenceLosses)
}

func (c ConferenceRecord) parseTies() uint {
	var conferenceTiesString string
	conferenceRecordText := scrapeConferenceRecordText(c.CBS)
	conferenceTiesSlice := strings.Split(conferenceRecordText, "-")

	if len(conferenceTiesSlice) == 3 {
		conferenceTiesString = conferenceTiesSlice[2]
	} else {
		conferenceTiesString = "0"
	}

	var conferenceTies int = utils.ConvertStringToInt(conferenceTiesString)
	return uint(conferenceTies)
}
