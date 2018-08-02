package ggriot

// Summoner is a struct used when returning information about the summoner.
type Summoner struct {
	ProfileIconID int    `json:"profileIconId"`
	Name          string `json:"name"`
	SummonerLevel int    `json:"summonerLevel"`
	AccountID     int    `json:"accountId"`
	ID            int    `json:"id"`
	RevisionDate  int64  `json:"revisionDate"`
}

// GetSummonerByAccID will get summoner information using Account ID
func GetSummonerByAccID(region string, accountID string) (s *Summoner, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseSummoner+"/by-account/"+accountID, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetSummonerBySumID will get summoner information using Summoner ID
func GetSummonerBySumID(region string, summonerID string) (s *Summoner, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseSummoner+"/"+summonerID, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetSummonerByIGN will get summoner information using IGN
func GetSummonerByIGN(region string, summonerIGN string) (s *Summoner, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseSummoner+"/by-name/"+summonerIGN, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
