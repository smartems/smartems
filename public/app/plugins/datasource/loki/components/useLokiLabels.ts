import { useState, useEffect } from 'react';
import { AbsoluteTimeRange } from '@smartems/data';

import LokiLanguageProvider from 'app/plugins/datasource/loki/language_provider';
import { CascaderOption } from 'app/plugins/datasource/loki/components/LokiQueryFieldForm';
import { useRefMounted } from 'app/core/hooks/useRefMounted';

/**
 *
 * @param languageProvider
 * @param languageProviderInitialised
 * @param activeOption rc-cascader provided option used to fetch option's values that hasn't been loaded yet
 *
 * @description Fetches missing labels and enables labels refresh
 */
export const useLokiLabels = (
  languageProvider: LokiLanguageProvider,
  languageProviderInitialised: boolean,
  activeOption: CascaderOption[],
  absoluteRange: AbsoluteTimeRange
) => {
  const mounted = useRefMounted();

  // State
  const [logLabelOptions, setLogLabelOptions] = useState([]);
  const [shouldTryRefreshLabels, setRefreshLabels] = useState(false);
  const [shouldForceRefreshLabels, setForceRefreshLabels] = useState(false);

  // Async
  const fetchOptionValues = async (option: string) => {
    await languageProvider.fetchLabelValues(option, absoluteRange);
    if (mounted.current) {
      setLogLabelOptions(languageProvider.logLabelOptions);
    }
  };

  const tryLabelsRefresh = async () => {
    await languageProvider.refreshLogLabels(absoluteRange, shouldForceRefreshLabels);

    if (mounted.current) {
      setRefreshLabels(false);
      setForceRefreshLabels(false);
      setLogLabelOptions(languageProvider.logLabelOptions);
    }
  };

  // Effects

  // This effect performs loading of options that hasn't been loaded yet
  // It's a subject of activeOption state change only. This is because of specific behavior or rc-cascader
  // https://github.com/react-component/cascader/blob/master/src/Cascader.jsx#L165
  useEffect(() => {
    if (languageProviderInitialised) {
      const targetOption = activeOption[activeOption.length - 1];
      if (targetOption) {
        const nextOptions = logLabelOptions.map(option => {
          if (option.value === targetOption.value) {
            return {
              ...option,
              loading: true,
            };
          }
          return option;
        });
        setLogLabelOptions(nextOptions); // to set loading
        fetchOptionValues(targetOption.value);
      }
    }
  }, [activeOption]);

  // This effect is performed on shouldTryRefreshLabels or shouldForceRefreshLabels state change only.
  // Since shouldTryRefreshLabels is reset AFTER the labels are refreshed we are secured in case of trying to refresh
  // when previous refresh hasn't finished yet
  useEffect(() => {
    if (shouldTryRefreshLabels || shouldForceRefreshLabels) {
      tryLabelsRefresh();
    }
  }, [shouldTryRefreshLabels, shouldForceRefreshLabels]);

  return {
    logLabelOptions,
    setLogLabelOptions,
    refreshLabels: () => setRefreshLabels(true),
  };
};
