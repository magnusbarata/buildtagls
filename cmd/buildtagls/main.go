package main

import (
        "buildtagls"
        "golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(buildtagls.Analyzer) }
