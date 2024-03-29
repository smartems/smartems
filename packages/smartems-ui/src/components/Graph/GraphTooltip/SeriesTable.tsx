import React from 'react';
import { stylesFactory } from '../../../themes/stylesFactory';
import { GrafanaTheme, GraphSeriesValue } from '@smartems/data';
import { css, cx } from 'emotion';
import { SeriesIcon } from '../../Legend/SeriesIcon';
import { useTheme } from '../../../themes';

interface SeriesTableRowProps {
  color?: string;
  label?: string;
  value: string | GraphSeriesValue;
  isActive?: boolean;
}

const getSeriesTableRowStyles = stylesFactory((theme: GrafanaTheme) => {
  return {
    icon: css`
      margin-right: ${theme.spacing.xs};
    `,
    seriesTable: css`
      display: table;
    `,
    seriesTableRow: css`
      display: table-row;
    `,
    seriesTableCell: css`
      display: table-cell;
    `,
    value: css`
      padding-left: ${theme.spacing.md};
    `,
    activeSeries: css`
      font-weight: ${theme.typography.weight.bold};
    `,
  };
});

const SeriesTableRow: React.FC<SeriesTableRowProps> = ({ color, label, value, isActive }) => {
  const theme = useTheme();
  const styles = getSeriesTableRowStyles(theme);
  return (
    <div className={cx(styles.seriesTableRow, isActive && styles.activeSeries)}>
      {color && (
        <div className={styles.seriesTableCell}>
          <SeriesIcon color={color} className={styles.icon} />
        </div>
      )}
      <div className={styles.seriesTableCell}>{label}</div>
      <div className={cx(styles.seriesTableCell, styles.value)}>{value}</div>
    </div>
  );
};

interface SeriesTableProps {
  timestamp?: string | GraphSeriesValue;
  series: SeriesTableRowProps[];
}

export const SeriesTable: React.FC<SeriesTableProps> = ({ timestamp, series }) => {
  return (
    <>
      {timestamp && <div aria-label="Timestamp">{timestamp}</div>}
      {series.map(s => {
        return <SeriesTableRow isActive={s.isActive} label={s.label} color={s.color} value={s.value} key={s.label} />;
      })}
    </>
  );
};
