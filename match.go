package main

import (
// "os"
// "path/filepath"
)

// // TODO: Use this for path matching
// func ExpandUser(path string) string {
// 	usr, err := user.Current()
// 	if err != nil {
// 		return path
// 	}
// 	home := usr.HomeDir
// 	if path == "~" {
// 		return home
// 	} else if strings.HasPrefix(path, "~/") {
// 		return filepath.Join(home, path[2:])
// 	}
// }

func MatchCurrentPath(path string, locations *[]Location) *Location {
	for _, l := range *locations {
		if l.CurrentPathRegex != nil && !l.CurrentPathRegex.MatchString(path) {
			continue
		}

		return &l
	}
	return nil
}

func MatchChangePath(path string, locations *[]Location) *Location {
	for _, l := range *locations {
		if l.ChangePathRegex != nil && !l.ChangePathRegex.MatchString(path) {
			continue
		}

		return &l
	}
	return nil
}
