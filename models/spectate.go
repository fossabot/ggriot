package models

// ActiveGame is the struct that holds information about the current game the summoner is in.
type ActiveGame struct {
	GameID            int64  `json:"gameId"`
	GameStartTime     int    `json:"gameStartTime"`
	PlatformID        string `json:"platformId"`
	GameMode          string `json:"gameMode"`
	MapID             int    `json:"mapId"`
	GameType          string `json:"gameType"`
	GameQueueConfigID int    `json:"gameQueueConfigId"`
	Observers         struct {
		EncryptionKey string `json:"encryptionKey"`
	} `json:"observers"`
	Participants []struct {
		ProfileIconID            int           `json:"profileIconId"`
		ChampionID               int           `json:"championId"`
		SummonerName             string        `json:"summonerName"`
		GameCustomizationObjects []interface{} `json:"gameCustomizationObjects"`
		Bot                      bool          `json:"bot"`
		Perks                    struct {
			PerkStyle    int   `json:"perkStyle"`
			PerkIds      []int `json:"perkIds"`
			PerkSubStyle int   `json:"perkSubStyle"`
		} `json:"perks"`
		Spell2ID   int `json:"spell2Id"`
		TeamID     int `json:"teamId"`
		Spell1ID   int `json:"spell1Id"`
		SummonerID int `json:"summonerId"`
	} `json:"participants"`
	GameLength      int `json:"gameLength"`
	BannedChampions []struct {
		TeamID     int `json:"teamId"`
		ChampionID int `json:"championId"`
		PickTurn   int `json:"pickTurn"`
	} `json:"bannedChampions"`
}

// FeaturedGames returns Riot's featured games.
type FeaturedGames struct {
	ClientRefreshInterval int `json:"clientRefreshInterval"`
	GameList              []struct {
		GameID            int64  `json:"gameId"`
		GameStartTime     int64  `json:"gameStartTime"`
		PlatformID        string `json:"platformId"`
		GameMode          string `json:"gameMode"`
		MapID             int    `json:"mapId"`
		GameType          string `json:"gameType"`
		GameQueueConfigID int    `json:"gameQueueConfigId"`
		Observers         struct {
			EncryptionKey string `json:"encryptionKey"`
		} `json:"observers"`
		Participants []struct {
			ProfileIconID int    `json:"profileIconId"`
			ChampionID    int    `json:"championId"`
			SummonerName  string `json:"summonerName"`
			Bot           bool   `json:"bot"`
			Spell2ID      int    `json:"spell2Id"`
			TeamID        int    `json:"teamId"`
			Spell1ID      int    `json:"spell1Id"`
		} `json:"participants"`
		GameLength      int           `json:"gameLength"`
		BannedChampions []interface{} `json:"bannedChampions"`
	} `json:"gameList"`
}
