FROM scratch

MAINTAINER Jakub Juszczak <jakub@posteo.de>

WORKDIR /

ADD ./bs3stat /bs3stat
ADD web/dist/ web/dist

CMD ["/bs3stat"]

EXPOSE 3000
