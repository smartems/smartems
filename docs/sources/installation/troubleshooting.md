+++
title = "Troubleshooting"
description = "Guide to troubleshooting smartEMS problems"
keywords = ["smartems", "troubleshooting", "documentation", "guide"]
type = "docs"
[menu.docs]
parent = "admin"
weight = 8
+++


# Troubleshooting

## Visualization and Query issues

{{< imgbox max-width="40%" img="/img/docs/v45/query_inspector.png" caption="Query Inspector" >}}

The most common problems are related to the query and response from you data source. Even if it looks
like a bug or visualization issue in smartEMS it is 99% of time a problem with the data source query or
the data source response.

To check this you should use Query Inspector (new in smartEMS v4.5). The query Inspector shows query requests and responses.

For more on the query inspector read [this guide here](https://community.smartems.com/t/using-smartemss-query-inspector-to-troubleshoot-issues/2630). For
older versions of smartEMS read the [how troubleshoot metric query issue](https://community.smartems.com/t/how-to-troubleshoot-metric-query-issues/50/2) article.

## Logging

If you encounter an error or problem it is a good idea to check the smartems server log. Usually
located at `/var/log/smartems/smartems.log` on Unix systems or in `<smartems_install_dir>/data/log` on
other platforms and manual installs.

You can enable more logging by changing log level in your smartems configuration file.

## Diagnostics

The `smartems-server` process can be instructued to enable certain diagnostics when it starts. This can be helpful
when experiencing/investigating certain performance problems. It's `not` recommended to have these enabled per default.

### Profiling

The `smartems-server` can be started with the arguments `-profile` to enable profiling and  `-profile-port` to override
the default HTTP port (`6060`) where the pprof debugging endpoints will be available, e.g.

```bash
./smartems-server -profile -profile-port=8080
```

Note that pprof debugging endpoints are served on a different port than the smartEMS HTTP server.

You can configure/override profiling settings using environment variables:

```bash
export GF_DIAGNOSTICS_PROFILING_ENABLED=true
export GF_DIAGNOSTICS_PROFILING_PORT=8080
```

See [Go command pprof](https://golang.org/cmd/pprof/) for more information about how to collect and analyze profiling data.

### Tracing

The `smartems-server` can be started with the arguments `-tracing` to enable tracing and `-tracing-file` to
override the default trace file (`trace.out`) where trace result will be written to, e.g.

```bash
./smartems-server -tracing -tracing-file=/tmp/trace.out
```

You can configure/override profiling settings using environment variables:

```bash
export GF_DIAGNOSTICS_TRACING_ENABLED=true
export GF_DIAGNOSTICS_TRACING_FILE=/tmp/trace.out
```

View the trace in a web browser (Go required to be installed):

```bash
go tool trace <trace file>
2019/11/24 22:20:42 Parsing trace...
2019/11/24 22:20:42 Splitting trace...
2019/11/24 22:20:42 Opening browser. Trace viewer is listening on http://127.0.0.1:39735
```

See [Go command trace](https://golang.org/cmd/trace/) for more information about how to analyze trace files.

## FAQ

Checkout the [FAQ](https://community.smartems.com/c/howto/faq) section on our community page for frequently
asked questions.

