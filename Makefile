default: get test out/example
clean:
	rm -rf out
get:
	go get
test:
	go vet && go test
out/example: implementation.go cmd/example/main.go
	mkdir -p out
	go build -ldflags="-X 'main.buildVersion=`git describe`'" -o out/example ./cmd/example
