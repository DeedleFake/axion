package theme

import "github.com/charmbracelet/lipgloss"

func c(str string) lipgloss.Color {
	return lipgloss.Color(str)
}

func ac(light, dark string) lipgloss.AdaptiveColor {
	return lipgloss.AdaptiveColor{
		Light: light,
		Dark:  dark,
	}
}

type Theme struct {
	editor lipgloss.Style
}

// The default theme is based on ra-silver and ra-dark for VSCode:
// https://github.com/rahmanyerli/ra-silver
var defaultTheme = Theme{
	editor: lipgloss.NewStyle().
		Background(ac("#D4D9DE", "#202020")).
		Foreground(ac("#161A1D", "#E0E0E0")).
		BorderBackground(ac("#D4D9DE", "#202020")).
		BorderForeground(ac("#161A1D", "#E0E0E0")),
}

func Default() Theme {
	return defaultTheme
}

func (t Theme) Editor() lipgloss.Style {
	return t.editor.Copy()
}

// ra-silver:
// "editor.background": "#D4D9DE",
// "editor.foreground": "#161A1D",
// "editor.lineHighlightBorder": "#00000011",
// "editorLineNumber.foreground": "#7E8C9B",
// "editorLineNumber.activeForeground": "#434D56",
// "editorIndentGuide.background": "#00000010",
// "editorIndentGuide.activeBackground": "#00000020",
// "activityBar.background": "#708090",
// "activityBar.foreground": "#161A1D",
// "activityBarBadge.background": "#3F51B5",
// "activityBar.dropBackground": "#3F51B5",
// "sideBar.background": "#C6CCD2",
// "sideBar.foreground": "#21262B",
// "sideBarSectionHeader.background": "#708090",
// "sideBarSectionHeader.foreground": "#21262B",
// "statusBar.background": "#708090",
// "statusBar.debuggingBackground": "#3F51B5",
// "list.activeSelectionBackground": "#9BA6B1",
// "list.activeSelectionForeground": "#F1F2F4",
// "list.inactiveSelectionBackground": "#B7BFC7",
// "list.inactiveSelectionForeground": "#21262B",
// "list.hoverBackground": "#B7BFC7",
// "editorGroupHeader.tabsBackground": "#708090",
// "editorGroupHeader.foreground": "#21262B",
// "tab.activeBackground": "#D4D9DE",
// "tab.inactiveBackground": "#708090",
// "tab.activeForeground": "#21262B",
// "tab.inactiveForeground": "#F1F2F4",
// "tab.border": "#00000010"

// ra-dark:
// "editor.background": "#202020",
// "editor.foreground": "#E0E0E0",
// "editorLineNumber.foreground": "#606060",
// "editorLineNumber.activeForeground": "#909090",
// "editorIndentGuide.background": "#FFFFFF10",
// "editorIndentGuide.activeBackground": "#FFFFFF20",
// "editor.lineHighlightBackground": "#FFFFFF08",
// "editor.selectionBackground": "#FFFFFF20",
// "editor.selectionHighlightBackground": "#FFFFFF30",
// "editor.wordHighlightBackground": "#FFFFFF30",
// "editor.wordHighlightStrongBackground": "#FFFFFF40",
// "editor.findMatchBackground": "#FFFFFF10",
// "editor.findMatchHighlightBackground": "#FFFFFF20",
// "editor.findRangeHighlightBackground": "#FFFFFF30",
// "editor.hoverHighlightBackground": "#FFFFFF40",
// "editor.rangeHighlightBackground": "#FFFFFF40",
// "editorBracketMatch.background": "#FFFFFF10",
// "editorWidget.background": "#202020",
// "activityBar.background": "#202020",
// "activityBar.border": "#FFFFFF10",
// "activityBarBadge.background": "#3366CC",
// "sideBar.background": "#202020",
// "sideBar.border": "#FFFFFF10",
// "sideBarSectionHeader.background": "#242424",
// "statusBar.background": "#202020",
// "statusBar.debuggingBackground": "#3366CC",
// "statusBar.border": "#FFFFFF10",
// "list.activeSelectionBackground": "#303030",
// "list.inactiveSelectionBackground": "#202020",
// "list.hoverBackground": "#282828",
// "editorGroupHeader.tabsBackground": "#202020",
// "editorGroupHeader.tabsBorder": "#FFFFFF10",
// "tab.activeBackground": "#282828",
// "tab.activeBorder": "#FF9999",
// "tab.inactiveBackground": "#202020",
// "tab.border": "#FFFFFF10"
