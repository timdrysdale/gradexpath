package gradexpath

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSetupPaths(t *testing.T) {

	SetTesting()

	root := Root()

	if root != "./tmp-delete-me" {
		t.Errorf("test root set up wrong %s", root)
	}

	// don't use GetRoot() here
	// JUST in case we kill a whole working installation
	os.RemoveAll("./tmp-delete-me")

	err := SetupGradexPaths()
	assert.NoError(t, err)

	SetupExamPaths("sample")
	assert.NoError(t, err)
}

//check we can move files without adjusting the modification time
func TestFileMod(t *testing.T) {

	d1 := []byte("Gradex Testing\n")
	basepath := filepath.Join(Root(), "tmp")
	err := EnsureDir(basepath)
	assert.NoError(t, err)
	testPath := filepath.Join(basepath, "test.txt")
	err = ioutil.WriteFile(testPath, d1, 0755)
	assert.NoError(t, err)
	err = os.Chmod(testPath, 0755)
	assert.NoError(t, err)

	info, err := os.Stat(testPath)
	assert.NoError(t, err)

	time.Sleep(10 * time.Millisecond)

	assert.NotEqual(t, info.ModTime(), time.Now())

	newPath := filepath.Join(Root(), "tmp", "new.txt")
	err = os.Rename(testPath, newPath)
	infoNew, err := os.Stat(newPath)

	assert.NoError(t, err)
	assert.Equal(t, info.ModTime(), infoNew.ModTime())
}
