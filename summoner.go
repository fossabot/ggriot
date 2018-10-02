package ggriot

import "github.com/soowan/ggriot/models"

// SummonerByAccID will get summoner information using Account ID
func SummonerByAccID(region string, accountID string) (s *models.Summoner, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseSummoner+"/by-account/"+accountID, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// SummonerBySumID will get summoner information using Summoner ID
func SummonerBySumID(region string, summonerID string) (s *models.Summoner, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseSummoner+"/"+summonerID, &s)
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
