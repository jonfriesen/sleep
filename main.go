package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

var defaultDelay = time.Second * 5

func main() {

	args := os.Args[1:]

	if len(args) > 0 {
		i, err := strconv.Atoi(args[0])
		if err != nil {
			log.Println("Argument ", args[0], "NaN")
			return
		}
		defaultDelay = time.Second * time.Duration(int64(i))
	}

	log.Println("Sleeping for", defaultDelay)
	time.Sleep(defaultDelay)
}
