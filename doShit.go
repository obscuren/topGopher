package main

import (
	"bufio"
	"os"
	"os/exec"
	"time"
)

func main() {
	//2 commits per minute
	for {
		loopThatShit()
		time.Sleep(1 * time.Minute)
	}
}

func loopThatShit() {

	filename := "silly.go"

	// open output file
	fo, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	// make a write buffer
	w := bufio.NewWriter(fo)
	w.WriteString("package main\n")
	w.Flush()

	//commit it
	gitAddAll()
	gitCommit()
	gitPush()

	//delete it
	os.Remove(filename)

	//commit it
	gitAddAll()
	gitRemove()
	gitPush()
}

func gitAddAll() {
	app := "git"
	arg0 := "add"
	arg1 := "."
	cmd := exec.Command(app, arg0, arg1)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(out))
}

func gitCommit() {
	app := "git"
	arg0 := "commit"
	arg1 := "-am"
	arg2 := "\"update\""
	cmd := exec.Command(app, arg0, arg1, arg2)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(out))
}

func gitRemove() {
	app := "git"
	arg0 := "commit"
	arg1 := "-am"
	arg2 := "\"remove\""
	cmd := exec.Command(app, arg0, arg1, arg2)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(out))
}

func gitPush() {
	app := "git"
	arg0 := "push"
	arg1 := "origin"
	arg2 := "master"
	cmd := exec.Command(app, arg0, arg1, arg2)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(out))
}
