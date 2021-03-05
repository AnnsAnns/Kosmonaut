/*
 * Kosmos Reborn Builder
 * Copyright (C) 2021 Nichole Mattera
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 2
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA
 * 02110-1301, USA.
 */

package main

import (
    "encoding/json"
    "errors"
    "io"
    "io/ioutil"
    "net/http"
    "net/url"
    "os"
    "path/filepath"
    "regexp"
    "strconv"
    "strings"
)

type GitHubUser struct {
    Login string `json:"login"`
    Id int `json:"id"`
    NodeId string `json:"node_id"`
    AvatarUrl string `json:"avatar_url"`
    GravatarId string `json:"gravatar_id"`
    Url string `json:"url"`
    HtmlUrl string `json:"html_url"`
    FollowersUrl string `json:"followers_url"`
    FollowingUrl string `json:"following_url"`
    GistsUrl string `json:"gists_url"`
    StarredUrl string `json:"starred_url"`
    SubscriptionsUrl string `json:"subscriptions_url"`
    OrganizationsUrl string `json:"organizations_url"`
    ReposUrl string `json:"repos_url"`
    EventsUrl string `json:"events_url"`
    ReceivedEventsUrl string `json:"received_events_url"`
    Type string `json:"type"`
    SiteAdmin bool `json:"site_admin"`
}

type GitHubAsset struct {
    Url string `json:"url"`
    Id int `json:"id"`
    NodeId string `json:"node_id"`
    Name string `json:"name"`
    Label string `json:"label"`
    Uploader GitHubUser `json:"uploader"`
    ContentType string `json:"content_type"`
    State string `json:"state"`
    Size int `json:"size"`
    DownloadCount int `json:"download_count"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
    BrowserDownloadUrl string `json:"browser_download_url"`
}
 
type GitHubRelease struct {
    Url string `json:"url"`
    AssetsUrl string `json:"assets_url"`
    UploadUrl string `json:"upload_url"`
    HtmlUrl string `json:"html_url"`
    Id int `json:"id"`
    Author GitHubUser `json:"author"`
    NodeId string `json:"node_id"`
    TagName string `json:"tag_name"`
    TargetCommitish string `json:"target_commitish"`
    Name string `json:"name"`
    Draft bool `json:"draft"`
    Prerelease bool `json:"prerelease"`
    CreatedAt string `json:"created_at"`
    PublishedAt string `json:"published_at"`
    Assets []GitHubAsset `json:"assets"`
    TarballUrl string `json:"tarball_url"`
    ZipballUrl string `json:"zipball_url"`
    Body string `json:"body"`
}

func GetLatestRelease(organization string, repository string, assetPattern string, config Config) (string, string, error) {
    resp, err := http.Get("https://" + config.GithubUsername + ":" + config.GithubPassword + "@api.github.com/repos/" + organization + "/" + repository + "/releases/latest")
    if err != nil {
        return "", "", err
    }
    if resp.StatusCode < 200 || resp.StatusCode >= 300 {
        return "", "", errors.New("Getting latest release returned status code: " + strconv.Itoa(resp.StatusCode))
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", "", err
    }

    var release GitHubRelease
    json.Unmarshal(body, &release)

    for _, asset := range release.Assets {
        matched, err := regexp.Match(assetPattern, []byte(asset.Name))
        if err != nil {
            return "", "", err
        }

        if matched {
            return release.TagName, asset.BrowserDownloadUrl, nil
        }
    }

    return "", "", errors.New("No assets")
}

func DownloadFile(rawUrl string, destination string) (string, error) {
    url, err := url.Parse(rawUrl)
    if err != nil {
        return "", err
    }

    segments := strings.Split(url.Path, "/")
    fileName := segments[len(segments)-1]

    path := filepath.Join(destination, fileName)
    file, err := os.Create(path)
    if err != nil {
        return "", err
    }

    defer file.Close()

    resp, err := http.Get(rawUrl)
    if err != nil {
        return "", err
    }
    if resp.StatusCode < 200 || resp.StatusCode >= 300 {
        return "", errors.New("Download file returned status code: " + strconv.Itoa(resp.StatusCode))
    }

    defer resp.Body.Close()

    _, err = io.Copy(file, resp.Body)
    if err != nil {
        return "", err
    }

    return path, nil
}
