package version

import (
	"encoding/json"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name        string
		versionStr  string
		wantMajor   int
		wantMinor   int
		wantPatch   int
		wantHash    string
		wantSuffix  string
		shouldError bool
	}{
		{
			name:       "simple version",
			versionStr: "1.2.3",
			wantMajor:  1,
			wantMinor:  2,
			wantPatch:  3,
		},
		{
			name:       "version with hash",
			versionStr: "0.2.0:EF06A10",
			wantMajor:  0,
			wantMinor:  2,
			wantPatch:  0,
			wantHash:   "EF06A10",
		},
		{
			name:       "version with hash and dev suffix",
			versionStr: "0.2.0:EF06A10-dev",
			wantMajor:  0,
			wantMinor:  2,
			wantPatch:  0,
			wantHash:   "EF06A10",
			wantSuffix: "dev",
		},
		{
			name:       "version with alpha suffix",
			versionStr: "1.0.0-alpha",
			wantMajor:  1,
			wantMinor:  0,
			wantPatch:  0,
			wantSuffix: "alpha",
		},
		{
			name:       "version with beta suffix",
			versionStr: "2.1.5-beta",
			wantMajor:  2,
			wantMinor:  1,
			wantPatch:  5,
			wantSuffix: "beta",
		},
		{
			name:       "version with rc suffix",
			versionStr: "3.0.0-rc1",
			wantMajor:  3,
			wantMinor:  0,
			wantPatch:  0,
			wantSuffix: "rc1",
		},
		{
			name:        "invalid format",
			versionStr:  "invalid",
			shouldError: true,
		},
		{
			name:        "incomplete version",
			versionStr:  "1.2",
			shouldError: true,
		},
		{
			name:       "version with canary suffix",
			versionStr: "0.0.1-canary",
			wantMajor:  0,
			wantMinor:  0,
			wantPatch:  1,
			wantSuffix: "canary",
		},
		{
			name:       "version with hash only",
			versionStr: "0.0.1:4f00",
			wantMajor:  0,
			wantMinor:  0,
			wantPatch:  1,
			wantHash:   "4F00",
		},
		{
			name:        "invalid format - hash without version",
			versionStr:  "4:ff00-beta",
			shouldError: true,
		},
		{
			name:        "invalid format - only hash",
			versionStr:  ":abc123",
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info, err := Parse(tt.versionStr)

			if tt.shouldError {
				if err == nil {
					t.Errorf("Parse() expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Parse() unexpected error: %v", err)
				return
			}

			if info.Major != tt.wantMajor {
				t.Errorf("Major = %d, want %d", info.Major, tt.wantMajor)
			}
			if info.Minor != tt.wantMinor {
				t.Errorf("Minor = %d, want %d", info.Minor, tt.wantMinor)
			}
			if info.Patch != tt.wantPatch {
				t.Errorf("Patch = %d, want %d", info.Patch, tt.wantPatch)
			}
			if info.Hash != tt.wantHash {
				t.Errorf("Hash = %q, want %q", info.Hash, tt.wantHash)
			}
			if info.Suffix != tt.wantSuffix {
				t.Errorf("Suffix = %q, want %q", info.Suffix, tt.wantSuffix)
			}
		})
	}
}

func TestInfoMethods(t *testing.T) {
	info := &Info{
		Version: "1.2.3:ABC123-dev",
		Major:   1,
		Minor:   2,
		Patch:   3,
		Hash:    "ABC123",
		Suffix:  "dev",
	}

	if info.String() != "1.2.3:ABC123-dev" {
		t.Errorf("String() = %q, want %q", info.String(), "1.2.3:ABC123-dev")
	}

	if info.Short() != "v1.2.3-dev" {
		t.Errorf("Short() = %q, want %q", info.Short(), "v1.2.3-dev")
	}

	if info.IsRelease() {
		t.Errorf("IsRelease() = true, want false")
	}

	if !info.IsDev() {
		t.Errorf("IsDev() = false, want true")
	}

	if info.IsPreRelease() {
		t.Errorf("IsPreRelease() = true, want false")
	}

	// Test release version
	releaseInfo := &Info{Major: 2, Minor: 0, Patch: 0}
	if !releaseInfo.IsRelease() {
		t.Errorf("IsRelease() = false, want true for release version")
	}

	// Test pre-release
	betaInfo := &Info{Major: 1, Minor: 0, Patch: 0, Suffix: "beta"}
	if !betaInfo.IsPreRelease() {
		t.Errorf("IsPreRelease() = false, want true for beta version")
	}
}

