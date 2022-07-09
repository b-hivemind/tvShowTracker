package tvmazeapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"fmt"
	"net/http"
)

/*
func auditSerialNumbers(episodes []Episode) {
	for s := 0; s < len(episodes); s++ {
		episodes[s].Serial = s + 1
	}
}
*/

func GetEpisodeInfoByID(episodeID int) (Episode, error) {
	var result Episode
	queryURL := fmt.Sprintf(TVMAZE_EPISODE_BY_ID_API, episodeID)
	resp, err := http.Get(queryURL)
	if err != nil {
		return result, err
	}

	if resp.StatusCode >= 400 {
		return result, errors.New(fmt.Sprintf("%d: %s", resp.StatusCode, resp.Status))
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func getEpisodeInfoByNumber(showID int, seasonNumber int, episodeNumber int) (Episode, error) {
	var result Episode
	queryURL := fmt.Sprintf(TVMAZE_EPISODE_BY_NUMBER_API, showID, seasonNumber, episodeNumber)
	resp, err := http.Get(queryURL)
	if err != nil {
		return result, err
	}

	if resp.StatusCode >= 400 {
		return result, errors.New(fmt.Sprintf("%d: %s", resp.StatusCode, resp.Status))
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
