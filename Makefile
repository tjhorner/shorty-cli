.PHONY: dist dist-win dist-macos dist-linux ensure-dist-dir build install uninstall

GOBUILD=go build -ldflags="-s -w"
INSTALLPATH=/usr/local/bin

ensure-dist-dir:
	@- mkdir -p dist

dist-win: ensure-dist-dir
	# Build for Windows x64
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o dist/shorty-windows-amd64.exe *.go

dist-macos: ensure-dist-dir
	# Build for macOS x64
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o dist/shorty-darwin-amd64 *.go

dist-linux: ensure-dist-dir
	# Build for Linux x64
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o dist/shorty-linux-amd64 *.go

dist: dist-win dist-macos dist-linux

build:
	@- mkdir -p bin
	$(GOBUILD) -o bin/shorty *.go
	@- chmod +x bin/shorty

install: build
	mv bin/shorty $(INSTALLPATH)/shorty
	@- rm -rf bin
	@echo "shorty was installed to $(INSTALLPATH)/shorty. Run make uninstall to get rid of it, or just remove the binary yourself."

uninstall:
	rm $(INSTALLPATH)/shorty

run:
	@- go run *.go