// sandbox.go
//
// Copyright © 2018 by Ollivier Robert <roberto@keltia.net>

package sandbox

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Dir struct {
	tag    string
	folder string
	old    string
}

func New(tag string) (*Dir, error) {
	// Extract in safe location
	sand, err := ioutil.TempDir("", tag)
	if err != nil {
		return &Dir{}, fmt.Errorf("unable to create sandbox %s: %v", sand, err)
	}
	fsand, err := filepath.Abs(sand)
	dir := &Dir{
		tag:    tag,
		folder: fsand,
	}
	return dir, nil
}

func (s *Dir) Enter() error {
	// Save
	old, err := os.Getwd()
	if err != nil {
		return err
	}
	s.old = old

	// Go on
	return os.Chdir(s.folder)
}

func (s *Dir) Exit() error {
	return os.Chdir(s.old)
}

func (s *Dir) Cleanup() error {
	err := os.RemoveAll(s.folder)
	return fmt.Errorf("cleanup failed for %s: %v", s.folder, err)
}

func (s *Dir) Cwd() string {
	return s.folder
}
