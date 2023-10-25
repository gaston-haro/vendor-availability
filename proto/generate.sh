#!/usr/bin/env bash

# Get current directory.
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# Find all directories containing at least one proto file.
# Based on: https://buf.build/docs/migration-prototool#prototool-generate.
for f in $(find ${DIR} -type f -name '*.proto'); do
  # Generate all files with protoc-gen-go.
  #protoc -I ${DIR} --go_out=plugins=grpc,paths=source_relative:${DIR}/gen/go ${files}
  protoc --proto_path=$DIR --go_out=. --go_opt=paths=source_relative:${DIR}/gen --go-grpc_out=. --go-grpc_opt=paths=source_relative:${DIR}/gen $f
#  printf "$f"
done