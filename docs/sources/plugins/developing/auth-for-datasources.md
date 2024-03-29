+++
title = "Authentication for Datasource Plugins"
type = "docs"
[menu.docs]
name = "Authentication for Datasource Plugins"
parent = "developing"
weight = 3
+++

# Authentication for data source plugins

smartEMS has a proxy feature that proxies all data requests through the smartEMS backend. This is very useful when your data source plugin calls an external/thirdy-party API. The smartEMS proxy adds CORS headers and can authenticate against the external API. This means that a data source plugin that proxies all requests via smartEMS can enable token authentication and the token will be renewed automatically for the user when it expires.

The plugin config page should save the API key/password to be encrypted (using the `secureJsonData` feature) and then when a request from the data source is made, the smartEMS Proxy will:

 1. decrypt the API key/password on the backend.
 2. carry out authentication and generate an OAuth token that will be added as an `Authorization` HTTP header to all requests (or it will add a HTTP header with the API key).
 3. renew the token if it expires.

This means that users that access the data source config page cannot access the API key or password after is saved the first time and that no secret keys are sent in plain text through the browser where they can be spied on.

For backend authentication to work, the external/third-party API must either have an OAuth endpoint or that the API accepts an API key as a HTTP header for authentication.

## Plugin Routes

You can specify routes in the `plugin.json` file for your data source plugin. [Here is an example](https://github.com/smartems/azure-monitor-datasource/blob/d74c82145c0a4af07a7e96cc8dde231bfd449bd9/src/plugin.json#L30-L95) with lots of routes (though most plugins will just have one route).

When you build your url to the third-party API in your data source class, the url should start with the text specified in the path field for a route. The proxy will strip out the path text and replace it with the value in the url field.

For example, if my code makes a call to url `azuremonitor/foo/bar` with this code:

```js
this.backendSrv.datasourceRequest({
  url: url,
  method: 'GET',
})
```

and this route:

```json
"routes": [{
  "path": "azuremonitor",
  "method": "GET",
  "url": "https://management.azure.com",
  ...
}]
```

then the smartEMS proxy will transform it into "https://management.azure.com/foo/bar" and add CORS headers.

The `method` parameter is optional. It can be set to any HTTP verb to provide more fine-grained control.

### Dynamic Routes

When using routes, you can also reference a variable stored in JsonData or SecureJsonData which will be interpolated when connecting to the data source.

With JsonData:
```json
"routes": [
  {
      "path": "custom/api/v5/*",
      "method": "*",
      "url": "{{.JsonData.dynamicUrl}}",
      ...
  },
]
```

With SecureJsonData:
```json
"routes": [{
      "path": "custom/api/v5/*",
      "method": "*",
      "url": "{{.SecureJsonData.dynamicUrl}}",
  ...
}]
```

In the above example, the app is able to set the value for `dynamicUrl` in JsonData or SecureJsonData and it will be replaced on-demand.

An app using this feature can be found [here](https://github.com/smartems/kentik-app).

## Encrypting Sensitive Data

When a user saves a password or secret with your data source plugin's Config page, then you can save data to a column in the data source table called `secureJsonData` that is an encrypted blob. Any data saved in the blob is encrypted by smartEMS and can only be decrypted by the smartEMS server on the backend. This means once a password is saved, no sensitive data is sent to the browser. If the password is saved in the `jsonData` blob or the `password` field then it is unencrypted and anyone with Admin access (with the help of Chrome Developer Tools) can read it.

This is an example of using the `secureJsonData` blob to save a property called `password`:

```html
<input type="password" class="gf-form-input" ng-model='ctrl.current.secureJsonData.password' placeholder="password"></input>
```

## API Key/HTTP Header Authentication

Some third-party API's accept a HTTP Header for authentication. The [example](https://github.com/smartems/azure-monitor-datasource/blob/d74c82145c0a4af07a7e96cc8dde231bfd449bd9/src/plugin.json#L91-L93) below has a `headers` section that defines the name of the HTTP Header that the API expects and it uses the `SecureJSONData` blob to fetch an encrypted API key. The smartEMS server proxy will decrypt the key, add the `X-API-Key` header to the request and forward it to the third-party API.

```json
{
  "path": "appinsights",
  "method": "GET",
  "url": "https://api.applicationinsights.io",
  "headers": [
    {"name": "X-API-Key", "content": "{{.SecureJsonData.appInsightsApiKey}}"}
  ]
}
```

## How Token Authentication Works

The token auth section in the `plugin.json` file looks like this:

```json
"tokenAuth": {
  "url": "https://login.microsoftonline.com/{{.JsonData.tenantId}}/oauth2/token",
  "params": {
    "grant_type":  "client_credentials",
    "client_id": "{{.JsonData.clientId}}",
    "client_secret": "{{.SecureJsonData.clientSecret}}",
    "resource": "https://management.azure.com/"
  }
}
```

This interpolates in data from both `jsonData`  and `secureJsonData` to generate the token request to the third-party API. It is common for tokens to have a short expiry period (30 minutes). The proxy in smartEMS server will automatically renew the token if it has expired.

## Always Restart the smartEMS Server After Route Changes

The plugin.json files are only loaded when the smartEMS server starts so when a route is added or changed then the smartEMS server has to be restarted for the changes to take effect.
