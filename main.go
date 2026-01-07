package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

var executables = make(map[string]string, 3)

func loadPath() string {
	err := godotenv.Load()
	assert(err == nil, "godotenv loading failed", err)

	pathToNpmProject := os.Getenv("PATH_TO_NPM_PROJECT")
	assert(pathToNpmProject != "", "empty PATH_TO_NPM_PROJECT", nil)

	return pathToNpmProject
}

func loadExeFiles() {
	exeNames := []string{"npm", "yarn", "git"}
	for _, exeName := range exeNames {
		exeFile, err := exec.LookPath(exeName)
		assert(err == nil, fmt.Sprintf("%s not installed", exeName), err)
		executables[exeName] = exeFile
	}
}

func main() {
	pathToApp := loadPath()
	loadExeFiles()

	command := askCommand()
	args := askArguments(command)

	execute(pathToApp, command, args)
}

func askCommand() string {
	logBlue("Input command name(npm_i, npm_b, yarn, git_pull, git_commit, git_stash_push):")

	var command string
	fmt.Scanf("%s", &command)

	for commandName, _ := range argsConfig {
		if commandName == command {
			return command
		}
	}

	logRed("Invalid command.")

	return askCommand()
}

func askArguments(command string) []string {
	commandArgs := argsConfig[command]
	args := make([]string, 0, len(commandArgs))

	for _, arg := range commandArgs {
		input := _askArgument(arg)
		args = append(args, input)
	}

	return args
}

func _askArgument(arg Argument) string {
	logBlue("Input " + arg.Name + ":")

	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	assert(err == nil, "[_askArgument_in.ReadString] failed: ", err)

	if line == "" && arg.Required {
		logRed("Argument required.")
		return _askArgument(arg)
	}

	return line
}
