import { sharedSingleStatPanelChangedHandler } from '@smartems/ui';
import { PanelPlugin } from '@smartems/data';
import { BarGaugePanel } from './BarGaugePanel';
import { BarGaugePanelEditor } from './BarGaugePanelEditor';
import { BarGaugeOptions, defaults } from './types';
import { barGaugePanelMigrationHandler } from './BarGaugeMigrations';

export const plugin = new PanelPlugin<BarGaugeOptions>(BarGaugePanel)
  .setDefaults(defaults)
  .setEditor(BarGaugePanelEditor)
  .setPanelChangeHandler(sharedSingleStatPanelChangedHandler)
  .setMigrationHandler(barGaugePanelMigrationHandler);
