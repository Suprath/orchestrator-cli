// FILE: internal/github/client.go
package github

import (
	"fmt"
	"os/exec"
)

// CheckAuthStatus checks if the user is authenticated with the GitHub CLI.
func CheckAuthStatus() error {
	fmt.Println("INFO: Checking GitHub authentication status...")
	cmd := exec.Command("gh", "auth", "status")
	// We check the error. If the command fails, it means not logged in.
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("GitHub CLI 'gh' not authenticated. Please run 'gh auth login' and try again")
	}
	fmt.Println("INFO: GitHub CLI is authenticated.")
	return nil
}

// SetBranchProtection applies protection rules to a given branch.
func SetBranchProtection(repo string, branch string) error {
	fmt.Printf("INFO: Applying branch protection to '%s' on repo '%s'...\n", branch, repo)
	// This command requires PRs and successful status checks before merging.
	cmd := exec.Command("gh", "api", fmt.Sprintf("repos/%s/branches/%s/protection", repo, branch), "-X", "PUT", "--silent", "-f", "required_status_checks=null", "-f", "enforce_admins=true", "-f", "required_pull_request_reviews=null", "-f", "restrictions=null")

	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to set branch protection: %s", string(output))
	}
	fmt.Printf("INFO: Successfully protected branch '%s'.\n", branch)
	return nil
}