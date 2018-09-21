package ggriot

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/soowan/ggriot/cache"
	"github.com/soowan/ggriot/models"
)

// GetMasteryList will return a struct with all the summoners champions and mastery exp/level.
func GetMasteryList(region string, summoner string) (ml *models.MasteryList, err error) {
	if useCache == true {
		var cc cache.MasteryList

		switch cache.CDB.Where("summoner_id = ?", summoner).First(&cc).Error {
		case gorm.ErrRecordNotFound:
			if err = apiRequest("https://"+region+"."+Base+BaseMastery+"/champion-masteries/by-summoner/"+summoner, &ml); err != nil {
				return ml, err
			}

			js, er := jsoniter.Marshal(&ml)
			if er != nil {
				return ml, er
			}
			dbe := cache.MasteryList{
				SummonerID: (*ml)[0].PlayerID,
				Region:     region,
				JSON:       string(js),
			}

			cache.CDB.Create(&dbe)

			return ml, err
		case nil:
			if time.Since(cc.UpdatedAt) > time.Duration(time.Second*5) {
				log.Println("outdated")
				if err = apiRequest("https://"+region+"."+Base+BaseMastery+"/champion-masteries/by-summoner/"+summoner, &ml); err != nil {
					return ml, err
				}

				js, er := jsoniter.Marshal(&ml)
				if er != nil {
					return ml, er
				}

				cache.CDB.Model(&cc).Update("json", string(js))
			}

			jsoniter.UnmarshalFromString(cc.JSON, &ml)
			return ml, nil
		default:
			return ml, errors.New("uh unkown error with case i suppose")
		}
	}

	if err = apiRequest("https://"+region+"."+Base+BaseMastery+"/champion-masteries/by-summoner/"+summoner, &ml); err != nil {
		return ml, err
	}

	return ml, nil
}

// GetChampionMastery will return a single champion mastery struct
func GetChampionMastery(region string, summoner string, championID string) (ml *models.MasteryList, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseMastery+"/champion-masteries/by-summoner/"+summoner+"/by-champion"+championID+apikey, &ml)
	if err != nil {
		return nil, err
	}
	return ml, nil
}

// GetTotalMasteryLevel gets the total mastery level.
func GetTotalMasteryLevel(region string, summoner string) (ml *models.MasteryLevel, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseMastery+"/scores/by-summoner/"+summoner, &ml)
	if err != nil {
		return nil, err
	}
	return ml, nil
}
