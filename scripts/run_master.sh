#!/usr/bin/env bash
docker run -d --name=roach1 \
--hostname=roach1 \
--net=roachnet \
-p 26257:26257 -p 8080:8080 \
cockroachdb/cockroach:v1.0.4 start --insecure
