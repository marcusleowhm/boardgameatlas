package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/marcusleowhm/boardgameatlas/api"
)

//go run . to run the file without building
//end goal is to be able to run the CLI by passing in flags
func main() {
	//boardgameatlas --query "boardname" --clientId "123" --skip 10 --limit 10
	//Define the command line argument
	query := flag.String("query", "", "Boardgame name to search")
	clientId := flag.String("clientId", "", "Boardgame Atlas client_id")

	skip := flag.Uint("skip", 0, "Skips the number of results provided")
	limit := flag.Uint("limit", 0, "Limits the number of results returned")

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

	bga := api.New(*clientId)
	bga.Search(context.Background(), *query, *limit, *skip)
}

func isNull(s string) bool {
	return len(strings.TrimSpace(s)) <= 0
}