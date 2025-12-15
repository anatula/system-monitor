# Simple Makefile for learning
.PHONY: test run clean

test:
	@echo "Running tests..."
	go test -v

run:
	go run main.go

clean:
	go clean