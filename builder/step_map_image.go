package builder

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
)

// Maps the system image to /dev/loopX
type StepMapImage struct {
	ResultKey  string
	loopDevice string
}

// Run.
func (s *StepMapImage) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	config := state.Get("config").(*Config)
	ui := state.Get("ui").(packer.Ui)
	image := config.ImageConfig.ImagePath

	// Ask losetup to find empty device and map image.
	// Manpage: https://man7.org/linux/man-pages/man8/losetup.8.html
	ui.Message(fmt.Sprintf("mapping image %s to free loopback device", image))

	out, err := exec.Command("losetup", "--find", "--partscan", "--show", image).CombinedOutput()

	if err != nil {
		ui.Error(fmt.Sprintf("error losetup --find --partscan %v: %s", err, string(out)))
		return multistep.ActionHalt
	}
	s.loopDevice = strings.Trim(string(out), "\n")

	state.Put(s.ResultKey, s.loopDevice)
	ui.Message(fmt.Sprintf("image %s mapped to %s", image, s.loopDevice))

	return multistep.ActionContinue
}

// Cleanup.
func (s *StepMapImage) Cleanup(state multistep.StateBag) {
	ui := state.Get("ui").(packer.Ui)

	// Warning: Busy device will prevent detaching loop device from file
	// https://github.com/util-linux/util-linux/issues/484
	out, err := exec.Command("losetup", "--detach", s.loopDevice).CombinedOutput()
	if err != nil {
		ui.Error(fmt.Sprintf("error while unmounting loop device %v: %s", err, string(out)))
	}
}
