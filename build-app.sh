#!/bin/bash

set -e

go generate
go build -o bs3stat

