#!/bin/bash

# remove old executable
rm working/blogkit

# compile by golang(use latest and not alpine version) with docker image
docker run --rm -v `pwd`:/go/src/blogkit -w /go/src/blogkit golang:1.12.6 \
go build -o working/blogkit -a -ldflags '-s -w -linkmode "external" -extldflags "-static"'

# reduce size of executable by upx(use latest version)
upx working/blogkit -9 -q > /dev/null

# remove old docker image
docker rmi blogkit:1.2.5

# build docker image
docker build -t blogkit:1.2.5 .
