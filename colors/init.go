package colors

import lg "github.com/charmbracelet/lipgloss"

const (
	PRIMARY   = lg.Color("#03F033") // green
	SECONDARY = lg.Color("#FF00FF") // magenta
	ACCENT    = lg.Color("#00A5FF") // blue
	DEFAULT   = lg.Color("#FFFFFF") // white

	SUCCESS = lg.Color("#00FF00") // green
	INFO    = lg.Color("#00FFFF") // cyan
	WARNING = lg.Color("#FFA500") // orange
	ERROR   = lg.Color("#FF0000") // red
)

func Render(s string, color lg.Color) string {
	return lg.NewStyle().Foreground(color).Render(s)
}
