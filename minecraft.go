package twitchappapi

import (
	"fmt"
	"net/http"
	"time"
)

const minecraftVersionTimestampURL = apiURL + "/minecraft/version/timestamp"
const minecraftVersionListURL = apiURL + "/minecraft/version"
const minecraftVersionInfoURL = apiURL + "/minecraft/version/%s"

type MinecraftVersion struct {
	ID                    int       `json:"id"`
	GameVersionID         int       `json:"gameVersionId"`
	VersionString         string    `json:"versionString"`
	JarDownloadURL        string    `json:"jarDownloadUrl"`
	JSONDownloadURL       string    `json:"jsonDownloadUrl"`
	Approved              bool      `json:"approved"`
	DateModified          time.Time `json:"dateModified"`
	GameVersionTypeID     int       `json:"gameVersionTypeId"`
	GameVersionStatus     int       `json:"gameVersionStatus"`
	GameVersionTypeStatus int       `json:"gameVersionTypeStatus"`
}

func MinecraftVersionTimestamp() (time.Time, error) {
	var versionTimestamp time.Time

	if err := requestJSONData(http.MethodGet, minecraftVersionTimestampURL, nil, &versionTimestamp); err != nil {
		return versionTimestamp, nil
	}

	return versionTimestamp, nil
}

func MinecraftVersionList() ([]MinecraftVersion, error) {
	var minecraftVersions []MinecraftVersion

	if err := requestJSONData(http.MethodGet, minecraftVersionListURL, nil, &minecraftVersions); err != nil {
		return minecraftVersions, nil
	}

	return minecraftVersions, nil
}

func MinecraftVersionInfo(version string) (MinecraftVersion, error) {
	var minecraftVersion MinecraftVersion

	if err := requestJSONData(http.MethodGet, fmt.Sprintf(minecraftVersionInfoURL, version), nil, &minecraftVersion); err != nil {
		return minecraftVersion, nil
	}

	return minecraftVersion, nil
}
