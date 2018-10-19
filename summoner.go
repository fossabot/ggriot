package ggriot

import (
	"strconv"

	"github.com/soowan/ggriot/models"
)

// SummonerByAccID will get summoner information using Account ID
func SummonerByAccID(region string, accountID int64) (s *models.Summoner, err error) {
	accID := strconv.FormatInt(accountID, 10)
	err = apiRequest("https://"+region+"."+Base+BaseSummoner+"/by-account/"+accID, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// SummonerBySumID will get summoner information using Summoner ID
func SummonerBySumID(region string, summonerID int64) (s *models.Summoner, err error) {
	sumID := strconv.FormatInt(summonerID, 10)
	err = apiRequest("https://"+region+"."+Base+BaseSummoner+"/"+sumID, &s)
	if err != nil {
		return s, err
	}
	return s, nil
}

// SummonerByIGN will get summoner information using IGN
func SummonerByIGN(region string, summonerIGN string) (s *models.Summoner, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseSummoner+"/by-name/"+summonerIGN, &s)
	if err != nil {
		return s, err
	}
	return s, nil
}
