//go:generate packer-sdc mapstructure-to-hcl2 -type QemuConfig
package config

import (
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
)

type QemuConfig struct {
	QemuBinarySourcePath      string `mapstructure:"qemu_binary_source_path" required:"true"`
	QemuBinaryDestinationPath string `mapstructure:"qemu_binary_destination_path"`
}

// Prepare qemu configuration.
func (q *QemuConfig) Prepare(_ *interpolate.Context) (warnings []string, errs []error) {
	// Those paths are usually the same.
	if q.QemuBinaryDestinationPath == "" {
		q.QemuBinaryDestinationPath = q.QemuBinarySourcePath
	}

	return warnings, errs
}
