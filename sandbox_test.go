package sandbox

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	snd, err := New("test")
	assert.NoError(t, err)
	assert.NotNil(t, snd)
	assert.Equal(t, "test", snd.tag)

	defer snd.Cleanup()

	tmpd := os.Getenv("TMPDIR")
	assert.True(t, strings.HasPrefix(snd.folder, tmpd))

	t.Logf("current sandbox is %s", snd.folder)
}

func TestNew2(t *testing.T) {
	snd, err := New("")
	assert.NoError(t, err)
	assert.NotNil(t, snd)
	snd.Cleanup()
}

func TestDir_Cwd(t *testing.T) {
	snd, err := New("test")
	assert.NoError(t, err)
	assert.NotNil(t, snd)

	defer snd.Cleanup()

	snd.Enter()
	spwd := snd.Cwd()

	pwd, err := os.Getwd()
	assert.NoError(t, err)
	fpwd, err := filepath.Abs(pwd)
	assert.NoError(t, err)

	assert.Equal(t, fpwd, spwd)
}

func TestDir_Enter(t *testing.T) {
	snd, err := New("test")
	assert.NoError(t, err)
	assert.NotNil(t, snd)

	// Get where we are
	opwd, err := os.Getwd()
	assert.NoError(t, err)
	fopwd, err := filepath.Abs(opwd)
	assert.NoError(t, err)

	snd.Enter()

	// Get where we are after Enter()
	apwd, err := os.Getwd()
	assert.NoError(t, err)
	fapwd, err := filepath.Abs(apwd)
	assert.NoError(t, err)

	cwd, err := filepath.Abs(snd.Cwd())
	assert.NoError(t, err)
	assert.Equal(t, fapwd, cwd)

	assert.Equal(t, opwd, snd.old)
	assert.Equal(t, fopwd, snd.old)
}

func TestDir_Exit(t *testing.T) {
	snd, err := New("test")
	assert.NoError(t, err)
	assert.NotNil(t, snd)

	// Get where we are
	opwd, err := os.Getwd()
	assert.NoError(t, err)
	fopwd, err := filepath.Abs(opwd)
	assert.NoError(t, err)

	snd.Enter()

	// Get where we are after Enter()
	apwd, err := os.Getwd()
	assert.NoError(t, err)
	fapwd, err := filepath.Abs(apwd)
	assert.NoError(t, err)

	assert.Equal(t, snd.old, fapwd)

	err = snd.Exit()
	assert.NoError(t, err)

	npwd, err := os.Getwd()
	assert.NoError(t, err)
	fnpwd, err := filepath.Abs(npwd)
	assert.NoError(t, err)

	assert.Equal(t, fopwd, fnpwd)
}
