package instance

import "errors"

type Instance struct {
	ID        uint64 `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	RAM       uint64 `json:"ram,omitempty"`
	Disk      string `json:"disk,omitempty"`
	Vcpus     uint64 `json:"vcpus,omitempty"`
	OsVariant string `json:"osvariant,omitempty"`
	Console   string `json:"console,omitempty"`
	Location  string `json:"location,omitempty"`
	ExtraArgs string `json:"extraargs,omitempty"`
}

func (instance Instance) Init() error {
	if err := instance.validate(); err != nil {
		return err
	}
	return nil
}

func (instance *Instance) validate() error {
	if instance.Name == "" {
		return errors.New("name field cannot be empty")
	}
	if instance.RAM < 1024 {
		return errors.New("RAM field must be at least 1024 MB")
	}
	if instance.Disk == "" {
		return errors.New("disk field cannot be empty")
	}
	if instance.Vcpus <= 0 {
		return errors.New("vcpus field cannot be 0")
	}
	if instance.OsVariant == "" {
		return errors.New("OSvariant field cannot be empty")
	}
	if instance.Console == "" {
		return errors.New("console field cannot be empty")
	}
	if instance.Location == "" {
		return errors.New("location field cannot be empty")
	}

	return nil
}
