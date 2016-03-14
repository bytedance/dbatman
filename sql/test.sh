#!/bin/bash

go test -coverprofile=$HOME/cover/coverage.out
cd ~/cover/ && go tool cover -html=coverage.out -o coverage.html
cp coverage.html /mnt/hgfs/ubuntu
cd -
