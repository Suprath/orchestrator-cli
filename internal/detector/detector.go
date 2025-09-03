// FILE: internal/detector/detector.go
package detector

import (
"os"
"path/filepath"
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

// DetectArchetype scans a directory to identify the project's archetype.
func DetectArchetype(dirPath string) Archetype {
// --- NEW: PHP Laravel Detection ---
// A strong indicator is the presence of 'artisan', 'composer.json', and the 'app' directory.
if fileExists(filepath.Join(dirPath, "artisan")) &&
fileExists(filepath.Join(dirPath, "composer.json")) &&
dirExists(filepath.Join(dirPath, "app")) {
return ArchetypePHPLaravel
}

// --- NEW: Node.js / Next.js Detection ---
// A strong indicator is 'package.json' and 'next.config.js'.
if fileExists(filepath.Join(dirPath, "package.json")) &&
fileExists(filepath.Join(dirPath, "next.config.js")) {
return ArchetypeNodeJSNextJS
}

// --- Existing Detection Logic ---
if fileExists(filepath.Join(dirPath, "pom.xml")) || fileExists(filepath.Join(dirPath, "build.gradle")) {
return ArchetypeJavaSpringBoot
}

if fileExists(filepath.Join(dirPath, "requirements.txt")) {
return ArchetypePythonFastAPI
}

return ArchetypeUnknown
}