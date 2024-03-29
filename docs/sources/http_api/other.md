+++
title = "Other HTTP API "
description = "smartEMS Other HTTP API"
keywords = ["smartems", "http", "documentation", "api", "other"]
aliases = ["/http_api/other/"]
type = "docs"
[menu.docs]
name = "Other"
parent = "http_api"
+++


# Frontend Settings API

## Get Settings

`GET /api/frontend/settings`

**Example Request**:

```http
GET /api/frontend/settings HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoiT0tTcG1pUlY2RnVKZTFVaDFsNFZXdE9ZWmNrMkZYbk
```

**Example Response**:

```http
HTTP/1.1 200
Content-Type: application/json

{
  "allowOrgCreate":true,
  "appSubUrl":"",
  "buildInfo":{
    "buildstamp":xxxxxx,
    "commit":"vyyyy",
    "version":"zzzzz"
  },
  "datasources":{
    "datasourcename":{
      "index":"smartems-dash",
      "meta":{
        "annotations":true,
        "module":"plugins/datasource/smartems/datasource",
        "name":"smartEMS",
        "partials":{
          "annotations":"app/plugins/datasource/smartems/partials/annotations.editor.html",
          "config":"app/plugins/datasource/smartems/partials/config.html"
        },
        "pluginType":"datasource",
        "serviceName":"smartEMS",
        "type":"smartemssearch"
      }
    }
  },
  "defaultDatasource": "smartEMS"
}
```

# Login API

## Renew session based on remember cookie

`GET /api/login/ping`

**Example Request**:

```http
GET /api/login/ping HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoiT0tTcG1pUlY2RnVKZTFVaDFsNFZXdE9ZWmNrMkZYbk
```

**Example Response**:

```http
HTTP/1.1 200
Content-Type: application/json

{"message": "Logged in"}
```

# Health API

## Returns health information about smartEMS

`GET /api/health`

**Example Request**

```http
GET /api/health
Accept: application/json
```

**Example Response**:

```http
HTTP/1.1 200 OK

{
  "commit": "087143285",
  "database": "ok",
  "version": "5.1.3"
}
```
