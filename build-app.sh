#!/bin/bash
set -e

echo "Building application"
godep go build -o bs3stat

