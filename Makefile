-include Makefile.local.mk

build:
	go build

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o mon-linux-amd64

build-and-scp:
ifndef TO
	$(error TO is undefined, use 'make build-and-scp TO=root@server.example.com:/mon')
endif

	@make build-linux-amd64
	scp mon-linux-amd64 ${TO}
