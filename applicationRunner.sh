#!/bin/bash

cd backEnd/
pwd
docker build -t dipugodocker/pdf-editor:$version .
# docker run --rm -d -p 80:8080 dipugodocker/pdf-editor:$version
docker run --rm \
  -dit \
  -e PORT \
  --name pdf-editor \
  -p 80:${PORT} \
  dipugodocker/pdf-editor:$version

docker ps