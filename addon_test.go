package twitchappapi

import (
	"testing"
)

func TestAddonInfo(t *testing.T) {
	if _, err := AddonInfo(310806); err != nil {
		t.Error(err)
	}
}

func TestAddonMultiple(t *testing.T) {
	if _, err := AddonMultiple(310806, 304026); err != nil {
		t.Error(err)
	}
}

func TestAddonSearch(t *testing.T) {
	searchMask := AddonSearchMask{
		CategoryID:   0,
		GameID:       432,
		GameVersion:  "1.12.2",
		Index:        0,
		PageSize:     25,
		SearchFilter: "ultimate",
		SectionID:    4471,
		Sort:         0,
	}

	if _, err := AddonSearch(searchMask); err != nil {
		t.Error(err)
	}
}

func TestAddonDescription(t *testing.T) {
	if _, err := AddonDescription(296062); err != nil {
		t.Error(err)
	}
}

func TestAddonFileChangelog(t *testing.T) {
	if _, err := AddonFileChangelog(296062, 2657461); err != nil {
		t.Error(err)
	}
}

func TestAddonFileInformation(t *testing.T) {
	if _, err := AddonFileInformation(310806, 2657461); err != nil {
		t.Error(err)
	}
}

func TestAddonFileDownloadURL(t *testing.T) {
	if _, err := AddonFileDownloadURL(296062, 2657461); err != nil {
		t.Error(err)
	}
}

func TestAddonFiles(t *testing.T) {
	if _, err := AddonFiles(304026); err != nil {
		t.Error(err)
	}
}

func TestAddonsFeatured(t *testing.T) {
	mask := AddonsFeaturedMask{
		GameID:        432,
		AddonIds:      nil,
		FeaturedCount: 6,
		PopularCount:  14,
		UpdatedCount:  14,
	}

	if _, err := AddonsFeatured(mask); err != nil {
		t.Error(err)
	}
}

func TestAddonsDatabaseTimestamp(t *testing.T) {
	if _, err := AddonsDatabaseTimestamp(); err != nil {
		t.Error(err)
	}
}

func TestAddonByFingerprint(t *testing.T) {
	if _, err := AddonByFingerprint(304026); err != nil {
		t.Error(err)
	}
}
