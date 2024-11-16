# ?-Party

A variant on a familiar quiz show game, implemented as an HTMX server in
[golang](https://golang.dev) and [echo](https://echo.labstack.com/),
with [ent](https://entgo.io/) and [sqlite](https://github.com/mattn/go-sqlite3)
for data management.  Uses WebSockets and SSE for realtime communication to the
players, host, and audience.


## Server (private, ask for access)

![TODO img](./app/public/screenshot_01.gif)


## Command Utilities

Some tools written in Go are provided
in the `cmd/` directory of this repo:

<ul>
<li>

**editor** - CLI tool for creating a new board and/or episode, exports as JSON or writes into the challenges database

</li><li>

**jarchive** - fetch a playable game from a single episode of the archive

</li><li>

**server** - Serves HTML and JSON to hypermedia clients

</li>
</ul>

