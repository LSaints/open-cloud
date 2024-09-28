package instancemanager

type InstanceManagerInterface interface {
	ProvisionInstanceFromTemplate(command string) error
	HaltInstanceFromTemplate(command string) error
	GetStatusInstanceFromTemplate(command string) (string, error)
	CreateDiskFromTemplate(command string) error
}
