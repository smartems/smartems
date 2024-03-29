import locationUtil from 'app/core/utils/location_util';

jest.mock('app/core/config', () => {
  return {
    appSubUrl: '/subUrl',
  };
});

describe('locationUtil', () => {
  describe('With /subUrl as appSubUrl', () => {
    it('/subUrl should be stripped', () => {
      const urlWithoutMaster = locationUtil.stripBaseFromUrl('/subUrl/smartems/');
      expect(urlWithoutMaster).toBe('/smartems/');
    });
  });
});
