// Copyright 2016 Egor Smolyakov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"flag"
	"os"

	"github.com/egorsmkv/feedly_search"
)

var (
	query, locale string
	number        int
)

func init() {
	flag.StringVar(&query, "q", "", "search `query`, can be a feed url, a site title, a site url or a #topic")
	flag.IntVar(&number, "n", 20, "`number` of results, minimum 1, maximum: 500")
	flag.StringVar(&locale, "l", "en", "hint the search engine to return feeds in that `locale` (e.g. pt, fr_FR)")
}

func main() {
	flag.Parse()

	if query == "" || number <= 0 || number > 500 {
		flag.Usage()
		os.Exit(1)
	}

	feedly_search.Process(query, locale, number)
}
