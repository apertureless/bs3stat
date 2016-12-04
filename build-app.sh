#!/bin/bash
set -e

echo "Building application"
#godep go build -o bs3stat .
go build --ldflags '-extldflags "-static"' -o bs3stat

