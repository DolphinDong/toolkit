FROM centos:7
WORKDIR /data/moni-server
COPY moni-server .
RUN chmod +x moni-server
#Dockerfile 时区设置
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

EXPOSE 80
ENTRYPOINT ./moni-server -c  config.json


# docker run -d -p 8083:80 --restart always -v /data/moni-server/config.json:/data/moni-server/config.json  -v /data/moni-server/logs:/data/moni-server/logs  moni-server