import React from 'react';
import config from 'app/core/config';

const loginServices: () => LoginServices = () => {
  const oauthEnabled = !!config.oauth;

  return {
    saml: {
      enabled: config.samlEnabled,
      name: 'SAML',
      className: 'github',
      icon: 'key',
    },
    google: {
      enabled: oauthEnabled && config.oauth.google,
      name: 'Google',
    },
    github: {
      enabled: oauthEnabled && config.oauth.github,
      name: 'GitHub',
    },
    gitlab: {
      enabled: oauthEnabled && config.oauth.gitlab,
      name: 'GitLab',
    },
    smartemscom: {
      enabled: oauthEnabled && config.oauth.smartems_com,
      name: 'Grafana.com',
      hrefName: 'smartems_com',
      icon: 'smartems_com',
    },
    oauth: {
      enabled: oauthEnabled && config.oauth.generic_oauth,
      name: oauthEnabled && config.oauth.generic_oauth ? config.oauth.generic_oauth.name : 'OAuth',
      icon: 'sign-in',
      hrefName: 'generic_oauth',
    },
  };
};

export interface LoginService {
  enabled: boolean;
  name: string;
  hrefName?: string;
  icon?: string;
  className?: string;
}

export interface LoginServices {
  [key: string]: LoginService;
}

const LoginDivider = () => {
  return (
    <>
      <div className="text-center login-divider">
        <div>
          <div className="login-divider-line" />
        </div>
        <div>
          <span className="login-divider-text">{config.disableLoginForm ? null : <span>or</span>}</span>
        </div>
        <div>
          <div className="login-divider-line" />
        </div>
      </div>
      <div className="clearfix" />
    </>
  );
};

export const LoginServiceButtons = () => {
  const keyNames = Object.keys(loginServices());
  const serviceElementsEnabled = keyNames.filter(key => {
    const service: LoginService = loginServices()[key];
    return service.enabled;
  });

  if (serviceElementsEnabled.length === 0) {
    return null;
  }

  const serviceElements = serviceElementsEnabled.map(key => {
    const service: LoginService = loginServices()[key];
    return (
      <a
        key={key}
        className={`btn btn-medium btn-service btn-service--${service.className || key} login-btn`}
        href={`login/${service.hrefName ? service.hrefName : key}`}
        target="_self"
      >
        <i className={`btn-service-icon fa fa-${service.icon ? service.icon : key}`} />
        Sign in with {service.name}
      </a>
    );
  });

  const divider = LoginDivider();
  return (
    <>
      {divider}
      <div className="login-oauth text-center">{serviceElements}</div>
    </>
  );
};
