#!/bin/bash

set -u
set -e

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

./upload_cc.sh $env
web="web`date +%Y%m%d`.tar.gz"
runcmd root@$cc_ip "cd /data/tex/tools/script && ./cc_install_web.sh $web"

exit 0
