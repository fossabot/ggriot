package ggriot

import (
	"errors"
	"time"

	"github.com/soowan/ggriot/models"

	"github.com/jinzhu/gorm"

	"github.com/json-iterator/go"
	"github.com/soowan/ggriot/cache"
)

var (
	// LeagueByIDExpire is how long this call should be considered "fresh".
	LeagueByIDExpire = time.Duration(15 * time.Minute)

	// ExpireGetChallengers is how long this call should be considered "fresh"
	ExpireGetChallengers = time.Duration(15 * time.Minute)

	// ExpireGetMasters is how long this call should be considered "fresh"
	ExpireGetMasters = time.Duration(15 * time.Minute)

	// ExpireGetPlayerPosition is how long this call should be considered "fresh"
	ExpireGetPlayerPosition = time.Duration(15 * time.Minute)
)

// GetChallengers will return all the challengers in the requested queue.
// Valid queues are, Ranked5s(RANKED_SOLO_5x5), Flex3s(RANKED_FLEX_TT), and Flex5s(RANKED_FLEX_SR)
func GetChallengers(region string, mode string) (lr *models.LeagueRoster, err error) {
	if cache.Enabled == true {
		ct := "league_challenger_by_queue"
		var cc cache.Cached

		switch cache.CDB.Table(ct+"_"+region).Where("call_key = ?", mode).First(&cc).Error {
		case gorm.ErrRecordNotFound:
			if err = apiRequest("https://"+region+"."+Base+BaseLeague+"/challengerleagues/by-queue/"+mode, &lr); err != nil {
				return lr, err
			}

			if err = cache.StoreCall(ct, region, mode, &lr); err != nil {
				return lr, err
			}

			return lr, err
		case nil:
			if time.Since(cc.UpdatedAt) > ExpireGetChallengers {

				if err = apiRequest("https://"+region+"."+Base+BaseLeague+"/challengerleagues/by-queue/"+mode, &lr); err != nil {
					return lr, err
				}

				cache.UpdateCall(ct, region, &cc, &lr)

				return lr, nil

			}

			jsoniter.UnmarshalFromString(cc.JSON, &lr)

			return lr, nil
		default:
			return lr, errors.New("uh unkown error with case i suppose")
		}
	}

	if err = apiRequest("https://"+region+"."+Base+BaseLeague+"/challengerleagues/by-queue/"+mode, &lr); err != nil {
		return lr, err
	}

	return lr, nil
}

// GetMasters returns the roster of all the players in the Masters tier for requested queue.
func GetMasters(region string, mode string) (lr *models.LeagueRoster, err error) {
	if cache.Enabled == true {
		ct := "league_master_by_queue"
		var cc cache.Cached

		er := cache.CDB.Table(ct+"_"+region).Where("call_key = ?", mode).First(&cc).Error
		switch er {
		case gorm.ErrRecordNotFound:
			if err = apiRequest("https://"+region+"."+Base+BaseLeague+"/masterleagues/by-queue/"+mode, &lr); err != nil {
				return lr, err
			}

			if err = cache.StoreCall(ct, region, mode, &lr); err != nil {
				return lr, err
			}

			return lr, err
		case nil:
			if time.Since(cc.UpdatedAt) > ExpireGetChallengers {

				if err = apiRequest("https://"+region+"."+Base+BaseLeague+"/masterleagues/by-queue/"+mode, &lr); err != nil {
					return lr, err
				}

				cache.UpdateCall(ct, region, &cc, &lr)

				return lr, nil
			}

			jsoniter.UnmarshalFromString(cc.JSON, &lr)

			return lr, nil
		default:
			return lr, errors.New("ggriot: unknown error, please open github issue: " + er.Error())
		}
	}

	if err = apiRequest("https://"+region+"."+Base+BaseLeague+"/masterleagues/by-queue/"+mode, &lr); err != nil {
		return lr, err
	}

	return lr, nil
}

// GetLeague will return the roster of the league requested, via UUID.
// You can and will get blocked from this call if you provide invalid UUIDs
func GetLeague(region string, leagueUUID string) (lr *models.LeagueRoster, err error) {
	if cache.Enabled == true {
		ct := "league_by_id"
		var cc cache.Cached

		er := cache.CDB.Table(ct+"_"+region).Where("call_key = ?", leagueUUID).First(&cc).Error
		switch er {
		case gorm.ErrRecordNotFound:
			if err = apiRequest("https://"+region+"."+Base+BaseLeague+"/leagues/"+leagueUUID, &lr); err != nil {
				return lr, err
			}

			if err = cache.StoreCall(ct, region, leagueUUID, &lr); err != nil {
				return lr, err
			}

			return lr, err
		case nil:
			if time.Since(cc.UpdatedAt) > LeagueByIDExpire {

				if err = apiRequest("https://"+region+"."+Base+BaseLeague+"/leagues/"+leagueUUID, &lr); err != nil {
					return lr, err
				}

				cache.UpdateCall(ct, region, &cc, &lr)

				return lr, nil

			}

			jsoniter.UnmarshalFromString(cc.JSON, &lr)

			return lr, nil
		default:
			return lr, errors.New("ggriot: unknown error, please open github issue: " + er.Error())
		}
	}

	if err = apiRequest("https://"+region+"."+Base+BaseLeague+"/leagues/"+leagueUUID, &lr); err != nil {
		return lr, err
	}

	return lr, nil
}

// GetPlayerPosition will return the requested players league position in each of the three ranked queues.
func GetPlayerPosition(region string, summonerID string) (lp *models.LeaguePosition, err error) {
	if cache.Enabled == true {
		ct := "league_position_by_summoner"
		var cc cache.Cached

		er := cache.CDB.Table(ct+"_"+region).Where("call_key = ?", summonerID).First(&cc).Error
		switch er {
		case gorm.ErrRecordNotFound:
			if err = apiRequest("https://"+region+"."+Base+BaseLeague+"/positions/by-summoner/"+summonerID, &lp); err != nil {
				return lp, err
			}

			if err = cache.StoreCall(ct, region, summonerID, &lp); err != nil {
				return lp, err
			}

			return lp, err
		case nil:
			if time.Since(cc.UpdatedAt) > ExpireGetChallengers {

				if err = apiRequest("https://"+region+"."+Base+BaseLeague+"/positions/by-summoner/"+summonerID, &lp); err != nil {
					return lp, err
				}

				cache.UpdateCall(ct, region, &cc, &lp)

				return lp, nil
			}

			jsoniter.UnmarshalFromString(cc.JSON, &lp)

			return lp, nil
		default:
			return lp, errors.New("ggriot: unknown error, please open github issue: " + er.Error())
		}
	}

	if err = apiRequest("https://"+region+"."+Base+BaseLeague+"/positions/by-summoner/"+summonerID, &lp); err != nil {
		return lp, err
	}

	return lp, nil
}
