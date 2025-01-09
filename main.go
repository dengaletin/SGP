package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	logFile, err := os.OpenFile("merge.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.Println("----- Starting the program -----")

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	repoURL := os.Getenv("REPO_URL")
	branchesEnv := os.Getenv("BRANCHES")
	mainBranchName := os.Getenv("MAIN_BRANCH_NAME")

	if repoURL == "" || branchesEnv == "" {
		log.Fatalf("REPO_URL or BRANCHES not set in .env file")
	}

	branches := strings.Split(branchesEnv, ",")

	for _, branch := range branches {
		branch = strings.TrimSpace(branch)
		if branch == "" {
			continue
		}

		log.Printf("Processing a branch: %s", branch)
		fmt.Printf("Processing a branch: %s\n", branch)

		if err := gitCommand("checkout", branch); err != nil {
			log.Printf("Error when switching to branch %s: %v", branch, err)
			continue
		}

		if err := gitCommand("merge", mainBranchName); err != nil {
			log.Printf("Error merging master into %s: %v", branch, err)
			continue
		}

		if err := gitCommand("push", "origin", branch); err != nil {
			log.Printf("Error when pushing a branch %s: %v", branch, err)
			continue
		}

		log.Printf("Branch %s has been successfully updated and pushed", branch)
		fmt.Printf("Branch %s has been successfully updated and pushed\n", branch)
	}

	log.Println("----- Completing the program -----")
}

func gitCommand(args ...string) error {
	cmd := exec.Command("git", args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("git %s command failed: %v\nOutput: %s", strings.Join(args, " "), err, string(output))
		return err
	}
	log.Printf("git %s command completed successfully.\nOutput: %s", strings.Join(args, " "), string(output))
	return nil
}
