package ggriot

import "github.com/soowan/ggriot/models"

// GetActiveGame will get the active game from the supplied id.
func GetActiveGame(region string, summoner string) (ag *models.ActiveGame, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseSpectator+"/active-games/by-summoner/"+summoner, &ag)

	if err != nil {
		return ag, err
	}

	return ag, nil
}

// GetFeaturedGames will get Riot's featured games.
func GetFeaturedGames(region string) (fg *models.FeaturedGames, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseSpectator+"/featured-games", &fg)

	if err != nil {
		return fg, err
	}

	return fg, nil
}
