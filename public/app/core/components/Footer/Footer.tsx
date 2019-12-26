import React, { FC } from 'react';
import { Tooltip } from '@smartems/ui';

interface Props {
  appName: string;
  buildVersion: string;
  buildCommit: string;
  newGrafanaVersionExists: boolean;
  newGrafanaVersion: string;
}

export const Footer: FC<Props> = React.memo(
  ({ appName, buildVersion, buildCommit, newGrafanaVersionExists, newGrafanaVersion }) => {
    return (
      <footer className="footer">
        <div className="text-center">
          <ul>
            <li>
              <a href="http://docs.smartems.org" target="_blank" rel="noopener">
                <i className="fa fa-file-code-o" /> Docs
              </a>
            </li>
            <li>
              <a
                href="https://smartems.com/products/enterprise/?utm_source=smartems_footer"
                target="_blank"
                rel="noopener"
              >
                <i className="fa fa-support" /> Support & Enterprise
              </a>
            </li>
            <li>
              <a href="https://community.smartems.com/" target="_blank" rel="noopener">
                <i className="fa fa-comments-o" /> Community
              </a>
            </li>
            <li>
              <a href="https://smartems.com" target="_blank" rel="noopener">
                {appName}
              </a>{' '}
              <span>
                v{buildVersion} (commit: {buildCommit})
              </span>
            </li>
            {newGrafanaVersionExists && (
              <li>
                <Tooltip placement="auto" content={newGrafanaVersion}>
                  <a href="https://smartems.com/get" target="_blank" rel="noopener">
                    New version available!
                  </a>
                </Tooltip>
              </li>
            )}
          </ul>
        </div>
      </footer>
    );
  }
);

export default Footer;
