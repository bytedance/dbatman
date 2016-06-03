#!/bin/bash

set -e

which godep || go get github.com/tools/godep

godep restore

cd cmd/dbatman && go build && cd -

mkdir -p output

cp cmd/dbatman/dbatman ./output
cp config/proxy.yml ./output
cp config/test.yml ./output
