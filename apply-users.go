package main

import (
	"fmt"
	"os/exec"
)

func ApplyUsers(users []*UserConfig) error {
	if len(users) == 0 {
		return nil
	}
	for _, user := range users {
		err := ApplyUser(user)
		if err != nil {
			return err
		}

	}
	return nil
}

func ApplyUser(user *UserConfig) error {
	if user.Username == "" {
		return fmt.Errorf("field \"Username\" is empty")
	}
	if user.Password != "" {
		cmd := exec.Command("net", "user", user.Username, user.Password)
		return cmd.Run()
	}
	return nil
}
