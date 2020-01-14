#!/usr/bin/env bash

protoc --go_out=plugins=grpc:. kvdb.proto
protoc --rust_out=. --grpc_out=. --plugin=protoc-gen-grpc=`which grpc_rust_plugin` kvdb.proto
