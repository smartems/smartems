+++
title = "Installing Plugins"
type = "docs"
[menu.docs]
parent = "plugins"
weight = 1
+++

# smartEMS Plugins

From smartEMS 3.0+ not only are data source plugins supported but also panel plugins and apps.
Having panels as plugins make it easy to create and add any kind of panel, to show your data
or improve your favorite dashboards. Apps is something new in smartEMS that enables
bundling of data sources, panels, dashboards and smartEMS pages into a cohesive experience.

smartEMS already have a strong community of contributors and plugin developers.
By making it easier to develop and install plugins we hope that the community
can grow even stronger and develop new plugins that we would never think about.

To discover plugins checkout the official [Plugin Repository](https://smartems.com/plugins).

# Installing Plugins

The easiest way to install plugins is by using the CLI tool smartems-cli which is bundled with smartems. Before any modification take place after modifying plugins, smartems-server needs to be restarted.

### smartEMS Plugin Directory

On Linux systems the smartems-cli will assume that the smartems plugin directory is `/var/lib/smartems/plugins`. It's possible to override the directory which smartems-cli will operate on by specifying the --pluginsDir flag. On Windows systems this parameter have to be specified for every call.

### smartEMS-cli Commands

List available plugins
```bash
smartems-cli plugins list-remote
```

Install the latest version of a plugin
```bash
smartems-cli plugins install <plugin-id>
```

Install a specific version of a plugin
```bash
smartems-cli plugins install <plugin-id> <version>
```

List installed plugins
```bash
smartems-cli plugins ls
```

Update all installed plugins
```bash
smartems-cli plugins update-all
```

Update one plugin
```bash
smartems-cli plugins update <plugin-id>
```

Remove one plugin
```bash
smartems-cli plugins remove <plugin-id>
```

### Installing Plugins Manually

If your smartEMS Server does not have access to the Internet, then the plugin will have to downloaded and manually copied to your smartEMS Server.

The Download URL from smartEMS.com API is in this form:

`https://smartems.com/api/plugins/<plugin id>/versions/<version number>/download`

You can specify a local URL by using the `--pluginUrl` option.
```bash
smartems-cli --pluginUrl https://nexus.company.com/smartems/plugins/<plugin-id>-<plugin-version>.zip plugins install <plugin-id>
```

To manually install a Plugin via the smartEMS.com API:

1. Find the plugin you want to download, the plugin id can be found on the Installation Tab on the plugin's page on smartEMS.com. In this example, the plugin id is `jdbranham-diagram-panel`:

    {{< imgbox img="/img/docs/installation-tab.png" caption="Installation Tab" >}}

2. Use the smartEMS API to find the plugin using this url `https://smartems.com/api/plugins/<plugin id from step 1>`. For example: https://smartems.com/api/plugins/jdbranham-diagram-panel should return:
    ```bash
    {
      "id": 145,
      "typeId": 3,
      "typeName": "Panel",
      "typeCode": "panel",
      "slug": "jdbranham-diagram-panel",
      "name": "Diagram",
      "description": "Diagram panel for smartems",
    ...
    ```

3. Find the download link:
    ```bash
    {
       "rel": "download",
       "href": "/plugins/jdbranham-diagram-panel/versions/1.4.0/download"
    }
    ```

4. Download the plugin with `https://smartems.com/api/plugins/<plugin id from step 1>/versions/<current version>/download` (for example: https://smartems.com/api/plugins/jdbranham-diagram-panel/versions/1.4.0/download). Unzip the downloaded file into the smartEMS Server's `plugins` directory.

5. Restart the smartEMS Server.
