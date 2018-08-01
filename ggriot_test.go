package ggriot

import (
	"sync"
	"testing"
	"time"
)

func TestActiveGame(t *testing.T) {
	SetAPIKey("RGAPI-6db8b12d-b073-4192-ae73-6fe228bd48af")
	_, err := GetActiveGame(NA, "44979352")
	if err != nil {
		t.Error("Error testing getting current game,", err)
	}
}

func TestRateLimit(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i <= 20; i++ {
		wg.Add(1)
		go func() {
			_, err := GetActiveGame(NA, "44979352")
			if err != nil {
				t.Error("Error testing getting current game,", err)
			}
			wg.Done()
		}()

	}
	time.Sleep(2 * time.Second)
	for i := 0; i <= 20; i++ {
		wg.Add(1)
		go func() {
			_, err := GetActiveGame(NA, "44979352")
			if err != nil {
				t.Error("Error testing getting current game,", err)
			}
			wg.Done()
		}()
		wg.Wait()
	}
}
