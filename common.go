// Copyright 2017 Jayson Grace. All rights reserved
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package main

import "regexp"

// strToSlice takes a string and a delimiter in the
// form of a regex. It will use this to split a string
// into a slice, and return it.
func strToSlice(s string, delimiter string) []string {
	r := regexp.MustCompile(delimiter)
	slice := r.FindAllString(s, -1)
	return slice
}
