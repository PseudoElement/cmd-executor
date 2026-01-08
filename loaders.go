package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
)

var executables = make(map[string]string, 3)

func loadPath() string {
	err := godotenv.Load()
	assert(err == nil, "godotenv loading failed", err)

	pathToNpmProject := checkPathToNpmProject()
	if pathToNpmProject != "" {
		logBlue("PATH_TO_NPM_PROJECT is " + pathToNpmProject)
		return pathToNpmProject
	} else {
		pathToNpmProject := os.Getenv("PATH_TO_NPM_PROJECT")
		assert(pathToNpmProject != "", "empty PATH_TO_NPM_PROJECT", nil)
		logBlue("PATH_TO_NPM_PROJECT is " + pathToNpmProject)
		return pathToNpmProject
	}
}

func loadExeFiles() {
	exeNames := []string{"npm", "yarn", "git"}

	for _, exeName := range exeNames {
		exeFile, err := exec.LookPath(exeName)
		assert(err == nil, fmt.Sprintf("%s not installed", exeName), err)
		executables[exeName] = exeFile
	}
}

func checkPathToNpmProject() string {
	for _, arg := range os.Args {
		pathToApp, found := strings.CutPrefix(arg, "--path=")
		if found {
			return pathToApp
		}
	}

	return ""
}
