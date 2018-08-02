package ggriot

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
type MasteryLevel struct {
	TotalMastery int
}

// GetMasteryList will return a struct with all the summoners champions and mastery exp/level.
func GetMasteryList(region string, summoner string) (ml *MasteryList, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseMastery+"/champion-masteries/by-summoner/"+summoner, &ml)
	if err != nil {
		return ml, err
	}

	return ml, nil
}

// GetChampionMastery will return a single champion mastery struct
func GetChampionMastery(region string, summoner string, championID string) (ml *MasteryList, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseMastery+"/champion-masteries/by-summoner/"+summoner+"/by-champion"+championID+apikey, &ml)
	if err != nil {
		return nil, err
	}
	return ml, nil
}

// GetTotalMasteryLevel gets the total mastery level.
func GetTotalMasteryLevel(region string, summoner string) (ml *MasteryLevel, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseMastery+"/scores/by-summoner/"+summoner, &ml)
	if err != nil {
		return nil, err
	}
	return ml, nil
}
