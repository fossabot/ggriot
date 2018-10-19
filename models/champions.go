package models

// MasteryList is a slice with every champion and its information in regards to mastery.
type MasteryList []struct {
	ChampionLevel                int   `json:"championLevel"`
	ChestGranted                 bool  `json:"chestGranted"`
	ChampionPoints               int   `json:"championPoints"`
	ChampionPointsSinceLastLevel int   `json:"championPointsSinceLastLevel"`
	PlayerID                     int64 `json:"playerId"`
	ChampionPointsUntilNextLevel int   `json:"championPointsUntilNextLevel"`
	TokensEarned                 int   `json:"tokensEarned"`
	ChampionID                   int   `json:"championId"`
	LastPlayTime                 int64 `json:"lastPlayTime"`
}

// MasteryLevel is the total champion mastery.
type MasteryLevel int

// ChampionRotation is used when returning the champion rotation of the week.
type ChampionRotation struct {
	FreeChampionIds              []int `json:"freeChampionIds"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
	FreeChampionIdsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
}
