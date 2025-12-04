#!/bin/bash

# Validate that go_version is provided
if [ -z "{{cookiecutter.go_version}}" ]; then
    echo "ERROR: go_version is required. Please provide a Go version (e.g., '1.25.4')."
    exit 1
fi

exit 0
