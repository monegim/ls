package ls

import (
	"strconv"
)

func getNames(uid, gid uint32) (string, string, error) {
	usr := strconv.FormatUint(uint64(uid), 10)
	group := strconv.FormatUint(uint64(gid), 10)
	_, _ = usr, group
	// if u, err := user.Lo
	return "", "", nil
}
