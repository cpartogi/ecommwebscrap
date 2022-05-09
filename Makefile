SHELL := /bin/bash

## Golang Stuff
GOCMD=go
GORUN=$(GOCMD) run

tidy:
	$(GOCMD) mod tidy

run:
	$(GORUN) main.go