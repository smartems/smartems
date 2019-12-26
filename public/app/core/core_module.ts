import angular from 'angular';

const coreModule = angular.module('smartems.core', ['ngRoute']);

// legacy modules
const angularModules = [
  coreModule,
  angular.module('smartems.controllers', []),
  angular.module('smartems.directives', []),
  angular.module('smartems.factories', []),
  angular.module('smartems.services', []),
  angular.module('smartems.filters', []),
  angular.module('smartems.routes', []),
];

export { angularModules, coreModule };

export default coreModule;
