package main

import (
	"bufio"
	"fmt"
	"os"
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
	file, err := os.OpenFile("CHANGELOG.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening CHANGELOG.md:", err)
		return
	}
	defer func() {
		cerr := file.Close()
		if err == nil { // only overwrite err if it's still nil
			err = cerr
		}
	}()

	if err := updateChangelog(file, version, shortDesc, longDesc); err != nil {
		fmt.Println("Error updating CHANGELOG:", err)
		return
	}

	// Commit changes
	commitMessage := fmt.Sprintf("%s: %s\n\n%s", selectedType, shortDesc, longDesc)
	executor := RealExecutor{}
	if err := commitChanges(commitMessage, executor, fileVersioner); err != nil {
		fmt.Println("Error committing changes:", err)
		return
	}

	fmt.Println("Commit and changelog update completed successfully.")
}
