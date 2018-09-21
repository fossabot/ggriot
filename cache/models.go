package cache

import "github.com/jinzhu/gorm"

/*

This file will contain how data is stored in the database.
Here's a outline on how it'll be stored.

* General data like, IDs and Name will be stored in relevant values
* The entire json response will be stored in the table as well.

*/

type LeagueTier struct {
	gorm.Model

	LeagueID string `sql:"string"`
	Region   string `sql:"string"`
	JSON     string `sql:"jsonb"`
}

type MasteryList struct {
	gorm.Model

	SummonerID int    `sql:"int"`
	Region     string `sql:"string"`
	JSON       string `pq:"jsonb"`
}
