rm working/blogkit

docker run --rm -v `pwd`:/go/src/blogkit -w /go/src/blogkit golang:1.12.6 \
go build -o working/blogkit -a -ldflags '-s -w -linkmode "external" -extldflags "-static"'

upx working/blogkit -9 -q > /dev/null

docker build -t blogkit:0.2.0 .
