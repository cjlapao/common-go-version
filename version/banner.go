package version

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	defaultBoxWidth = 80
	borderChar      = "*"
	paddingLeft     = 2
)

// BannerOptions configures banner generation behavior
type BannerOptions struct {
	// UseASCII determines whether to use ASCII art (true) or simple text (false) for the app name
	UseASCII bool
	// AutoWidth makes the banner auto-size to fit content instead of fixed width
	AutoWidth bool
	// FixedWidth sets a custom fixed width (ignored if AutoWidth is true, defaults to 80 if 0)
	FixedWidth int
	// FontStyle specifies the ASCII art font (only used if UseASCII is true)
	FontStyle FontStyle
	// ShowBorder determines whether to show the *** border box (defaults to true for fixed width, false for auto-width)
	ShowBorder *bool
}

// Banner generates a complete banner with simple text (no ASCII art), auto-width, and metadata
func Banner(appName string, info *Info) string {
	return BannerWithStyle(appName, info, FontStyleSlant)
}

// BannerWithStyle generates a complete banner with simple text (no ASCII art), auto-width, and metadata
// The style parameter is ignored when using default simple text mode
func BannerWithStyle(appName string, info *Info, style FontStyle) string {
	opts := BannerOptions{
		UseASCII:   false, // Default to simple text
		AutoWidth:  true,  // Default to auto-width
		FixedWidth: defaultBoxWidth,
		FontStyle:  style,
	}
	return BannerWithOptions(appName, info, opts)
}

// BannerWithOptions generates a complete banner with customizable options
func BannerWithOptions(appName string, info *Info, opts BannerOptions) string {
	// Set defaults
	// Default ShowBorder based on AutoWidth if not explicitly set
	showBorder := true
	if opts.ShowBorder != nil {
		showBorder = *opts.ShowBorder
	} else if opts.AutoWidth {
		showBorder = false // Default to no border for auto-width
	}

	var lines []string
	var titleLines []string

	// Generate title lines (ASCII art or simple text)
	if opts.UseASCII {
		switch {
		case opts.AutoWidth:
			// Generate without width constraint to get natural size
			titleLines = GenerateASCIIArtWithStyle(appName, 0, opts.FontStyle)
		case opts.FixedWidth > 0:
			available := opts.FixedWidth - 4 - paddingLeft*2
			if available < 0 {
				available = 0
			}
			// Generate with width constraint
			titleLines = GenerateASCIIArtWithStyle(appName, available, opts.FontStyle)
		default:
			// No explicit width - let ASCII art determine width
			titleLines = GenerateASCIIArtWithStyle(appName, 0, opts.FontStyle)
		}
	} else {
		// Simple text title - split by newlines if multi-line
		titleLines = splitManualTitleLines(appName)
	}

	if len(titleLines) == 0 {
		titleLines = []string{""}
	}

	// Calculate the longest line in title for centering (trim trailing spaces)
	maxTitleWidth := longestLineWidth(titleLines)
	contentWidth := maxContentWidth(titleLines, info)

	// Calculate box width
	var boxWidth int
	switch {
	case opts.AutoWidth:
		boxWidth = calculateAutoWidth(titleLines, info)
	case opts.FixedWidth > 0:
		boxWidth = opts.FixedWidth
	default:
		boxWidth = contentWidth + 4
		if boxWidth < defaultBoxWidth {
			boxWidth = defaultBoxWidth
		}
	}

	// Top border
	if showBorder {
		lines = append(lines, strings.Repeat(borderChar, boxWidth))
		// Empty line after top border
		lines = append(lines, formatBoxLineWithWidth("", boxWidth))
	}

		// Add title lines (centered within box when border is shown)
		for _, line := range titleLines {
			if showBorder {
				centered := centerText(line, boxWidth-4)
				lines = append(lines, formatBoxLineWithWidth(centered, boxWidth))
			} else {
				lines = append(lines, line)
			}
		}

	// Empty line after title
	if showBorder {
		lines = append(lines, formatBoxLineWithWidth("", boxWidth))
	} else {
		lines = append(lines, "")
	}

	// Version line (centered based on longest title line)
	versionLine := formatVersionLine(info)
	if showBorder {
		lines = append(lines, formatBoxLineWithWidth(centerText(versionLine, boxWidth-4), boxWidth))
	} else {
		// Center based on the longest title line
		lines = append(lines, centerText(versionLine, maxTitleWidth))
	}

	// Empty line after version
	if showBorder {
		lines = append(lines, formatBoxLineWithWidth("", boxWidth))
	} else {
		lines = append(lines, "")
	}

	// Metadata lines
	if info.Author != "" {
		line := fmt.Sprintf("Author:    %s", info.Author)
		if showBorder {
			lines = append(lines, formatBoxLineWithWidth(line, boxWidth))
		} else {
			lines = append(lines, line)
		}
	}
	if info.Company != "" {
		line := fmt.Sprintf("Company:   %s", info.Company)
		if showBorder {
			lines = append(lines, formatBoxLineWithWidth(line, boxWidth))
		} else {
			lines = append(lines, line)
		}
	}
	if info.Copyright != "" {
		line := fmt.Sprintf("Copyright: %s", info.Copyright)
		if showBorder {
			lines = append(lines, formatBoxLineWithWidth(line, boxWidth))
		} else {
			lines = append(lines, line)
		}
	}
	if info.Repo != "" {
		line := fmt.Sprintf("Repo:      %s", info.Repo)
		if showBorder {
			lines = append(lines, formatBoxLineWithWidth(line, boxWidth))
		} else {
			lines = append(lines, line)
		}
	}

	// Empty line before bottom border (only if we have metadata)
	if info.Author != "" || info.Company != "" || info.Copyright != "" || info.Repo != "" {
		if showBorder {
			lines = append(lines, formatBoxLineWithWidth("", boxWidth))
		} else {
			lines = append(lines, "")
		}
	}

	// Bottom border
	if showBorder {
		lines = append(lines, strings.Repeat(borderChar, boxWidth))
	}

	return strings.Join(lines, "\n")
}

