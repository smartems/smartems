import { configure } from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
import 'jquery';
import $ from 'jquery';

const global = window as any;
global.$ = global.jQuery = $;

import '../vendor/flot/jquery.flot';
import '../vendor/flot/jquery.flot.time';
import 'angular';
import angular from 'angular';

angular.module('smartems', ['ngRoute']);
angular.module('smartems.services', ['ngRoute', '$strap.directives']);
angular.module('smartems.panels', []);
angular.module('smartems.controllers', []);
angular.module('smartems.directives', []);
angular.module('smartems.filters', []);
angular.module('smartems.routes', ['ngRoute']);

jest.mock('app/core/core', () => ({}));
jest.mock('app/features/plugins/plugin_loader', () => ({}));

configure({ adapter: new Adapter() });

const localStorageMock = (() => {
  let store: any = {};
  return {
    getItem: (key: string) => {
      return store[key];
    },
    setItem: (key: string, value: any) => {
      store[key] = value.toString();
    },
    clear: () => {
      store = {};
    },
    removeItem: (key: string) => {
      delete store[key];
    },
  };
})();

global.localStorage = localStorageMock;

HTMLCanvasElement.prototype.getContext = jest.fn() as any;

const throwUnhandledRejections = () => {
  process.on('unhandledRejection', err => {
    throw err;
  });
};

throwUnhandledRejections();
