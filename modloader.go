package twitchappapi

import (
	"fmt"
	"net/http"
	"time"
)

const modloaderTimestampURL = apiURL + "/minecraft/modloader/timestamp"
const modloaderListURL = apiURL + "/minecraft/modloader"
const modloaderInfoURL = apiURL + "/minecraft/modloader/%s"

type Modloader struct {
	ID                             int         `json:"id"`
	GameVersionID                  int         `json:"gameVersionId"`
	MinecraftGameVersionID         int         `json:"minecraftGameVersionId"`
	ForgeVersion                   string      `json:"forgeVersion"`
	Name                           string      `json:"name"`
	Type                           int         `json:"type"`
	DownloadURL                    string      `json:"downloadUrl"`
	Filename                       string      `json:"filename"`
	InstallMethod                  int         `json:"installMethod"`
	Latest                         bool        `json:"latest"`
	Recommended                    bool        `json:"recommended"`
	Approved                       bool        `json:"approved"`
	DateModified                   time.Time   `json:"dateModified"`
	MavenVersionString             string      `json:"mavenVersionString"`
	VersionJSON                    string      `json:"versionJson"`
	LibrariesInstallLocation       string      `json:"librariesInstallLocation"`
	MinecraftVersion               string      `json:"minecraftVersion"`
	AdditionalFilesJSON            interface{} `json:"additionalFilesJson"`
	ModLoaderGameVersionID         int         `json:"modLoaderGameVersionId"`
	ModLoaderGameVersionTypeID     int         `json:"modLoaderGameVersionTypeId"`
	ModLoaderGameVersionStatus     int         `json:"modLoaderGameVersionStatus"`
	ModLoaderGameVersionTypeStatus int         `json:"modLoaderGameVersionTypeStatus"`
	McGameVersionID                int         `json:"mcGameVersionId"`
	McGameVersionTypeID            int         `json:"mcGameVersionTypeId"`
	McGameVersionStatus            int         `json:"mcGameVersionStatus"`
	McGameVersionTypeStatus        int         `json:"mcGameVersionTypeStatus"`
	InstallProfileJSON             interface{} `json:"installProfileJson"`
}

func ModloaderTimestamp() (time.Time, error) {
	var timestamp time.Time

	if err := requestJSONData(http.MethodGet, modloaderTimestampURL, nil, &timestamp); err != nil {
		return timestamp, nil
	}

	return timestamp, nil
}

func ModloaderList() ([]Modloader, error) {
	var modloaders []Modloader

	if err := requestJSONData(http.MethodGet, modloaderListURL, nil, &modloaders); err != nil {
		return modloaders, err
	}

	return modloaders, nil
}

func ModloaderInfo(name string) (Modloader, error) {
	var modloader Modloader

	if err := requestJSONData(http.MethodGet, fmt.Sprintf(modloaderInfoURL, name), nil, &modloader); err != nil {
		return modloader, err
	}

	return modloader, nil
}
