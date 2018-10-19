package ggriot

import (
	"errors"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/json-iterator/go"
	"github.com/soowan/ggriot/cache"
	"github.com/soowan/ggriot/models"
)

// Match will return the data for the match ID requested.
func Match(region string, matchID int64) (md *models.MatchData, err error) {
	mtID := strconv.FormatInt(matchID, 10)
	if cache.Enabled == true {
		ct := "league_match_by_id"
		var cc cache.Cached

		er := cache.CDB.Table(ct+"_"+region).Where("call_key = ?", mtID).First(&cc).Error
		switch er {
		case gorm.ErrRecordNotFound:
			if err = apiRequest("https://"+region+"."+Base+BaseMatch+"/matches/"+mtID, &md); err != nil {
				return md, err
			}

			if err = cache.StoreCall(ct, region, mtID, &md); err != nil {
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

	if err = apiRequest("https://"+region+"."+Base+BaseMatch+"/matches/"+mtID, &md); err != nil {
		return md, err
	}

	return md, nil
}

// MatchHistory will return an array of matches based on the parameters given.
// This doesn't actually return wins or loses, just basic information about the game.
// In order to get stats you have to request every game separately.
// TODO: Add ability to fully use the options when doing a matches call.
// TODO: Figure out if/how this can/should be cached.
func MatchHistory(region string, accountID int64) (ms *models.MatchHistory, err error) {
	accID := strconv.FormatInt(accountID, 10)
	err = apiRequest("https://"+region+"."+Base+BaseMatch+"/matchlists/by-account/"+accID, &ms)
	if err != nil {
		return nil, err
	}

	return ms, nil
}

// MatchTimeline will return the full timeline of the requested match ID.
// This is a pretty big struct that is returned so make sure you understand how to use this data first.
func MatchTimeline(region string, matchID int64) (mt *models.MatchTimeline, err error) {
	mtID := strconv.FormatInt(matchID, 10)
	if cache.Enabled == true {
		ct := "league_match_tl_by_id"
		var cc cache.Cached

		er := cache.CDB.Table(ct+"_"+region).Where("call_key = ?", mtID).First(&cc).Error
		switch er {
		case gorm.ErrRecordNotFound:
			if err = apiRequest("https://"+region+"."+Base+BaseMatch+"/timelines/by-match/"+mtID, &mt); err != nil {
				return mt, err
			}

			if err = cache.StoreCall(ct, region, mtID, &mt); err != nil {
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

	if err = apiRequest("https://"+region+"."+Base+BaseMatch+"/timelines/by-match/"+mtID, &mt); err != nil {
		return mt, err
	}

	return mt, nil
}
