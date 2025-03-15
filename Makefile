build:
	@go build -o bin/$(shell basename $(PWD)) *.go

run: build
	@./bin/$(shell basename $(PWD))

tidy:
	@go mod tidy


redis-run:
	docker run -d \
		--name lucid_heisenberg \
		-v ./data:/data \
		-p 6379:6379 \
		redis:7.2.7

stop-redis:
	docker stop lucid_heisenberg
	docker rm lucid_heisenberg