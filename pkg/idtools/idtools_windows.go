// +build windows

package idtools

import (
	"os"

	"github.com/docker/docker/pkg/system/filesys"
)

// Platforms such as Windows do not support the UID/GID concept. So make this
// just a wrapper around filesys.MkdirAll.
func mkdirAs(path string, mode os.FileMode, ownerUID, ownerGID int, mkAll, chownExisting bool) error {
	if err := filesys.MkdirAll(path, mode); err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}
