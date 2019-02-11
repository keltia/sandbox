# Main Makefile for sandbox
#
# Copyright 2018-2019 © by Ollivier Robert <roberto@keltia.net>
#

GOBIN=	${GOPATH}/bin

GO=		go
SRCS=	sandbox.go

OPTS=	-ldflags="-s -w" -v

all: build

build: ${SRCS}
	${GO} build ${OPTS} .

test: build
	${GO} test .

install:
	${GO} install ${OPTS} .

clean:
	${GO} clean .

push:
	git push --all
	git push --tags
