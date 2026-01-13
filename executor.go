package main

import (
	"fmt"
	"log"
	"os/exec"
)

func execute(pathToApp, command string, args []string) error {
	switch command {
	case "npm_i":
		return tryCommand(pathToApp, func() ([]byte, error) {
			return npmInstall(pathToApp, args[0])
		})
	case "yarn_add":
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
	case "git_push":
		return tryCommand(pathToApp, func() ([]byte, error) {
			return gitPush(pathToApp)
		})
	case "git_commit":
		return tryCommand(pathToApp, func() ([]byte, error) {
			return gitCommit(pathToApp, args[0])
		})
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
		logRed("[tryCommand_err] err: " + err.Error())
		logRed("[tryCommand_out] out: " + string(out))

		logBlue("Check invalid permission.")

		removePermissionTrackGlobally()
		givePermission(pathToApp)

		out, err = commandCall()
		if err == nil {
			logGreen("Execution succeeded.")
			log.Println("Output: ", string(out))

			return nil
		}

		logRed("Execution failed.")
		log.Println("Output: ", string(out))
		log.Fatal(err)

		return err
	}

	logGreen("Execution succeeded.")
	log.Println("Output: ", string(out))

	return nil
}

func npmBuild(pathToApp string) ([]byte, error) {
	npm := executables["npm"]
	cmd := exec.Command(npm, "run", "build", "--prefix", pathToApp)
	logBlue(fmt.Sprintf("Executable command: %s\n", cmd.String()))

	return cmd.Output()
}

func npmInstall(pathToApp, packageName string) ([]byte, error) {
	npm := executables["npm"]

	if packageName == "" {
		cmd := exec.Command("sudo", npm, "install", "--prefix", pathToApp, "--legacy-peer-deps")
		logBlue(fmt.Sprintf("Executable command: %s\n", cmd.String()))

		return cmd.Output()
	}

	cmd := exec.Command("sudo", npm, "install", packageName, "--prefix", pathToApp, "--legacy-peer-deps")
	logBlue(fmt.Sprintf("Executable command: %s\n", cmd.String()))

	return cmd.Output()
}

func yarnInstall(pathToApp, packageName string) ([]byte, error) {
	yarn := executables["yarn"]

	if packageName == "" {
		cmd := exec.Command("sudo", yarn, "--cwd", pathToApp)
		logBlue(fmt.Sprintf("Executable command: %s\n", cmd.String()))

		return cmd.Output()
	}

	cmd := exec.Command("sudo", yarn, "add", packageName, "--cwd", pathToApp)
	logBlue(fmt.Sprintf("Executable command: %s\n", cmd.String()))

	return cmd.Output()
}

func gitPull(pathToApp string) ([]byte, error) {
	git := executables["git"]

	gitDir := "--git-dir=" + pathToApp + "/.git"
	gitWorkTree := "--work-tree=" + pathToApp
	cmd := exec.Command(git, gitDir, gitWorkTree, "pull")
	logBlue(fmt.Sprintf("Executable command: %s\n", cmd.String()))

	return cmd.Output()
}

func gitPush(pathToApp string) ([]byte, error) {
	git := executables["git"]

	gitDir := "--git-dir=" + pathToApp + "/.git"
	gitWorkTree := "--work-tree=" + pathToApp
	cmd := exec.Command(git, gitDir, gitWorkTree, "push")
	logBlue(fmt.Sprintf("Executable command: %s\n", cmd.String()))

	return cmd.Output()
}

func gitCommit(pathToApp, commitMsg string) ([]byte, error) {
	git := executables["git"]

	gitDir := "--git-dir=" + pathToApp + "/.git"
	gitWorkTree := "--work-tree=" + pathToApp
	cmd := exec.Command(git, gitDir, gitWorkTree, "add", pathToApp)
	logBlue(fmt.Sprintf("Executable command: %s\n", cmd.String()))

	out, err := cmd.Output()
	if err != nil {
		return out, err
	}

	cmd = exec.Command(git, gitDir, "commit", "-m", commitMsg)
	logBlue(fmt.Sprintf("Executable command: %s\n", cmd.String()))

	return cmd.Output()
}

func gitStashPush(pathToApp string, stashMsg string) ([]byte, error) {
	git := executables["git"]

	gitDir := "--git-dir=" + pathToApp + "/.git"
	gitWorkTree := "--work-tree=" + pathToApp
	cmd := exec.Command(git, gitDir, gitWorkTree, "stash", "push", "-u", "-m", stashMsg)
	logBlue(fmt.Sprintf("Executable command: %s\n", cmd.String()))

	return cmd.Output()
}
