.PHONY: all
all:
	clean test

.PHONY: clean
clean:
	rm -rf bin

.PHONY: test
test:
	go test ./...