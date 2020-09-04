package main

import (
        "github.com/magnusbarata/buildtagls"
        "golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(buildtagls.Analyzer) }
