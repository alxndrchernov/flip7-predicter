package main

import (
	"flag"
	"flip7/predicter"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		playerCards  string
		removedCards string
	)
	flag.StringVar(&playerCards, "pc", "", "player cards (example: -pc 1,2,3)")
	flag.StringVar(&removedCards, "rc", "", "removed cards from pack (example: -rc 1,2,3)")
	flag.Parse()
	failPercents := predicter.NewPlayer(parseArrayFlag(playerCards)).Predict(parseArrayFlag(removedCards))
	fmt.Fprintf(os.Stdout, "%.2f", failPercents*100)
}

func parseArrayFlag(flag string) []string {
	return strings.Split(flag, ",")
}
