package ls

import (
	"os"
	"strconv"
	"syscall"
)

func getNames(uid, gid uint32) (string, string, error) {
	usr := strconv.FormatUint(uint64(uid), 10)
	group := strconv.FormatUint(uint64(gid), 10)
	_, _ = usr, group
	// if u, err := user.Lo
	return "", "", nil
}

func getOwnership(info os.FileInfo) (uid, gid uint32) {
	stat := info.Sys().(*syscall.Stat_t)
	return stat.Uid, stat.Gid
}
