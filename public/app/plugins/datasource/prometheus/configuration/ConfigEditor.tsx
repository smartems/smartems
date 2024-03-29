import React from 'react';
import { DataSourceHttpSettings } from '@smartems/ui';
import { DataSourcePluginOptionsEditorProps } from '@smartems/data';
import { PromSettings } from './PromSettings';
import { PromOptions } from '../types';

export type Props = DataSourcePluginOptionsEditorProps<PromOptions>;
export const ConfigEditor = (props: Props) => {
  const { options, onOptionsChange } = props;
  return (
    <>
      <DataSourceHttpSettings
        defaultUrl="http://localhost:9090"
        dataSourceConfig={options}
        onChange={onOptionsChange}
      />

      <PromSettings value={options} onChange={onOptionsChange} />
    </>
  );
};
