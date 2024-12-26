package presentation

import "github.com/charmbracelet/lipgloss"

type styles struct {
	bold          lipgloss.Style
	faint         lipgloss.Style
	italic        lipgloss.Style
	underline     lipgloss.Style
	strikethrough lipgloss.Style
	red           lipgloss.Style
	green         lipgloss.Style
	yellow        lipgloss.Style
	blue          lipgloss.Style
	magenta       lipgloss.Style
	cyan          lipgloss.Style
	gray          lipgloss.Style
}

func makeStyles(r *lipgloss.Renderer) styles {
	return styles{
		bold:          r.NewStyle().SetString("bold").Bold(true),
		faint:         r.NewStyle().SetString("faint").Faint(true),
		italic:        r.NewStyle().SetString("italic").Italic(true),
		underline:     r.NewStyle().SetString("underline").Underline(true),
		strikethrough: r.NewStyle().SetString("strikethrough").Strikethrough(true),
		red:           r.NewStyle().SetString("red").Foreground(lipgloss.Color("#E88388")),
		green:         r.NewStyle().SetString("green").Foreground(lipgloss.Color("#A8CC8C")),
		yellow:        r.NewStyle().SetString("yellow").Foreground(lipgloss.Color("#DBAB79")),
		blue:          r.NewStyle().SetString("blue").Foreground(lipgloss.Color("#71BEF2")),
		magenta:       r.NewStyle().SetString("magenta").Foreground(lipgloss.Color("#D290E4")),
		cyan:          r.NewStyle().SetString("cyan").Foreground(lipgloss.Color("#66C2CD")),
		gray:          r.NewStyle().SetString("gray").Foreground(lipgloss.Color("#B9BFCA")),
	}
}
