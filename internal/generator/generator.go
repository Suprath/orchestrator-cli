// FILE: internal/generator/generator.go
package generator
import (
    "os"
    "text/template"
    "github.com/Suprath/orchestrator-cli/internal/templates"
)

// Data struct holds the user's answers
type TemplateData struct {
    AppName string
    // ... more fields later
}

func GenerateFile(templatePath string, outputPath string, data TemplateData) error {
    // Read the template from the embedded filesystem
    tmpl, err := template.ParseFS(templates.TemplateFS, templatePath)
    if err != nil {
        return err
    }

    // Create the output file
    outputFile, err := os.Create(outputPath)
    if err != nil {
        return err
    }
    defer outputFile.Close()

    // Execute the template, writing the result to the file
    return tmpl.Execute(outputFile, data)
}