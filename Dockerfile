FROM alpine:3.9.4
MAINTAINER DL myownroc@163.com
COPY working/ /blogkit/
WORKDIR /blogkit
CMD ["/blogkit/blogkit"]
