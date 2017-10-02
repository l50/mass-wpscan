// Copyright 2017 Jayson Grace. All rights reserved
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"sync"

	"github.com/fatih/color"
)

// Execute an input command
// Takes cmd, the command to run
// Takes wg, a sync.WaitGroup
// Returns a string with the output result of the command
// TODO: Add concurrent operations to speed things up
func ExeCmd(cmd string, wg *sync.WaitGroup) string {
	fmt.Println("Running: ", cmd)
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
	wg.Done()
	return string(out)
}

// Convert a file into an array
// Takes filename, the file to read
// Returns a string array with the content of the file
func FileToArray(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		println(err.Error())
		return nil
	}
	fileArr := strings.Split(string(content), "\n")
	return fileArr
}

// Remove --url from arguments if included and convert arguments into a string
// Takes inputArgs, the parameters that were input
// Returns args, the inputArgs in string format without the --url parameter (if it was input)
func processUrlParams(inputArgs []string) string {
	for i, v := range inputArgs {
		if v == "--url" {
			inputArgs = append(inputArgs[:i], inputArgs[i+1:]...)
			break
		}
	}
	args := strings.Join(inputArgs, " ")
	return args
}

func scanTargets(targets []string, inputArgs []string, output string, wg *sync.WaitGroup) string {
	var wp_out string
	args := processUrlParams(inputArgs)
	for _, target := range targets {
		color.Green("Scanning %s with wpscan, please wait...", target)
		cmd := "wpscan" + " --url " + target + " " + args
		wg.Add(1)
		wp_out += ExeCmd(string(cmd), wg)
		wp_out += "\n"
		wg.Wait()
	}
	fmt.Print("hi")
	return output + wp_out
}

func ArrayToFile(contents []byte, file string) {
	ioutil.WriteFile(file, contents, 0644)
}

//func usage() {
//	fmt.Println("Usage: mass-wpscan [options] keyword")
//	flag.PrintDefaults()
//}

func main() {
	//flag.Usage = usage
	//flag.BoolVar(&config.help, "h", false, "Show help page.")
	//flag.BoolVar(&config.args, "a", false, "Arguments to run with wpscan.")
	//flag.BoolVar(&config.input, "i", false, "Input file with targets.")
	//flag.BoolVar(&config.output, "o", false, "Output file.")

	//flag.Parse()

	//if len(flag.Args()) != 1 {
	//	flag.Usage()
	//	os.Exit(0)
	//}
	wg := new(sync.WaitGroup)
	color.Green("Updating wpscan, please wait...")
	wg.Add(1)
	output := ExeCmd("wpscan --update", wg)
	wg.Wait()
	// TODO: Make this parameter an input argument
	targets := FileToArray("vuln_targets.txt")
	// TODO: Make this parameter an input argument
	wpscanArgs := []string{"-r", "--batch", "-e", "vt,tt,u,vp"}
	// TODO: Make this parameter an input argument
	outfile := "output.txt"
	// TODO: require input file be specified
	output = scanTargets(targets, wpscanArgs, output, wg)
	fmt.Println(output)
	ArrayToFile([]byte(output), outfile)
	// TODO: if output file specified, write to output file else, write to stdout
}
