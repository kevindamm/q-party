# ?-Party command-line interpreter

Initializes the environment for serving from a local DB, building on existing
index tables or creating all the tables for the service to function.  This
includes user tables and session tables, from a hard-coded list of lobby names.

## Example commands

```sh
go build cmd/jarchive
./jarchive -embed-init -write_db=jarchive.sqlite
```

This will use the default data path (`./.data` from cwd) and will write the
embedded index to a sqlite file in that directory.  Further interactions may
require additional fetches, the command takes care of rate-limiting these
fetches to be considerate to the original hosts, and only the bare minimum of
external pages will be fetched to populate a local board or two.

To further reduce the load on the host, the local DB will cache the episode's
content (when -write_db is given a value).  The host must comply with the terms
of service for j-archive.com if persisting the contents of this database.

If making changes to the code, use `-log_debug=debug.log` (with any file path)
to save the verbose log to a file.  Pass "-" as a filename to print to STDOUT.


