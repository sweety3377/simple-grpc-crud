#!/usr/bin/env bash

OUT_DIR="../pb"
GRPC_OUT_DIR="../pb"
PROTO_DIR="../proto/*.proto"

protoc -I . --go_out=${OUT_DIR} --go-grpc_out=${GRPC_OUT_DIR} "${PROTO_DIR}"