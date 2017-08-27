#!/usr/bin/env bash
docker run -d \
--name=$1 \
--hostname=$1 \
--net=roachnet \
-p $2:26257 \
cockroachdb/cockroach:v1.0.4 start --insecure --join=roach1
