package main

import (
	"errors"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf"
)

type region struct {
	name string
	url  string
}

func Regions(country string) (r []region, e error) {
	b := surf.NewBrowser()
	e = b.Open(countryUrl(country))
	if e != nil {
		return
	}
	b.Find("a[href*=\"regions\"]").Each(func(_ int, s *goquery.Selection) {
		if url, hasUrl := s.Attr("href"); hasUrl {
			rg := region{name: s.Text(), url: url}
			r = append(r, rg)
		} else {
			e = errors.New("Region link missing href attribute")
			return
		}
	})
	return
}

const countryUrlBase = "https://www.tide-forecast.com/countries/$/regions"

func countryUrl(country string) string {
	country = strings.Replace(country, " ", "-", -1)
	return strings.Replace(countryUrlBase, "$", country, -1)
}
