// Libraries
import React, { Component } from 'react';

// Types
import { ThemeContext } from '@smartems/ui';
import { PanelProps } from '@smartems/data';
import { Options } from './types';
import Table from '@smartems/ui/src/components/Table/Table';

interface Props extends PanelProps<Options> {}

export class TablePanel extends Component<Props> {
  constructor(props: Props) {
    super(props);
  }

  render() {
    const { data, options } = this.props;

    if (data.series.length < 1) {
      return <div>No Table Data...</div>;
    }

    return (
      <ThemeContext.Consumer>
        {theme => <Table {...this.props} {...options} theme={theme} data={data.series[0]} />}
      </ThemeContext.Consumer>
    );
  }
}
