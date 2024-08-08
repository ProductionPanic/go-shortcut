package main

import (
	"fmt"
	lg "github.com/charmbracelet/lipgloss"
	"os"
	"shortcut/colors"
	"shortcut/store"
)

func main() {
	store.Install()

	args := os.Args[1:]

	if len(args) == 0 {
		printHelp()
		return
	}

	switch args[0] {
	case "add":
		if len(args) != 3 {
			printHelp()
			return
		}
		store.Add(args[1], args[2])
		fmt.Println(colors.Render("Added alias "+args[1], colors.SUCCESS))
	case "remove":
		if len(args) != 2 {
			printHelp()
			return
		}
		store.Remove(args[1])
		fmt.Println(colors.Render("Removed alias "+args[1], colors.SUCCESS))
	default:
		printHelp()
	}
}

func printHelp() {
	s := []string{}
	s = append(s, colors.Render("Usage:", colors.PRIMARY))
	s = append(s, colors.Render("  shortcut add ", colors.DEFAULT)+colors.Render("<alias> <command>", colors.ACCENT))
	s = append(s, colors.Render("  shortcut remove ", colors.DEFAULT)+colors.Render("<alias>", colors.ACCENT))
	o := lg.JoinVertical(lg.Left, s...)
	fmt.Println(o)
}
