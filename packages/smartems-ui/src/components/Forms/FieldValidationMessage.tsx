import React from 'react';
import { useTheme, stylesFactory } from '../../themes';
import { GrafanaTheme } from '@smartems/data';
import { css, cx } from 'emotion';

export interface FieldValidationMessageProps {
  children: string;
  className?: string;
}

export const getFieldValidationMessageStyles = stylesFactory((theme: GrafanaTheme) => {
  return {
    fieldValidationMessage: css`
      font-size: ${theme.typography.size.sm};
      font-weight: ${theme.typography.weight.semibold};
      margin: ${theme.spacing.formLabelMargin};
      padding: ${theme.spacing.formValidationMessagePadding};
      color: ${theme.colors.formValidationMessageText};
      background: ${theme.colors.formValidationMessageBg};
      border-radius: ${theme.border.radius.sm};
      position: relative;

      &:before {
        content: '';
        position: absolute;
        left: 9px;
        top: -5px;
        width: 0;
        height: 0;
        border-left: 5px solid transparent;
        border-right: 5px solid transparent;
        border-bottom: 5px solid ${theme.colors.formValidationMessageBg};
      }
    `,
    fieldValidationMessageIcon: css`
      margin-right: ${theme.spacing.formSpacingBase}px;
    `,
  };
});

export const FieldValidationMessage: React.FC<FieldValidationMessageProps> = ({ children, className }) => {
  const theme = useTheme();
  const styles = getFieldValidationMessageStyles(theme);

  return (
    <div className={cx(styles.fieldValidationMessage, className)}>
      <i className={cx(styles.fieldValidationMessageIcon, 'fa', 'fa-warning')} />
      {children}
    </div>
  );
};
