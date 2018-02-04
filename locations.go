package main

import (
	"errors"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf"
)

type location struct {
	name string
	url  string
}

func Locations(region string) (loc []location, e error) {
	b := surf.NewBrowser()
	e = b.Open(regionUrl(region))
	if e != nil {
		return
	}
	b.Find("a[href*=\"tides/latest\"]").Each(func(_ int, s *goquery.Selection) {
		if url, hasUrl := s.Attr("href"); hasUrl {
			ll := location{name: s.Text(), url: url}
			loc = append(loc, ll)
		} else {
			e = errors.New("Location link missing href attribute")
			return
		}
	})
	return
}

const regionUrlBase = "https://www.tide-forecast.com/regions/"

func regionUrl(region string) string {
	region = strings.Replace(region, " ", "-", -1)
	return regionUrlBase + region
}
