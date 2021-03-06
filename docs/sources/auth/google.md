+++
title = "Google OAuth2 Authentication"
description = "smartEMS OAuthentication Guide "
keywords = ["smartems", "configuration", "documentation", "oauth"]
type = "docs"
[menu.docs]
name = "Google"
identifier = "ggogle_oauth2"
parent = "authentication"
weight = 3
+++

# Google OAuth2 Authentication

To enable the Google OAuth2 you must register your application with Google. Google will generate a client ID and secret key for you to use.

## Create Google OAuth keys

First, you need to create a Google OAuth Client:

1. Go to https://console.developers.google.com/apis/credentials
2. Click the 'Create Credentials' button, then click 'OAuth Client ID' in the menu that drops down
3. Enter the following:
   - Application Type: Web Application
   - Name: smartEMS
   - Authorized Javascript Origins: https://smartems.mycompany.com
   - Authorized Redirect URLs: https://smartems.mycompany.com/login/google
   - Replace https://smartems.mycompany.com with the URL of your smartEMS instance.
4. Click Create
5. Copy the Client ID and Client Secret from the 'OAuth Client' modal

## Enable Google OAuth in smartEMS

Specify the Client ID and Secret in the [smartEMS configuration file]({{< relref "installation/configuration.md#config-file-locations" >}}). For example:

```bash
[auth.google]
enabled = true
client_id = CLIENT_ID
client_secret = CLIENT_SECRET
scopes = https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/userinfo.email
auth_url = https://accounts.google.com/o/oauth2/auth
token_url = https://accounts.google.com/o/oauth2/token
allowed_domains = mycompany.com mycompany.org
allow_sign_up = true
```

You may have to set the `root_url` option of `[server]` for the callback URL to be 
correct. For example in case you are serving smartEMS behind a proxy.

Restart the smartEMS back-end. You should now see a Google login button
on the login page. You can now login or sign up with your Google
accounts. The `allowed_domains` option is optional, and domains were separated by space.

You may allow users to sign-up via Google authentication by setting the
`allow_sign_up` option to `true`. When this option is set to `true`, any
user successfully authenticating via Google authentication will be
automatically signed up.
