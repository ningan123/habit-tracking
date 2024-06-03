COMMIT := $(shell git rev-parse --short HEAD)
BASE_VERSION ?= v0.0.1
VERSION := ${BASE_VERSION}-${COMMIT}

arch=amd64
archARM=arm64

habit:
	GOOS=linux GOARCH=$(arch) go build -ldflags "-X main.Version=${VERSION}" -o bin/habit-tracking 

habitRAM:
	GOOS=linux GOARCH=$(archARM) go build -ldflags "-X main.Version=${VERSION}" -o bin/habit-tracking
	
# 测试可用
# go build -ldflags "-X main.Version=v1.0.0" -o bin/habit-tracking
# GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=v1.0.0" -o bin/habit-tracking
# ./bin/habit-tracking -h