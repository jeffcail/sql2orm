FROM alpine

# 设置工作目录
WORKDIR /www/wwwroot/sql2orm

# 添加可执行文件
COPY ./sql2orm .

EXPOSE 7892

ENTRYPOINT ["./sql2orm"]