package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	pathToApp := loadPath()
	loadExeFiles()

	command := askCommand()
	args := askArguments(command)

	logBlue(fmt.Sprintf("arguments: %v\n", args))

	execute(pathToApp, command, args)
}

func askCommand() string {
	logBlue("Input command name(npm_i, npm_b, yarn_add, git_pull, git_push, git_commit, git_stash_push):")

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

	re := regexp.MustCompile("\n")
	replaced := re.ReplaceAllString(line, "")

	if replaced == "" && arg.Required {
		logRed("Argument required.")
		return _askArgument(arg)
	}

	return replaced
}
