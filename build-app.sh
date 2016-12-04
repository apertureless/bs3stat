#!/bin/bash
set -e

echo "Building application"
go build --ldflags '-extldflags "-static"' -o bs3stat

