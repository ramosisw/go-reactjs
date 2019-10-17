/*
* MIT License
*
* Copyright (c) 2019 Julio C. Ramos
*
 */

//Script to build application
package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

//Command type to exec on command line
type Command struct {
	Name string
	Args []string
	Dir  string
}

func main() {
	commands := []Command{
		{"go", []string{"get", "github.com/go-bindata/go-bindata/go-bindata"}, ""},
		{"go", []string{"get", "github.com/elazarl/go-bindata-assetfs/go-bindata-assetfs"}, ""},
		{"npm", []string{"install"}, "frontend"},
		{"npm", []string{"run", "build"}, "frontend"},
		{"go-bindata-assetfs", []string{"-pkg=frontend", "-nocompress=false", "-o=frontend/frontend.go", "frontend/build/..."}, ""},
		{"go", []string{"build", "-v", "-ldflags", "-s -w"}, ""},
	}
	//Remove previus frontend if exists

	if _, err := os.Stat("frontend/frontend.go"); os.IsExist(err) {
		if err := os.Remove("frontend/frontend.go"); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	log.Println("Building please wait...")
	for _, command := range commands {
		log.Printf("Runing: %s %s", command.Name, strings.Join(command.Args, " "))
		cmd := exec.Command(command.Name, command.Args...)
		if command.Dir != "" {
			cmd.Dir = command.Dir
		}
		if errStr, err := cmd.Output(); err != nil {
			log.Fatalf("Error: %v", string(errStr))
		}
	}
}
