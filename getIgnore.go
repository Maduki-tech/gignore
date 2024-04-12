package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func executeSystemCommand(URL string) (success bool, error error) {
	cmd := exec.Command("curl", "-o", ".gitignore", URL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, nil
}

// TODO: Add validierung that String can only be a valid language
func DownloadIgnore(language string)(success bool, error error){
	URL := fmt.Sprintf("https://raw.githubusercontent.com/github/gitignore/main/%s.gitignore", language)
	success, err := executeSystemCommand(URL)

	if err != nil {
		log.Fatal(err)
		return false, err
	}

	return success, nil
}
