#!/bin/bash
set -e

echo "Building application"
CGO_ENABLED=0 GOOS=linux godep go build -a -installsuffix cgo -o bs3stat .

