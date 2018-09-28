// sandbox.go
//
// Copyright © 2018 by Ollivier Robert <roberto@keltia.net>

package sandbox

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

const (
	// MyVersion is the SemVer version
	MyVersion = "0.9.2"

	// MyName is for when we need it
	MyName = "sandbox"
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
		return &Dir{}, errors.Wrapf(err, "create sandbox %s", sand)
	}
	fsand, err := filepath.Abs(sand)
	if err != nil {
		return &Dir{}, errors.Wrapf(err, "inconsistent %s", sand)
	}

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

	// Store absolute path
	fold, err := filepath.Abs(old)
	if err != nil {
		return errors.Wrap(err, "filepath.Abs")
	}
	s.old = fold

	// Go on
	return os.Chdir(s.folder)
}

func (s *Dir) Exit() error {
	return os.Chdir(s.old)
}

func (s *Dir) Cleanup() error {
	err := os.RemoveAll(s.folder)
	return errors.Wrapf(err, "cleanup failed for %s", s.folder)
}

func (s *Dir) Cwd() string {
	return s.folder
}

func (s *Dir) Run(f func() error) error {
	err := s.Enter()
	if err != nil {
		return errors.Wrap(err, "Run")
	}
	defer s.Exit()

	err = f()
	return errors.Wrap(err, "Run/f")
}

func Version() string {
	return fmt.Sprintf("%s/%s", MyName, MyVersion)
}
