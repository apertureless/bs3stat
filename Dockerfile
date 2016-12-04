FROM golang:latest

MAINTAINER Jakub Juszczak <jakub@posteo.de>

WORKDIR /

ADD ./bs3stat /bs3stat
ADD web/dist/ web/dist

ENTRYPOINT main

EXPOSE 3000
