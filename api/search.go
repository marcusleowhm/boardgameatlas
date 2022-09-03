package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

const SEARCH_URL = "https://api.boardgameatlas.com/api/search"

type BoardgameAtlas struct {
	clientId string
}

func New(clientId string) BoardgameAtlas {
	return BoardgameAtlas{ clientId }
}

//Receiver is bga
//if need to change value in struct, pass in pointer *BoardgameAtlas
func (bga BoardgameAtlas) Search(ctx context.Context, query string, limit uint, skip uint) error {
	
	//create http client
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, SEARCH_URL, nil)
	if err != nil {
		return fmt.Errorf("cannot create HTTP client: %v", err)
	}

	//URL encode queries
	//Get the params string object
	qs := req.URL.Query()
	//Populate URL with params
	qs.Add("name", query)
	qs.Add("limit", fmt.Sprintf("%d", limit))
	qs.Add("skip", strconv.Itoa(int(skip)))
	qs.Add("client_id", bga.clientId)

	//Encode query params and add to URL
	req.URL.RawQuery = qs.Encode()

	fmt.Printf("URL = %s\n", req.URL.String())

	return nil
}