#!/bin/bash

set -eu -o pipefail

_log() {
  echo "[$(date '+%s')] $*"
}

rm -f test.go test-text2go text2go

_log "make sure we're in test directory"
cd "$(git rev-parse --show-toplevel)/test"

input=../generate.go

_log "building text2go binary"
go build -o text2go ..

_log "generating test.go"
./text2go -input="$input" -output="test.go" -prefix=Test

_log "building test binary with generated data included"
go build -o test-text2go

_log "comparing source data with embedded data from test binary"
src="$(cat "$input")"
gen="$(./test-text2go)"

if ! [ "$src" = "$gen" ] ; then
  _log "BAD! input not equal to generated output! diff follows:"
  echo
  diff -u <(echo "$src") <(echo "$gen")
else
  _log "GOOD! input equal to generated output"
fi
