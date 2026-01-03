package main

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
)

func loadPath() string {
	err := godotenv.Load()
	assert(err == nil, "godotenv loading failed")

	pathToNpmProject := os.Getenv("PATH_TO_NPM_PROJECT")
	assert(pathToNpmProject != "", "empty PATH_TO_NPM_PROJECT")

	return pathToNpmProject
}

func givePermission(path string) error {
	cmd := exec.Command("chmod", "-R", "777", path)

	out, err := cmd.Output()
	assert(err == nil, "chmod failed")

	log.Println("Permission changed successfully.")
	if len(out) > 0 {
		log.Println("Output: ", string(out))
	}

	return nil
}

func tryBuildProject(path string, firstTry bool) error {
	npm, err := exec.LookPath("npm")
	assert(err == nil, "npm not installed")

	cmd := exec.Command(npm, "run", "build", "--prefix", path)

	out, err := cmd.Output()
	if err != nil {
		log.Println("Execution failed.")
		if strings.Contains(err.Error(), "exit status 243") && firstTry {
			givePermission(path)
			tryBuildProject(path, false)
		} else {
			log.Fatal(err)
		}
	}

	log.Println("Execution succeded.")
	if len(out) > 0 {
		log.Println("Output: ", string(out))
	}

	return nil
}

func main() {
	tryBuildProject(loadPath(), true)
	// tryCommand(buildNpm(loadPath()))
}
