// Kosmos Reborn Builder
// Copyright (C) 2022 Nichole Mattera
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA
// 02110-1301, USA.

package internal

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

type GiteaUser struct {
	Active            bool   `json:"active"`
	AvatarUrl         string `json:"avatar_url"`
	Created           string `json:"created"`
	Description       string `json:"description"`
	Email             string `json:"email"`
	FullName          string `json:"full_name"`
	FollowersCount    int    `json:"followers_count"`
	FollowingCount    int    `json:"following_count"`
	Id                int    `json:"id"`
	IsAdmin           bool   `json:"is_admin"`
	Language          string `json:"language"`
	LastLogin         string `json:"last_login"`
	Location          string `json:"location"`
	Login             string `json:"login"`
	ProhibitLogin     bool   `json:"prohibit_login"`
	Restricted        bool   `json:"restricted"`
	StarredReposCount int    `json:"starred_repos_count"`
	Username          string `json:"username"`
	Visibility        string `json:"visibility"`
	Website           string `json:"website"`
}

type GiteaAsset struct {
	BrowserDownloadUrl string `json:"browser_download_url"`
	CreatedAt          string `json:"created_at"`
	DownloadCount      int    `json:"download_count"`
	Id                 int    `json:"id"`
	Name               string `json:"name"`
	Size               int    `json:"size"`
	UUID               string `json:"uuid"`
}

type GiteaRelease struct {
	Assets          []GiteaAsset `json:"assets"`
	Author          GiteaUser    `json:"author"`
	Body            string       `json:"body"`
	CreatedAt       string       `json:"created_at"`
	Draft           bool         `json:"draft"`
	HtmlUrl         string       `json:"html_url"`
	Id              int          `json:"id"`
	Name            string       `json:"name"`
	Prerelease      bool         `json:"prerelease"`
	PublishedAt     string       `json:"published_at"`
	TagName         string       `json:"tag_name"`
	TarballUrl      string       `json:"tarball_url"`
	TargetCommitish string       `json:"target_commitish"`
	Url             string       `json:"url"`
	ZipballUrl      string       `json:"zipball_url"`
}

func GetLatestGiteaRelease(domain string, organization string, repository string, assetPattern string) (string, string, string, error) {
	resp, err := http.Get("https://" + domain + "/api/v1/repos/" + organization + "/" + repository + "/releases?limit=1")
	if err != nil {
		return "", "", "", err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", "", "", errors.New("getting releases for " + organization + "/" + repository + " returned status code: " + strconv.Itoa(resp.StatusCode))
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", "", err
	}

	var releases []GiteaRelease
	err = json.Unmarshal(body, &releases)
	if err != nil {
		return "", "", "", err
	}

	if len(releases) != 1 {
		return "", "", "", errors.New("no releases for " + organization + "/" + repository)
	}

	for _, asset := range releases[0].Assets {
		matched, err := regexp.Match(assetPattern, []byte(asset.Name))
		if err != nil {
			return "", "", "", err
		}

		if matched {
			return releases[0].TagName, asset.BrowserDownloadUrl, asset.Name, nil
		}
	}

	return "", "", "", errors.New("no assets")
}
