package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf"
	"github.com/headzoo/surf/browser"
)

var input = []string{
	"Half Moon Bay, California",
	"Huntington Beach, California",
	"Providence, Rhode Island",
	"Wrightsville Beach, North Carolina",
}

func main() {
	b := surf.NewBrowser()
	for _, l := range input {
		printDaylightLowTides(b, l)
	}
}

const locationBaseUrl = "https://www.tide-forecast.com/locations/$/tides/latest"

func guessLocationByWholeString(l string) string {
	l = strings.Replace(l, ",", "", -1)
	l = strings.Replace(l, " ", "-", -1)
	l = strings.Replace(locationBaseUrl, "$", l, -1)
	return l
}

func guessLocationBeforeComma(l string) string {
	l = strings.TrimRight(l, ",")
	l = strings.Replace(l, "$", locationBaseUrl, -1)
	return l
}

func printDaylightLowTides(b *browser.Browser, l string) {
	var daylight bool
	e := b.Open(guessLocationByWholeString(l))
	if e != nil {
		e := b.Open(guessLocationBeforeComma(l))
		if e != nil {
			panic(e)
		}
	}
	fmt.Printf("Daylight low tide times and heights for %s\n========\n", l)
	b.Find("section.tide-events > table tr").Each(func(_ int, tr *goquery.Selection) {
		if "Sunrise" == tr.Find("td:last-child").Text() {
			daylight = true
			return
		}
		if "Sunset" == tr.Find("td:last-child").Text() {
			daylight = false
			return
		}
		if daylight && "Low Tide" == tr.Find("td:last-child").Text() {
			time := tr.Find("td.time").Text()
			height := tr.Find("td.level.metric").Text()
			fmt.Printf("%s: %s\n", time, height)
		}
	})
	fmt.Println()
}
