# Plugin development 

This document is not meant as a complete guide for developing plugins but more as a changelog for changes in
smartEMS that can impact plugin development. Whenever you as a plugin author encounter an issue with your plugin after
upgrading smartEMS please check here before creating an issue. 

## Plugin development resources

- [smartEMS plugin developer guide](http://docs.smartems.org/plugins/developing/development/)
- [Webpack smartEMS plugin template project](https://github.com/CorpGlory/smartems-plugin-template-webpack)
- [Simple JSON datasource plugin](https://github.com/smartems/simple-json-datasource)

## Changes in smartEMS v4.6

This version of smartEMS has big changes that will impact a limited set of plugins. We moved from systemjs to webpack
for built-in plugins and everything internal. External plugins still use systemjs but now with a limited 
set of smartEMS components they can import. Plugins can depend on libs like lodash & moment and internal components 
like before using the same import paths. However since everything in smartEMS is no longer accessible, a few plugins could encounter issues when importing a smartEMS dependency. 

[List of exposed components plugins can import/require](https://github.com/smartems/smartems/blob/master/public/app/features/plugins/plugin_loader.ts#L48)

If you think we missed exposing a crucial lib or smartEMS component let us know by opening an issue.  

### Deprecated components 

The angular directive `<spectrum-picker>` is now deprecated (will still work for a version more) but we recommend plugin authors
upgrade to new `<color-picker color="ctrl.color" onChange="ctrl.onSparklineColorChange"></color-picker>`

## Changes in smartEMS v6.0

### DashboardSrv.ts

If you utilize [DashboardSrv](https://github.com/smartems/smartems/commit/8574dca081002f36e482b572517d8f05fd44453f#diff-1ab99561f9f6a10e1fafcddc39bc1d65) in your plugin code, `dash` was renamed to `dashboard`.
