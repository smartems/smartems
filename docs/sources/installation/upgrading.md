+++
title = "Upgrading"
description = "Upgrading smartEMS guide"
keywords = ["smartems", "configuration", "documentation", "upgrade"]
type = "docs"
[menu.docs]
name = "Upgrading"
identifier = "upgrading"
parent = "installation"
weight = 10
+++

# Upgrading smartEMS

We recommend everyone to upgrade smartEMS often to stay up to date with the latest fixes and enhancements.
In order make this a reality smartEMS upgrades are backward compatible and the upgrade process is simple and quick.

Upgrading is generally always safe (between many minor and one major version) and dashboards and graphs will look the same. There can be minor breaking changes in some edge cases which are usually outlined in the [Release Notes](https://community.smartems.com/c/releases) and [Changelog](https://github.com/smartems/smartems/blob/master/CHANGELOG.md)

## Update plugins

After you have upgraded it is highly recommended that you update all your plugins as a new version of smartEMS
can make older plugins stop working properly.

You can update all plugins using

```bash
smartems-cli plugins update-all
```

## Database backup

Before upgrading it can be a good idea to backup your smartEMS database. This will ensure that you can always rollback to your previous version. During startup, smartEMS will automatically migrate the database schema (if there are changes or new tables). Sometimes this can cause issues if you later want to downgrade.

#### sqlite

If you use sqlite you only need to make a backup of your `smartems.db` file. This is usually located at `/var/lib/smartems/smartems.db` on Unix systems.
If you are unsure what database you use and where it is stored check you smartems configuration file. If you
installed smartems to custom location using a binary tar/zip it is usually in `<smartems_install_dir>/data`.

#### mysql

```bash
backup:
> mysqldump -u root -p[root_password] [smartems] > smartems_backup.sql

restore:
> mysql -u root -p smartems < smartems_backup.sql
```

#### postgres

```bash
backup:
> pg_dump smartems > smartems_backup

restore:
> psql smartems < smartems_backup
```

### Ubuntu / Debian

If you installed smartems by downloading a debian package (`.deb`) you can just follow the same installation guide
and execute the same `dpkg -i` command but with the new package. It will upgrade your smartEMS install.

If you used our APT repository:

```bash
sudo apt-get update
sudo apt-get install smartems
```

#### Upgrading from binary tar file

If you downloaded the binary tar package you can just download and extract a new package
and overwrite all your existing files. But this might overwrite your config changes. We
recommend you place your config changes in a file named `<smartems_install_dir>/conf/custom.ini`
as this will make upgrades easier without risking losing your config changes.

### Centos / RHEL

If you installed smartems by downloading a rpm package you can just follow the same installation guide
and execute the same `yum install` or `rpm -i` command but with the new package. It will upgrade your smartEMS install.

If you used our YUM repository:

```bash
sudo yum update smartems
```

### Docker

This just an example, details depend on how you configured your smartems container.

```bash
docker pull smartems
docker stop my-smartems-container
docker rm my-smartems-container
docker run --name=my-smartems-container --restart=always -v /var/lib/smartems:/var/lib/smartems
```

### Windows

If you downloaded the Windows binary package you can just download a newer package and extract
to the same location (and overwrite the existing files). This might overwrite your config changes. We
recommend you place your config changes in a file named `<smartems_install_dir>/conf/custom.ini`
as this will make upgrades easier without risking losing your config changes.

## Upgrading from 1.x

[Migrating from 1.x to 2.x]({{< relref "installation/migrating_to2.md" >}})

## Upgrading from 2.x

We are not aware of any issues upgrading directly from 2.x to 4.x but to be on the safe side go via 3.x => 4.x.

## Upgrading to v5.0

The dashboard grid layout engine has changed. All dashboards will be automatically upgraded to new
positioning system when you load them in v5. Dashboards saved in v5 will not work in older versions of smartEMS. Some
external panel plugins might need to be updated to work properly.

For more details on the new panel positioning system, [click here]({{< relref "reference/dashboard.md#panel-size-position" >}})

## Upgrading to v5.2

One of the database migrations included in this release will update all annotation timestamps from second to millisecond precision. If you have a large amount of annotations the database migration may take a long time to complete which may cause problems if you use systemd to run smartEMS.

We've got one report where using systemd, PostgreSQL and a large amount of annotations (table size 1645mb) took 8-20 minutes for the database migration to complete. However, the smartems-server process was killed after 90 seconds by systemd. Any database migration queries in progress when systemd kills the smartems-server process continues to execute in database until finished.

If you're using systemd and have a large amount of annotations consider temporary adjusting the systemd `TimeoutStartSec` setting to something high like `30m` before upgrading.

## Upgrading to v6.0

If you have text panels with script tags they will no longer work due to a new setting that per default disallow unsanitized HTML.
Read more [here](/installation/configuration/#disable-sanitize-html) about this new setting.


### Authentication and security

If your using smartEMS's builtin, LDAP (without Auth Proxy) or OAuth authentication all users will be required to login upon the next visit after the upgrade.

If you have `cookie_secure` set to `true` in the `session` section you probably want to change the `cookie_secure` to `true` in the `security` section as well. Ending up with a configuration like this:

```ini
[session]
cookie_secure = true

[security]
cookie_secure = true
```

The `login_remember_days`, `cookie_username` and `cookie_remember_name` settings in the `security` section are no longer being used so they're safe to remove.

If you have `login_remember_days` configured to 0 (zero) you should change your configuration to this to accomplish similar behavior, i.e. a logged in user will maximum be logged in for 1 day until being forced to login again:

```ini
[auth]
login_maximum_inactive_lifetime_days = 1
login_maximum_lifetime_days = 1
```

The default cookie name for storing the auth token is `smartems_session`. you can configure this with `login_cookie_name` in `[auth]` settings.

## Upgrading to v6.2

### Ensure encryption of data source secrets

Data sources store passwords and basic auth passwords in secureJsonData encrypted (AES-256 in CFB mode) by default. Existing data source
will keep working with unencrypted passwords. If you want to migrate to encrypted storage for your existing data sources
you can do that by:

- For data sources created through UI, you need to go to data source config, re enter the password or basic auth
  password and save the data source.
- For data sources created by provisioning, you need to update your config file and use secureJsonData.password or
  secureJsonData.basicAuthPassword field. See [provisioning docs](/administration/provisioning) for example of current
  configuration.

### Embedding smartEMS

If you're embedding smartEMS in a `<frame>`, `<iframe>`, `<embed>` or `<object>` on a different website it will no longer work due to a new setting
that per default instructs the browser to not allow smartEMS to be embedded. Read more [here](/installation/configuration/#allow-embedding) about
this new setting.

### Session storage is no longer used

In 6.2 we completely removed the backend session storage since we replaced the previous login session implementation with an auth token.
If you are using Auth proxy with LDAP an shared cached is used in smartEMS so you might want configure [remote_cache] instead. If not
smartEMS will fallback to using the database as an shared cache.

### Upgrading Elasticsearch to v7.0+

The semantics of `max concurrent shard requests` changed in Elasticsearch v7.0, see [release notes](https://www.elastic.co/guide/en/elasticsearch/reference/7.0/breaking-changes-7.0.html#semantics-changed-max-concurrent-shared-requests) for reference.

If you upgrade Elasticsearch to v7.0+ you should make sure to update the data source configuration in smartEMS so that version
is `7.0+` and `max concurrent shard requests` properly configured. 256 was the default in pre v7.0 versions. In v7.0 and above 5 is the default.

## Upgrading to v6.4

### Annotations database migration

One of the database migrations included in this release will merge multiple rows used to represent an annotation range into a single row. If you have a large number of region annotations the database migration may take a long time to complete. See [Upgrading to v5.2](#upgrading-to-v5-2) for tips on how to manage this process.

### Docker

smartEMS’s docker image is now based on [Alpine](http://alpinelinux.org) instead of [Ubuntu](https://ubuntu.com/).

### Plugins that need updating

- [Splunk](https://smartems.com/smartems/plugins/smartems-splunk-datasource)

## Upgrading to v6.5

Pre smartEMS 6.5.0, the CloudWatch datasource used the GetMetricStatistics API for all queries that did not have an ´id´ and did not have an ´expression´ defined in the query editor. The GetMetricStatistics API has a limit of 400 transactions per second (TPS). In this release, all queries use the GetMetricData API which has a limit of 50 TPS and 100 metrics per transaction. We expect this transition to be smooth for most of our users, but in case you do face throttling issues we suggest you increase the TPS quota. To do that, please visit the [AWS Service Quotas console](https://console.aws.amazon.com/servicequotas/home?r#!/services/monitoring/quotas/L-5E141212). For more details around CloudWatch API limits, [see CloudWatch docs](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_limits.html).

Each request to the GetMetricData API can include 100 queries. This means that each panel in smartEMS will only issue one GetMetricData request, regardless of the number of query rows that are present in the panel. Consequently as it is no longer possible to set `HighRes` on a per query level anymore, this switch is now removed from the query editor. High resolution can still be achieved by choosing a smaller minimum period in the query editor.

The handling of multi template variables in dimension values has been changed in smartEMS 6.5. When a multi template variable is being used, smartEMS will generate a search expression. In the GetMetricData API, expressions are limited to 1024 characters, so it might be the case that this limit is reached when a multi template variable that has a lot of values is being used. If this is the case, we suggest you start using `*` wildcard as dimension value instead of a multi template variable.
