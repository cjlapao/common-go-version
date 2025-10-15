// Package version provides utilities for parsing and formatting software version information.
package version

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Info holds version information
type Info struct {
	Version   string // Full version string (e.g., "0.2.0:EF06A10-dev")
	Major     int
	Minor     int
	Patch     int
	Hash      string // Git commit hash
	Suffix    string // alpha, beta, dev, or empty for release
	Author    string
	Company   string
	Copyright string
	Repo      string
}

// VersionJSON represents version information in JSON format
type VersionJSON struct {
	Version   string  `json:"version"`
	Hash      *string `json:"hash,omitempty"`
	BuildType *string `json:"build_type,omitempty"`
}

var versionRegex = regexp.MustCompile(`^(\d+)\.(\d+)\.(\d+)(?::([A-Fa-f0-9]+))?(?:-(alpha|beta|dev|rc\d+|canary))?$`)

// Parse parses a version string in various formats
// Supported formats:
//   - "0.0.1"              -> version only
//   - "0.0.1-canary"       -> version with build_type
//   - "0.0.1:4f00"         -> version with hash
//   - "0.0.1:4f00-beta"    -> version with hash and build_type
//
// Invalid formats (will return error):
//   - "4:ff00-beta"        -> missing proper version format
//   - ":abc123"            -> missing version
//   - "1.2"                -> incomplete version (needs X.Y.Z)
func Parse(versionStr string) (*Info, error) {
	versionStr = strings.TrimSpace(versionStr)
	versionStr = strings.TrimPrefix(versionStr, "v")
	versionStr = strings.TrimPrefix(versionStr, "V")

	if versionStr == "" {
		return nil, fmt.Errorf("version string cannot be empty")
	}
	// Validate that it starts with a proper version number
	if !strings.Contains(versionStr, ".") || !regexp.MustCompile(`^\d+\.\d+\.\d+`).MatchString(versionStr) {
		return nil, fmt.Errorf("invalid version format: %s (version must start with X.Y.Z format)", versionStr)
	}

	matches := versionRegex.FindStringSubmatch(versionStr)
	if matches == nil {
		return nil, fmt.Errorf("invalid version format: %s (expected: X.Y.Z[:HASH][-suffix])", versionStr)
	}

	info := &Info{
		Version: versionStr,
	}

	// Parse major, minor, patch
	info.Major, _ = strconv.Atoi(matches[1])
	info.Minor, _ = strconv.Atoi(matches[2])
	info.Patch, _ = strconv.Atoi(matches[3])

	// Optional hash
	if matches[4] != "" {
		info.Hash = strings.ToUpper(matches[4])
	}

	// Optional suffix
	if matches[5] != "" {
		info.Suffix = matches[5]
	}

	return info, nil
}

// String returns the full version string
func (i *Info) String() string {
	return i.Version
}

// Short returns a short version string (e.g., "v0.2.0-dev")
func (i *Info) Short() string {
	ver := fmt.Sprintf("v%d.%d.%d", i.Major, i.Minor, i.Patch)
	if i.Suffix != "" {
		ver += "-" + i.Suffix
	}
	return ver
}

// IsRelease returns true if this is a release version (no suffix)
func (i *Info) IsRelease() bool {
	return i.Suffix == ""
}

// IsDev returns true if this is a dev version
func (i *Info) IsDev() bool {
	return i.Suffix == "dev"
}

// IsPreRelease returns true if this is alpha, beta, or rc
func (i *Info) IsPreRelease() bool {
	return strings.HasPrefix(i.Suffix, "alpha") ||
		strings.HasPrefix(i.Suffix, "beta") ||
		strings.HasPrefix(i.Suffix, "rc")
}

// Text returns the version in plain text format
// Examples:
//   - "v0.1.0"
//   - "v0.1.0 (4fd00)"
//   - "v0.1.0-beta"
//   - "v0.1.0 (4fd00) [BETA]"
func (i *Info) Text() string {
	base := fmt.Sprintf("v%d.%d.%d", i.Major, i.Minor, i.Patch)

	// Add hash if present
	if i.Hash != "" {
		base += fmt.Sprintf(" (%s)", i.Hash)
	}

	// Add suffix if present
	if i.Suffix != "" {
		base += fmt.Sprintf(" [%s]", strings.ToUpper(i.Suffix))
	}

	return base
}

// JSON returns the version information as a JSON string
func (i *Info) JSON() (string, error) {
	vj := VersionJSON{
		Version: fmt.Sprintf("%d.%d.%d", i.Major, i.Minor, i.Patch),
	}

	if i.Hash != "" {
		vj.Hash = &i.Hash
	}

	if i.Suffix != "" {
		vj.BuildType = &i.Suffix
	}

	bytes, err := json.Marshal(vj)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// JSONPretty returns the version information as a pretty-printed JSON string
func (i *Info) JSONPretty() (string, error) {
	vj := VersionJSON{
		Version: fmt.Sprintf("%d.%d.%d", i.Major, i.Minor, i.Patch),
	}

	if i.Hash != "" {
		vj.Hash = &i.Hash
	}

	if i.Suffix != "" {
		vj.BuildType = &i.Suffix
	}

	bytes, err := json.MarshalIndent(vj, "", "  ")
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// ToJSON converts Info to VersionJSON struct
func (i *Info) ToJSON() VersionJSON {
	vj := VersionJSON{
		Version: fmt.Sprintf("%d.%d.%d", i.Major, i.Minor, i.Patch),
	}

	if i.Hash != "" {
		vj.Hash = &i.Hash
	}

	if i.Suffix != "" {
		vj.BuildType = &i.Suffix
	}

	return vj
}
