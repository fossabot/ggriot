package ggriot

import "github.com/soowan/ggriot/models"

// ChampionRotation is for getting the champion rotation for the week.
// This is the new api endpoint added. Subject to change.
func ChampionRotation(region string) (cr *models.ChampionRotation, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseChampionR, &cr)

	if err != nil {
		return cr, err
	}
	return cr, nil
}
