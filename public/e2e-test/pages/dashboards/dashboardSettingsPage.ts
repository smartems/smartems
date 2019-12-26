import { ClickablePageObject, ClickablePageObjectType, Selector, TestPage } from '@smartems/toolkit/src/e2e';

export interface DashboardSettingsPage {
  deleteDashBoard: ClickablePageObjectType;
}

export const dashboardSettingsPage = new TestPage<DashboardSettingsPage>({
  pageObjects: {
    deleteDashBoard: new ClickablePageObject(Selector.fromAriaLabel('Dashboard settings page delete dashboard button')),
  },
});
