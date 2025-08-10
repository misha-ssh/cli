#!/bin/sh
set -e
rm -rf manpages
mkdir manpages
go run ./cmd/misha man | gzip -c -9 >manpages/misha.1.gz