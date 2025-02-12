# Single-Board Computer (SBC) Bakery

A [Packer](https://www.packer.io/) plugin for baking ARM-based Operating System (OS) images for Single-Board
Compute (SBC) modules, such as [Raspberry Pi](https://www.raspberrypi.com/).

This project is mostly based off of [packer-builder-arm](https://github.com/mkaczanowski/packer-builder-arm). The
primary motivation for this project was to:

1.  Learn [Go](https://go.dev/).
2.  Learn how to write Packer plugins (following Hashicorp's fairly opinionated naming and organizational requirements.
    Read more [here](https://developer.hashicorp.com/packer/docs/plugins/creation)).
3.  Provide myself the freedom and flexibility to bring my own development workflow and separate the functional plugin
    code (this repo) from the Packer configuration that will implement it.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing
purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

*   [Docker](https://www.docker.com/) - Build, share, run, and verify applications anywhere.
*   [Devcontainer](https://containers.dev/) - Full-featured containerized development environment specification.

### Installing

1.  Open your favorite IDE that utilizes devcontainer (tested with [VS Code](https://code.visualstudio.com/)).

That's it.

For more information about developing code inside a container, please refer to
[VS Code's excellent documentation](https://code.visualstudio.com/docs/devcontainers/containers).

### Example Implementation

An example containing the necessary packer configuration (written in [HCL](https://github.com/hashicorp/hcl)) can be
found in `example/` at the root of this project.

The example can be run directly inside the devcontainer.

1.  `cd example`
2.  `packer init .`
3.  `packer build .`

## Testing

TBD

## Code Style

TBD

## Deployment

This project uses Github Workflows to build and release new versions of this plugin. To cut a release, simply:

1.  Create a git tag (either on the command-line or on Github) named like `v1.2.3` (Follow
    [SemVer](https://semver.org/)).

This will initiate the release workflow, which utilizes [goreleaser](https://goreleaser.com/) to manage the creation of
binaries, checksums, etc. For more detail, have a look at `.github/workflows/release.yml` and `.goreleaser.yaml` in this
project.

## Version Management

The version of Go is set and managed in the `Makefile`. To update the Go runtime version used in the
[Dev Container](https://containers.dev/), go.mod file, and anywhere else it may be set, simply set the `GO_VERSION`
variable in Makefile to the desired version, then run:

```sh
make set-go-version
```

## Acknowledgements

Hats off to @mkaczanowski for their great work on [packer-builder-arm](https://github.com/mkaczanowski/packer-builder-arm).
