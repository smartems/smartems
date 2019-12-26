+++
title = "Installing on RPM-based Linux"
description = "smartEMS Installation guide for Centos, Fedora, OpenSuse, Redhat."
keywords = ["grafana", "installation", "documentation", "centos", "fedora", "opensuse", "redhat"]
aliases = ["installation/installation/rpm"]
type = "docs"
[menu.docs]
name = "Installing on Centos / Redhat"
identifier = "rpm"
parent = "installation"
weight = 2
+++

# Installing on RPM-based Linux (CentOS, Fedora, OpenSuse, RedHat)

Read [Upgrading smartEMS]({{< relref "installation/upgrading.md" >}}) for tips and guidance on updating an existing
installation.

## Download

Go to the [download page](https://grafana.com/grafana/download?platform=linux) for the latest download
links.


You can install smartEMS using Yum directly.

```bash
$ sudo yum install <rpm package url>
```

You will find package urls on the [download page](https://grafana.com/grafana/download?platform=linux).

Or install manually using `rpm`. First execute

```bash
$ wget <rpm package url>
```

### On CentOS / Fedora / Redhat:

```bash
$ sudo yum install initscripts urw-fonts
$ sudo rpm -Uvh <local rpm package>
```

### On OpenSuse:

```bash
$ sudo rpm -i --nodeps <local rpm package>
```

## Install via YUM Repository

Add the following to a new file at `/etc/yum.repos.d/grafana.repo`

```bash
[grafana]
name=grafana
baseurl=https://packages.grafana.com/oss/rpm
repo_gpgcheck=1
enabled=1
gpgcheck=1
gpgkey=https://packages.grafana.com/gpg.key
sslverify=1
sslcacert=/etc/pki/tls/certs/ca-bundle.crt
```

There is a separate repository if you want beta releases.

```bash
[grafana]
name=grafana
baseurl=https://packages.grafana.com/oss/rpm-beta
repo_gpgcheck=1
enabled=1
gpgcheck=1
gpgkey=https://packages.grafana.com/gpg.key
sslverify=1
sslcacert=/etc/pki/tls/certs/ca-bundle.crt
```

Then install smartEMS via the `yum` command.

```bash
$ sudo yum install grafana
```

### RPM GPG Key

The RPMs are signed, you can verify the signature with this [public GPG
key](https://packages.grafana.com/gpg.key).

## Package details

- Installs binary to `/usr/sbin/smartems-server`
- Copies init.d script to `/etc/init.d/smartems-server`
- Installs default file (environment vars) to `/etc/sysconfig/smartems-server`
- Copies configuration file to `/etc/grafana/grafana.ini`
- Installs systemd service (if systemd is available) name `smartems-server.service`
- The default configuration uses a log file at `/var/log/grafana/grafana.log`
- The default configuration specifies an sqlite3 database at `/var/lib/grafana/grafana.db`

## Start the server (init.d service)

You can start smartEMS by running:

```bash
$ sudo service smartems-server start
```

This will start the `smartems-server` process as the `grafana` user,
which is created during package installation. The default HTTP port is
`3000`, and default user and group is `admin`.

Default login and password `admin`/ `admin`

To configure the smartEMS server to start at boot time:

```bash
$ sudo /sbin/chkconfig --add smartems-server
```

## Start the server (via systemd)

```bash
$ systemctl daemon-reload
$ systemctl start smartems-server
$ systemctl status smartems-server
```

### Enable the systemd service to start at boot

```bash
sudo systemctl enable smartems-server.service
```

## Environment file

The systemd service file and init.d script both use the file located at
`/etc/sysconfig/smartems-server` for environment variables used when
starting the back-end. Here you can override log directory, data
directory and other variables.

### Logging

By default smartEMS will log to `/var/log/grafana`

### Database

The default configuration specifies a sqlite3 database located at
`/var/lib/grafana/grafana.db`. Please backup this database before
upgrades. You can also use MySQL or Postgres as the smartEMS database, as detailed on [the configuration page]({{< relref "configuration.md#database" >}}).

## Configuration

The configuration file is located at `/etc/grafana/grafana.ini`.  Go the
[Configuration]({{< relref "configuration.md" >}}) page for details on all
those options.

### Adding data sources

- [Graphite]({{< relref "features/datasources/graphite.md" >}})
- [InfluxDB]({{< relref "features/datasources/influxdb.md" >}})
- [OpenTSDB]({{< relref "features/datasources/opentsdb.md" >}})
- [Prometheus]({{< relref "features/datasources/prometheus.md" >}})

### Server side image rendering

Server side image (png) rendering is a feature that is optional but very useful when sharing visualizations,
for example in alert notifications.

If the image is missing text make sure you have font packages installed.

```bash
yum install fontconfig
yum install freetype*
yum install urw-fonts
```

## Installing from binary tar file

Download [the latest `.tar.gz` file](https://grafana.com/get) and
extract it.  This will extract into a folder named after the version you
downloaded. This folder contains all files required to run smartEMS.  There are
no init scripts or install scripts in this package.

To configure smartEMS add a configuration file named `custom.ini` to the
`conf` folder and override any of the settings defined in
`conf/defaults.ini`.

Start smartEMS by executing `./bin/smartems-server web`. The `smartems-server`
binary needs the working directory to be the root install directory (where the
binary and the `public` folder is located).

## Logging in for the first time

To run smartEMS open your browser and go to http://localhost:3000/. 3000 is the default HTTP port that smartEMS listens to if you haven't [configured a different port](/installation/configuration/#http-port).
Then follow the instructions [here](/guides/getting_started/).
