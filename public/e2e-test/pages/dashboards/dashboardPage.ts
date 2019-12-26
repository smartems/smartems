import { ClickablePageObject, ClickablePageObjectType, Selector, TestPage } from '@smartems/toolkit/src/e2e';

export interface DashboardPage {
  settings: ClickablePageObjectType;
}

export const dashboardPage = new TestPage<DashboardPage>({
  pageObjects: {
    settings: new ClickablePageObject(Selector.fromAriaLabel('Dashboard settings navbar button')),
  },
});
