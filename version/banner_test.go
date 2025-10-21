package version

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func TestBannerCentersVersionForMultilineTitle(t *testing.T) {
	appName := `
MONEYGROW AI
BACKEND SERVER
PLATFORM
`

	info := &Info{
		Major: 0,
		Minor: 2,
		Patch: 0,
	}

	opts := BannerOptions{
		UseASCII:  false,
		AutoWidth: true,
	}

	output := BannerWithOptions(appName, info, opts)

	lines := strings.Split(output, "\n")
	titleLines := splitManualTitleLines(appName)

	versionIdx := len(titleLines) + 1
	if versionIdx >= len(lines) {
		t.Fatalf("version line missing in banner output: %q", output)
	}

	versionLine := lines[versionIdx]
	if strings.TrimSpace(versionLine) != "v0.2.0" {
		t.Fatalf("expected version line to contain v0.2.0, got %q", versionLine)
	}

	expectedWidth := longestLineWidth(titleLines)
	actualWidth := utf8.RuneCountInString(versionLine)
	if actualWidth != expectedWidth {
		t.Fatalf("expected centered version line width %d, got %d (line %q)", expectedWidth, actualWidth, versionLine)
	}

	leftPadding := len(versionLine) - len(strings.TrimLeft(versionLine, " "))
	rightPadding := len(versionLine) - len(strings.TrimRight(versionLine, " "))
	diff := leftPadding - rightPadding
	if diff < 0 {
		diff = -diff
	}

	if diff > 1 {
		t.Fatalf("version line not centered: left padding %d, right padding %d (line %q)", leftPadding, rightPadding, versionLine)
	}
}

func TestBannerExpandsBorderForWideContent(t *testing.T) {
	appName := strings.Repeat("A", 100)
	info := &Info{
		Major: 0,
		Minor: 2,
		Patch: 0,
	}

	showBorder := true
	opts := BannerOptions{
		UseASCII:   false,
		AutoWidth:  false,
		ShowBorder: &showBorder,
	}

	output := BannerWithOptions(appName, info, opts)
	lines := strings.Split(output, "\n")
	if len(lines) == 0 {
		t.Fatalf("expected banner output, got empty string")
	}

	topBorder := lines[0]
	expectedWidth := utf8.RuneCountInString(appName) + 4
	if len(topBorder) != expectedWidth {
		t.Fatalf("expected top border width %d, got %d (line %q)", expectedWidth, len(topBorder), topBorder)
	}

	versionLine := ""
	for _, line := range lines {
		if strings.Contains(line, "v0.2.0") {
			versionLine = line
			break
		}
	}
	if versionLine == "" {
		t.Fatalf("version line not found in banner output")
	}

	innerWidth := expectedWidth - 4
	if utf8.RuneCountInString(strings.Trim(versionLine, "* ")) > innerWidth {
		t.Fatalf("version line exceeds available width: %q", versionLine)
	}
}

func TestBannerKeepsMultibyteCharactersIntact(t *testing.T) {
	appName := strings.Join([]string{
		"██████╗  █████╗ ███╗   ██╗",
		"██╔══██╗██╔══██╗████╗  ██║",
	}, "\n")

	info := &Info{Major: 1, Minor: 0, Patch: 0}

	opts := BannerOptions{UseASCII: false, AutoWidth: false}

	output := BannerWithOptions(appName, info, opts)
	lines := strings.Split(output, "\n")

	for _, expected := range strings.Split(appName, "\n") {
		found := false
		for _, line := range lines {
			if strings.Contains(line, expected) {
				if !utf8.ValidString(line) {
					t.Fatalf("line contains invalid UTF-8: %q", line)
				}
				if strings.Contains(line, "�") {
					t.Fatalf("line contains replacement character: %q", line)
				}
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("expected to find art line %q in banner output", expected)
		}
	}
}
