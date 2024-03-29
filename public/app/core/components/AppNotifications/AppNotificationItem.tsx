import React, { Component } from 'react';
import { AppNotification } from 'app/types';
import { Alert } from '@smartems/ui';

interface Props {
  appNotification: AppNotification;
  onClearNotification: (id: number) => void;
}

export default class AppNotificationItem extends Component<Props> {
  shouldComponentUpdate(nextProps: Props) {
    return this.props.appNotification.id !== nextProps.appNotification.id;
  }

  componentDidMount() {
    const { appNotification, onClearNotification } = this.props;
    setTimeout(() => {
      onClearNotification(appNotification.id);
    }, appNotification.timeout);
  }

  render() {
    const { appNotification, onClearNotification } = this.props;

    return (
      <Alert
        severity={appNotification.severity}
        title={appNotification.title}
        children={appNotification.component || appNotification.text}
        onRemove={() => onClearNotification(appNotification.id)}
      />
    );
  }
}
