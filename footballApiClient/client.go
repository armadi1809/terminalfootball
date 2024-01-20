package footballApiClient

import (
	"net/http"
	"time"
)

const baseUrl string = "https://api.football-data.org/v4/matches"

type FootballApiClient struct {
	authKey string
	client  *http.Client
}

type Match struct {
	Competition struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Code   string `json:"code"`
		Type   string `json:"type"`
		Emblem string `json:"emblem"`
	} `json:"competition"`
	ID          int         `json:"id"`
	UtcDate     time.Time   `json:"utcDate"`
	Status      string      `json:"status"`
	Matchday    int         `json:"matchday"`
	Stage       string      `json:"stage"`
	Group       interface{} `json:"group"`
	LastUpdated time.Time   `json:"lastUpdated"`
	HomeTeam    struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		ShortName string `json:"shortName"`
		Tla       string `json:"tla"`
		Crest     string `json:"crest"`
	} `json:"homeTeam"`
	AwayTeam struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		ShortName string `json:"shortName"`
		Tla       string `json:"tla"`
		Crest     string `json:"crest"`
	} `json:"awayTeam"`
	Score struct {
		Winner   string `json:"winner"`
		Duration string `json:"duration"`
		FullTime struct {
			Home int `json:"home"`
			Away int `json:"away"`
		} `json:"fullTime"`
		HalfTime struct {
			Home int `json:"home"`
			Away int `json:"away"`
		} `json:"halfTime"`
	} `json:"score"`
}

type matchesCallResult struct {
	Matches []Match `json:"matches"`
}

func New(authKey string) *FootballApiClient {
	return &FootballApiClient{authKey: authKey, client: &http.Client{}}
}

func (c *FootballApiClient) GetAllTodaysMatches() ([]Match, error) {
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Auth-Token", c.authKey)
	return nil, nil
}
