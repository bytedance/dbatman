#!/usr/bin/env bash

set -e
echo "" > coverage.txt

for d in $(find ./* -maxdepth 3 -type d); do
    if ls $d/*.go &> /dev/null; then
        IMPORT_LIST=`go list -f "{{.Imports}}" $d | sed -e "s/\[//g" | sed -e "s/\]//g"`
        BYTEDANCE_PKG=''
		for i in ${IMPORT_LIST[@]}; do
            if [[ ${i} == github\.com\/bytedance\/dbatman* ]]; then
				BYTEDANCE_PKG=${BYTEDANCE_PKG},${i}
            fi
        done
		
		BYTEDANCE_PKG=`echo ${BYTEDANCE_PKG} | sed -e "s/^,//g"`

		if [[ ${BYTEDANCE_PKG} == '' ]]; then
			go test -coverprofile=profile.out -covermode=atomic $d
		else
        		go test -coverprofile=profile.out -covermode=atomic $d -coverpkg=${BYTEDANCE_PKG}
        	fi
			
		if [ -f profile.out ]; then
            cat profile.out >> coverage.txt
            rm profile.out
        fi
    fi
done
