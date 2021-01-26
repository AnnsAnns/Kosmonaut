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
    "os"
    "path/filepath"
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
        err = os.Remove(match)
        if err != nil {
            return err
        }
    }

    return nil
}

func ExtractInstruction(module Module, instruction Instruction, moduleTempDirectory string, tempDirectory string) error {
    return nil
}

func MkdirInstruction(module Module, instruction Instruction, moduleTempDirectory string, tempDirectory string) error {
    return os.Mkdir(filepath.Join(tempDirectory, instruction.Destination), 0755)
}
