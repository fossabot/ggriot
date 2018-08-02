package ggriot

// LeagueRoster returns everyone in the  tier for the mode requested.
type LeagueRoster struct {
	Tier     string `json:"tier"`
	Queue    string `json:"queue"`
	LeagueID string `json:"leagueId"`
	Name     string `json:"name"`
	Entries  []struct {
		HotStreak        bool   `json:"hotStreak"`
		Wins             int    `json:"wins"`
		Veteran          bool   `json:"veteran"`
		Losses           int    `json:"losses"`
		Rank             string `json:"rank"`
		PlayerOrTeamName string `json:"playerOrTeamName"`
		Inactive         bool   `json:"inactive"`
		PlayerOrTeamID   string `json:"playerOrTeamId"`
		FreshBlood       bool   `json:"freshBlood"`
		LeaguePoints     int    `json:"leaguePoints"`
	} `json:"entries"`
}

// LeaguePosition is what's returned when requesting a players league stats.
type LeaguePosition []struct {
	QueueType        string `json:"queueType"`
	HotStreak        bool   `json:"hotStreak"`
	Wins             int    `json:"wins"`
	Veteran          bool   `json:"veteran"`
	Losses           int    `json:"losses"`
	PlayerOrTeamID   string `json:"playerOrTeamId"`
	LeagueName       string `json:"leagueName"`
	PlayerOrTeamName string `json:"playerOrTeamName"`
	Inactive         bool   `json:"inactive"`
	Rank             string `json:"rank"`
	FreshBlood       bool   `json:"freshBlood"`
	LeagueID         string `json:"leagueId"`
	Tier             string `json:"tier"`
	LeaguePoints     int    `json:"leaguePoints"`
}

// GetChallengers will return all the challengers in the requested queue.
// Valid queues are, Ranked5s(RANKED_SOLO_5x5), Flex3s(RANKED_FLEX_TT), and Flex5s(RANKED_FLEX_SR)
func GetChallengers(region string, mode string) (lr *LeagueRoster, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseLeague+"/challengerleagues/by-queue/"+mode, &lr)
	if err != nil {
		return nil, err
	}

	return lr, nil
}

// GetMasters returns the roster of all the players in the Masters tier for requested queue.
func GetMasters(region string, mode string) (lr *LeagueRoster, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseLeague+"/masterleagues/by-queue/"+mode, &lr)
	if err != nil {
		return nil, err
	}

	return lr, nil
}

// GetLeague will return the roster of the league requested, via UUID.
// You can and will get blocked from this call if you provide invalid UUIDs
func GetLeague(region string, leagueUUID string) (lr *LeagueRoster, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseLeague+"/leagues/"+leagueUUID, &lr)
	if err != nil {
		return nil, err
	}

	return lr, nil
}

// GetPlayerPosition will return the requested players league position in each of the three ranked queues.
func GetPlayerPosition(region string, summonerID string) (lp *LeaguePosition, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseLeague+"/positions/by-summoner/"+summonerID, &lp)
	if err != nil {
		return nil, err
	}

	return lp, nil
}
