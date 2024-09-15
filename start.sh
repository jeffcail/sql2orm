#! /bin/bash

if [ ! -f 'sql2orm' ]; then
  echo 文件不存在! 待添加的安装包: 'sql2orm'
  exit
fi

echo "sql2orm..."
sleep 3
docker stop sql2orm

sleep 2
docker rm sql2orm

docker rmi sql2orm
echo ""

echo "sql2orm packing..."
sleep 3
docker build -t sql2orm .
echo ""

echo "sql2orm running..."
sleep 3

docker run \
  -p 7892:7892 \
  --name sql2orm \
  -v /www/wwwroot/sql2orm/dist:/www/wwwroot/sql2orm/dist \
  -v /etc/localtime:/etc/localtime \
  -d sql2orm \

  docker logs -f sql2orm | sed '/Started sql2orm Application/q'

  echo ""
