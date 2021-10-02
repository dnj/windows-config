package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/sys/windows"

	"github.com/alexflint/go-arg"
)

type ApplyCmd struct {
	ConfigFile string `arg:"-f" help:"Path to config.json file"`
}

type Args struct {
	Apply *ApplyCmd `arg:"subcommand:apply"`
}

func (Args) Version() string {
	return "windows-config 1.0.0"
}

func main() {

	var args Args
	parser := arg.MustParse(&args)

	switch {
	case args.Apply != nil:
		if !amAdmin() {
			runMeElevated()
			return
		}
		apply(parser, args)
	default:
		parser.WriteUsage(os.Stderr)
	}

}

func runMeElevated() {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		fmt.Println(err)
	}
}

func amAdmin() bool {
	fd, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	fd.Close()
	if err != nil {
		return false
	}
	return true
}
