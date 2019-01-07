// Copyright 2017 Jayson Grace and Ron Minnich. All rights reserved
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/fatih/color"
)

var (
	inputFile string
	wpParams  string
	outfile   string
	errmsg    = color.Red
	warn      = color.Yellow
	msg       = color.Green
)

type cmdResult struct {
	err error
	out []byte
}

// scanTargets runs wpscan against a specified set of targets concurrently.
// Once it has finished, it returns the output results of the
// various wpscan instances in the form of a slice.
func scanTargets(targets []string, wpParams string, res chan *cmdResult) {
	wg := new(sync.WaitGroup)
	for _, target := range targets {
		wg.Add(1)
		msg("Scanning %s with wpscan, please wait...", target)
		go func(t string) {
			defer wg.Done()
			opts := append([]string{"--url", t}, strings.Fields(wpParams)...)
			out, err := exec.Command("wpscan", opts...).CombinedOutput()
			res <- &cmdResult{out: out, err: err}
		}(target)
	}
	wg.Wait()
	close(res)
}

func main() {
	flag.Parse()

	if flag.NFlag() == 0 || validateInput() == false {
		usage()
	}

	paramSlice := strToSlice(wpParams, "[^\\s]+")

	validateWpParams(paramSlice)

	msg("Start wpscan --update")
	output, err := exec.Command("wpscan", "--update").Output()
	if err != nil {
		errmsg("%s", err)
	}
	warn("%s", output)

	targets, err := readLines(inputFile)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	res := make(chan *cmdResult, 64)
	res <- &cmdResult{out: output, err: err}

	go scanTargets(targets, wpParams, res)

	f := os.Stdout
	if outfile != "" {
		// create the output file if one is specified as an input parameter
		f, err = os.Create(outfile)
		if err != nil {
			log.Fatalf("writeLines: %s", err)
		}
	}
	for r := range res {
		if _, err := f.Write(r.out); err != nil {
			log.Fatal(err)
		}
	}
}

// init specifies the input parameters which mass-wpscan can take.
func init() {
	flag.StringVar(&inputFile, "i", "", "Input file with targets.")
	flag.StringVar(&wpParams, "p", "", "Arguments to run with wpscan.")
	flag.StringVar(&outfile, "o", "", "File to output information to.")
}

// usage prints the usage instructions for mass-wpscan.
func usage() {
	os.Args[0] = os.Args[0] + " [options]"
	flag.Usage()
	os.Exit(1)
}

// validateInput ensures that the user inputs proper arguments into mass-wpscan.
func validateInput() bool {
	if inputFile == "" || wpParams == "" {
		errmsg("You must specify an input file with targets and parameters for wpscan!")
		errmsg("Example: mass-wpscan -i vuln_targets.txt -p \"--rua -e vt,tt,u,vp\"")
		errmsg("Another Example: mass-wpscan -i vuln_targets.txt -p \" \" -o output.txt")
		return false
	}
	return true
}

// validateWpParams ensures that the --url parameter is omitted if specified.
// This is due to the nature of the program, and why it was created in the first
// place. If you want to specify --url, you're probably only scanning one system
// and should just use wpscan directly.
func validateWpParams(parameters []string) {
	for _, p := range parameters {
		if p == "--url" {
			log.Fatal("You can not include the --url parameter, all targets should be in your input file!")
		}
	}
}
