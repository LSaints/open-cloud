package disk

import (
	config "backend/configs"
	"fmt"
)

type QemuImg struct {
}

func (qemuImg QemuImg) CreateDisk(name string, size uint64) (string, string) {
	root_path := fmt.Sprintf("%s/%s/disk", config.DiskDir, name)
	command := fmt.Sprintf(
		`mkdir -p %s && qemu-img create -f qcow2 %s/%s.qcow2 %dG`,
		root_path, root_path, name, size,
	)
	dirOutput := fmt.Sprintf("path=%s/%s.qcow2,size=%d", root_path, name, size)

	return command, dirOutput
}
