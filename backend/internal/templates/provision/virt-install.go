package provision

import (
	"fmt"
)

type VirtInstallTemplate struct{}

func (vTemplate VirtInstallTemplate) GenerateTemplate(name string, ram uint64, disk string, vcpus uint64, osVariant string, console string, location string, extraArgs string) string {
	template := fmt.Sprintf(
		`virt-install --name %s \
		--ram %v \
		--disk %s \
		--vcpus %v \
		--os-variant %s \
		--console %s \
		--location '%s' \
		--extra-args '%s' \
		--noreboot`,
		name,
		ram,
		disk,
		vcpus,
		osVariant,
		console,
		location,
		extraArgs,
	)

	return template
}

func (vTemplate VirtInstallTemplate) DeleteInstanceFromTemplate(instanceName string) string {
	template := fmt.Sprintf(`
		virsh destroy %s && \
		virsh undefine %s && \
		rm -rf ./var/lib/opencloud/instances/%s
	`, instanceName, instanceName, instanceName)
	return template
}
