package models

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
