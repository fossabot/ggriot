package ggriot

import (
	"fmt"
	"log"
	"testing"

	"github.com/soowan/ggriot/cache"
)

func TestActiveGame(t *testing.T) {
	SetAPIKey("")
	cache.UseCache("")
}

func TestGetActiveGame(t *testing.T) {
	e, err := GetMasteryList(NA, "21669611")
	if err != nil {
		t.Error(err)
	}

	log.Println((*e)[0].PlayerID)
}

func TestGetTotalMasteryLevel(t *testing.T) {
	e, err := GetTotalMasteryLevel(NA, "21669611")
	if err != nil {
		t.Error(err)
	}

	log.Println(e)
}

func TestGetChallengers(t *testing.T) {
	e, err := GetChallengers(NA, Ranked5s)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(e.Name)
}

func TestGetMasters(t *testing.T) {
	e, err := GetMasters(NA, Flex3s)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(e.Name)
}

func TestGetPlayerPosition(t *testing.T) {
	e, err := GetPlayerPosition(NA, "21669611")
	if err != nil {
		t.Error(err)
	}

	fmt.Println((*e)[0].LeagueName)
}

func TestGetMatch(t *testing.T) {
	e, err := GetMatch(NA, "2872782472")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(e.GameVersion)
}

func TestGetMatchTimeline(t *testing.T) {
	e, err := GetMatchTimeline(NA, "2872782472")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(e.Frames[10].ParticipantFrames.Num1.CurrentGold)
}
