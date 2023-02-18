#!/bin/sh
wait-for "3306:3306" -- "$@"
CompileDaemon --build="go build main.go" --command=./main
