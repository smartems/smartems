import React, { PureComponent } from 'react';
import { ExploreStartPageProps } from '@smartems/data';
import InfluxCheatSheet from './InfluxCheatSheet';

export default class InfluxStartPage extends PureComponent<ExploreStartPageProps> {
  render() {
    return <InfluxCheatSheet onClickExample={this.props.onClickExample} />;
  }
}
