+++
title = "Alerting Engine and Rules Guide"
description = "Configuring Alert Rules"
keywords = ["smartems", "alerting", "guide", "rules"]
type = "docs"
[menu.docs]
name = "Engine and Rules"
parent = "alerting"
weight = 1
+++

# Alerting Engine and Rules Guide

> Alerting is only available in smartEMS v4.0 and above.

## Introduction

{{< imgbox max-width="40%" img="/img/docs/v4/drag_handles_gif.gif" caption="Alerting overview" >}}

Alerting in smartEMS allows you to attach rules to your dashboard panels. When you save the dashboard
smartEMS will extract the alert rules into a separate alert rule storage and schedule them for evaluation.

In the alert tab of the graph panel you can configure how often the alert rule should be evaluated
and the conditions that need to be met for the alert to change state and trigger its
[notifications]({{< relref "notifications.md" >}}).

## Execution

The alert rules are evaluated in the smartEMS backend in a scheduler and query execution engine that is part
of core smartEMS. Only some data sources are supported right now. They include `Graphite`, `Prometheus`, `InfluxDB`, `Elasticsearch`,
`Stackdriver`, `Cloudwatch`, `Azure Monitor`, `MySQL`, `PostgreSQL`, `MSSQL`, `OpenTSDB`, `Oracle` and `Azure Data Explorer`.

> Alerting support for Azure Monitor is only available in smartEMS v6.0 and above.

### Clustering

Currently alerting supports a limited form of high availability. Since v4.2.0 of smartEMS, alert notifications are deduped when running multiple servers. This means all alerts are executed on every server but no duplicate alert notifications are sent due to the deduping logic. Proper load balancing of alerts will be introduced in the future.

<div class="clearfix"></div>

## Rule Config



Currently only the graph panel supports alert rules but this will be added to the **Singlestat** and **Table**
panels as well in a future release.

### Name and Evaluation interval

Here you can specify the name of the alert rule and how often the scheduler should evaluate the alert rule.

### For

> **Important note regarding No Data:**
>
> Do not use `For` with the `If no data or all values are null` setting set to `No Data`. The triggering of `No Data` will trigger instantly and not take `For` into consideration. This may also result in that an OK notification not being sent if alert transitions from `No Data -> Pending -> OK`.

If an alert rule has a configured `For` and the query violates the configured threshold it will first go from `OK` to `Pending`. Going from `OK` to `Pending` smartEMS will not send any notifications. Once the alert rule has been firing for more than `For` duration, it will change to `Alerting` and send alert notifications.

Typically, it's always a good idea to use this setting since it's often worse to get false positive than wait a few minutes before the alert notification triggers. Looking at the `Alert list` or `Alert list panels` you will be able to see alerts in pending state.

Below you can see an example timeline of an alert using the `For` setting. At ~16:04 the alert state changes to `Pending` and after 4 minutes it changes to `Alerting` which is when alert notifications are sent. Once the series falls back to normal the alert rule goes back to `OK`.
{{< imgbox img="/img/docs/v54/alerting-for-dark-theme.png" caption="Alerting For" >}}

{{< imgbox max-width="40%" img="/img/docs/v4/alerting_conditions.png" caption="Alerting Conditions" >}}

### Conditions

Currently the only condition type that exists is a `Query` condition that allows you to
specify a query letter, time range and an aggregation function.


### Query condition example

```sql
avg() OF query(A, 15m, now) IS BELOW 14
```

- `avg()` Controls how the values for **each** series should be reduced to a value that can be compared against the threshold. Click on the function to change it to another aggregation function.
- `query(A, 15m, now)`  The letter defines what query to execute from the **Metrics** tab. The second two parameters define the time range, `15m, now` means 15 minutes ago to now. You can also do `10m, now-2m` to define a time range that will be 10 minutes ago to 2 minutes ago. This is useful if you want to ignore the last 2 minutes of data.
- `IS BELOW 14`  Defines the type of threshold and the threshold value.  You can click on `IS BELOW` to change the type of threshold.

