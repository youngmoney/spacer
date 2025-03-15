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

func MatchChangePath(name string, path string, locations *[]Location) *Location {
	for _, l := range *locations {
		if l.ChangePathRegex != nil && !l.ChangePathRegex.MatchString(path) {
			continue
		}
		if l.Name != name {
			continue
		}
		return &l
	}
	return nil
}

func MatchName(name string, locations *[]Location) *Location {
	for _, l := range *locations {
		if l.Name != name {
			continue
		}
		return &l
	}
	return nil
}

func MatchCreatorName(name string, creators *[]Creator) *Creator {
	for _, c := range *creators {
		if c.Name != name {
			continue
		}
		return &c
	}
	return nil
}

func MatchLayoutName(name string, layouts *[]Layout) *Layout {
	for _, l := range *layouts {
		if l.Name != name {
			continue
		}
		return &l
	}
	return nil
}
