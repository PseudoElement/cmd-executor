package main

import (
	"log"
	"os/exec"
)

func removePermissionTrackGlobally() {
	git := executables["git"]

	cmd := exec.Command(git, "config", "core.filemode", "false")
	out, err := cmd.Output()
	if err != nil {
		logRed("core.filemode change failed.")
		logRed("Output: " + string(out))
		panic(err)
	}

	logGreen("Variable core.filemode set to false.")
}

func givePermission(path string) {
	cmd := exec.Command("sudo", "chmod", "-R", "777", path)

	out, err := cmd.Output()
	assert(err == nil, "chmod failed", err)

	logGreen("Permission changed successfully.")
	if len(out) > 0 {
		log.Println("Output: ", string(out))
	}
}
