package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type PokemonAbilitiesResponse struct {
	Count    int              `json:"count"`
	Next     string           `json:"next"`
	Previous *string          `json:"previous"`
	Results  []PokemonAbility `json:"results"`
}

type PokemonAbility struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonAbilityResponse struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Id             int            `json:"id"`
	Name           string         `json:"string"`
	BaseExperience int            `json:"base_experience"`
	Weight         int            `json:"weight"`
	Stats          []PokemonStats `json:"stats"`
}

type PokemonStats struct {
	BaseStat int         `json:"base_stat"`
	Effort   string      `json:"effort"`
	Stat     PokemonStat `json:"stat"`
}

type PokemonStat struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func GetPokemonAbilities() (PokemonAbilitiesResponse, error) {
	var response PokemonAbilitiesResponse

	url := "https://pokeapi.co/api/v2/ability/"

	req, reqErr := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Set("Content-Type", "application/json")

	if reqErr != nil {
		return response, reqErr
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	clientResponse, clientErr := client.Do(req)

	if clientErr != nil {
		return response, clientErr
	}

	if clientResponse.StatusCode != http.StatusOK {
		return response, errors.New("Request failed: " + clientResponse.Status)
	}

	defer clientResponse.Body.Close()

	responseBody, _ := io.ReadAll(clientResponse.Body)

	if marshErr := json.Unmarshal(responseBody, &response); marshErr != nil {
		return response, marshErr
	}

	return response, nil
}

func GetPokemonByAbility(abilityName string) (PokemonAbilityResponse, error) {
	var response PokemonAbilityResponse

	url := fmt.Sprintf("https://pokeapi.co/api/v2/ability/%s", abilityName)

	req, reqErr := http.NewRequest(http.MethodGet, url, nil)

	if reqErr != nil {
		return response, reqErr
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	clientResponse, clientErr := client.Do(req)

	if clientErr != nil {
		return response, clientErr
	}

	if clientResponse.StatusCode != http.StatusOK {
		return response, errors.New("Request failed - " + clientResponse.Status)
	}

	defer clientResponse.Body.Close()

	responseBody, _ := io.ReadAll(clientResponse.Body)

	if marshErr := json.Unmarshal(responseBody, &response); marshErr != nil {
		return response, marshErr
	}

	return response, nil
}
