package ggriot

import "github.com/soowan/ggriot/models"

// GetMatch will return the data for the match ID requested.
func GetMatch(region string, matchID string) (md *models.MatchData, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseMatch+"/matches/"+matchID+apikey, &md)
	if err != nil {
		return nil, err
	}

	return md, nil
}

// GetMatchHistory will return an array of matches based on the parameters given.
// This doesn't actually return wins or loses, just basic information about the game.
// In order to get stats you have to request every game separately.
// TODO: Add ability to fully use the options when doing a matches call.
func GetMatchHistory(region string, accountID string) (ms *models.MatchHistory, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseMatch+"/matchlists/by-account/"+accountID+apikey, &ms)
	if err != nil {
		return nil, err
	}

	return ms, nil
}

// GetMatchTimeline will return the full timeline of the requested match ID.
// This is a pretty big struct that is returned so make sure you understand how to use this data first.
func GetMatchTimeline(region string, matchID string) (mt *models.MatchTimeline, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseMatch+"/timelines/by-match/"+matchID, &mt)
	if err != nil {
		return nil, err
	}

	return mt, nil
}
