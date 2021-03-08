// Kosmos Reborn Builder
// Copyright (C) 2021 Nichole Mattera
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
	"archive/zip"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Action string
const (
	Copy Action = "copy"
	Delete Action = "delete"
	Extract Action = "extract"
	Mkdir Action = "mkdir"
)

type Instruction struct {
	Action Action
	Source string
	Destination string
}

func CopyInstruction(module Module, instruction Instruction, moduleTempDirectory string, tempDirectory string) error {
	matches, err := filepath.Glob(filepath.Join(moduleTempDirectory, instruction.Source))
	if err != nil {
		return err
	}

	for _, match := range matches {
		matchInfo, err := os.Stat(match)
		if err != nil {
			return err
		}

		if matchInfo.IsDir() {
			CopyDirectory(match, filepath.Join(tempDirectory, instruction.Destination))
		} else {
			CopyFile(match, filepath.Join(tempDirectory, instruction.Destination))
		}
	}

	return nil
}

func DeleteInstruction(module Module, instruction Instruction, moduleTempDirectory string, tempDirectory string) error {
	matches, err := filepath.Glob(filepath.Join(moduleTempDirectory, instruction.Source))
	if err != nil {
		return err
	}

	for _, match := range matches {
		if Exists(match) {
			err = os.Remove(match)
			if err != nil {
				return err
			}    
		}
	}

	return nil
}

func ExtractInstruction(module Module, instruction Instruction, moduleTempDirectory string, tempDirectory string) error {
	matches, err := filepath.Glob(filepath.Join(moduleTempDirectory, instruction.Source))
	if err != nil {
		return err
	}
	if len(matches) < 1 {
		return errors.New("Nothing to unzip for pattern: " + instruction.Source)
	}

	for _, match := range matches {
		zipReader, err := zip.OpenReader(match)
		if err != nil {
			return err
		}
		defer zipReader.Close()

		for _, file := range zipReader.File {
			path := filepath.Join(moduleTempDirectory, file.Name)
			if !strings.HasPrefix(path, filepath.Clean(moduleTempDirectory) + string(os.PathSeparator)) {
				return errors.New("Illegal file path: " + path)
			}

			// Extract folder
			if file.FileInfo().IsDir() {
				os.MkdirAll(path, os.ModePerm)
				continue
			}

			// Extract file
			err = os.MkdirAll(filepath.Dir(path), os.ModePerm);
			if err != nil {
				return err
			}

			outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				return err
			}

			defer outFile.Close()

			inFile, err := file.Open()
			if err != nil {
				return err
			}

			defer inFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func MkdirInstruction(module Module, instruction Instruction, moduleTempDirectory string, tempDirectory string) error {
	return os.MkdirAll(filepath.Join(tempDirectory, instruction.Destination), os.ModePerm)
}
