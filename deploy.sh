#!/bin/bash

#Get dependencies
go get github.com/constabulary/gb/...
go get -u github.com/golang/lint/golint

#Build
make vendor_get
make

#Deploy
./bin/fastestroute