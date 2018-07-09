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

	tmpd := os.Getenv("TMPDIR")
	assert.True(t, strings.HasPrefix(snd.folder, tmpd))

	t.Logf("current sandbox is %s", snd.folder)
}

func TestNew2(t *testing.T) {
	snd, err := New("")
	assert.NoError(t, err)
	assert.NotNil(t, snd)
}

func TestDir_Cwd(t *testing.T) {
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

}
