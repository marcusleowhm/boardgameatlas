package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const SEARCH_URL = "https://api.boardgameatlas.com/api/search"

type BoardgameAtlas struct {
	clientId string
}

// Game
type Game struct {
	Id            string `json:"id"` //Automapping
	Name          string `json:"name"`
	Price         string `json:"price"`
	YearPublished uint   `json:"year_published"`
	Description   string `json:"description"`
	Url           string `json:"official_url"`
	ImageUrl      string `json:"image_url"`
	RulesUrl      string `json:"rules_url"`
}

type SearchResult struct {
	Games []Game `json:"games"`
	Count uint   `json:"count"`
}

func New(clientId string) BoardgameAtlas {
	return BoardgameAtlas{clientId}
}

// Receiver is bga
// if need to change value in struct, pass in pointer *BoardgameAtlas
func (bga BoardgameAtlas) Search(ctx context.Context, query string, limit uint, skip uint) (*SearchResult, error) {

	//create http client
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, SEARCH_URL, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot create HTTP client: %v", err)
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

	//Make the call
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Cannot create HTTP client for invocation: %v", err)
	}

	//HTTP status code > 400
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("Error HTTP status: %s", resp.Status)
	}


	//Deserialize object
	var result SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("cannot deserialize JSON payload: %v", err)
	}
	
	return &result, nil
}
