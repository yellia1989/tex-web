#!/bin/bash

old_path=`pwd`
cd `dirname $0`
path=`pwd`
server=web

$path/stop.sh

nohup $path/$server >> log 2>&1  &

sleep 1
if [ -n "`ps -ef|grep "$path/$server"|grep -v "grep"`" ]
then
    echo "start $path/$server ok ...."
else
    echo "start $path/$server failed ...."
fi

cd $old_path 1>/dev/null
