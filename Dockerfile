FROM alpine:3.9.4
MAINTAINER DL myownroc@163.com
COPY working/blogkit /blogkit/blogkit
COPY working/articles /blogkit/articles/
COPY working/config.json /blogkit/config.json
COPY working/favicon.ico /blogkit/favicon.ico
COPY working/static/files/blogkit.png /blogkit/static/files/blogkit.png
COPY working/static/files/blogkit.zip /blogkit/static/files/blogkit.zip
WORKDIR /blogkit
CMD ["/blogkit/blogkit"]
