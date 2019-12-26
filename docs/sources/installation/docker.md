+++
title = "Installing using Docker"
description = "Installing smartEMS using Docker guide"
keywords = ["smartems", "configuration", "documentation", "docker"]
type = "docs"
[menu.docs]
name = "Installing using Docker"
identifier = "docker"
parent = "installation"
weight = 4
+++

# Installing using Docker

smartEMS is very easy to install and run using the official docker container.

```bash
$ docker run -d -p 3000:3000 smartems/smartems
```

## Configuration

All options defined in `conf/smartems.ini` can be overridden using environment
variables by using the syntax `GF_<SectionName>_<KeyName>`.
For example:

```bash
$ docker run -d \
  -p 3000:3000 \
  --name=smartems \
  -e "GF_SERVER_ROOT_URL=http://smartems.server.name" \
  -e "GF_SECURITY_ADMIN_PASSWORD=secret" \
  smartems/smartems
```

The back-end web server has a number of configuration options. Go to the
[Configuration]({{< relref "configuration.md" >}}) page for details on all
those options.

> For any changes to `conf/smartems.ini` (or corresponding environment variables) to take effect you need to restart smartEMS by restarting the Docker container.

### Default Paths

The following settings are hard-coded when launching the smartEMS Docker container and can only be overridden using environment variables, not in `conf/smartems.ini`.

Setting               | Default value
----------------------|---------------------------
GF_PATHS_CONFIG       | /etc/smartems/smartems.ini
GF_PATHS_DATA         | /var/lib/smartems
GF_PATHS_HOME         | /usr/share/smartems
GF_PATHS_LOGS         | /var/log/smartems
GF_PATHS_PLUGINS      | /var/lib/smartems/plugins
GF_PATHS_PROVISIONING | /etc/smartems/provisioning

## Image Variants

The official smartEMS Docker image comes in two variants.

**`smartems/smartems:<version>`:**

