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
	e, err := MasteryList(NA, "21669611")
	if err != nil {
		t.Error(err)
	}

	log.Println((*e)[0].PlayerID)
}

func TestGetTotalMasteryLevel(t *testing.T) {
	e, err := TotalMasteryLevel(NA, "21669611")
	if err != nil {
		t.Error(err)
	}

	log.Println(e)
}

func TestGetChallengers(t *testing.T) {
	e, err := Challengers(NA, Ranked5s)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(e.Name)
}

func TestGetMasters(t *testing.T) {
	e, err := Masters(NA, Flex3s)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(e.Name)
}

func TestGetPlayerPosition(t *testing.T) {
	e, err := PlayerPosition(NA, "21669611")
	if err != nil {
		t.Error(err)
	}

	fmt.Println((*e)[0].LeagueName)
}

func TestGetMatch(t *testing.T) {
	e, err := Match(NA, "2872782472")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(e.GameVersion)
}

func TestGetMatchTimeline(t *testing.T) {
	e, err := MatchTimeline(NA, "2872782472")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(e.Frames[10].ParticipantFrames.Num1.CurrentGold)
}
