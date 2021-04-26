#!/bin/bash -e

go test -tags testrunmain -coverpkg=./... -coverprofile=c.out  ./... &
./wait-for-it 127.0.0.1:8090
curl -s 127.0.0.1:8090/hello > /dev/null
sleep 3
go tool cover -func=c.out
