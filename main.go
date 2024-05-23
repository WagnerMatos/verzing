package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Commit type
	fmt.Println("Select commit type:")
	commitTypes := []string{"Feat", "Fix", "Docs", "Style", "Refactor", "Perf", "Test", "Build", "CI", "Chore", "Revert"}
	for i, t := range commitTypes {
		fmt.Printf("%d. %s\n", i+1, t)
	}
	choice, _ := reader.ReadString('\n')
	index, _ := strconv.Atoi(strings.TrimSpace(choice))
	selectedType := commitTypes[index-1]

	// Short description
	fmt.Println("Enter short commit description:")
	shortDesc, _ := reader.ReadString('\n')
	shortDesc = strings.TrimSpace(shortDesc)

	// Long description
	fmt.Println("Enter long commit description (press ENTER to skip):")
	longDesc, _ := reader.ReadString('\n')
	longDesc = strings.TrimSpace(longDesc)

	// Breaking changes
	fmt.Println("Are there any breaking changes? (yes/no)")
	breakingChangeInput, _ := reader.ReadString('\n')
	breakingChange := strings.TrimSpace(breakingChangeInput) == "yes"

	// Create a FileVersioner instance
	fileVersioner := FileVersioner{FilePath: "VERSION.md"}

	// Handle versioning
	version, err := updateVersion(fileVersioner, selectedType, breakingChange)
	if err != nil {
		fmt.Println("Error updating version:", err)
		return
	}

	// Update CHANGELOG.md
	if err := updateChangelog(version, shortDesc, longDesc); err != nil {
		fmt.Println("Error updating CHANGELOG:", err)
		return
	}

	// Commit changes
	commitMessage := fmt.Sprintf("%s: %s\n\n%s", selectedType, shortDesc, longDesc)
	if err := commitChanges(commitMessage); err != nil {
		fmt.Println("Error committing changes:", err)
		return
	}

	fmt.Println("Commit and changelog update completed successfully.")
}

//func updateVersion(commitType string, breakingChange bool) (string, error) {
//	// Dummy version update logic
//	return "v1.0.0", nil // Replace with actual versioning logic
//}

func commitChanges(message string) error {
	cmd := exec.Command("git", "add", ".")
	if err := cmd.Run(); err != nil {
		return err
	}
	cmd = exec.Command("git", "commit", "-m", message)
	return cmd.Run()
}

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
