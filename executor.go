package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func execute(pathToApp, command string, args ...string) error {
	switch command {
	case "npm_i":
		return tryCommand(pathToApp, func() ([]byte, error) {
			return npmInstall(pathToApp, args[0])
		})
	case "yarn":
		return tryCommand(pathToApp, func() ([]byte, error) {
			return yarnInstall(pathToApp, args[0])
		})
	case "npm_b":
		return tryCommand(pathToApp, func() ([]byte, error) {
			return npmBuild(pathToApp)
		})
	case "git_pull":
		return tryCommand(pathToApp, func() ([]byte, error) {
			return gitPull(pathToApp)
		})
	case "git_commit":
		return tryCommand(pathToApp, func() ([]byte, error) {
			return gitCommit(pathToApp, args[0])
		})
	// @FIX check problem with stash
	case "git_stash_push":
		return tryCommand(pathToApp, func() ([]byte, error) {
			return gitStashPush(pathToApp, args[0])
		})
	}

	return fmt.Errorf("unknown commandName %s", command)
}

func tryCommand(pathToApp string, commandCall func() ([]byte, error)) error {
	out, err := commandCall()
	if err != nil {
		if strings.Contains(err.Error(), "exit status 243") {
			logBlue("Invalid permission.")
			givePermission(pathToApp)

			out, err = commandCall()
			if err == nil {
				logRed("Execution succeded.")
				log.Println("Output: ", string(out))
			}

			return err
		} else {
			logRed("Execution failed.")
			log.Println("Output: ", string(out))
			log.Fatal(err)
		}
	}

	return nil
}

func givePermission(path string) {
	cmd := exec.Command("chmod", "-R", "777", path)

	out, err := cmd.Output()
	assert(err == nil, "chmod failed")

	logGreen("Permission changed successfully.")
	if len(out) > 0 {
		log.Println("Output: ", string(out))
	}
}

func npmBuild(pathToApp string) ([]byte, error) {
	npm := executables["npm"]
	cmd := exec.Command(npm, "run", "build", "--prefix", pathToApp)

	return cmd.Output()
}

func npmInstall(pathToApp, packageName string) ([]byte, error) {
	npm := executables["npm"]

	if packageName == "" {
		cmd := exec.Command(npm)
		return cmd.Output()
	}

	cmd := exec.Command(npm, "install", packageName, "--prefix", pathToApp, "--legacy-peer-deps")

	return cmd.Output()
}

func yarnInstall(pathToApp, packageName string) ([]byte, error) {
	yarn := executables["yarn"]

	if packageName == "" {
		cmd := exec.Command(yarn)
		return cmd.Output()
	}

	cmd := exec.Command(yarn, "add", packageName, "--cwd", pathToApp)

	return cmd.Output()
}

func gitPull(pathToApp string) ([]byte, error) {
	git := executables["git"]

	gitDir := "--git-dir=" + pathToApp + "/.git"
	cmd := exec.Command(git, gitDir, "pull")

	return cmd.Output()
}

func gitCommit(pathToApp, commitMsg string) ([]byte, error) {
	git := executables["git"]

	gitDir := "--git-dir=" + pathToApp + "/.git"
	cmd := exec.Command(git, gitDir, "add", pathToApp)
	out, err := cmd.Output()
	if err != nil {
		return out, err
	}

	// cmd = exec.Command(git, gitDir, "commit", "-m", commitMsg)

	return cmd.Output()
}

func gitStashPush(pathToApp string, stashMsg string) ([]byte, error) {
	git := executables["git"]

	gitDir := "--git-dir=" + pathToApp + "/.git"
	cmd := exec.Command(git, gitDir, "stash", "push", "-u", "-m", stashMsg)

	return cmd.Output()
}
