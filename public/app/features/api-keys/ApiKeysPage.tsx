import React, { PureComponent } from 'react';
import ReactDOMServer from 'react-dom/server';
import { connect } from 'react-redux';
import { hot } from 'react-hot-loader';
import { ApiKey, NewApiKey, OrgRole } from 'app/types';
import { getNavModel } from 'app/core/selectors/navModel';
import { getApiKeys, getApiKeysCount } from './state/selectors';
import { loadApiKeys, deleteApiKey, setSearchQuery, addApiKey } from './state/actions';
import Page from 'app/core/components/Page/Page';
import { SlideDown } from 'app/core/components/Animations/SlideDown';
import ApiKeysAddedModal from './ApiKeysAddedModal';
import config from 'app/core/config';
import appEvents from 'app/core/app_events';
import EmptyListCTA from 'app/core/components/EmptyListCTA/EmptyListCTA';
import { DeleteButton, EventsWithValidation, FormLabel, Input, ValidationEvents } from '@smartems/ui';
import { NavModel, dateTime, isDateTime } from '@grafana/data';
import { FilterInput } from 'app/core/components/FilterInput/FilterInput';
import { store } from 'app/store/store';
import kbn from 'app/core/utils/kbn';

// Utils
import { CoreEvents } from 'app/types';
import { getTimeZone } from 'app/features/profile/state/selectors';

const timeRangeValidationEvents: ValidationEvents = {
  [EventsWithValidation.onBlur]: [
    {
      rule: value => {
        if (!value) {
          return true;
        }
        try {
          kbn.interval_to_seconds(value);
          return true;
        } catch {
          return false;
        }
      },
      errorMessage: 'Not a valid duration',
    },
  ],
};

export interface Props {
  navModel: NavModel;
  apiKeys: ApiKey[];
  searchQuery: string;
  hasFetched: boolean;
  loadApiKeys: typeof loadApiKeys;
  deleteApiKey: typeof deleteApiKey;
  setSearchQuery: typeof setSearchQuery;
  addApiKey: typeof addApiKey;
  apiKeysCount: number;
}

export interface State {
  isAdding: boolean;
  newApiKey: NewApiKey;
}

enum ApiKeyStateProps {
  Name = 'name',
  Role = 'role',
  SecondsToLive = 'secondsToLive',
}

const initialApiKeyState = {
  name: '',
  role: OrgRole.Viewer,
  secondsToLive: '',
};

const tooltipText =
  'The api key life duration. For example 1d if your key is going to last for one day. All the supported units are: s,m,h,d,w,M,y';

export class ApiKeysPage extends PureComponent<Props, any> {
  constructor(props: Props) {
    super(props);
    this.state = { isAdding: false, newApiKey: initialApiKeyState };
  }

  componentDidMount() {
    this.fetchApiKeys();
  }

  async fetchApiKeys() {
    await this.props.loadApiKeys();
  }

  onDeleteApiKey(key: ApiKey) {
    this.props.deleteApiKey(key.id);
  }

  onSearchQueryChange = (value: string) => {
    this.props.setSearchQuery(value);
  };

  onToggleAdding = () => {
    this.setState({ isAdding: !this.state.isAdding });
  };

  onAddApiKey = async (evt: any) => {
    evt.preventDefault();

    const openModal = (apiKey: string) => {
      const rootPath = window.location.origin + config.appSubUrl;
      const modalTemplate = ReactDOMServer.renderToString(<ApiKeysAddedModal apiKey={apiKey} rootPath={rootPath} />);

      appEvents.emit(CoreEvents.showModal, {
        templateHtml: modalTemplate,
      });
    };

    // make sure that secondsToLive is number or null
    const secondsToLive = this.state.newApiKey['secondsToLive'];
    this.state.newApiKey['secondsToLive'] = secondsToLive ? kbn.interval_to_seconds(secondsToLive) : null;
    this.props.addApiKey(this.state.newApiKey, openModal);
    this.setState((prevState: State) => {
      return {
        ...prevState,
        newApiKey: initialApiKeyState,
        isAdding: false,
      };
    });
  };

  onApiKeyStateUpdate = (evt: any, prop: string) => {
    const value = evt.currentTarget.value;
    this.setState((prevState: State) => {
      const newApiKey: any = {
        ...prevState.newApiKey,
      };
      newApiKey[prop] = value;

      return {
        ...prevState,
        newApiKey: newApiKey,
      };
    });
  };

  renderEmptyList() {
    const { isAdding } = this.state;
    return (
      <>
        {!isAdding && (
          <EmptyListCTA
            title="You haven't added any API Keys yet."
            buttonIcon="gicon gicon-apikeys"
            buttonLink="#"
            onClick={this.onToggleAdding}
            buttonTitle=" New API Key"
            proTip="Remember you can provide view-only API access to other applications."
          />
        )}
        {this.renderAddApiKeyForm()}
      </>
    );
  }

