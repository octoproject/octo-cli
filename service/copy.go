package service

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func copy(src, dest string) error {
	i, err := os.Stat(src)
	if err != nil {
		return err
	}

	if i.IsDir() {
		return copyDir(src, dest)
	}
	return copyFile(src, dest)
}

func copyDir(src, dest string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	_, err = os.Stat(dest)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if exists(dest) {
		return fmt.Errorf("folder: %s already exists", dest)
	}

	err = os.MkdirAll(dest, srcInfo.Mode())
	if err != nil {
		return fmt.Errorf("error creating folder: %s with %s", dest, err.Error())
	}

	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		err := copy(filepath.Join(src, entry.Name()),
			filepath.Join(dest, entry.Name()))
		if err != nil {
			return err
		}
	}
	return nil
}

func copyFile(src, dest string) error {
	si, err := os.Stat(src)
	if err != nil {
		return err
	}

	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()

	err = os.Chmod(dest, si.Mode())
	if err != nil {
		return err
	}

	s, err := os.Open(src)
	if err != nil {
		return err
	}
	defer s.Close()

	_, err = io.Copy(f, s)
	if err != nil {
		return err
	}
	return nil
}

func exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}
