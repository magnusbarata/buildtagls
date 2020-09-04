package buildtagls_test

import (
	"testing"

	"github.com/magnusbarata/buildtagls"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, buildtagls.Analyzer, "a")
}
