package vcs

import (
	"fmt"
	"runtime/debug"
)

func Version() string {
	var (
		revision string
		modified bool
	)

	bi, ok := debug.ReadBuildInfo()
	if ok {
		for _, s := range bi.Settings {
			switch s.Key {
			case "vcs.revision":
				revision = s.Value
			case "vcs.modified":
				modified = true
			}
		}
	}

	if modified {
		return fmt.Sprintf("%s-dirty", revision)
	}

	return revision
}