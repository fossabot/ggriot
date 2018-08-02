package ggriot

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

// GetAllChampions will return all the current champions in the game.
// This isn't affected by API limit rates, however it's good practice to still obey them.
// However this data isn't exactly all that useful as it still only returns champion IDs.
// You are better off creating a cache for champions yourself, and use that as needed.
func GetAllChampions(region string) (cl *ChampionsList, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseChampion, &cl)

	if err != nil {
		return nil, err
	}

	return cl, nil
}

// GetChampion returns the information on one champion.
// Like GetAllChampions this doesn't count towards rate limiting.
func GetChampion(region string, championID string) (c *Champion, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseChampion+"/"+championID, &c)

	if err != nil {
		return nil, err
	}

	return c, nil
}
