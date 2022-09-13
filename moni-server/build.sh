#!/bin/bash
time=`date +%Y%m%d`
make
docker tag moni-server:latest moni-server:${time}
docker build -t moni-server:latest .
