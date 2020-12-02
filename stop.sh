#!/bin/bash

server=web

old_path=`pwd`
cd `dirname $0`
path=`pwd`
exefile=$(readlink -f $path/$server)
kill $(/usr/sbin/lsof -n -P -d txt | grep "$exefile" | awk -vexefile="$exefile" '$9==exefile{print $2}') 2>/dev/null

until [ -z "`ps -ef|grep "$path/$server"|grep -v "grep"`" ]
do
	sleep 1
done

echo "stop $path/$server ok ...."

cd $old_path 1>/dev/null
