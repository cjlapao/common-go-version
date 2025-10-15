# common-go-version

Utilities for parsing semantic version strings, formatting release metadata, and printing friendly banners for Go CLIs and services. Built on top of [`go-figure`](https://github.com/common-nighthawk/go-figure) so it ships with dozens of embedded ASCII art fonts.

## Features
- Parse versions such as `1.2.3`, `1.2.3:ABC123`, or `1.2.3-beta` into a rich `Info` struct.
- Render release information as human-readable text, compact strings, or JSON/pretty JSON.
- Generate ASCII art for application names and assemble configurable launch banners.
- Drop-in helpers (`QuickPrint*`) that parse a version and print a banner in one call.
- Works without external font files thanks to the embedded `go-figure` typefaces.

## Installation
```bash
go get github.com/cjlapao/common-go-version/version
```

## Quick Start
```go
package main

import (
    "log"
    "github.com/cjlapao/common-go-version/version"
)

func main() {
    info, err := version.Parse("0.2.0:EF06A10-dev")
    if err != nil {
        log.Fatal(err)
    }

    info.Author = "MoneyGrow Team"
    info.Company = "MoneyGrow AI Inc."
    info.Repo = "https://github.com/example/moneygrow-ai"

    version.Print("MoneyGrow AI", info)
}
```

Example banner output:
```
********************************************************************************
*                                                                              *
*                              MoneyGrow AI                                    *
*                                                                              *
*                                v0.2.0 EF06A10 [DEV]                          *
*                                                                              *
* Author:    MoneyGrow Team                                                    *
* Company:   MoneyGrow AI Inc.                                                 *
* Repo:      https://github.com/example/moneygrow-ai                           *
*                                                                              *
********************************************************************************
```

## Working with versions
```go
info, _ := version.Parse("1.2.3:ABC123-beta")

info.String()      // "1.2.3:ABC123-beta"
info.Short()       // "v1.2.3-beta"
info.Text()        // "v1.2.3 (ABC123) [BETA]"
info.IsRelease()   // false
info.IsDev()       // false
info.IsPreRelease()// true

jsonStr, _ := info.JSON()       // {"version":"1.2.3","hash":"ABC123","build_type":"beta"}
pretty, _ := info.JSONPretty()  // Multi-line JSON string
```

## ASCII art and banners
Use ASCII art to highlight your application name or stay with simple text. The library supports auto-width banners, custom borders, and dozens of fonts.

```go
info := &version.Info{Major: 1, Minor: 0, Patch: 0}
opts := version.BannerOptions{
    UseASCII:   true,
    AutoWidth:  true,
    FontStyle:  version.FontStyleStarwars, // any font in ascii.go
}

banner := version.BannerWithOptions("MoneyGrow AI", info, opts)
```

Generate ASCII art directly:
```go
lines := version.GenerateASCIIArt("MoneyGrow", 80)
for _, line := range lines {
    fmt.Println(line)
}
```
Set `maxWidth` to `0` for the natural width, or limit the width to wrap long names. For alternate styles pick any `version.FontStyle*` constant.

## Convenience printing
`QuickPrint`, `QuickPrintWithStyle`, and `QuickPrintWithOptions` parse a version string, fill metadata, and print the banner for you:
```go
err := version.QuickPrint(
    "MoneyGrow AI",
    "0.2.0:EF06A10-dev",
    "MoneyGrow Team",
    "MoneyGrow AI Inc.",
    "2025 MoneyGrow AI Inc.",
    "https://github.com/example/moneygrow-ai",
)
```

## Development
- The repository tracks the current release in the `VERSION` file.
- Run the test suite with `go test ./...`.
- Examples under `example/` double as executable documentation; run them via `go test ./example -run Example`.

## Contributing
Issues and pull requests are welcome. Please make sure tests and examples continue to pass before submitting changes.
