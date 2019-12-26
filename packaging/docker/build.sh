#!/bin/sh
set -e

BUILD_FAST=0
UBUNTU_BASE=0
TAG_SUFFIX=""

while [ "$1" != "" ]; do
  case "$1" in
    "--fast")
      BUILD_FAST=1
      echo "Fast build enabled"
      shift
      ;;
    "--ubuntu")
      UBUNTU_BASE=1
      TAG_SUFFIX="-ubuntu"
      echo "Ubuntu base image enabled"
      shift
      ;;
    * )
      # unknown param causes args to be passed through to $@
      break
      ;;
  esac
done

_smartems_tag=${1:-}
_docker_repo=${2:-smartems/smartems}

# If the tag starts with v, treat this as a official release
if echo "$_smartems_tag" | grep -q "^v"; then
  _smartems_version=$(echo "${_smartems_tag}" | cut -d "v" -f 2)
else
  _smartems_version=$_smartems_tag
fi

echo "Building ${_docker_repo}:${_smartems_version}${TAG_SUFFIX}"

export DOCKER_CLI_EXPERIMENTAL=enabled

# Build smartems image for a specific arch
docker_build () {
  arch=$1

  case "$arch" in
    "x64")
      base_arch=""
      repo_arch=""
      ;;
    "armv7")
      base_arch="arm32v7/"
      repo_arch="-arm32v7-linux"
      ;;
    "arm64")
      base_arch="arm64v8/"
      repo_arch="-arm64v8-linux"
      ;;
  esac
  if [ $UBUNTU_BASE = "0" ]; then
    libc="-musl"
    dockerfile="Dockerfile"
    base_image="${base_arch}alpine:3.10"
  else
    libc=""
    dockerfile="ubuntu.Dockerfile"
    base_image="${base_arch}ubuntu:18.10"
  fi

  smartems_tgz="smartems-latest.linux-${arch}${libc}.tar.gz"
  tag="${_docker_repo}${repo_arch}:${_smartems_version}${TAG_SUFFIX}"

  docker build \
    --build-arg BASE_IMAGE=${base_image} \
    --build-arg SMARTEMS_TGZ=${smartems_tgz} \
    --tag "${tag}" \
    --no-cache=true \
    -f "${dockerfile}" \
    .
}

docker_tag_linux_amd64 () {
  tag=$1
  docker tag "${_docker_repo}:${_smartems_version}${TAG_SUFFIX}" "${_docker_repo}:${tag}${TAG_SUFFIX}"
}

# Tag docker images of all architectures
docker_tag_all () {
  tag=$1
  docker_tag_linux_amd64 $1
  if [ $BUILD_FAST = "0" ]; then
    docker tag "${_docker_repo}-arm32v7-linux:${_smartems_version}${TAG_SUFFIX}" "${_docker_repo}-arm32v7-linux:${tag}${TAG_SUFFIX}"
    docker tag "${_docker_repo}-arm64v8-linux:${_smartems_version}${TAG_SUFFIX}" "${_docker_repo}-arm64v8-linux:${tag}${TAG_SUFFIX}"
  fi
}

docker_build "x64"
if [ $BUILD_FAST = "0" ]; then
  docker_build "armv7"
  docker_build "arm64"
fi

# Tag as 'latest' for official release; otherwise tag as smartems/smartems:master
if echo "$_smartems_tag" | grep -q "^v"; then
  docker_tag_all "latest"
  # Create the expected tag for running the end to end tests successfully
  docker tag "${_docker_repo}:${_smartems_version}${TAG_SUFFIX}" "smartems/smartems-dev:${_smartems_tag}${TAG_SUFFIX}"
else
  docker_tag_all "master"
  docker tag "${_docker_repo}:${_smartems_version}${TAG_SUFFIX}" "smartems/smartems-dev:${_smartems_version}${TAG_SUFFIX}"
fi
