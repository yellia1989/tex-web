#!/bin/bash

web="$1"

if [ -f ./stop.sh ]; then
    ./stop.sh && rm -rf front
fi

tar -xjvf backup/$web && ./start.sh
