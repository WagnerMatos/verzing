package main

import (
	"fmt"
	"os"
)

func updateChangelog(version, shortDesc, longDesc string) error {
	file, err := os.OpenFile("CHANGELOG.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	entry := fmt.Sprintf("## %s\n- %s\n", version, shortDesc)
	if longDesc != "" {
		entry += fmt.Sprintf("  %s\n", longDesc)
	}

	if _, err := file.WriteString(entry); err != nil {
		return err
	}
	return nil
}
