#!/usr/bin/env bash

mkdir -p output/conf
cp conf/* output/conf/

go build -o output/WebService
