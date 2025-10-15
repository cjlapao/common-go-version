package version

import (
	"fmt"
	"strings"
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
	if opts.FixedWidth == 0 {
		opts.FixedWidth = defaultBoxWidth
	}

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
		if opts.AutoWidth {
			// Generate without width constraint to get natural size
			titleLines = GenerateASCIIArtWithStyle(appName, 0, opts.FontStyle)
		} else {
			// Generate with width constraint
			titleLines = GenerateASCIIArtWithStyle(appName, opts.FixedWidth-4-paddingLeft*2, opts.FontStyle)
		}
	} else {
		// Simple text title - split by newlines if multi-line
		titleLines = strings.Split(appName, "\n")
	}

	// Calculate the longest line in title for centering (trim trailing spaces)
	maxTitleWidth := 0
	for _, line := range titleLines {
		trimmedLen := len(strings.TrimRight(line, " "))
		if trimmedLen > maxTitleWidth {
			maxTitleWidth = trimmedLen
		}
	}

	// Calculate box width
	boxWidth := opts.FixedWidth
	if opts.AutoWidth {
		boxWidth = calculateAutoWidth(appName, titleLines, info)
	}

	// Top border
	if showBorder {
		lines = append(lines, strings.Repeat(borderChar, boxWidth))
		// Empty line after top border
		lines = append(lines, formatBoxLineWithWidth("", boxWidth))
	}

	// Add title lines (centered based on longest title line)
	for _, line := range titleLines {
		if showBorder {
			lines = append(lines, formatBoxLineWithWidth(line, boxWidth))
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

// calculateAutoWidth determines the optimal width based on content
func calculateAutoWidth(appName string, titleLines []string, info *Info) int {
	maxWidth := len(appName) + 4 // minimum for simple text

	// Check ASCII art width
	for _, line := range titleLines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	// Check version line width
	versionLine := formatVersionLine(info)
	if len(versionLine) > maxWidth {
		maxWidth = len(versionLine)
	}

	// Check metadata widths
	if info.Author != "" {
		w := len(fmt.Sprintf("Author:    %s", info.Author))
		if w > maxWidth {
			maxWidth = w
		}
	}
	if info.Company != "" {
		w := len(fmt.Sprintf("Company:   %s", info.Company))
		if w > maxWidth {
			maxWidth = w
		}
	}
	if info.Copyright != "" {
		w := len(fmt.Sprintf("Copyright: %s", info.Copyright))
		if w > maxWidth {
			maxWidth = w
		}
	}
	if info.Repo != "" {
		w := len(fmt.Sprintf("Repo:      %s", info.Repo))
		if w > maxWidth {
			maxWidth = w
		}
	}

	// Add padding for borders and margins
	return maxWidth + 4 // 2 for borders, 2 for padding
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

	// Pad or truncate content to fit
	contentLen := len(content)
	if contentLen > availableWidth {
		content = content[:availableWidth]
	} else if contentLen < availableWidth {
		content = content + strings.Repeat(" ", availableWidth-contentLen)
	}

	// Add border and padding
	return fmt.Sprintf("%s %s %s", borderChar, content, borderChar)
}

// centerText centers text within a given width
func centerText(text string, width int) string {
	textLen := len(text)
	if textLen >= width {
		return text
	}

	leftPadding := (width - textLen) / 2
	rightPadding := width - textLen - leftPadding

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
