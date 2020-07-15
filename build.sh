#!/bin/bash
# 脚本发生错误终止执行
set -e
IMAGE_NAME="registry.cn-shanghai.aliyuncs.com/release-lib/proxy-pool"

IMAGE_FULL_NAME="$IMAGE_NAME:$1"

echo ${IMAGE_FULL_NAME}

echo "Building image..."

docker build -f Dockerfile -t ${IMAGE_FULL_NAME} .

echo "Push ..."
docker login --username=kangxinhappy2016 registry.cn-shanghai.aliyuncs.com
docker push ${IMAGE_FULL_NAME}
echo "Done"