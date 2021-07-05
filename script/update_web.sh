#!/bin/bash

if [ $# -ne 1 ] ;then
	echo "Usage: $0 env (u/124.156.205.111 r/101.32.168.86)"
	exit 100
fi

env="$1"
source remote_cmd.sh

case "$env" in
    u)
    ip=124.156.205.111
    cp ../conf_u.cfg conf.cfg
    ;;
    r)
    ip=101.32.168.86
    cp ../conf_r.cfg conf.cfg
    ;;
    *)
    echo "invalid env"
    exit 0
    ;;
esac

runcmd root@$ip "[ ! -f /data/web ] && mkdir /data/web"
runcmd root@$ip "mkdir /data/web/backup"

web="web`date +%Y%m%d`.tar.gz"

tar -cjvf $web conf.cfg ../front ../web ../start.sh ../stop.sh ../sql

if [ ! -f $web ]; then
    echo '打包web失败'
    exit 0
fi

echo "拷贝文件时间较长， 请耐心等待。。。"
putfile root@$ip ../update.sh /data/web/
putfile root@$ip $web /data/web/backup/

echo "更新web中。。。"
runcmd root@$ip "cd /data/web && ./update.sh $web"

rm -rf $web
rm -rf conf.cfg
exit 0
