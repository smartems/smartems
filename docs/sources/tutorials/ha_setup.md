+++
title = "Setup smartEMS for high availability"
type = "docs"
keywords = ["smartems", "tutorials", "HA", "high availability"]
[menu.docs]
parent = "tutorials"
weight = 10
+++

# How to setup smartEMS for high availability

Setting up smartEMS for high availability is fairly simple. You just need a shared database for storing dashboard, users,
and other persistent data. So the default embedded SQLite database will not work, you will have to switch to
MySQL or Postgres.

<div class="text-center">
  <img src="/img/docs/tutorials/smartems-high-availability.png"  max-width= "800px" class="center" />
</div>

## Configure multiple servers to use the same database

First, you need to do is to setup MySQL or Postgres on another server and configure smartEMS to use that database.
You can find the configuration for doing that in the [[database]]({{< relref "../installation/configuration.md" >}}#database) section in the smartems config.
smartEMS will now persist all long term data in the database. How to configure the database for high availability is out of scope for this guide. We recommend finding an expert on for the database you're using.

## Alerting

Currently alerting supports a limited form of high availability. Since v4.2.0, alert notifications are deduped when running multiple servers. This means all alerts are executed on every server but alert notifications are only sent once per alert. smartEMS does not support load distribution between servers.

## User sessions

> After smartEMS 6.2 you don't need to configure session storage since the database will be used by default.
> If you want to offload the login session data from the database you can configure [remote_cache]({{< relref "../installation/configuration.md" >}}#remote-cache)

The second thing to consider is how to deal with user sessions and how to configure your load balancer in front of smartEMS.
smartEMS supports two ways of storing session data: locally on disk or in a database/cache-server.
If you want to store sessions on disk you can use `sticky sessions` in your load balancer. If you prefer to store session data in a database/cache-server
you can use any stateless routing strategy in your load balancer (ex round robin or least connections).

### Sticky sessions
Using sticky sessions, all traffic for one user will always be sent to the same server. Which means that session related data can be
stored on disk rather than on a shared database. This is the default behavior for smartEMS and if only want multiple servers for fail over this is a good solution since it requires the least amount of work.

### Stateless sessions
You can also choose to store session data in a Redis/Memcache/Postgres/MySQL which means that the load balancer can send a user to any smartEMS server without having to log in on each server. This requires a little bit more work from the operator but enables you to remove/add smartems servers without impacting the user experience.
If you use MySQL/Postgres for session storage, you first need a table to store the session data in. More details about that in [[sessions]]({{< relref "../installation/configuration.md" >}}#session)

For smartEMS itself it doesn't really matter if you store the session data on disk or database/redis/memcache. But we recommend using a database/redis/memcache since it makes it easier manage the smartems servers.


