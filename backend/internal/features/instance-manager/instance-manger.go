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
