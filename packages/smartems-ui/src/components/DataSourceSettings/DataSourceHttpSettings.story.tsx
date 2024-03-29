import React from 'react';
import { storiesOf } from '@storybook/react';
import { DataSourceHttpSettings } from './DataSourceHttpSettings';
import { DataSourceSettings } from '@smartems/data';
import { UseState } from '../../utils/storybook/UseState';

const settingsMock: DataSourceSettings<any, any> = {
  id: 4,
  orgId: 1,
  name: 'gdev-influxdb',
  type: 'influxdb',
  typeLogoUrl: '',
  access: 'direct',
  url: 'http://localhost:8086',
  password: '',
  user: 'smartems',
  database: 'site',
  basicAuth: false,
  basicAuthUser: '',
  basicAuthPassword: '',
  withCredentials: false,
  isDefault: false,
  jsonData: {
    timeInterval: '15s',
    httpMode: 'GET',
    keepCookies: ['cookie1', 'cookie2'],
  },
  secureJsonData: {
    password: true,
  },
  readOnly: true,
};

const DataSourceHttpSettingsStories = storiesOf('UI/DataSource/DataSourceHttpSettings', module);

DataSourceHttpSettingsStories.add('default', () => {
  return (
    <UseState initialState={settingsMock} logState>
      {(dataSourceSettings, updateDataSourceSettings) => {
        return (
          <DataSourceHttpSettings
            defaultUrl="http://localhost:9999"
            dataSourceConfig={dataSourceSettings}
            onChange={updateDataSourceSettings}
            showAccessOptions={true}
          />
        );
      }}
    </UseState>
  );
});
