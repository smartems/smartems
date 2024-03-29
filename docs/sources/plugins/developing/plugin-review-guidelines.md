+++
title = "Plugin Review Guidelines"
type = "docs"
[menu.docs]
name = "Plugin Review Guidelines"
parent = "developing"
weight = 2
+++

# Plugin Review Guidelines

The smartEMS team reviews all plugins that are published on smartEMS.com. There are two areas we review, the metadata for the plugin and the plugin functionality.

## Metadata

The plugin metadata consists of a `plugin.json` file and the README.md file. These `plugin.json` file is used by smartEMS to load the plugin and the README.md file is shown in the plugins section of smartEMS and the plugins section of smartEMS.com.

### README.md

The README.md file is shown on the plugins page in smartEMS and the plugin page on smartEMS.com. There are some differences between the GitHub markdown and the markdown allowed in smartEMS/smartEMS.com:

- Cannot contain inline HTML.
- Any image links should be absolute links. For example: https://raw.githubusercontent.com/smartems/azure-monitor-datasource/master/dist/img/smartems_cloud_install.png

The README should:

- describe the purpose of the plugin.
- contain steps on how to get started.

### Plugin.json

The `plugin.json` file is the same concept as the `package.json` file for an npm package. When the smartEMS server starts it will scan the plugin folders (all folders in the data/plugins subfolder) and load every folder that contains a `plugin.json` file unless the folder contains a subfolder named `dist`. In that case, the smartEMS server will load the `dist` folder instead.

A minimal `plugin.json` file:

```json
{
  "type": "panel",
  "name": "Clock",
  "id": "yourorg-clock-panel",

  "info": {
    "description": "Clock panel for smartems",
    "author": {
      "name": "Author Name",
      "url": "http://yourwebsite.com"
    },
    "keywords": ["clock", "panel"],
    "version": "1.0.0",
    "updated": "2018-03-24"
  },

  "dependencies": {
    "smartemsVersion": "3.x.x",
    "plugins": [ ]
  }
}
```

- The convention for the plugin id is [github username/org]-[plugin name]-[datasource|app|panel] and it has to be unique. Although if org and plugin name are the same then [plugin name]-[datasource|app|panel] is also valid. The org **cannot** be `smartems` unless it is a plugin created by the smartEMS core team.

    Examples:

    - raintank-worldping-app
    - ryantxu-ajax-panel
    - alexanderzobnin-zabbix-app
    - hawkular-datasource

- The `type` field should be either `datasource` `app` or `panel`.
- The `version` field should be in the form: x.x.x e.g. `1.0.0` or `0.4.1`.