func TestTextOutput(t *testing.T) {
	tests := []struct {
		name string
		info *Info
		want string
	}{
		{
			name: "version only",
			info: &Info{Major: 0, Minor: 1, Patch: 0},
			want: "v0.1.0",
		},
		{
			name: "version with hash",
			info: &Info{Major: 0, Minor: 1, Patch: 0, Hash: "4fd00"},
			want: "v0.1.0 (4fd00)",
		},
		{
			name: "version with suffix",
			info: &Info{Major: 0, Minor: 1, Patch: 0, Suffix: "beta"},
			want: "v0.1.0 [BETA]",
		},
		{
			name: "version with hash and suffix",
			info: &Info{Major: 0, Minor: 1, Patch: 0, Hash: "4fd00", Suffix: "beta"},
			want: "v0.1.0 (4fd00) [BETA]",
		},
		{
			name: "version with canary",
			info: &Info{Major: 0, Minor: 0, Patch: 1, Suffix: "canary"},
			want: "v0.0.1 [CANARY]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.info.Text()
			if got != tt.want {
				t.Errorf("Text() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestJSONOutput(t *testing.T) {
	tests := []struct {
		name    string
		info    *Info
		wantVer string
		wantHash *string
		wantType *string
	}{
		{
			name:    "version only",
			info:    &Info{Major: 0, Minor: 1, Patch: 0},
			wantVer: "0.1.0",
		},
		{
			name:     "version with hash",
			info:     &Info{Major: 0, Minor: 1, Patch: 0, Hash: "4fd00"},
			wantVer:  "0.1.0",
			wantHash: strPtr("4fd00"),
		},
		{
			name:     "version with suffix",
			info:     &Info{Major: 0, Minor: 1, Patch: 0, Suffix: "beta"},
			wantVer:  "0.1.0",
			wantType: strPtr("beta"),
		},
		{
			name:     "version with hash and suffix",
			info:     &Info{Major: 0, Minor: 1, Patch: 0, Hash: "4fd00", Suffix: "beta"},
			wantVer:  "0.1.0",
			wantHash: strPtr("4fd00"),
			wantType: strPtr("beta"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonStr, err := tt.info.JSON()
			if err != nil {
				t.Fatalf("JSON() error = %v", err)
			}

			var got VersionJSON
			if err := json.Unmarshal([]byte(jsonStr), &got); err != nil {
				t.Fatalf("json.Unmarshal() error = %v", err)
			}

			if got.Version != tt.wantVer {
				t.Errorf("Version = %q, want %q", got.Version, tt.wantVer)
			}

			if (got.Hash == nil) != (tt.wantHash == nil) {
				t.Errorf("Hash presence mismatch: got %v, want %v", got.Hash, tt.wantHash)
			} else if got.Hash != nil && *got.Hash != *tt.wantHash {
				t.Errorf("Hash = %q, want %q", *got.Hash, *tt.wantHash)
			}

			if (got.BuildType == nil) != (tt.wantType == nil) {
				t.Errorf("BuildType presence mismatch: got %v, want %v", got.BuildType, tt.wantType)
			} else if got.BuildType != nil && *got.BuildType != *tt.wantType {
				t.Errorf("BuildType = %q, want %q", *got.BuildType, *tt.wantType)
			}
		})
	}
}

func TestToJSON(t *testing.T) {
	info := &Info{
		Major:  1,
		Minor:  2,
		Patch:  3,
		Hash:   "ABC123",
		Suffix: "dev",
	}

	vj := info.ToJSON()

	if vj.Version != "1.2.3" {
		t.Errorf("Version = %q, want %q", vj.Version, "1.2.3")
	}
	if vj.Hash == nil || *vj.Hash != "ABC123" {
		t.Errorf("Hash = %v, want %q", vj.Hash, "ABC123")
	}
	if vj.BuildType == nil || *vj.BuildType != "dev" {
		t.Errorf("BuildType = %v, want %q", vj.BuildType, "dev")
	}
}

// Helper function to create string pointers
func strPtr(s string) *string {
	return &s
}
