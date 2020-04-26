package gradexpath

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupPaths(t *testing.T) {

	setTesting()

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
