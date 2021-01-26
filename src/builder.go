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
    "flag"
    "fmt"
    "os"
    "path/filepath"
)

type Config struct {
    GithubUsername string
    GithubPassword string
}

func main() {
    var version string
    var output string

    flag.StringVar(&version, "v", "", "The Kosmos Reborn version. (Required)")
    flag.StringVar(&output, "o", "", "Path of where to generate the zip file. (Required)")

    flag.Parse()
    
    if version == "" || output == "" {
        fmt.Println("Usage:")
        flag.PrintDefaults()
        return
    }

    config := GetConfig()
    if config.GithubUsername == "" || config.GithubPassword == "" {
        fmt.Println("Error: Make sure you have the following environment variables setup.")
        fmt.Printf("\tGH_USERNAME - Github Username\n")
        fmt.Printf("\tGH_PASSWORD - Github Password\n")
        return
    }

    tempDirectory := GenerateTempPath()
    os.Mkdir(tempDirectory, 0755)

    fmt.Printf("Kosmos Reborn %s built with:\n", version)

    buildMessage, err := BuildModules(tempDirectory, version)

    os.RemoveAll(output)

    if err == nil {
        // TODO: Compress all the files together.

        fmt.Println(buildMessage)
    }

    cwd, _ := os.Getwd()
    os.RemoveAll(filepath.Join(cwd, "tmp"))        
}

func GetConfig() Config {
    return Config {
        os.Getenv("GH_USERNAME"),
        os.Getenv("GH_PASSWORD"),
    }
}
