package ggriot

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/json-iterator/go"
	"github.com/soowan/ggriot/cache"
	"github.com/soowan/ggriot/models"
)

// ExpireMMR is how long cache should be kept
var ExpireMMR = time.Duration(1440 * time.Minute)

// MMR will return the data for the MMR ID requested.
func MMR(region string, summoner string) (mmr *models.MMR, err error) {
	s, err := SummonerByIGN(region, summoner)
	if err != nil {
		return mmr, err
	}
	rrr := region
	switch region {
	case NA:
		region = "na"
	case EUN:
		region = "eun"
	case EUW:
		region = "euw"
	default:
		return mmr, errors.New("ggriot: mmr doesnt support " + rrr)
	}

	if cache.Enabled == true {
		ct := "mmr"
		var cc cache.Cached

		er := cache.CDB.Table(ct+"_"+rrr).Where("call_key = ?", summoner).First(&cc).Error
		switch er {
		case gorm.ErrRecordNotFound:
			if err = apiRequest("https://"+region+".whatismymmr.com/api/v1/summoner?name="+s.Name, &mmr); err != nil {
				return mmr, err
			}

			if err = cache.StoreCall(ct, rrr, summoner, &mmr); err != nil {
				return mmr, err
			}

			return mmr, err
		case nil:
			if time.Since(cc.UpdatedAt) > ExpireMMR {

				if err = apiRequest("https://"+region+".whatismymmr.com/api/v1/summoner?name="+s.Name, &mmr); err != nil {
					return mmr, err
				}

				cache.UpdateCall(ct, rrr, &cc, &mmr)

				return mmr, nil
			}

			jsoniter.UnmarshalFromString(cc.JSON, &mmr)
			return mmr, nil
		default:
			return mmr, errors.New("ggriot: unknown error, please open github issue: " + er.Error())
		}
	}

	if err = apiRequest("https://"+region+".whatismymmr.com/api/v1/summoner?name="+s.Name, &mmr); err != nil {
		return mmr, err
	}

	return mmr, nil
}
