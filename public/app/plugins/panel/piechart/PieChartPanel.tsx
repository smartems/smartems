// Libraries
import React, { PureComponent } from 'react';

// Services & Utils
import { config } from 'app/core/config';

// Components
import { PieChart } from '@smartems/ui';
import { getFieldDisplayValues } from '@smartems/data';

// Types
import { PieChartOptions } from './types';
import { PanelProps } from '@smartems/data';

interface Props extends PanelProps<PieChartOptions> {}

export class PieChartPanel extends PureComponent<Props> {
  render() {
    const { width, height, options, data, replaceVariables } = this.props;

    const values = getFieldDisplayValues({
      fieldOptions: options.fieldOptions,
      data: data.series,
      theme: config.theme,
      replaceVariables: replaceVariables,
    }).map(v => v.display);

    return (
      <PieChart
        width={width}
        height={height}
        values={values}
        pieType={options.pieType}
        strokeWidth={options.strokeWidth}
        theme={config.theme}
      />
    );
  }
}
