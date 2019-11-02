package twitchappapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const addonInfoURL = apiURL + "/addon/%d"
const addonMultipleURL = apiURL + "/addon"
const addonSearchURL = apiURL + "/addon/search?categoryId=%d&gameId=%d&gameVersion=%s&index=%d&pageSize=%d&searchFilter=%s&sectionId=%d&sort=%d"
const addonDescriptionURL = apiURL + "/addon/%d/description"
const addonFileChangelogURL = apiURL + "/addon/%d/file/%d/changelog"
const addonFileInfoURL = apiURL + "/addon/%d/file/%d"
const addonFileDownloadURL = apiURL + "/addon/%d/file/%d/download-url"
const addonFilesURL = apiURL + "/addon/%d/files"
const addonFeaturedURL = apiURL + "/addon/featured"
const addonDatabaseTimestampURL = apiURL + "/addon/timestamp"
const addonByFingerprintURL = apiURL + "/fingerprint"

type Addon struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Authors []struct {
		Name              string      `json:"name"`
		URL               string      `json:"url"`
		ProjectID         int         `json:"projectId"`
		ID                int         `json:"id"`
		ProjectTitleID    interface{} `json:"projectTitleId"`
		ProjectTitleTitle interface{} `json:"projectTitleTitle"`
		UserID            int         `json:"userId"`
		TwitchID          int         `json:"twitchId"`
	} `json:"authors"`
	Attachments []struct {
		ID           int    `json:"id"`
		ProjectID    int    `json:"projectId"`
		Description  string `json:"description"`
		IsDefault    bool   `json:"isDefault"`
		ThumbnailURL string `json:"thumbnailUrl"`
		Title        string `json:"title"`
		URL          string `json:"url"`
		Status       int    `json:"status"`
	} `json:"attachments"`
	WebsiteURL    string  `json:"websiteUrl"`
	GameID        int     `json:"gameId"`
	Summary       string  `json:"summary"`
	DefaultFileID int     `json:"defaultFileId"`
	DownloadCount float64 `json:"downloadCount"`
	LatestFiles   []File  `json:"latestFiles"`
	Categories    []struct {
		CategoryID int    `json:"categoryId"`
		Name       string `json:"name"`
		URL        string `json:"url"`
		AvatarURL  string `json:"avatarUrl"`
		ParentID   int    `json:"parentId"`
		RootID     int    `json:"rootId"`
		ProjectID  int    `json:"projectId"`
		AvatarID   int    `json:"avatarId"`
		GameID     int    `json:"gameId"`
	} `json:"categories"`
	Status            int `json:"status"`
	PrimaryCategoryID int `json:"primaryCategoryId"`
	CategorySection   struct {
		ID                      int         `json:"id"`
		GameID                  int         `json:"gameId"`
		Name                    string      `json:"name"`
		PackageType             int         `json:"packageType"`
		Path                    string      `json:"path"`
		InitialInclusionPattern string      `json:"initialInclusionPattern"`
		ExtraIncludePattern     interface{} `json:"extraIncludePattern"`
		GameCategoryID          int         `json:"gameCategoryId"`
	} `json:"categorySection"`
	Slug                   string `json:"slug"`
	GameVersionLatestFiles []struct {
		GameVersion       string      `json:"gameVersion"`
		ProjectFileID     int         `json:"projectFileId"`
		ProjectFileName   string      `json:"projectFileName"`
		FileType          int         `json:"fileType"`
		GameVersionFlavor interface{} `json:"gameVersionFlavor"`
	} `json:"gameVersionLatestFiles"`
	IsFeatured         bool      `json:"isFeatured"`
	PopularityScore    float64   `json:"popularityScore"`
	GamePopularityRank int       `json:"gamePopularityRank"`
	PrimaryLanguage    string    `json:"primaryLanguage"`
	GameSlug           string    `json:"gameSlug"`
	GameName           string    `json:"gameName"`
	PortalName         string    `json:"portalName"`
	DateModified       time.Time `json:"dateModified"`
	DateCreated        time.Time `json:"dateCreated"`
	DateReleased       time.Time `json:"dateReleased"`
	IsAvailable        bool      `json:"isAvailable"`
	IsExperiemental    bool      `json:"isExperiemental"`
}

type AddonSearchMask struct {
	CategoryID   int
	GameID       int
	GameVersion  string
	Index        int
	PageSize     int
	SearchFilter string
	SectionID    int
	Sort         int
}