  formatDate(date: any, format?: string) {
    if (!date) {
      return 'No expiration date';
    }
    date = isDateTime(date) ? date : dateTime(date);
    format = format || 'YYYY-MM-DD HH:mm:ss';
    const timezone = getTimeZone(store.getState().user);

    return timezone === 'utc' ? date.utc().format(format) : date.format(format);
  }

  renderAddApiKeyForm() {
    const { newApiKey, isAdding } = this.state;

    return (
      <SlideDown in={isAdding}>
        <div className="cta-form">
          <button className="cta-form__close btn btn-transparent" onClick={this.onToggleAdding}>
            <i className="fa fa-close" />
          </button>
          <h5>Add API Key</h5>
          <form className="gf-form-group" onSubmit={this.onAddApiKey}>
            <div className="gf-form-inline">
              <div className="gf-form max-width-21">
                <span className="gf-form-label">Key name</span>
                <Input
                  type="text"
                  className="gf-form-input"
                  value={newApiKey.name}
                  placeholder="Name"
                  onChange={evt => this.onApiKeyStateUpdate(evt, ApiKeyStateProps.Name)}
                />
              </div>
              <div className="gf-form">
                <span className="gf-form-label">Role</span>
                <span className="gf-form-select-wrapper">
                  <select
                    className="gf-form-input gf-size-auto"
                    value={newApiKey.role}
                    onChange={evt => this.onApiKeyStateUpdate(evt, ApiKeyStateProps.Role)}
                  >
                    {Object.keys(OrgRole).map(role => {
                      return (
                        <option key={role} label={role} value={role}>
                          {role}
                        </option>
                      );
                    })}
                  </select>
                </span>
              </div>
              <div className="gf-form max-width-21">
                <FormLabel tooltip={tooltipText}>Time to live</FormLabel>
                <Input
                  type="text"
                  className="gf-form-input"
                  placeholder="1d"
                  validationEvents={timeRangeValidationEvents}
                  value={newApiKey.secondsToLive}
                  onChange={evt => this.onApiKeyStateUpdate(evt, ApiKeyStateProps.SecondsToLive)}
                />
              </div>
              <div className="gf-form">
                <button className="btn gf-form-btn btn-primary">Add</button>
              </div>
            </div>
          </form>
        </div>
      </SlideDown>
    );
  }

  renderApiKeyList() {
    const { isAdding } = this.state;
    const { apiKeys, searchQuery } = this.props;

    return (
      <>
        <div className="page-action-bar">
          <div className="gf-form gf-form--grow">
            <FilterInput
              labelClassName="gf-form--has-input-icon gf-form--grow"
              inputClassName="gf-form-input"
              placeholder="Search keys"
              value={searchQuery}
              onChange={this.onSearchQueryChange}
            />
          </div>

          <div className="page-action-bar__spacer" />
          <button className="btn btn-primary pull-right" onClick={this.onToggleAdding} disabled={isAdding}>
            Add API key
          </button>
        </div>

        {this.renderAddApiKeyForm()}

        <h3 className="page-heading">Existing Keys</h3>
        <table className="filter-table">
          <thead>
            <tr>
              <th>Name</th>
              <th>Role</th>
              <th>Expires</th>
              <th style={{ width: '34px' }} />
            </tr>
          </thead>
          {apiKeys.length > 0 ? (
            <tbody>
              {apiKeys.map(key => {
                return (
                  <tr key={key.id}>
                    <td>{key.name}</td>
                    <td>{key.role}</td>
                    <td>{this.formatDate(key.expiration)}</td>
                    <td>
                      <DeleteButton onConfirm={() => this.onDeleteApiKey(key)} />
                    </td>
                  </tr>
                );
              })}
            </tbody>
          ) : null}
        </table>
      </>
    );
  }

  render() {
    const { hasFetched, navModel, apiKeysCount } = this.props;

    return (
      <Page navModel={navModel}>
        <Page.Contents isLoading={!hasFetched}>
          {hasFetched && (apiKeysCount > 0 ? this.renderApiKeyList() : this.renderEmptyList())}
        </Page.Contents>
      </Page>
    );
  }
}

function mapStateToProps(state: any) {
  return {
    navModel: getNavModel(state.navIndex, 'apikeys'),
    apiKeys: getApiKeys(state.apiKeys),
    searchQuery: state.apiKeys.searchQuery,
    apiKeysCount: getApiKeysCount(state.apiKeys),
    hasFetched: state.apiKeys.hasFetched,
  };
}

const mapDispatchToProps = {
  loadApiKeys,
  deleteApiKey,
  setSearchQuery,
  addApiKey,
};

export default hot(module)(
  connect(
    mapStateToProps,
    mapDispatchToProps
  )(ApiKeysPage)
);
