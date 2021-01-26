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
    "io/ioutil"
    "os"
)

type Module struct {
    Name string
    Org string
    Repo string
    AssetPattern string
    Instructions []Instruction
}

func BuildModules(tempDirectory string, version string) (string, error) {
    modules, err := ioutil.ReadDir("./modules")
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

        byteValue, _ := ioutil.ReadAll(moduleFile)

        var module Module
        json.Unmarshal(byteValue, &module)

        moduleTempDirectory := GenerateTempPath()

        // TODO: Download asset.
        
        for _, instruction := range module.Instructions {
            switch (instruction.Action) {
                case Copy:
                    err = CopyInstruction(module, instruction, moduleTempDirectory, tempDirectory)
                    break

                case Delete:
                    err = DeleteInstruction(module, instruction, moduleTempDirectory, tempDirectory)
                    break

                case Extract:
                    err = ExtractInstruction(module, instruction, moduleTempDirectory, tempDirectory)
                    break

                case Mkdir:
                    err = MkdirInstruction(module, instruction, moduleTempDirectory, tempDirectory)
                    break
            }

            if err != nil {
                return "", err
            }
        }

        os.RemoveAll(moduleTempDirectory)
        buildMessage += "\t" + module.Name + ": " + "\n"
    }

    return buildMessage, nil
}
