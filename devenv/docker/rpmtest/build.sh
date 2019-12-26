#!/bin/bash

cp Dockerfile ../../dist
cd ../../dist

docker build --tag "smartems/rpmtest" .

rm Dockerfile

docker run -i -t smartems/rpmtest /bin/bash
