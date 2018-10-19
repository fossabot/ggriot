package ggriot

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/soowan/ggriot/cache"
	"github.com/soowan/ggriot/models"
)

var (
	// GetMasteryListExpire sets the time it takes for this cached call to be considered "expired"
	GetMasteryListExpire = time.Duration(240 * time.Minute)

	// GetTotalMasteryLevelExpire sets the time it takes for this cached call to be considered "expired"
	GetTotalMasteryLevelExpire = time.Duration(240 * time.Minute)
)

// MasteryList will return a struct with all the summoners champions and mastery exp/level.
func MasteryList(region string, summonerID int64) (ml *models.MasteryList, err error) {
	summID := strconv.FormatInt(summonerID, 10)
	if cache.Enabled == true {
		ct := "mastery_by_summoner"
		var cc cache.Cached

		er := cache.CDB.Table(ct+"_"+region).Where("call_key = ?", summID).First(&cc).Error
		switch er {
		case gorm.ErrRecordNotFound:
			if err = apiRequest("https://"+region+"."+Base+BaseMastery+"/champion-masteries/by-summoner/"+summID, &ml); err != nil {
				return ml, err
			}

			if err = cache.StoreCall(ct, region, summID, &ml); err != nil {
				return ml, err
			}

			return ml, err
		case nil:
			if time.Since(cc.UpdatedAt) > GetMasteryListExpire {

				if err = apiRequest("https://"+region+"."+Base+BaseMastery+"/champion-masteries/by-summoner/"+summID, &ml); err != nil {
					return ml, err
				}

				cache.UpdateCall(ct, region, &cc, &ml)

				return ml, nil

			}

			jsoniter.UnmarshalFromString(cc.JSON, &ml)

			return ml, nil
		default:
			return ml, errors.New("ggriot: unknown error, please open github issue: " + er.Error())
		}
	}

	if err = apiRequest("https://"+region+"."+Base+BaseMastery+"/champion-masteries/by-summoner/"+summID, &ml); err != nil {
		return ml, err
	}

	return ml, nil
}

// ChampionMastery will return a single champion mastery struct
// TODO: Add special case for this, as it uses two inputs.
func ChampionMastery(region string, summonerID int64, championID int) (ml *models.MasteryList, err error) {
	summID := strconv.FormatInt(summonerID, 10)
	chamID := strconv.Itoa(championID)
	err = apiRequest("https://"+region+"."+Base+BaseMastery+"/champion-masteries/by-summoner/"+summID+"/by-champion"+chamID+apikey, &ml)
	if err != nil {
		return ml, err
	}
	return ml, nil
}

// TotalMasteryLevel gets the total mastery level.
func TotalMasteryLevel(region string, summonerID int64) (ml int, err error) {
	summID := strconv.FormatInt(summonerID, 10)
	if cache.Enabled == true {
		ct := "mastery_level"
		var cc cache.Cached

		er := cache.CDB.Table(ct+"_"+region).Where("call_key = ?", summID).First(&cc).Error
		switch er {
		case gorm.ErrRecordNotFound:
			if err = apiRequest("https://"+region+"."+Base+BaseMastery+"/scores/by-summoner/"+summID, &ml); err != nil {
				return ml, err
			}

			if err = cache.StoreCall(ct, region, summID, &ml); err != nil {
				return ml, err
			}

			return ml, err
		case nil:
			if time.Since(cc.UpdatedAt) > GetTotalMasteryLevelExpire {

				if err = apiRequest("https://"+region+"."+Base+BaseMastery+"/scores/by-summoner/"+summID, &ml); err != nil {
					return ml, err
				}

				cache.UpdateCall(ct, region, &cc, &ml)

				return ml, nil

			}

			jsoniter.UnmarshalFromString(cc.JSON, &ml)

			return ml, nil
		default:
			return ml, errors.New("ggriot: unknown error, please open github issue: " + er.Error())
		}
	}

	if err = apiRequest("https://"+region+"."+Base+BaseMastery+"/scores/by-summoner/"+summID, &ml); err != nil {
		return ml, err
	}

	return ml, nil
}
