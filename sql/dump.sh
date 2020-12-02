#!/bin/bash

if [ $# -ne 4 ]; then
    echo "Usage: $0 ip port user pwd"
    exit 100
fi

IP=$1
PORT=$2
USER=$3
PWD=$4

mysqldump -h$IP -P$PORT -u$USER -p$PWD --databases db_stat -d | sed 's/ AUTO_INCREMENT=[0-9]*\b//g' > db.sql
