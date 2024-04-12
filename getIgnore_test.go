package main

import (
	"testing"
)

func TestSyScommand(t *testing.T) {
	success, err := executeSystemCommand("https://raw.githubusercontent.com/github/gitignore/main/Go.gitignore")
	if err != nil {
		t.Fatal(err)
	}

	if success != true {
		t.Fatal("Expected true, got ", success)
	}
}

func TestDownloadIgnore(t *testing.T) {
	success, err := DownloadIgnore("Go")
	if err != nil {
		t.Fatal(err)
	}

	if success != true {
		t.Fatal("Expected true, got ", success)
	}
}
