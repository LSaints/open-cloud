package shell

import (
	"bytes"
	"fmt"
	"os/exec"
)

type ShellExec struct {
}

func (shellExec ShellExec) ExecuteCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return stderr.String(), fmt.Errorf("failed to execute command: %w, stderr: %s", err, stderr.String())
	}

	output := stdout.String()

	return output, nil
}
