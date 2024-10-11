#!/usr/bin/env just --justfile

update:
  go get -u
  go mod tidy -v

test:
    go test ./... -count=1