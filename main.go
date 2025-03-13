package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/charmbracelet/huh"
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
	commitType  string
	title       string
	description string
}

func NewCommit(commitType string, title, description string) Commit {
	return Commit{
		commitType:  commitType,
		title:       title,
		description: description,
	}
}

func isGitInstalled() bool {
	cmd := exec.Command("git", "-v")
	_, err := cmd.Output()

	return err == nil
}

func GitCommit(commit Commit) {
	title := fmt.Sprintf("%s: %s", commit.commitType, commit.title)
	// log.Println(title)
	// return ""

	var cmd *exec.Cmd

	if commit.description == "" {
		cmd = exec.Command("git", "commit", "-m", title)
	} else {

		cmd = exec.Command("git", "-m", title, "-m", commit.description)
	}

	log.Println(cmd.String())

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(out))
}

func main() {
	if !isGitInstalled() {
		fmt.Println(
			`git is not installed on your computer 
go to https://git-scm.com/book/en/v2/Getting-Started-Installing-Git`)
		return
	}

	commit := Commit{}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Commit type").
				Options(
					huh.NewOption(FEAT, "feat"),
					huh.NewOption(FIX, "fix"),
					huh.NewOption(CHORE, "chore"),
					huh.NewOption(REFACTOR, "refactor"),
					huh.NewOption(DOCS, "docs"),
					huh.NewOption(STYLE, "style"),
					huh.NewOption(TEST, "test"),
					huh.NewOption(PERF, "perf"),
					huh.NewOption(CI, "ci"),
					huh.NewOption(BUILD, "build"),
					huh.NewOption(REVERT, "revert"),
				).
				Value(&commit.commitType),
		),

		huh.NewGroup(
			huh.NewInput().
				Title("commit title").
				Value(&commit.title),

			huh.NewText().Title("commit description").Value(&commit.description),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	GitCommit(commit)
}
