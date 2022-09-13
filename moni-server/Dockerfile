FROM centos:7
WORKDIR /data/moni-server
COPY moni-server .
RUN chmod +x moni-server
EXPOSE 80
ENTRYPOINT ./moni-server -c  config.json


# docker run -d -p 8083:80 --restart always -v config.json:/data/moni-server  -v /data/moni-server/logs:/data/moni-server/logs  moni-server