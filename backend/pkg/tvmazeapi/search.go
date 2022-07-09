package tvmazeapi

import (
	"github.com/b-hivemind/preparer/pkg/util"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type MutliSearchResult struct {
	Score float32 `json:"score"`
	Show  `json:"show"`
}

func SearchMultiShows(title string) ([]Show, error) {
	var results []MutliSearchResult
	var shows []Show

	resp, err := http.Get(TVMAZE_MULTI_SEARCH_API + url.PathEscape(title))
	util.FatalIfErr(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	util.FatalIfErr(err)

	err = json.Unmarshal(body, &results)

	for _, result := range results {
		shows = append(shows, result.Show)
	}

	return shows, nil
}

func SearchSingleShow(title string) (Show, error) {
	var show Show
	resp, err := http.Get(TVMAZE_SINGLE_SEARCH_API + url.PathEscape(title))
	util.FatalIfErr(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	util.FatalIfErr(err)

	err = json.Unmarshal(body, &show)
	util.FatalIfErr(err)

	fmt.Println(show)

	return show, nil
}
