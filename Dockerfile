FROM ubuntu:14.04

RUN apt-get update

RUN mkdir /home/cid
VOLUME ["/home/cid"]

ENV CID_ROOT /home/cid
ENV CID_REDIS_PORT 6379

EXPOSE 5678

RUN apt-get install redis-server -y
RUN apt-get install golang -y
RUN apt-get install git -y

ENV GOPATH /home/cid

RUN go get github.com/garyburd/redigo/redis

ENTRYPOINT /home/cid/server.sh
