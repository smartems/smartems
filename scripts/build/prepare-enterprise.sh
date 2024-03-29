#!/bin/bash

cd ..

if [ -z "$CIRCLE_TAG" ]; then
  _target="master"
else
  _target="$CIRCLE_TAG"
fi

git clone -b "$_target" --single-branch git@github.com:smartems/smartems-enterprise.git --depth 1

cd smartems-enterprise || exit
./build.sh
