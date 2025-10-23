#!/usr/bin/env bash

set -e
set -o pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No colour

echo -e "${YELLOW} Running tests for Core service...${NC}"
cd "$(dirname "$0")/.."  # Move to /client directory

if go test ./... -v; then
    echo -e "${GREEN} Tests passed successfully!${NC}"
else
    echo -e "${RED} Tests failed. Aborting generation.${NC}"
    exit 1
fi

echo " " 

echo -e "${YELLOW} Generating gRPC Go files from proto definitions...${NC}"

if ! command -v protoc >/dev/null 2>&1; then
    echo -e "${RED}Error: protoc not found. Please install Protocol Buffers compiler.${NC}"
    exit 1
fi

if ! command -v protoc-gen-go >/dev/null 2>&1; then
    echo -e "${RED}Error: protoc-gen-go not found. Install with:${NC}"
    echo -e "${YELLOW}go install google.golang.org/protobuf/cmd/protoc-gen-go@latest${NC}"
    exit 1
fi

if ! command -v protoc-gen-go-grpc >/dev/null 2>&1; then
    echo -e "${RED}Error: protoc-gen-go-grpc not found. Install with:${NC}"
    echo -e "${YELLOW}go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest${NC}"
    exit 1
fi

protoc \
  --go_out=. \
  --go-grpc_out=. \
  --proto_path=proto \
  proto/*.proto

echo -e "${GREEN} gRPC files generated successfully into internal/api!${NC}"
