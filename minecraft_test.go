package twitchappapi

import (
	"testing"
)

func TestMinecraftVersionTimestamp(t *testing.T) {
	if _, err := MinecraftVersionTimestamp(); err != nil {
		t.Error(err)
	}
}

func TestMinecraftVersionList(t *testing.T) {
	if _, err := MinecraftVersionList(); err != nil {
		t.Error(err)
	}
}

func TestMinecraftVersionInfo(t *testing.T) {
	if _, err := MinecraftVersionInfo("1.12.2"); err != nil {
		t.Error(err)
	}
}
