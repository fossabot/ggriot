package ggriot

import (
	"io/ioutil"
	"net/http"

	"strconv"

	"time"

	"github.com/json-iterator/go"
	"github.com/pkg/errors"
)

var (
	// IgnoreLimit can be set to true to disable
	IgnoreLimit = false
	// ShortLimit is how many request in one second.
	ShortLimit int64 = 20
	// LongLimit is how many request in two minutes.
	LongLimit int64 = 200

	// apikey is APIkey for Riot's API. Not exported to prevent leaks.
	apikey string

	// Base is the base url w/o a server.
	Base = "api.riotgames.com/lol"

	// BaseMastery is the path for mastery api calls.
	BaseMastery = "/champion-mastery/v3"

	// BaseChampion is the path for champion api calls.
	BaseChampion = "/platform/v3/champions"

	// BaseLeague is the path for ranked leagues api calls.
	BaseLeague = "/league/v3"

	// BaseStatus is the path for service status api calls.
	BaseStatus = "/status/v3/shard-data"

	// BaseMatch is the path for match data api calls.
	BaseMatch = "/match/v3"

	// BaseSpectator is the path for spectator api calls.
	BaseSpectator = "/spectator/v3"

	// BaseSummoner is the path for summoner information api calls.
	BaseSummoner = "/summoner/v3/summoners"

	// BaseTPC is the path for using third party codes to verify summoners via api calls.
	BaseTPC = "/platform/v3/third-party-code/by-summoner"

	// BaseTournamentS is the path for using tournament functions via api calls.
	// This is using the STUB version.
	BaseTournamentS = "/tournament-stub/v3"

	// BaseTournament is the path for using tournament functions via api calls.
	// This can only be used with a production API key.
	BaseTournament = "/tournament/v3"

	// RU is the server code for the Russian servers.
	RU = "RU"
	// KR is the server code for the South Korean servers.
	KR = "KR"
	// BR is the server code for the Brazil servers.
	BR = "BR1"
	// OC is the server code for the Oceania servers.
	OC = "OC1"
	// JP is the server code for the Japan servers.
	JP = "JP1"
	// NA is the server code for the North America servers.
	NA = "NA1"
	// EUN is the server code for the European Union North servers.
	EUN = "EUN1"
	// EUW is the server code for the European Union West servers.
	EUW = "EUW1"
	// TR is the server code for the Turkish servers.
	TR = "TR1"
	// LA1 is the server code for the first Latin America server.
	LA1 = "LA1"
	// LA2 is the server code for the second Latin America server.
	LA2 = "LA2"

	// Ranked5s is Solo/Duo ranked.
	Ranked5s = "RANKED_SOLO_5x5"
	// Flex3s is flex for twisted treeline
	Flex3s = "RANKED_FLEX_TT"
	// Flex5s is flex for summoners rift.
	Flex5s = "RANKED_FLEX_SR"

	// NoAPIKeySet is the error that is returned if no API key was set.
	errNoAPI = errors.New("ggriot: No API key was set")
)

// SetAPIKey will set the api key for the global session.
func SetAPIKey(key string) {
	// apikey is the global api key.
	apikey = key
}

// apiRequest is the function that does all the heavy lifting.
// It also employs a rate limiter that can be changed, depending on the limits of the api key.
// This drops the connections if the limit is reached, however in the future there maybe an option to use a queue system.
func apiRequest(request string, s interface{}) (err error) {
	if CheckKeySet() == false {
		return errNoAPI
	}

	if CheckRateLimit() == false {
		return errors.New("rate limit reached")
	}

	ggriotClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, request, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "ggriot")
	req.Header.Set("X-Riot-Token", apikey)

	res, err := ggriotClient.Do(req)
	if err != nil {
		return errors.New(err.Error())
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, er := ioutil.ReadAll(res.Body)
		if er != nil {
			return errors.New("ggriot: " + er.Error())
		}

		var jsonError JSONError

		er = jsoniter.Unmarshal(body, &jsonError)
		if er != nil {
			return errors.New("ggriot: " + er.Error())
		}

		return errors.New("ggriot: HTTP Status: " + strconv.Itoa(jsonError.Status.StatusCode) + " - " + jsonError.Status.Message)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return // Add return error here.
	}
	err = jsoniter.Unmarshal(body, s)
	if err != nil {
		return
	}
	return
}

// CheckKeySet checks if an API key was set.
func CheckKeySet() bool {
	return !(apikey == "")
}

// DisableRateLimiting will stop the limiter.
func DisableRateLimiting() {
	close(StopRateLimit)
}

// EnableRateLimiting will restart the rate limiter if it was stopped.
func EnableRateLimiting() {
	if RateLimitingActive == false {
		go ShortRateTime()
		go LongRateTime()
	}
}
