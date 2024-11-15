# ?-Party

A variant on a familiar quiz show game, implemented as an HTMX server in
[golang](https://golang.dev) and [echo](https://echo.labstack.com/),
with [ent](https://entgo.io/) and [sqlite](https://github.com/mattn/go-sqlite3)
for data management.  Uses WebSockets and SSE for realtime communication to the
players and audience.


## Server (private, ask for access)

![TODO img](./app/public/screenshot_01.gif)

## Building

The server is built from `/cmd/server` and the client is presently HTMX so it
is built within the server.

TODO there is room for a SPA interface, either for a higher-fidelity host, for an editor interface in the browser, and/or for more dynamic interfaces (hexagonal baord, social category selection, etc.).

Contact me if you're interested in contributing!

## Commands

Some tools written in Go are provided
in the `cmd/` directory of this repo:

<ul>
<li>

**editor** - CLI tool for creating a new board and/or episode, exports as JSON or writes into the challenges database.

</li><li>

**jarchive** - fetch a playable game from a single episode of the archive

</li><li>

**server** - Serves HTML and JSON to hypermedia clients (tested on Chrome).

</li>
</ul>

