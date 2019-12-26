import { DataQuery } from '@smartems/data';

export interface GraphiteQuery extends DataQuery {
  target?: string;
}

export enum GraphiteType {
  Default = 'default',
  Metrictank = 'metrictank',
}
