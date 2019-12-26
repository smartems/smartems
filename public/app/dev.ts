import app from './app';

/*
Import theme CSS based on env vars, e.g.: `env SMARTEMS_THEME=light yarn start`
*/
declare var SMARTEMS_THEME: any;
require('../sass/grafana.' + SMARTEMS_THEME + '.scss');

app.init();
