+++
title = "Configuration"
description = "Configuration Docs"
keywords = ["smartems", "configuration", "documentation"]
type = "docs"
[menu.docs]
name = "Configuration"
identifier = "config"
parent = "admin"
weight = 1
+++

# Configuration

The smartEMS back-end has a number of configuration options that can be
specified in a `.ini` configuration file or specified using environment variables.

> **Note.** smartEMS needs to be restarted for any configuration changes to take effect.

## Comments In .ini Files

Semicolons (the `;` char) are the standard way to comment out lines in a `.ini` file.

A common problem is forgetting to uncomment a line in the `custom.ini` (or `smartems.ini`) file which causes the configuration option to be ignored.

## Config file locations

- Default configuration from `$WORKING_DIR/conf/defaults.ini`
- Custom configuration from `$WORKING_DIR/conf/custom.ini`
- The custom configuration file path can be overridden using the `--config` parameter

> **Note.** If you have installed smartEMS using the `deb` or `rpm`
> packages, then your configuration file is located at
> `/etc/smartems/smartems.ini` and a separate `custom.ini` is not
> used. This path is specified in the smartEMS
> init.d script using `--config` file parameter.

## Using environment variables

All options in the configuration file (listed below) can be overridden
using environment variables using the syntax:

```bash
GF_<SectionName>_<KeyName>
```

Where the section name is the text within the brackets. Everything
should be upper case, `.` should be replaced by `_`. For example, given these configuration settings:

```bash
# default section
instance_name = ${HOSTNAME}

[security]
admin_user = admin

[auth.google]
client_secret = 0ldS3cretKey
```

Then you can override them using:

```bash
export GF_DEFAULT_INSTANCE_NAME=my-instance
export GF_SECURITY_ADMIN_USER=true
export GF_AUTH_GOOGLE_CLIENT_SECRET=newS3cretKey
```

<hr />

## instance_name

Set the name of the smartems-server instance. Used in logging and internal metrics and in
clustering info. Defaults to: `${HOSTNAME}`, which will be replaced with
environment variable `HOSTNAME`, if that is empty or does not exist smartEMS will try to use
system calls to get the machine name.

## [paths]

### data

Path to where smartEMS stores the sqlite3 database (if used), file based
sessions (if used), and other data.  This path is usually specified via
command line in the init.d script or the systemd service file.

### temp_data_lifetime

How long temporary images in `data` directory should be kept. Defaults to: `24h`. Supported modifiers: `h` (hours),
`m` (minutes), for example: `168h`, `30m`, `10h30m`. Use `0` to never clean up temporary files.

### logs

Path to where smartEMS will store logs. This path is usually specified via
command line in the init.d script or the systemd service file.  It can
be overridden in the configuration file or in the default environment variable
file.

### plugins

Directory where smartems will automatically scan and look for plugins

### provisioning

Folder that contains [provisioning](/administration/provisioning) config files that smartems will apply on startup. Dashboards will be reloaded when the json files changes

## [server]

### http_addr

The IP address to bind to. If empty will bind to all interfaces

### http_port

The port to bind to, defaults to `3000`. To use port 80 you need to
either give the smartEMS binary permission for example:

```bash
$ sudo setcap 'cap_net_bind_service=+ep' /usr/sbin/smartems-server
```

Or redirect port 80 to the smartEMS port using:

```bash
$ sudo iptables -t nat -A PREROUTING -p tcp --dport 80 -j REDIRECT --to-port 3000
```

Another way is put a webserver like Nginx or Apache in front of smartEMS and have them proxy requests to smartEMS.

### protocol

`http`,`https`,`h2` or `socket`

