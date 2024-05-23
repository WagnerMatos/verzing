package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// updateVersion reads the current version, increments it based on the commit type and breaking changes, and writes it back.
func updateVersion(commitType string, breakingChange bool) (string, error) {
	// Read the current version from the VERSION.md file
	version, err := readVersion()
	if err != nil {
		return "", err
	}

	// Parse the version numbers
	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid version format")
	}

	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])
	patch, _ := strconv.Atoi(parts[2])

	// Increment the version based on the commit type and breaking changes
	if breakingChange {
		major++
		minor = 0
		patch = 0
	} else {
		switch commitType {
		case "Fix":
			patch++
		default:
			minor++
			patch = 0
		}
	}

	// Construct the new version string
	newVersion := fmt.Sprintf("%d.%d.%d", major, minor, patch)
	if err := writeVersion(newVersion); err != nil {
		return "", err
	}

	return newVersion, nil
}

// readVersion reads the version from the VERSION.md file
func readVersion() (string, error) {
	file, err := os.Open("VERSION.md")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text()), nil
	}

	return "", fmt.Errorf("failed to read version")
}

// writeVersion writes the updated version to the VERSION.md file
func writeVersion(version string) error {
	file, err := os.OpenFile("VERSION.md", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(version + "\n")
	return err
}
