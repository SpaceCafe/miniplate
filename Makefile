.PHONY: clean test
.DEFAULT_GOAL = all

all: build/miniplate test

build/miniplate:
	go build -a -buildmode=exe -trimpath -ldflags="-s -w" -o build/ ./cmd/miniplate/
	upx --best build/miniplate

test:
	go test ./... -v -race
	SECRET_FILE=test/secret.txt ./build/miniplate --in test/example.tmpl --context user1=test/testdata.json --context user2=test/testdata.yaml

clean:
	@rm -rf "build"
