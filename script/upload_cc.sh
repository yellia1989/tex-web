#!/bin/bash

# 上传包到cc机器
env_helper="env (d/192.168.0.15 u/106.15.139.153 r/47.241.161.10 robot/101.133.160.60)"
if [ $# -ne 1 ] ;then
	echo "Usage: $0 $env_helper"
	exit 100
fi

env="$1"
source remote_cmd.sh

case "$env" in
    d)
    cc_ip=192.168.0.15
    ;;
    u)
    cc_ip=106.15.139.153
    ;;
    r)
    cc_ip=47.241.161.10
    ;;
    t)
    cc_ip=101.132.43.124
    ;;
    robot)
    cc_ip=101.133.160.60
    ;;
    *)
    echo "invalid env, $env_helper"
    exit 0
    ;;
esac

path=/data/web/backup/tmp
runcmd root@$cc_ip "mkdir -p $path"

web="web`date +%Y%m%d`.tar.gz"

tar -cjf $web ../front ../web ../data ../start.sh ../stop.sh

if [ ! -f $web ]; then
    echo '打包web失败'
    exit 0
fi

echo "拷贝文件时间较长， 请耐心等待。。。"
putfile root@$cc_ip ../conf.cfg $path
putfile root@$cc_ip ../update.sh $path
putfile root@$cc_ip $web $path

rm -rf $web
exit 0
