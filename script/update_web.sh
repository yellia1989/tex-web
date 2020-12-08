#!/bin/bash

if [ $# -ne 1 ] ;then
	echo "Usage: $0 env (d/47.103.96.228 u/139.224.58.141 r/47.91.45.105)"
	exit 100
fi

env="$1"
source remote_cmd.sh

case "$env" in
    d)
    ip=47.103.96.228
    cp ../conf_d.cfg conf.cfg
    ;;
    u)
    ip=139.224.58.141
    cp ../conf_u.cfg conf.cfg
    ;;
    r)
    ip=47.91.45.105
    cp ../conf_r.cfg conf.cfg
    ;;
    *)
    echo "invalid env"
    exit 0
    ;;
esac

runcmdroot root@$ip "[ ! -f /data/web ] && mkdir /data/web"
runcmdroot root@$ip "mkdir /data/web/backup"

web="web`date +%Y%m%d`.tar.gz"

tar -cjvf $web conf.cfg ../front ../data ../web ../start.sh ../stop.sh ../sql

if [ ! -f $web ]; then
    echo '打包web失败'
    exit 0
fi

echo "拷贝文件时间较长， 请耐心等待。。。"
putfile root@$ip ../update.sh /data/web/
putfile root@$ip $web /data/web/backup/

echo "更新web中。。。"
runcmdroot root@$ip "cd /data/web && ./update.sh $web"

rm -rf $web
rm -rf conf.cfg
exit 0
