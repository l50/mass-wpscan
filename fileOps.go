// Copyright 2017 Jayson Grace and Ron Minnich. All rights reserved
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"io/ioutil"
	"strings"
)

// readLines reads an input file into memory from the specified path.
// Upon successfully reading the file in, it will return the lines which
// make up the file in the form of a []string.
func readLines(filePath string) ([]string, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return RemoveTrailingEmptyStringsInStringArray(strings.Split(string(b), "\n")), nil
}

// RemoveTrailingEmptyStringsinStringArray removes any empty strings that are trailing
// a given slice.
func RemoveTrailingEmptyStringsInStringArray(sa []string) []string {
	lastNonEmptyStringIndex := len(sa) - 1
	for i := lastNonEmptyStringIndex; i >= 0; i-- {
		if sa[i] == "" {
			lastNonEmptyStringIndex--
		} else {
			break
		}
	}
	return sa[0 : lastNonEmptyStringIndex+1]
}
