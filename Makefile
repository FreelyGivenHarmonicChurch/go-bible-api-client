
.PHONY: env build run

# env:
# 	export $(grep -v '^#' .env | xargs)

build:
	go build .

run: env
	./go-bible-api-client
