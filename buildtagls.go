package buildtagls

import (
        "fmt"
        "golang.org/x/tools/go/analysis"
        "regexp"

        "strings"
)

const doc = "buildtagls list all valid and available build tags of a package"

// Analyzer list all valid and available build tags of a package.
var Analyzer = &analysis.Analyzer{
        Name: "buildtagls",
        Doc:  doc,
        Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
        tags := map[string]bool{}

        for _, f := range pass.Files {
                for _, group := range f.Comments {
                        // A +build comment afer package declaration is not valid
                        if group.End()+1 >= f.Package {
                                continue
                        }

                        // Check comments before package declaration
                        for _, comment := range group.List {
                                // Skipping comments not containing +build
                                if !strings.Contains(comment.Text, "+build") {

                                        continue
                                }

                                line := strings.TrimPrefix(comment.Text, "//")
                                line = strings.TrimSpace(line)
                                if strings.HasPrefix(line, "+build") {
                                        regex := regexp.MustCompile(`[, ]+`)
                                        toks := regex.Split(line, -1)
                                        if toks[0] != "+build" {
                                                continue
                                        }

                                        for _, tag := range toks[1:] {
                                                tag := strings.TrimPrefix(tag, "!")
                                                tags[tag] = true
                                        }
                                }
                        }
                }
        }

        for tag := range tags {
                fmt.Println(tag)
        }

        return nil, nil
}
