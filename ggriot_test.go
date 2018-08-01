package ggriot

import "testing"

func init() {
	SetAPIKey("RGAPI-8596d5a2-b650-4929-bb8d-c2baf8fc2468")
}

func TestActiveGame(t *testing.T) {
	_, err := GetActiveGame(NA, "92742955")
	if err != nil {
		t.Error("Error testing getting current game,", err)
	}
}
