package ggriot

import "github.com/soowan/ggriot/models"

// TODO: Remove this. (These api calls will be deprecated in october.)

// GetAllChampions will return all the current champions in the game.
// This isn't affected by API limit rates, however it's good practice to still obey them.
// However this data isn't exactly all that useful as it still only returns champion IDs.
// You are better off creating a cache for champions yourself, and use that as needed.
func GetAllChampions(region string) (cl *models.ChampionsList, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseChampion, &cl)

	if err != nil {
		return cl, err
	}

	return cl, nil
}

// GetChampion returns the information on one champion.
// Like GetAllChampions this doesn't count towards rate limiting.
func GetChampion(region string, championID string) (c *models.Champion, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseChampion+"/"+championID, &c)

	if err != nil {
		return c, err
	}

	return c, nil
}

// GetChampionRotation is for getting the champion rotation for the week.
func GetChampionRotation(region string) (cr *models.ChampionRotation, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseChampionR, &cr)

	if err != nil {
		return cr, err
	}
	return cr, nil
}
