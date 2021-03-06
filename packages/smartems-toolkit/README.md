
> **WARNING: @smartems/toolkit is currently in ALPHA**. The core API is unstable and can be a subject of breaking changes!

# smartems-toolkit
smartems-toolkit is a CLI that enables efficient development of smartEMS plugins. We want to help our community focus on the core value of their plugins rather than all the setup required to develop them.

## Getting started

Set up a new plugin with `smartems-toolkit plugin:create` command:

```sh
npx @smartems/toolkit plugin:create my-smartems-plugin
cd my-smartems-plugin
yarn install
yarn dev
```

## Update your plugin to use smartems-toolkit

Follow the steps below to start using smartems-toolkit in your existing plugin.

1. Add `@smartems/toolkit` package to your project by running `yarn add @smartems/toolkit` or `npm install @smartems/toolkit`.
2. Create `tsconfig.json` file in the root dir of your plugin and paste the code below:

```json
{
  "extends": "./node_modules/@smartems/toolkit/src/config/tsconfig.plugin.json",
  "include": ["src", "types"],
  "compilerOptions": {
    "rootDir": "./src",
    "baseUrl": "./src",
    "typeRoots": ["./node_modules/@types"]
  }
}
```

3. Create `.prettierrc.js` file in the root dir of your plugin and paste the code below:
```js
module.exports = {
  ...require("./node_modules/@smartems/toolkit/src/config/prettier.plugin.config.json"),
};
```

4. In your `package.json` file add following scripts:
```json
"scripts": {
  "build": "smartems-toolkit plugin:build",
  "test": "smartems-toolkit plugin:test",
  "dev": "smartems-toolkit plugin:dev",
  "watch": "smartems-toolkit plugin:dev --watch"
},
```

## Usage
With smartems-toolkit, we give you a  CLI that addresses common tasks performed when working on smartEMS plugin:

- `smartems-toolkit plugin:create`
- `smartems-toolkit plugin:dev`
- `smartems-toolkit plugin:test`
- `smartems-toolkit plugin:build`

### Create your plugin
`smartems-toolkit plugin:create plugin-name`

This command creates a new smartEMS plugin from template.

If `plugin-name` is provided, then the template is downloaded to `./plugin-name` directory. Otherwise, it will be downloaded to the current directory.

### Develop your plugin
`smartems-toolkit plugin:dev`

This command creates a development build that's easy to play with and debug using common browser tooling.

Available options:
- `-w`, `--watch` - run development task in a watch mode

### Test your plugin
`smartems-toolkit plugin:test`

This command runs Jest against your codebase.

