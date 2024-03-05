package footballApiClient

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
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

func (c *FootballApiClient) GetAllTodaysMatches(dateFrom, dateTo string) ([]Match, error) {
	req, err := http.NewRequest("GET", baseUrl+"/matches", nil)
	if err != nil {
		log.Println("Unable to create request")
		return nil, err
	}
	req.Header.Add("X-Auth-Token", c.authKey)

	q := &url.Values{
		"dateFrom": []string{dateFrom},
		"dateTo":   []string{dateTo},
	}
	req.URL.RawQuery = q.Encode()
	resp, err := c.client.Do(req)
	if err != nil {
		log.Println("Failed to send request to server")
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read server response")
		return nil, err
	}

	res := &matchesCallResult{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Println("Failed to unmarshal server response")
		return nil, err
	}
	return res.Matches, nil
}

func (c *FootballApiClient) GetTodayMatchesForLeagues(leagues []string) ([]Match, error) {
	req, err := http.NewRequest("GET", baseUrl+"/matches", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Auth-Token", c.authKey)

	q := &url.Values{
		"competitions": []string{strings.Join(leagues, ",")},
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
