import './query_parameter_ctrl';
import { DataSourcePlugin } from '@smartems/data';
import { ConfigEditor } from './components/ConfigEditor';
import { QueryEditor } from './components/QueryEditor';
import CloudWatchDatasource from './datasource';
import { CloudWatchJsonData, CloudWatchQuery } from './types';

class CloudWatchAnnotationsQueryCtrl {
  static templateUrl = 'partials/annotations.editor.html';
}

export const plugin = new DataSourcePlugin<CloudWatchDatasource, CloudWatchQuery, CloudWatchJsonData>(
  CloudWatchDatasource
)
  .setConfigEditor(ConfigEditor)
  .setQueryEditor(QueryEditor)
  .setExploreQueryField(QueryEditor)
  .setAnnotationQueryCtrl(CloudWatchAnnotationsQueryCtrl);
