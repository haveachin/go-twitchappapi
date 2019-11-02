package twitchappapi

import "testing"

func TestGameTimestamp(t *testing.T) {
	if _, err := GameTimestamp(); err != nil {
		t.Error(err)
	}
}

func TestGameList(t *testing.T) {
	if _, err := GameList(); err != nil {
		t.Error(err)
	}
}

func TestGameInfo(t *testing.T) {
	if _, err := GameInfo(423); err != nil {
		t.Error(err)
	}
}
