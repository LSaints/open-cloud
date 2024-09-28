package disk

type Disk interface {
	CreateDisk(name string, size uint64) string
}
