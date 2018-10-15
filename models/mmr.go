package models

/*
PLEASE READ:
This is not using Riot Game's API for MMR, they have not, and will likely never make MMR public.

This is using the public API from whatismymmr.com.
*/

// MMR is the exported struct for mmr calls.
type MMR struct {
	Ranked struct {
		Avg      int    `json:"avg"`
		Err      int    `json:"err"`
		Warn     bool   `json:"warn"`
		Summary  string `json:"summary"`
		TierData []struct {
			Name string `json:"name"`
			Avg  int    `json:"avg"`
			Min  int    `json:"min"`
			Max  int    `json:"max"`
		} `json:"tierData"`
		Timestamp  int `json:"timestamp"`
		Historical []struct {
			Avg       int  `json:"avg"`
			Err       int  `json:"err"`
			Warn      bool `json:"warn"`
			Timestamp int  `json:"timestamp"`
		} `json:"historical"`
		HistoricalTierData []struct {
			Name     string `json:"name"`
			Avg      int    `json:"avg"`
			AvgCount int    `json:"avgCount"`
			Min      int    `json:"min"`
			Max      int    `json:"max"`
		} `json:"historicalTierData"`
	} `json:"ranked"`
	Normal struct {
		Avg        int  `json:"avg"`
		Err        int  `json:"err"`
		Warn       bool `json:"warn"`
		Timestamp  int  `json:"timestamp"`
		Historical []struct {
			Avg       int  `json:"avg"`
			Err       int  `json:"err"`
			Warn      bool `json:"warn"`
			Timestamp int  `json:"timestamp"`
		} `json:"historical"`
	} `json:"normal"`
	ARAM struct {
		Avg        int  `json:"avg"`
		Err        int  `json:"err"`
		Warn       bool `json:"warn"`
		Timestamp  int  `json:"timestamp"`
		Historical []struct {
			Avg       int  `json:"avg"`
			Err       int  `json:"err"`
			Warn      bool `json:"warn"`
			Timestamp int  `json:"timestamp"`
		} `json:"historical"`
	} `json:"ARAM"`
}
