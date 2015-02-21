FROM ubuntu:14.04

RUN apt-get update

RUN mkdir /home/cid
VOLUME ["/home/cid"]

EXPOSE 5678

RUN apt-get install redis-server -y
RUN apt-get install golang -y


ENTRYPOINT /home/cid/server.sh
