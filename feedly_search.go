// Copyright 2016 Egor Smolyakov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package feedly_search

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

const SEARCH_API_URL = "https://cloud.feedly.com/v3/search/feeds"

type feed struct {
	FeedID      string `json:"feedId"`
	Subscribers int    `json:"subscribers"`
	Title       string `json:"title"`
	Website     string `json:"website"`
}

type response struct {
	Results []feed   `json:"results"`
	Hint    string   `json:"hint,omitempty"`
	Related []string `json:"related,omitempty"`
}

func Process(query, locale string, number int) {
	resp, err := http.Get(constructURL(query, locale, number))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var r response
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Fatalln(err)
	}

	if r.Hint != "" {
		fmt.Println("Feedly hint:", r.Hint)
	}

	if len(r.Related) != 0 {
		fmt.Println("Feedly related queries:", strings.Join(r.Related, ", "))
		fmt.Println()
	}

	if len(r.Results) == 0 {
		fmt.Println("No results.")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Title", "Subscribers", "RSS"})
	table.SetRowLine(true)

	// Generate feeds list
	for _, feed := range r.Results {
		data := []string{
			title(feed.Title, feed.Website),
			subscribers(feed.Subscribers),
			rssUrl(feed.FeedID),
		}
		table.Append(data)
	}

	table.Render()
}

func title(feedTitle, feedWebsite string) string {
	r := ""

	feedTitle = strings.TrimSpace(feedTitle)
	if len(feedTitle) > 0 {
		r = r + feedTitle
	}
	if len(feedWebsite) > 0 {
		r = r + ", " + feedWebsite
	}
	return r
}

func rssUrl(feedID string) string {
	return strings.Replace(feedID, "feed/", "", 1)
}

func subscribers(count int) string {
	if s := strconv.Itoa(count); s != "" {
		return s
	}

	return "-"
}

func constructURL(query, locale string, number int) string {
	req, err := http.NewRequest("GET", SEARCH_API_URL, nil)
	if err != nil {
		log.Fatalln(err)
	}

	q := req.URL.Query()
	q.Add("query", query)
	q.Add("locale", locale)
	q.Add("count", strconv.Itoa(number))
	req.URL.RawQuery = q.Encode()

	return req.URL.String()
}
