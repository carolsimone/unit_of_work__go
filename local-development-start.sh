#!/bin/bash
set -e

# Run your build and bring up services
make build && make up
