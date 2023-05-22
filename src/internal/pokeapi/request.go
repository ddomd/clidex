package pokeapi

import (
	"errors"
	"net/http"
	"io/ioutil"
)

func (c* Client) request(endpoint string) ([]byte, error){
	data, ok := c.cache.Get(endpoint); if ok {
		return data, nil
	}

	request, err := http.NewRequest("GET", endpoint, nil)

	if err != nil {
		return []byte{}, err
	}

	response, err := c.client.Do(request)

	if err != nil {
		return []byte{}, err
	}

	if response.StatusCode > 399 {
		return []byte{}, errors.New("\nCouldn't find resource")
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return []byte{}, err
	}

	c.cache.Add(endpoint, responseData)

	return responseData, nil
}