// FILE: internal/detector/detector.go
package detector

import (
	"encoding/json" // NEW
	"fmt"
	"os"
	"path/filepath"
	"strings" // NEW

	"github.com/Masterminds/semver" // NEW
)

// Archetype represents the detected type of the project.
type Archetype string

// Define constants for each supported project type.
const (
	ArchetypeUnknown        Archetype = "unknown"
	ArchetypeJavaSpringBoot Archetype = "java_spring_boot"
	ArchetypePythonFastAPI  Archetype = "python_fastapi"
	// --- NEW ARYCHETYPES ---
	ArchetypePHPLaravel     Archetype = "php_laravel"
	ArchetypeNodeJSNextJS   Archetype = "nodejs_nextjs"
)

// ProjectProfile struct
type ProjectProfile struct {
	Archetype           Archetype
	LanguageVersion     string // e.g., "8.2", "18", "3.10"
	DatabaseType        string
	DeploymentEnvironment string
}

// fileExists is a helper function to check if a file exists at a given path.
func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// dirExists is a helper function to check if a directory exists.
func dirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// GetProjectProfile scans a directory to identify the project's archetype and language version.
func GetProjectProfile(dirPath string) (*ProjectProfile, error) {
	// --- PHP Laravel Detection ---
	composerPath := filepath.Join(dirPath, "composer.json")
	if fileExists(composerPath) && fileExists(filepath.Join(dirPath, "artisan")) {
		version, err := parsePhpVersionFromComposer(composerPath)
		if err != nil {
			// Could not parse, but it's still a PHP project. Fallback to a default.
			return &ProjectProfile{Archetype: ArchetypePHPLaravel, LanguageVersion: "8.2"}, nil
		}
		return &ProjectProfile{Archetype: ArchetypePHPLaravel, LanguageVersion: version}, nil
	}

	// --- Java Spring Boot Detection ---
	if fileExists(filepath.Join(dirPath, "pom.xml")) || fileExists(filepath.Join(dirPath, "build.gradle")) {
		return &ProjectProfile{Archetype: ArchetypeJavaSpringBoot, LanguageVersion: "17"}, nil // Default to Java 17
	}

	// --- Python FastAPI Detection ---
	if fileExists(filepath.Join(dirPath, "requirements.txt")) {
		content, err := os.ReadFile(filepath.Join(dirPath, "requirements.txt"))
		if err == nil && (strings.Contains(string(content), "fastapi") || strings.Contains(string(content), "FastAPI")) {
			return &ProjectProfile{Archetype: ArchetypePythonFastAPI, LanguageVersion: "3.9"}, nil // Default to Python 3.9
		}
	}

	// --- NodeJS NextJS Detection ---
	packageJSONPath := filepath.Join(dirPath, "package.json")
	if fileExists(packageJSONPath) {
		content, err := os.ReadFile(packageJSONPath)
		if err == nil && strings.Contains(string(content), "\"next\"") {
			return &ProjectProfile{Archetype: ArchetypeNodeJSNextJS, LanguageVersion: "18"}, nil // Default to Node 18
		}
	}

	return nil, fmt.Errorf("could not determine project type")
}

// composer.json structure for PHP version
type Composer struct {
	Require struct {
		Php string `json:"php"`
	} `json:"require"`
}

func parsePhpVersionFromComposer(composerPath string) (string, error) {
	data, err := os.ReadFile(composerPath)
	if err != nil {
		return "", fmt.Errorf("failed to read composer.json: %w", err)
	}

	var composer Composer
	if err := json.Unmarshal(data, &composer); err != nil {
		return "", fmt.Errorf("failed to parse composer.json: %w", err)
	}

	if composer.Require.Php == "" {
		return "", fmt.Errorf("php version not found in composer.json 'require' section")
	}

	// Use Masterminds/semver to parse the version constraint
	// For simplicity, we'll try to find a compatible major.minor version.
	// This is a simplified logic and might need more sophistication for real-world scenarios.
	constraint, err := semver.NewConstraint(composer.Require.Php)
	if err != nil {
		return "", fmt.Errorf("invalid PHP version constraint in composer.json: %w", err)
	}

	// Define a list of common PHP versions to check against
	// In a real scenario, this might come from a configuration or a more dynamic source
	phpVersions := []string{"7.4.0", "8.0.0", "8.1.0", "8.2.0", "8.3.0"}
	
	for _, v := range phpVersions {
		version, err := semver.NewVersion(v)
		if err != nil {
			continue // Skip invalid versions in our list
		}
		if constraint.Check(version) {
			// Return major.minor version
			return fmt.Sprintf("%d.%d", version.Major(), version.Minor()), nil
		}
	}

	return "", fmt.Errorf("no compatible PHP version found for constraint %s", composer.Require.Php)
}