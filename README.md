# Single-Board Computer (SBC) Bakery

A [Packer](https://www.packer.io/) plugin for baking ARM-based Operating System (OS) images for Single-Board
Compute(SBC) modules.

**Note:** This project is very much a work-in-progress. Please see
[packer-builder-arm](https://github.com/mkaczanowski/packer-builder-arm), which is the original inspiration for this
work.

## Version Management

The version of Go is set and managed in the `Makefile`. To update the Go runtime version in Dev Container Dockerfile,
go.mod, and anywhere else it may be set, simply set the `GO_VERSION` variable in Makefile to the desired version, then
run:

```sh
make set-go-version
```
