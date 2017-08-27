#!/usr/bin/env bash

docker run -d --name=roachproxy \
--net=roachnet \
-p 26256:26257 \
roachhaproxy