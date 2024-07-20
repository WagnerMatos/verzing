package main

import (
	"fmt"
	"os/exec"
)

type Executor interface {
	Run(cmd *exec.Cmd) error
}

type RealExecutor struct{}

func (e RealExecutor) Run(cmd *exec.Cmd) error {
	return cmd.Run()
}

func commitChanges(message string, executor Executor, versioner FileVersioner) error {
	cmd := exec.Command("git", "add", ".")
	if err := executor.Run(cmd); err != nil {
		return err
	}

	cmd = exec.Command("git", "commit", "-m", message)
	if err := executor.Run(cmd); err != nil {
		return err
	}

	// Read the version for tagging
	tag, err := versioner.ReadVersion()
	if err != nil {
		return fmt.Errorf("failed to read version: %w", err)
	}

	// After a successful commit, create a tag
	if err := createTag(tag, executor); err != nil {
		return fmt.Errorf("failed to create tag: %w", err)
	}

	return nil
}

func createTag(tag string, executor Executor) error {
	cmd := exec.Command("git", "tag", tag)
	return executor.Run(cmd)
}
