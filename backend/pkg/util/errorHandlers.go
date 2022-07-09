package util

import (
	"log"
	"os"
)

func FatalIfErr(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

}
