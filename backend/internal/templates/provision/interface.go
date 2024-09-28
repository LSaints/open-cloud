package provision

type ProvisionInterface interface {
	GenerateTemplate(name string, ram uint64, disk string, vcpus uint64, osVariant string, console string, location string, extraArgs string) string
	DeleteInstanceFromTemplate(instanceName string) string
}
