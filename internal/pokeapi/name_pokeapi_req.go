package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/adominguez793/pokedexcli/internal/pokecache"
)

func (c *Client) NamePokeapiReq(name string, cache *pokecache.Cache) (NameLocationArea, error) {
	endpointURL := "/location-area/"
	fullURL := baseURL + endpointURL + name
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return NameLocationArea{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return NameLocationArea{}, err
	}
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return NameLocationArea{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		return NameLocationArea{}, err
	}

	var NameLocation NameLocationArea
	err = json.Unmarshal(dat, &NameLocation)
	if err != nil {
		return NameLocationArea{}, err
	}

	cache.Add(name, dat)
	return NameLocation, nil
}
