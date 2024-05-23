package main

import (
	"fmt"
	"os/exec"
	"testing"
)

// MockExecutor is used to simulate command execution in tests
type MockExecutor struct {
	ShouldFail bool
}

func (m MockExecutor) Run(cmd *exec.Cmd) error {
	if m.ShouldFail {
		return fmt.Errorf("mock error")
	}
	return nil
}

func TestCommitChanges(t *testing.T) {
	executor := MockExecutor{ShouldFail: false}
	err := commitChanges("Add dummy file", executor)
	if err != nil {
		t.Errorf("commitChanges should not fail: %s", err)
	}

	executor = MockExecutor{ShouldFail: true}
	err = commitChanges("Add dummy file", executor)
	if err == nil {
		t.Errorf("commitChanges should fail")
	}
}
