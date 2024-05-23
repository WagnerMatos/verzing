package main

import (
	"os/exec"
)

type Executor interface {
	Run(cmd *exec.Cmd) error
}

type RealExecutor struct{}

func (e RealExecutor) Run(cmd *exec.Cmd) error {
	return cmd.Run()
}

func commitChanges(message string, executor Executor) error {
	cmd := exec.Command("git", "add", ".")
	if err := executor.Run(cmd); err != nil {
		return err
	}
	cmd = exec.Command("git", "commit", "-m", message)
	return executor.Run(cmd)
}
