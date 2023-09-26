GOFMT=gofmt
GC=go build

service: service-windows service-linux service-mac

service-windows:
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 $(GC) -o ./bin/service-windows-amd64.exe .

service-linux:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 $(GC) -o ./bin/service-linux-amd64 .

service-mac:
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 $(GC) -o ./bin/service-darwin-amd64 .


clean:
	rm -rf *.8 *.o *.out *.6 *exe coverage
	rm ./bin/log/*.log ./bin/sqlite.db ./bin/service*

