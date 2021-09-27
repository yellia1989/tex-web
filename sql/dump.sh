#!/bin/bash

if [ $# -ne 4 ]; then
    echo "Usage: $0 ip port user pwd"
    exit 100
fi

IP=$1
PORT=$2
USER=$3
PWD=$4

mysqldump -h$IP -P$PORT -u$USER -p$PWD --databases db_stat -d --ignore-table=db_stat.sys_user --ignore-table=db_stat.sys_role_perm --ignore-table=db_stat.sys_role --ignore-table=db_stat.sys_perms --ignore-table=db_stat.sys_menu_role | sed 's/ AUTO_INCREMENT=[0-9]*\b//g' > db.sql
mysqldump -h$IP -P$PORT -u$USER -p$PWD db_stat sys_user sys_role_perm sys_role sys_perms sys_menu_role | sed 's/ AUTO_INCREMENT=[0-9]*\b//g' > db2.sql