The full file format for the `plugin.json` file is described [here](http://docs.smartems.org/plugins/developing/plugin.json/).

## Plugin Language

JavaScript, TypeScript, ES6 (or any other language) are all fine as long as the contents of the `dist` subdirectory are transpiled to JavaScript (ES5).

## File and Directory Structure Conventions

Here is a typical directory structure for a plugin.

```bash
johnnyb-awesome-datasource
|-- dist
|-- src
|   |-- img
|   |   |-- logo.svg
|   |-- partials
|   |   |-- annotations.editor.html
|   |   |-- config.html
|   |   |-- query.editor.html
|   |-- datasource.js
|   |-- module.js
|   |-- plugin.json
|   |-- query_ctrl.js
|-- Gruntfile.js
|-- LICENSE
|-- package.json
|-- README.md
```

Most JavaScript projects have a build step. The generated JavaScript should be placed in the `dist` directory and the source code in the `src` directory. We recommend that the plugin.json file be placed in the src directory and then copied over to the dist directory when building. The `README.md` can be placed in the root or in the dist directory.

Directories:

- `src/` contains plugin source files.
- `src/partials` contains html templates.
- `src/img` contains plugin logos and other images.
- `dist/` contains built content.

## HTML and CSS

For the HTML on editor tabs, we recommend using the inbuilt smartEMS styles rather than defining your own. This makes plugins feel like a more natural part of smartEMS. If done correctly, the html will also be responsive and adapt to smaller screens. The `gf-form` css classes should be used for labels and inputs.

Below is a minimal example of an editor row with one form group and two fields, a dropdown and a text input:

```html
<div class="editor-row">
  <div class="section gf-form-group">
    <h5 class="section-heading">My Plugin Options</h5>
    <div class="gf-form">
      <label class="gf-form-label width-10">Label1</label>
      <div class="gf-form-select-wrapper max-width-10">
        <select class="input-small gf-form-input" ng-model="ctrl.panel.mySelectProperty" ng-options="t for t in ['option1', 'option2', 'option3']" ng-change="ctrl.onSelectChange()"></select>
      </div>
      <div class="gf-form">
        <label class="gf-form-label width-10">Label2</label>
        <input type="text" class="input-small gf-form-input width-10" ng-model="ctrl.panel.myProperty" ng-change="ctrl.onFieldChange()" placeholder="suggestion for user" ng-model-onblur />
      </div>
    </div>
  </div>
</div>
```

Use the `width-x` and `max-width-x` classes to control the width of your labels and input fields. Try to get labels and input fields to line up neatly by having the same width for all the labels in a group and the same width for all inputs in a group if possible.

## Data Sources

A basic guide for data sources can be found [here](http://docs.smartems.org/plugins/developing/datasources/).

### Config Page Guidelines

- It should be as easy as possible for a user to configure a url. If the data source is using the `datasource-http-settings` component, it should use the `suggest-url` attribute to suggest the default url or a url that is similar to what it should be (especially important if the url refers to a REST endpoint that is not common knowledge for most users e.g. `https://yourserver:4000/api/custom-endpoint`).

    ```html
    <datasource-http-settings
      current="ctrl.current"
      suggest-url="http://localhost:8080">
    </datasource-http-settings>
    ```

- The `testDatasource` function should make a query to the data source that will also test that the authentication details are correct. This is so the data source is correctly configured when the user tries to write a query in a new dashboard.

#### Password Security

If possible, any passwords or secrets should be be saved in the `secureJsonData` blob. To encrypt sensitive data, the smartEMS server's proxy feature must be used. The smartEMS server has support for token authentication (OAuth) and HTTP Header authentication. If the calls have to be sent directly from the browser to a third-party API then this will not be possible and sensitive data will not be encrypted.

Read more here about how [authentication for data sources]({{< relref "auth-for-datasources.md" >}}) works.

If using the proxy feature then the Config page should use the `secureJsonData` blob like this:

  - good: `<input type="password" class="gf-form-input" ng-model='ctrl.current.secureJsonData.password' placeholder="password"></input>`
  - bad: `<input type="password" class="gf-form-input" ng-model='ctrl.current.password' placeholder="password"></input>`

### Query Editor

Each query editor is unique and can have a unique style. It should be adapted to what the users of the data source are used to.

- Should use the smartEMS CSS `gf-form` classes.
- Should be neat and tidy. Labels and fields in columns should be aligned and should be the same width if possible.
- The data source should be able to handle when a user toggles a query (by clicking on the eye icon) and not execute the query. This is done by checking the `hide` property - an [example](https://github.com/smartems/smartems/blob/master/public/app/plugins/datasource/postgres/datasource.ts#L35-L38).
- Should not execute queries if fields in the Query Editor are empty and the query will throw an exception (defensive programming).
- Should handle errors. There are two main ways to do this:
  - use the notification system in smartEMS to show a toaster popup with the error message. Example [here](https://github.com/alexanderzobnin/smartems-zabbix/blob/fdbbba2fb03f5f2a4b3b0715415e09d5a4cf6cde/src/panel-triggers/triggers_panel_ctrl.js#L467-L471).
  - provide an error notification in the query editor like the MySQL/Postgres data sources do. Example code in the `query_ctrl`  [here](https://github.com/smartems/azure-monitor-datasource/blob/b184d077f082a69f962120ef0d1f8296a0d46f03/src/query_ctrl.ts#L36-L51) and in the [html](https://github.com/smartems/azure-monitor-datasource/blob/b184d077f082a69f962120ef0d1f8296a0d46f03/src/partials/query.editor.html#L190-L193).
