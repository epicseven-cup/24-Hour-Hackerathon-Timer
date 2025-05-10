package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strings"
	"time"
)

var hour int
var round bool
var wordlist []string

func init() {
	flag.IntVar(&hour, "h", 24, "Count down hours (shorthand)")
	flag.IntVar(&hour, "hour", 24, "Count down hours")
	flag.BoolVar(&round, "r", false, "Round hours to closest hour for start time, round down (shorthand)")
	flag.BoolVar(&round, "round", false, "Round hours to closest hour for start time, round down")
}

func hackerathonTimer(h int, r bool) error {
	c := time.NewTicker(time.Hour)
	startTime := time.Now()

	if r {
		startTime = startTime.Round(time.Hour)
	}

	hoursPassed := 0

	for next := range c.C {
		if next.After(startTime) {
			index := rand.IntN(len(wordlist))
			fmt.Printf("It been %v hours for the hackerathon\n", next)
			fmt.Printf("Current Random world of the hour: %s \n", wordlist[index])
			hoursPassed++
		}

		if hoursPassed >= h {
			return nil
		}
	}
	return fmt.Errorf("could not reach hackerathon")
}

func main() {
	data, err := os.ReadFile("/usr/share/dict/words")
	if err != nil {
		log.Fatal(err)
		return
	}
	wordlist = strings.Split(string(data), "\n")

	err = hackerathonTimer(hour, round)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Hackerathon timer finished.")
}