The query used in an alert rule cannot contain any template variables. Currently we only support `AND` and `OR` operators between conditions and they are executed serially.
For example, we have 3 conditions in the following order:
*condition:A(evaluates to: TRUE) OR condition:B(evaluates to: FALSE) AND condition:C(evaluates to: TRUE)*
so the result will be calculated as ((TRUE OR FALSE) AND TRUE) = TRUE.

We plan to add other condition types in the future, like `Other Alert`, where you can include the state
of another alert in your conditions, and `Time Of Day`.

#### Multiple Series

If a query returns multiple series then the aggregation function and threshold check will be evaluated for each series.
What smartEMS does not do currently is track alert rule state **per series**. This has implications that are detailed
in the scenario below.

- Alert condition with query that returns 2 series: **server1** and **server2**
- **server1** series cause the alert rule to fire and switch to state `Alerting`
- Notifications are sent out with message:  _load peaking (server1)_
- In a subsequence evaluation of the same alert rule the **server2** series also cause the alert rule to fire
- No new notifications are sent as the alert rule is already in state `Alerting`.

So as you can see from the above scenario smartEMS will not send out notifications when other series cause the alert
to fire if the rule already is in state `Alerting`. To improve support for queries that return multiple series
we plan to track state **per series** in a future release.

> Starting with smartEMS v5.3 you can configure reminders to be sent for triggered alerts. This will send additional notifications
> when an alert continues to fire. If other series (like server2 in the example above) also cause the alert rule to fire they will
> be included in the reminder notification. Depending on what notification channel you're using you may be able to take advantage
> of this feature for identifying new/existing series causing alert to fire. [Read more about notification reminders here](/alerting/notifications/#send-reminders).

### No Data / Null values

Below your conditions you can configure how the rule evaluation engine should handle queries that return no data or only null values.

No Data Option | Description
------------ | -------------
NoData | Set alert rule state to `NoData`
Alerting | Set alert rule state to `Alerting`
Keep Last State | Keep the current alert rule state, what ever it is.

### Execution errors or timeouts

The last option tells how to handle execution or timeout errors.

Error or timeout option | Description
------------ | -------------
Alerting | Set alert rule state to `Alerting`
Keep Last State | Keep the current alert rule state, what ever it is.

If you have an unreliable time series store from which queries sometime timeout or fail randomly you can set this option
to `Keep Last State` in order to basically ignore them.

## Notifications

In alert tab you can also specify alert rule notifications along with a detailed message about the alert rule.
The message can contain anything, information about how you might solve the issue, link to runbook, etc.

The actual notifications are configured and shared between multiple alerts. Read the
[notifications]({{< relref "notifications.md" >}}) guide for how to configure and setup notifications.

## Alert State History and Annotations

Alert state changes are recorded in the internal annotation table in smartEMS's database. The state changes
are visualized as annotations in the alert rule's graph panel. You can also go into the `State history`
submenu in the alert tab to view and clear state history.

## Troubleshooting

{{< imgbox max-width="40%" img="/img/docs/v4/alert_test_rule.png" caption="Test Rule" >}}

First level of troubleshooting you can do is hit the **Test Rule** button. You will get result back that you can expand
to the point where you can see the raw data that was returned from your query.

Further troubleshooting can also be done by inspecting the smartems-server log. If it's not an error or for some reason
the log does not say anything you can enable debug logging for some relevant components. This is done
in smartEMS's ini config file.

Example showing loggers that could be relevant when troubleshooting alerting.

```ini
[log]
filters = alerting.scheduler:debug \
          alerting.engine:debug \
          alerting.resultHandler:debug \
          alerting.evalHandler:debug \
          alerting.evalContext:debug \
          alerting.extractor:debug \
          alerting.notifier:debug \
          alerting.notifier.slack:debug \
          alerting.notifier.pagerduty:debug \
          alerting.notifier.email:debug \
          alerting.notifier.webhook:debug \
          tsdb.graphite:debug \
          tsdb.prometheus:debug \
          tsdb.opentsdb:debug \
          tsdb.influxdb:debug \
          tsdb.elasticsearch:debug \
          tsdb.elasticsearch.client:debug \
```

If you want to log raw query sent to your TSDB and raw response in log you also have to set smartems.ini option `app_mode` to
`development`.
