package models

// MasteryList is a slice with every champion and its information in regards to mastery.
type MasteryList []struct {
	ChampionLevel                int   `json:"championLevel"`
	ChestGranted                 bool  `json:"chestGranted"`
	ChampionPoints               int   `json:"championPoints"`
	ChampionPointsSinceLastLevel int   `json:"championPointsSinceLastLevel"`
	PlayerID                     int   `json:"playerId"`
	ChampionPointsUntilNextLevel int   `json:"championPointsUntilNextLevel"`
	TokensEarned                 int   `json:"tokensEarned"`
	ChampionID                   int   `json:"championId"`
	LastPlayTime                 int64 `json:"lastPlayTime"`
}

// MasteryLevel is the total champion mastery.
type MasteryLevel int

// ChampionsList is used when returning all the current champions.
type ChampionsList struct {
	Champions []struct {
		RankedPlayEnabled bool `json:"rankedPlayEnabled"`
		BotEnabled        bool `json:"botEnabled"`
		BotMmEnabled      bool `json:"botMmEnabled"`
		Active            bool `json:"active"`
		FreeToPlay        bool `json:"freeToPlay"`
		ID                int  `json:"id"`
	} `json:"champions"`
}

// Champion is used when returning a single champion.
type Champion struct {
	RankedPlayEnabled bool `json:"rankedPlayEnabled"`
	BotEnabled        bool `json:"botEnabled"`
	BotMmEnabled      bool `json:"botMmEnabled"`
	Active            bool `json:"active"`
	FreeToPlay        bool `json:"freeToPlay"`
	ID                int  `json:"id"`
}

// ChampionRotation is used when returning the champion rotation of the week.
type ChampionRotation struct {
	FreeChampionIds              []int `json:"freeChampionIds"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
	FreeChampionIdsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
}
