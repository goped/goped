# Makefile

# Build Data
IS_DIRTY    := $(test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE  := `date +%FT%T%z`
COMMIT_HASH := `git rev-parse --short HEAD 2>/dev/null`
BRANCH      := $(shell git rev-parse --abbrev-ref HEAD)
APP_VERSION := $(shell cat `git rev-parse --show-toplevel`/VERSION)
MACHINE     := $(shell hostname -f)
GO_VERSION  := $(go version | cut -d ' ' -f3)

OS=$(shell uname)
OS_SIMPLE = $(subst Darwin,darwin,$(subst Linux,linux,$(subst FreeBSD,freebsd,$(OS))))
ARCH=$(shell uname -m)

# Compliler Flags
LDFLAGS := -ldflags \
	" -X main.buildGoVersion ${GO_VERSION}\
		-X main.buildMachine ${MACHINE}\
		-X main.buildOS ${OS_SIMPLE}\
		-X main.buildAppVersion ${APP_VERSION}\
		-X main.buildBranch ${BRANCH}\
		-X main.buildCommitHash ${COMMIT_HASH}${IS_DIRTY}\
		-X main.buildDate ${BUILD_DATE}"

all: build

build-info:
	echo ${COMMIT_HASH}${IS_DIRTY}
	echo ${BUILD_DATE}

build:
	${GOROOT}/bin/go build ${LDFLAGS} 

install:
	${GOROOT}/bin/go install -race ${LDFLAGS} 
