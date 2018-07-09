package main

// import (
// 	"log"
// 	"os"
// 	"os/exec"
// 	"path/filepath"

// 	"github.com/sbrow/skirmish"
// )

// func main2() {
// 	log.SetPrefix("[main] ")
// 	args := os.Args[1:]
// 	cmd := os.Args[0]
// 	switch cmd {
// 	case "ps":
// 		comm := exec.Command(filepath.Join(os.Getenv("GOBIN"), "cmd.exe"), args...)
// 		comm.Stdin = os.Stdin
// 		comm.Stderr = os.Stderr
// 		comm.Run()
// 		return
// 	case "db":
// 		opt := args[0]
// 		if opt == "save" {
// 			skirmish.Dump(skirmish.DataDir)
// 		} else if opt == "load" {
// 			skirmish.Recover(skirmish.DataDir)
// 		}
// 	}
// }
