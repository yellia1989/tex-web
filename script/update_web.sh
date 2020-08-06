#!/bin/bash

if [ $# -ne 1 ] ;then
	echo "Usage: $0 env (d/192.168.0.7 t/101.133.141.46)"
	exit 100
fi

env="$1"
source remote_cmd.sh

case "$env" in
    d)
    ip=192.168.0.7
    cp ../conf_d.cfg conf.cfg
    ;;
    t)
    ip=101.133.141.46
    cp ../conf_t.cfg conf.cfg
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
putfile root@$ip $web /data/web/backup/

echo "更新web中。。。"
runcmdroot root@$ip "cd /data/web && ./stop.sh && rm -rf front && tar -xjvf backup/$web && ./start.sh"

rm -rf $web
rm -rf conf.cfg
exit 0
