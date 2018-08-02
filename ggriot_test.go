package ggriot

import (
	"testing"
)

func TestActiveGame(t *testing.T) {
	SetAPIKey("RGAPI-6db8b12d-b073-4192-ae73-6fe228bd48af")
	_, err := GetActiveGame(NA, "19887289")
	if err != nil {
		t.Error("Error testing getting current game,", err)
	}
}

func TestMasteryList(t *testing.T) {
	_, err := GetMasteryList(NA, "19887289")
	if err != nil {
		t.Error("Error getting mastery list,", err)
	}
}
