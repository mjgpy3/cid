#!/bin/bash

export GOPATH=$CID_ROOT

redis-server - <<REDIS_CONFIG
  port $CID_REDIS_PORT
  dbfilename cid.rdb
  dir $CID_ROOT/data
REDIS_CONFIG

go run $GOPATH/src/cid/cid.go
