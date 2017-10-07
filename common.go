// Copyright 2017 Jayson Grace. All rights reserved
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"sync"
)

// exeCmd executes an input command.
// Once the command have successfully ben run, it will
// return a string with the output result of the command.
// TODO: Add concurrent operations to speed things up
func exeCmd(cmd string, wg *sync.WaitGroup) string {
	fmt.Println("Running: ", cmd)
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		errmsg("%s", err)
	}
	warn("%s", out)
	wg.Done()
	return string(out)
}

// strToSlice takes a string and a delimiter in the
// form of a regex. It will use this to split a string
// into a slice, and return it.
func strToSlice(s string, delimiter string) []string {
	r := regexp.MustCompile("[^\\s]+")
	slice := r.FindAllString(s, -1)
	return slice
}
