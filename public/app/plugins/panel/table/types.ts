import TableModel from 'app/core/table_model';
import { Column } from '@smartems/data';
import { ColumnStyle } from '@smartems/ui/src/components/Table/TableCellBuilder';

export interface TableTransform {
  description: string;
  getColumns(data?: any): any[];
  transform(data: any, panel: any, model: TableModel): void;
}

export interface ColumnRender extends Column {
  title: string;
  style: ColumnStyle;
  hidden: boolean;
}

export interface TableRenderModel {
  columns: ColumnRender[];
  rows: any[][];
}
