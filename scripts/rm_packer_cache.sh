#!/bin/bash

set -euo pipefail

# Remove the packer cache
if [ -d .packer_cache ]; then
  rm -rf .packer_cache
fi
