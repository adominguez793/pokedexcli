package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CatchPokeapiReq(pokemonName string) (PokeInfo, error) {
	fullURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemonName)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		fmt.Printf("%s\n", err)
		return PokeInfo{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Printf("%s\n", err)
		return PokeInfo{}, err
	}
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s\n", err)
		return PokeInfo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		fmt.Printf("Status Code: %d\n", resp.StatusCode)
		return PokeInfo{}, errors.New("Status code error")
	}

	var pokemonInfo PokeInfo
	err = json.Unmarshal(dat, &pokemonInfo)
	if err != nil {
		fmt.Printf("%s\n", err)
		return PokeInfo{}, err
	}
	return pokemonInfo, nil
}
