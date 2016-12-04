FROM alpine:3.2

MAINTAINER Jakub Juszczak <jakub@posteo.de>

WORKDIR /

ADD ./bs3stat /bs3stat
ADD web/dist/ web/dist

RUN chmod +x bs3stat
ENTRYPOINT ["/bs3stat"]

EXPOSE 3000
