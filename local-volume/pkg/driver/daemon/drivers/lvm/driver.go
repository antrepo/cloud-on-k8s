package lvm

import (
	"fmt"
	"github.com/elastic/stack-operators/local-volume/pkg/driver/daemon/cmdutil"
)

// DriverKind in LVM
const DriverKind = "LVM"

// Default driver options
const (
	DefaultVolumeGroup    = "elastic-local-vg"
	DefaultUseThinVolumes = false
	DefaultThinPoolName   = "elastic-local-thinpool"
)

// Driver handles LVM mounts
type Driver struct {
	options Options
}

// Info returns some information about the driver
func (d *Driver) Info() string {
	return fmt.Sprintf("LVM Driver: %+v", d)
}

// Options defines parameters for the LVM driver
type Options struct {
	ExecutableFactory cmdutil.ExecutableFactory
	VolumeGroupName   string
	UseThinVolumes    bool
	ThinPoolName      string
}

// NewDriver creates a new lvm.Driver with the given options
func NewDriver(options Options) *Driver {
	return &Driver{options: options}
}