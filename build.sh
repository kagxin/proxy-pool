#!/bin/bash
# 脚本发生错误终止执行
set -e
echo "Build ..."
docker build -f Dockerfile -t proxypool .
echo "Clean .."
docker image prune -f
echo "Done"