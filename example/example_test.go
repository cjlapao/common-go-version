package version_test

import (
	"fmt"

	"github.com/cjlapao/common-go-version/version"
)

func ExampleParse() {
	info, err := version.Parse("1.2.3:ABC123-dev")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Version: %s\n", info.String())
	fmt.Printf("Short: %s\n", info.Short())
	fmt.Printf("Major: %d, Minor: %d, Patch: %d\n", info.Major, info.Minor, info.Patch)
	fmt.Printf("Hash: %s\n", info.Hash)
	fmt.Printf("Suffix: %s\n", info.Suffix)
	fmt.Printf("IsDev: %t\n", info.IsDev())

	// Output:
	// Version: 1.2.3:ABC123-dev
	// Short: v1.2.3-dev
	// Major: 1, Minor: 2, Patch: 3
	// Hash: ABC123
	// Suffix: dev
	// IsDev: true
}

func ExampleBanner() {
	info, _ := version.Parse("0.2.0:EF06A10-dev")
	info.Author = "MoneyGrow Team"
	info.Company = "MoneyGrow AI Inc."
	info.Copyright = "2025 MoneyGrow AI Inc."
	info.Repo = "https://github.com/example/moneygrow-ai"

	banner := version.Banner("MoneyGrow AI", info)
	fmt.Println(banner)
}

func ExampleQuickPrint() {
	err := version.QuickPrint(
		"MoneyGrow AI",
		"0.2.0:EF06A10-dev",
		"MoneyGrow Team",
		"MoneyGrow AI Inc.",
		"2025 MoneyGrow AI Inc.",
		"https://github.com/example/moneygrow-ai",
	)
	if err != nil {
		panic(err)
	}
}
