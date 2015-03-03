#!/bin/sh

GOGC=1000 ./cmd/proxy/proxy --config=etc/proxy_single.yaml --logfile=log/proxy.log --loglevel=0