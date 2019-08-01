#!/bin/bash

# run blogkit with docker image
docker run --rm --name blogkit -v `pwd`/working:/blogkit -p 80:80 blogkit:1.2.4