type File struct {
	ID              int       `json:"id"`
	DisplayName     string    `json:"displayName"`
	FileName        string    `json:"fileName"`
	FileDate        time.Time `json:"fileDate"`
	FileLength      int       `json:"fileLength"`
	ReleaseType     int       `json:"releaseType"`
	FileStatus      int       `json:"fileStatus"`
	DownloadURL     string    `json:"downloadUrl"`
	IsAlternate     bool      `json:"isAlternate"`
	AlternateFileID int       `json:"alternateFileId"`
	Dependencies    []struct {
		AddonID int `json:"addonId"`
		Type    int `json:"type"`
	} `json:"dependencies"`
	IsAvailable bool `json:"isAvailable"`
	Modules     []struct {
		Foldername  string `json:"foldername"`
		Fingerprint int64  `json:"fingerprint"`
	} `json:"modules"`
	PackageFingerprint      int         `json:"packageFingerprint"`
	GameVersion             []string    `json:"gameVersion"`
	InstallMetadata         interface{} `json:"installMetadata"`
	ServerPackFileID        interface{} `json:"serverPackFileId"`
	HasInstallScript        bool        `json:"hasInstallScript"`
	GameVersionDateReleased time.Time   `json:"gameVersionDateReleased"`
	GameVersionFlavor       interface{} `json:"gameVersionFlavor"`
}

type AddonsFeaturedMask struct {
	GameID        int           `json:"GameId"`
	AddonIds      []interface{} `json:"addonIds"`
	FeaturedCount int           `json:"featuredCount"`
	PopularCount  int           `json:"popularCount"`
	UpdatedCount  int           `json:"updatedCount"`
}

type FeaturedAddons struct {
	Featured        []Addon `json:"Featured"`
	Popular         []Addon `json:"Popular"`
	RecentlyUpdated []Addon `json:"RecentlyUpdated"`
}

type Fingerprints struct {
	IsCacheBuilt bool `json:"isCacheBuilt"`
	ExactMatches []struct {
		ID          int    `json:"id"`
		File        File   `json:"file"`
		LatestFiles []File `json:"latestFiles"`
	} `json:"exactMatches"`
	ExactFingerprints        []int64 `json:"exactFingerprints"`
	PartialMatches           []int64 `json:"partialMatches"`
	PartialMatchFingerprints struct {
	} `json:"partialMatchFingerprints"`
	InstalledFingerprints []int64 `json:"installedFingerprints"`
	UnmatchedFingerprints []int64 `json:"unmatchedFingerprints"`
}

func AddonInfo(addonID int) (Addon, error) {
	var addon Addon

	if err := requestJSONData(http.MethodGet, fmt.Sprintf(addonInfoURL, addonID), nil, &addon); err != nil {
		return addon, err
	}

	return addon, nil
}

func AddonMultiple(addonIDs ...int) ([]Addon, error) {
	var addons []Addon

	body, err := json.Marshal(addonIDs)
	if err != nil {
		return addons, err
	}

	if err := requestJSONData(http.MethodPost, addonMultipleURL, body, &addons); err != nil {
		return addons, err
	}

	return addons, nil
}

func AddonSearch(mask AddonSearchMask) ([]Addon, error) {
	var addons []Addon
	url := fmt.Sprintf(addonSearchURL, mask.CategoryID, mask.GameID, mask.GameVersion, mask.Index, mask.PageSize, mask.SearchFilter, mask.SectionID, mask.Sort)

	if err := requestJSONData(http.MethodGet, url, nil, &addons); err != nil {
		return addons, err
	}

	return addons, nil
}

func AddonDescription(addonID int) (string, error) {
	return requestPlainText(http.MethodGet, fmt.Sprintf(addonDescriptionURL, addonID), nil)
}

func AddonFileChangelog(addonID, fileID int) (string, error) {
	return requestPlainText(http.MethodGet, fmt.Sprintf(addonFileChangelogURL, addonID, fileID), nil)
}

func AddonFileInformation(addonID, fileID int) (File, error) {
	var file File

	if err := requestJSONData(http.MethodGet, fmt.Sprintf(addonFileInfoURL, addonID, fileID), nil, &file); err != nil {
		return file, err
	}

	return file, nil
}

func AddonFileDownloadURL(addonID, fileID int) (string, error) {
	return requestPlainText(http.MethodGet, fmt.Sprintf(addonFileDownloadURL, addonID, fileID), nil)
}

func AddonFiles(addonID int) ([]File, error) {
	var files []File

	if err := requestJSONData(http.MethodGet, fmt.Sprintf(addonFilesURL, addonID), nil, &files); err != nil {
		return files, err
	}

	return files, nil
}

func AddonsFeatured(mask AddonsFeaturedMask) (FeaturedAddons, error) {
	var featuredAddons FeaturedAddons

	body, err := json.Marshal(mask)
	if err != nil {
		return featuredAddons, err
	}

	if err := requestJSONData(http.MethodPost, addonFeaturedURL, body, &featuredAddons); err != nil {
		return featuredAddons, err
	}

	return featuredAddons, nil
}

func AddonsDatabaseTimestamp() (time.Time, error) {
	var timestamp time.Time

	if err := requestJSONData(http.MethodGet, addonDatabaseTimestampURL, nil, &timestamp); err != nil {
		return timestamp, nil
	}

	return timestamp, nil
}

func AddonByFingerprint(addonID int) (Fingerprints, error) {
	var fingerprints Fingerprints

	body, err := json.Marshal([]int{addonID})
	if err != nil {
		return fingerprints, err
	}

	if err := requestJSONData(http.MethodPost, addonByFingerprintURL, body, &fingerprints); err != nil {
		return fingerprints, nil
	}

	return fingerprints, nil
}
