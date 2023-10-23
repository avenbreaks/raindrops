package main

import (
	_ "embed"

	"github.com/avenbreaks/raindrops/command/root"
	"github.com/avenbreaks/raindrops/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
