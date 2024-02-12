package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

const baseURL = "https://pokeapi.co/api/v2"

func (c *Client) PokeapiReq() (LocationArea, error) {
	endpointURL := "/location-area"
	fullURL := baseURL + endpointURL
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	if resp.StatusCode > 299 {
		return LocationArea{}, errors.New("Bad status code...")
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}
	var Location LocationArea
	err = json.Unmarshal(dat, &Location)
	if err != nil {
		return LocationArea{}, err
	}

	return Location, nil
}
