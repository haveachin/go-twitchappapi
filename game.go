package twitchappapi

import (
	"fmt"
	"net/http"
	"time"
)

const gameTimestampURL = apiURL + "/game/timestamp"
const gameListURL = apiURL + "/game"
const gameInfoURL = apiURL + "/game/%d"

type Game struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Slug         string    `json:"slug"`
	DateModified time.Time `json:"dateModified"`
	GameFiles    []struct {
		ID           int    `json:"id"`
		GameID       int    `json:"gameId"`
		IsRequired   bool   `json:"isRequired"`
		FileName     string `json:"fileName"`
		FileType     int    `json:"fileType"`
		PlatformType int    `json:"platformType"`
	} `json:"gameFiles"`
	GameDetectionHints []struct {
		ID          int         `json:"id"`
		HintType    int         `json:"hintType"`
		HintPath    string      `json:"hintPath"`
		HintKey     interface{} `json:"hintKey"`
		HintOptions int         `json:"hintOptions"`
		GameID      int         `json:"gameId"`
	} `json:"gameDetectionHints"`
	FileParsingRules []interface{} `json:"fileParsingRules"`
	CategorySections []struct {
		ID                      int    `json:"id"`
		GameID                  int    `json:"gameId"`
		Name                    string `json:"name"`
		PackageType             int    `json:"packageType"`
		Path                    string `json:"path"`
		InitialInclusionPattern string `json:"initialInclusionPattern"`
		ExtraIncludePattern     string `json:"extraIncludePattern"`
		GameCategoryID          int    `json:"gameCategoryId"`
	} `json:"categorySections"`
	MaxFreeStorage                 int         `json:"maxFreeStorage"`
	MaxPremiumStorage              int         `json:"maxPremiumStorage"`
	MaxFileSize                    int         `json:"maxFileSize"`
	AddonSettingsFolderFilter      interface{} `json:"addonSettingsFolderFilter"`
	AddonSettingsStartingFolder    interface{} `json:"addonSettingsStartingFolder"`
	AddonSettingsFileFilter        interface{} `json:"addonSettingsFileFilter"`
	AddonSettingsFileRemovalFilter interface{} `json:"addonSettingsFileRemovalFilter"`
	SupportsAddons                 bool        `json:"supportsAddons"`
	SupportsPartnerAddons          bool        `json:"supportsPartnerAddons"`
	SupportedClientConfiguration   int         `json:"supportedClientConfiguration"`
	SupportsNotifications          bool        `json:"supportsNotifications"`
	ProfilerAddonID                int         `json:"profilerAddonId"`
	TwitchGameID                   int         `json:"twitchGameId"`
	ClientGameSettingsID           int         `json:"clientGameSettingsId"`
}

func GameTimestamp() (time.Time, error) {
	var timestamp time.Time

	if err := requestJSONData(http.MethodGet, gameTimestampURL, nil, &timestamp); err != nil {
		return timestamp, nil
	}

	return timestamp, nil
}

func GameList() ([]Game, error) {
	var games []Game

	if err := requestJSONData(http.MethodGet, gameListURL, nil, &games); err != nil {
		return games, nil
	}

	return games, nil
}

func GameInfo(gameID int) (Game, error) {
	var game Game

	if err := requestJSONData(http.MethodGet, fmt.Sprintf(gameInfoURL, gameID), nil, &game); err != nil {
		return game, nil
	}

	return game, nil
}
