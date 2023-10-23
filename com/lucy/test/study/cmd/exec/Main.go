package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	command := exec.Command("java", "-version")
	var stdout, stderr bytes.Buffer
	command.Stdout = &stdout
	command.Stderr = &stderr
	err := command.Run()
	if err != nil {
		fmt.Println("err:" + err.Error())
		return
	}
	fmt.Println("stdout:" + stdout.String())
	fmt.Println("stderr:" + stderr.String())
}
