+++
title = "Installing on Debian / Ubuntu"
description = "Install guide for smartEMS"
keywords = ["smartems", "installation", "documentation"]
type = "docs"
aliases = ["/installation/installation/debian"]
[menu.docs]
name = "Installing on Ubuntu / Debian"
identifier = "debian"
parent = "installation"
weight = 1
+++

# Installing on Debian / Ubuntu

Read [Upgrading smartEMS]({{< relref "installation/upgrading.md" >}}) for tips and guidance on updating an existing
installation.

## Download

Go to the [download page](https://smartems.com/smartems/download?platform=linux) for the latest download
links.

```bash
wget <debian package url>
sudo apt-get install -y adduser libfontconfig1
sudo dpkg -i smartems_<version>_amd64.deb
```

You will find package urls on the [download page](https://smartems.com/smartems/download?platform=linux).

## APT Repository

The command `add-apt-repository` isn't a default app on Debian 9 and requires
```bash
apt-get install -y software-properties-common
```

Install the repository for stable releases

```bash
sudo add-apt-repository "deb https://packages.smartems.com/oss/deb stable main"
```

There is a separate repository if you want beta releases.

```bash
sudo add-apt-repository "deb https://packages.smartems.com/oss/deb beta main"
```

Use the above line even if you are on Ubuntu or another Debian version. Then add our gpg key. This allows you to install signed packages.

```bash
wget -q -O - https://packages.smartems.com/gpg.key | sudo apt-key add -
```

Update your Apt repositories and install smartEMS

```bash
sudo apt-get update
sudo apt-get install smartems
```

On some older versions of Ubuntu and Debian you may need to install the
`apt-transport-https` package which is needed to fetch packages over
HTTPS.

```bash
sudo apt-get install -y apt-transport-https
```

## Package details

- Installs binary to `/usr/sbin/smartems-server`
- Installs Init.d script to `/etc/init.d/smartems-server`
- Creates default file (environment vars) to `/etc/default/smartems-server`
- Installs configuration file to `/etc/smartems/smartems.ini`
- Installs systemd service (if systemd is available) name `smartems-server.service`
- The default configuration sets the log file at `/var/log/smartems/smartems.log`
- The default configuration specifies an sqlite3 db at `/var/lib/smartems/smartems.db`
- Installs HTML/JS/CSS and other smartEMS files at `/usr/share/smartems`

## Start the server (init.d service)

Start smartEMS by running:

```bash
sudo service smartems-server start
```

This will start the `smartems-server` process as the `smartems` user,
which was created during the package installation. The default HTTP port
is `3000` and default user and group is `admin`.

Default login and password `admin`/ `admin`

To configure the smartEMS server to start at boot time:

```bash
sudo update-rc.d smartems-server defaults
```

## Start the server (via systemd)

To start the service using systemd:

```bash
systemctl daemon-reload
systemctl start smartems-server
systemctl status smartems-server
```

Enable the systemd service so that smartEMS starts at boot.

```bash
sudo systemctl enable smartems-server.service
```

## Environment file

The systemd service file and init.d script both use the file located at
`/etc/default/smartems-server` for environment variables used when
starting the back-end. Here you can override log directory, data
directory and other variables.

### Logging

By default smartEMS will log to `/var/log/smartems`

### Database

The default configuration specifies a sqlite3 database located at
`/var/lib/smartems/smartems.db`. Please backup this database before
upgrades. You can also use MySQL or Postgres as the smartEMS database, as detailed on [the configuration page]({{< relref "configuration.md#database" >}}).

## Configuration

The configuration file is located at `/etc/smartems/smartems.ini`.  Go the
[Configuration]({{< relref "configuration.md" >}}) page for details on all
those options.

### Adding data sources

- [Graphite]({{< relref "../features/datasources/graphite.md" >}})
- [InfluxDB]({{< relref "../features/datasources/influxdb.md" >}})
- [OpenTSDB]({{< relref "../features/datasources/opentsdb.md" >}})
- [Prometheus]({{< relref "../features/datasources/prometheus.md" >}})

## Installing from binary tar file

Download [the latest `.tar.gz` file](https://smartems.com/get) and
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
