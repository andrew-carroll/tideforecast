package main

import "testing"

func TestLocations(t *testing.T) {
	var loc []location
	loc, e := Locations("California")
	if e != nil {
		return
	}
	if 0 >= len(loc) {
		t.Fatal("No locations found")
	}
	want, got := "San Diego", loc[0].name
	if got != want {
		t.Errorf("Locations(%s)[0].name => %s, want %s", "California", got, want)
	}
	want, got = "/locations/San-Diego-California/tides/latest", loc[0].url
	if got != want {
		t.Errorf("Locations(%s)[0].url => %s, want %s", "California", got, want)
	}
}
