package ls

import (
	"os"
	"os/user"
	"strconv"
	"syscall"
)

func getNames(uid, gid uint32) (string, string, error) {
	usr := strconv.FormatUint(uint64(uid), 10)
	group := strconv.FormatUint(uint64(gid), 10)
	u, err := user.LookupId(usr)
	if err != nil {
		return "", "", err
	}
	usr = u.Username
	g, err := user.LookupGroupId(group)
	if err != nil {
		return "", "", err
	}
	group = g.Name
	return usr, group, nil
}

func getOwnership(info os.FileInfo) (uid, gid uint32) {
	stat := info.Sys().(*syscall.Stat_t)
	return stat.Uid, stat.Gid
}