Available options:
- `--watch` - Runs tests in interactive watch mode.
- `--coverage` - Reports code coverage.
- `-u`, `--updateSnapshot` - Performs snapshots update.
- `--testNamePattern=<regex>` - Runs test with names that match provided regex (https://jestjs.io/docs/en/cli#testnamepattern-regex).
- `--testPathPattern=<regex>` - Runs test with paths that match provided regex (https://jestjs.io/docs/en/cli#testpathpattern-regex).


### Build your plugin
`smartems-toolkit plugin:build`

This command creates a production-ready build of your plugin.

## FAQ

### Which version of smartems-toolkit should I use?
See [smartEMS packages versioning guide](https://github.com/smartems/smartems/blob/master/packages/README.md#versioning).

### What tools does smartems-toolkit use?
smartems-toolkit comes with Typescript, TSLint, Prettier, Jest, CSS and SASS support.

### How to start using smartems-toolkit in my plugin?
See [Updating your plugin to use smartems-toolkit](#updating-your-plugin-to-use-smartems-toolkit).

### Can I use Typescript to develop smartEMS plugins?
Yes! smartems-toolkit supports Typescript by default.

### How can I test my plugin?
smartems-toolkit comes with Jest as a test runner.

Internally at smartEMS we use Enzyme. If you are developing React plugin and you want to configure Enzyme as a testing utility, then you need to configure `enzyme-adapter-react`. To do so, create `<YOUR_PLUGIN_DIR>/config/jest-setup.ts` file that will provide necessary setup. Copy the following code into that file to get Enzyme working with React:

```ts
import { configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';

configure({ adapter: new Adapter() });
```

You can also set up Jest with shims of your needs by creating `jest-shim.ts` file in the same directory: `<YOUR_PLUGIN_DIR_>/config/jest-shim.ts`

### Can I provide custom setup for Jest?
You can provide custom Jest configuration with a `package.json` file. For more details, see [Jest docs](https://jest-bot.github.io/jest/docs/configuration.html).

Currently we support following Jest configuration properties:
- [`snapshotSerializers`](https://jest-bot.github.io/jest/docs/configuration.html#snapshotserializers-array-string)
- [`moduleNameMapper`](https://jestjs.io/docs/en/configuration#modulenamemapper-object-string-string)

### How can I style my plugin?
We support pure CSS, SASS, and CSS-in-JS approach (via [Emotion](https://emotion.sh/)).

#### Single CSS or SASS file
Create your CSS or SASS file and import it in your plugin entry point (typically `module.ts`):

```ts
import 'path/to/your/css_or_sass'
```
The styles will be injected via `style` tag during runtime.

> Note that imported static assets will be inlined as base64 URIs. *This can be subject of change in the future!*

#### Theme-specific stylesheets
If you want to provide different stylesheets for dark/light theme, then create `dark.[css|scss]` and `light.[css|scss]` files in the `src/styles` directory of your plugin. smartems-toolkit generates theme-specific stylesheets that are stored in `dist/styles` directory.

In order for smartEMS to pick up your theme stylesheets, you need to use `loadPluginCss` from `@smartems/runtime` package. Typically you would do that in the entry point of your plugin:

```ts
import { loadPluginCss } from '@smartems/runtime';

loadPluginCss({
  dark: 'plugins/<YOUR-PLUGIN-ID>/styles/dark.css',
  light: 'plugins/<YOUR-PLUGIN-ID>/styles/light.css',
});
```

You must add `@smartems/runtime` to your plugin dependencies by running `yarn add @smartems/runtime` or `npm instal @smartems/runtime`.

> Note that in this case static files (png, svg, json, html) are all copied to dist directory when the plugin is bundled. Relative paths to those files does not change!

#### Emotion
Starting from smartEMS 6.2 *our suggested way* for styling plugins is by using [Emotion](https://emotion.sh). It's a CSS-in-JS library that we use internally at smartEMS. The biggest advantage of using Emotion is that you can access smartEMS Theme variables.

To start using Emotion, you first must add it to your plugin dependencies:

```
  yarn add "@emotion/core"@10.0.14
```

Then, import `css` function from Emotion:

```ts
import { css } from 'emotion'
```

Now you are ready to implement your styles:

```tsx
const MyComponent = () => {
  return <div className={css`background: red;`} />
}
```
To learn more about using smartEMS theme please refer to [Theme usage guide](https://github.com/smartems/smartems/blob/master/style_guides/themes.md#react)

> We do not support Emotion's `css` prop. Use className instead!

### Can I adjust Typescript configuration to suit my needs?
Yes! However, it's important that your `tsconfig.json` file contains the following lines:

```json
{
  "extends": "./node_modules/@smartems/toolkit/src/config/tsconfig.plugin.json",
  "include": ["src"],
  "compilerOptions": {
    "rootDir": "./src",
    "typeRoots": ["./node_modules/@types"]
  }
}
```

### Can I adjust TSLint configuration to suit my needs?
smartems-toolkit comes with [default config for TSLint](https://github.com/smartems/smartems/blob/master/packages/smartems-toolkit/src/config/tslint.plugin.json). For now, there is now way to customise TSLint config.

### How is Prettier integrated into smartems-toolkit workflow?
When building plugin with [`smartems-toolkit plugin:build`](#building-plugin) task, smartems-toolkit performs Prettier check. If the check detects any Prettier issues, the build will not pass. To avoid such situation we suggest developing plugin with [`smartems-toolkit plugin:dev --watch`](#developing-plugin) task running. This task tries to fix Prettier issues automatically.

### My editor does not respect Prettier config, what should I do?
In order for your editor to pick up our Prettier config you need to create `.prettierrc.js` file in the root directory of your plugin with following content:

```js
module.exports = {
  ...require("./node_modules/@smartems/toolkit/src/config/prettier.plugin.config.json"),
};
```

### How do I add 3rd party dependencies that are not npm packages?
You can add such dependencies by putting them in `static` directory in the root of your project. The `static` directory will be copied when building the plugin.

## Contribute to smartems-toolkit
You can contribute to smartems-toolkit in the by helping develop it or by debugging it.

### Develop smartems-toolkit
Typically plugins should be developed using the `@smartems/toolkit` installed from npm. However, when working on the toolkit, you might want to use the local version. Follow the steps below to develop with a local version:

1. Clone [smartEMS repository](https://github.com/smartems/smartems).
2. Navigate to the directory you have cloned smartEMS repo to and then run `yarn install --pure-lockfile`.
3. Navigate to `<SMARTEMS_DIR>/packages/smartems-toolkit` and then run `yarn link`.
2. Navigate to the directory where your plugin code is and then run `npx smartems-toolkit plugin:dev --yarnlink`. This adds all dependencies required by smartems-toolkit to your project, as well as link your local smartems-toolkit version to be used by the plugin.


### Debug smartems-toolkit
To debug smartems-toolkit you can use standard [NodeJS debugging methods](https://nodejs.org/de/docs/guides/debugging-getting-started/#enable-inspector) (`node --inspect`, `node --inspect-brk`).

To run smartems-toolkit in a debugging session use the following command in the toolkit's directory:

`node --inspect-brk ./bin/smartems-toolkit.js [task]`
