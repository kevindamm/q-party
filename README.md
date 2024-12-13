# ?-Party

A variant on a familiar quiz show game, implemented as an HTMX server in
[golang](https://golang.dev) and [echo](https://echo.labstack.com/),
with [ent](https://entgo.io/) and [sqlite](https://github.com/mattn/go-sqlite3)
for data management.  Uses WebSockets and SSE for realtime communication to the
players, host, and audience.


## Server (private, ask for access)

![TODO img](./app/public/screenshot_01.gif)

## Host Interface

As part of the quiz game interface, a unique board-construction editor is
provided using an inverted-selection mechanic:

1. the host chooses a category from those available (from previous episodes),

2. then a possible answer is entered, with autocomplete & suggestions given by
   the server; if not found, new entries are still accepted.

3. with an existing selection, the associated clue is automatically pulled
   from the database.  The host can edit it (or must provide a clue if it's a new answer entry).

Sometimes an existing clue needs to be updated because it is based on
constraints that are no longer accurate in today's world.  The above process is
one place that this kind of accuracy-check can be done, or it can be raised
during gameplay whenever a challenge for fact-checking comes up.


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

