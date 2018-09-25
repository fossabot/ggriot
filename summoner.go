package ggriot

import "github.com/soowan/ggriot/models"

// GetSummonerByAccID will get summoner information using Account ID
func GetSummonerByAccID(region string, accountID string) (s *models.Summoner, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseSummoner+"/by-account/"+accountID, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetSummonerBySumID will get summoner information using Summoner ID
func GetSummonerBySumID(region string, summonerID string) (s *models.Summoner, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseSummoner+"/"+summonerID, &s)
	if err != nil {
		return s, err
	}
	return s, nil
}

// GetSummonerByIGN will get summoner information using IGN
func GetSummonerByIGN(region string, summonerIGN string) (s *models.Summoner, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseSummoner+"/by-name/"+summonerIGN, &s)
	if err != nil {
		return s, err
	}
	return s, nil
}
