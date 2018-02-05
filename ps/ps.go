package ps

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

const (
	Cmd  = "cscript.exe"
	Opts = "/nologo"
)

var Folder = path.Join(os.Getenv("GOPATH"), "/src/github.com/sbrow/skirmish/ps", "vbs")

func Start() error {
	_, err := run("start")
	return err
}

func Open(path string) ([]byte, error) {
	return run("open", path)
}

func Close() error {
	_, err := run("close")
	return err
}

func Quit() ([]byte, error) {
	return run("quit")
}

func Js(args ...string) ([]byte, error) {
	return run("dojs", args...)
}
func Wait(msg string) {
	fmt.Print(msg)
	var input string
	fmt.Scanln(&input)
}

func run(name string, args ...string) ([]byte, error) {
	ext := ".txt"
	switch runtime.GOOS {
	case "windows":
		ext = ".vbs"
	case "darwin":
		ext = ".applescript"
	}
	if !strings.HasSuffix(name, ext) {
		name += ext
	}
	args = append([]string{Opts, path.Join(Folder, name)}, args...)
	cmd := exec.Command(Cmd, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}
	return out.Bytes(), err
}
