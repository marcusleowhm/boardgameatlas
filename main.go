package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/marcusleowhm/boardgameatlas/api"
	"github.com/fatih/color"
)

// go run . to run the file without building
// end goal is to be able to run the CLI by passing in flags
func main() {
	//boardgameatlas --query "boardname" --clientId "123" --skip 10 --limit 10
	//Define the command line argument
	query := flag.String("query", "", "Boardgame name to search")
	clientId := flag.String("clientId", "", "Boardgame Atlas client_id")

	skip := flag.Uint("skip", 0, "Skips the number of results provided")
	limit := flag.Uint("limit", 10, "Limits the number of results returned")
	timeout := flag.Uint("timeout", 10, "Timeout")

	//Parse the command line
	flag.Parse()

	//Make sure that --query and --clientId are set
	if isNull(*query) {
		log.Fatalln("Please use --query to set a boardgame name to search")
	}
	if isNull(*clientId) {
		log.Fatalln("PLease use --clientId to set the clientId")
	}
	fmt.Printf("query=%s, clientId=%s, limit=%d, skip=%d\n", *query, *clientId, *limit, *skip)

	//Create a new istance of Boardgame Atlas client
	bga := api.New(*clientId)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeout*uint(time.Second)))
	defer cancel() //Call cancel when program exits

	//Make the invocation
	result, err := bga.Search(ctx, *query, *limit, *skip)
	if err != nil {
		log.Fatalf("Cannot search for boardgame: %v", err)
	}

	boldGreen := color.New(color.Bold).Add(color.FgHiGreen).SprintFunc()
	boldBlue := color.New(color.Bold).Add(color.FgBlue).SprintFunc()
	boldRed := color.New(color.Bold).Add(color.FgRed).SprintFunc()

	for _, g := range result.Games {
		fmt.Printf("Name: %s\n", boldGreen(g.Name))
		fmt.Printf("Description: %s\n", boldBlue(g.Description))
		fmt.Printf("Url: %s\n\n", boldRed(g.Url))
	}
}

func isNull(s string) bool {
	return len(strings.TrimSpace(s)) <= 0
}
