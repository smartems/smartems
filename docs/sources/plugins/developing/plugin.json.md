+++
title = "plugin.json Schema"
keywords = ["smartems", "plugins", "documentation"]
type = "docs"
[menu.docs]
name = "plugin.json Schema"
parent = "developing"
weight = 8
+++

# Plugin.json

The plugin.json file is mandatory for all plugins. When smartEMS starts it will scan the plugin folders and mount every folder that contains a plugin.json file unless the folder contains a subfolder named `dist`. In that case smartems will mount the `dist` folder instead.

## Plugin JSON Schema

| Property | Description |
| ------------- |-------------|
| id | unique name of the plugin - [conventions described in styleguide]({{< relref "code-styleguide.md" >}}) |
| type | panel/datasource/app |
| name | Human readable name of the plugin |
| info.description | Description of plugin. Used for searching smartems.com plugins |
| info.author | |
| info.keywords | plugin keywords. Used for search on smartems net|
| info.logos | link to project logos |
| info.version | project version of this commit. Must be semver |
| dependencies.smartemsVersion | Required smartems backend version for this plugin |
| dependencies.plugins | required plugins for this plugin. |
