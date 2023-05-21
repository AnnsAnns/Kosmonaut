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

type GitHubUser struct {
	Login             string `json:"login"`
	Id                int    `json:"id"`
	NodeId            string `json:"node_id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type GitHubAsset struct {
	Url                string     `json:"url"`
	Id                 int        `json:"id"`
	NodeId             string     `json:"node_id"`
	Name               string     `json:"name"`
	Label              string     `json:"label"`
	Uploader           GitHubUser `json:"uploader"`
	ContentType        string     `json:"content_type"`
	State              string     `json:"state"`
	Size               int        `json:"size"`
	DownloadCount      int        `json:"download_count"`
	CreatedAt          string     `json:"created_at"`
	UpdatedAt          string     `json:"updated_at"`
	BrowserDownloadUrl string     `json:"browser_download_url"`
}

type GitHubRelease struct {
	Url             string        `json:"url"`
	AssetsUrl       string        `json:"assets_url"`
	UploadUrl       string        `json:"upload_url"`
	HtmlUrl         string        `json:"html_url"`
	Id              int           `json:"id"`
	Author          GitHubUser    `json:"author"`
	NodeId          string        `json:"node_id"`
	TagName         string        `json:"tag_name"`
	TargetCommitish string        `json:"target_commitish"`
	Name            string        `json:"name"`
	Draft           bool          `json:"draft"`
	Prerelease      bool          `json:"prerelease"`
	CreatedAt       string        `json:"created_at"`
	PublishedAt     string        `json:"published_at"`
	Assets          []GitHubAsset `json:"assets"`
	TarballUrl      string        `json:"tarball_url"`
	ZipballUrl      string        `json:"zipball_url"`
	Body            string        `json:"body"`
}

func GetLatestGitHubRelease(organization string, repository string, assetPattern string, githubUsername string, githubPassword string) (string, string, string, error) {
	resp, err := http.Get("https://" + githubUsername + ":" + githubPassword + "@api.github.com/repos/" + organization + "/" + repository + "/releases/latest")
	if err != nil {
		return "", "", "", err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", "", "", errors.New("getting latest release for " + organization + "/" + repository + " returned status code: " + strconv.Itoa(resp.StatusCode))
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", "", err
	}

	var release GitHubRelease
	err = json.Unmarshal(body, &release)
	if err != nil {
		return "", "", "", err
	}

	for _, asset := range release.Assets {
		matched, err := regexp.Match(assetPattern, []byte(asset.Name))
		if err != nil {
			return "", "", "", err
		}

		if matched {
			return release.TagName, asset.BrowserDownloadUrl, asset.Name, nil
		}
	}

	return "", "", "", errors.New("no assets")
}
