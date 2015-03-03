all: build run 

build:
	go install ./...

run:
	cd cmd/proxy && go build && cd -
	GOGC=1000 ./cmd/proxy/proxy --config=etc/proxy_single.yaml --logfile=log/proxy.log --loglevel=0 

clean:
	go clean -i ./...

test:
	go test ./...

package: build
	tar cvf output.tar etc/ cmd/proxy/proxy run.sh 
