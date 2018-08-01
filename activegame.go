package ggriot

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

// GetActiveGame will get the active game from the supplied id.
func GetActiveGame(region string, summoner string) (ag *ActiveGame, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseSpectator+"/active-games/by-summoner/"+summoner+apikey, &ag)
	if err != nil {
		return nil, err
	}
	return ag, err
}
