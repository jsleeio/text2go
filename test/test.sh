#!/bin/bash

set -eu -o pipefail

rm -f test.go test-text2go text2go

cd "$(git rev-parse --show-toplevel)/test"

input=../generate.go
go build -o text2go ..
./text2go -input="$input" -output="test.go" -prefix=Test
go build -o test-text2go

src="$(cat "$input")"
gen="$(./test-text2go)"

if ! [ "$src" = "$gen" ] ; then
  echo "input not equal to generated output! diff follows:"
  echo
  diff -u <(echo "$src") <(echo "$gen")
fi
