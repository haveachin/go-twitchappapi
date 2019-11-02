package twitchappapi

import (
	"fmt"
	"net/http"
	"time"
)

const categoryTimestampURL = apiURL + "/category/timestamp"
const categoryListURL = apiURL + "/category"
const categoryInfoURL = apiURL + "/category/%d"

type Category struct {
	ID                   int       `json:"id"`
	Name                 string    `json:"name"`
	Slug                 string    `json:"slug"`
	AvatarURL            string    `json:"avatarUrl"`
	DateModified         time.Time `json:"dateModified"`
	ParentGameCategoryID int       `json:"parentGameCategoryId"`
	RootGameCategoryID   int       `json:"rootGameCategoryId"`
	GameID               int       `json:"gameId"`
}

func CategoryTimestamp() (time.Time, error) {
	var timestamp time.Time

	if err := requestJSONData(http.MethodGet, categoryTimestampURL, nil, &timestamp); err != nil {
		return timestamp, nil
	}

	return timestamp, nil
}

func CategoryList() ([]Category, error) {
	var categories []Category

	if err := requestJSONData(http.MethodGet, categoryListURL, nil, &categories); err != nil {
		return categories, nil
	}

	return categories, nil
}

func CategoryInfo(categoryID int) (Category, error) {
	var category Category

	if err := requestJSONData(http.MethodGet, fmt.Sprintf(categoryInfoURL, categoryID), nil, &category); err != nil {
		return category, nil
	}

	return category, nil
}
