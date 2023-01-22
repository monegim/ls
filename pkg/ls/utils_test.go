package ls

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNames(t *testing.T) {
	// path := "ls.go"
	// // info, err := os.Stat(path)
	// f, err := os.Open(path)
	// assert.NoError(t, err)
	// defer f.Close()
	// file, err := f.ReadDir(1)
	// info := file[0].Info()
	// assert.NoError(t, err)
	// stat := info.Sys().(*syscall.Stat_t)
	// _ = stat
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
