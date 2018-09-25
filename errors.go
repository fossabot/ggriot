package ggriot

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/json-iterator/go"
)

var (
	// ErrDataNotFound is the error returned when data is not found.
	ErrDataNotFound = errors.New("ggriot: 404 - Data not found")
	// ErrNoPlayerData is the error which is similar to 404 but used when looking at old summoners.
	ErrNoPlayerData = errors.New("ggriot: 422 - Player exists, but hasn't played since match history collection began")
	// ErrRateLimitExceed this is the error message that happens when the hard limit happens.
	ErrRateLimitExceed = errors.New("ggriot: 429 - Hard rate limit exceeded")
	// ErrSoftRateLimitExceed is the error when the soft limit happens.
	ErrSoftRateLimitExceed = errors.New("ggriot: Soft rate limit reached")
)

func returnErr(res *http.Response) error {
	var jsonError JSONError

	er := jsoniter.NewDecoder(res.Body).Decode(&jsonError)
	if er != nil {
		return errors.New("ggriot: " + er.Error())
	}

	switch jsonError.Status.StatusCode {
	case 404:
		return ErrDataNotFound
	case 422:
		return ErrNoPlayerData
	case 429:
		return ErrRateLimitExceed
	}

	return errors.New("ggriot: HTTP Status: " + strconv.Itoa(jsonError.Status.StatusCode) + " - " + jsonError.Status.Message)
}
