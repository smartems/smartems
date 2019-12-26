import { PanelPlugin } from '@smartems/data';
import { PieChartPanelEditor } from './PieChartPanelEditor';
import { PieChartPanel } from './PieChartPanel';
import { PieChartOptions, defaults } from './types';

export const plugin = new PanelPlugin<PieChartOptions>(PieChartPanel)
  .setDefaults(defaults)
  .setEditor(PieChartPanelEditor);
