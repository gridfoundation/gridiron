#!/bin/sh

if ! ./scripts/setup_gridiron.sh; then
  exit 1
fi
./scripts/run_gridiron.sh

