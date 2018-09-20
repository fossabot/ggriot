package cache

import (
	"errors"

	"github.com/jinzhu/gorm"

	// This is how GORM does dialects, must use a blank import.
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// CDB is the exported pointer to the opened cache server.
var CDB *gorm.DB

// UseCache will open a connection to a postgres server that will be used as a cache server.
func UseCache() (err error) {
	CDB, err = gorm.Open("postgres", "")
	if err != nil {
		return errors.New("ggriot: error opening cache cb, " + err.Error())
	}
	return nil
}
