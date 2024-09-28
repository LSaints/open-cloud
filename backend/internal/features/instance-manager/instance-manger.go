package instancemanager

import (
	shellexec "backend/internal/features/shell"
)

type InstanceManager struct{}

func (im InstanceManager) ProvisionInstanceFromTemplate(command string) error {
	shellexec := shellexec.ShellExec{}
	shellexec.ExecuteCommand(command)

	return nil
}

func (im InstanceManager) CreateDiskFromTemplate(command string) error {
	shellexec := shellexec.ShellExec{}
	shellexec.ExecuteCommand(command)

	return nil
}

func (im InstanceManager) DeleteFromTemplate(command string) error {
	shellexec := shellexec.ShellExec{}
	shellexec.ExecuteCommand(command)

	return nil
}

func (im InstanceManager) HaltInstanceFromTemplate(command string) error {
	shellexec := shellexec.ShellExec{}
	shellexec.ExecuteCommand(command)

	return nil
}

func (im InstanceManager) GetStatusInstanceFromTemplate(command string) (string, error) {
	shellexec := shellexec.ShellExec{}
	output, _ := shellexec.ExecuteCommand(command)

	return output, nil
}

func (im InstanceManager) StartInstanceFromTemplate(command string) (string, error) {
	shellexec := shellexec.ShellExec{}
	output, _ := shellexec.ExecuteCommand(command)

	return output, nil
}
