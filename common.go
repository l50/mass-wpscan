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

// Execute an input command
// Takes cmd, the command to run
// Takes wg, a sync.WaitGroup
// Returns a string with the output result of the command
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

func splitStringSpaceSlice(s string) []string {
	r := regexp.MustCompile("[^\\s]+")
	sl := r.FindAllString(s, -1)
	return sl
}
