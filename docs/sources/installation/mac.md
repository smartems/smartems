+++
title = "Installing on macOS"
description = "Installing smartEMS on macOS"
keywords = ["smartems", "configuration", "documentation", "mac", "homebrew", "osx"]
type = "docs"
[menu.docs]
parent = "installation"
weight = 4
+++


# Installing on macOS

## Install using homebrew

Installation can be done using [homebrew](http://brew.sh/)

Install latest stable:

```bash
brew update
brew install smartems
```

To start smartems look at the command printed after the homebrew install completes.

To upgrade use the reinstall command

```bash
brew update
brew reinstall smartems
```

-------------

You can also install the latest unstable smartems from git:


```bash
brew install --HEAD smartems/smartems/smartems
```

To upgrade smartems if you've installed from HEAD:

```bash
brew reinstall --HEAD smartems/smartems/smartems
```

### Starting smartEMS

To start smartEMS using homebrew services first make sure homebrew/services is installed.

```bash
brew tap homebrew/services
```

Then start smartEMS using:

```bash
brew services start smartems
```

Default login and password `admin`/ `admin`


### Configuration

The Configuration file should be located at `/usr/local/etc/smartems/smartems.ini`.

### Logs

The log file should be located at `/usr/local/var/log/smartems/smartems.log`.

### Plugins

If you want to manually install a plugin place it here: `/usr/local/var/lib/smartems/plugins`.

### Database

The default sqlite database is located at `/usr/local/var/lib/smartems`

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
