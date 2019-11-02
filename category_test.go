package twitchappapi

import "testing"

func TestCategoryTimestamp(t *testing.T) {
	if _, err := CategoryTimestamp(); err != nil {
		t.Error(err)
	}
}

func TestCategoryList(t *testing.T) {
	if _, err := CategoryList(); err != nil {
		t.Error(err)
	}
}

func TestCategoryInfo(t *testing.T) {
	if _, err := CategoryInfo(423); err != nil {
		t.Error(err)
	}
}
