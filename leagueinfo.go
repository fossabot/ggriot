package ggriot

// SeasonStatus is a struct that has the current information about a summoners league information.
type SeasonStatus []struct {
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