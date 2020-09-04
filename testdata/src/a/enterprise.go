// Program explanation
// Author info, etc

// +buildmiss
// +build !pro,!enterprise

// +build nonewline
package main

var _ = `
// +build insidevar
`

/* Multiline comment
 * +build multiline
 */
func init() {
  features = append(
    features,
    "Enterprise Feature 1",
    "Enterprise Feature 2",
  )
}
