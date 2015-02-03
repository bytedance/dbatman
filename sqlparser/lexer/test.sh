#!/bin/bash

go test -coverprofile=/home/work/cover/coverage.out
cd /home/work/cover/ && go tool cover -html=coverage.out -o coverage.html && cp coverage.html /mnt/hgfs/ubuntu
cd -
