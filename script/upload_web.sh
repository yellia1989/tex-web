#!/bin/bash

# 上传包到cc机器
env_helper="env (d/124.222.61.82 u/106.15.36.228 u/47.254.134.86)"
if [ $# -ne 1 ] ;then
	echo "Usage: $0 $env_helper"
	exit 100
fi

env="$1"
source remote_cmd.sh

case "$env" in
    d)
    cc_ip=124.222.61.82
    ;;
    u)
    cc_ip=106.15.36.228
    ;;
    r)
    cc_ip=47.254.134.86
    ;;
    *)
    echo "invalid env, $env_helper"
    exit 0
    ;;
esac
web="web`date +%Y%m%d`.tar.gz"

tar -czf $web ../front ../web ../data ../sql ../conf.cfg ../start.sh ../stop.sh

if [ ! -f $web ]; then
    echo '打包web失败'
    exit 0
fi

echo "拷贝文件时间较长， 请耐心等待。。。"

path=/data/patch/web
runcmd yk@$cc_ip "mkdir -p $path"
putfile yk@$cc_ip $web $path

rm -rf $web
exit 0
