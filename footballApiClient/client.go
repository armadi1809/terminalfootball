package footballApiClient

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

const baseUrl string = "https://api.football-data.org/v4"

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
	req, err := http.NewRequest("GET", baseUrl+"/matches", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Auth-Token", c.authKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	res := &matchesCallResult{}
	json.Unmarshal(body, &res)
	return res.Matches, nil
}

func (c *FootballApiClient) GetTodayMatchesForLeagues(leagues []string) ([]Match, error) {
	req, err := http.NewRequest("GET", baseUrl+"/matches", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Auth-Token", c.authKey)

	q := &url.Values{
		"competitions": leagues,
	}

	req.URL.RawQuery = q.Encode()
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	res := &matchesCallResult{}
	json.Unmarshal(body, &res)
	return res.Matches, nil
}

func constructLeagueQueries(leagues []string) string {
	query := ""

	for _, league := range leagues {
		if query == "" {
			query += league
			continue
		}
		query += "," + league
	}
	return query
}
