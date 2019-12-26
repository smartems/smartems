// const context = require.context('./', true, /_specs\.ts/);
// context.keys().forEach(context);
// module.exports = context;

import '@babel/polyfill';
import 'jquery';
import angular from 'angular';
import 'angular-mocks';
import 'app/app';

// configure enzyme
import Enzyme from 'enzyme';
import Adapter from 'enzyme-adapter-react-16';
Enzyme.configure({ adapter: new Adapter() });

angular.module('smartems', ['ngRoute']);
angular.module('smartems.services', ['ngRoute', '$strap.directives']);
angular.module('smartems.panels', []);
angular.module('smartems.controllers', []);
angular.module('smartems.directives', []);
angular.module('smartems.filters', []);
angular.module('smartems.routes', ['ngRoute']);

const context = (require as any).context('../', true, /specs\.(tsx?|js)/);
for (const key of context.keys()) {
  context(key);
}
