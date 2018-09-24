package ggriot

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/json-iterator/go"
	"github.com/soowan/ggriot/cache"
	"github.com/soowan/ggriot/models"
)

// GetMatch will return the data for the match ID requested.
func GetMatch(region string, matchID string) (md *models.MatchData, err error) {
	if cache.Enabled == true {
		ct := "league_match_by_id"
		var cc cache.Cached

		er := cache.CDB.Table(ct+"_"+region).Where("call_key = ?", matchID).First(&cc).Error
		switch er {
		case gorm.ErrRecordNotFound:
			if err = apiRequest("https://"+region+"."+Base+BaseMatch+"/matches/"+matchID, &md); err != nil {
				return md, err
			}

			if err = cache.StoreCall(ct, region, matchID, &md); err != nil {
				return md, err
			}

			return md, err
		case nil:
			jsoniter.UnmarshalFromString(cc.JSON, &md)

			return md, nil
		default:
			return md, errors.New("ggriot: unknown error, please open github issue: " + er.Error())
		}
	}

	if err = apiRequest("https://"+region+"."+Base+BaseMatch+"/matches/"+matchID, &md); err != nil {
		return md, err
	}

	return md, nil
}

// GetMatchHistory will return an array of matches based on the parameters given.
// This doesn't actually return wins or loses, just basic information about the game.
// In order to get stats you have to request every game separately.
// TODO: Add ability to fully use the options when doing a matches call.
// TODO: Figure out if/how this can/should be cached.
func GetMatchHistory(region string, accountID string) (ms *models.MatchHistory, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseMatch+"/matchlists/by-account/"+accountID, &ms)
	if err != nil {
		return nil, err
	}

	return ms, nil
}

// GetMatchTimeline will return the full timeline of the requested match ID.
// This is a pretty big struct that is returned so make sure you understand how to use this data first.
func GetMatchTimeline(region string, matchID string) (mt *models.MatchTimeline, err error) {
	if cache.Enabled == true {
		ct := "league_match_tl_by_id"
		var cc cache.Cached

		er := cache.CDB.Table(ct+"_"+region).Where("call_key = ?", matchID).First(&cc).Error
		switch er {
		case gorm.ErrRecordNotFound:
			if err = apiRequest("https://"+region+"."+Base+BaseMatch+"/timelines/by-match/"+matchID, &mt); err != nil {
				return mt, err
			}

			if err = cache.StoreCall(ct, region, matchID, &mt); err != nil {
				return mt, err
			}

			return mt, err
		case nil:
			jsoniter.UnmarshalFromString(cc.JSON, &mt)

			return mt, nil
		default:
			return mt, errors.New("ggriot: unknown error, please open github issue: " + er.Error())
		}
	}

	if err = apiRequest("https://"+region+"."+Base+BaseMatch+"/timelines/by-match/"+matchID, &mt); err != nil {
		return mt, err
	}

	return mt, nil
}
