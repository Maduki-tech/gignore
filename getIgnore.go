package main

import (
	"fmt"
	"os"
	"os/exec"
)

func executeSystemCommand(URL string) {
	cmd := exec.Command("curl", "-o", ".gitignore", URL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

}

// TODO: Add validierung that String can only be a valid language
func DownloadIgnore(language string) {
	URL := fmt.Sprintf("https://raw.githubusercontent.com/github/gitignore/main/%s.gitignore", language)
	executeSystemCommand(URL)
}

