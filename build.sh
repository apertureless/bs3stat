#!/bin/bash
set -e

echo "Building frontend assets"
cd web
npm install
npm run build
cd ..
