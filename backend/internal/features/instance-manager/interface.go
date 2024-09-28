package instancemanager

type InstanceManagerInterface interface {
	ProvisionInstanceFromTemplate(command string) error
	CreateDiskFromTemplate(command string) error
}
