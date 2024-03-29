import { renderHook, act } from 'react-hooks-testing-library';
import LanguageProvider from 'app/plugins/datasource/loki/language_provider';
import { useLokiLabels } from './useLokiLabels';
import { AbsoluteTimeRange } from '@smartems/data';
import { makeMockLokiDatasource } from '../mocks';

describe('useLokiLabels hook', () => {
  it('should refresh labels', async () => {
    const datasource = makeMockLokiDatasource({});
    const languageProvider = new LanguageProvider(datasource);
    const logLabelOptionsMock = ['Holy mock!'];
    const rangeMock: AbsoluteTimeRange = {
      from: 1560153109000,
      to: 1560153109000,
    };

    languageProvider.refreshLogLabels = () => {
      languageProvider.logLabelOptions = logLabelOptionsMock;
      return Promise.resolve();
    };

    const { result, waitForNextUpdate } = renderHook(() => useLokiLabels(languageProvider, true, [], rangeMock));
    act(() => result.current.refreshLabels());
    expect(result.current.logLabelOptions).toEqual([]);
    await waitForNextUpdate();
    expect(result.current.logLabelOptions).toEqual(logLabelOptionsMock);
  });
});