// splitManualTitleLines breaks manual multi-line titles into individual lines while preserving indentation.
func splitManualTitleLines(title string) []string {
	if title == "" {
		return []string{""}
	}

	normalized := strings.ReplaceAll(title, "\r\n", "\n")
	normalized = strings.ReplaceAll(normalized, "\r", "")

	lines := strings.Split(normalized, "\n")

	// Remove leading and trailing empty lines introduced by raw string literals.
	start := 0
	end := len(lines)

	for start < end && strings.TrimSpace(lines[start]) == "" {
		start++
	}
	for end > start && strings.TrimSpace(lines[end-1]) == "" {
		end--
	}

	if start >= end {
		return []string{""}
	}

	return lines[start:end]
}

// longestLineWidth returns the maximum visible width among provided lines.
func longestLineWidth(lines []string) int {
	maxWidth := 0
	for _, line := range lines {
		width := visibleWidth(line)
		if width > maxWidth {
			maxWidth = width
		}
	}
	return maxWidth
}

// visibleWidth calculates the width of a line ignoring trailing whitespace that doesn't affect rendering.
func visibleWidth(line string) int {
	trimmed := strings.TrimRight(line, " \t")
	return utf8.RuneCountInString(trimmed)
}

// maxContentWidth calculates the maximum width among title, version, and metadata lines.
func maxContentWidth(titleLines []string, info *Info) int {
	maxWidth := longestLineWidth(titleLines)

	versionLine := formatVersionLine(info)
	if width := utf8.RuneCountInString(versionLine); width > maxWidth {
		maxWidth = width
	}

	if info.Author != "" {
		if width := utf8.RuneCountInString(fmt.Sprintf("Author:    %s", info.Author)); width > maxWidth {
			maxWidth = width
		}
	}
	if info.Company != "" {
		if width := utf8.RuneCountInString(fmt.Sprintf("Company:   %s", info.Company)); width > maxWidth {
			maxWidth = width
		}
	}
	if info.Copyright != "" {
		if width := utf8.RuneCountInString(fmt.Sprintf("Copyright: %s", info.Copyright)); width > maxWidth {
			maxWidth = width
		}
	}
	if info.Repo != "" {
		if width := utf8.RuneCountInString(fmt.Sprintf("Repo:      %s", info.Repo)); width > maxWidth {
			maxWidth = width
		}
	}

	return maxWidth
}

