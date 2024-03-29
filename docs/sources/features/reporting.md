+++
title = "Reporting"
description = ""
keywords = ["smartems", "reporting"]
type = "docs"
aliases = ["/administration/reports"]
[menu.docs]
parent = "features"
weight = 8
+++

# Reporting

> Reporting is only available in smartEMS Enterprise. Read more about [smartEMS Enterprise]({{< relref "enterprise" >}}).

> Only available in smartEMS v6.4+

Reporting allows you to generate PDFs from any of your Dashboards and have them sent out to interested parties on a schedule.

{{< docs-imagebox img="/img/docs/enterprise/reports_list.png" max-width="500px" class="docs-image--no-shadow" >}}

## Dashboard as a Report

With Reports there are a few things to keep in mind, most importantly, any changes you make to the Dashboard used in a report will be reflected in the report. If you change the time range in the Dashboard the time range will be the same in the report as well.

## Setup

> SMTP must be configured for reports to be sent

We recommend using the new image rendering plugin with reporting as it supports a wider range of panels than the built-in image rendering. Read more about it [here]({{< relref "administration/image_rendering.md#smartems-image-renderer-plugin" >}})

## Usage

{{< docs-imagebox img="/img/docs/enterprise/reports_create_new.png" max-width="500px" class="docs-image--no-shadow" >}}

Currently only Organisation Admins can create reports. To get to report click on the reports icon in the side menu. This will allow you to list, create and update your reports.

| Setting       | Description                                                       |
| --------------|------------------------------------------------------------------ |
| Name          | name of the Report                                                |
| Dashboard     | what dashboard to generate the report from                        |
| Recipients    | emails of the people who will receive this report                 |
| ReplyTo       | your email address, so that the recipient can respond             |
| Message       | message body in the email with the report                         |
| Schedule      | how often do you want the report generated and sent               |

## Debugging errors

If you have problems with the reporting feature you can enable debug logging by switching the logger to debug (`filters = report:debug`). Learn more about making configuration changes [here]({{< relref "installation/configuration.md#filters" >}}).
