debug ?= gctrace=0
addr  ?= ws://127.0.0.1:9999
limit ?= 5

export GODEBUG=$(debug)

.PHONY: server
server: dependencies
	go run ./examples/server/server.go \
		--debug                    \
		--addr=$(addr)


.PHONY: client
client: dependencies
	go run ./examples/client/client.go \
		--debug                    \
		--addr=$(addr)             \
		--limit=$(limit)
