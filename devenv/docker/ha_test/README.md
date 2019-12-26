# smartEMS High Availability (HA) test setup

A set of docker compose services which together creates a smartEMS HA test setup with capability of easily
scaling up/down number of smartEMS instances.

Included services

* smartEMS
* Mysql - smartEMS configuration database and session storage
* Prometheus - Monitoring of smartEMS and used as data source of provisioned alert rules
* Nginx - Reverse proxy for smartEMS and Prometheus. Enables browsing smartEMS/Prometheus UI using a hostname

## Prerequisites

### Build smartems docker container

Build a smartEMS docker container from current branch and commit and tag it as smartems/smartems:dev.

```bash
$ cd <smartems repo>
$ make build-docker-full
```

### Virtual host names

#### Alternative 1 - Use dnsmasq

```bash
$ sudo apt-get install dnsmasq
$ echo 'address=/loc/127.0.0.1' | sudo tee /etc/dnsmasq.d/dnsmasq-loc.conf > /dev/null
$ sudo /etc/init.d/dnsmasq restart
$ ping whatever.loc
PING whatever.loc (127.0.0.1) 56(84) bytes of data.
64 bytes from localhost (127.0.0.1): icmp_seq=1 ttl=64 time=0.076 ms
--- whatever.loc ping statistics ---
1 packet transmitted, 1 received, 0% packet loss, time 1998ms
```

#### Alternative 2 - Manually update /etc/hosts

Update your `/etc/hosts` to be able to access smartEMS and/or Prometheus UI using a hostname.

```bash
$ cat /etc/hosts
127.0.0.1       smartems.loc
127.0.0.1       prometheus.loc
```

## Start services

```bash
$ docker-compose up -d
```

Browse
* http://smartems.loc/
* http://prometheus.loc/

Check for any errors

```bash
$ docker-compose logs | grep error
```

### Scale smartEMS instances up/down

Scale number of smartEMS instances to `<instances>`

```bash
$ docker-compose up --scale smartems=<instances> -d
# for example 3 instances
$ docker-compose up --scale smartems=3 -d
```

## Test alerting

### Create notification channels

Creates default notification channels, if not already exists

```bash
$ ./alerts.sh setup
```

### Slack notifications

Disable

```bash
$ ./alerts.sh slack -d
```

Enable and configure url

```bash
$ ./alerts.sh slack -u https://hooks.slack.com/services/...
```

Enable, configure url and enable reminders

```bash
$ ./alerts.sh slack -u https://hooks.slack.com/services/... -r -e 10m
```

### Provision alert dashboards with alert rules

Provision 1 dashboard/alert rule (default)

```bash
$ ./alerts.sh provision
```

Provision 10 dashboards/alert rules

```bash
$ ./alerts.sh provision -a 10
```

Provision 10 dashboards/alert rules and change condition to `gt > 100`

```bash
$ ./alerts.sh provision -a 10 -c 100
```

### Pause/unpause all alert rules

Pause

```bash
$ ./alerts.sh pause
```

Unpause

```bash
$ ./alerts.sh unpause
```
