package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/alexflint/go-arg"
)

func apply(parser *arg.Parser, args Args) {
	if args.Apply.ConfigFile == "" {
		self, err := os.Executable()
		if err == nil {
			args.Apply.ConfigFile = filepath.Join(filepath.Dir(self), "config.json")
		}
	}
	if _, err := os.Stat(args.Apply.ConfigFile); err != nil {
		parser.Fail("Cannot open " + args.Apply.ConfigFile + ". Provide a valid path to config.json file")
		return
	}
	config, err := ParseConfig(args.Apply.ConfigFile)
	if err != nil {
		parser.Fail("Cannot parse config.json. Provide a path to a valid config.json file")
		return
	}

	err = ApplyNetwork(config.Network)
	if err != nil {
		log.Println("Error during network config: " + err.Error())
	}

	err = ApplyUsers(config.Users)
	if err != nil {
		log.Println("Error during users config: " + err.Error())
	}
}
