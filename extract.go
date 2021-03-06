package main

import "os"

// extract contains the extraction from Readdir().
type extract struct {
	Files []string
	Dirs  []string

	HiddenFiles int
	HiddenDirs  int
}

// appendIfMatching considers the passed file and appends it to extract.Files or extract.Dirs
// if it is matching the hidden setting.
func (e *extract) appendIfMatching(file os.FileInfo, displayOnlyHidden bool) {
	name := file.Name()
	isHidden := name[:1] == "."

	if file.IsDir() {
		e.appendDir(name, displayOnlyHidden)

		if isHidden {
			e.HiddenDirs++
		}
	} else {
		e.appendFile(name, displayOnlyHidden)

		if isHidden && name != ".DS_Store" {
			e.HiddenFiles++
		}
	}
}

func (e *extract) appendDir(dir string, displayOnlyHidden bool) {
	if shouldAppendForDisplay(dir, displayOnlyHidden) {
		e.Dirs = append(e.Dirs, dir)
	}
}

func (e *extract) appendFile(fileName string, displayOnlyHidden bool) {
	if shouldAppendForDisplay(fileName, displayOnlyHidden) {
		e.Files = append(e.Files, fileName)
	}
}

func shouldAppendForDisplay(fileName string, displayOnlyHidden bool) bool {
	if displayOnlyHidden {
		if fileName[:1] == "." {
			return true
		}
	} else if fileName[:1] != "." {
		return true
	}

	return false
}
