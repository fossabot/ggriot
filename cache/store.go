package cache

import (
	"errors"

	"github.com/json-iterator/go"
)

// StoreCall will store the interface into JSON then store that json into the correct table.
func StoreCall(call string, region string, key string, s interface{}) error {
	js, err := jsoniter.MarshalToString(&s)
	if err != nil {
		return errors.New("ggriot: error marshalling to string, " + err.Error())
	}

	c := Cached{
		CallKey: key,
		JSON:    js,
	}

	CDB.Table(call + "_" + region).Create(&c)

	return nil
}

// UpdateCall will update the entry in the database.
func UpdateCall(call string, region string, gorm *Cached, s interface{}) error {
	js, err := jsoniter.MarshalToString(&s)
	if err != nil {
		return errors.New("ggriot: error marshalling to string, " + err.Error())
	}

	CDB.Table(call+"_"+region).Model(&gorm).Update("json", js)

	return nil
}
