+++
title = "How to integrate Hubot and smartEMS"
type = "docs"
keywords = ["smartems", "tutorials", "hubot", "slack", "hipchat", "setup", "install", "config"]
[menu.docs]
parent = "tutorials"
weight = 10
+++

# How to integrate Hubot with smartEMS

smartEMS 2.0 shipped with a great feature that enables it to render any graph or panel to a PNG image.
No matter what data source you are using, the PNG image of the Graph will look the same
as it does in your browser.

This guide will show you how to install and configure the [Hubot-smartEMS](https://github.com/stephenyeargin/hubot-smartems)
plugin. This plugin allows you to tell hubot to render any dashboard or graph right from a channel in
Slack, Hipchat or Basecamp. The bot will respond with an image of the graph and a link that will
take you to the graph.

> *Amazon S3 Required*: The hubot-smartems script will upload the rendered graphs to Amazon S3. This
> is so Hipchat and Slack can show them reliably (they require the image to be publicly available).

<div class="text-center">
  <img src="/img/docs/tutorials/hubot_smartems.png" class="center"></a>
</div>

## What is Hubot?

[Hubot](https://hubot.github.com/) is an universal and extensible chat bot that can be used with many chat
services and has a huge library of third party plugins that allow you to automate anything from your
chat rooms.

## Install Hubot

Hubot is very easy to install and host. If you do not already have a bot up and running please
read the official [Getting Started With Hubot](https://hubot.github.com/docs/) guide.

## Install Hubot-smartEMS script

In your Hubot project repo install the smartEMS plugin using `npm`:
```bash
npm install hubot-smartems --save
```
Edit the file external-scripts.json, and add hubot-smartems to the list of plugins.

```json
[
"hubot-pugme",
"hubot-shipit",
"hubot-smartems"
]
```

## Configure

The `hubot-smartems` plugin requires a number of environment variables to be set in order to work properly.

```bash
export HUBOT_SMARTEMS_HOST=http://play.smartems.org
export HUBOT_SMARTEMS_API_KEY=abcd01234deadbeef01234
export HUBOT_SMARTEMS_S3_BUCKET=mybucket
export HUBOT_SMARTEMS_S3_ACCESS_KEY_ID=ABCDEF123456XYZ
export HUBOT_SMARTEMS_S3_SECRET_ACCESS_KEY=aBcD01234dEaDbEef01234
export HUBOT_SMARTEMS_S3_PREFIX=graphs
export HUBOT_SMARTEMS_S3_REGION=us-standard
```

### smartEMS server side rendering

The hubot plugin will take advantage of the smartEMS server side rendering feature that can
render any panel on the server using phantomjs. smartEMS ships with a phantomjs binary (Linux only).

To verify that this feature works try the `Direct link to rendered image` link in the panel share dialog.
If you do not get an image when opening this link verify that the required font packages are installed for phantomjs to work.

### smartEMS API Key

{{< docs-imagebox img="/img/docs/v2/orgdropdown_api_keys.png" max-width="150px" class="docs-image--right">}}

You need to set the environment variable `HUBOT_SMARTEMS_API_KEY` to a smartEMS API Key.
You can add these from the API Keys page which you find in the Organization dropdown.

### Amazon S3

The `S3` options are optional but for the images to work properly in services like Slack and Hipchat they need
to publicly available. By specifying the `S3` options the hubot-smartems script will publish the rendered
panel to `S3` and it will use that URL when it posts to Slack or Hipchat.

## Hubot commands

- `hubot graf list`
    - Lists the available dashboards
- `hubot graf db graphite-carbon-metrics`
    - Graph all panels in the dashboard
- `hubot graf db graphite-carbon-metrics:3`
    - Graph only panel with id 3 of a particular dashboard
- `hubot graf db graphite-carbon-metrics:cpu`
    - Graph only the panels containing "cpu" (case insensitive) in the title
- `hubot graf db graphite-carbon-metrics now-12hr`
    - Get a dashboard with a window of 12 hours ago to now
- `hubot graf db graphite-carbon-metrics now-24hr now-12hr`
    - Get a dashboard with a window of 24 hours ago to 12 hours ago
- `hubot graf db graphite-carbon-metrics:3 now-8d now-1d`
    - Get only the third panel of a particular dashboard with a window of 8 days ago to yesterday
- `hubot graf db graphite-carbon-metrics host=carbon-a`
    - Get a templated dashboard with the `$host` parameter set to `carbon-a`

## Aliases

Some of the hubot commands above can lengthy and you might have to remember the dashboard slug (url id).
If you have a few favorite graphs you want to be able check up on often (let's say from your mobile) you
can create hubot command aliases with the hubot script `hubot-alias`.

Install it:

```bash
npm i --save hubot-alias
```

Now add `hubot-alias` to the list of plugins in `external-scripts.json` and restart hubot.

Now you can add an alias like this:

- `hubot alias graf-lb=graf db loadbalancers:2 now-20m`

<div class="text-center">
  Using the alias:<br>
  <img src="/img/docs/tutorials/hubot_smartems2.png" class="center"></a>
</div>

## Summary

smartEMS is going to ship with integrated Slack and Hipchat features some day but you do
not have to wait for that. smartEMS 2 shipped with a very clever server side rendering feature
that can render any panel to a png using phantomjs. The hubot plugin for smartEMS is something
you can install and use today!


