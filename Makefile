#!/bin/bash

rootDir = /git/go/src/golinkie
mainRoot = $(rootDir)/core
serverRoot = $(rootDir)/server
project = golinkie
all: $(mainRoot)/main $(serverRoot)/server

$(mainRoot)/main: $(mainRoot)/main.go 
	go build -o $(mainRoot)/core $(project)/core

$(serverRoot)/server: $(serverRoot)/server.go
	go build -o $(serverRoot)/server $(project)/server
    
