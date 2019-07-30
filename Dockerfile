# Base operating system: alpine, use latest version
FROM alpine:3.9.4

# How to contact me: via myownroc@163.com or myownroc@live.com
MAINTAINER DL myownroc@163.com

# Copy work units into docker image
COPY working/blogkit /blogkit/blogkit
COPY working/articles /blogkit/articles/
COPY working/config.json /blogkit/config.json
COPY working/favicon.ico /blogkit/favicon.ico
COPY working/static /blogkit/static/

# specify work directory
WORKDIR /blogkit

# default command
CMD ["/blogkit/blogkit"]
