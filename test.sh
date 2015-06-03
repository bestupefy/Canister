#!/bin/bash -e

source ./build

echo "Running tests..."
go test -timeout 60s ${COVER} $@ ${TEST} --race

echo "Checking gofmt..."
fmtRes=$(gofmt -l $FMT)
if [ -n "${fmtRes}" ]; then
	echo -e "gofmt checking failed:\n${fmtRes}"
	exit 255
fi

echo "Success"
