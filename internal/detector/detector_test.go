package detector

import (
"os"
"path/filepath"
"testing"
)

func TestGetProjectProfile_PHPLaravel(t *testing.T) {
    // Define our test cases
    testCases := []struct {
        name                string // Name of the test
        composerContent     string // The content of the dummy composer.json
        expectedVersion     string // The PHP version we expect the detector to choose
        expectError         bool   // Whether we expect an error
    }{
        {
            name: "Simple Laravel Project with PHP 8.2",
            composerContent: `{
                "require": {
                    "php": "^8.2"
                }
            }`,
            expectedVersion: "8.2",
            expectError: false,
        },
        {
            name: "Laravel Project with a Range Constraint",
            composerContent: `{
                "require": {
                    "php": ">=8.1 <8.4"
                }
            }`,
            expectedVersion: "8.2", // Should pick a stable version within the range
            expectError: false,
        },
        {
            name: "Laravel Project with No PHP Version",
            composerContent: `{
                "require": {
                    "laravel/framework": "^9.0"
                }
            }`,
            expectedVersion: "8.2", // Should return our safe default
            expectError: false,
        },
        {
            name: "Malformed composer.json",
            composerContent: `{ "require": { "php": `, // Intentionally broken JSON
            expectedVersion: "8.2", // Should still detect it's PHP and return the default
            expectError: false, // Our current parser is simple and will fall back
        },
    }

    // --- Run the test loop ---
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            // Create a temporary directory for the test
            tempDir, err := os.MkdirTemp("", "orchestrator-test-*")
            if err != nil {
                t.Fatalf("Failed to create temp dir: %v", err)
            }
            defer os.RemoveAll(tempDir) // Clean up the directory after the test

            // Create the dummy fingerprint files
            os.WriteFile(filepath.Join(tempDir, "composer.json"), []byte(tc.composerContent), 0644)
            os.WriteFile(filepath.Join(tempDir, "artisan"), []byte(""), 0644)
            os.Mkdir(filepath.Join(tempDir, "app"), 0755)

            // Run the function we want to test
            profile, err := GetProjectProfile(tempDir)

            if tc.expectError {
                if err == nil {
                    t.Errorf("Expected an error, but got none")
                }
            } else {
                if err != nil {
                    t.Errorf("Did not expect an error, but got: %v", err)
                }
                if profile == nil {
                    t.Fatalf("Expected a profile, but got nil")
                }
                if profile.Archetype != ArchetypePHPLaravel {
                    t.Errorf("Expected archetype %s, but got %s", ArchetypePHPLaravel, profile.Archetype)
                }
                if profile.LanguageVersion != tc.expectedVersion {
                    t.Errorf("Expected language version %s, but got %s", tc.expectedVersion, profile.LanguageVersion)
                }
            }
        })
    }
}

func TestGetProjectProfile_JavaSpringBoot(t *testing.T) {
	// Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("", "orchestrator-test-java-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up the directory after the test

	// Create dummy fingerprint files for Java Spring Boot
	os.WriteFile(filepath.Join(tempDir, "pom.xml"), []byte("<project></project>"), 0644)
	os.MkdirAll(filepath.Join(tempDir, "src", "main", "java"), 0755)

	// Run the function we want to test
	profile, err := GetProjectProfile(tempDir)

	if err != nil {
		t.Errorf("Did not expect an error, but got: %v", err)
	}
	if profile == nil {
		t.Fatalf("Expected a profile, but got nil")
	}
	if profile.Archetype != ArchetypeJavaSpringBoot {
		t.Errorf("Expected archetype %s, but got %s", ArchetypeJavaSpringBoot, profile.Archetype)
	}
	if profile.LanguageVersion != "17" {
		t.Errorf("Expected language version %s, but got %s", "17", profile.LanguageVersion)
	}
}

func TestGetProjectProfile_PythonFastAPI(t *testing.T) {
	// Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("", "orchestrator-test-python-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up the directory after the test

	// Create dummy fingerprint files for Python FastAPI
	os.WriteFile(filepath.Join(tempDir, "requirements.txt"), []byte("fastapi\n"), 0644)

	// Run the function we want to test
	profile, err := GetProjectProfile(tempDir)

	if err != nil {
		t.Errorf("Did not expect an error, but got: %v", err)
	}
	if profile == nil {
		t.Fatalf("Expected a profile, but got nil")
	}
	if profile.Archetype != ArchetypePythonFastAPI {
		t.Errorf("Expected archetype %s, but got %s", ArchetypePythonFastAPI, profile.Archetype)
	}
	if profile.LanguageVersion != "3.9" {
		t.Errorf("Expected language version %s, but got %s", "3.9", profile.LanguageVersion)
	}
}

func TestGetProjectProfile_NodeJSNextJS(t *testing.T) {
	// Create a temporary directory for the test
	tempDir, err := os.MkdirTemp("", "orchestrator-test-nodejs-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up the directory after the test

	// Create dummy fingerprint files for NodeJS NextJS
	os.WriteFile(filepath.Join(tempDir, "package.json"), []byte(`{"dependencies": {"next": "^12.0.0"}}`), 0644)

	// Run the function we want to test
	profile, err := GetProjectProfile(tempDir)

	if err != nil {
		t.Errorf("Did not expect an error, but got: %v", err)
	}
	if profile == nil {
		t.Fatalf("Expected a profile, but got nil")
	}
	if profile.Archetype != ArchetypeNodeJSNextJS {
		t.Errorf("Expected archetype %s, but got %s", ArchetypeNodeJSNextJS, profile.Archetype)
	}
	if profile.LanguageVersion != "18" {
		t.Errorf("Expected language version %s, but got %s", "18", profile.LanguageVersion)
	}
}