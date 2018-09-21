package ggriot

import (
	"fmt"
	"log"
	"time"

	"github.com/soowan/ggriot/models"

	"github.com/jinzhu/gorm"

	"github.com/json-iterator/go"
	"github.com/soowan/ggriot/cache"
)

// GetChallengers will return all the challengers in the requested queue.
// Valid queues are, Ranked5s(RANKED_SOLO_5x5), Flex3s(RANKED_FLEX_TT), and Flex5s(RANKED_FLEX_SR)
func GetChallengers(region string, mode string) (lr *models.LeagueRoster, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseLeague+"/challengerleagues/by-queue/"+mode, &lr)
	if err != nil {
		return nil, err
	}

	return lr, nil
}

// GetMasters returns the roster of all the players in the Masters tier for requested queue.
func GetMasters(region string, mode string) (lr *models.LeagueRoster, err error) {
	err = apiRequest("https://"+region+"."+Base+BaseLeague+"/masterleagues/by-queue/"+mode, &lr)
	if err != nil {
		return nil, err
	}

	return lr, nil
}

// GetLeague will return the roster of the league requested, via UUID.
// You can and will get blocked from this call if you provide invalid UUIDs
func GetLeague(region string, leagueUUID string) (lr *models.LeagueRoster, err error) {
	switch useCache {
	case true:
		var clt cache.LeagueTier
		if c := cache.CDB.Where("league_id = ?", leagueUUID).First(&clt); c.Error != nil {
			switch c.Error {
			case gorm.ErrRecordNotFound:
				if err = apiRequest("https://"+region+"."+Base+BaseLeague+"/leagues/"+leagueUUID, &lr); err != nil {
					return &models.LeagueRoster{}, err
				}
				go func() {
					js, er := jsoniter.Marshal(&lr)
					if er != nil {
						log.Println("error marshal", er)
					}

					dbe := cache.LeagueTier{
						LeagueID: lr.LeagueID,
						Region:   region,
						JSON:     string(js),
					}

					cache.CDB.Create(&dbe)
				}()
				return lr, nil
			default:
				return &models.LeagueRoster{}, c.Error
			}
		}
		fmt.Println("age of cached entry:", time.Since(clt.CreatedAt))
		if er := jsoniter.UnmarshalFromString(clt.JSON, &lr); er != nil {
			log.Println(er)
			return &models.LeagueRoster{}, er
		}
		return lr, nil
	}
	if err = apiRequest("https://"+region+"."+Base+BaseLeague+"/leagues/"+leagueUUID, &lr); err != nil {
		return &models.LeagueRoster{}, err
	}
	return lr, nil
}

// GetPlayerPosition will return the requested players league position in each of the three ranked queues.
func GetPlayerPosition(region string, summonerID string) (lp *models.LeaguePosition, err error) {
	if err = apiRequest("https://"+region+"."+Base+BaseLeague+"/positions/by-summoner/"+summonerID, &lp); err != nil {
		return nil, err
	}

	return lp, nil
}
