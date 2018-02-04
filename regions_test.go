package main

import "testing"

func TestRegions(t *testing.T) {
	var r []region
	r, e := Regions("United States")
	if e != nil {
		t.Fatal(e)
	}
	if 0 >= len(r) {
		t.Fatal("No regions found")
	}
	if r[0].name != "Alabama" {
		t.Errorf("Regions(%s)[0].name => %s, want %s", "United States", r[0].name, "Alabama")
	}
	if r[0].url != "/regions/Alabama" {
		t.Errorf("Regions(%s)[0].url => %s, want %s", "United States", r[0].url, "/regions/Alabama")
	}
}
