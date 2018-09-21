package ggriot

import (
	"log"
	"testing"

	"github.com/soowan/ggriot/cache"
)

func TestActiveGame(t *testing.T) {
	SetAPIKey("")
	cache.UseCache()
	useCache = true
}

func TestGetActiveGame(t *testing.T) {
	e, err := GetMasteryList(NA, "21669611")
	if err != nil {
		t.Error(err)
	}

	log.Println((*e)[0].PlayerID)
}
