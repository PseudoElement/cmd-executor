package main

import "os/exec"

func tryCommand(cmd *exec.Cmd) error {
	return nil
}

func buildNpm(path string) *exec.Cmd {
	npm, err := exec.LookPath("npm")
	assert(err == nil, "npm not installed")

	return exec.Command(npm, "run", "build", "--prefix", path)
}
