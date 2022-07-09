package tvmazeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"

	"github.com/b-hivemind/preparer/pkg/core"
	"github.com/b-hivemind/preparer/pkg/util"
)

type Episode struct {
	ID     int    `json:"id"`
	ShowID int    `json:"showid"`
	Status bool   `json:"status"`
	Serial int    `json:"serial"`
	Season int    `json:"season"`
	Number int    `json:"number"`
	Name   string `json:"name"`
}

type ShowImage struct {
	Medium string `json:"medium"`
}

type ShowEpisodes struct {
	All       map[int]bool `json:"all"`
	Favorites []int        `json:"favorites"`
}

type Show struct {
	ID       int          `json:"id" bson:"_id" uri:"showid" binding:"required"`
	Name     string       `json:"name"`
	Status   string       `json:"status"`
	Image    ShowImage    `json:"image"`
	Seasons  []int        `json:"seasons"`
	Episodes ShowEpisodes `json:"episodes"`
}

func (show Show) GetEpisodeBySerial(serial int) (Episode, error) {
	var epIDs []int
	for k := range show.Episodes.All {
		epIDs = append(epIDs, k)
	}
	sort.SliceStable(epIDs, func(i, j int) bool { return epIDs[i] < epIDs[j] })
	episodeID := epIDs[serial-1]
	result, err := GetEpisodeInfoByID(episodeID)
	if err != nil {
		return result, err
	}
	result.Serial = serial
	result.ShowID = show.ID
	result.Status = show.Episodes.All[result.ID]
	return result, nil
}

func (show Show) GetEpisodeBySeason(season int, episode int) (Episode, error) {
	partial, err := getEpisodeInfoByNumber(show.ID, season, episode)
	if err != nil {
		return partial, err
	}
	partial.Serial = core.GetSerialNumber(partial.ID, show.Episodes.All)
	partial.ShowID = show.ID
	partial.Status = show.Episodes.All[partial.ID]
	return partial, nil
}

func (show Show) GetNextEpisodeID() (int, error) {
	var sortedIDs []int
	for k := range show.Episodes.All {
		sortedIDs = append(sortedIDs, k)
	}
	sort.SliceStable(sortedIDs, func(i, j int) bool { return sortedIDs[i] < sortedIDs[j] })

	for _, v := range sortedIDs {
		if !show.Episodes.All[v] {
			return v, nil
		}
	}
	return -1, errors.New("All Episodes have been watched")
}

func (show Show) GetNextEpisode() (Episode, error) {
	var result Episode
	id, err := show.GetNextEpisodeID()
	if err != nil {
		return result, err
	}
	result, err = GetEpisodeInfoByID(id)
	if err != nil {
		return result, err
	}
	result.Serial = core.GetSerialNumber(result.ID, show.Episodes.All)
	result.ShowID = show.ID
	result.Status = show.Episodes.All[result.ID]
	return result, nil
}

func (show *Show) ToggleEpisode(epID int) {
	state, exists := show.Episodes.All[epID]
	if exists {
		show.Episodes.All[epID] = !state
	} else {
		log.Printf("Episode ID %d not found", epID)
	}
}

func (show *Show) ToggleNextEpisode() {
	if id, err := show.GetNextEpisodeID(); err != nil {
		log.Fatal(err)
	} else {
		show.ToggleEpisode(id)
	}
}

func (show *Show) PopulateEpisodes() {
	var results []Episode
	queryURL := fmt.Sprintf(TVMAZE_ALL_EPISODES_API, show.ID)
	resp, err := http.Get(queryURL)
	util.FatalIfErr(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	util.FatalIfErr(err)

	err = json.Unmarshal(body, &results)
	util.FatalIfErr(err)

	show.Episodes.All = make(map[int]bool)
	for _, ep := range results {
		show.Episodes.All[ep.ID] = false
	}

	show.Seasons = append(show.Seasons, results[0].ID)
	for i := 1; i < len(results); i++ {
		if results[i].Season > results[i-1].Season {
			show.Seasons = append(show.Seasons, results[i].ID)
		}
	}
}
