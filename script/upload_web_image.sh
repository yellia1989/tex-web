#!/bin/bash

# 上传包到cc机器
env_helper="env (stg pro)"
if [ $# -ne 1 ] ;then
	echo "Usage: $0 $env_helper"
	exit 100
fi

env="$1"

case "$env" in
    stg)
    profile=hametsu_stg
    instanceId=cri-fe5eqsg73zl4vd4m
    repo=staging-legolas-registry.ap-northeast-1.cr.aliyuncs.com
    ;;
    pro)
    profile=hametsu_pro
    instanceId=cri-azy285s79zw3ti5v
    repo=production-legolas-registry.ap-northeast-1.cr.aliyuncs.com
    ;;
    *)
    echo "invalid env, $env_helper"
    exit 0
    ;;
esac

webVer="`date +%Y%m%d%H%M%S`"

cd .. && ./build_image.sh $webVer

password=$(aliyun cr GetAuthorizationToken --profile $profile --region ap-northeast-1 --InstanceId "$instanceId" --version 2018-12-01 --force | jq .AuthorizationToken | tr -d '"')
docker login --username=cr_temp_user --password=$password $repo
docker tag web:$webVer $repo/hametsu/web:$webVer
docker push $repo/hametsu/web:$webVer
docker image rm $repo/hametsu/web:$webVer

exit 0
