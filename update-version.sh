#!/bin/bash

version=$(golangci-lint version | grep -oP 'version \Kv[0-9]+\.[0-9]+\.[0-9]+')

yaml_file=".custom-gcl.yaml"

if grep -qE "^version:" "$yaml_file"; then
  # Update existing version line
  sed -i -E "s/^version:.*$/version: $version/" "$yaml_file"
else
  # Insert version at the top of the file
  sed -i "1s/^/version: $version\n/" "$yaml_file"
fi
