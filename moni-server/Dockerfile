FROM centos:7
WORKDIR /data/moni-server
COPY moni-server .
COPY config.json .
RUN chmod +x moni-server
EXPOSE 80
ENTRYPOINT ./moni-server -c  config.json