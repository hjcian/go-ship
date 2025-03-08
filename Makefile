build:
	@go build -o bin/$(shell basename $(PWD)) *.go

run: build
	@./bin/$(shell basename $(PWD))

tidy:
	@go mod tidy