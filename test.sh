#!/bin/bash

docker run --rm --name blogkit -v `pwd`/working:/blogkit -p 80:80 blogkit:0.3.0
