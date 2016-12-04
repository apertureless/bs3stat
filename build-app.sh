#!/bin/bash
set -e

echo "Building application"
godep CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bs3stat .

