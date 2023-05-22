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
	"os"
	"path/filepath"
	"strconv"
)

type Module struct {
	Source       string
	Name         string
	Org          string
	Repo         string
	AssetPattern string
	Instructions []Instruction
}

func BuildModules(tempDirectory string, version string, githubUsername string, githubPassword string) (string, error) {
	modules, err := os.ReadDir("./modules")
	if err != nil {
		return "", err
	}

	buildMessage := ""

	for _, file := range modules {
		moduleFile, err := os.Open("./modules/" + file.Name())
		if err != nil {
			return "", err
		}
		defer moduleFile.Close()

		byteValue, _ := io.ReadAll(moduleFile)

		var module Module
		json.Unmarshal(byteValue, &module)

		moduleTempDirectory := GenerateTempPath()
		os.MkdirAll(moduleTempDirectory, os.ModePerm)

		version, downloadURL, fileName, err := GetLatestRelease(module.Source, module.Org, module.Repo, module.AssetPattern, githubUsername, githubPassword)
		if err != nil {
			return "", err
		}

		_, err = DownloadFile(downloadURL, moduleTempDirectory, fileName)
		if err != nil {
			return "", err
		}

		for _, instruction := range module.Instructions {
			switch instruction.Action {
			case Copy:
				err = CopyInstruction(module, instruction, moduleTempDirectory, tempDirectory)

			case Delete:
				err = DeleteInstruction(module, instruction, moduleTempDirectory, tempDirectory)

			case Extract:
				err = ExtractInstruction(module, instruction, moduleTempDirectory, tempDirectory)

			case Mkdir:
				err = MkdirInstruction(module, instruction, moduleTempDirectory, tempDirectory)
			}

			if err != nil {
				return "", err
			}
		}

		os.RemoveAll(moduleTempDirectory)
		buildMessage += "\t" + module.Name + ": " + version + "\n"
	}

	return buildMessage, nil
}

func GetLatestRelease(source string, organization string, repository string, assetPattern string, githubUsername string, githubPassword string) (string, string, string, error) {
	if source == "NicholeMattera" {
		return GetLatestGiteaRelease("git.nicholemattera.com", organization, repository, assetPattern)
	}

	return GetLatestGitHubRelease(organization, repository, assetPattern, githubUsername, githubPassword)
}

func DownloadFile(rawUrl string, destination string, fileName string) (string, error) {
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
		return "", errors.New("download file returned status code: " + strconv.Itoa(resp.StatusCode))
	}

	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}

	return path, nil
}
