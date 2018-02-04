This scraper finds and prints daylight time and height records for a list of hardcoded tide pool locations.

To build the scraper yourself (if you have already installed Go):

```golang
go get -u github.com/andrew-carroll/tideforecast
cd $GOPATH/github.com/andrew-carroll/tideforecast
go build
```

Then to run the scraper, simply run `./tideforecast`. Note that this scraper does not include paging, so I recommend outputting to either a file or `less`.
