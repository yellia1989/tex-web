#!/bin/bash

web="$1"

if [ -f ./stop.sh ]; then
    ./stop.sh && rm -rf front
fi

tar -xjf backup/$web
chown root:root ./*
oldpwd=`pwd`
cd /data/tex/tools/script
source func.sh
source config.sh
cd $oldpwd
sql="select endpoint from db_tex.t_service where app='tex' and server='mfwregistry' and service='QueryObj'"
locatorpoint=$(echo "$sql"|mysql -h$registry_db_host -u$registry_db_user -p$registry_db_pass -N)
locapoint=$(echo $locatorpoint|sed 's/[ ]\+/:/7')
sed -i -e "s/TEMPLATE_LOCATOR/$locapoint/g" conf.cfg

LOCALIP=""
getLocalIp
sed -i -e "s/TEMPLATE_IP/$LOCALIP/g" conf.cfg
./start.sh
