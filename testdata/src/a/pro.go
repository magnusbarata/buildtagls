// +build !pro

// comment +build

/*
 * Multiline comment
 */

package main

// +build afterpackage

func init() {
  features = append(
    features,
    "Pro Feature 1", // Comment here
    "Pro Feature 2",
  )
}
