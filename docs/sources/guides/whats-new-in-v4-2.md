+++
title = "What's New in smartEMS v4.2"
description = "Feature and improvement highlights for smartEMS v4.2"
keywords = ["smartems", "new", "documentation", "4.2.0"]
type = "docs"
[menu.docs]
name = "Version 4.2"
identifier = "v4.2"
parent = "whatsnew"
weight = -1
+++

## What's new in smartEMS v4.2

smartEMS v4.2 Beta is now [available for download](https://smartems.com/smartems/download/4.2.0).
Just like the last release this one contains lots bug fixes and minor improvements.
We are very happy to say that 27 of 40 issues was closed by pull requests from the community.
Big thumbs up!

## Release Highlights

- **Hipchat**: Adds support for sending alert notifications to hipchat [#6451](https://github.com/smartems/smartems/issues/6451), thx [@jregovic](https://github.com/jregovic)
- **Telegram**: Added Telegram alert notifier [#7098](https://github.com/smartems/smartems/pull/7098), thx [@leonoff](https://github.com/leonoff)
- **LINE**: Add LINE as alerting notification channel [#7301](https://github.com/smartems/smartems/pull/7301), thx [@huydx](https://github.com/huydx)
- **Templating**: Make $__interval and $__interval_ms global built in variables that can be used in by any data source (in panel queries), closes [#7190](https://github.com/smartems/smartems/issues/7190), closes [#6582](https://github.com/smartems/smartems/issues/6582)
- **Alerting**: Adds deduping of alert notifications [#7632](https://github.com/smartems/smartems/pull/7632)
- **Alerting**: Better information about why an alert triggered [#7035](https://github.com/smartems/smartems/issues/7035)
- **Orgs**: Sharing dashboards using smartEMS share feature will now redirect to correct org. [#6948](https://github.com/smartems/smartems/issues/6948)
- [Full changelog](https://github.com/smartems/smartems/blob/master/CHANGELOG.md)

### New alert notification channels

This release adds **five** new alert notifications channels, all of them contributed by the community.

* Hipchat
* Telegram
* LINE
* Pushover
* Threema

### Templating

We added two new global built in variables in smartems. `$__interval` and `$__interval_ms` are now reserved template names in smartems and can be used by any data source.
We might add more global built in variables in the future and if we do we will prefix them with `$__`. So please avoid using that in your template variables.

### Dedupe alert notifications when running multiple servers

In this release we will dedupe alert notifications when you are running multiple servers.
This makes it possible to run alerting on multiple servers and only get one notification.

We currently solve this with sql transactions which puts some limitations for how many servers you can use to execute the same rules.
3-5 servers should not be a problem but as always, it depends on how many alerts you have and how frequently they execute.

Next up for a better HA situation is to add support for workload balancing between smartEMS servers.

### Alerting more info

You can now see the reason why an alert triggered in the alert history. Its also easier to detect when an alert is set to `alerting` due to the `no_data` option.

### Improved support for multi-org setup

When loading dashboards we now set an query parameter called orgId. So we can detect from which org an user shared a dashboard.
This makes it possible for users to share dashboards between orgs without changing org first.

We aim to introduce [dashboard groups](https://github.com/smartems/smartems/issues/1611) sometime in the future which will introduce access control and user groups within one org.
Making it possible to have users in multiple groups and have detailed access control.

## Upgrade and Breaking changes

If you're using https in smartems we now force you to use tls 1.2 and the most secure ciphers.
We think its better to be secure by default rather then making it configurable.
If you want to run https with lower versions of tls we suggest you put a reserve proxy in front of smartems.

If you have template variables name `$__interval` or `$__interval_ms` they will no longer work since these keywords
are reserved as global built in variables. We might add more global built in variables in the future and if we do, we will prefix them with `$__`. So please avoid using that in your template variables.

## Changelog

Checkout the [CHANGELOG.md](https://github.com/smartems/smartems/blob/master/CHANGELOG.md) file for a complete list
of new features, changes, and bug fixes.

## Download

Head to [v4.2-beta download page](/download/4_2_0/) for download links and instructions.

## Thanks

A big thanks to all the smartEMS users who contribute by submitting PRs, bug reports and feedback!
