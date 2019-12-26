import { DataQuery } from '@smartems/data';

export const SHARED_DASHBODARD_QUERY = '-- Dashboard --';

export interface DashboardQuery extends DataQuery {
  panelId?: number;
}
