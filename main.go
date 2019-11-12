package gourban

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var baseURL = "https://api.urbandictionary.com/v0/define?term="

type Entry struct {
	Definition string    `json:"definition"`
	Permalink  string    `json:"permalink"`
	Upvotes    int       `json:"thumbs_up"`
	Downvotes  int       `json:"thumbs_down"`
	Word       string    `json:"word"`
	Defid      int       `json:"defid"`
	Date       time.Time `json:"written_on"`
	Example    string    `json:"example"`
}

func Query(s string) ([]Entry, error) {
	queryURL := baseURL + url.QueryEscape(s)
	res, err := http.Get(queryURL)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var qres map[string][]Entry
	if err := json.Unmarshal(body, &qres); err != nil {
		return nil, err
	}
	return qres["list"], nil
}
