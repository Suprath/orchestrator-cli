// FILE: cmd/init.go
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Suprath/orchestrator-cli/internal/detector"
	"github.com/Suprath/orchestrator-cli/internal/generator"
	"github.com/Suprath/orchestrator-cli/internal/github" // <-- NEW IMPORT
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new project with a production-ready CI/CD architecture.",
	Run: func(cmd *cobra.Command, args []string) {
		// --- AUTHENTICATION CHECK ---
		if err := github.CheckAuthStatus(); err != nil {
			fmt.Printf("❌ %v\n", err)
			os.Exit(1)
		}

		fmt.Println(" Scanning current directory for project type...")
		currentDir, _ := os.Getwd()
		archetype := detector.DetectArchetype(currentDir)
		if archetype == detector.ArchetypeUnknown {
			fmt.Println("❌ Could not determine project type.")
			os.Exit(1)
		}
		fmt.Printf("✅ Detected a %s project.\n", archetype)

		reader := bufio.NewReader(os.Stdin)
		fmt.Print(" Enter a short, lowercase name for your application (e.g., 'my-api'): ")
		appName, _ := reader.ReadString('\n')
		appName = strings.TrimSpace(appName)
		if appName == "" {
			fmt.Println("❌ App name cannot be empty.")
			os.Exit(1)
		}

		data := generator.TemplateData{
			AppName: appName,
		}

		fmt.Println("\n Generating architectural files...")

		// Expanded list of files to generate
		filesToGenerate := []struct {
			TemplatePath string
			OutputPath   string
			IsCommon     bool
		}{
			{TemplatePath: "common/docker-compose.yml.tmpl", OutputPath: "docker-compose.yml", IsCommon: true},
			{TemplatePath: "common/terraform/eks_fargate.tf.tmpl", OutputPath: "terraform/main.tf", IsCommon: true},
			{TemplatePath: "common/kubernetes/deployment.yml.tmpl", OutputPath: "kubernetes/deployment.yml", IsCommon: true},
			{TemplatePath: "Dockerfile.tmpl", OutputPath: "Dockerfile", IsCommon: false},
			{TemplatePath: "pipeline.yml.tmpl", OutputPath: ".github/workflows/pipeline.yml", IsCommon: false},
		}

		for _, file := range filesToGenerate {
			var templatePath string
			if file.IsCommon {
				templatePath = file.TemplatePath
			} else {
				templatePath = filepath.Join(string(archetype), file.TemplatePath)
			}

			outputDir := filepath.Dir(file.OutputPath)
			if outputDir != "." {
				_ = os.MkdirAll(outputDir, os.ModePerm)
			}

			if err := generator.GenerateFile(templatePath, file.OutputPath, data); err != nil {
				fmt.Printf("❌ Error generating file %s: %v\n", file.OutputPath, err)
				os.Exit(1)
			}
			fmt.Printf("   ✅ Successfully generated %s\n", file.OutputPath)
		}

		// --- GITHUB API INTERACTION ---
		fmt.Print("\n Do you want to apply branch protection rules to this repository on GitHub? (y/n): ")
		applyProtection, _ := reader.ReadString('\n')
		if strings.TrimSpace(strings.ToLower(applyProtection)) == "y" {
			fmt.Print("   Enter the GitHub repository name (e.g., YourUser/YourRepo): ")
			repoName, _ := reader.ReadString('\n')
			repoName = strings.TrimSpace(repoName)

			if repoName != "" {
				// We protect 'main' and 'develop' branches by default
				github.SetBranchProtection(repoName, "main")
				github.SetBranchProtection(repoName, "develop")
			} else {
				fmt.Println("   Skipping branch protection, no repository name provided.")
			}
		}

		fmt.Println("\n Setup complete! Please review the generated files and commit them to your repository.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
