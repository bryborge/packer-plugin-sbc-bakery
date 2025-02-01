package builder

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/mholt/archiver/v3"
)

// Creates the filesystem on an already-partitioned image.
type StepExtractAndCopyImage struct {
	FromKey string
}

// Run.
func (s *StepExtractAndCopyImage) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	config := state.Get("config").(*Config)
	archivePath := state.Get(s.FromKey).(string)

	var err error
	var out []byte

	// Step 1: Create temporary directory.
	dir, err := os.MkdirTemp(config.TmpDirLocation, "image")
	if err != nil {
		ui.Error(fmt.Sprintf("error while creating temporary directory %v", err))
		return multistep.ActionHalt
	}
	defer os.RemoveAll(dir)

	// Step 2: Copy archive to temporary directory.
	dst := filepath.Join(dir, filepath.Base(archivePath))
	out, err = exec.Command("cp", "-rf", archivePath, dst).CombinedOutput()
	if err != nil {
		ui.Error(fmt.Sprintf("error while copying file %v: %s", err, out))
		return multistep.ActionHalt
	}

	// Skip unarchive logic if provided a raw image.
	if config.RemoteFileConfig.TargetExtension == "img" || config.RemoteFileConfig.TargetExtension == "iso" {
		ui.Message("using raw image")
	} else {
		// Step 3: Unarchive file within the temporary dir.
		ui.Message(fmt.Sprintf("unpacking %s to %s", archivePath, config.ImageConfig.ImagePath))
		if len(config.RemoteFileConfig.FileUnarchiveCmd) != 0 {
			cmd := make([]string, len(config.RemoteFileConfig.FileUnarchiveCmd))
			vars := map[string]string{
				"$ARCHIVE_PATH": dst,
				"$TMP_DIR":      dir,
			}

			for i, elem := range config.RemoteFileConfig.FileUnarchiveCmd {
				if _, ok := vars[elem]; ok {
					cmd[i] = vars[elem]
				} else {
					cmd[i] = elem
				}
			}

			ui.Message(fmt.Sprintf("unpacking with custom command: %s", cmd))
			out, err = exec.Command(cmd[0], cmd[1:]...).CombinedOutput()
		} else {
			out, err = []byte("N/A"), archiver.Unarchive(archivePath, dir)
		}

		if err != nil {
			ui.Error(fmt.Sprintf("error while unpacking %v: %s", err, out))
			return multistep.ActionHalt
		}

		// Step 4: If previously-copied archive still exists, remove it.
		if _, err := os.Stat(dst); err == nil {
			os.RemoveAll(dst)
		}
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		ui.Error(fmt.Sprintf("error while reading temporary directory %v", err))
		return multistep.ActionHalt
	}

	if len(files) != 1 {
		ui.Error(fmt.Sprintf("only one file is expected to be present after unarchiving, found: %d", len(files)))
		return multistep.ActionHalt
	}

	// Step 5: Move single file to destination path.
	out, err = exec.Command("mv", filepath.Join(dir, files[0].Name()), config.ImageConfig.ImagePath).CombinedOutput()
	if err != nil {
		ui.Error(fmt.Sprintf("error while copying file %v: %s", err, out))
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

// Cleanup.
func (s *StepExtractAndCopyImage) Cleanup(_ multistep.StateBag) {}
