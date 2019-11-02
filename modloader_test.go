package twitchappapi

import "testing"

func TestModloaderTimestamp(t *testing.T) {
	if _, err := ModloaderTimestamp(); err != nil {
		t.Error(err)
	}
}

func TestModloaderList(t *testing.T) {
	if _, err := ModloaderList(); err != nil {
		t.Error(err)
	}
}

func TestModloaderInfo(t *testing.T) {
	if _, err := ModloaderInfo("forge-9.11.1.965"); err != nil {
		t.Error(err)
	}
}
