BINARY=bin/omake

ifeq ($(OS),Windows_NT)
	BINARY:=$(BINARY).exe
endif

default:
	go build -o $(BINARY) main.go