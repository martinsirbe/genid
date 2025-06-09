PROJECT_NAME := genid

.PHONY: build
build:
	 go build -o genid cmd/genid/main.go

.PHONY: release-local
release-local:
	 goreleaser  release --snapshot --skip-publish --clean --rm-dist