> **Note:** This image was based on [Ubuntu](https://ubuntu.com/) before version 6.4.0.

This is the default image. This image is based on the popular [Alpine Linux project](http://alpinelinux.org), available in [the alpine official image](https://hub.docker.com/_/alpine). Alpine Linux is much smaller than most distribution base images, and thus leads to slimmer and more secure images.

This variant is highly recommended when security and final image size being as small as possible is desired. The main caveat to note is that it does use [musl libc](http://www.musl-libc.org) instead of [glibc and friends](http://www.etalabs.net/compare_libcs.html), so certain software might run into issues depending on the depth of their libc requirements. However, most software doesn't have an issue with this, so this variant is usually a very safe choice.

**`smartems/smartems:<version>-ubuntu`:**

> **Note:** This image is available since version 6.5.0.

This image is based on [Ubuntu](https://ubuntu.com/), available in [the ubuntu official image](https://hub.docker.com/_/ubuntu).
This is an alternative image for those who prefer an [Ubuntu](https://ubuntu.com/) based image and/or who are dependent on certain
tooling not available for Alpine.

## Running a specific version of smartEMS

```bash
# specify right tag, e.g. 6.5.0 - see Docker Hub for available tags
$ docker run -d -p 3000:3000 --name smartems smartems/smartems:6.5.0
# ubuntu based images available since smartEMS 6.5.0
$ docker run -d -p 3000:3000 --name smartems smartems/smartems:6.5.0-ubuntu
```

## Running the master branch

For every successful build of the master branch we update the `smartems/smartems:master` and `smartems/smartems:master-ubuntu`. Additionally, two new tags are created, `smartems/smartems-dev:master-<commit hash>` and `smartems/smartems-dev:master-<commit hash>-ubuntu`, which includes the hash of the git commit that was built. This means you can always get the latest version of smartEMS.

When running smartEMS master in production we **strongly** recommend that you use the `smartems/smartems-dev:master-<commit hash>` tag as that will guarantee that you use a specific version of smartEMS instead of whatever was the most recent commit at the time.

For a list of available tags, check out [smartems/smartems](https://hub.docker.com/r/smartems/smartems/tags/) and [smartems/smartems-dev](https://hub.docker.com/r/smartems/smartems-dev/tags/).

## Installing Plugins for smartEMS

Pass the plugins you want installed to docker with the `GF_INSTALL_PLUGINS` environment variable as a comma separated list. This will pass each plugin name to `smartems-cli plugins install ${plugin}` and install them when smartEMS starts.

```bash
docker run -d \
  -p 3000:3000 \
  --name=smartems \
  -e "GF_INSTALL_PLUGINS=smartems-clock-panel,smartems-simple-json-datasource" \
  smartems/smartems
```

> If you need to specify the version of a plugin, you can add it to the `GF_INSTALL_PLUGINS` environment variable. Otherwise, the latest will be assumed. For example: `-e "GF_INSTALL_PLUGINS=smartems-clock-panel 1.0.1,smartems-simple-json-datasource 1.3.5"`

## Building a custom smartEMS image

In the [smartEMS GitHub repository](https://github.com/smartems/smartems/tree/master/packaging/docker) there is a folder called `custom/` which two includes Dockerfiles, `Dockerfile` and `ubuntu.Dockerfile`, that can be used to build a custom smartEMS image.
It accepts `SMARTEMS_VERSION`, `GF_INSTALL_PLUGINS` and `GF_INSTALL_IMAGE_RENDERER_PLUGIN` as build arguments.

### With pre-installed plugins

> If you need to specify the version of a plugin, you can add it to the `GF_INSTALL_PLUGINS` build argument. Otherwise, the latest will be assumed. For example: `--build-arg "GF_INSTALL_PLUGINS=smartems-clock-panel 1.0.1,smartems-simple-json-datasource 1.3.5"`

Example of how to build and run:
```bash
cd custom
docker build \
  --build-arg "SMARTEMS_VERSION=latest" \
  --build-arg "GF_INSTALL_PLUGINS=smartems-clock-panel,smartems-simple-json-datasource" \
  -t smartems-custom -f Dockerfile .

docker run -d -p 3000:3000 --name=smartems smartems-custom
```

Replace `Dockerfile` in above example with `ubuntu.Dockerfile` to build a custom Ubuntu based image (smartEMS 6.5+).

### With smartEMS Image Renderer plugin pre-installed

> Only available in smartEMS v6.5+ and experimental.

The [smartEMS Image Renderer plugin](/administration/image_rendering/#smartems-image-renderer-plugin) does not
currently work if it is installed in smartEMS docker image.
You can build a custom docker image by using the `GF_INSTALL_IMAGE_RENDERER_PLUGIN` build argument.
This will install additional dependencies needed for the smartEMS Image Renderer plugin to run.

Example of how to build and run:
```bash
cd custom
docker build \
  --build-arg "SMARTEMS_VERSION=latest" \
  --build-arg "GF_INSTALL_IMAGE_RENDERER_PLUGIN=true" \
  -t smartems-custom -f Dockerfile .

docker run -d -p 3000:3000 --name=smartems smartems-custom
```

Replace `Dockerfile` in above example with `ubuntu.Dockerfile` to build a custom Ubuntu based image.

## Installing Plugins from other sources

> Only available in smartEMS v5.3.1+

It's possible to install plugins from custom url:s by specifying the url like this: `GF_INSTALL_PLUGINS=<url to plugin zip>;<plugin name>`

```bash
docker run -d \
  -p 3000:3000 \
  --name=smartems \
  -e "GF_INSTALL_PLUGINS=http://plugin-domain.com/my-custom-plugin.zip;custom-plugin" \
  smartems/smartems
```

## Configuring AWS Credentials for CloudWatch Support

```bash
$ docker run -d \
  -p 3000:3000 \
  --name=smartems \
  -e "GF_AWS_PROFILES=default" \
  -e "GF_AWS_default_ACCESS_KEY_ID=YOUR_ACCESS_KEY" \
  -e "GF_AWS_default_SECRET_ACCESS_KEY=YOUR_SECRET_KEY" \
  -e "GF_AWS_default_REGION=us-east-1" \
  smartems/smartems
```

You may also specify multiple profiles to `GF_AWS_PROFILES` (e.g.
`GF_AWS_PROFILES=default another`).

Supported variables:

- `GF_AWS_${profile}_ACCESS_KEY_ID`: AWS access key ID (required).
- `GF_AWS_${profile}_SECRET_ACCESS_KEY`: AWS secret access  key (required).
- `GF_AWS_${profile}_REGION`: AWS region (optional).

## smartEMS container with persistent storage (recommended)

```bash
# create a persistent volume for your data in /var/lib/smartems (database and plugins)
docker volume create smartems-storage

# start smartems
docker run -d -p 3000:3000 --name=smartems -v smartems-storage:/var/lib/smartems smartems/smartems
```

## smartEMS container using bind mounts

You may want to run smartEMS in Docker but use folders on your host for the database or configuration. When doing so it becomes important to start the container with a user that is able to access and write to the folder you map into the container.

```bash
mkdir data # creates a folder for your data
ID=$(id -u) # saves your user id in the ID variable

# starts smartems with your user id and using the data folder
docker run -d --user $ID --volume "$PWD/data:/var/lib/smartems" -p 3000:3000 smartems/smartems:5.1.0
```

## Reading secrets from files (support for Docker Secrets)

> Only available in smartEMS v5.2+.

It's possible to supply smartEMS with configuration through files. This works well with [Docker Secrets](https://docs.docker.com/engine/swarm/secrets/) as the secrets by default gets mapped into `/run/secrets/<name of secret>` of the container.

You can do this with any of the configuration options in conf/smartems.ini by setting `GF_<SectionName>_<KeyName>__FILE` to the path of the file holding the secret.

Let's say you want to set the admin password this way.

- Admin password secret: `/run/secrets/admin_password`
- Environment variable: `GF_SECURITY_ADMIN_PASSWORD__FILE=/run/secrets/admin_password`


## Migration from a previous version of the docker container to 5.1 or later

The docker container for smartEMS has seen a major rewrite for 5.1.

**Important changes**

* file ownership is no longer modified during startup with `chown`
* default user id `472` instead of `104`
* no more implicit volumes
  - `/var/lib/smartems`
  - `/etc/smartems`
  - `/var/log/smartems`

### Removal of implicit volumes

Previously `/var/lib/smartems`, `/etc/smartems` and `/var/log/smartems` were defined as volumes in the `Dockerfile`. This led to the creation of three volumes each time a new instance of the smartEMS container started, whether you wanted it or not.

You should always be careful to define your own named volume for storage, but if you depended on these volumes you should be aware that an upgraded container will no longer have them.

**Warning**: when migrating from an earlier version to 5.1 or later using docker compose and implicit volumes you need to use `docker inspect` to find out which volumes your container is mapped to so that you can map them to the upgraded container as well. You will also have to change file ownership (or user) as documented below.

### User ID changes

In 5.1 we switched the id of the smartems user. Unfortunately this means that files created prior to 5.1 won't have the correct permissions for later versions. We made this change so that it would be more likely that the smartems users id would be unique to smartEMS. For example, on Ubuntu 16.04 `104` is already in use by the syslog user.

Version | User    | User ID
--------|---------|---------
< 5.1   | smartems | 104
>= 5.1  | smartems | 472

There are two possible solutions to this problem. Either you start the new container as the root user and change ownership from `104` to `472` or you start the upgraded container as user `104`.

#### Running docker as a different user

```bash
docker run --user 104 --volume "<your volume mapping here>" smartems/smartems:5.1.0
```

##### Specifying a user in docker-compose.yml
```yaml
version: "2"

services:
  smartems:
    image: smartems/smartems:5.1.0
    ports:
      - 3000:3000
    user: "104"
```

#### Modifying permissions

The commands below will run bash inside the smartEMS container with your volume mapped in. This makes it possible to modify the file ownership to match the new container. Always be careful when modifying permissions.

```bash
$ docker run -ti --user root --volume "<your volume mapping here>" --entrypoint bash smartems/smartems:5.1.0

# in the container you just started:
chown -R root:root /etc/smartems && \
  chmod -R a+r /etc/smartems && \
  chown -R smartems:smartems /var/lib/smartems && \
  chown -R smartems:smartems /usr/share/smartems
```

## Migration from a previous version of the docker container to 6.4 or later

smartEMS’s docker image was changed to be based on [Alpine](http://alpinelinux.org) instead of [Ubuntu](https://ubuntu.com/).

## Migration from a previous version of the docker container to 6.5 or later

smartEMS Docker image now comes in two variants, one [Alpine](http://alpinelinux.org) based and one [Ubuntu](https://ubuntu.com/) based, see [Image Variants](#image-variants) for details.

## Logging in for the first time

To run smartEMS open your browser and go to http://localhost:3000/. 3000 is the default HTTP port that smartEMS listens to if you haven't [configured a different port](/installation/configuration/#http-port).
Then follow the instructions [here](/guides/getting_started/).
