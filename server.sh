#!/bin/bash

export GOPATH=`dirname $0`

go run $GOPATH/src/cid/cid.go
