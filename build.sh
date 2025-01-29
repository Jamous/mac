#!/bin/sh
#Build Linux version
go build -o "bin/linux_x86/" .
GOOS=windows GOARCH=amd64 go build -o "bin/windows_x86/" .

#Reset build envireoment
unset GOOS unset GOARCH 
