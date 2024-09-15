#!/usr/bin/env bash

echo "开始编译文件"
buildFileName="sql2xorm"
BuildTime=`date +'%Y.%m.%d.%H:%M:%S'`
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags "-w -s -X 'sql2xorm/conf.BuildTime=${BuildTime}' -X 'sql2xorm/conf.BuildHash=$(git rev-parse --short HEAD)'" \
    -o ${buildFileName}

echo "编译完成, 编译时间${BuildTime}"
echo "开始压缩文件"
upx -9 ${buildFileName}
echo "完成压缩"