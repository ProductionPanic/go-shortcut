package store

import (
	"path"
	"shortcut/store/fs"
)

var fileName string = ".shortcuts"

func Install() {
	fs.EnsureFile(path.Join(fs.Home(), fileName))
	appendStr := "source " + path.Join(fs.Home(), fileName) + "\n"

	rcFiles := []string{
		path.Join(fs.Home(), ".bashrc"),
		path.Join(fs.Home(), ".zshrc"),
		path.Join(fs.Home(), ".bash_profile"),
		path.Join(fs.Home(), ".profile"),
	}

	for _, rcFile := range rcFiles {
		if fs.FileExists(rcFile) && !fs.FileHasString(rcFile, appendStr) {
			fs.AppendToFile(rcFile, appendStr)
		}
	}
}

func Add(alias string, command string) {
	fs.AppendToFile(path.Join(fs.Home(), fileName), "alias "+alias+"="+command+"\n")
}

func Remove(alias string) {
	fs.RemoveFromFile(path.Join(fs.Home(), fileName), "alias "+alias+"=.*\n")
}
