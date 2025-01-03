GO_VERSION := 1.23.4

update-go-version:
	go mod edit -go=$(GO_VERSION)
	sed -i "s/FROM golang:.*/FROM golang:$(GO_VERSION)/" .devcontainer/Dockerfile
	echo "$(GO_VERSION)" > .go-version
