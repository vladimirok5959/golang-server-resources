VERSION="1.0.0"

default: test

test:
	@go test -v ./...

build: clean
	@-mkdir ./bin
	@cd ./bin
	@cd ..
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o ./bin/file2bytes.linux-amd64 -ldflags='-X main.Version=$(VERSION) -extldflags "-static"' ./tools/file2bytes
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -o ./bin/file2bytes.darwin-amd64 -ldflags='-X main.Version=$(VERSION) -extldflags "-static"' ./tools/file2bytes
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -o ./bin/file2bytes.windows-amd64.exe -ldflags='-X main.Version=$(VERSION) -extldflags "-static"' ./tools/file2bytes
	cd ./bin && find . -name 'file2bytes*' | xargs -I{} tar czf {}.tar.gz {}
	cd ./bin && shasum -a 256 * > sha256sum.txt
	cat ./bin/sha256sum.txt

clean:
	@-rm -r ./bin
