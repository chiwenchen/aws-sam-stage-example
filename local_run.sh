#!/usr/bin/env bash

GOOS=linux go build -o main
SomeString=iamfromlocal sam local start-api
