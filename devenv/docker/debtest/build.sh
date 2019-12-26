#!/bin/bash

cp Dockerfile ../../dist
cd ../../dist

docker build --tag "smartems/debtest" .

rm Dockerfile

docker run -i -t smartems/debtest /bin/bash
