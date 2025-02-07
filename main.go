package main

import "fmt"

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

func main() {

	fmt.Println("Consign is a git commit helper")
}
