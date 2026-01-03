package main

import (
	"os/exec"
)

type MonitorMacOS struct{}

func (m *MonitorMacOS) readLogs() {}

func (m *MonitorMacOS) test() {}

func (m *MonitorMacOS) executeAndMonitor(cmd *exec.Cmd) {

}
