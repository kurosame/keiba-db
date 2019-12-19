package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"default/domain/model"
	"default/domain/service"
)

const baseURL = "https://db.netkeiba.com"
const term = 3

var now time.Time

func getEventDateLinks() []string {
	var links []string
	for i := 0; i < term; i++ {
		time.Sleep(1 * time.Second)

		target := now.AddDate(0, -i, 0)
		url := fmt.Sprintf("%s/?pid=race_top&date=%d%d", baseURL, target.Year(), target.Month())
		doc := service.GetHTMLDoc(url)
		links = append(links, service.GetLinks(doc, `^/race/list/\d+/$`)...)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(links)))

	return service.UniqStr(links)
}

func getRaceLinks(eventDateLinks []string) []string {
	var links []string
	for _, l := range eventDateLinks {
		time.Sleep(1 * time.Second)

		doc := service.GetHTMLDoc(baseURL + l)
		links = append(links, service.GetLinks(doc, `^/race/\d+/$`)...)
	}

	return links
}

func main() {
	now = time.Now()

	// eventDateLinks := getEventDateLinks()
	// raceLinks := getRaceLinks(eventDateLinks)

	doc := service.GetHTMLDoc("https://db.netkeiba.com/race/201910021212/")
	texts := service.GetTexts(doc, ".race_table_01 > tbody > tr")

	var raceResults []*model.RaceResult
	for i, t := range texts {
		// Exclude header
		if i == 0 {
			continue
		}

		tr := strings.NewReplacer("</td>\n", "</td>\n", "\n", "").Replace(t)
		ts := service.SanitizeHTML(tr)

		raceResult := &model.RaceResult{
			Rank:         ts[0],
			FrameNumber:  ts[1],
			HorseNumber:  ts[2],
			HorseName:    ts[3],
			Age:          ts[4],
			Weight:       ts[5],
			JockeyName:   ts[6],
			Time:         ts[7],
			Gap:          ts[8],
			TimeIndex:    ts[9],
			Passing:      ts[10],
			Rise:         ts[11],
			Odds:         ts[12],
			Popular:      ts[13],
			HorseWeight:  ts[14],
			TrainingTime: ts[15],
			Comment:      ts[16],
			Remarks:      ts[17],
			TrainerName:  ts[18],
			OwnerName:    ts[19],
			Prize:        ts[20]}
		raceResults = append(raceResults, raceResult)
	}
	race := &model.Race{RaceResults: raceResults}

	service.OutputJSONL([]string{race.ToJSON()})
}
