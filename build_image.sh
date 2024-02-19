#!/bin/bash

if [ $# -ne 1 ] ;then
	echo "Usage: $0 version"
	exit -1
fi

version=$1

source make.sh

docker build -t "web:$version" -f web.Dockerfile .
