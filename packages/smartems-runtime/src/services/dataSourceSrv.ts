import { ScopedVars, DataSourceApi } from '@smartems/data';

export interface DataSourceSrv {
  get(name?: string, scopedVars?: ScopedVars): Promise<DataSourceApi>;
}

let singletonInstance: DataSourceSrv;

export function setDataSourceSrv(instance: DataSourceSrv) {
  singletonInstance = instance;
}

export function getDataSourceSrv(): DataSourceSrv {
  return singletonInstance;
}
