package pokeapi

import (
	"encoding/json"
	"strings"
)

func normalizeFlavorText(s string) string {
	removeEscapes := strings.NewReplacer(
		"\f", " ",
		"\n", "\n\t| ")
	format := strings.NewReplacer(". ", "\n\t| ")

	s = removeEscapes.Replace(s)

	return format.Replace(s)
}

func (c *Client) GetDex(query string) (Pokedex, error) {
	statsEndpoint := apiUrl + "pokemon/" + query
	flavorEndpoint := apiUrl + "pokemon-species/" + query

	statsData, err := c.request(statsEndpoint); if err != nil {
		return Pokedex{}, err
	}

	flavorData, err := c.request(flavorEndpoint); if err != nil {
		return Pokedex{}, err
	}

	var statsObj Pokemon
	var flavorObj PokemonSpecies

	json.Unmarshal(statsData, &statsObj)
	json.Unmarshal(flavorData, &flavorObj)
	
	return Pokedex{
		Name: statsObj.Name,
		Number: statsObj.ID,
		Height: statsObj.Height,
		Weight: statsObj.Weight,
		Flavor: normalizeFlavorText(flavorObj.FlavorTextEntries[0].FlavorText),
	}, nil

}