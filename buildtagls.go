package buildtagls

import (
	"golang.org/x/tools/go/analysis"
)

const doc = "buildtagls list all valid and available build tags of a package"

// Analyzer list all valid and available build tags of a package.
var Analyzer = &analysis.Analyzer{
	Name: "buildtagls",
	Doc:  doc,
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	return nil, nil
}