// calculateAutoWidth determines the optimal width based on content
func calculateAutoWidth(titleLines []string, info *Info) int {
	// Add padding for borders and margins
	return maxContentWidth(titleLines, info) + 4 // 2 for borders, 2 for padding
}

// formatVersionLine creates a formatted version string
func formatVersionLine(info *Info) string {
	var parts []string

	// Base version
	versionStr := fmt.Sprintf("v%d.%d.%d", info.Major, info.Minor, info.Patch)
	parts = append(parts, versionStr)

	// Add hash if present
	if info.Hash != "" {
		parts = append(parts, info.Hash)
	}

	// Add suffix if present
	if info.Suffix != "" {
		suffixUpper := strings.ToUpper(info.Suffix)
		parts = append(parts, fmt.Sprintf("[%s]", suffixUpper))
	}

	return strings.Join(parts, " ")
}

// formatBoxLineWithWidth formats a line to fit within the box with borders and padding using specified width
func formatBoxLineWithWidth(content string, boxWidth int) string {
	// Calculate available width (excluding borders and padding)
	availableWidth := boxWidth - 4 // 2 for borders (* *) and 2 for padding (space on each side)
	if availableWidth < 0 {
		availableWidth = 0
	}

	// Pad or truncate content to fit while respecting rune boundaries
	runes := []rune(content)
	if len(runes) > availableWidth {
		runes = runes[:availableWidth]
	}

	trimmed := string(runes)
	currentWidth := utf8.RuneCountInString(trimmed)
	if currentWidth < availableWidth {
		trimmed += strings.Repeat(" ", availableWidth-currentWidth)
	}

	// Add border and padding
	return fmt.Sprintf("%s %s %s", borderChar, trimmed, borderChar)
}

// centerText centers text within a given width
func centerText(text string, width int) string {
	textLen := utf8.RuneCountInString(text)
	if textLen >= width || width <= 0 {
		return text
	}

	totalPadding := width - textLen
	leftPadding := totalPadding / 2
	rightPadding := totalPadding - leftPadding

	return strings.Repeat(" ", leftPadding) + text + strings.Repeat(" ", rightPadding)
}

// Print prints the banner to stdout using default slant font
func Print(appName string, info *Info) {
	fmt.Println(Banner(appName, info))
}

// PrintWithStyle prints the banner to stdout using specified font style
func PrintWithStyle(appName string, info *Info, style FontStyle) {
	fmt.Println(BannerWithStyle(appName, info, style))
}

// PrintWithOptions prints the banner to stdout using custom options
func PrintWithOptions(appName string, info *Info, opts BannerOptions) {
	fmt.Println(BannerWithOptions(appName, info, opts))
}

// QuickPrint is a convenience function that parses version string and prints banner using default slant font
func QuickPrint(appName, versionStr, author, company, copyright, repo string) error {
	return QuickPrintWithStyle(appName, versionStr, author, company, copyright, repo, FontStyleSlant)
}

// QuickPrintWithStyle is a convenience function that parses version string and prints banner using specified font style
func QuickPrintWithStyle(appName, versionStr, author, company, copyright, repo string, style FontStyle) error {
	info, err := Parse(versionStr)
	if err != nil {
		return err
	}

	info.Author = author
	info.Company = company
	info.Copyright = copyright
	info.Repo = repo

	PrintWithStyle(appName, info, style)
	return nil
}

// QuickPrintWithOptions is a convenience function that parses version string and prints banner using custom options
func QuickPrintWithOptions(appName, versionStr, author, company, copyright, repo string, opts BannerOptions) error {
	info, err := Parse(versionStr)
	if err != nil {
		return err
	}

	info.Author = author
	info.Company = company
	info.Copyright = copyright
	info.Repo = repo

	PrintWithOptions(appName, info, opts)
	return nil
}
