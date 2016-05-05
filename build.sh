#!/bin/bash

if test "`which godep`" -ne "0"; then 
	go get github.com/tools/godep
fi

godep restore

cd cmd/dbatman && go build && cd -

mkdir -p output

cp comd/dbatman/dbatman ./output
cp config/proxy.yml ./output