package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	command := "pwd"
	cmd := exec.Command("/bin/sh", "-c", command)
	bytes, err := cmd.Output()
	if err != nil {
		log.Println(err)
		os.Exit(3)
	}
	resp := string(bytes)
	log.Println(resp)
	os.Exit(0)
}
