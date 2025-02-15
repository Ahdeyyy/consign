package main

import (
	"fmt"
	"os/exec"
)

type CommitType string

const (
	FEAT     = "feat"
	FIX      = "fix"
	CHORE    = "chore"
	REFACTOR = "refactor"
	DOCS     = "docs"
	STYLE    = "style"
	TEST     = "test"
	PERF     = "perf"
	CI       = "ci"
	BUILD    = "build"
	REVERT   = "revert"
)

type Commit struct {
	commitType  CommitType
	title       string
	description string
}

func isGitInstalled() bool {
	cmd := exec.Command("git", "-v")
	_, err := cmd.Output()
	if err != nil {
		return false
	}
	return true
}

func main() {
	if isGitInstalled() {
		fmt.Println(
			`git is not installed on your computer 
go to https://git-scm.com/book/en/v2/Getting-Started-Installing-Git`)
	}
	fmt.Println("Consign is a git commit message helper")

}
