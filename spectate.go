package ggriot

import (
	"strconv"

	"github.com/soowan/ggriot/models"
)

// ActiveGame will get the active game from the supplied id.
func ActiveGame(region string, summonerID int64) (ag *models.ActiveGame, err error) {
	summID := strconv.FormatInt(summonerID, 10)
	err = apiRequest("https://"+region+"."+Base+BaseSpectator+"/active-games/by-summoner/"+summID, &ag)

	if err != nil {
		return ag, err
	}

	return ag, nil
}

// FeaturedGames will get Riot's featured games.
func FeaturedGames(region string) (fg *models.FeaturedGames, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseSpectator+"/featured-games", &fg)

	if err != nil {
		return fg, err
	}

	return fg, nil
}
