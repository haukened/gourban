/*
Pacakge gourban provides a golang interface to the urbandictionary.com API.
It can be imported into your codebase and used like `gourban.Query(string)` or as a standalone go binary for fun.
*/
package gourban

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var baseURL = "https://api.urbandictionary.com/v0/define?term="

// An Entry represents a single definition from urban dictionary.  Each word can have multile entries.
// number of up/down votes can be used to sort definitions to most (un)popular if desired.
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

// Query accepts a string (one word or multiples) and is URL encoded and fed to the urbandictionary API
// queries that do not return results return an empty slice.  Upstream errors are propogated and returned, so many types may exist.
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