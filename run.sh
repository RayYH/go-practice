#!/usr/bin/env bash

WORKING_DIR=$(dirname "${BASH_SOURCE[0]}")
WORKING_DIR=$(cd "$WORKING_DIR" && pwd)

function run() {
  [ -d "$WORKING_DIR/$1" ] || return
  cd "$1" && go run .
  cd "$WORKING_DIR" || return
}

run "$@"
