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

type Action int
const (
    Copy Action = iota
    Delete
    Extract
    Mkdir
)

func (a Action) String() string {
    return [...]string {
        "copy",
        "delete",
        "extract",
        "mkdir",
    }[a]
}

type Instruction struct {
    Action Action
    Source string
    Destination string
}

type Module struct {
    Name string
    Org string
    Repo string
    AssetPattern string
    Instructions []Instruction
}

func BuildModules(tempDirectory string, version string, output string) (string, error) {
    return "", nil
}
