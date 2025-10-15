package version

import (
	"fmt"

	"github.com/cjlapao/common-go-version/version"
)

// Example demonstrating simple text banner (no ASCII art)
func ExampleBannerWithOptions_simpleText() {
	info := &version.Info{
		Major:     1,
		Minor:     0,
		Patch:     0,
		Author:    "John Doe",
		Company:   "Acme Corp",
		Copyright: "2025 Acme Corp",
		Repo:      "github.com/acme/product",
	}

	opts := version.BannerOptions{
		UseASCII:   false, // Use simple text instead of ASCII art
		AutoWidth:  false,
		FixedWidth: 80,
		FontStyle:  version.FontStyleSlant, // Ignored when UseASCII is false
	}

	fmt.Println(version.BannerWithOptions("MyApp", info, opts))
}

// Example demonstrating auto-width banner (fits content)
func ExampleBannerWithOptions_autoWidth() {
	info := &version.Info{
		Major:     1,
		Minor:     2,
		Patch:     3,
		Author:    "Jane Smith",
		Company:   "Tech Inc",
		Copyright: "2025 Tech Inc",
		Repo:      "github.com/tech/app",
	}

	opts := version.BannerOptions{
		UseASCII:  true,
		AutoWidth: true, // Banner will auto-size to fit content
		FontStyle: version.FontStyleBig,
	}

	fmt.Println(version.BannerWithOptions("App", info, opts))
}

// Example demonstrating custom fixed width banner
func ExampleBannerWithOptions_customWidth() {
	info := &version.Info{
		Major:  2,
		Minor:  0,
		Patch:  0,
		Suffix: "beta",
	}

	opts := version.BannerOptions{
		UseASCII:   true,
		AutoWidth:  false,
		FixedWidth: 120, // Custom width of 120 characters
		FontStyle:  version.FontStyleStarwars,
	}

	fmt.Println(version.BannerWithOptions("MoneyGrow AI", info, opts))
}

// Example demonstrating simple text with auto-width
func ExampleBannerWithOptions_simpleAutoWidth() {
	info := &version.Info{
		Major:     3,
		Minor:     5,
		Patch:     7,
		Hash:      "abc1234",
		Author:    "Development Team",
		Copyright: "2025",
	}

	opts := version.BannerOptions{
		UseASCII:  false, // Simple text
		AutoWidth: true,  // Auto-size to fit
	}

	fmt.Println(version.BannerWithOptions("Microservice API", info, opts))
}
