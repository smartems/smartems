+++
title = "Running smartEMS behind a reverse proxy"
description = "Guide for running smartEMS behind a reverse proxy"
keywords = ["smartems", "nginx", "documentation", "haproxy", "reverse"]
type = "docs"
[menu.docs]
name = "Running smartEMS behind a reverse proxy"
parent = "tutorials"
weight = 1
+++


# Running smartEMS behind a reverse proxy

It should be straight forward to get smartEMS up and running behind a reverse proxy. But here are some things that you might run into.

Links and redirects will not be rendered correctly unless you set the server.domain setting.
```bash
[server]
domain = foo.bar
```

To use sub *path* ex `http://foo.bar/smartems` make sure to include `/smartems` in the end of root_url.
Otherwise smartEMS will not behave correctly. See example below.

## Examples
Here are some example configurations for running smartEMS behind a reverse proxy.

### smartEMS configuration (ex http://foo.bar)

```bash
[server]
domain = foo.bar
```

### Nginx configuration

Nginx is a high performance load balancer, web server and reverse proxy: https://www.nginx.com/

```bash
server {
  listen 80;
  root /usr/share/nginx/www;
  index index.html index.htm;

  location / {
   proxy_pass http://localhost:3000/;
  }
}
```

### Examples with **sub path** (ex http://foo.bar/smartems)

#### smartEMS configuration with sub path
```bash
[server]
domain = foo.bar
root_url = %(protocol)s://%(domain)s/smartems/
```

#### Nginx configuration with sub path
```bash
server {
  listen 80;
  root /usr/share/nginx/www;
  index index.html index.htm;

  location /smartems/ {
   proxy_pass http://localhost:3000/;
  }
}
```

#### HAProxy configuration with sub path
```bash
frontend http-in
  bind *:80
  use_backend smartems_backend if { path /smartems } or { path_beg /smartems/ }

backend smartems_backend
  # Requires haproxy >= 1.6
  http-request set-path %[path,regsub(^/smartems/?,/)]

  # Works for haproxy < 1.6
  # reqrep ^([^\ ]*\ /)smartems[/]?(.*) \1\2

  server smartems localhost:3000
```

### IIS URL Rewrite Rule (Windows) with Subpath

IIS requires that the URL Rewrite module is installed.

Given:

- subpath `smartems`
- smartEMS installed on `http://localhost:3000`
- server config:

    ```bash
    [server]
    domain = localhost:8080
    root_url = %(protocol)s://%(domain)s/smartems/
    ```

Create an Inbound Rule for the parent website (localhost:8080 in this example) in IIS Manager with the following settings:

- pattern: `smartems(/)?(.*)`
- check the `Ignore case` checkbox
- rewrite url set to `http://localhost:3000/{R:2}`
- check the `Append query string` checkbox
- check the `Stop processing of subsequent rules` checkbox

This is the rewrite rule that is generated in the `web.config`:

```xml
  <rewrite>
      <rules>
          <rule name="smartEMS" enabled="true" stopProcessing="true">
              <match url="smartems(/)?(.*)" />
              <action type="Rewrite" url="http://localhost:3000/{R:2}" logRewrittenUrl="false" />
          </rule>
      </rules>
  </rewrite>
```

See the [tutorial on IIS Url Rewrites](http://docs.smartems.org/tutorials/iis/) for more in-depth instructions.
