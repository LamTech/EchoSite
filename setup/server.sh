#!/bin/sh
./echosited -log="./log/server.txt" -log-level="DEBUG" -tlsKey="/etc/letsencrypt/live/online.echosite.cn/privkey.pem" -tlsCrt="/etc/letsencrypt/live/online.echosite.cn/fullchain.pem" -domain="online.echosite.cn" -httpAddr=":80" -httpsAddr=":443" -echositeAddr="https://www.echosite.cn/check/permission"
