package models

// Summoner is a struct used when returning information about the summoner.
type Summoner struct {
	ProfileIconID int    `json:"profileIconId"`
	Name          string `json:"name"`
	SummonerLevel int    `json:"summonerLevel"`
	AccountID     int64  `json:"accountId"`
	ID            int64  `json:"id"`
	RevisionDate  int64  `json:"revisionDate"`
}
