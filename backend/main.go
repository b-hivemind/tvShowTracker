package main

import (
	"flag"
	"strings"

	"github.com/b-hivemind/preparer/pkg/api"
	"github.com/b-hivemind/preparer/pkg/db"
)

func main() {
	mode := flag.String("mode", "live", "Execution mode")
	flag.Parse()
	if strings.ToLower(*mode) == "test" {
		db.DatabaseName = "tvShowTrackerTest"
	}
	api.Start()
}
