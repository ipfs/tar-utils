// +build !linux
// +build !freebsd
// +build !netbsd
// +build !openbsd
// +build !dragonfly

package tar

import (
	"os"
	"time"
)

func updateMtime(path string, mtime time.Time) error {
	return os.Chtimes(path, mtime, mtime)
}
