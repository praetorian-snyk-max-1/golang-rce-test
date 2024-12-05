#!/bin/bash

# First, generate the code
echo "Generating code..."
go generate ./...

# Then build
echo "Building project..."
go build .

# Run the resulting binary
echo "Running binary..."
./buildtest
