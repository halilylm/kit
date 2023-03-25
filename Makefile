.PHONY: test test-no-cache

test:
	go test ./...

test-no-cache:
	go test ./... -count=1