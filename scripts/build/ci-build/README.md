# smartems-build-container
smartEMS build container

## Description

This is a container for cross-platform builds of smartEMS. You can run it locally using the Makefile.

## Makefile targets

* `make run-with-local-source-copy`
  - Starts the container locally and copies your local sources into the container
* `make run-with-local-source-live`
  - Starts the container (as your user) locally and maps your smartEMS project dir into the container
* `make update-source`
  - Updates the sources in the container from your local sources
* `make stop`
  - Kills the container
* `make attach`
  - Opens bash within the running container

## Build/Publish Docker Image
In order to build and publish the smartEMS build Docker image, execute the following:
`./build-deploy.sh`.