> **Note** smartEMS versions earlier than 3.0 are vulnerable to [POODLE](https://en.wikipedia.org/wiki/POODLE). So we strongly recommend to upgrade to 3.x or use a reverse proxy for ssl termination.

### socket
Path where the socket should be created when `protocol=socket`. Please make sure that smartEMS has appropriate permissions.

### domain

This setting is only used in as a part of the `root_url` setting (see below). Important if you
use GitHub or Google OAuth.

### enforce_domain

Redirect to correct domain if host header does not match domain.
Prevents DNS rebinding attacks. Default is `false`.

### root_url

This is the full URL used to access smartEMS from a web browser. This is
important if you use Google or GitHub OAuth authentication (for the
callback URL to be correct).

> **Note** This setting is also important if you have a reverse proxy
> in front of smartEMS that exposes it through a subpath. In that
> case add the subpath to the end of this URL setting.

### serve_from_sub_path
> Available in 6.3 and above

Serve smartEMS from subpath specified in `root_url` setting. By
default it is set to `false` for compatibility reasons.

By enabling this setting and using a subpath in `root_url` above, e.g.
`root_url = http://localhost:3000/smartems`, smartEMS will be accessible on
`http://localhost:3000/smartems`.

### static_root_path

The path to the directory where the front end files (HTML, JS, and CSS
files). Default to `public` which is why the smartEMS binary needs to be
executed with working directory set to the installation path.

### enable_gzip

Set this option to `true` to enable HTTP compression, this can improve
transfer speed and bandwidth utilization. It is recommended that most
users set it to `true`. By default it is set to `false` for compatibility
reasons.

### cert_file

Path to the certificate file (if `protocol` is set to `https` or `h2`).

### cert_key

Path to the certificate key file (if `protocol` is set to `https` or `h2`).

### router_logging

Set to `true` for smartEMS to log all HTTP requests (not just errors). These are logged as Info level events
to smartems log.

<hr />

## [database]

smartEMS needs a database to store users and dashboards (and other
things). By default it is configured to use `sqlite3` which is an
embedded database (included in the main smartEMS binary).

### url

Use either URL or the other fields below to configure the database
Example: `mysql://user:secret@host:port/database`

### type

Either `mysql`, `postgres` or `sqlite3`, it's your choice.

### path

Only applicable for `sqlite3` database. The file path where the database
will be stored.

### host

Only applicable to MySQL or Postgres. Includes IP or hostname and port or in case of Unix sockets the path to it.
For example, for MySQL running on the same host as smartEMS: `host =
127.0.0.1:3306` or with Unix sockets: `host = /var/run/mysqld/mysqld.sock`

### name

The name of the smartEMS database. Leave it set to `smartems` or some
other name.

### user

The database user (not applicable for `sqlite3`).

### password

The database user's password (not applicable for `sqlite3`). If the password contains `#` or `;` you have to wrap it with triple quotes. For example `"""#password;"""`

### ssl_mode

For Postgres, use either `disable`, `require` or `verify-full`.
For MySQL, use either `true`, `false`, or `skip-verify`.

### ca_cert_path

The path to the CA certificate to use. On many Linux systems, certs can be found in `/etc/ssl/certs`.

### client_key_path

The path to the client key. Only if server requires client authentication.

### client_cert_path

The path to the client cert. Only if server requires client authentication.

### server_cert_name

The common name field of the certificate used by the `mysql` or `postgres` server. Not necessary if `ssl_mode` is set to `skip-verify`.

### max_idle_conn
The maximum number of connections in the idle connection pool.

### max_open_conn
The maximum number of open connections to the database.

### conn_max_lifetime

Sets the maximum amount of time a connection may be reused. The default is 14400 (which means 14400 seconds or 4 hours). For MySQL, this setting should be shorter than the [`wait_timeout`](https://dev.mysql.com/doc/refman/5.7/en/server-system-variables.html#sysvar_wait_timeout) variable.

### log_queries

Set to `true` to log the sql calls and execution times.

### cache_mode

For "sqlite3" only. [Shared cache](https://www.sqlite.org/sharedcache.html) setting used for connecting to the database. (private, shared)
Defaults to `private`.

<hr />

## [remote_cache]

### type

Either `redis`, `memcached` or `database`. Defaults to `database`

### connstr

The remote cache connection string. The format depends on the `type` of the remote cache.

#### Database

Leave empty when using `database` since it will use the primary database.

#### Redis

Example connstr: `addr=127.0.0.1:6379,pool_size=100,db=0,ssl=false`

- `addr` is the host `:` port of the redis server.
- `pool_size` (optional) is the number of underlying connections that can be made to redis.
- `db` (optional) is the number indentifer of the redis database you want to use.
- `ssl` (optional) is if SSL should be used to connect to redis server. The value may be `true`, `false`, or `insecure`. Setting the value to `insecure` skips verification of the certificate chain and hostname when making the connection.

#### Memcache

Example connstr: `127.0.0.1:11211`

<hr />

## [security]

### disable_initial_admin_creation

> Only available in smartEMS v6.5+.

Disable creation of admin user on first start of smartems.

### admin_user

The name of the default smartEMS admin user (who has full permissions).
Defaults to `admin`.

### admin_password

The password of the default smartEMS admin. Set once on first-run.  Defaults to `admin`.

### login_remember_days

The number of days the keep me logged in / remember me cookie lasts.

### secret_key

Used for signing some data source settings like secrets and passwords, the encryption format used is AES-256 in CFB mode. Cannot be changed without requiring an update
to data source settings to re-encode them.

### disable_gravatar

Set to `true` to disable the use of Gravatar for user profile images.
Default is `false`.

### data_source_proxy_whitelist

Define a white list of allowed ips/domains to use in data sources. Format: `ip_or_domain:port` separated by spaces.

### cookie_secure

Set to `true` if you host smartEMS behind HTTPS. Default is `false`.

### cookie_samesite

Sets the `SameSite` cookie attribute and prevents the browser from sending this cookie along with cross-site requests. The main goal is mitigate the risk of cross-origin information leakage. It also provides some protection against cross-site request forgery attacks (CSRF),  [read more here](https://www.owasp.org/index.php/SameSite). Valid values are `lax`, `strict` and `none`. Default is `lax`.

### allow_embedding

When `false`, the HTTP header `X-Frame-Options: deny` will be set in smartEMS HTTP responses which will instruct
browsers to not allow rendering smartEMS in a `<frame>`, `<iframe>`, `<embed>` or `<object>`. The main goal is to
mitigate the risk of [Clickjacking](https://www.owasp.org/index.php/Clickjacking). Default is `false`.

### strict_transport_security

Set to `true` if you want to enable HTTP `Strict-Transport-Security` (HSTS) response header. This is only sent when HTTPS is enabled in this configuration. HSTS tells browsers that the site should only be accessed using HTTPS. The default value is `false` until the next minor release, `6.3`.

### strict_transport_security_max_age_seconds

Sets how long a browser should cache HSTS in seconds. Only applied if strict_transport_security is enabled. The default value is `86400`.

### strict_transport_security_preload

Set to `true` if to enable HSTS `preloading` option. Only applied if strict_transport_security is enabled. The default value is `false`.

### strict_transport_security_subdomains

Set to `true` if to enable the HSTS includeSubDomains option. Only applied if strict_transport_security is enabled. The default value is `false`.

### x_content_type_options

Set to `true` to enable the X-Content-Type-Options response header. The X-Content-Type-Options response HTTP header is a marker used by the server to indicate that the MIME types advertised in the Content-Type headers should not be changed and be followed. The default value is `false` until the next minor release, `6.3`.

### x_xss_protection

Set to `false` to disable the X-XSS-Protection header, which tells browsers to stop pages from loading when they detect reflected cross-site scripting (XSS) attacks. The default value is `false` until the next minor release, `6.3`.

<hr />

## [users]

### allow_sign_up

Set to `false` to prohibit users from being able to sign up / create
user accounts. Defaults to `false`.  The admin user can still create
users from the [smartEMS Admin Pages](../../reference/admin)

### allow_org_create

Set to `false` to prohibit users from creating new organizations.
Defaults to `false`.

### auto_assign_org

Set to `true` to automatically add new users to the main organization
(id 1). When set to `false`, new users will automatically cause a new
organization to be created for that new user.

### auto_assign_org_id

Set this value to automatically add new users to the provided org.
This requires `auto_assign_org` to be set to `true`. Please make sure
that this organization does already exists.

### auto_assign_org_role

The role new users will be assigned for the main organization (if the
above setting is set to true).  Defaults to `Viewer`, other valid
options are `Admin` and `Editor`. e.g. :

`auto_assign_org_role = Viewer`

### viewers_can_edit

Viewers can edit/inspect dashboard settings in the browser, but not save the dashboard.
Defaults to `false`.

### editors_can_admin

Editors can administrate dashboards, folders and teams they create.
Defaults to `false`.

### login_hint

Text used as placeholder text on login page for login/username input.

### password_hint

Text used as placeholder text on login page for password input.

<hr>

## [auth]

smartEMS provides many ways to authenticate users. The docs for authentication has been split in to many different pages
below.

- [Authentication Overview]({{< relref "../auth/overview.md" >}}) (anonymous access options, hide login and more)
- [Google OAuth]({{< relref "../auth/google.md" >}}) (auth.google)
- [GitHub OAuth]({{< relref "../auth/github.md" >}}) (auth.github)
- [Gitlab OAuth]({{< relref "../auth/gitlab.md" >}}) (auth.gitlab)
- [Generic OAuth]({{< relref "../auth/generic-oauth.md" >}}) (auth.generic_oauth, okta2, auth0, bitbucket, azure)
- [Basic Authentication]({{< relref "../auth/overview.md" >}}) (auth.basic)
- [LDAP Authentication]({{< relref "../auth/ldap.md" >}}) (auth.ldap)
- [Auth Proxy]({{< relref "../auth/auth-proxy.md" >}}) (auth.proxy)

## [dataproxy]

### logging

This enables data proxy logging, default is `false`.

### timeout

How long the data proxy should wait before timing out. Default is `30` (seconds)

### send_user_header

If enabled and user is not anonymous, data proxy will add X-smartEMS-User header with username into the request. Default is `false`.

<hr />

## [analytics]

### reporting_enabled

When enabled smartEMS will send anonymous usage statistics to
`stats.smartems.org`. No IP addresses are being tracked, only simple counters to
track running instances, versions, dashboard and error counts. It is very helpful
to us, so please leave this enabled. Counters are sent every 24 hours. Default
value is `true`.

### google_analytics_ua_id

If you want to track smartEMS usage via Google analytics specify *your* Universal
Analytics ID here. By default this feature is disabled.

### check_for_updates

Set to false to disable all checks to https://smartems.com for new versions of installed plugins and to the smartEMS GitHub repository to check for a newer version of smartEMS. The version information is used in some UI views to notify that a new smartEMS update or a plugin update exists. This option does not cause any auto updates, nor send any sensitive information. The check is run every 10 minutes.

<hr />

## [dashboards]

### versions_to_keep

Number dashboard versions to keep (per dashboard). Default: `20`, Minimum: `1`.

## [dashboards.json]

> This have been replaced with dashboards [provisioning](/administration/provisioning) in 5.0+

### enabled
`true` or `false`. Is disabled by default.

### path
The full path to a directory containing your json dashboards.

## [smtp]
Email server settings.

### enabled
defaults to `false`

### host
defaults to `localhost:25`

### user
In case of SMTP auth, defaults to `empty`

### password
In case of SMTP auth, defaults to `empty`

### cert_file
File path to a cert file, defaults to `empty`

### key_file
File path to a key file, defaults to `empty`

### skip_verify
Verify SSL for smtp server? defaults to `false`

### from_address
Address used when sending out emails, defaults to `admin@smartems.localhost`

### from_name
Name to be used when sending out emails, defaults to `smartEMS`

### ehlo_identity
Name to be used as client identity for EHLO in SMTP dialog, defaults to instance_name.

## [log]

### mode
Either "console", "file", "syslog". Default is "console" and "file".
Use spaces to separate multiple modes, e.g. `console file`

### level
Either "debug", "info", "warn", "error", "critical", default is `info`

### filters
optional settings to set different levels for specific loggers.
For example `filters = sqlstore:debug`

## [metrics]

### enabled
Enable metrics reporting. defaults true. Available via HTTP API `/metrics`.

### basic_auth_username
If set configures the username to use for basic authentication on the metrics endpoint.

### basic_auth_password
If set configures the password to use for basic authentication on the metrics endpoint.

### disable_total_stats
If set to `true`, then total stats generation (`stat_totals_*` metrics) is disabled. The default is `false`.

### interval_seconds

Flush/Write interval when sending metrics to external TSDB. Defaults to 10s.

## [metrics.graphite]
Include this section if you want to send internal smartEMS metrics to Graphite.

### address
Format `<Hostname or ip>`:port

### prefix
Graphite metric prefix. Defaults to `prod.smartems.%(instance_name)s.`

## [snapshots]

### external_enabled
Set to `false` to disable external snapshot publish endpoint (default `true`)

### external_snapshot_url
Set root URL to a smartEMS instance where you want to publish external snapshots (defaults to https://snapshots-origin.raintank.io)

### external_snapshot_name
Set name for external snapshot button. Defaults to `Publish to snapshot.raintank.io`

### snapshot_remove_expired
Enabled to automatically remove expired snapshots

## [external_image_storage]
These options control how images should be made public so they can be shared on services like slack.

### provider
You can choose between (s3, webdav, gcs, azure_blob, local). If left empty smartEMS will ignore the upload action.

## [external_image_storage.s3]

### bucket
Bucket name for S3. e.g. smartems.snapshot

### region
Region name for S3. e.g. 'us-east-1', 'cn-north-1', etc

### path
Optional extra path inside bucket, useful to apply expiration policies

### bucket_url
(for backward compatibility, only works when no bucket or region are configured)
Bucket URL for S3. AWS region can be specified within URL or defaults to 'us-east-1', e.g.
- http://smartems.s3.amazonaws.com/
- https://smartems.s3-ap-southeast-2.amazonaws.com/

### access_key
Access key. e.g. AAAAAAAAAAAAAAAAAAAA

Access key requires permissions to the S3 bucket for the 's3:PutObject' and 's3:PutObjectAcl' actions.

### secret_key
Secret key. e.g. AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA

## [external_image_storage.webdav]

### url
Url to where smartEMS will send PUT request with images

### public_url
Optional parameter. Url to send to users in notifications. If the string contains the sequence ${file}, it will be replaced with the uploaded filename. Otherwise, the file name will be appended to the path part of the url, leaving any query string unchanged.

### username
basic auth username

### password
basic auth password

## [external_image_storage.gcs]

### key_file
Path to JSON key file associated with a Google service account to authenticate and authorize.
Service Account keys can be created and downloaded from https://console.developers.google.com/permissions/serviceaccounts.

Service Account should have "Storage Object Writer" role. The access control model of the bucket needs to be "Set object-level and bucket-level permissions". smartEMS itself will make the images public readable.

### bucket
Bucket Name on Google Cloud Storage.

### path
Optional extra path inside bucket

## [external_image_storage.azure_blob]

### account_name
Storage account name

### account_key
Storage account key

### container_name
Container name where to store "Blob" images with random names. Creating the blob container beforehand is required. Only public containers are supported.

## [alerting]

### enabled
Defaults to `true`. Set to `false` to disable alerting engine and hide Alerting from UI.

### execute_alerts

Makes it possible to turn off alert rule execution.

### error_or_timeout
> Available in 5.3 and above

Default setting for new alert rules. Defaults to categorize error and timeouts as alerting. (alerting, keep_state)

### nodata_or_nullvalues
> Available in 5.3  and above

Default setting for how smartEMS handles nodata or null values in alerting. (alerting, no_data, keep_state, ok)

### concurrent_render_limit

> Available in 5.3  and above

Alert notifications can include images, but rendering many images at the same time can overload the server.
This limit will protect the server from render overloading and make sure notifications are sent out quickly. Default
value is `5`.


### evaluation_timeout_seconds

Default setting for alert calculation timeout. Default value is `30`

### notification_timeout_seconds

Default setting for alert notification timeout. Default value is `30`

### max_attempts

Default setting for max attempts to sending alert notifications. Default value is `3`

## [rendering]

Options to configure a remote HTTP image rendering service, e.g. using https://github.com/smartems/smartems-image-renderer.

### server_url

URL to a remote HTTP image renderer service, e.g. http://localhost:8081/render, will enable smartEMS to render panels and dashboards to PNG-images using HTTP requests to an external service.

### callback_url

If the remote HTTP image renderer service runs on a different server than the smartEMS server you may have to configure this to a URL where smartEMS is reachable, e.g. http://smartems.domain/.

## [panels]

### disable_sanitize_html

If set to true smartEMS will allow script tags in text panels. Not recommended as it enable XSS vulnerabilities. Default
is false. This settings was introduced in smartEMS v6.0.

## [plugins]

### enable_alpha

Set to true if you want to test alpha plugins that are not yet ready for general usage.

## [feature_toggles]
### enable

Keys of alpha features to enable, separated by space. Available alpha features are: `transformations`

<hr />

# Removed options
Please note that these options have been removed.

## [session]
**Removed starting from smartEMS v6.2. Please use [remote_cache](#remote-cache) option instead.**

### provider

Valid values are `memory`, `file`, `mysql`, `postgres`, `memcache` or `redis`. Default is `file`.

### provider_config

This option should be configured differently depending on what type of
session provider you have configured.

- **file:** session file path, e.g. `data/sessions`
- **mysql:** go-sql-driver/mysql dsn config string, e.g. `user:password@tcp(127.0.0.1:3306)/database_name`
- **postgres:** ex:  `user=a password=b host=localhost port=5432 dbname=c sslmode=verify-full`
- **memcache:** ex:  `127.0.0.1:11211`
- **redis:** ex: `addr=127.0.0.1:6379,pool_size=100,prefix=smartems`. For Unix socket, use for example: `network=unix,addr=/var/run/redis/redis.sock,pool_size=100,db=smartems`

Postgres valid `sslmode` are `disable`, `require`, `verify-ca`, and `verify-full` (default).

### cookie_name

The name of the smartEMS session cookie.

### cookie_secure

Set to true if you host smartEMS behind HTTPS only. Defaults to `false`.

### session_life_time

How long sessions lasts in seconds. Defaults to `86400` (24 hours).

