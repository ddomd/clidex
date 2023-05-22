package pokeapi

import (
	"net/http"
	"time"
	"github.com/ddomd/clidex/internal/cache"
)

//Defines a custom client with a cache
type Client struct {
	client http.Client
	cache cache.Cache
}

//Generates a new client with a specified timeout and cache duration
func NewClient(timeout time.Duration, cacheDuration time.Duration) Client{
	return Client{
		client: http.Client{
			Timeout: timeout,
		},
		cache: cache.NewCache(cacheDuration),
	}
}