package ls

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNames(t *testing.T) {
	path := "ls.go"
	info, err := os.Stat(path)
	assert.NoError(t, err)
	uid, gid := getOwnership(info)
	user, group, err := getNames(uid, gid)
	assert.NoError(t, err)
	assert.Equal(t, user, "mostafa")
	assert.Equal(t, group, "mostafa")
}

func TestGetOwnerShip(t *testing.T) {
	path := "ls.go"
	info, err := os.Stat(path)
	assert.NoError(t, err)
	uid, gid := getOwnership(info)
	t.Log(uid)
	assert.Equal(t, int(uid), 1000)
	assert.Equal(t, int(gid), 1000)
}
