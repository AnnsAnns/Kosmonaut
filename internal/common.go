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
	"archive/zip"
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/google/uuid"
)

func GenerateTempPath() string {
	cwd, _ := os.Getwd()
	return filepath.Join(cwd, "tmp", uuid.New().String())
}

func Compress(src, dst string) error {
	zipFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		f, err := zipWriter.Create(path[len(src) + 1:])
		if err != nil {
			return err
		}

		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	return nil
}

func CopyFile(src, dst string) error {
	srcfd, err := os.Open(src);
	if err != nil {
		return err
	}
	defer srcfd.Close()

	dstfd, err := os.Create(dst);
	if err != nil {
		return err
	}
	defer dstfd.Close()

	_, err = io.Copy(dstfd, srcfd);
	if err != nil {
		return err
	}

	srcinfo, err := os.Stat(src);
	if err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

func CopyDirectory(src, dst string) error {
	srcinfo, err := os.Stat(src);
	if err != nil {
		return err
	}

	err = os.MkdirAll(dst, srcinfo.Mode());
	if err != nil {
		return err
	}

	fds, err := ioutil.ReadDir(src);
	if err != nil {
		return err
	}

	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			err = CopyDirectory(srcfp, dstfp);
			if err != nil {
				return err
			}
		} else {
			err = CopyFile(srcfp, dstfp);
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func Exists(dst string) bool {
	_, err := os.Stat(dst)
	return !os.IsNotExist(err)
}

func WriteToFile(dst, value string) error {
	file, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(value)
	writer.Flush()

	return nil
}
