package cache

import (
	"github.com/json-iterator/go"
	"github.com/soowan/ggriot/models"
)

func StoreMasteryList(s *models.MasteryList, region string) (err error) {
	js, er := jsoniter.Marshal(&s)
	if er != nil {
		return er
	}
	dbe := MasteryList{
		SummonerID: (*s)[0].PlayerID,
		Region:     region,
		JSON:       string(js),
	}

	CDB.Create(&dbe)
	return nil
}
