import { PanelModel } from '@grafana/data';
import { sharedSingleStatMigrationHandler } from '@smartems/ui';
import { BarGaugeOptions } from './types';

export const barGaugePanelMigrationHandler = (panel: PanelModel<BarGaugeOptions>): Partial<BarGaugeOptions> => {
  return sharedSingleStatMigrationHandler(panel);
};
