package ggriot

import (
	"errors"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/json-iterator/go"
	"github.com/soowan/ggriot/cache"

	"github.com/soowan/ggriot/models"
)

var (
	// ExpireSummoner is how long the summoner by ign cache is saved.
	ExpireSummoner = time.Duration(60 * time.Minute)
)

// SummonerByAccID will get summoner information using Account ID
func SummonerByAccID(region string, accountID int64) (s *models.Summoner, err error) {
	accID := strconv.FormatInt(accountID, 10)
	err = apiRequest("https://"+region+"."+Base+BaseSummoner+"/by-account/"+accID, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// SummonerBySumID will get summoner information using Summoner ID
func SummonerBySumID(region string, summonerID int64) (s *models.Summoner, err error) {
	sumID := strconv.FormatInt(summonerID, 10)
	err = apiRequest("https://"+region+"."+Base+BaseSummoner+"/"+sumID, &s)
	if err != nil {
		return s, err
	}
	return s, nil
}

// SummonerByIGN will get summoner information using IGN
func SummonerByIGN(region string, summonerIGN string) (s *models.Summoner, err error) {
	if cache.Enabled == true {
		ct := "summoner_by_ign"
		var cc cache.Cached

		er := cache.CDB.Table(ct+"_"+region).Where("call_key = ?", summonerIGN).First(&cc).Error
		switch er {
		case gorm.ErrRecordNotFound:
			if err = apiRequest("https://"+region+"."+Base+BaseSummoner+"/by-name/"+summonerIGN, &s); err != nil {
				return s, err
			}

			if err = cache.StoreCall(ct, region, summonerIGN, &s); err != nil {
				return s, err
			}

			return s, err
		case nil:
			if time.Since(cc.UpdatedAt) > ExpireGetPlayerPosition {

				if err = apiRequest("https://"+region+"."+Base+BaseSummoner+"/by-name/"+summonerIGN, &s); err != nil {
					return s, err
				}

				cache.UpdateCall(ct, region, &cc, &s)

				return s, nil
			}

			jsoniter.UnmarshalFromString(cc.JSON, &s)

			return s, nil
		default:
			return s, errors.New("ggriot: unknown error, please open github issue: " + er.Error())
		}
	}

	if err = apiRequest("https://"+region+"."+Base+BaseSummoner+"/by-name/"+summonerIGN, &s); err != nil {
		return s, err
	}

	return s, nil
}
