#!/bin/bash

_version="1.2.2"
_tag="smartems/smartems-ci-deploy:${_version}"

docker build -t $_tag .
docker push $_tag
