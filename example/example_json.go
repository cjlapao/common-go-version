package version

import (
	"fmt"
	"log"

	"github.com/cjlapao/common-go-version/version"
)

// Example demonstrating parsing different version formats
func ExampleParse_formats() {
	// Version only
	v1, _ := version.Parse("0.0.1")
	fmt.Println("Version only:", v1.Text())

	// Version with build type
	v2, _ := version.Parse("0.0.1-canary")
	fmt.Println("With build type:", v2.Text())

	// Version with hash
	v3, _ := version.Parse("0.0.1:4f00")
	fmt.Println("With hash:", v3.Text())

	// Version with hash and build type
	v4, _ := version.Parse("0.0.1:4f00-beta")
	fmt.Println("With hash and build type:", v4.Text())

	// Output:
	// Version only: v0.0.1
	// With build type: v0.0.1 [CANARY]
	// With hash: v0.0.1 (4F00)
	// With hash and build type: v0.0.1 (4F00) [BETA]
}

// Example demonstrating invalid version formats
func ExampleParse_invalid() {
	// These will return errors:
	formats := []string{
		"4:ff00-beta", // Missing proper version
		":abc123",     // Missing version
		"1.2",         // Incomplete version
	}

	for _, format := range formats {
		_, err := version.Parse(format)
		if err != nil {
			fmt.Printf("%s: %v\n", format, err)
		}
	}
}

// Example demonstrating JSON output
func ExampleInfo_JSON() {
	// Version only
	v1, _ := version.Parse("0.1.0")
	json1, _ := v1.JSON()
	fmt.Println("Version only:", json1)

	// Version with hash
	v2, _ := version.Parse("0.1.0:4fd00")
	json2, _ := v2.JSON()
	fmt.Println("With hash:", json2)

	// Version with build type
	v3, _ := version.Parse("0.1.0-beta")
	json3, _ := v3.JSON()
	fmt.Println("With build type:", json3)

	// Output:
	// Version only: {"version":"0.1.0"}
	// With hash: {"version":"0.1.0","hash":"4FD00"}
	// With build type: {"version":"0.1.0","build_type":"beta"}
}

// Example demonstrating pretty JSON output
func ExampleInfo_JSONPretty() {
	info, _ := version.Parse("1.2.3:ABC123-beta")
	pretty, _ := info.JSONPretty()
	fmt.Println(pretty)

	// Output:
	// {
	//   "version": "1.2.3",
	//   "hash": "ABC123",
	//   "build_type": "beta"
	// }
}

// Example demonstrating text output
func ExampleInfo_Text() {
	// Different version formats
	versions := []string{
		"0.1.0",
		"0.1.0:4fd00",
		"0.1.0-beta",
		"0.1.0:4fd00-beta",
	}

	for _, v := range versions {
		info, err := version.Parse(v)
		if err != nil {
			log.Printf("Error parsing %s: %v", v, err)
			continue
		}
		fmt.Printf("%-20s -> %s\n", v, info.Text())
	}

	// Output:
	// 0.1.0                -> v0.1.0
	// 0.1.0:4fd00          -> v0.1.0 (4FD00)
	// 0.1.0-beta           -> v0.1.0 [BETA]
	// 0.1.0:4fd00-beta     -> v0.1.0 (4FD00) [BETA]
}
