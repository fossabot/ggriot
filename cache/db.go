package cache

import (
	"errors"

	"github.com/jinzhu/gorm"

	// This is how GORM does dialects, must use a blank import.
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// CDB is the exported pointer to the opened cache server.
var CDB *gorm.DB

// Enabled this is checked to see if ggriot should call the postgres db first or skip calling cache.
var Enabled = false

// UseCache will open a connection to a postgres server that will be used as a cache server.
func UseCache(gostring string) (err error) {
	CDB, err = gorm.Open("postgres", gostring)
	if err != nil {
		return errors.New("ggriot: error opening cache cb, " + err.Error())
	}

	Enabled = true

	// This will create all the tables needed, I'm not sure how to make this better yet, but I know it probably
	// needs to be changed.
	r := []string{
		"RU",
		"KR",
		"BR1",
		"OC1",
		"JP1",
		"NA1",
		"EUN1",
		"EUW1",
		"TR1",
		"LA1",
		"LA2",
	}

	c := []string{
		"mastery_by_summoner",
		"mastery_by_champion",
		"mastery_level",
		"champion_rotation",
		"league_by_queue",
		"league_by_id",
		"league_master_by_queue",
		"league_challenger_by_queue",
		"league_position_by_summoner",
		"league_match_by_id",
		"league_match_tl_by_id",
		"summoner_by_ign",
		"mmr",
	}

	for rr := range r {
		for cc := range c {
			if CDB.Table(c[cc]+"_"+r[rr]).HasTable(&Cached{}) == false {
				CDB.Table(c[cc] + "_" + r[rr]).CreateTable(&Cached{})
			}
		}
	}

	return nil
}
