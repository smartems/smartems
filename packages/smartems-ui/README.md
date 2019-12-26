# smartEMS UI components library

> **@smartems/toolkit is currently in ALPHA**. Core API is unstable and can be a subject of breaking changes!

@smartems/ui is a collection of components used by [smartEMS](https://github.com/smartems/smartems)

Our goal is to deliver smartEMS's common UI elements for plugins developers and contributors.

See [package source](https://github.com/smartems/smartems/tree/master/packages/grafana-ui) for more details.

## Installation

`yarn add @smartems/ui`

`npm install @smartems/ui`

## Development

For development purposes we suggest using `yarn link` that will create symlink to @smartems/ui lib. To do so navigate to `packages/grafana-ui` and run `yarn link`. Then, navigate to your project and run `yarn link @smartems/ui` to use the linked version of the lib. To unlink follow the same procedure, but use `yarn unlink` instead.
