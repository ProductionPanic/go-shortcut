package store

import (
	lg "github.com/charmbracelet/lipgloss"
	"os"
	"os/exec"
	"path"
	"regexp"
	"shortcut/colors"
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

func List() {
	file, err := fs.ReadFile(path.Join(fs.Home(), fileName))
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile("alias\\s(.*?)=(.*)\n")
	matches := re.FindAllStringSubmatch(string(file), -1)
	for _, match := range matches {
		lg.JoinHorizontal(lg.Left,
			lg.NewStyle().Foreground(colors.PRIMARY).Render(match[1]),
			lg.NewStyle().Foreground(colors.DEFAULT).Render(" -> "),
			lg.NewStyle().Foreground(colors.ACCENT).Render(match[2]),
		)
	}

}

func Run(alias string) {
	file, err := fs.ReadFile(path.Join(fs.Home(), fileName))
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile("alias\\s" + alias + "=(.*)\n")
	match := re.FindStringSubmatch(string(file))
	if len(match) == 0 {
		panic("alias not found")
	}
	runCommand(match[1])
}

func runCommand(command string) {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
